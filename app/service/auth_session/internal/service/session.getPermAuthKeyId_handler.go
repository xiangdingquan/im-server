package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) SessionGetPermAuthKeyId(ctx context.Context, request *mtproto.Int64) (*mtproto.Int64, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("session.getPermAuthKeyId - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	reply := &mtproto.TLInt64{Data2: &mtproto.Int64{
		V: s.AuthSessionCore.GetPermAuthKeyId(ctx, request.V),
	}}

	log.Debugf("session.getPermAuthKeyId - reply: {%s}", reply.DebugString())
	return reply.To_Int64(), nil
}
