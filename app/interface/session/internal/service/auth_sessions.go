package service

import (
	"context"
	"fmt"
	"math"
	"reflect"
	"runtime/debug"
	"sync"
	"time"

	"github.com/go-kratos/kratos/pkg/conf/env"

	status_client "open.chat/app/service/status/client"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
	"open.chat/pkg/queue2"
	"open.chat/pkg/sync2"
)

const (
	clientUnknown  = 0
	clientAndroid  = 1
	clientiOS      = 2
	clientTdesktop = 3
	clientMacSwift = 4
	clientWebogram = 5
)

type rpcApiMessage struct {
	sessionId int64
	clientIp  string
	reqMsgId  int64
	reqMsg    mtproto.TLObject
	rpcResult *mtproto.TLRpcResult
}

func (m *rpcApiMessage) DebugString() string {
	if m.rpcResult == nil {
		return fmt.Sprintf("{session_id: %d, req_msg_id: %d, req_msg: %s}",
			m.sessionId,
			m.reqMsgId,
			logger.JsonDebugData(m.reqMsg))
	} else {
		return fmt.Sprintf("{session_id: %d, req_msg_id: %d, req_msg: %s, rpc_result: %s}",
			m.sessionId,
			m.reqMsgId,
			logger.JsonDebugData(m.reqMsg),
			logger.JsonDebugData(m.rpcResult.Result))
	}
}

type sessionData struct {
	gatewayId string
	clientIp  string
	sessionId int64
	salt      int64
	buf       []byte
}

type sessionHttpData struct {
	gatewayId  string
	clientIp   string
	sessionId  int64
	salt       int64
	buf        []byte
	resChannel chan interface{}
}

type syncRpcResultData struct {
	sessionId   int64
	clientMsgId int64
	data        []byte
}

type syncSessionData struct {
	sessionId int64
	data      *messageData
}

type syncData struct {
	needAndroidPush bool
	data            *messageData
}

func makeSyncData(needAndroidPush bool, data *messageData) *syncData {
	return &syncData{
		needAndroidPush: needAndroidPush,
		data:            data,
	}
}

type connData struct {
	isNew     bool
	gatewayId string
	sessionId int64
}

const (
	keyIdNew     = 0
	keyLoaded    = 1
	unauthorized = 2
	userIdLoaded = 3
	offline      = 4
	closed       = 5
	unknownError = 6
)

type authSessionsCallback interface {
	SendDataToGate(ctx context.Context, serverId int32, authKeyId, sessionId int64, payload []byte) error
}

type authSessions struct {
	authKeyId       int64
	Layer           int32
	IpAddress       string
	Client          string
	AuthUserId      int32
	cacheSalt       *mtproto.TLFutureSalt
	cacheLastSalt   *mtproto.TLFutureSalt
	pushSessionId   int64
	sessions        map[int64]*session
	closeChan       chan struct{}
	sessionDataChan chan interface{}
	rpcDataChan     chan interface{}
	rpcQueue        *queue2.SyncQueue
	finish          sync.WaitGroup
	running         sync2.AtomicInt32
	state           int
	onlineExpired   int64
	clientType      int
	nextNotifyId    int64
	nextPushId      int64
	*Service
}

func newAuthSessions(authKeyId int64, s2 *Service) (*authSessions, error) {
	keyData, err := s2.GetKeyStateData(context.Background(), authKeyId)
	if err != nil {
		log.Errorf("getKeyStateData error: %v", err)
		return nil, err
	}

	s := &authSessions{
		authKeyId:       authKeyId,
		Layer:           keyData.Layer,
		AuthUserId:      keyData.UserId,
		sessions:        make(map[int64]*session),
		closeChan:       make(chan struct{}),
		sessionDataChan: make(chan interface{}, 1024),
		rpcDataChan:     make(chan interface{}, 1024),
		rpcQueue:        queue2.NewSyncQueue(),
		finish:          sync.WaitGroup{},
		clientType:      int(keyData.ClientType),
		nextPushId:      0,
		nextNotifyId:    math.MaxInt32,
		Service:         s2,
	}

	s.Start()
	return s, err
}

