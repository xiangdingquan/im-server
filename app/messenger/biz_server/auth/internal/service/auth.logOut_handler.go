package service

import (
	"context"

	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) AuthLogOut(ctx context.Context, request *mtproto.TLAuthLogOut) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("auth.logOut#5717da40 - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))
	// unbind auth_key and user_id
	s.AuthSessionRpcClient.SessionUnbindAuthKeyUser(ctx, &authsessionpb.TLSessionUnbindAuthKeyUser{
		AuthKeyId: md.AuthId,
		UserId:    md.UserId,
	})

	log.Debugf("auth.logOut#5717da40 - reply: {true}")
	return mtproto.ToBool(true), nil
}
