package server

import (
	"context"

	"open.chat/app/interface/gateway/egatepb"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

func (s *Server) ReceiveData(ctx context.Context, r *egatepb.SessionRawData) (reply *mtproto.Bool, err error) {
	log.Debugf("ReceiveData - request: {kId: %d, sessionId: %d, payloadLen: %d}", r.AuthKeyId, r.SessionId, len(r.Payload))

	var (
		authKey *authKeyUtil
	)

	isWebsocket, connIdList := s.authSessionMgr.FoundSessionConnIdList(r.AuthKeyId, r.SessionId)
	if connIdList == nil {
		log.Errorf("ReceiveData - not found connIdList - keyId: %d, sessionId: %d", r.AuthKeyId, r.SessionId)
		return mtproto.ToBool(false), nil
	}
	for _, connId := range connIdList {
		log.Debugf("[keyId: %d, sessionId: %d, isWebsocket: %v]: %v", r.AuthKeyId, r.SessionId, isWebsocket, connId)
		conn2 := s.server.GetConnection(connId)
		if conn2 != nil {
			ctx, _ := conn2.Context.(*connContext)
			authKey = ctx.getAuthKey(r.AuthKeyId)
			if authKey == nil {
				log.Warn("invalid authKeyId, authKeyId = %d", r.AuthKeyId)
				continue
			}
			if !isWebsocket && ctx.isHttp {
				if !ctx.canSend {
					continue
				}
			}
			if err = s.SendToClient(conn2, authKey, r.Payload); err == nil {
				log.Debugf("ReceiveData -  result: {auth_key_id = %d, session_id = %d, conn = %s}",
					r.AuthKeyId,
					r.SessionId,
					conn2)

				if !isWebsocket && ctx.isHttp {
					s.authSessionMgr.PushBackHttpData(r.AuthKeyId, r.SessionId, r.Payload)
				}
				return mtproto.ToBool(true), nil
			} else {
				log.Errorf("ReceiveData - sendToClient error (%v), auth_key_id = %d, session_id = %d, conn_id_list = %v",
					err,
					r.AuthKeyId,
					r.SessionId,
					connIdList)
			}
		}
	}

	log.Errorf("ReceiveData - conn closed, auth_key_id = %d, session_id = %d, conn_id_list = %v",
		r.AuthKeyId,
		r.SessionId,
		connIdList)
	return mtproto.ToBool(false), nil
}
