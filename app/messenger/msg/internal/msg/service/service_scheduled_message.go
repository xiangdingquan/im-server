package service

import (
	"context"

	"open.chat/app/messenger/msg/msgpb"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

func (s *Service) sendScheduledMessage(ctx context.Context, r *msgpb.OutgoingMessage) (reply *mtproto.Updates, err error) {
	switch r.PeerType {
	case model.PEER_USER:
		var users model.MutableUsers
		users = s.UserFacade.GetMutableUsers(ctx, r.From.Id, r.PeerId)
		sender, _ := users.GetImmutableUser(r.From.Id)
		if sender == nil || sender.Deleted() {
			err = mtproto.ErrInputUserDeactivated
			log.Errorf("sendUserOutgoingMessage - error: %v", err)
			return
		}
		peerUser, _ := users.GetImmutableUser(r.PeerId)
		if peerUser == nil || peerUser.Deleted() {
			err = mtproto.ErrPeerIdInvalid
			log.Errorf("sendUserOutgoingMessage - error: %v", err)
			return
		}
		if s.UserFacade.IsBlockedByUser(ctx, r.PeerId, r.From.Id) {
			err = mtproto.ErrYouBlockedUser
			log.Errorf("sendUserOutgoingMessage - error: %v", err)
			return
		}
	case model.PEER_CHAT:
		var users model.MutableUsers
		users = s.UserFacade.GetMutableUsers(ctx, r.From.Id)
		sender, _ := users.GetImmutableUser(r.From.Id)
		if sender == nil || sender.Deleted() {
			err = mtproto.ErrInputUserDeactivated
			log.Errorf("sendUserOutgoingMessage - error: %v", err)
			return
		}
	case model.PEER_CHANNEL:
		var users model.MutableUsers
		users = s.UserFacade.GetMutableUsers(ctx, r.From.Id)
		sender, _ := users.GetImmutableUser(r.From.Id)
		if sender == nil || sender.Deleted() {
			err = mtproto.ErrInputUserDeactivated
			log.Errorf("sendUserOutgoingMessage - error: %v", err)
			return
		}
	default:
		err = mtproto.ErrPeerIdInvalid
	}

	hasDuplicateMessage, err := s.MsgCore.HasDuplicateMessage(ctx, r.From.Id, r.Message.RandomId)
	if err != nil {
		log.Errorf("checkDuplicateMessage error - %v", err)
		return nil, err
	} else if hasDuplicateMessage {
		upd, err := s.MsgCore.GetDuplicateMessage(ctx, r.From.Id, r.Message.RandomId)
		if err != nil {
			log.Errorf("checkDuplicateMessage error - %v", err)
			return nil, err
		} else if upd != nil {
			return upd, nil
		}
	}

	outBox, err := s.MsgCore.SendScheduledMessage(ctx,
		r.From.Id,
		&model.PeerUtil{PeerType: r.PeerType, PeerId: r.PeerId},
		r.Message.RandomId,
		r.Message.ScheduleDate.Value,
		r.Message.Message)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	updatesHelper := model.MakeUpdatesHelper(mtproto.MakeTLUpdateNewScheduledMessage(&mtproto.Update{
		RandomId:        outBox.RandomId,
		Message_MESSAGE: model.MessageUpdate(outBox.Message),
	}).To_Update())

	replyUpdates := updatesHelper.ToReplyUpdates(ctx, r.From.Id, s.UserFacade, s.ChatFacade, s.ChannelFacade)

	s.MsgCore.PutDuplicateMessage(ctx, r.From.Id, r.Message.RandomId, replyUpdates)
	return replyUpdates, nil
}
