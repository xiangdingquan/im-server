package model

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/log"
)

type MessageBoxList []*MessageBox

func (m MessageBoxList) ToMessagesPeersList(ctx context.Context, selfUserId int32, cb1 UserHelper, cb2 ChatHelper, cb3 ChannelHelper) (messages []*mtproto.Message, users []*mtproto.User, chats []*mtproto.Chat) {
	var (
		userIdList, chatIdList, channelIdList IDList
	)
	userIdList.AddIfNot(selfUserId)
	messages = make([]*mtproto.Message, 0, len(m))
	for _, box := range m {
		userIdList.AddIfNot(box.SelfUserId)
		messages = append(messages, MessageUpdate(box.ToMessage(selfUserId)))
	}
	var l1 IDList
	l1, chatIdList, channelIdList = PickAllIdListByMessages(messages)
	userIdList.AddIfNot(l1...)

	//log.Warnf("ToMessagesPeersList:[%v,%v,%v]", userIdList, chatIdList, channelIdList)

	users = make([]*mtproto.User, 0, len(userIdList))
	if cb1 != nil {
		mutableUsers := cb1.GetMutableUsers(ctx, userIdList...)
		for _, id := range userIdList {
			user, _ := mutableUsers.ToUnsafeUser(selfUserId, id)
			if user != nil {
				users = append(users, user)
			}
		}
	}

	chats = make([]*mtproto.Chat, 0, len(chatIdList)+len(channelIdList))
	if cb2 != nil {
		for _, id := range chatIdList {
			if mutableChat, err := cb2.GetMutableChat(ctx, id, selfUserId); err != nil {
				log.Errorf("getMutableChat - not found chat (%d), error: %v", id, err)
			} else {
				if chat := mutableChat.ToUnsafeChat(selfUserId); chat != nil {
					chats = append(chats, chat)
				}
			}
		}
	}

	if cb3 != nil {
		for _, id := range channelIdList {
			if mutableChannel, err := cb3.GetMutableChannel(ctx, id, selfUserId); err != nil {
				log.Errorf("getMutableChannel - not found chat (%d), error: %v", id, err)
			} else {
				if channel := mutableChannel.ToUnsafeChat(selfUserId); channel != nil {
					chats = append(chats, channel)
				}
			}
		}
	}
	return
}