func (s *authSessions) getNextNotifyId() (id int64) {
	id = s.nextNotifyId
	s.nextNotifyId--
	return
}

func (s *authSessions) getNextPushId() (id int64) {
	id = s.nextPushId
	s.nextPushId++
	return
}

func (s *authSessions) getAuthKeyId() int64 {
	return s.authKeyId
}

func (s *authSessions) getTempAuthKeyId() int64 {
	return s.authKeyId
}

func (s *authSessions) getUserId() int32 {
	return s.AuthUserId
}

func (s *authSessions) setUserId(userId int32) {
	s.AuthUserId = userId
	s.onBindUser(userId)
}

func (s *authSessions) getCacheSalt() *mtproto.TLFutureSalt {
	return s.cacheSalt
}

func (s *authSessions) getLayer() int32 {
	if s.Layer == 0 {
		s.Layer, _ = s.GetCacheApiLayer(context.Background(), s.authKeyId)
	}
	return s.Layer
}

func (s *authSessions) setLayer(layer int32) {
	if layer != 0 {
		s.Layer = layer
		s.PutCacheApiLayer(context.Background(), s.authKeyId, layer)
	}
}

func (s *authSessions) getClient() string {
	if s.Client == "" {
		s.Client = s.GetCacheClient(context.Background(), s.authKeyId)
	}
	return s.Client
}

func (s *authSessions) setClient(c string) {
	if c != "" {
		s.Client = c
		s.PutCacheClient(context.Background(), s.authKeyId, c)
	}
}

func (s *authSessions) getIpAddress() string {
	if s.IpAddress == "" {
		s.IpAddress = s.GetCacheIpAddress(context.Background(), s.authKeyId)
	}
	return s.IpAddress
}

func (s *authSessions) setIpAddress(c string) {
	if c != "" {
		s.IpAddress = c
		s.PutCacheIpAddress(context.Background(), s.authKeyId, c)
	}
}

func (s *authSessions) destroySession(sessionId int64) bool {
	delete(s.sessions, sessionId)
	return true
}

func (s *authSessions) sendToRpcQueue(rpcMessage *rpcApiMessage) {
	s.rpcQueue.Push(rpcMessage)
}

func (s *authSessions) getPushSessionId() int64 {
	if s.pushSessionId == 0 && s.AuthUserId != 0 {
		s.pushSessionId, _ = s.GetCachePushSessionID(context.Background(), s.AuthUserId, s.authKeyId)
	}
	return s.pushSessionId
}

func (s *authSessions) onBindUser(userId int32) {
	s.state = userIdLoaded
	s.AuthUserId = userId

	s.getPushSessionId()

	if s.Layer == 0 {
		layer, _ := s.GetCacheApiLayer(context.Background(), s.authKeyId)
		if layer != 0 {
			s.onBindLayer(layer)
		}
	}
}

func (s *authSessions) onBindPushSessionId(sessionId int64) {
	if s.pushSessionId == 0 {
		s.pushSessionId = sessionId
	}
	sess := s.sessions[sessionId]
	if sess != nil {
		sess.isAndroidPush = true
		sess.cb.setOnline()
	}
}

func (s *authSessions) onBindLayer(layer int32) {
	s.Layer = layer
}

func (s *authSessions) setOnline() {
	date := time.Now().Unix()
	if (s.onlineExpired == 0 || date > s.onlineExpired-kPingAddTimeout) && s.AuthUserId != 0 {
		status_client.AddOnline(context.Background(), s.AuthUserId, s.authKeyId, env.Hostname)
		s.onlineExpired = int64(date + 60)
	} else {
	}
}

