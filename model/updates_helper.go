package model

import (
	"context"
	"time"

	"open.chat/mtproto"
	"open.chat/pkg/log"
)

const (
	BroadcastTypeChat          = 1
	BroadcastTypeChannel       = 2
	BroadcastTypeChannelAdmins = 3
)

var (
	UpdatesTooLong = mtproto.MakeTLUpdatesTooLong(nil).To_Updates()
)

func MakeEmptyUpdates() *mtproto.Updates {
	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates()
}

// ////////////////////////////////////////////////////////////////////////////
type UserUpdates struct {
	UserId int32
	*mtproto.Updates
}

func MakeUserUpdates(userId int32, updates *mtproto.Updates) *UserUpdates {
	return &UserUpdates{
		UserId:  userId,
		Updates: updates,
	}
}

type UserUpdatesList []*UserUpdates

// ////////////////////////////////////////////////////////////////////////////
type UpdateList []*mtproto.Update

func (m UpdateList) PickAllIdList(selfUserId int32) (users, chats, channels IDList) {
	users.AddIfNot(selfUserId)
	messages := make([]*mtproto.Message, 0, len(m))
	for _, m2 := range m {
		msg := m2.Message_MESSAGE
		if msg == nil {
			continue
		}
		messages = append(messages, m2.Message_MESSAGE)
	}
	var l1 IDList
	l1, chats, channels = PickAllIdListByMessages(messages)
	users.AddIfNot(l1...)
	return
}

// ////////////////////////////////////////////////////////////////////////////
func MakeUserUpdatesList(sz int) UserUpdatesList {
	return make([]*UserUpdates, 0, sz)
}

func (m UserUpdatesList) Add(userId int32, updates *mtproto.Updates) UserUpdatesList {
	return append(m, MakeUserUpdates(userId, updates))
}

type UpdatesHelper struct {
	updates       UpdateList
	cacheUsers    MutableUsers
	cacheChats    []*MutableChat
	cacheChannels []*MutableChannel
	date          int32
}

func MakeUpdatesHelper(update ...*mtproto.Update) *UpdatesHelper {
	return &UpdatesHelper{
		date:    int32(time.Now().Unix()),
		updates: update,
	}
}

func (m *UpdatesHelper) DebugString() string {
	return ""
}

func (m *UpdatesHelper) GetUpdates() *mtproto.Updates {
	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: m.updates,
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    m.date,
		Seq:     0,
	}).To_Updates()
}

func (m *UpdatesHelper) SetUsers(users MutableUsers) {
	m.cacheUsers = users
}

func (m *UpdatesHelper) PushBackUpdate(update ...*mtproto.Update) {
	m.updates = append(m.updates, update...)
}

func (m *UpdatesHelper) PushFrontUpdate(update ...*mtproto.Update) {
	m.updates = append(update, m.updates...)
}

func (m *UpdatesHelper) PushChannel(channel ...*MutableChannel) {
	m.cacheChannels = append(m.cacheChannels, channel...)
}

func (m *UpdatesHelper) PushChat(chat ...*MutableChat) {
	m.cacheChats = append(m.cacheChats, chat...)
}

func (m *UpdatesHelper) getCacheChat(ctx context.Context, selfUserId, id int32, cb ChatHelper) (*mtproto.Chat, error) {
	var (
		mutableChat *MutableChat
		err         error
	)

	for _, c := range m.cacheChats {
		if c.Chat.Id == id {
			mutableChat = c
			break
		}
	}
	if mutableChat == nil {
		if mutableChat, err = cb.GetMutableChat(ctx, id, selfUserId); err != nil {
			return nil, err
		}
		m.cacheChats = append(m.cacheChats, mutableChat)
	}
	return mutableChat.ToUnsafeChat(selfUserId), nil
}

func (m *UpdatesHelper) getCacheChannel(ctx context.Context, selfUserId, id int32, cb ChannelHelper) (*mtproto.Chat, error) {
	var (
		mutableChannel *MutableChannel
		err            error
	)

	for _, c := range m.cacheChannels {
		if c.GetId() == id {
			mutableChannel = c
			break
		}
	}
	if mutableChannel == nil {
		if mutableChannel, err = cb.GetMutableChannel(ctx, id, selfUserId); err != nil {
			return nil, err
		}
		m.cacheChannels = append(m.cacheChannels, mutableChannel)
	}
	return mutableChannel.ToUnsafeChat(selfUserId), nil
}

func (m *UpdatesHelper) pickAllIdList(selfUserId int32) (userIdList, chatIdList, channelIdList IDList) {
	userIdList, chatIdList, channelIdList = m.updates.PickAllIdList(selfUserId)
	for _, c := range m.cacheChats {
		chatIdList.AddIfNot(c.Chat.Id)
	}

	for _, c := range m.cacheChannels {
		channelIdList.AddIfNot(c.GetId())
	}

	return
}

