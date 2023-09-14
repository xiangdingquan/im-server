package service

import (
	"context"

	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"

	"open.chat/mtproto"
)

// messages.getReplies#24b581ba peer:InputPeer msg_id:int offset_id:int offset_date:int add_offset:int limit:int max_id:int min_id:int hash:int = messages.Messages;
func (s *Service) MessagesGetReplies(ctx context.Context, request *mtproto.TLMessagesGetReplies) (*mtproto.Messages_Messages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getReplies - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// Sorry: not impl MessagesGetReplies logic
	log.Warn("messages.getReplies - error: method MessagesGetReplies not impl")

	return nil, mtproto.ErrMethodNotImpl
}
