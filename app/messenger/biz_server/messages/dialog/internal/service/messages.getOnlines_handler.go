package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetOnlines(ctx context.Context, request *mtproto.TLMessagesGetOnlines) (*mtproto.ChatOnlines, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getOnlines#6e2be050 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	reply := mtproto.MakeTLChatOnlines(&mtproto.ChatOnlines{
		Onlines: 1,
	})

	log.Debugf("messages.getOnlines#6e2be050 - reply: {%s}", reply.DebugString())
	return reply.To_ChatOnlines(), nil
}
