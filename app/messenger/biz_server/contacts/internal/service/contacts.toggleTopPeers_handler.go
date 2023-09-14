package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ContactsToggleTopPeers(ctx context.Context, request *mtproto.TLContactsToggleTopPeers) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("contacts.toggleTopPeers#8514bdda - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	log.Debugf("contacts.toggleTopPeers#8514bdda - reply: {true}")
	return mtproto.ToBool(true), nil
}
