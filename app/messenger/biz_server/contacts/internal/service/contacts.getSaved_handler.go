package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ContactsGetSaved(ctx context.Context, request *mtproto.TLContactsGetSaved) (*mtproto.Vector_SavedContact, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("contacts.getSaved#82f1e39f - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return &mtproto.Vector_SavedContact{
		Datas: []*mtproto.SavedContact{},
	}, nil
}
