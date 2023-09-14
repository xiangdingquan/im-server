package service

import (
	"context"
	"time"

	"github.com/gogo/protobuf/types"

	idgen "open.chat/app/service/idgen/client"

	"open.chat/app/messenger/msg/msgpb"
	"open.chat/app/sysconfig"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesSendMultiMedia(ctx context.Context, request *mtproto.TLMessagesSendMultiMedia) (reply *mtproto.Updates, err error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.sendMultiMedia#2095512f - metadata: %s, request: %s", md.DebugString(), request.DebugString())

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

	if peer.PeerType == model.PEER_USER {
		err = s.checkUserSendRight(ctx, peer.PeerId, md.UserId, md.GetAuthId())
		if err != nil {
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

	groupedId := &types.Int64Value{Value: idgen.GetUUID()}
	outboxMultiMedia := make([]*msgpb.OutboxMessage, 0, len(request.MultiMedia))
	for _, media := range request.MultiMedia {
		if len(media.Message) > 4000 {
			err = mtproto.ErrMessageTooLong
			log.Errorf("messages.sendMultiMedia: %v", err)
			return
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
			Message:           media.Message,
			ReplyMarkup:       nil, // request.ReplyMarkup,
			Entities:          media.Entities,
			Views:             nil,
			Forwards:          nil,
			EditDate:          nil,
			PostAuthor:        nil,
			GroupedId:         groupedId,
			RestrictionReason: nil,
			TtlSeconds:        request.GetTtlSeconds(),
		}).To_Message()

		if request.ReplyToMsgId != nil {
			outMessage.ReplyTo = mtproto.MakeTLMessageReplyHeader(
				&mtproto.MessageReplyHeader{
					ReplyToMsgId: request.ReplyToMsgId.GetValue(),
				}).To_MessageReplyHeader()
		}

		outMessage.Media, err = s.makeMediaByInputMedia(ctx, md.UserId, md.AuthId, peer, media.GetMedia())
		if err != nil {
			if err.Error() == string(sysconfig.ConfigKeysBanImages) {
				media.Message = media.Message + s.getIllegalMessage(ctx)
				return
			} else {
				log.Errorf("messages.sendMultiMedia: %v", err)
				return
			}
		}
		outMessage, _ = s.fixMessageEntities(ctx, md.UserId, peer, true, outMessage, hasBot)
		outboxMultiMedia = append(outboxMultiMedia, &msgpb.OutboxMessage{
			NoWebpage:    true,
			Background:   request.Background,
			RandomId:     media.RandomId,
			Message:      outMessage,
			ScheduleDate: request.ScheduleDate,
		})
	}

	reply, err = s.MsgFacade.SendMultiMessage(ctx, md.UserId, md.AuthId, peer, outboxMultiMedia)

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
