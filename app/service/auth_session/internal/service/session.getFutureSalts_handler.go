package service

import (
	"context"

	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

const (
	kDefaultSaltNum = 32
)

func (s *Service) SessionGetFutureSalts(ctx context.Context, request *authsessionpb.TLSessionGetFutureSalts) (*mtproto.FutureSalts, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("session.getFutureSalts - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	num := request.GetNum()
	if num == 0 {
		num = kDefaultSaltNum
	}
	futureSalts, err := s.AuthSessionCore.GetFutureSalts(ctx, request.GetAuthKeyId(), num)
	if err != nil {
		log.Errorf("session.getFutureSalts - %v", err)
		return nil, err
	}

	log.Debugf("session.getFutureSalts - reply: {%s}", futureSalts.DebugString())
	return futureSalts.To_FutureSalts(), nil
}
