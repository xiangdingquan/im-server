package core

import (
	"context"
	"encoding/json"
	"math"

	"open.chat/app/service/biz_service/chat/dal/dataobject"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/hack"
)

func makeImmutableChatByDO(chatsDO *dataobject.ChatsDO) (chat *model.ImmutableChat) {
	chat = &model.ImmutableChat{
		Id:                  chatsDO.Id,
		Creator:             chatsDO.CreatorUserId,
		Title:               chatsDO.Title,
		Photo:               nil,
		Deactivated:         chatsDO.Deactivated,
		ParticipantsCount:   chatsDO.ParticipantCount,
		Date:                chatsDO.Date,
		Version:             chatsDO.Version,
		MigratedTo:          nil,
		DefaultBannedRights: model.ToChatBannedRights(chatsDO.DefaultBannedRights, math.MaxInt32),
		CanSetUsername:      false,
		About:               chatsDO.About,
		Notice:              chatsDO.Notice,
		ChatPhoto:           mtproto.MakeTLPhotoEmpty(nil).To_Photo(),
		Link:                chatsDO.Link,
		BotInfo:             nil,
	}

	if chatsDO.MigratedToId != 0 && chatsDO.MigratedToAccessHash != 0 {
		chat.MigratedTo = mtproto.MakeTLInputChannel(&mtproto.InputChannel{
			ChannelId:  chatsDO.MigratedToId,
			AccessHash: chatsDO.MigratedToAccessHash,
		}).To_InputChannel()
	}

	if chatsDO.Photo != "" {
		json.Unmarshal(hack.Bytes(chatsDO.Photo), chat.ChatPhoto)
	}
	chat.Photo = model.MakeChatPhotoByPhoto(chat.ChatPhoto)
	return
}

func makeImmutableChatParticipant(chat *model.ImmutableChat, chatParticipantsDO *dataobject.ChatParticipantsDO) (participant *model.ImmutableChatParticipant) {
	participant = &model.ImmutableChatParticipant{
		Id:              chatParticipantsDO.Id,
		State:           int(chatParticipantsDO.State),
		ChatParticipant: nil,
		AdminRights:     nil,
		PinnedMsgId:     chatParticipantsDO.PinnedMsgId,
		Dialog:          makeDialog(chatParticipantsDO),
		Chat:            chat,
	}

	switch chatParticipantsDO.ParticipantType {
	case model.ChatMemberNormal:
		participant.ChatParticipant = mtproto.MakeTLChatParticipant(&mtproto.ChatParticipant{
			UserId:    chatParticipantsDO.UserId,
			InviterId: chatParticipantsDO.InviterUserId,
			Date:      chatParticipantsDO.InvitedAt,
		}).To_ChatParticipant()
	case model.ChatMemberAdmin:
		participant.ChatParticipant = mtproto.MakeTLChatParticipantAdmin(&mtproto.ChatParticipant{
			UserId:    chatParticipantsDO.UserId,
			InviterId: chatParticipantsDO.InviterUserId,
			Date:      chatParticipantsDO.InvitedAt,
		}).To_ChatParticipant()
		participant.AdminRights = model.MakeDefaultChatAdminRights()
	case model.ChatMemberCreator:
		participant.ChatParticipant = mtproto.MakeTLChatParticipantCreator(&mtproto.ChatParticipant{
			UserId: chatParticipantsDO.UserId,
		}).To_ChatParticipant()
	default:
	}

	return
}

func (m *ChatCore) GetChatListByIdList(ctx context.Context, selfUserId int32, idList []int32) (chats []*mtproto.Chat) {
	chats = make([]*mtproto.Chat, 0, len(idList))

	for _, id := range idList {
		chat, err := m.GetMutableChat(ctx, id, selfUserId)
		if err != nil {
			chats = append(chats, mtproto.MakeTLChatEmpty(&mtproto.Chat{
				Id: id,
			}).To_Chat())
		} else {
			chats = append(chats, chat.ToUnsafeChat(selfUserId))
		}
	}

	return
}

