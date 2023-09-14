package service

import (
	"context"
	"time"

	"open.chat/app/interface/session/sessionpb"
	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/mtproto"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) CreateSession(ctx context.Context, r *sessionpb.SessionClientEvent) (res *mtproto.Bool, err error) {
	log.Debugf("CreateSession - request: %s", logger.JsonDebugData(r))
	var (
		sessList *authSessions
		ok       bool
	)

	s.mu.Lock()
	defer s.mu.Unlock()

	if sessList, ok = s.sessionsManager[r.AuthKeyId]; !ok {
		sessList, err = newAuthSessions(r.AuthKeyId, s)
		if err != nil {
			return
		}
		s.sessionsManager[r.AuthKeyId] = sessList
	}
	sessList.sessionClientNew(r.ServerId, r.SessionId)
	res = mtproto.ToBool(true)
	return
}

func (s *Service) CloseSession(ctx context.Context, r *sessionpb.SessionClientEvent) (*mtproto.Bool, error) {
	log.Debugf("CloseSession - request: %s", logger.JsonDebugData(r))
	var (
		sessList *authSessions
		ok       bool
	)

	s.mu.RLock()
	defer s.mu.RUnlock()

	if sessList, ok = s.sessionsManager[r.AuthKeyId]; !ok {
		log.Warn("not found sessList by keyId: %d", r.AuthKeyId)
	} else {
		sessList.sessionClientClosed(r.ServerId, r.SessionId)
	}

	return mtproto.ToBool(true), nil
}

func (s *Service) SendAsyncSessionData(ctx context.Context, r *sessionpb.SessionClientData) (res *mtproto.Bool, err error) {
	r2 := &sessionpb.SessionClientData{
		ServerId:  r.ServerId,
		ClientIp:  r.ClientIp,
		ConnType:  r.ConnType,
		AuthKeyId: r.AuthKeyId,
		QuickAck:  r.QuickAck,
		SessionId: r.SessionId,
		Salt:      r.Salt,
		Payload:   r.Payload[:32],
	}
	log.Debugf("SendAsyncSessionData - request: %s,Payload:%d", logger.JsonDebugData(r2), len(r.Payload))
	var (
		sessList *authSessions
		ok       bool
	)

	s.mu.Lock()
	defer s.mu.Unlock()

	if sessList, ok = s.sessionsManager[r.AuthKeyId]; !ok {
		sessList, err = newAuthSessions(r.AuthKeyId, s)
		if err != nil {
			return
		}

		s.sessionsManager[r.AuthKeyId] = sessList
	}
	sessList.sessionDataArrived(r.ServerId, r.ClientIp, r.SessionId, r.Salt, r.Payload)
	res = mtproto.ToBool(true)
	return
}

func (s *Service) SendSyncSessionData(ctx context.Context, r *sessionpb.SessionClientData) (res *sessionpb.SessionData, err error) {
	r2 := &sessionpb.SessionClientData{
		ServerId:  r.ServerId,
		ClientIp:  r.ClientIp,
		ConnType:  r.ConnType,
		AuthKeyId: r.AuthKeyId,
		QuickAck:  r.QuickAck,
		SessionId: r.SessionId,
		Salt:      r.Salt,
		Payload:   r.Payload[:32],
	}
	log.Debugf("SendSyncSessionData - request: %s,Payload:%d", logger.JsonDebugData(r2), len(r.Payload))
	respChan := make(chan interface{}, 2)

	s.mu.Lock()
	var (
		sessList *authSessions
		ok       bool
	)

	if sessList, ok = s.sessionsManager[r.AuthKeyId]; !ok {
		sessList, err = newAuthSessions(r.AuthKeyId, s)
		if err != nil {
			s.mu.Unlock()
			log.Errorf("sendSyncSessionData - error: %v", err)
			return
		}
		s.sessionsManager[r.AuthKeyId] = sessList
	}
	sessList.sessionHttpDataArrived(r.ServerId, r.ClientIp, r.SessionId, r.Salt, r.Payload, respChan)
	s.mu.Unlock()

	timer := time.NewTimer(60 * time.Second)
	select {
	case rc := <-respChan:
		if rc != nil {
			if payload, ok := rc.([]byte); ok {
				log.Debugf("recv http data: %d", len(payload))
				res = &sessionpb.SessionData{
					Payload: payload,
				}
				return
			}
		}
		err = mtproto.ErrInternelServerError
		log.Errorf("sendSyncSessionData - error: %v", err)
	case <-timer.C:
		err = mtproto.ErrTimeOut503
		log.Errorf("sendSyncSessionData - error: %v", err)
	}

	return
}