func (s *authSessions) trySetOffline() {
	for _, sess := range s.sessions {
		if (sess.isGeneric && sess.sessionOnline()) ||
			(sess.isAndroidPush && sess.sessionOnline()) {
			return
		}
	}
	status_client.DelOnline(context.Background(), s.AuthUserId, s.authKeyId)
	s.onlineExpired = 0
}

func (s *authSessions) String() string {
	return fmt.Sprintf("{auth_key_id: %d, user_id: %d, layer: %d}", s.authKeyId, s.AuthUserId, s.Layer)
}

func (s *authSessions) Start() {
	s.running.Set(1)
	s.finish.Add(1)
	go s.rpcRunLoop()
	go s.runLoop()
}

func (s *authSessions) Stop() {
	s.trySetOffline()
	s.running.Set(0)
	s.rpcQueue.Close()
}

func (s *authSessions) runLoop() {
	defer func() {
		s.finish.Done()
		close(s.closeChan)
		s.finish.Wait()
	}()

	for s.running.Get() == 1 {
		select {
		case <-s.closeChan:
			return

		case sessionMsg := <-s.sessionDataChan:
			switch sm := sessionMsg.(type) {
			case *sessionData:
				s.onSessionData(sm)
			case *sessionHttpData:
				s.onSessionHttpData(sm)
			case *syncRpcResultData:
				s.onSyncRpcResultData(sm)
			case *syncData:
				s.onSyncData(sm)
			case *syncSessionData:
				s.onSyncSessionData(sm)
			case *connData:
				if sm.isNew {
					s.onSessionNew(sm)
				} else {
					s.onSessionClosed(sm)
				}
			default:
				panic("receive invalid type msg")
			}
		case rpcMessages := <-s.rpcDataChan:
			result, _ := rpcMessages.(*rpcApiMessage)
			s.onRpcResult(result)
		case <-time.After(time.Second):
			s.onTimer()
		}
	}

	log.Info("quit runLoop...")
}

func (s *authSessions) rpcRunLoop() {
	for {
		apiRequest := s.rpcQueue.Pop()
		if apiRequest == nil {
			log.Info("quit rpcRunLoop...")
			return
		} else {
			request, _ := apiRequest.(*rpcApiMessage)
			if s.onRpcRequest(request) {
				s.rpcDataChan <- request
			}
		}
	}
}

func (s *authSessions) onTimer() {
	for _, sess := range s.sessions {
		if (sess.isGeneric && sess.sessionOnline()) ||
			sess.isAndroidPush && sess.sessionOnline() {
			s.setOnline()
		}

		sess.onTimer()
	}

	for _, sess := range s.sessions {
		if !sess.sessionClosed() {
			return
		}
	}

	go func() {
		s.DeleteByAuthKeyId(s.authKeyId)
	}()
}

func (s *authSessions) sessionClientNew(gatewayId string, sessionId int64) error {
	s.sessionDataChan <- &connData{true, gatewayId, sessionId}
	return nil
}

func (s *authSessions) sessionDataArrived(gatewayId, clientIp string, sessionId, salt int64, buf []byte) error {
	s.sessionDataChan <- &sessionData{gatewayId, clientIp, sessionId, salt, buf}
	return nil
}

func (s *authSessions) sessionHttpDataArrived(gatewayId, clientIp string, sessionId, salt int64, buf []byte, resChan chan interface{}) error {
	s.sessionDataChan <- &sessionHttpData{gatewayId, clientIp, sessionId, salt, buf, resChan}
	return nil
}

func (s *authSessions) sessionClientClosed(gatewayId string, sessionId int64) error {
	s.sessionDataChan <- &connData{false, gatewayId, sessionId}
	return nil
}

func (s *authSessions) syncRpcResultDataArrived(sessionId, clientMsgId int64, data []byte) error {
	s.sessionDataChan <- &syncRpcResultData{sessionId, clientMsgId, data}
	return nil
}

func (s *authSessions) syncSessionDataArrived(sessionId int64, data *messageData) error {
	s.sessionDataChan <- &syncSessionData{sessionId, data}
	return nil
}

