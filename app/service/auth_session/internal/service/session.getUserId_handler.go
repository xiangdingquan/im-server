package service

import (
	"context"

	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) SessionGetUserId(ctx context.Context, request *authsessionpb.TLSessionGetUserId) (*mtproto.Int32, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("session.getUserId - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	keyData, err := s.AuthSessionCore.Dao.GetAuthKey(ctx, request.GetAuthKeyId())
	if err != nil {
		log.Errorf("session.getUserId - error: %v", err)
		return nil, err
	}

	userId := s.AuthSessionCore.GetAuthKeyUserId(ctx, keyData.PermAuthKeyId)
	reply := &mtproto.TLInt32{Data2: &mtproto.Int32{
		V: userId,
	}}

	log.Debugf("session.getUserId - reply: {%s}", reply.DebugString())
	return reply.To_Int32(), nil
}
