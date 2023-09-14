package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsGetGroupsForDiscussion(ctx context.Context, request *mtproto.TLChannelsGetGroupsForDiscussion) (*mtproto.Messages_Chats, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.getGroupsForDiscussion - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("channels.getGroupsForDiscussion - error: %v", err)
		return nil, err
	}

	var chatList []*mtproto.Chat

	chatList = append(chatList, s.ChatFacade.GetMyChatList(ctx, md.UserId, true)...)

	channelList := s.ChannelFacade.GetMyAdminChannelList(ctx, md.UserId)
	for _, mChannel := range channelList {
		if mChannel.Channel.IsMegagroup() {
			chatList = append(chatList, mChannel.ToUnsafeChat(md.UserId))
		}
	}

	if chatList == nil {
		chatList = []*mtproto.Chat{}
	}

	result := mtproto.MakeTLMessagesChats(&mtproto.Messages_Chats{
		Chats: chatList,
	}).To_Messages_Chats()

	log.Debugf("channels.getGroupsForDiscussion - reply %s", result.DebugString())
	return result, nil
}
