package core

import (
	"context"
	"encoding/json"
	"math"
	"math/rand"
	"time"

	"open.chat/app/pkg/env2"
	"open.chat/app/service/biz_service/chat/dal/dataobject"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
	"open.chat/pkg/random2"
)

const (
	createChatFlood = 10 // 10s
)

func (m *ChatCore) CreateChat2(ctx context.Context, creatorId int32, userIdList []int32, title string) (chat *model.MutableChat, err error) {
	var (
		chatsDO *dataobject.ChatsDO
		date    = int32(time.Now().Unix())
	)

	if chatsDO, err = m.ChatsDAO.SelectLastCreator(ctx, creatorId); err != nil {
		return
	} else if chatsDO != nil {
		if date-chatsDO.Date < createChatFlood {
			err = mtproto.NewErrFloodWaitX(date - chatsDO.Date)
			log.Errorf("createChat error: %v. lastCreate = ", err, chatsDO.Date)
			return
		}
	}

	chatsDO = &dataobject.ChatsDO{
		CreatorUserId:    creatorId,
		AccessHash:       rand.Int63(),
		ParticipantCount: int32(1 + len(userIdList)),
		Title:            title,
		PhotoId:          0,
		Version:          1,
		Date:             date,
	}
	b, _ := json.Marshal(mtproto.MakeTLChatPhotoEmpty(nil).To_ChatPhoto())
	chatsDO.Photo = hack.String(b)
	b, _ = json.Marshal(mtproto.MakeTLPhotoEmpty(nil).To_Photo())
	chatsDO.ChatPhoto = hack.String(b)

	participantDOList := make([]*dataobject.ChatParticipantsDO, 1+len(userIdList))
	for i := 0; i < len(userIdList)+1; i++ {
		if i == 0 {
			participantDOList[i] = &dataobject.ChatParticipantsDO{
				UserId:          creatorId,
				ParticipantType: model.ChatMemberCreator,
			}
		} else {
			participantDOList[i] = &dataobject.ChatParticipantsDO{
				UserId:          userIdList[i-1],
				ParticipantType: model.ChatMemberNormal,
				InviterUserId:   creatorId,
				InvitedAt:       date,
			}
		}
	}

	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		// 1. insert chat
		chatId, _, err := m.ChatsDAO.InsertTx(tx, chatsDO)
		if err != nil {
			result.Err = err
			return
		}
		chatsDO.Id = int32(chatId)
		for i := 0; i < len(participantDOList); i++ {
			participantDOList[i].ChatId = chatsDO.Id
		}

		_, _, err = m.ChatParticipantsDAO.InsertBulkTx(tx, participantDOList)
		if err != nil {
			result.Err = err
			return
		}
		return
	})

	if tR.Err != nil {
		err = tR.Err
		return
	}

	chat = &model.MutableChat{
		Chat:         makeImmutableChatByDO(chatsDO),
		Participants: make(map[int32]*model.ImmutableChatParticipant, len(participantDOList)),
	}

	for i := 0; i < len(participantDOList); i++ {
		chat.Participants[participantDOList[i].UserId] = makeImmutableChatParticipant(chat.Chat, participantDOList[i])
	}

	chat.Chat.ParticipantsCount = int32(len(participantDOList))

	return
}

