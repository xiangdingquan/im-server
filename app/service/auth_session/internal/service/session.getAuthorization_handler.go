package service

import (
	"context"

	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) SessionGetAuthorization(ctx context.Context, request *authsessionpb.TLSessionGetAuthorization) (*mtproto.Authorization, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("session.getAuthorization - metadata: %s, request: %s", md.DebugString(), logger.JsonDebugData(request))

	myKeyData, err := s.AuthSessionCore.Dao.GetAuthKey(ctx, request.AuthKeyId)
	if err != nil {
		log.Errorf("session.getAuthorization - error: %v", err)
		return nil, err
	}

	authorization, err := s.AuthSessionCore.GetAuthorization(ctx, myKeyData.PermAuthKeyId)
	if err != nil {
		log.Errorf("session.getAuthorization - error: %v", err)
		return nil, err
	}

	log.Debugf("session.getAuthorization - reply: {%s}", authorization.DebugString())
	return authorization, nil
}
