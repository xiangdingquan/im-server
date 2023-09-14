package service

import (
	"context"
	"math"
	sync_client "open.chat/app/messenger/sync/client"
	"time"

	"open.chat/app/json/helper"
	"open.chat/app/messenger/msg/msgpb"
	"open.chat/app/sysconfig"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesSendMessage(ctx context.Context, request *mtproto.TLMessagesSendMessage) (reply *mtproto.Updates, err error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.sendMessage#fa88427a - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		hasBot = md.IsBot
		peer   = model.FromInputPeer2(md.UserId, request.Peer)
	)

	switch peer.PeerType {
	case model.PEER_SELF:
		peer.PeerType = model.PEER_USER
	case model.PEER_USER:
		if !md.IsBot {
			hasBot = s.UserFacade.IsBot(ctx, peer.PeerId)
		}
	case model.PEER_CHAT:
	case model.PEER_CHANNEL:
	default:
		log.Errorf("invalid peer: %v", request.Peer)
		err = mtproto.ErrPeerIdInvalid
		return
	}

	if request.Message == "" {
		err = mtproto.ErrMessageEmpty
		log.Errorf("message empty: %v", err)
		return
	}

	if len(request.Message) > 4000 {
		err = mtproto.ErrMessageTooLong
		log.Errorf("messages.sendMessage: %v", err)
		return
	}

	if peer.PeerType == model.PEER_USER {
		err = s.checkUserSendRight(ctx, peer.PeerId, md.UserId, md.GetAuthId())
		if err != nil {
			log.Errorf("messages.sendMessage, checkUserSendRight, error: %v", err)
			return
		}
	}

	if peer.PeerType == model.PEER_CHAT {
		err = s.checkChatSendRight(ctx, peer.PeerId, md.UserId, nil)
		if err != nil {
			return
		}
	}

	if peer.PeerType == model.PEER_CHANNEL {
		err = s.checkChannelSendRight(ctx, peer.PeerId, md.UserId, nil)
		if err != nil {
			return
		}
	}

	if peer.PeerType == model.PEER_CHAT || peer.PeerType == model.PEER_CHANNEL {
		interval := sysconfig.GetConfig2Uint32(ctx, sysconfig.ConfigKeysGroupChatSendInterval, 0, 0)
		if interval > 0 && sysconfig.KeyIsExist(ctx, md.UserId, "limit_send_interval", interval) {
			err = mtproto.ErrChatRestricted
			log.Errorf("send message too quick: %v", err)
			return
		}
	}

	if request.Message != "" {
		msg := helper.TJsonMessage{
			Msg: request.Message,
		}
		var sMsg string
		//过滤红包消息
		if msg.Parse(2, 0, &sMsg) {
			request.Message = "1"
		}
	}

	// 由客户端实现就可以了
	//_, err = s.kickWhoSendBanWord(ctx, peer, md.UserId, md.GetAuthId(), request.Message)
	//if err != nil {
	//	log.Errorf("kickWhoSendBanWord, error: %v", err)
	//	return
	//}

	outMessage := mtproto.MakeTLMessage(&mtproto.Message{
		Out:               true,
		Mentioned:         false,
		MediaUnread:       false,
		Silent:            request.Silent,
		Post:              false,
		FromScheduled:     false,
		Legacy:            false,
		EditHide:          false,
		Id:                0,
		FromId_FLAGPEER:   model.MakePeerUser(md.UserId),
		ToId:              peer.ToPeer(),
		FwdFrom:           nil,
		ViaBotId:          nil,
		ReplyTo:           nil,
		Date:              int32(time.Now().Unix()),
		Message:           request.Message,
		Media:             nil,
		ReplyMarkup:       request.ReplyMarkup,
		Entities:          request.Entities,
		Views:             nil,
		Forwards:          nil,
		EditDate:          nil,
		PostAuthor:        nil,
		GroupedId:         nil,
		RestrictionReason: nil,
		TtlSeconds:        request.GetTtlSeconds(),
	}).To_Message()

	if request.ReplyToMsgId != nil {
		outMessage.ReplyTo = mtproto.MakeTLMessageReplyHeader(
			&mtproto.MessageReplyHeader{
				ReplyToMsgId: request.ReplyToMsgId.GetValue(),
			}).To_MessageReplyHeader()
	}

	outMessage, _ = s.fixMessageEntities(ctx, md.UserId, peer, request.NoWebpage, outMessage, hasBot)
	outboxMsg := &msgpb.OutboxMessage{
		NoWebpage:    request.NoWebpage,
		Background:   request.Background,
		RandomId:     request.RandomId,
		Message:      outMessage,
		ScheduleDate: request.ScheduleDate,
	}

	reply, err = s.MsgFacade.SendMessage(ctx, md.UserId, md.AuthId, peer, outboxMsg)

	if err != nil {
		log.Errorf("messages.sendMessage#fa88427a - error: %v", err)
	} else {
		go func() {
			if request.ClearDraft {
				s.doClearDraft(context.Background(), md.UserId, md.AuthId, peer)
			}
			s.UserFacade.UpdateUserStatus(context.Background(), md.UserId, time.Now().Unix())
		}()

		log.Debugf("messages.sendMessage#fa88427a - reply: %s", reply.DebugString())
	}
	return
}

func (s *Service) kickWhoSendBanWord(ctx context.Context, peer *model.PeerUtil, whoSend int32, auth int64, message string) (kicked bool, err error) {
	log.Debugf("sendMessage - kickWhoSendBanWord, whoSend:%d", whoSend)
	isContain, err := s.isContainBanWord(ctx, peer, message)
	if err != nil {
		return false, err
	}
	if !isContain {
		return false, err
	}
	log.Debugf("sendMessage - kickWhoSendBanWord, is contain ban word")

	var creator int32
	var chatId int32
	if peer.PeerType == model.PEER_CHAT {
		chat, err := s.ChatFacade.GetMutableChat(ctx, peer.PeerId)
		if err != nil {
			return false, err
		}
		creator = chat.Chat.Creator
		chatId = chat.Chat.Id
	} else if peer.PeerType == model.PEER_CHANNEL {
		channel, err := s.ChannelFacade.GetMutableChannel(ctx, peer.PeerId)
		if err != nil {
			return false, err
		}
		creator = channel.Channel.CreatorId
		chatId = channel.Channel.Id
	}
	log.Debug("sendMessage - kickWhoSendBanWord, creator:%d, chatId:%d", creator, chatId)

	if creator == 0 {
		return false, nil
	}

	if creator == whoSend {
		return false, nil
	}

	bannedRights := model.ChatBannedRights{
		Rights:    math.MaxInt32,
		UntilDate: math.MaxInt32,
	}
	channel, deleted, err := s.ChannelFacade.EditBanned(ctx, chatId, creator, whoSend, bannedRights)
	log.Debug("sendMessage - kickWhoSendBanWord, EditBanned")
	if err != nil {
		return false, err
	}
	log.Debugf("sendMessage - kickWhoSendBanWord, deleted:%b", deleted)
	helper.MakeSender(uint32(whoSend), auth, 6, 100)
	go func() {
		sync_client.PushUpdates(context.Background(), whoSend, mtproto.MakeTLUpdates(&mtproto.Updates{
			Updates: []*mtproto.Update{model.MakeUpdateChannel(channel.GetChannelId())},
			Users:   []*mtproto.User{},
			Chats:   []*mtproto.Chat{channel.ToUnsafeChat(whoSend)},
			Date:    int32(time.Now().Unix()),
			Seq:     0,
		}).To_Updates())
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

	return true, nil
}
