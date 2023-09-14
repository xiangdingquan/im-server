package service

import (
	"context"
	"time"

	"github.com/gogo/protobuf/types"

	"open.chat/app/messenger/msg/msgpb"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) checkForwardPrivacy(ctx context.Context, selfUserId, checkId int32) bool {
	rules, _ := s.UserFacade.GetPrivacy(ctx, selfUserId, model.FORWARDS)
	if len(rules) == 0 {
		return true
	}
	return model.CheckPrivacyIsAllow(selfUserId,
		rules,
		checkId,
		func(id, checkId int32) bool {
			return s.UserFacade.CheckContact(ctx, id, checkId)
		},
		func(checkId int32, idList []int32) bool {
			chatIdList, channelIdList := model.SplitChatAndChannelIdList(idList)
			return s.ChatFacade.CheckParticipantIsExist(ctx, checkId, chatIdList) ||
				s.ChannelFacade.CheckParticipantIsExist(ctx, checkId, channelIdList)
		})
}

func (s *Service) makeForwardMessages(
	ctx context.Context,
	md *grpc_util.RpcMetadata,
	fromPeer, toPeer *model.PeerUtil,
	saved bool,
	request *mtproto.TLMessagesForwardMessages) ([]*msgpb.OutboxMessage, error) {

	var (
		idList  = request.Id
		ridList = request.RandomId
		now     = int32(time.Now().Unix())
	)

	findRandomIdById := func(id int32) int64 {
		for i := 0; i < len(idList); i++ {
			if id == idList[i] {
				return ridList[i]
			}
		}
		return 0
	}

	var messageList model.MessageBoxList
	if fromPeer.PeerType == model.PEER_CHANNEL {
		messageList = s.MessageFacade.GetChannelMessageList(ctx, md.UserId, fromPeer.PeerId, idList)
	} else {
		messageList = s.MessageFacade.GetUserMessageList(ctx, md.UserId, idList)
	}

	fwdOutboxList := make([]*msgpb.OutboxMessage, 0, len(messageList))
	for _, box := range messageList {
		m := box.Message
		if m.FwdFrom == nil {
			fwdFrom := mtproto.MakeTLMessageFwdHeader(&mtproto.MessageFwdHeader{
				Date: m.GetDate(),
			}).To_MessageFwdHeader()

			if m.Views != nil {
				fwdFrom.ChannelId = &types.Int32Value{Value: fromPeer.PeerId}
				fwdFrom.ChannelPost = &types.Int32Value{Value: m.Id}
				fwdFrom.PostAuthor = m.PostAuthor
			} else {
				fromId := box.SendUserId
				if s.checkForwardPrivacy(ctx, fromId, md.UserId) {
					fwdFrom.FromId_FLAGPEER = model.MakePeerUser(fromId)
				} else {
					fwdFrom.FromName = &types.StringValue{Value: s.UserFacade.GetUserName(ctx, fromId)}
				}
				m.Post = false
				m.PostAuthor = nil
			}

			if saved {
				if m.Views != nil {
					fwdFrom.SavedFromPeer = box.Message.ToId
				} else {
					fwdFrom.SavedFromPeer = model.MakePeerUser(box.SendUserId)
				}
				fwdFrom.SavedFromMsgId = &types.Int32Value{Value: m.Id}
			}
			m.FwdFrom = fwdFrom
		} else {
			if saved {
				if m.Views != nil {
					m.FwdFrom.SavedFromPeer = box.Message.ToId
				} else {
					m.FwdFrom.SavedFromPeer = model.MakePeerUser(box.SendUserId)
				}
				m.FwdFrom.SavedFromMsgId = &types.Int32Value{Value: m.Id}
			} else {
				m.FwdFrom.SavedFromPeer = nil
				m.FwdFrom.SavedFromMsgId = nil
			}
		}
		m.ToId = toPeer.ToPeer()
		m.FromId_FLAGPEER = model.MakePeerUser(md.UserId)
		m.Date = now
		m.Silent = request.Silent
		m.Post = false

		fwdOutboxList = append(fwdOutboxList, &msgpb.OutboxMessage{
			NoWebpage:    true,
			Background:   false,
			RandomId:     findRandomIdById(m.GetId()),
			Message:      m,
			ScheduleDate: request.GetScheduleDate(),
		})
	}

	return fwdOutboxList, nil
}

func (s *Service) MessagesForwardMessages(ctx context.Context, request *mtproto.TLMessagesForwardMessages) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.forwardMessages#708e0195 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		fromPeer      = model.FromInputPeer2(md.UserId, request.FromPeer)
		toPeer        = model.FromInputPeer2(md.UserId, request.ToPeer)
		err           error
		resultUpdates *mtproto.Updates
		saved         = false
	)

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.forwardMessages#708e0195 - error: %v", err)
		return nil, err
	}

	switch toPeer.PeerType {
	case model.PEER_SELF:
		toPeer.PeerType = model.PEER_USER
		saved = true
	case model.PEER_USER:
		if toPeer.PeerId == md.UserId {
			saved = true
		}
	case model.PEER_CHAT:
	case model.PEER_CHANNEL:
	default:
		err = mtproto.ErrPeerIdInvalid
		log.Errorf("messages.forwardMessages#708e0195 - error: %v", err)
		return nil, err
	}

	if len(request.Id) == 0 ||
		len(request.RandomId) == 0 ||
		len(request.Id) != len(request.RandomId) {

		err = mtproto.ErrInputRequestInvalid
		log.Errorf("invalid id or random_id")
		return nil, err
	}

	fwdOutboxList, err := s.makeForwardMessages(ctx, md, fromPeer, toPeer, saved, request)
	if err != nil {
		log.Errorf("messages.forwardMessages#708e0195 - error: %v", err)
		return nil, err
	}

	resultUpdates, err = s.MsgFacade.SendMultiMessage(ctx, md.UserId, md.AuthId, toPeer, fwdOutboxList)
	if err != nil {
		log.Errorf("messages.forwardMessages#708e0195 - error: %v", err)
	} else {
		log.Debugf("messages.forwardMessages#708e0195 - reply: %s", resultUpdates.DebugString())
	}
	return resultUpdates, err
}
