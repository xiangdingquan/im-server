package service

import (
	"context"

	"open.chat/app/interface/session/sessionpb"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

func (s *Service) PushUpdates(ctx context.Context, r *sessionpb.PushUpdatesData) (*mtproto.Bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var (
		sessList *authSessions
		ok       bool
	)
	if sessList, ok = s.sessionsManager[r.AuthKeyId]; !ok {
		log.Errorf("not found authKeyId")
		return mtproto.ToBool(false), nil
	}

	sessList.syncDataArrived(r.Notification, &messageData{obj: r.Updates})
	return mtproto.ToBool(true), nil
}

func (s *Service) PushSessionUpdates(ctx context.Context, r *sessionpb.PushSessionUpdatesData) (*mtproto.Bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var (
		sessList *authSessions
		ok       bool
	)
	if sessList, ok = s.sessionsManager[r.AuthKeyId]; !ok {
		log.Errorf("not found authKeyId")
		return mtproto.ToBool(false), nil
	}

	sessList.syncSessionDataArrived(r.SessionId, &messageData{obj: r.Updates})
	return mtproto.ToBool(true), nil
}

func (s *Service) PushRpcResult(ctx context.Context, r *sessionpb.PushRpcResultData) (*mtproto.Bool, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var (
		sessList *authSessions
		ok       bool
	)
	if sessList, ok = s.sessionsManager[r.AuthKeyId]; !ok {
		log.Errorf("not found authKeyId")
		return mtproto.ToBool(false), nil
	}

	sessList.syncRpcResultDataArrived(r.SessionId, r.ClientReqMsgId, r.RpcResultData)
	return mtproto.ToBool(true), nil
}