func (m *ChatCore) DeleteChatUser(ctx context.Context, chatId, operatorId, deleteUserId int32) (*model.MutableChat, error) {
	var (
		now             = int32(time.Now().Unix())
		chat            *model.MutableChat
		me, deletedUser *model.ImmutableChatParticipant
		err             error
		kicked          = operatorId != deleteUserId
	)

	chat, err = m.GetMutableChat(ctx, chatId, operatorId, deleteUserId)
	if err != nil {
		return nil, err
	}

	me = chat.GetImmutableChatParticipant(operatorId)
	if me == nil {
		err = mtproto.ErrInputUserDeactivated
		return nil, err
	}

	if kicked {
		if me.State != model.ChatMemberStateNormal {
			err = mtproto.ErrPeerIdInvalid
			return nil, err
		}

		if !me.CanAdminBanUsers() {
			err = mtproto.ErrChatAdminRequired
			return nil, err
		}

		deletedUser = chat.GetImmutableChatParticipant(deleteUserId)
		if deletedUser == nil {
			// USER_NOT_PARTICIPANT
			err = mtproto.ErrUserNotParticipant
			return nil, err
		} else if deletedUser.State != model.ChatMemberStateNormal {
			err = mtproto.ErrPeerIdInvalid
			return nil, err
		}
	} else {
		// left
		deletedUser = me
		if me.State != model.ChatMemberStateNormal {
			err = mtproto.ErrPeerIdInvalid
			return nil, err
		}
	}

	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		if kicked {
			_, result.Err = m.ChatParticipantsDAO.UpdateKickedTx(tx, now, deletedUser.Id)
			if result.Err != nil {
				return
			}
			deletedUser.State = model.ChatMemberStateKicked
		} else {
			_, result.Err = m.ChatParticipantsDAO.UpdateLeft(ctx, now, deletedUser.Id)
			if result.Err != nil {
				return
			}
			deletedUser.State = model.ChatMemberStateLeft
		}
		chat.Chat.ParticipantsCount -= 1
		chat.Chat.Date = now
		chat.Chat.Version += 1
		_, result.Err = m.ChatsDAO.UpdateParticipantCount(ctx, chat.Chat.ParticipantsCount, chat.Chat.Id)
	})

	if tR.Err != nil {
		err = tR.Err
		return nil, err
	}

	return chat, nil
}

func (m *ChatCore) EditChatTitle(ctx context.Context, chatId, editUserId int32, title string) (*model.MutableChat, error) {
	var (
		now  = int32(time.Now().Unix())
		chat *model.MutableChat
		me   *model.ImmutableChatParticipant
		err  error
	)

	if title == "" {
		err = mtproto.ErrChatTitleEmpty
		return nil, err
	}

	chat, err = m.GetMutableChat(ctx, chatId, editUserId)
	if err != nil {
		return nil, err
	}
	if chat.Chat.Title == title {
		err = mtproto.ErrChatNotModified
		return nil, err
	}

	me = chat.GetImmutableChatParticipant(editUserId)
	if me == nil || me.State != model.ChatMemberStateNormal {
		err = mtproto.ErrInputUserDeactivated
		return nil, err
	}

	// 400	CHAT_ADMIN_REQUIRED	You must be an admin in this chat to do this
	if !me.CanChangeInfo() {
		err = mtproto.ErrChatAdminRequired
		return nil, err
	}

	_, err = m.ChatsDAO.UpdateTitle(ctx, title, chat.Chat.Id)
	if err != nil {
		return nil, err
	}

	chat.Chat.Title = title
	chat.Chat.Version += 1
	chat.Chat.Date = now
	return chat, nil
}

func (m *ChatCore) EditChatAbout(ctx context.Context, chatId, editUserId int32, about string) (*model.MutableChat, error) {
	var (
		now  = int32(time.Now().Unix())
		chat *model.MutableChat
		me   *model.ImmutableChatParticipant
		err  error
	)

	chat, err = m.GetMutableChat(ctx, chatId, editUserId)
	if err != nil {
		return nil, err
	}
	if chat.Chat.About == about {
		err = mtproto.ErrChatNotModified
		return nil, err
	}

	me = chat.GetImmutableChatParticipant(editUserId)
	if me == nil || me.State != model.ChatMemberStateNormal {
		err = mtproto.ErrInputUserDeactivated
		return nil, err
	}

	// 400	CHAT_ADMIN_REQUIRED	You must be an admin in this chat to do this
	if !me.CanChangeInfo() {
		err = mtproto.ErrChatAdminRequired
		return nil, err
	}

	_, err = m.ChatsDAO.UpdateAbout(ctx, about, chat.Chat.Id)
	if err != nil {
		return nil, err
	}

	chat.Chat.About = about
	chat.Chat.Version += 1
	chat.Chat.Date = now
	return chat, nil
}

