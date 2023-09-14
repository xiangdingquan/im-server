package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsDeleteMessages(ctx context.Context, request *mtproto.TLChannelsDeleteMessages) (*mtproto.Messages_AffectedMessages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Infof("channels.deleteMessages - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if !IsInputChannel(request.Channel) {
		err := mtproto.ErrChannelInvalid
		log.Errorf("channels.deleteMessages - error: %v", err)
		return nil, err
	}

	channel, err := s.ChannelFacade.GetMutableChannel(ctx, request.Channel.ChannelId, md.UserId)
	if err != nil {
		log.Errorf("channels.deleteMessages - error: %v", err)
		return nil, err
	}

	me := channel.GetImmutableChannelParticipant(md.UserId)
	if me == nil || !me.IsStateOk() {
		err = mtproto.ErrUserNotParticipant
		log.Errorf("channels.deleteMessages - error: %v", err)
		return nil, err
	}

	deleteBoxList := s.MessageFacade.GetChannelMessageList(ctx, md.UserId, request.Channel.ChannelId, request.Id)
	if len(deleteBoxList) != len(request.Id) {
	}

	if len(deleteBoxList) == 0 {
		affectedMessages := mtproto.MakeTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
			Pts:      channel.Channel.Pts,
			PtsCount: 0,
		}).To_Messages_AffectedMessages()

		log.Debugf("channels.deleteMessages - reply: {%s}", affectedMessages.DebugString())
		return affectedMessages, nil
	}

	if !me.CanAdminDeleteMessages() {
		for _, box := range deleteBoxList {
			if box.SendUserId != md.UserId {
				err = mtproto.ErrMessageDeleteForbidden
				log.Errorf("channels.deleteMessages - error: %v", err)
				return nil, err
			} else if box.MessageType == model.MESSAGE_TYPE_MESSAGE_SERVICE {
				err = mtproto.ErrMessageDeleteForbidden
				log.Errorf("channels.deleteMessages - error: %v", err)
				return nil, err
			}
		}
	}

	affectedMessages, err := s.MsgFacade.DeleteMessages(ctx,
		md.UserId,
		md.AuthId,
		model.MakeChannelPeerUtil(request.Channel.ChannelId),
		true,
		request.Id)

	if err != nil {
		log.Errorf("channels.deleteMessages - error: %v", err)
		return nil, err
	}

	log.Debugf("channels.deleteMessages - reply: %s", affectedMessages.DebugString())
	return affectedMessages, nil
}