func (m *ChatCore) GetChatBySelfId(ctx context.Context, selfUserId, chatId int32) (chat *mtproto.Chat) {
	mutableChat, _ := m.GetMutableChat(ctx, chatId, selfUserId)
	if mutableChat == nil {
		chat = mtproto.MakeTLChatEmpty(&mtproto.Chat{
			Id: chatId,
		}).To_Chat()
	} else {
		chat = mutableChat.ToUnsafeChat(selfUserId)
	}
	return
}

func (m *ChatCore) GetChatParticipantIdList(ctx context.Context, chatId int32) ([]int32, error) {
	doList, err := m.ChatParticipantsDAO.SelectList(ctx, chatId)
	if err != nil {
		return nil, err
	}

	idList := make([]int32, 0, len(doList))
	for i := 0; i < len(doList); i++ {
		if doList[i].State == 0 {
			idList = append(idList, doList[i].UserId)
		}
	}
	return idList, nil
}

func (m *ChatCore) getImmutableChat(ctx context.Context, chatId int32) (chat *model.ImmutableChat, err error) {
	var chatsDO *dataobject.ChatsDO
	chatsDO, err = m.ChatsDAO.Select(ctx, chatId)
	if err != nil {
		return
	} else if chatsDO == nil {
		err = mtproto.ErrChatIdInvalid
		return
	}
	chat = makeImmutableChatByDO(chatsDO)
	return
}

func (m *ChatCore) getImmutableChatParticipants(ctx context.Context, chat *model.ImmutableChat, id ...int32) (participants map[int32]*model.ImmutableChatParticipant, err error) {
	var doList []dataobject.ChatParticipantsDO
	doList, err = m.ChatParticipantsDAO.SelectList(ctx, chat.Id)
	participants = make(map[int32]*model.ImmutableChatParticipant, len(doList))
	if err != nil {
		return
	} else if len(doList) == 0 {
		return
	}

	for i := 0; i < len(doList); i++ {
		participants[doList[i].UserId] = makeImmutableChatParticipant(chat, &doList[i])
	}

	return
}

func (m *ChatCore) GetMutableChat(ctx context.Context, chatId int32, id ...int32) (chat *model.MutableChat, err error) {
	chat = new(model.MutableChat)
	chat.Chat, err = m.getImmutableChat(ctx, chatId)
	if err != nil {
		return
	}

	if len(id) > 0 {
		chat.Participants, err = m.getImmutableChatParticipants(ctx, chat.Chat, id...)
		if err != nil {
			return
		}
	}
	return
}

func (m *ChatCore) GetMutableChatListByIdList(ctx context.Context, selfUserId int32, chatId ...int32) (chats []*model.MutableChat, err error) {
	chats = make([]*model.MutableChat, 0, len(chatId))
	for _, id := range chatId {
		var (
			chat *model.MutableChat
		)
		if chat, err = m.GetMutableChat(ctx, id, selfUserId); chat != nil {
			chats = append(chats, chat)
		}
	}
	return
}

func (m *ChatCore) GetMutableChatByLink(ctx context.Context, link string, id ...int32) (chat *model.MutableChat, err error) {
	if link == "" {
		err = mtproto.ErrInviteHashEmpty
		return
	}

	var (
		chatDO       *dataobject.ChatsDO
		participants map[int32]*model.ImmutableChatParticipant
	)

	chatDO, err = m.ChatsDAO.SelectByLink(ctx, link)
	if err != nil {
		return
	} else if chatDO == nil {
		err = mtproto.ErrInviteHashInvalid
		return
	}

	immutableChat := makeImmutableChatByDO(chatDO)

	if participants, err = m.getImmutableChatParticipants(ctx, immutableChat, id...); err != nil {
		return
	}

	chat = &model.MutableChat{
		Chat:         immutableChat,
		Participants: participants,
	}

	return
}

func (m *ChatCore) MigratedToChannel(ctx context.Context, chat *model.MutableChat, id int32, accessHash int64) error {
	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		_, err := m.ChatsDAO.UpdateMigratedToTx(tx, id, accessHash, chat.Chat.Id)
		if err != nil {
			result.Err = err
			return
		}

		chat.Chat.Deactivated = 1
		chat.Chat.MigratedTo = mtproto.MakeTLInputChannel(&mtproto.InputChannel{
			ChannelId:  id,
			AccessHash: accessHash,
		}).To_InputChannel()
	})
	return tR.Err
}
