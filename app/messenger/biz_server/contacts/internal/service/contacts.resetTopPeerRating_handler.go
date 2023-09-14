package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ContactsResetTopPeerRating(ctx context.Context, request *mtproto.TLContactsResetTopPeerRating) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("ContactsResetTopPeerRating - metadata: %s, request: %s", md.DebugString(), request.DebugString())
	log.Debugf("ContactsResetTopPeerRating - reply: {true}")
	return mtproto.ToBool(true), nil
}
