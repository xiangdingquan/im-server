package chats

import (
	"context"
	"errors"
	"math"
	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"time"

	"open.chat/app/json/consts"
	"open.chat/app/json/helper"
	svc "open.chat/app/json/service"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"

	"open.chat/app/json/service/handler"

	"open.chat/app/json/services/handler/chats/core"

	channel_facade "open.chat/app/service/biz_service/channel/facade"
	chat_facade "open.chat/app/service/biz_service/chat/facade"
)

type cls struct {
	chat_facade.ChatFacade
	channel_facade.ChannelFacade
	*core.ChatsCore
}

// New .
func New(s *svc.Service) {
	service := &cls{
		ChatsCore: core.New(nil),
	}
	var err error
	service.ChatFacade, err = chat_facade.NewChatFacade("local")
	helper.CheckErr(err)
	service.ChannelFacade, err = channel_facade.NewChannelFacade("local")
	helper.CheckErr(err)
	s.AppendServices(handler.RegisterChats(service))
}

func (s *cls) checkChatMember(ctx context.Context, admin bool, chatId uint32, self int32) error {
	chat, err := s.ChatFacade.GetMutableChat(ctx, int32(chatId), self)
	if err != nil {
		return errors.New("chat id invalid")
	}
	participant := chat.GetImmutableChatParticipant(self)
	if participant == nil {
		return errors.New("not in the chat")
	}
	if admin {
		if !participant.IsChatMemberCreator() && !participant.IsChatMemberAdmin() {
			return errors.New("not a chat admin")
		}
	} else {
		if !participant.IsChatMemberStateNormal() {
			return errors.New("not a chat member")
		}
	}
	return nil
}

func (s *cls) checkChannelMember(ctx context.Context, admin bool, chatId uint32, self int32) error {
	channel, err := s.ChannelFacade.GetMutableChannel(ctx, int32(chatId), self)
	if err != nil {
		return errors.New("chat id invalid")
	}
	participant := channel.GetImmutableChannelParticipant(self)
	if participant == nil {
		return errors.New("not in the chat")
	}
	if admin {
		if !participant.IsCreatorOrAdmin() {
			return errors.New("not a chat admin")
		}
	} else {
		if !participant.IsStateOk() {
			return errors.New("not a chat member")
		}
	}
	return nil
}

func (s *cls) GetBannedRightEx(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TChatID) *helper.ResultJSON {
	if r.ChatID == 0 {
		return &helper.ResultJSON{Code: -1, Msg: "chat id invalid"}
	}
	var (
		isChannel    bool   = r.ChatID > 0
		chatId       uint32 = uint32(-r.ChatID)
		bannedRights *core.BannedRights
		err          error
	)
	if isChannel {
		chatId = uint32(r.ChatID)
		err = s.checkChannelMember(ctx, false, chatId, md.UserId)
		if err != nil {
			return &helper.ResultJSON{Code: -2, Msg: err.Error()}
		}
		bannedRights = s.ChatsCore.GetChannelBannedRights(ctx, chatId)
	} else {
		err = s.checkChatMember(ctx, false, chatId, md.UserId)
		if err != nil {
			return &helper.ResultJSON{Code: -2, Msg: err.Error()}
		}
		bannedRights = s.ChatsCore.GetChatBannedRights(ctx, chatId)
	}
	if bannedRights == nil {
		return &helper.ResultJSON{Code: -3, Msg: "fail"}
	}
	return &helper.ResultJSON{Code: 200, Msg: "success", Data: bannedRights}
}

