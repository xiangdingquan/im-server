package service

import (
	"context"

	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"

	"open.chat/mtproto"
)

// messages.readDiscussion#f731a9f4 peer:InputPeer msg_id:int read_max_id:int = Bool;
func (s *Service) MessagesReadDiscussion(ctx context.Context, request *mtproto.TLMessagesReadDiscussion) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.readDiscussion - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// Sorry: not impl MessagesReadDiscussion logic
	log.Warn("messages.readDiscussion - error: method MessagesReadDiscussion not impl")

	return nil, mtproto.ErrMethodNotImpl
}