func (m *UpdatesHelper) ToReplyUpdates(ctx context.Context, selfUserId int32, cb1 UserHelper, cb2 ChatHelper, cb3 ChannelHelper) (updates *mtproto.Updates) {
	var (
		userIdList, chatIdList, channelIdList = m.pickAllIdList(selfUserId)
		updateList                            = make([]*mtproto.Update, 0, len(m.updates))
	)

	for _, upd := range m.updates {
		switch upd.PredicateName {
		case mtproto.Predicate_updateNewChannelMessage:
			switch upd.Message_MESSAGE.PredicateName {
			case mtproto.Predicate_messageService:
				switch upd.Message_MESSAGE.Action.PredicateName {
				case mtproto.Predicate_messageActionChannelCreate:
					updateList = append(updateList, MakeUpdateChannel(upd.Message_MESSAGE.ToId.ChannelId))
				}
			}
		}
	}

	for _, upd := range m.updates {
		switch upd.PredicateName {
		case mtproto.Predicate_updateNewMessage,
			mtproto.Predicate_updateNewChannelMessage,
			mtproto.Predicate_updateNewScheduledMessage:
			updateList = append(updateList, mtproto.MakeTLUpdateMessageID(&mtproto.Update{
				Id_INT32: upd.Message_MESSAGE.Id,
				RandomId: upd.RandomId,
			}).To_Update())
		case mtproto.Predicate_updateNewBlog:
			updateList = append(updateList, mtproto.MakeTLUpdateBlogID(&mtproto.Update{
				Id_INT32: upd.Blog.Id,
				RandomId: upd.RandomId,
			}).To_Update())
		}
	}

	for _, upd := range m.updates {
		switch upd.PredicateName {
		case mtproto.Predicate_updateNewChannelMessage:
			updateList = append(updateList, mtproto.MakeTLUpdateReadChannelInbox(&mtproto.Update{
				FolderId:         nil,
				ChannelId:        upd.Message_MESSAGE.ToId.ChannelId,
				MaxId:            upd.Message_MESSAGE.Id,
				StillUnreadCount: 0,
				Pts_INT32:        upd.Pts_INT32,
			}).To_Update())
		}
		updateList = append(updateList, upd)
	}

	users := make([]*mtproto.User, 0, len(userIdList))
	if cb1 != nil {
		mutableUsers := cb1.GetMutableUsers(ctx, userIdList...)
		for _, id := range userIdList {
			if user := GetFirstValue(mutableUsers.ToUnsafeUser(selfUserId, id)).(*mtproto.User); user != nil {
				users = append(users, user)
			}
		}
	}

	chats := make([]*mtproto.Chat, 0, len(chatIdList)+len(channelIdList))
	if cb2 != nil {
		for _, id := range chatIdList {
			if chat, err := m.getCacheChat(ctx, selfUserId, id, cb2); err != nil {
				log.Errorf("getCacheChat error: %v", err)
			} else {
				chats = append(chats, chat)
			}
		}
	}

	if cb3 != nil {
		for _, id := range channelIdList {
			if channel, err := m.getCacheChannel(ctx, selfUserId, id, cb3); err != nil {
				log.Errorf("getCacheChannel error: %v", err)
			} else {
				chats = append(chats, channel)
			}
		}
	}

	updates = mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: updateList,
		Users:   users,
		Chats:   chats,
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates()
	return
}

