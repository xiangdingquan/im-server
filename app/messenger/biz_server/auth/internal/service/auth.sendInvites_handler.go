package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AuthSendInvites(ctx context.Context, request *mtproto.TLAuthSendInvites) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("auth.sendInvites#771c1d97 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	log.Debugf("auth.sendInvites#771c1d97 - reply: {true}")
	return mtproto.ToBool(true), nil
}
