package service

import (
	"context"

	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) SessionGetPushSessionId(ctx context.Context, request *authsessionpb.TLSessionGetPushSessionId) (*mtproto.Int64, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("session.getPushSessionId - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	sessionId := s.AuthSessionCore.GetPushSessionId(ctx, request.GetUserId(), request.GetAuthKeyId(), request.GetTokenType())
	reply := &mtproto.TLInt64{Data2: &mtproto.Int64{
		V: sessionId,
	}}

	log.Debugf("session.getPushSessionId - reply: {%s}", reply.DebugString())
	return reply.To_Int64(), nil
}