func (m *ChatCore) EditChatNotice(ctx context.Context, chatId, editUserId int32, notice string) (*model.MutableChat, error) {
	var (
		now  = int32(time.Now().Unix())
		chat *model.MutableChat
		me   *model.ImmutableChatParticipant
		err  error
	)

	chat, err = m.GetMutableChat(ctx, chatId, editUserId)
	if err != nil {
		return nil, err
	}
	if chat.Chat.Notice == notice {
		err = mtproto.ErrChatNotModified
		return nil, err
	}

	me = chat.GetImmutableChatParticipant(editUserId)
	if me == nil || me.State != model.ChatMemberStateNormal {
		err = mtproto.ErrInputUserDeactivated
		return nil, err
	}

	// 400	CHAT_ADMIN_REQUIRED	You must be an admin in this chat to do this
	if !me.CanChangeInfo() {
		err = mtproto.ErrChatAdminRequired
		return nil, err
	}

	_, err = m.ChatsDAO.UpdateNotice(ctx, notice, chat.Chat.Id)
	if err != nil {
		return nil, err
	}

	chat.Chat.Notice = notice
	chat.Chat.Version += 1
	chat.Chat.Date = now
	return chat, nil
}

func (m *ChatCore) EditChatPhoto(ctx context.Context, chatId, editUserId int32, chatPhoto *mtproto.Photo) (*model.MutableChat, error) {
	var (
		now  = int32(time.Now().Unix())
		chat *model.MutableChat
		me   *model.ImmutableChatParticipant
		err  error
	)

	chat, err = m.GetMutableChat(ctx, chatId, editUserId)
	if err != nil {
		return nil, err
	}

	me = chat.GetImmutableChatParticipant(editUserId)
	if me == nil || me.State != model.ChatMemberStateNormal {
		err = mtproto.ErrInputUserDeactivated
		return nil, err
	}

	// 400	CHAT_ADMIN_REQUIRED	You must be an admin in this chat to do this
	if !me.CanChangeInfo() {
		err = mtproto.ErrChatAdminRequired
		return nil, err
	}

	m.ChatsDAO.UpdatePhoto(ctx, hack.String(model.TLObjectToJson(chatPhoto)), chatId)
	if err != nil {
		return nil, err
	}

	chat.Chat.Version += 1
	chat.Chat.Date = now
	return chat, nil
}

func (m *ChatCore) EditChatAdmin(ctx context.Context, chatId, operatorId, editChatAdminId int32, isAdmin bool) (*model.MutableChat, error) {
	var (
		now           = int32(time.Now().Unix())
		chat          *model.MutableChat
		me, editAdmin *model.ImmutableChatParticipant
		err           error
	)

	chat, err = m.GetMutableChat(ctx, chatId, operatorId, editChatAdminId)
	if err != nil {
		return nil, err
	}

	me = chat.GetImmutableChatParticipant(operatorId)
	if me == nil || me.State != model.ChatMemberStateNormal {
		err = mtproto.ErrInputUserDeactivated
		return nil, err
	}

	editAdmin = chat.GetImmutableChatParticipant(editChatAdminId)
	if editAdmin != nil && editAdmin.State != model.ChatMemberStateNormal {
		err = mtproto.ErrPeerIdInvalid
		return nil, err
	}

	if !me.CanAdminAddAdmins() {
		err = mtproto.ErrChatAdminRequired
		return nil, err
	}

	if isAdmin {
		_, err = m.ChatParticipantsDAO.UpdateParticipantType(ctx, model.ChatMemberAdmin, editAdmin.Id)
		if err != nil {
			return nil, err
		}
		editAdmin.AdminRights = model.MakeDefaultChatAdminRights()
		editAdmin.ChatParticipant.PredicateName = mtproto.Predicate_chatParticipantAdmin
	} else {
		_, err = m.ChatParticipantsDAO.UpdateParticipantType(ctx, model.ChatMemberNormal, editAdmin.Id)
		if err != nil {
			return nil, err
		}
		editAdmin.AdminRights = nil
		editAdmin.ChatParticipant.PredicateName = mtproto.Predicate_chatParticipant
	}

	chat.Chat.Version += 1
	chat.Chat.Date = now
	return chat, nil
}

