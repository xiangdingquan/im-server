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

func (s *Service) ChannelsInviteToChannel(ctx context.Context, request *mtproto.TLChannelsInviteToChannel) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.inviteToChannel#199f3a6c - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("channels.inviteToChannel - error: %v", err)
		return nil, err
	}

	if !IsInputChannel(request.Channel) {
		err := mtproto.ErrChannelInvalid
		log.Errorf("channels.inviteToChannel - error: %v", err)
		return nil, err
	}

	idList := make([]int32, 0, len(request.Users))
	for _, u := range request.Users {
		peer := model.FromInputUser(md.UserId, u)
		switch peer.PeerType {
		case model.PEER_USER:
			idList = append(idList, u.GetUserId())
		default:
			log.Warnf("invalid user - %v", u)
		}
	}

	if len(idList) == 0 {
		err := mtproto.ErrUserIdInvalid
		log.Errorf("channels.inviteToChannel#199f3a6c - error: %v", err)
		return nil, err
	} else if len(idList) == 1 {
		users := s.UserFacade.GetMutableUsers(ctx, md.UserId, idList[0])
		_ = users
	} else {
		users := s.UserFacade.GetMutableUsers(ctx, append(idList, md.UserId)...)
		_ = users
	}

	channel, invitedList, err := s.ChannelFacade.InviteToChannel(ctx,
		request.Channel.ChannelId,
		md.UserId,
		idList...)
	if err != nil {
		log.Errorf("inviteToChannel - error: %v", err)
		return nil, err
	}

	var meUpdates *mtproto.Updates

	if len(invitedList) == 0 {
		meUpdates = mtproto.MakeTLUpdates(&mtproto.Updates{
			Updates: []*mtproto.Update{},
			Users:   []*mtproto.User{},
			Chats:   []*mtproto.Chat{channel.ToUnsafeChat(md.UserId)},
			Date:    int32(time.Now().Unix()),
			Seq:     0,
		}).To_Updates()
	} else if channel.Channel.Broadcast {
		meUpdates = mtproto.MakeTLUpdates(&mtproto.Updates{
			Updates: []*mtproto.Update{},
			Users:   []*mtproto.User{},
			Chats:   []*mtproto.Chat{channel.ToUnsafeChat(md.UserId)},
			Date:    int32(time.Now().Unix()),
			Seq:     0,
		}).To_Updates()

		go func() {
			sync_client.SyncUpdatesNotMe(ctx, md.UserId, md.AuthId, mtproto.MakeTLUpdates(&mtproto.Updates{
				Updates: []*mtproto.Update{model.MakeUpdateChannel(channel.GetChannelId())},
				Users:   []*mtproto.User{},
				Chats:   []*mtproto.Chat{channel.ToUnsafeChat(md.UserId)},
				Date:    int32(time.Now().Unix()),
				Seq:     0,
			}).To_Updates())
			for _, id := range idList {
				sync_client.PushUpdates(ctx, id, mtproto.MakeTLUpdates(&mtproto.Updates{
					Updates: []*mtproto.Update{model.MakeUpdateChannel(channel.GetChannelId())},
					Users:   []*mtproto.User{},
					Chats:   []*mtproto.Chat{channel.ToUnsafeChat(id)},
					Date:    int32(time.Now().Unix()),
					Seq:     0,
				}).To_Updates())
			}
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
				Message:      channel.MakeMessageService(md.UserId, false, 0, model.MakeMessageActionChatAddUser(idList...)),
				ScheduleDate: nil,
			})
		if err != nil {
			log.Errorf("channels.inviteToChannel - push editMessage error: %v", err)
			return nil, err
		}
	}

	log.Debugf("channels.inviteToChannel - reply: %s", meUpdates.DebugString())
	return model.WrapperGoFunc(meUpdates, func() {
		for _, id := range invitedList {
			s.AdminLogClient.PutChannelAdminLogEventAction(context.Background(),
				&adminlogpb.ChannelAdminLogEventData{
					LogUserId: md.UserId,
					ChannelId: channel.GetChannelId(),
					Event: mtproto.MakeTLChannelAdminLogEventActionParticipantInvite(&mtproto.ChannelAdminLogEventAction{
						Participant: mtproto.MakeTLChannelParticipant(&mtproto.ChannelParticipant{
							UserId: id,
							Date:   int32(time.Now().Unix()),
						}).To_ChannelParticipant(),
					}).To_ChannelAdminLogEventAction(),
				})

		}
	}).(*mtproto.Updates), nil
}
