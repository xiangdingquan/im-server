package server

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/go-kratos/kratos/pkg/conf/env"
	"strings"
	"sync"

	"open.chat/app/interface/session/sessionpb"
	"open.chat/mtproto"
	"open.chat/pkg/log"
	"open.chat/pkg/net2"
	"open.chat/pkg/time2"
	"open.chat/pkg/util"
)

type HandshakeStateCtx struct {
	State         int32
	ResState      int32
	Nonce         []byte
	ServerNonce   []byte
	NewNonce      []byte
	A             []byte
	P             []byte
	handshakeType int
	ExpiresIn     int32
}

func (m *HandshakeStateCtx) DebugString() string {
	return fmt.Sprintf(`{"state":%d,"res_state":%d,"nonce":"%s","server_nonce":"%s"}`,
		m.State,
		m.ResState,
		hex.EncodeToString(m.Nonce),
		hex.EncodeToString(m.ServerNonce))
}

type connContext struct {
	sync.Mutex
	state      int
	authKeys   []*authKeyUtil
	sessionId  int64
	isHttp     bool
	canSend    bool
	trd        *time2.TimerData
	handshakes []*HandshakeStateCtx
	clientIp   string
}

func newConnContext(clientIp string) *connContext {
	return &connContext{
		state:    STATE_CONNECTED2,
		clientIp: clientIp,
	}
}

func (ctx *connContext) getState() int {
	ctx.Lock()
	defer ctx.Unlock()
	return ctx.state
}

func (ctx *connContext) setState(state int) {
	ctx.Lock()
	defer ctx.Unlock()
	if ctx.state != state {
		ctx.state = state
	}
}

func (ctx *connContext) getAuthKey(id int64) *authKeyUtil {
	ctx.Lock()
	defer ctx.Unlock()
	for _, key := range ctx.authKeys {
		if key.AuthKeyId() == id {
			return key
		}
	}

	return nil
}

func (ctx *connContext) putAuthKey(k *authKeyUtil) {
	ctx.Lock()
	defer ctx.Unlock()
	for _, key := range ctx.authKeys {
		if key.Equal(k) {
			return
		}
	}

	ctx.authKeys = append(ctx.authKeys, k)
}

func (ctx *connContext) getAllAuthKeyId() (idList []int64) {
	ctx.Lock()
	defer ctx.Unlock()

	idList = make([]int64, len(ctx.authKeys))
	for i, key := range ctx.authKeys {
		idList[i] = key.AuthKeyId()
	}

	return
}

func (ctx *connContext) getHandshakeStateCtx(nonce []byte) *HandshakeStateCtx {
	ctx.Lock()
	defer ctx.Unlock()

	for _, state := range ctx.handshakes {
		if bytes.Equal(nonce, state.Nonce) {
			return state
		}
	}

	return nil
}

func (ctx *connContext) putHandshakeStateCt(state *HandshakeStateCtx) {
	ctx.Lock()
	defer ctx.Unlock()

	ctx.handshakes = append(ctx.handshakes, state)
}

func (ctx *connContext) encryptedMessageAble() bool {
	ctx.Lock()
	defer ctx.Unlock()
	return true
}

func (ctx *connContext) DebugString() string {
	s := make([]string, 0, 4)
	s = append(s, fmt.Sprintf(`"state":%d`, ctx.state))
	return "{" + strings.Join(s, ",") + "}"
}

func (s *Server) OnNewConnection(conn *net2.TcpConnection) {
	addr := conn.RemoteAddr().String()
	ctx := newConnContext(strings.Split(addr, ":")[0])

	log.Debugf("onNewConnection - {peer: %s, ctx: {%s}}", conn, ctx.DebugString())
	conn.Context = ctx
}

