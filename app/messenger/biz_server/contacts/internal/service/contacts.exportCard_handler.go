package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ContactsExportCard(ctx context.Context, request *mtproto.TLContactsExportCard) (*mtproto.Vector_Int, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("contacts.exportCard#84e53737 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	exports := &mtproto.Vector_Int{
		Datas: []int32{},
	}

	log.Debugf("contacts.exportCard#84e53737 - not impl ContactsExportCard, reply: {}")
	return exports, nil
}
