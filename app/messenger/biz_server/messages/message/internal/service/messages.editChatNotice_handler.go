package service

import (
	"context"

	"open.chat/pkg/log"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
)

// messages.editChatNotice#8ca0d616 peer:InputPeer notice:string = Bool;
func (s *Service) MessagesEditChatNotice(ctx context.Context, request *mtproto.TLMessagesEditChatNotice) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.editChatNotice - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// Sorry: not impl MessagesEditChatNotice logic
	log.Warn("messages.editChatNotice - error: method MessagesEditChatNotice not impl")

	return nil, mtproto.ErrMethodNotImpl
}
