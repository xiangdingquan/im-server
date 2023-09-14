package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ContactsGetContactIDs(ctx context.Context, request *mtproto.TLContactsGetContactIDs) (*mtproto.Vector_Int, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("contacts.getContactIDs - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return &mtproto.Vector_Int{
		Datas: []int32{},
	}, nil
}
