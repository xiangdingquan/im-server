package service

import (
	"context"
	"errors"
	"strconv"
	"time"

	"open.chat/pkg/hack"

	"open.chat/app/messenger/msg/msgpb"
	"open.chat/app/sysconfig"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesSendMedia(ctx context.Context, request *mtproto.TLMessagesSendMedia) (reply *mtproto.Updates, err error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.sendMedia#c8f16791 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// peer
	var (
		hasBot = md.IsBot
		peer   *model.PeerUtil
	)

	peer = model.FromInputPeer2(md.UserId, request.Peer)
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

	if len(request.Message) > 4000 {
		err = mtproto.ErrMessageTooLong
		log.Errorf("messages.sendMedia: %v", err)
		return
	}

	if peer.PeerType == model.PEER_USER {
		err = s.checkUserSendRight(ctx, peer.PeerId, md.UserId, md.GetAuthId())
		if err != nil {
			log.Errorf("messages.sendMedia#c8f16791 - %v", err)
			return
		}
	}

	if peer.PeerType == model.PEER_CHAT {
		err = s.checkChatSendRight(ctx, peer.PeerId, md.UserId, request.Media)
		if err != nil {
			log.Errorf("messages.sendMedia#c8f16791 - %v", err)
			return
		}
	}

	if peer.PeerType == model.PEER_CHANNEL {
		err = s.checkChannelSendRight(ctx, peer.PeerId, md.UserId, request.Media)
		if err != nil {
			log.Errorf("messages.sendMedia#c8f16791 - %v", err)
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
		Media:             nil,
		Message:           request.Message,
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

	outMessage.Media, err = s.makeMediaByInputMedia(ctx, md.UserId, md.AuthId, peer, request.Media)
	if err != nil {
		if err.Error() == string(sysconfig.ConfigKeysBanImages) {
			return nil, errors.New(s.getIllegalMessage(ctx))
			//outMessage.Message = outMessage.Message + "[违规发图片,被屏蔽]"
		} else {
			log.Errorf("messages.sendMedia - error: %v", err)
			return nil, err
		}
	}

	poll, _ := model.GetPollByMessage(outMessage.Media)
	if poll != nil {
		var correctAnswers []int
		for _, v := range request.Media.CorrectAnswers {
			if iV, err := strconv.ParseInt(hack.String(v), 10, 64); err == nil {
				correctAnswers = append(correctAnswers, int(iV))
			}
		}

		mediaPoll, err := s.PollFacade.CreateMediaPoll(ctx, md.UserId, correctAnswers, poll)
		if err != nil {
			log.Errorf("createMediaPoll error - %v", err)
			return nil, err
		}
		outMessage.Media = mediaPoll.ToMessageMedia()
	}

	outMessage, _ = s.fixMessageEntities(ctx, md.UserId, peer, true, outMessage, hasBot)
	outboxMsg := &msgpb.OutboxMessage{
		NoWebpage:    true,
		Background:   request.Background,
		RandomId:     request.RandomId,
		Message:      outMessage,
		ScheduleDate: request.ScheduleDate,
	}

	reply, err = s.MsgFacade.SendMessage(ctx, md.UserId, md.AuthId, peer, outboxMsg)

	if err != nil {
		log.Errorf("messages.sendMedia#c8f16791 - error: %v", err)
	} else {
		go func() {
			if request.ClearDraft {
				s.doClearDraft(context.Background(), md.UserId, md.AuthId, peer)
			}

			s.UserFacade.UpdateUserStatus(context.Background(), md.UserId, time.Now().Unix())
		}()

		log.Debugf("messages.sendMedia#c8f16791 - reply: %s", reply.DebugString())
	}
	return
}