func (s *Service) Heartbeat(ctx context.Context, r *sessionpb.SessionClientEvent) (*mtproto.Bool, error) {
	log.Debugf("Heartbeat - request: %s", logger.JsonDebugData(r))
	return mtproto.ToBool(true), nil
}

func (s *Service) DestroySessions(ctx context.Context, r *sessionpb.AuthId) (*mtproto.Bool, error) {
	log.Debugf("DestroySessions - request: %s", logger.JsonDebugData(r))
	return mtproto.ToBool(true), nil
}

func (s *Service) ImportSessions(ctx context.Context, r *sessionpb.AuthSessionIdList) (*mtproto.Bool, error) {
	log.Debugf("ImportSessions - request: %s", logger.JsonDebugData(r))
	return mtproto.ToBool(true), nil
}

func (s *Service) SessionQueryAuthKey(ctx context.Context, r *sessionpb.AuthId) (*sessionpb.AuthKeyInfo, error) {
	log.Debugf("session.queryAuthKey - request: %s", logger.JsonDebugData(r))
	key, err := s.Dao.AuthSessionRpcClient.SessionQueryAuthKey(ctx, &authsessionpb.TLSessionQueryAuthKey{
		AuthKeyId: r.AuthKeyId,
	})
	if err != nil {
		log.Errorf("queryAuthKey error: %v", err)
		return nil, err
	}
	reply := &sessionpb.AuthKeyInfo{
		AuthKeyId:          key.AuthKeyId,
		AuthKey:            key.AuthKey,
		AuthKeyType:        key.AuthKeyType,
		PermAuthKeyId:      key.PermAuthKeyId,
		TempAuthKeyId:      key.TempAuthKeyId,
		MediaTempAuthKeyId: key.MediaTempAuthKeyId,
	}
	log.Debugf("session.setAuthKey - reply: {%s}", logger.JsonDebugData(reply))
	return reply, nil
}

func (s *Service) SessionSetAuthKey(ctx context.Context, r *sessionpb.AuthKeySalts) (*mtproto.Bool, error) {
	log.Debugf("session.setAuthKey - request: %s", logger.JsonDebugData(r))
	if r.AuthKey == nil {
		log.Errorf("setAuthKey error: auth_key is nil")
		return nil, mtproto.ErrInputRequestInvalid
	}
	reply, err := s.Dao.AuthSessionRpcClient.SessionSetAuthKey(ctx, &authsessionpb.TLSessionSetAuthKey{
		AuthKey: &authsessionpb.AuthKeyInfo{
			AuthKeyId:          r.AuthKey.AuthKeyId,
			AuthKey:            r.AuthKey.AuthKey,
			AuthKeyType:        r.AuthKey.AuthKeyType,
			PermAuthKeyId:      r.AuthKey.PermAuthKeyId,
			TempAuthKeyId:      r.AuthKey.TempAuthKeyId,
			MediaTempAuthKeyId: r.AuthKey.MediaTempAuthKeyId,
		},
		FutureSalt: r.FutureSalt,
	})
	if err != nil {
		log.Errorf("session.setAuthKey - error: %v", err)
	}
	log.Debugf("session.setAuthKey - reply: {%s}", reply.DebugString())
	return reply, err
}
