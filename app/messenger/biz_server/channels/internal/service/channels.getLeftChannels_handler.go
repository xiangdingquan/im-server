package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsGetLeftChannels(ctx context.Context, request *mtproto.TLChannelsGetLeftChannels) (*mtproto.Messages_Chats, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.getLeftChannels - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("channels.getLeftChannels - error: %v", err)
		return nil, err
	}

	channels, err := s.ChannelFacade.GetLeftChannelList(ctx, md.UserId, request.Offset)
	if err != nil {
		log.Errorf("channels.getLeftChannels - error: %v", err)
		return nil, err
	}

	messagesChats := mtproto.MakeTLMessagesChats(&mtproto.Messages_Chats{
		Chats: channels.ToChats(md.UserId),
		Count: int32(len(channels)),
	}).To_Messages_Chats()

	log.Debugf("channels.getLeftChannels - reply: {%s}", messagesChats.DebugString())
	return messagesChats, nil
}
