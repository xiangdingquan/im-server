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

func (s *Service) ChannelsReadHistory(ctx context.Context, request *mtproto.TLChannelsReadHistory) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.readHistory - metadata: %s, request: %s", md.DebugString(), request.DebugString())
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("channels.readHistory - error: %v", err)
		return nil, err
	}

	if !IsInputChannel(request.Channel) {
		err := mtproto.ErrChannelInvalid
		log.Errorf("channels.readHistory - error: %v", err)
		return nil, err
	}

	var (
		channelId = request.Channel.ChannelId
		date      = int32(time.Now().Unix())
	)

	channel, err := s.ChannelFacade.GetMutableChannel(ctx, channelId, md.UserId)
	if err != nil {
		log.Errorf("channels.readHistory - error: %v", err)
		return nil, err
	}

	me := channel.GetImmutableChannelParticipant(md.UserId)
	if me == nil || !me.CanViewMessages(date) {
		log.Debugf("channels.readHistory - reply: false")
		return mtproto.BoolFalse, nil
	}

	if request.MaxId <= 0 {
		err := mtproto.ErrMsgIdInvalid
		log.Errorf("channels.readHistory - error: %v", err)
		return nil, err
	}

	readInboxMaxId := s.ChannelFacade.GetReadInboxMaxId(ctx, md.UserId, channelId)
	s.ChannelFacade.ReadOutboxHistory(ctx, channelId, md.UserId, request.MaxId)
	s.MessageFacade.ReadEphemeralMsgByBetween(ctx, md.UserId, model.MakeChannelPeerUtil(channelId), readInboxMaxId, request.MaxId)

	log.Debugf("channels.readHistory - reply: true")
	return model.WrapperGoFunc(mtproto.BoolTrue, func() {
		sync_client.SyncUpdatesNotMe(context.Background(),
			md.UserId,
			md.AuthId,
			model.MakeUpdatesByUpdates(mtproto.MakeTLUpdateReadChannelInbox(&mtproto.Update{
				ChannelId: channelId,
				MaxId:     request.GetMaxId(),
			}).To_Update()))

		sync_client.BroadcastChannelUpdates(context.Background(),
			channelId,
			model.MakeUpdatesByUpdates(mtproto.MakeTLUpdateReadChannelOutbox(&mtproto.Update{
				ChannelId: channelId,
				MaxId:     request.GetMaxId(),
			}).To_Update()),
			md.UserId)
	}).(*mtproto.Bool), nil
}
