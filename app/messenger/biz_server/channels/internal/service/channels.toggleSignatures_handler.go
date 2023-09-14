package service

import (
	"context"

	"time"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsToggleSignatures(ctx context.Context, request *mtproto.TLChannelsToggleSignatures) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Infof("channels.toggleSignatures#1f69b606 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("channels.toggleSignatures - error: %v", err)
		return nil, err
	}

	if !IsInputChannel(request.Channel) {
		err := mtproto.ErrChannelInvalid
		log.Errorf("channels.toggleSignatures - error: %v", err)
		return nil, err
	}

	channel, err := s.ChannelFacade.ToggleSignatures(ctx, request.Channel.ChannelId, md.UserId, mtproto.FromBool(request.Enabled))
	if err != nil {
		log.Errorf("channels.toggleSignatures - error: %v", err)
		return nil, err
	}

	result := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{channel.ToUnsafeChat(md.UserId)},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates()

	log.Debugf("channels.toggleSignatures - reply: {%s}", result.DebugString())
	return model.WrapperGoFunc(result, func() {
		sync_client.SyncUpdatesNotMe(context.Background(), md.UserId, md.AuthId, mtproto.MakeTLUpdates(&mtproto.Updates{
			Updates: []*mtproto.Update{model.MakeUpdateChannel(channel.GetChannelId())},
			Users:   []*mtproto.User{},
			Chats:   []*mtproto.Chat{channel.ToUnsafeChat(md.UserId)},
			Date:    int32(time.Now().Unix()),
			Seq:     0,
		}).To_Updates())

		sync_client.BroadcastChannelAdminsUpdates(context.Background(), request.Channel.ChannelId, mtproto.MakeTLUpdates(&mtproto.Updates{
			Updates: []*mtproto.Update{model.MakeUpdateChannel(channel.GetChannelId())},
			Users:   []*mtproto.User{},
			Chats:   []*mtproto.Chat{},
			Date:    int32(time.Now().Unix()),
			Seq:     0,
		}).To_Updates(), md.UserId)
	}).(*mtproto.Updates), nil
}