func (s *authSessions) syncDataArrived(needAndroidPush bool, data *messageData) error {
	s.sessionDataChan <- makeSyncData(needAndroidPush, data)
	return nil
}

func (s *authSessions) onSessionNew(connMsg *connData) {
	sess, ok := s.sessions[connMsg.sessionId]
	if !ok {
		sess = newSession(connMsg.sessionId, s)
		s.sessions[connMsg.sessionId] = sess
	}
	sess.onSessionConnNew(connMsg.gatewayId)
}

func (s *authSessions) onSessionData(sessionMsg *sessionData) {
	var (
		err      error
		message2 = &mtproto.TLMessage2{}
		now      = int32(time.Now().Unix())
	)

	err = message2.Decode(mtproto.NewDecodeBuf(sessionMsg.buf))
	if err != nil {
		log.Errorf("onSessionData - error: {%s}, data: {sessions: %s, gate_id: %d}", err, s, sessionMsg.gatewayId)
		return
	}

	if s.cacheSalt == nil {
		s.cacheSalt, s.cacheLastSalt, _ = s.GetOrFetchNewSalt(context.Background(), s.authKeyId)
	} else {
		if now >= s.cacheSalt.GetValidUntil() {
			s.cacheSalt, s.cacheLastSalt, _ = s.GetOrFetchNewSalt(context.Background(), s.authKeyId)
		}
	}

	if s.cacheSalt == nil {
		log.Errorf("onSessionData - getOrFetchNewSalt nil error, data: {sessions: %s, conn_id: %s}", s, sessionMsg.gatewayId)
		return
	}

	sess, ok := s.sessions[sessionMsg.sessionId]
	if !ok {
		sess = newSession(sessionMsg.sessionId, s)
		s.sessions[sessionMsg.sessionId] = sess
	}

	sess.onSessionConnNew(sessionMsg.gatewayId)
	sess.onSessionMessageData(sessionMsg.gatewayId, sessionMsg.clientIp, sessionMsg.salt, message2)
}

func (s *authSessions) onSessionHttpData(sessionMsg *sessionHttpData) {
	var (
		err      error
		message2 = &mtproto.TLMessage2{}
		now      = int32(time.Now().Unix())
	)

	err = message2.Decode(mtproto.NewDecodeBuf(sessionMsg.buf))
	if err != nil {
		log.Errorf("onSessionData - error: {%s}, data: {sessions: %s, gate_id: %d}", err, s, sessionMsg.gatewayId)
		return
	}

	if s.cacheSalt == nil {
		s.cacheSalt, s.cacheLastSalt, _ = s.GetOrFetchNewSalt(context.Background(), s.authKeyId)
	} else {
		if now >= s.cacheSalt.GetValidUntil() {
			s.cacheSalt, s.cacheLastSalt, _ = s.GetOrFetchNewSalt(context.Background(), s.authKeyId)
		}
	}

	if s.cacheSalt == nil {
		log.Errorf("onSessionData - getOrFetchNewSalt nil error, data: {sessions: %s, conn_id: %s}", s, sessionMsg.gatewayId)
		return
	}

	sess, ok := s.sessions[sessionMsg.sessionId]
	if !ok {
		sess = newSession(sessionMsg.sessionId, s)
		s.sessions[sessionMsg.sessionId] = sess
	}

	sess.isHttp = true
	sess.httpQueue.Push(sessionMsg.resChannel)
	sess.onSessionConnNew(sessionMsg.gatewayId)
	sess.onSessionMessageData(sessionMsg.gatewayId, sessionMsg.clientIp, sessionMsg.salt, message2)
}

func (s *authSessions) onSessionClosed(connMsg *connData) {
	if sess, ok := s.sessions[connMsg.sessionId]; !ok {
		log.Warnf("session conn closed -  conn: ", connMsg, ", sess: ", sess)
	} else {
		sess.onSessionConnClose(connMsg.gatewayId)
	}
}