func (s *cls) ModifyBannedRightEx(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TChatIDModifyBannedRights) *helper.ResultJSON {
	if r.ChatID == 0 {
		return &helper.ResultJSON{Code: -1, Msg: "chat id invalid"}
	}
	var err error
	isChannel := r.ChatID > 0
	chatId := uint32(-r.ChatID)
	bannedRights := &core.BannedRights{
		BanWhisper:         r.BanWhisper,
		BanSendWebLink:     r.BanSendWebLink,
		BanSendQRcode:      r.BanSendQRcode,
		BanSendKeyword:     r.BanSendKeyword,
		BanSendDmMention:   r.BanSendDmMention,
		KickWhoSendKeyword: r.KickWhoSendKeyword,
		ShowKickMessage:    r.ShowKickMessage,
	}
	if isChannel {
		chatId = uint32(r.ChatID)
		err = s.checkChannelMember(ctx, true, chatId, md.UserId)
		if err != nil {
			return &helper.ResultJSON{Code: -2, Msg: err.Error()}
		}
		err = s.ChatsCore.UpdateChannelBannedRights(ctx, chatId, bannedRights)
	} else {
		err = s.checkChatMember(ctx, true, chatId, md.UserId)
		if err != nil {
			return &helper.ResultJSON{Code: -2, Msg: err.Error()}
		}
		err = s.ChatsCore.UpdateChatBannedRights(ctx, chatId, bannedRights)
	}
	if err != nil {
		log.Errorf("ModifyBannedRightEx, update fail, error: %v", err)
		if err.Error() == "update fail" {
			return &helper.ResultJSON{Code: 201, Msg: "not changed"}
		} else {
			return &helper.ResultJSON{Code: -3, Msg: "update fail"}
		}
	}
	push := &helper.PushUpdate{
		Action: consts.ActionChatsRightOnUpdate,
		From:   uint32(md.UserId),
	}
	data := struct {
		IsChannel bool   `json:"isChannel"`
		ChatId    uint32 `json:"chatId"`
		*core.BannedRights
	}{
		IsChannel:    isChannel,
		ChatId:       chatId,
		BannedRights: bannedRights,
	}
	if isChannel {
		err = push.ToChannel(ctx, uint32(r.ChatID), data)
	} else {
		err = push.ToChat(ctx, uint32(-r.ChatID), data)
	}
	if err != nil {
		return &helper.ResultJSON{Code: -4, Msg: "push to chat fail"}
	}
	return &helper.ResultJSON{Code: 200, Msg: "success"}
}

func (s *cls) GetFilterKeywords(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TChatID) *helper.ResultJSON {
	if r.ChatID == 0 {
		return &helper.ResultJSON{Code: -1, Msg: "chat id invalid"}
	}
	var (
		isChannel bool   = r.ChatID > 0
		chatId    uint32 = uint32(-r.ChatID)
		keywords  []string
		err       error
	)
	if isChannel {
		chatId = uint32(r.ChatID)
		err = s.checkChannelMember(ctx, false, chatId, md.UserId)
		if err != nil {
			return &helper.ResultJSON{Code: -2, Msg: err.Error()}
		}
		keywords, err = s.ChatsDAO.SelectChannelBannedKeywords(ctx, chatId)
	} else {
		err = s.checkChatMember(ctx, false, chatId, md.UserId)
		if err != nil {
			return &helper.ResultJSON{Code: -2, Msg: err.Error()}
		}
		keywords, err = s.ChatsDAO.SelectChatBannedKeywords(ctx, chatId)
	}
	if err != nil {
		return &helper.ResultJSON{Code: -3, Msg: "fail"}
	}
	return &helper.ResultJSON{Code: 200, Msg: "success", Data: keywords}
}

func (s *cls) SetFilterKeywords(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TChatIDSetFilterKeywords) *helper.ResultJSON {
	if r.ChatID == 0 {
		return &helper.ResultJSON{Code: -1, Msg: "chat id invalid"}
	}
	var err error
	isChannel := r.ChatID > 0
	chatId := uint32(-r.ChatID)
	if isChannel {
		chatId = uint32(r.ChatID)
		err = s.checkChannelMember(ctx, true, chatId, md.UserId)
		if err != nil {
			return &helper.ResultJSON{Code: -2, Msg: err.Error()}
		}
		err = s.ChatsDAO.UpdateChannelKeywords(ctx, chatId, r.Keywords)
	} else {
		err = s.checkChatMember(ctx, true, chatId, md.UserId)
		if err != nil {
			return &helper.ResultJSON{Code: -2, Msg: err.Error()}
		}
		err = s.ChatsDAO.UpdateChatKeywords(ctx, chatId, r.Keywords)
	}
	if err != nil {
		return &helper.ResultJSON{Code: -3, Msg: "update fail"}
	}
	push := &helper.PushUpdate{
		Action: consts.ActionChatsKeywordsOnUpdate,
		From:   uint32(md.UserId),
	}
	data := struct {
		IsChannel bool     `json:"isChannel"`
		ChatId    uint32   `json:"chatId"`
		Keywords  []string `json:"keywords"`
	}{
		IsChannel: isChannel,
		ChatId:    chatId,
		Keywords:  r.Keywords,
	}
	if isChannel {
		err = push.ToChannel(ctx, uint32(r.ChatID), data)
	} else {
		err = push.ToChat(ctx, uint32(-r.ChatID), data)
	}
	if err != nil {
		return &helper.ResultJSON{Code: -4, Msg: "push to chat fail"}
	}
	return &helper.ResultJSON{Code: 200, Msg: "success"}
}