func (m *UpdatesHelper) ToSyncNotMeUpdates(ctx context.Context, selfUserId int32, cb1 UserHelper, cb2 ChatHelper, cb3 ChannelHelper) (updates *mtproto.Updates) {
	var (
		userIdList, chatIdList, channelIdList = m.pickAllIdList(selfUserId)
		updateList                            = make([]*mtproto.Update, 0, len(m.updates))
	)

	for _, upd := range m.updates {
		switch upd.PredicateName {
		case mtproto.Predicate_updateNewChannelMessage:
			switch upd.Message_MESSAGE.PredicateName {
			case mtproto.Predicate_messageService:
				switch upd.Message_MESSAGE.Action.PredicateName {
				case mtproto.Predicate_messageActionChannelCreate:
					updateList = append(updateList, MakeUpdateChannel(upd.Message_MESSAGE.ToId.ChannelId))
				}
			}
		}
	}

	for _, upd := range m.updates {
		switch upd.PredicateName {
		case mtproto.Predicate_updateNewChannelMessage:
			updateList = append(updateList, mtproto.MakeTLUpdateReadChannelInbox(&mtproto.Update{
				FolderId:         nil,
				ChannelId:        upd.Message_MESSAGE.ToId.ChannelId,
				MaxId:            upd.Message_MESSAGE.Id,
				StillUnreadCount: 0,
				Pts_INT32:        upd.Pts_INT32,
			}).To_Update())
		}
		updateList = append(updateList, upd)
	}

	users := make([]*mtproto.User, 0, len(userIdList))
	if cb1 != nil {
		mutableUsers := cb1.GetMutableUsers(ctx, userIdList...)
		for _, id := range userIdList {
			if user := GetFirstValue(mutableUsers.ToUnsafeUser(selfUserId, id)).(*mtproto.User); user != nil {
				users = append(users, user)
			}
		}
	}

	chats := make([]*mtproto.Chat, 0, len(chatIdList)+len(channelIdList))
	if cb2 != nil {
		for _, id := range chatIdList {
			if chat, err := m.getCacheChat(ctx, selfUserId, id, cb2); err != nil {
				log.Errorf("getCacheChat error: %v", err)
			} else {
				chats = append(chats, chat)
			}
		}
	}

	if cb3 != nil {
		for _, id := range channelIdList {
			if channel, err := m.getCacheChannel(ctx, selfUserId, id, cb3); err != nil {
				log.Errorf("getCacheChannel error: %v", err)
			} else {
				chats = append(chats, channel)
			}
		}
	}

	updates = mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: updateList,
		Users:   users,
		Chats:   chats,
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates()
	return
}

func (m *UpdatesHelper) ToPushUpdates(ctx context.Context, selfUserId int32, cb1 UserHelper, cb2 ChatHelper, cb3 ChannelHelper) (updates *mtproto.Updates) {
	var (
		userIdList, chatIdList, channelIdList = m.pickAllIdList(selfUserId)
		updateList                            = make([]*mtproto.Update, 0, len(m.updates))
	)

	for _, upd := range m.updates {
		switch upd.PredicateName {
		case mtproto.Predicate_updateNewChannelMessage:
			switch upd.Message_MESSAGE.PredicateName {
			case mtproto.Predicate_message:
				upd.Message_MESSAGE.Mentioned = CheckHasMention(upd.Message_MESSAGE.Entities, selfUserId)
			case mtproto.Predicate_messageService:
				switch upd.Message_MESSAGE.Action.PredicateName {
				case mtproto.Predicate_messageActionChatAddUser:
					updateList = append(updateList, MakeUpdateChannel(upd.Message_MESSAGE.ToId.ChannelId))
				}
			}
		case mtproto.Predicate_updateChannel:
			channelIdList.AddIfNot(upd.ChannelId)
		case mtproto.Predicate_updateNewBlog:
			userIdList.AddIfNot(upd.Blog.GetUserId())
		case mtproto.Predicate_updateBlogFollow, mtproto.Predicate_updateBlogLike:
			userIdList.AddIfNot(upd.UserId)
		}
	}

	updateList = append(updateList, m.updates...)

	users := make([]*mtproto.User, 0, len(userIdList))
	if cb1 != nil {
		mutableUsers := cb1.GetMutableUsers(ctx, userIdList...)
		for _, id := range userIdList {
			user, _ := mutableUsers.ToUnsafeUser(selfUserId, id)
			if user != nil {
				users = append(users, user)
			}
		}
	}

	chats := make([]*mtproto.Chat, 0, len(chatIdList)+len(channelIdList))
	if cb2 != nil {
		for _, id := range chatIdList {
			if chat, err := m.getCacheChat(ctx, selfUserId, id, cb2); err != nil {
				log.Errorf("getCacheChat error: %v", err)
			} else {
				chats = append(chats, chat)
			}
		}
	}

	if cb3 != nil {
		for _, id := range channelIdList {
			if channel, err := m.getCacheChannel(ctx, selfUserId, id, cb3); err != nil {
				log.Errorf("getCacheChannel error: %v", err)
			} else {
				chats = append(chats, channel)
			}
		}
	}

	updates = mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: updateList,
		Users:   users,
		Chats:   chats,
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates()
	return
}

func MakeUpdatesByUpdates(updates ...*mtproto.Update) *mtproto.Updates {
	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: updates,
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates()
}

func MakeUpdatesByUpdatesUsers(users []*mtproto.User, updates ...*mtproto.Update) *mtproto.Updates {
	if users == nil {
		users = []*mtproto.User{}
	}
	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: updates,
		Users:   users,
		Chats:   []*mtproto.Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates()
}

func MakeUpdatesByUpdatesUsersChats(users []*mtproto.User, chats []*mtproto.Chat, updates ...*mtproto.Update) *mtproto.Updates {
	if users == nil {
		users = []*mtproto.User{}
	}
	if chats == nil {
		chats = []*mtproto.Chat{}
	}

	return mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: updates,
		Users:   users,
		Chats:   chats,
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates()
}