func (m *ChatCore) EditChatDefaultBannedRights(ctx context.Context, chatId, operatorId int32, bannedRights *mtproto.ChatBannedRights) (*model.MutableChat, error) {
	var (
		now  = int32(time.Now().Unix())
		chat *model.MutableChat
		me   *model.ImmutableChatParticipant
		err  error
	)

	chat, err = m.GetMutableChat(ctx, chatId, operatorId)
	if err != nil {
		return nil, err
	}

	me = chat.GetImmutableChatParticipant(operatorId)
	if me == nil || me.State != model.ChatMemberStateNormal {
		err = mtproto.ErrInputUserDeactivated
		return nil, err
	}

	if me.IsChatMemberNormal() {
		err = mtproto.ErrChatAdminRequired
		return nil, err
	}

	b, d := model.FromChatBannedRights(bannedRights.To_ChatBannedRights())
	_ = d
	_, err = m.ChatsDAO.UpdateDefaultBannedRights(ctx, b, chatId)
	if err != nil {
		return nil, err
	}

	chat.Chat.DefaultBannedRights = model.ToChatBannedRights(b, math.MaxInt32)
	chat.Chat.Version += 1
	chat.Chat.Date = now
	return chat, nil
}

func (m *ChatCore) ToggleChatAdmins(ctx context.Context, chatId, operatorId int32, adminsEnabled bool) (*model.MutableChat, error) {
	var (
		now  = int32(time.Now().Unix())
		chat *model.MutableChat
		me   *model.ImmutableChatParticipant
		err  error
	)

	chat, err = m.GetMutableChat(ctx, chatId, operatorId)
	if err != nil {
		return nil, err
	}

	me = chat.GetImmutableChatParticipant(operatorId)
	if me == nil || me.State != model.ChatMemberStateNormal {
		err = mtproto.ErrInputUserDeactivated
		return nil, err
	}

	if me.IsChatMemberNormal() {
		err = mtproto.ErrChatAdminRequired
		return nil, err
	}

	if adminsEnabled {
		b, d := model.FromChatBannedRights(mtproto.MakeTLChatBannedRights(&mtproto.ChatBannedRights{
			ViewMessages: false,
			SendMessages: false,
			SendMedia:    false,
			SendStickers: false,
			SendGifs:     false,
			SendGames:    false,
			SendInline:   false,
			EmbedLinks:   false,
			SendPolls:    false,
			ChangeInfo:   false,
			InviteUsers:  false,
			PinMessages:  false,
			UntilDate:    0,
		}))
		_ = d
		_, err = m.ChatsDAO.UpdateDefaultBannedRights(ctx, b, chatId)
		if err != nil {
			return nil, err
		}
		chat.Chat.DefaultBannedRights = model.ToChatBannedRights(b, math.MaxInt32)
	} else {
		b, d := model.FromChatBannedRights(mtproto.MakeTLChatBannedRights(&mtproto.ChatBannedRights{
			ViewMessages: false,
			SendMessages: false,
			SendMedia:    false,
			SendStickers: false,
			SendGifs:     false,
			SendGames:    false,
			SendInline:   false,
			EmbedLinks:   false,
			SendPolls:    false,
			ChangeInfo:   true,
			InviteUsers:  true,
			PinMessages:  true,
			UntilDate:    0,
		}))
		_ = d
		_, err = m.ChatsDAO.UpdateDefaultBannedRights(ctx, b, chatId)
		if err != nil {
			return nil, err
		}
		chat.Chat.DefaultBannedRights = model.ToChatBannedRights(b, math.MaxInt32)
	}

	chat.Chat.Version += 1
	chat.Chat.Date = now
	return chat, nil
}

