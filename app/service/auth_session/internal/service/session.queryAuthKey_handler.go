package service

import (
	"context"

	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) SessionQueryAuthKey(ctx context.Context, request *authsessionpb.TLSessionQueryAuthKey) (*authsessionpb.AuthKeyInfo, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Infof("session.queryAuthKey - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	reply, err := s.AuthSessionCore.QueryAuthKey(ctx, request.GetAuthKeyId())
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	log.Infof("session.queryAuthKey - reply: {%s}", reply.DebugString())
	return reply, nil
}
