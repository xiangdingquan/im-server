package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ContactsResetSaved(ctx context.Context, request *mtproto.TLContactsResetSaved) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("ContactsResetSaved - metadata: %s, request: %s", md.DebugString(), request.DebugString())
	log.Debugf("ContactsResetSaved - reply: {true}")
	return mtproto.ToBool(true), nil
}
