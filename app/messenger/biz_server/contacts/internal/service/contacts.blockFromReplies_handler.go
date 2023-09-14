package service

import (
	"context"

	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"

	"open.chat/mtproto"
)

// contacts.blockFromReplies#29a8962c flags:# delete_message:flags.0?true delete_history:flags.1?true report_spam:flags.2?true msg_id:int = Updates;
func (s *Service) ContactsBlockFromReplies(ctx context.Context, request *mtproto.TLContactsBlockFromReplies) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("contacts.blockFromReplies - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// Sorry: not impl ContactsBlockFromReplies logic
	log.Warn("contacts.blockFromReplies - error: method ContactsBlockFromReplies not impl")

	return nil, mtproto.ErrMethodNotImpl
}
