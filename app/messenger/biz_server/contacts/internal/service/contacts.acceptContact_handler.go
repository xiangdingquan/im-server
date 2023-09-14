package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ContactsAcceptContact(ctx context.Context, request *mtproto.TLContactsAcceptContact) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("contacts.acceptContact - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return model.MakeEmptyUpdates(), nil
}
