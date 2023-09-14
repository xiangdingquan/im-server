package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

func (s *Service) MessagesGetCommonChats(ctx context.Context, request *mtproto.TLMessagesGetCommonChats) (*mtproto.Messages_Chats, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getCommonChats - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.getCommonChats - error: %v", err)
		return nil, err
	}

	var (
		userId              = request.GetUserId().GetUserId()
		total_count   int32 = 0
		chats               = []*mtproto.Chat{}
		messagesChats *mtproto.Messages_Chats
	)

	// 400	MSG_ID_INVALID	Invalid message ID provided
	// 400	USER_ID_INVALID	The provided user ID is invalid

	if userId == md.UserId {
		err := mtproto.ErrUserIdInvalid
		log.Errorf("messages.getCommonChats - error: %v", err)
		return nil, err
	} else {
		find_commonChats := func(me, user []int32) (int32, []int32) {
			commonChats := util.Int32Intersect(me, user)
			count := int32(len(commonChats))
			log.Debugf("messages.getCommonChats - commonChats: %v", commonChats)
			findChats := make([]int32, 0)
			for i, id := range commonChats {
				if id > request.MaxId {
					findChats = commonChats[i:]
					break
				}
			}
			return count, findChats
		}

		chatIdList := s.ChatFacade.GetUsersChatIdList(ctx, []int32{md.UserId, request.UserId.GetUserId()})
		if len(chatIdList) == 2 {
			total, commonChats := find_commonChats(chatIdList[md.UserId], chatIdList[request.UserId.GetUserId()])
			if len(commonChats) > 0 {
				chats = append(chats, s.ChatFacade.GetChatListByIdList(ctx, md.UserId, commonChats)...)
			}
			total_count += total
		}

		chatIdList = s.ChannelFacade.GetUsersChannelIdList(ctx, []int32{md.UserId, request.UserId.GetUserId()})
		if len(chatIdList) == 2 {
			total, commonChats := find_commonChats(chatIdList[md.UserId], chatIdList[request.UserId.GetUserId()])
			if len(commonChats) > 0 {
				chats = append(chats, s.ChannelFacade.GetChannelListByIdList(ctx, md.UserId, commonChats...)...)
			}
			total_count += total
		}
	}

	if len(chats) > 0 {
		if int(request.Limit) < len(chats) {
			chats = chats[:request.Limit]
		}
		messagesChats = mtproto.MakeTLMessagesChats(&mtproto.Messages_Chats{
			Chats: chats,
		}).To_Messages_Chats()
	} else {
		messagesChats = mtproto.MakeTLMessagesChatsSlice(&mtproto.Messages_Chats{
			Count: total_count,
			Chats: []*mtproto.Chat{},
		}).To_Messages_Chats()
	}

	log.Debugf("messages.getCommonChats - reply: %s", messagesChats.DebugString())
	return messagesChats, nil
}
