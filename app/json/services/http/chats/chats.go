package chats

import (
	"context"

	"open.chat/app/json/helper"
	svc "open.chat/app/json/service"
	"open.chat/mtproto"

	wallet_dao "open.chat/app/json/services/handler/wallet/dao"

	"open.chat/app/json/service/http"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"open.chat/app/messenger/biz_server/channels"
	"open.chat/app/messenger/biz_server/messages/chat"
	"open.chat/app/messenger/biz_server/messages/message"
	sync_client "open.chat/app/messenger/sync/client"
	channel_facade "open.chat/app/service/biz_service/channel/facade"
	chat_facade "open.chat/app/service/biz_service/chat/facade"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type cls struct {
	Wallet *wallet_dao.Dao
	chat_facade.ChatFacade
	channel_facade.ChannelFacade
	mtproto.RPCMessagesChatServer
	mtproto.RPCChannelsServer
	mtproto.RPCMessagesMessageServer
}

// New .
func New(s *svc.Service, rg *bm.RouterGroup) {
	service := &cls{
		Wallet:                   wallet_dao.New(),
		RPCMessagesChatServer:    chat.New(),
		RPCChannelsServer:        channels.New(),
		RPCMessagesMessageServer: message.New(),
	}
	var err error
	service.ChatFacade, err = chat_facade.NewChatFacade("local")
	checkErr(err)
	service.ChannelFacade, err = channel_facade.NewChannelFacade("local")
	checkErr(err)
	http.RegisterChats(service, rg)
}

func (s *cls) Disband(ctx context.Context, r *http.TChatsDisband) *helper.ResultJSON {
	var (
		chatId, channelId, ownerId uint32
	)
	if r.ChatID < 0 {
		chatId = uint32(-r.ChatID)
		chat, err := s.ChatFacade.GetMutableChat(ctx, int32(chatId))
		if err != nil {
			return &helper.ResultJSON{Code: -1, Msg: err.Error()}
		}
		if chat.Chat != nil {
			return &helper.ResultJSON{Code: -2, Msg: "chat is not exist"}
		}
		if chat.Chat.MigratedTo != nil {
			channelId = uint32(chat.Chat.MigratedTo.GetChannelId())
		}
		ownerId = uint32(chat.Chat.Creator)
	} else {
		channelId = uint32(r.ChatID)
		chat, err := s.ChannelFacade.GetMutableChannel(ctx, int32(channelId))
		if err != nil {
			return &helper.ResultJSON{Code: -3, Msg: err.Error()}
		}
		if chat.Channel == nil {
			return &helper.ResultJSON{Code: -4, Msg: "chat is not exist"}
		}
		chatId = uint32(chat.Channel.MigratedFromChatId)
		ownerId = uint32(chat.Channel.CreatorId)
	}

	ctx, err := helper.DefaultMetadata(ctx, ownerId, 0)
	if err != nil {
		return &helper.ResultJSON{Code: -5, Msg: err.Error()}
	}
	//解散超级群
	if channelId != 0 {
		request := &mtproto.TLChannelsDeleteChannel{
			Channel: mtproto.MakeTLInputChannel(&mtproto.InputChannel{ChannelId: int32(channelId)}).To_InputChannel(),
		}
		_, err := s.RPCChannelsServer.ChannelsDeleteChannel(ctx, request)
		if err != nil {
			return &helper.ResultJSON{Code: -6, Msg: err.Error()}
		}
		//sync_client.SyncUpdatesMe(ctx, (int32)(ownerId), 0, 0, "", updates)
	}

	//解散群
	if chatId != 0 {
		request := &mtproto.TLMessagesDeleteChatUser{
			ChatId: int32(chatId),
			UserId: mtproto.MakeTLInputUser(&mtproto.InputUser{UserId: int32(ownerId)}).To_InputUser(),
		}
		updates, err := s.RPCMessagesChatServer.MessagesDeleteChatUser(ctx, request)
		if err != nil {
			return &helper.ResultJSON{Code: -7, Msg: err.Error()}
		}
		sync_client.PushUpdates(ctx, (int32)(ownerId), updates)
	}

	return &helper.ResultJSON{Code: 0, Msg: "success"}
}

func (s *cls) DeleteDialog(ctx context.Context, r *http.TChatsDeleteDialog) *helper.ResultJSON {
	ctx, err := helper.DefaultMetadata(ctx, r.UserId, 0)
	if err != nil {
		return &helper.ResultJSON{Code: -1, Msg: err.Error()}
	}
	request := &mtproto.TLMessagesDeleteHistory{
		Revoke: true,
		Peer: mtproto.MakeTLInputPeerUser(&mtproto.InputPeer{
			UserId: int32(r.TargetUid),
		}).To_InputPeer(),
		MaxId: 2147483647,
	}
	affectedHistory, err := s.MessagesDeleteHistory(ctx, request)
	if err != nil {
		return &helper.ResultJSON{Code: -2, Msg: err.Error()}
	}
	_ = affectedHistory
	return &helper.ResultJSON{Code: 0, Msg: "success"}
}

func (s *cls) Join(ctx context.Context, r *http.TChatAndUser) *helper.ResultJSON {
	if r.ChatID <= 0 {
		return &helper.ResultJSON{Code: -1, Msg: "chat id invalid"}
	}
	go func() {
		for _, uid := range r.UserIds {
			ctx, _ := helper.DefaultMetadata(context.TODO(), uid, 0)
			s.RPCChannelsServer.ChannelsJoinChannel(ctx, &mtproto.TLChannelsJoinChannel{
				Channel: mtproto.MakeTLInputChannel(&mtproto.InputChannel{
					ChannelId: int32(r.ChatID),
				}).To_InputChannel(),
			})
		}
	}()
	return &helper.ResultJSON{Code: 0, Msg: "success"}
}

func (s *cls) Leave(ctx context.Context, r *http.TChatAndUser) *helper.ResultJSON {
	if r.ChatID <= 0 {
		return &helper.ResultJSON{Code: -1, Msg: "chat id invalid"}
	}
	go func() {
		for _, uid := range r.UserIds {
			ctx, _ := helper.DefaultMetadata(context.TODO(), uid, 0)
			s.RPCChannelsServer.ChannelsLeaveChannel(ctx, &mtproto.TLChannelsLeaveChannel{
				Channel: mtproto.MakeTLInputChannel(&mtproto.InputChannel{
					ChannelId: int32(r.ChatID),
				}).To_InputChannel(),
			})
		}
	}()
	return &helper.ResultJSON{Code: 0, Msg: "success"}
}