func (s *cls) DisableInviteLink(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TChatID) *helper.ResultJSON {
	if r.ChatID == 0 {
		return &helper.ResultJSON{Code: -1, Msg: "chat id invalid"}
	}
	var err error
	isChannel := r.ChatID > 0
	chatId := uint32(-r.ChatID)
	now := int32(time.Now().Unix())
	if isChannel {
		chatId = uint32(r.ChatID)
		err = s.checkChannelMember(ctx, true, chatId, md.UserId)
		if err != nil {
			return &helper.ResultJSON{Code: -2, Msg: err.Error()}
		}
		err = s.ClearChannelLink(ctx, now, int32(chatId))
	} else {
		err = s.checkChatMember(ctx, true, chatId, md.UserId)
		if err != nil {
			return &helper.ResultJSON{Code: -2, Msg: err.Error()}
		}
		err = s.ClearChatLink(ctx, now, int32(chatId))
	}
	if err != nil {
		return &helper.ResultJSON{Code: -3, Msg: "disable fail"}
	}
	return &helper.ResultJSON{Code: 200, Msg: "success"}
}

func (s *cls) KickWhoSendKeyword(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TChatID) *helper.ResultJSON {
	if r.ChatID == 0 {
		return &helper.ResultJSON{Code: -1, Msg: "chat id invalid"}
	}
	var err error
	isChannel := r.ChatID > 0
	whoSend := md.UserId

	log.Debugf("KickWhoSendKeyword, whoSend:%d, chat:%d", whoSend, r.ChatID)

	var creator int32
	var chatId int32
	if isChannel {
		channel, err := s.ChannelFacade.GetMutableChannel(ctx, r.ChatID)
		if err != nil {
			log.Errorf("KickWhoSendKeyword, GetMutableChannel, error: %v", err)
			return &helper.ResultJSON{Code: -2, Msg: "internal error"}
		}
		creator = channel.Channel.CreatorId
		chatId = channel.Channel.Id
	} else {
		chat, err := s.ChatFacade.GetMutableChat(ctx, -r.ChatID)
		if err != nil {
			log.Errorf("KickWhoSendKeyword, GetMutableChat, error: %v", err)
			return &helper.ResultJSON{Code: -3, Msg: "internal error"}
		}
		creator = chat.Chat.Creator
		chatId = chat.Chat.Id
	}
	log.Debug("KickWhoSendKeyword, creator:%d, chatId:%d", creator, chatId)

	if creator == whoSend {
		return &helper.ResultJSON{Code: -4, Msg: "can not kick creator"}
	}

	var channel *model.MutableChannel
	if creator == 0 {
		channel, err = s.ChannelFacade.LeaveChannel(ctx, chatId, whoSend)
		if err != nil {
			log.Errorf("KickWhoSendKeyword, LeaveChannel, error: %v", err)
			return &helper.ResultJSON{Code: -5, Msg: "internal error"}
		}
		log.Debug("KickWhoSendKeyword, LeaveChannel")
	} else {
		deleted := false
		bannedRights := model.ChatBannedRights{
			Rights:    math.MaxInt32,
			UntilDate: math.MaxInt32,
		}
		log.Debug("KickWhoSendKeyword, try EditBanned")
		channel, deleted, err = s.ChannelFacade.EditBanned(ctx, chatId, creator, whoSend, bannedRights)
		log.Debug("KickWhoSendKeyword, EditBanned")

		if err != nil {
			log.Errorf("KickWhoSendKeyword, EditBanned, error: %v", err)
			return &helper.ResultJSON{Code: -5, Msg: "internal error"}
		}
		log.Debugf("KickWhoSendKeyword, deleted:%b", deleted)
	}
	go func() {
		sync_client.PushUpdates(context.Background(), whoSend, mtproto.MakeTLUpdates(&mtproto.Updates{
			Updates: []*mtproto.Update{model.MakeUpdateChannel(channel.GetChannelId())},
			Users:   []*mtproto.User{},
			Chats:   []*mtproto.Chat{channel.ToUnsafeChat(whoSend)},
			Date:    int32(time.Now().Unix()),
			Seq:     0,
		}).To_Updates())
	}()

	func() {
		var bannedRights *core.BannedRights
		if isChannel {
			bannedRights = s.ChatsCore.GetChannelBannedRights(ctx, uint32(r.ChatID))
		} else {
			bannedRights = s.ChatsCore.GetChatBannedRights(ctx, uint32(-r.ChatID))
		}
		log.Debugf("KickWhoSendKeyword, bannedRights:%v", bannedRights)
		if bannedRights != nil && bannedRights.ShowKickMessage {
			sender := helper.MakeSender(uint32(whoSend), md.GetAuthId(), 6, 100)
			msg := struct {
				UID int32 `json:"UID"`
			}{UID: whoSend}
			sender.SendToChannel(ctx, uint32(chatId), &msg)
		}
	}()

	//chat, err := s.ChatFacade.DeleteChatUser(ctx, chatId, creator, whoSend)
	//if err != nil {
	//	return false, err
	//}
	//
	//log.Debug("sendMessage - kickWhoSendBanWord, user deleted")
	//
	//helper.MakeSender(uint32(whoSend), auth, 6, 100)
	//
	//go func() {
	//	updateChatParticipants := mtproto.MakeTLUpdateChatParticipants(&mtproto.Update{
	//		Participants: chat.ToChatParticipants(0),
	//	}).To_Update()
	//
	//	updatesHelper := model.MakeUpdatesHelper(updateChatParticipants)
	//
	//	chat.Walk(func(userId int32, participant *model.ImmutableChatParticipant) error {
	//		sync_client.PushUpdates(ctx, userId,
	//			updatesHelper.ToPushUpdates(context.Background(), userId, s.UserFacade, s.ChatFacade, s.ChannelFacade))
	//		return nil
	//	})
	//
	//}()

	return &helper.ResultJSON{Code: 200, Msg: "success"}
}