func (s *authSessions) onSyncRpcResultData(syncMsg *syncRpcResultData) {
	log.Infof("onSyncRpcResultData - receive data: {sess: %s}",
		s)

	sess := s.sessions[syncMsg.sessionId]
	if sess != nil {
		sess.onSyncRpcResultData(syncMsg.clientMsgId, syncMsg.data)
	}
}

func (s *authSessions) onSyncSessionData(syncMsg *syncSessionData) {
	log.Infof("onSyncSessionData - receive data: {sess: %s}",
		s)
	sess := s.sessions[syncMsg.sessionId]
	if sess != nil {
		sess.onSyncSessionData(syncMsg.data.obj)
	}
}

func (s *authSessions) onSyncData(syncMsg *syncData) {
	log.Info("authSessions - ", reflect.TypeOf(syncMsg.data.obj))
	if upds, ok := syncMsg.data.obj.(*mtproto.Updates); ok {
		if upds.PredicateName == mtproto.Predicate_updateAccountResetAuthorization {
			log.Info("recv updateAccountResetAuthorization - ", reflect.TypeOf(syncMsg.data.obj))
			if s.AuthUserId != upds.GetUserId() {
				log.Errorf("upds -- ", upds)
			}
			s.PutCacheUserId(context.Background(), s.authKeyId, 0)
			s.DeleteByAuthKeyId(s.authKeyId)
			s.AuthUserId = 0
			return
		} else {
		}
	}

	for _, sess2 := range s.sessions {
		if sess2.isGeneric {
			sess2.onSyncData(syncMsg.data.obj)
		}
		if sess2.isAndroidPush {
			if syncMsg.needAndroidPush {
				sess2.onSyncData(nil)
			}
		}
	}
}

func (s *authSessions) onRpcResult(rpcResult *rpcApiMessage) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("tcp_server handle panic: %v\n%s", err, debug.Stack())
		}
	}()

	if sess, ok := s.sessions[rpcResult.sessionId]; ok {
		sess.onRpcResult(rpcResult)
	} else {
		log.Warnf("onRpcResult - not found rpcSession by sessionId: ", rpcResult.sessionId)
	}
}

func (s *authSessions) onRpcRequest(request *rpcApiMessage) bool {
	var (
		err       error
		rpcResult mtproto.TLObject
	)

	rpcMetadata := &grpc_util.RpcMetadata{
		ServerId:    env.Hostname,
		ClientAddr:  request.clientIp,
		AuthId:      s.authKeyId,
		SessionId:   request.sessionId,
		ReceiveTime: time.Now().Unix(),
		UserId:      s.AuthUserId,
		ClientMsgId: request.reqMsgId,
		Layer:       s.Layer,
		Client:      s.getClient(),
	}

	switch req := request.reqMsg.(type) {
	case *mtproto.TLAuthBindTempAuthKey:
		rpcResult, err = s.Service.Dao.AuthSessionRpcClient.AuthBindTempAuthKey(context.Background(), req)
	default:
		rpcResult, err = s.Service.Dao.Invoke(rpcMetadata, req)
	}

	reply := &mtproto.TLRpcResult{
		ReqMsgId: request.reqMsgId,
	}

	if err != nil {
		log.Error(err.Error())
		if rpcErr, ok := err.(*mtproto.TLRpcError); ok {
			reply.Result = rpcErr
		} else {
			reply.Result = mtproto.NewRpcError(int32(mtproto.TLRpcErrorCodes_INTERNAL), "INTERNAL_SERVER_ERROR")
		}
	} else {
		log.Debugf("invokeRpcRequest - rpc_result: {%s}\n", reflect.TypeOf(rpcResult))
		reply.Result = rpcResult
	}

	request.rpcResult = reply

	if _, ok := request.reqMsg.(*mtproto.TLAuthLogOut); ok {
		log.Debugf("authLogOut - %#v", rpcMetadata)
		s.PutCacheUserId(context.Background(), s.authKeyId, 0)
	}
	return true
}
