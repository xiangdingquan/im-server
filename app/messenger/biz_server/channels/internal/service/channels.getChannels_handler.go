package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsGetChannels(ctx context.Context, request *mtproto.TLChannelsGetChannels) (*mtproto.Messages_Chats, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.getChannels - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	messagesChats := mtproto.MakeTLMessagesChats(&mtproto.Messages_Chats{
		Chats: make([]*mtproto.Chat, 0, len(request.Id)),
		Count: 0,
	}).To_Messages_Chats()

	if len(request.Id) == 0 {
		log.Debugf("channels.getChannels - reply: {%s}", messagesChats.DebugString())
		return messagesChats, nil
	}

	for i, id := range request.Id {
		if !IsInputChannel(id) {
			log.Errorf("channels.getChannels - error: %v", mtproto.ErrChannelInvalid)
			if i == 0 && len(request.Id) == 1 {
				return nil, mtproto.ErrChannelInvalid
			}
			continue
		}

		if channel, err := s.ChannelFacade.GetMutableChannel(ctx, id.ChannelId); err != nil {
			log.Errorf("channels.getChannels - error: %v", err)
			if i == 0 && len(request.Id) == 1 {
				return nil, err
			}
		} else {
			messagesChats.Chats = append(messagesChats.Chats, channel.ToUnsafeChat(md.UserId))
			messagesChats.Count += 1
		}
	}

	log.Debugf("channels.getChannels - reply: {%s}", messagesChats.DebugString())
	return messagesChats, nil
}
