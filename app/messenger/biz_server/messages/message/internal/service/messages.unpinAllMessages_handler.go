package service

import (
	"context"

	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"

	"open.chat/mtproto"
)

// messages.unpinAllMessages#f025bc8b peer:InputPeer = messages.AffectedHistory;
func (s *Service) MessagesUnpinAllMessages(ctx context.Context, request *mtproto.TLMessagesUnpinAllMessages) (*mtproto.Messages_AffectedHistory, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.unpinAllMessages - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// Sorry: not impl MessagesUnpinAllMessages logic
	log.Warn("messages.unpinAllMessages - error: method MessagesUnpinAllMessages not impl")

	return nil, mtproto.ErrMethodNotImpl
}
