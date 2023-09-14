package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) ContactsGetTopPeers(ctx context.Context, request *mtproto.TLContactsGetTopPeers) (*mtproto.Contacts_TopPeers, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("contacts.getTopPeers#d4982db5 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	topPeers := &mtproto.TLContactsTopPeers{Data2: &mtproto.Contacts_TopPeers{
		Categories: []*mtproto.TopPeerCategoryPeers{},
		Chats:      []*mtproto.Chat{},
		Users:      []*mtproto.User{},
	}}

	log.Debugf("contacts.getTopPeers#d4982db5 - reply: %s", logger.JsonDebugData(topPeers))
	return topPeers.To_Contacts_TopPeers(), nil
}
