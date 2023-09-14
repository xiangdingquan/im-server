package service

import (
	"context"

	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) SessionGetLayer(ctx context.Context, request *authsessionpb.TLSessionGetLayer) (*mtproto.Int32, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("session.getLayer - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	keyData, err := s.AuthSessionCore.Dao.GetAuthKey(ctx, request.GetAuthKeyId())
	if err != nil {
		log.Errorf("session.getLayer - error: %v", err)
		return nil, err
	}

	layer := s.AuthSessionCore.GetApiLayer(ctx, keyData.PermAuthKeyId)
	reply := &mtproto.TLInt32{Data2: &mtproto.Int32{
		V: layer,
	}}

	log.Debugf("session.getLayer - reply: {%s}", reply.DebugString())
	return reply.To_Int32(), nil
}
