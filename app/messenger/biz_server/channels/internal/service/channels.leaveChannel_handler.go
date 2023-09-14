package service

import (
	"context"

	"math/rand"
	"time"

	"open.chat/app/job/admin_log/adminlogpb"
	"open.chat/app/messenger/msg/msgpb"
	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsLeaveChannel(ctx context.Context, request *mtproto.TLChannelsLeaveChannel) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.leaveChannel#f836aa95 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if !IsInputChannel(request.Channel) {
		err := mtproto.ErrChannelInvalid
		log.Errorf("channels.leaveChannel - error: %v", err)
		return nil, err
	}

	channel, err := s.ChannelFacade.LeaveChannel(ctx, request.Channel.ChannelId, md.UserId)
	if err != nil {
		log.Errorf("channels.leaveChannel - error: %v", err)
		return nil, err
	}

	var meUpdates *mtproto.Updates

	if channel.Channel.Broadcast {
		meUpdates = mtproto.MakeTLUpdates(&mtproto.Updates{
			Updates: []*mtproto.Update{},
			Users:   []*mtproto.User{},
			Chats:   []*mtproto.Chat{channel.ToUnsafeChat(md.UserId)},
			Date:    int32(time.Now().Unix()),
			Seq:     0,
		}).To_Updates()

		go func() {
			pushUpdates := mtproto.MakeTLUpdates(&mtproto.Updates{
				Updates: []*mtproto.Update{model.MakeUpdateChannel(channel.GetChannelId())},
				Users:   []*mtproto.User{},
				Chats:   []*mtproto.Chat{channel.ToUnsafeChat(md.UserId)},
				Date:    int32(time.Now().Unix()),
				Seq:     0,
			}).To_Updates()
			sync_client.SyncUpdatesNotMe(ctx, md.UserId, md.AuthId, pushUpdates)
		}()
	} else {
		meUpdates, err = s.MsgFacade.SendMessage(ctx,
			md.UserId,
			md.AuthId,
			model.MakeChannelPeerUtil(channel.GetChannelId()),
			&msgpb.OutboxMessage{
				NoWebpage:    true,
				Background:   false,
				RandomId:     rand.Int63(),
				Message:      channel.MakeMessageService(md.UserId, false, 0, model.MakeMessageActionChatDeleteUser(md.UserId)),
				ScheduleDate: nil,
			})
		if err != nil {
			log.Errorf("channels.leaveChannel - send leaveChannel error: %v", err)
			return nil, err
		}
	}

	log.Debugf("channels.leaveChannel - reply: {%s}", meUpdates.DebugString())
	return model.WrapperGoFunc(meUpdates, func() {
		s.AdminLogClient.PutChannelAdminLogEventAction(context.Background(), &adminlogpb.ChannelAdminLogEventData{
			LogUserId: md.UserId,
			ChannelId: channel.GetChannelId(),
			Event:     mtproto.MakeTLChannelAdminLogEventActionParticipantLeave(&mtproto.ChannelAdminLogEventAction{}).To_ChannelAdminLogEventAction(),
		})
	}).(*mtproto.Updates), nil
}