func (s *Server) OnConnectionDataArrived(conn *net2.TcpConnection, msg interface{}) error {
	msg2, ok := msg.(*mtproto.MTPRawMessage)
	if !ok {
		err := fmt.Errorf("recv invalid MTPRawMessage: {peer: %s, msg: %v", conn, msg2)
		log.Error(err.Error())
		return err
	}

	ctx, _ := conn.Context.(*connContext)

	log.Debugf("onConnectionDataArrived - receive data: {peer: %s, ctx: %s, msg: %s}", conn, ctx.DebugString(), msg2.DebugString())

	if msg2.ConnType() == mtproto.TRANSPORT_HTTP {
		ctx.isHttp = true
	}

	var err error
	if msg2.AuthKeyId() == 0 {
		err = s.onUnencryptedMessage(ctx, conn, msg2)
	} else {
		authKey := ctx.getAuthKey(msg2.AuthKeyId())
		if authKey == nil {
			key, _ := s.dao.GetAuthKey(context.Background(), msg2.AuthKeyId())
			if key == nil {
				err = fmt.Errorf("invalid auth_key_id: {%d}", msg2.AuthKeyId())
				log.Error("invalid auth_key_id: {%v} - {peer: %s, ctx: %s, msg: %s}", err, conn, ctx.DebugString(), msg2.DebugString())
				var code = int32(-404)
				cData := make([]byte, 4)
				binary.LittleEndian.PutUint32(cData, uint32(code))
				conn.Send(&mtproto.MTPRawMessage{Payload: cData})
				return err
			}
			authKey = newAuthKeyUtil(key)
			ctx.putAuthKey(authKey)
		}

		err = s.onEncryptedMessage(ctx, conn, authKey, msg2)
	}

	return err
}

func (s *Server) OnConnectionClosed(conn *net2.TcpConnection) {
	ctx, _ := conn.Context.(*connContext)
	log.Info("onServerConnectionClosed - {peer:%s, ctx:%s}", conn, ctx.DebugString())

	if ctx.trd != nil {
		s.timer.Del(ctx.trd)
		ctx.trd = nil
	}

	sessId, connId := ctx.sessionId, conn.GetConnID()
	for _, id := range ctx.getAllAuthKeyId() {
		bDeleted := s.authSessionMgr.RemoveSession(id, sessId, connId)
		if bDeleted {
			s.sendClientClosed(conn, id, sessId)
			log.Debugf("onServerConnectionClosed - sendClientClosed: {peer:%s, ctx:%s}", conn, ctx.DebugString())
		}
	}
}