func (m *ChatCore) ExportChatInvite(ctx context.Context, chatId, inviteUserId int32) (string, error) {
	var (
		now  = int32(time.Now().Unix())
		chat *model.MutableChat
		me   *model.ImmutableChatParticipant
		err  error
	)

	chat, err = m.GetMutableChat(ctx, chatId, inviteUserId)
	if err != nil {
		return "", err
	}

	me = chat.GetImmutableChatParticipant(inviteUserId)
	if me == nil || me.State != model.ChatMemberStateNormal {
		err = mtproto.ErrInputUserDeactivated
		return "", err
	}

	if !me.CanAdminInviteUsers() {
		err = mtproto.ErrChatAdminRequired
		return "", err
	}

	link := random2.RandomAlphanumeric(22)
	_, err = m.ChatsDAO.UpdateLink(ctx, link, now, chatId)
	if err != nil {
		return "", err
	}

	return env2.T_ME + "/joinchat?link=" + link, nil
}

func (m *ChatCore) AddChatUser(ctx context.Context, chatId, inviterId, userId int32) (*model.MutableChat, error) {
	var (
		now         = int32(time.Now().Unix())
		chat        *model.MutableChat
		me, willAdd *model.ImmutableChatParticipant
		err         error
	)

	chat, err = m.GetMutableChat(ctx, chatId, inviterId, userId)
	if err != nil {
		return nil, err
	}

	me = chat.GetImmutableChatParticipant(inviterId)
	if me == nil || (me.State != model.ChatMemberStateNormal && !me.IsChatMemberCreator()) {
		err = mtproto.ErrInputUserDeactivated
		return nil, err
	}

	willAdd = chat.GetImmutableChatParticipant(userId)
	if willAdd != nil && willAdd.State == model.ChatMemberStateNormal {
		err = mtproto.ErrUserAlreadyParticipant
		return nil, err
	}

	// 400	CHAT_ADMIN_REQUIRED	You must be an admin in this chat to do this
	if !me.CanInviteUsers() {
		err = mtproto.ErrChatAdminRequired
		return nil, err
	}

	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		chatParticipantDO := &dataobject.ChatParticipantsDO{
			ChatId:          chat.Chat.Id,
			UserId:          userId,
			ParticipantType: model.ChatMemberNormal,
			InviterUserId:   inviterId,
			InvitedAt:       now,
		}
		if chat.Chat.Creator == userId {
			chatParticipantDO.ParticipantType = model.ChatMemberCreator
		}
		if willAdd == nil {
			lastInsertId, _, err := m.ChatParticipantsDAO.InsertTx(tx, chatParticipantDO)
			if err != nil {
				result.Err = err
				return
			}
			chatParticipantDO.Id = int32(lastInsertId)
			willAdd = makeImmutableChatParticipant(chat.Chat, chatParticipantDO)
		} else {
			chatParticipantDO.Id = willAdd.Id
			if _, err := m.ChatParticipantsDAO.Update(ctx, chatParticipantDO.ParticipantType, inviterId, now, chatParticipantDO.Id); err != nil {
				result.Err = err
				return
			}
		}
		chat.Chat.ParticipantsCount += 1
		chat.Chat.Version += 1
		chat.Chat.Date = now
		_, result.Err = m.ChatsDAO.UpdateParticipantCount(ctx, chat.Chat.ParticipantsCount, chat.Chat.Id)
	})

	if tR.Err != nil {
		return nil, tR.Err
	}
	return chat, nil
}

func (m *ChatCore) ImportChatInvite(ctx context.Context, hash string, userId int32) (*model.MutableChat, error) {
	var (
		chat *model.MutableChat
		err  error
	)

	chat, err = m.GetMutableChatByLink(ctx, hash, userId)
	if err != nil {
		return nil, err
	}

	return chat, nil
}

