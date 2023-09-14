package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsGetBroadcastsForDiscussion(ctx context.Context, request *mtproto.TLChannelsGetBroadcastsForDiscussion) (*mtproto.Messages_Chats, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.getBroadcastsForDiscussion - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("channels.getBroadcastsForDiscussion - error: %v", err)
		return nil, err
	}

	result := mtproto.MakeTLMessagesChats(&mtproto.Messages_Chats{
		Chats: []*mtproto.Chat{},
	}).To_Messages_Chats()

	log.Debugf("channels.getBroadcastsForDiscussion - reply %s", result.DebugString())
	return result, nil
}
