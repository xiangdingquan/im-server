package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ContactsDeleteByPhones(ctx context.Context, request *mtproto.TLContactsDeleteByPhones) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("contacts.deleteByPhones - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return mtproto.BoolTrue, nil
}