func (m *ChatCore) UpdatePinnedMessage(ctx context.Context, userId, chatId int32, chatPinnedList map[int32]int32) (*model.MutableChat, error) {
	var (
		chat *model.MutableChat
		err  error
	)

	chat, err = m.GetMutableChat(ctx, chatId, userId)
	if err != nil {
		return nil, err
	}

	sqlx.TxWrapper(ctx, m.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		for k, v := range chatPinnedList {
			m.ChatParticipantsDAO.UpdatePinnedMsgIdTx(tx, v, k, chatId)
		}
		chat.Chat.Version += 1
		_, result.Err = m.ChatsDAO.UpdateVersionTx(tx, chatId)
	})

	return chat, nil
}

func (m *ChatCore) UpdateUnPinnedMessage(ctx context.Context, userId, chatId int32) (*model.MutableChat, error) {
	var (
		chat *model.MutableChat
		err  error
	)

	chat, err = m.GetMutableChat(ctx, chatId, userId)
	if err != nil {
		return nil, err
	}

	sqlx.TxWrapper(ctx, m.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		m.ChatParticipantsDAO.UpdateUnPinnedMsgId(ctx, chatId)
		chat.Chat.Version += 1
		_, result.Err = m.ChatsDAO.UpdateVersionTx(tx, chatId)
	})
	return chat, nil
}

func (m *ChatCore) GetUsersChatIdList(ctx context.Context, id []int32) map[int32][]int32 {
	var idList2 = make(map[int32][]int32, len(id))

	chatDOList, err := m.ChatParticipantsDAO.SelectUsersChatIdList(ctx, id)
	if err != nil {
		log.Errorf("getMyChatIdList - error: %v", err)
	}
	for i := 0; i < len(chatDOList); i++ {
		idList2[chatDOList[i].UserId] = append(idList2[chatDOList[i].UserId], chatDOList[i].ChatId)
	}

	return idList2
}

func (m *ChatCore) GetMyChatList(ctx context.Context, userId int32, isCreator bool) []*mtproto.Chat {
	var chatList []*mtproto.Chat
	//
	if isCreator {
		if chatIdList, err := m.ChatParticipantsDAO.SelectMyAdminList(ctx, userId); err != nil {
			log.Errorf("getMyIsCreatorChatList error: %v", err)
		} else {
			log.Debugf("getMyIsCreatorChatList - {%v}", chatIdList)
			chatList = make([]*mtproto.Chat, 0, len(chatIdList))
			for _, id := range chatIdList {
				chat, _ := m.GetMutableChat(ctx, id, userId)
				if chat.Chat.Deactivated == 0 {
					chatList = append(chatList, chat.ToUnsafeChat(userId))
				}
			}
		}
	} else {
	}

	if chatList == nil {
		chatList = []*mtproto.Chat{}
	}

	return chatList
}

func (m *ChatCore) GetAllChats(ctx context.Context, selfUserId int32) (chats []*mtproto.Chat) {
	var chatList []*mtproto.Chat
	if chatIdList, err := m.ChatParticipantsDAO.SelectMyAllList(ctx, selfUserId); err != nil {
		log.Errorf("getAllChats error: %v", err)
	} else {
		log.Debugf("getAllChats - {%v}", chatIdList)
		chatList = make([]*mtproto.Chat, 0, len(chatIdList))
		for _, id := range chatIdList {
			chat, _ := m.GetMutableChat(ctx, id, selfUserId)
			if chat.Chat.Deactivated == 0 {
				chatList = append(chatList, chat.ToUnsafeChat(selfUserId))
			}
		}
	}

	if chatList == nil {
		chatList = []*mtproto.Chat{}
	}

	return chatList
}

func (m *ChatCore) GetFilterKeywords(ctx context.Context, id uint32) (keywords []string, err error) {
	return m.Dao.SelectChatBannedKeywords(ctx, id)
}