func (s *cls) SetNickname(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TChatIdSetNickname) *helper.ResultJSON {
	if r.ChatID == 0 {
		return &helper.ResultJSON{Code: -1, Msg: "chat id invalid"}
	}
	var err error

	if s.isChannel(r.ChatID) {
		err = s.ChatsCore.ChatsDAO.UpdateChannelNickname(ctx, s.getChannelId(r.ChatID), uint32(md.UserId), r.Nickname)
	} else {
		err = s.ChatsCore.ChatsDAO.UpdateChatNickname(ctx, s.getChatId(r.ChatID), uint32(md.UserId), r.Nickname)
	}

	if err != nil {
		return &helper.ResultJSON{Code: -1, Msg: "internal error"}
	}

	err = s.pushUpdate(ctx, r.ChatID, md.UserId, consts.ActionChatsNicknameOnUpdate, struct {
		Nickname string `json:"nickname"`
		ChatId   int32  `json:"chatId"`
		UserId   int32  `json:"userId"`
	}{
		Nickname: r.Nickname,
		ChatId:   r.ChatID,
		UserId:   md.UserId,
	})

	if err != nil {
		return &helper.ResultJSON{Code: -2, Msg: "internal error"}
	}

	return &helper.ResultJSON{Code: 200, Msg: "success"}
}

func (s *cls) pushUpdate(ctx context.Context, chatId, from int32, action consts.OnAction, data interface{}) (err error) {
	push := &helper.PushUpdate{
		Action: action,
		From:   uint32(from),
	}
	if s.isChannel(chatId) {
		err = push.ToChannel(ctx, s.getChannelId(chatId), data)
	} else {
		err = push.ToChat(ctx, s.getChatId(chatId), data)
	}
	return
}

func (s *cls) isChannel(chatId int32) bool {
	return chatId > 0
}

func (s *cls) getChatId(chatId int32) uint32 {
	return uint32(-chatId)
}

func (s *cls) getChannelId(chatId int32) uint32 {
	return uint32(chatId)
}
