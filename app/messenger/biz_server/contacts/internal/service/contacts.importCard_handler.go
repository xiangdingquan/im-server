package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) ContactsImportCard(ctx context.Context, request *mtproto.TLContactsImportCard) (*mtproto.User, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("contacts.importCard#4fe196fe - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	user := mtproto.MakeTLUserEmpty(nil)

	log.Debugf("contacts.importCard#4fe196fe - reply: %s", logger.JsonDebugData(user))
	return user.To_User(), nil
}
