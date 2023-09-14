package service

import (
	"context"

	"math/rand"

	"open.chat/app/job/admin_log/adminlogpb"
	"open.chat/app/messenger/msg/msgpb"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsEditTitle(ctx context.Context, request *mtproto.TLChannelsEditTitle) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.editTitle - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if !IsInputChannel(request.Channel) {
		err := mtproto.ErrChannelInvalid
		log.Errorf("channels.editTitle - error: %v", err)
		return nil, err
	}

	if request.Title == "" {
		err := mtproto.ErrChatTitleEmpty
		log.Errorf("channels.editTitle - error: %v", err)
		return nil, err
	}

	channel, err := s.ChannelFacade.EditTitle(ctx, request.Channel.ChannelId, md.UserId, request.Title)
	if err != nil {
		log.Errorf("channels.editTitle - error: %v", err)
		return nil, err
	}

	result, err := s.MsgFacade.SendMessage(
		ctx,
		md.UserId,
		md.AuthId,
		model.MakeChannelPeerUtil(request.Channel.ChannelId),
		&msgpb.OutboxMessage{
			NoWebpage:  true,
			Background: false,
			RandomId:   rand.Int63(),
			Message: channel.MakeMessageService(md.UserId,
				false,
				0,
				model.MakeMessageActionChatEditTitle(request.Title)),
			ScheduleDate: nil,
		})

	if err != nil {
		log.Errorf("channels.editTitle - error: %v", err)
		return nil, err
	}

	log.Debugf("channels.editTitle - reply: {%s}", result.DebugString())
	return model.WrapperGoFunc(result, func() {
		s.AdminLogClient.PutChannelAdminLogEventAction(context.Background(), &adminlogpb.ChannelAdminLogEventData{
			LogUserId: md.UserId,
			ChannelId: channel.GetChannelId(),
			Event: mtproto.MakeTLChannelAdminLogEventActionChangeTitle(&mtproto.ChannelAdminLogEventAction{
				PrevValue_STRING: request.Title,
				NewValue_STRING:  request.Title,
			}).To_ChannelAdminLogEventAction(),
		})
	}).(*mtproto.Updates), nil
}
