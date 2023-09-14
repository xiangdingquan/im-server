package service

import (
	"context"
	"reflect"
	"time"

	"github.com/go-kratos/kratos/pkg/conf/env"

	"fmt"
	"open.chat/app/interface/session/sessionpb"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"strings"
)

func (s *Service) Invoke(ctx context.Context, token string, r mtproto.TLObject) (mtproto.TLObject, error) {
	var (
		rpcResult mtproto.TLObject
	)

	a := strings.Split(token, ":")
	if len(a) != 2 {
		return nil, fmt.Errorf("invalid tolen: %s", token)
	}

	auth, err := s.dao.GetCacheAuthUser(ctx, a[0], a[1])
	if err != nil {
		log.Errorf("getBotSessionByToken error: %v", err)
		return nil, err
	}

	botSession := s.sessions.Get(auth.AuthKeyId())
	if botSession == nil {
		botSession = newCacheBotSession(auth.UserId(), auth.AuthKeyId(), auth.Layer())
		s.sessions.Put(botSession)
	}

	rpcMetadata := &grpc_util.RpcMetadata{
		ServerId:    env.Hostname,
		ClientAddr:  env.Hostname,
		AuthId:      botSession.AuthKeyId(),
		SessionId:   botSession.SessionId(),
		ReceiveTime: time.Now().Unix(),
		UserId:      botSession.UserId(),
		ClientMsgId: nextMessageId(true),
		IsBot:       true,
		Layer:       botSession.Layer(),
	}

	rpcResult, err = s.dao.Invoke(rpcMetadata, r)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	log.Debugf("callSession - rpc_result: {%s}\n", reflect.TypeOf(rpcResult))
	return rpcResult, nil
}

func (s *Service) PushUpdates(ctx context.Context, r *sessionpb.PushUpdatesData) (*mtproto.Bool, error) {
	botSession := s.sessions.Get(r.AuthKeyId)
	if botSession == nil {
		log.Errorf("botSession down: %d", r.AuthKeyId)
		return mtproto.ToBool(false), nil
	}

	err := s.sessions.reqCache.shoot(r.AuthKeyId, r.Updates)
	if err != nil {
		botSession.PushUpdates(r.Updates)
	}

	return mtproto.ToBool(true), nil
}

func (s *Service) PushRpcResult(ctx context.Context, r *sessionpb.PushRpcResultData) (*mtproto.Bool, error) {
	return mtproto.ToBool(true), nil
}

func (s *Service) PushSessionUpdates(ctx context.Context, r *sessionpb.PushSessionUpdatesData) (*mtproto.Bool, error) {
	return mtproto.ToBool(true), nil
}