func (s *Server) onUnencryptedMessage(ctx *connContext, conn *net2.TcpConnection, mmsg *mtproto.MTPRawMessage) error {
	log.Info("receive unencryptedRawMessage: {peer: %s, ctx: %s, mmsg: %s}", conn, ctx.DebugString(), mmsg.DebugString())

	if len(mmsg.Payload) < 8 {
		err := fmt.Errorf("invalid data len < 8")
		log.Error(err.Error())
		return err
	}

	_, obj, err := ParseFromIncomingMessage(mmsg.Payload[8:])
	if err != nil {
		err := fmt.Errorf("invalid data len < 8")
		log.Errorf(err.Error())
	}

	var rData []byte

	switch request := obj.(type) {
	case *mtproto.TLReqPq:
		resPQ, err := s.handshake.onReqPq(request)
		if err != nil {
			log.Errorf("onHandshake error: {%v} - {peer: %s, ctx: %s, mmsg: %s}", err, conn, ctx.DebugString(), mmsg.DebugString())
			conn.Close()
			return err
		}
		ctx.putHandshakeStateCt(&HandshakeStateCtx{
			State:       STATE_pq_res,
			Nonce:       resPQ.GetNonce(),
			ServerNonce: resPQ.GetServerNonce(),
		})
		rData = SerializeToBuffer(mtproto.GenerateMessageId(), resPQ)
	case *mtproto.TLReqPqMulti:
		resPQ, err := s.handshake.onReqPqMulti(request)
		if err != nil {
			log.Errorf("onHandshake error: {%v} - {peer: %s, ctx: %s, mmsg: %s}", err, conn, ctx.DebugString(), mmsg.DebugString())
			conn.Close()
			return err
		}
		ctx.putHandshakeStateCt(&HandshakeStateCtx{
			State:       STATE_pq_res,
			Nonce:       resPQ.GetNonce(),
			ServerNonce: resPQ.GetServerNonce(),
		})
		rData = SerializeToBuffer(mtproto.GenerateMessageId(), resPQ)
	case *mtproto.TLReq_DHParams:
		if state := ctx.getHandshakeStateCtx(request.Nonce); state != nil {
			resServerDHParam, err := s.handshake.onReqDHParams(state, obj.(*mtproto.TLReq_DHParams))
			if err != nil {
				log.Errorf("onHandshake error: {%v} - {peer: %s, ctx: %s, mmsg: %s}", err, conn, ctx.DebugString(), mmsg.DebugString())
				conn.Close()
				return err
			}
			state.State = STATE_DH_params_res
			rData = SerializeToBuffer(mtproto.GenerateMessageId(), resServerDHParam)
		} else {
			log.Errorf("onHandshake error: {invalid nonce} - {peer: %s, ctx: %s, mmsg: %s}", conn, ctx.DebugString(), mmsg.DebugString())
			return conn.Close()
		}
	case *mtproto.TLSetClient_DHParams:
		if state := ctx.getHandshakeStateCtx(request.Nonce); state != nil {
			resSetClientDHParamsAnswer, err := s.handshake.onSetClientDHParams(state, obj.(*mtproto.TLSetClient_DHParams))
			if err != nil {
				log.Errorf("onHandshake error: {%v} - {peer: %s, ctx: %s, mmsg: %s}", err, conn, ctx.DebugString(), mmsg.DebugString())
				return conn.Close()
			}
			state.State = STATE_dh_gen_res
			rData = SerializeToBuffer(mtproto.GenerateMessageId(), resSetClientDHParamsAnswer)
		} else {
			log.Errorf("onHandshake error: {invalid nonce} - {peer: %s, ctx: %s, mmsg: %s}", conn, ctx.DebugString(), mmsg.DebugString())
			return conn.Close()
		}
	case *mtproto.TLMsgsAck:
		return nil
	default:
		err = fmt.Errorf("invalid handshake type")
		return conn.Close()
	}
	return conn.Send(&mtproto.MTPRawMessage{Payload: rData})
}

func (s *Server) onEncryptedMessage(ctx *connContext, conn *net2.TcpConnection, authKey *authKeyUtil, mmsg *mtproto.MTPRawMessage) error {
	log.Info("receive encryptedRawMessage - receive data: {peer: %s, ctx: %s, msg: %s}", conn, ctx.DebugString(), mmsg.DebugString())

	mtpRwaData, err := authKey.AesIgeDecrypt(mmsg.Payload[8:8+16], mmsg.Payload[24:])
	if err != nil {
		log.Errorf("decrypt error: {%v} - {peer: %s, ctx: %s, mmsg: %s}", err, conn, ctx.DebugString(), mmsg.DebugString())
		conn.Close()
		return err
	}

	sessionId := int64(binary.LittleEndian.Uint64(mtpRwaData[8:]))
	if ctx.sessionId == 0 {
		if s.authSessionMgr.AddNewSession(authKey.AuthKeyId(), sessionId, conn.IsWebsocket(), conn.GetConnID()) {
			s.sendClientNew(conn, authKey.AuthKeyId(), sessionId)
		}
		ctx.sessionId = sessionId
	} else {
		if ctx.sessionId != sessionId {
			log.Warnf("%d != %d - {peer: %s, ctx: %s, mmsg: %s}", ctx.sessionId, sessionId, conn, ctx.DebugString(), mmsg.DebugString())
			if s.authSessionMgr.AddNewSession(authKey.AuthKeyId(), sessionId, conn.IsWebsocket(), conn.GetConnID()) {
				s.sendClientNew(conn, authKey.AuthKeyId(), sessionId)
			}
			ctx.sessionId = sessionId
		}
	}

	c, err := s.GetSessionClient(util.Int64ToString(mmsg.AuthKeyId()))
	if err != nil {
		log.Errorf("getSessionClient error: {%v} - {peer: %s, ctx: %s, mmsg: %s}", err, conn, ctx.DebugString(), mmsg.DebugString())
		return err
	}

	cSessData := &sessionpb.SessionClientData{
		ServerId:  env.Hostname,
		AuthKeyId: authKey.AuthKeyId(),
		SessionId: sessionId,
		Salt:      int64(binary.LittleEndian.Uint64(mtpRwaData)),
		Payload:   mtpRwaData[16:],
		ClientIp:  ctx.clientIp,
	}

	if ctx.isHttp {
		ctx.canSend = true
		rData, err := c.SendSyncDataToSession(context.Background(), cSessData)
		if err != nil {
			log.Warnf("recv error: %v", err)
			conn.Close()
		} else {
			s.SendToClient(conn, authKey, rData.Payload)
		}
	} else {
		c.SendAsyncDataToSession(context.Background(), cSessData)
	}

	return nil
}

