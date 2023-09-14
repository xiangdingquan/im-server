package service

import (
	"context"

	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"

	"open.chat/mtproto"
)

// messages.getDiscussionMessage#446972fd peer:InputPeer msg_id:int = messages.DiscussionMessage;
func (s *Service) MessagesGetDiscussionMessage(ctx context.Context, request *mtproto.TLMessagesGetDiscussionMessage) (*mtproto.Messages_DiscussionMessage, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getDiscussionMessage - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// Sorry: not impl MessagesGetDiscussionMessage logic
	log.Warn("messages.getDiscussionMessage - error: method MessagesGetDiscussionMessage not impl")

	return nil, mtproto.ErrMethodNotImpl
}
