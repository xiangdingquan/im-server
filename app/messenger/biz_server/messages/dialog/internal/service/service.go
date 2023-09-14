package service

import (
	"context"
	"sort"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"

	sync_client "open.chat/app/messenger/sync/client"
	channel_facade "open.chat/app/service/biz_service/channel/facade"
	chat_facade "open.chat/app/service/biz_service/chat/facade"
	message_facade "open.chat/app/service/biz_service/message/facade"
	private_facade "open.chat/app/service/biz_service/private/facade"
	report_facade "open.chat/app/service/biz_service/report/facade"
	updates_facade "open.chat/app/service/biz_service/updates/facade"
	user_client "open.chat/app/service/biz_service/user/client"
	media_client "open.chat/app/service/media/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

type Service struct {
	user_client.UserFacade
	updates_facade.UpdatesFacade
	report_facade.ReportFacade
	private_facade.PrivateFacade
	chat_facade.ChatFacade
	channel_facade.ChannelFacade
	message_facade.MessageFacade
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func New() *Service {
	var (
		ac struct {
			c *warden.ClientConfig
		}
		err error
		s   = new(Service)
	)

	checkErr(paladin.Get("application.toml").UnmarshalTOML(&ac))

	s.UserFacade, err = user_client.NewUserFacade("local")
	checkErr(err)
	s.UpdatesFacade, err = updates_facade.NewUpdatesFacade("local")
	checkErr(err)
	s.ReportFacade, err = report_facade.NewReportFacade("local")
	checkErr(err)
	s.PrivateFacade, err = private_facade.NewPrivateFacade("local")
	checkErr(err)
	s.ChatFacade, err = chat_facade.NewChatFacade("local")
	checkErr(err)
	s.ChannelFacade, err = channel_facade.NewChannelFacade("local")
	checkErr(err)
	s.MessageFacade, err = message_facade.NewMessageFacade("local")
	checkErr(err)

	sync_client.New()
	media_client.New()
	return s
}

func (s *Service) Close() {
}

// ///////////////////////////////////////////////////////////////////////////////////////
func mergeLimitDialogs(limit int32, idList ...model.DialogExtList) model.DialogExtList {
	var dialogs model.DialogExtList
	for _, id := range idList {
		dialogs = append(dialogs, id...)
	}
	if len(dialogs) == 0 {
		return dialogs
	}

	c := sort.Reverse(dialogs)
	sort.Sort(c)

	if limit == 0 {
		return dialogs
	} else {
		if len(dialogs) <= int(limit) {
			return dialogs
		} else {
			return dialogs[:limit]
		}
	}
}

type dialogsData struct {
	Dialogs  []*mtproto.Dialog
	Messages []*mtproto.Message
	Chats    []*mtproto.Chat
	Users    []*mtproto.User
}

func (s *Service) ToMessageDialogs(ctx context.Context, selfUserId int32, dialogList model.DialogExtList) (*dialogsData, error) {
	dialogs := &dialogsData{
		Dialogs:  []*mtproto.Dialog{},
		Messages: []*mtproto.Message{},
		Chats:    []*mtproto.Chat{},
		Users:    []*mtproto.User{},
	}
	if len(dialogList) == 0 {
		return dialogs, nil
	}

	settings, err := s.UserFacade.GetAllNotifySettings(ctx, selfUserId)
	if err != nil {
		log.Errorf("getNotifySettingsByPeerList - %v", err)
	}

	var (
		msgIdList            model.IDList
		channelMsgDataIdList model.ID64List
		dialogsMessages      model.MessageBoxList
	)
	var chatIds model.IDList
	var channelIds model.IDList
	for _, d := range dialogList {
		peer := model.FromPeer(d.Peer)
		if s, ok := settings[int64(peer.PeerType)<<32|int64(peer.PeerId)]; ok {
			d.NotifySettings = s
		}
		switch peer.PeerType {
		case model.PEER_USER:
			msgIdList.AddIfNot(d.TopMessage)
		case model.PEER_CHAT:
			msgIdList.AddIfNot(d.TopMessage)
			chatIds.AddIfNot(peer.PeerId)
		case model.PEER_CHANNEL:
			channelIds.AddIfNot(peer.PeerId)
			if d.AvailableMinId == d.TopMessage {
				dialogsMessages = append(dialogsMessages, &model.MessageBox{
					SelfUserId:     selfUserId,
					MessageId:      d.TopMessage,
					MessageBoxType: model.MESSAGE_BOX_TYPE_CHANNEL,
					Message: mtproto.MakeTLMessageService(&mtproto.Message{
						Out:             true,
						Id:              d.TopMessage,
						FromId_FLAGPEER: model.MakePeerUser(selfUserId),
						PeerId:          d.Peer,
						ToId:            d.Peer,
						Date:            d.Date,
						Action:          model.MakeMessageActionHistoryClear(),
					}).To_Message(),
				})
			} else {
				channelMsgDataIdList.AddIfNot(int64(d.Peer.ChannelId)<<32 | int64(d.TopMessage))
			}
		default:
			log.Errorf("fatal error")
		}
	}

	// 1. dialogs
	dialogs.Dialogs = dialogList.ToDialogList()

	// 2. messages
	dialogsMessages = append(dialogsMessages, s.MessageFacade.GetUserMessageList(ctx, selfUserId, msgIdList)...)
	dialogsMessages = append(dialogsMessages, s.MessageFacade.GetChannelMessageListByDataIdList(ctx, selfUserId, channelMsgDataIdList)...)
	dialogs.Messages, dialogs.Users, dialogs.Chats = dialogsMessages.ToMessagesPeersList(ctx, selfUserId, s.UserFacade, s.ChatFacade, s.ChannelFacade)
	//for _, m := range dialogs.Messages {
	//	if m.GetPeerId() != nil && m.GetPeerId().PredicateName == mtproto.Predicate_peerUser {
	//		if m.GetFromId_FLAGPEER() != nil && m.GetFromId_FLAGPEER().PredicateName == mtproto.Predicate_peerUser {
	//			if m.GetPeerId().UserId == selfUserId {
	//				m.GetPeerId().UserId = m.GetFromId_FLAGPEER().UserId
	//			}
	//		}
	//	}
	//}
	/*
		for _, chat := range dialogs.Chats {
			switch chat.PredicateName {
			case mtproto.Predicate_chat:
				for i := 0; i < len(chatIds); i++ {
					if chatIds[i] == chat.GetId() {
						chatIds = append(chatIds[:i], chatIds[i+1:]...)
						break
					}
				}
			case mtproto.Predicate_channel:
				for i := 0; i < len(channelIds); i++ {
					if channelIds[i] == chat.GetId() {
						channelIds = append(channelIds[:i], channelIds[i+1:]...)
						break
					}
				}
			}
		}

		for _, id := range chatIds {
			if mutableChat, err := s.ChatFacade.GetMutableChat(ctx, id, selfUserId); err != nil {
				log.Errorf("getMutableChat - not found chat (%d), error: %v", id, err)
			} else {
				if chat := mutableChat.ToUnsafeChat(selfUserId); chat != nil {
					dialogs.Chats = append(dialogs.Chats, chat)
				}
			}
		}

		for _, id := range channelIds {
			if mutableChannel, err := s.ChannelFacade.GetMutableChannel(ctx, id, selfUserId); err != nil {
				log.Errorf("getMutableChannel - not found chat (%d), error: %v", id, err)
			} else {
				if channel := mutableChannel.ToUnsafeChat(selfUserId); channel != nil {
					dialogs.Chats = append(dialogs.Chats, channel)
				}
			}
		}
	*/
	return dialogs, nil
}