func (s *Server) sendClientNew(conn *net2.TcpConnection, authKeyId, sessionId int64) {
	ctx, _ := conn.Context.(*connContext)
	if ctx == nil {
		log.Warnf("sendClientNew: ctx is nil - {peer: %s}", conn)
		return
	}

	c, err := s.GetSessionClient(util.Int64ToString(authKeyId))
	if err != nil {
		log.Errorf("getSessionClient error: {%v} - {peer: %s, authKeyId: %d}", err, conn, authKeyId)
		return
	}

	c.CreateSession(context.Background(), &sessionpb.SessionClientEvent{
		ServerId:  env.Hostname,
		AuthKeyId: authKeyId,
		SessionId: sessionId,
		ClientIp:  ctx.clientIp,
	})
}

func (s *Server) sendClientClosed(conn *net2.TcpConnection, authKeyId, sessionId int64) {
	ctx, _ := conn.Context.(*connContext)
	if ctx == nil {
		log.Warnf("sendClientClosed: ctx is nil - {peer: %s}", conn)
		return
	}

	c, err := s.GetSessionClient(util.Int64ToString(authKeyId))
	if err != nil {
		log.Errorf("getSessionClient error: {%v} - {peer: %s, authKeyId: %d}", err, conn, authKeyId)
		return
	}

	c.CloseSession(context.Background(), &sessionpb.SessionClientEvent{
		ServerId:  env.Hostname,
		AuthKeyId: authKeyId,
		SessionId: sessionId,
		ClientIp:  ctx.clientIp,
	})
}

func (s *Server) GetConnByConnID(id uint64) *net2.TcpConnection {
	return s.server.GetConnection(id)
}

func (s *Server) SendToClient(conn *net2.TcpConnection, authKey *authKeyUtil, b []byte) error {
	ctx, _ := conn.Context.(*connContext)
	if ctx.trd != nil {
		log.Info("del conn timeout")
		s.timer.Del(ctx.trd)
		ctx.trd = nil
	}

	msgKey, mtpRawData, _ := authKey.AesIgeEncrypt(b)
	x := mtproto.NewEncodeBuf(8 + len(msgKey) + len(mtpRawData))
	x.Long(authKey.AuthKeyId())
	x.Bytes(msgKey)
	x.Bytes(mtpRawData)
	msg := &mtproto.MTPRawMessage{Payload: x.GetBuf()}
	if ctx.isHttp {
		ctx.canSend = false
	}

	err := conn.Send(msg)
	if err != nil {
		log.Errorf("send error: %v", err)
		return err
	}

	return nil
}
