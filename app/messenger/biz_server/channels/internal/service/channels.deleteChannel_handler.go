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

func (s *Service) ChannelsDeleteChannel(ctx context.Context, request *mtproto.TLChannelsDeleteChannel) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.deleteChannel#c0111fe3 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("channels.deleteChannel - error: %v", err)
		return nil, err
	}

	if !IsInputChannel(request.Channel) {
		err := mtproto.ErrChannelInvalid
		log.Errorf("channels.deleteChannel - error: %v", err)
		return nil, err
	}

	channel, err := s.ChannelFacade.DeleteChannel(ctx, request.Channel.ChannelId, md.UserId)
	if err != nil {
		log.Errorf("channels.deleteChannel - error: %v", err)
		return nil, err
	}

	s.UsernameFacade.DeleteUsernameByPeer(ctx, model.PEER_CHANNEL, channel.GetChannelId())
	channelForbidden := channel.ToChannelForbidden()
	result := mtproto.MakeTLUpdates(&mtproto.Updates{
		Updates: []*mtproto.Update{},
		Users:   []*mtproto.User{},
		Chats:   []*mtproto.Chat{channelForbidden},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates()

	log.Debugf("channels.deleteChannel#c0111fe3b - reply: %s", result.DebugString())
	return model.WrapperGoFunc(result, func() {
		pushUpdates := result

		sync_client.SyncUpdatesNotMe(context.Background(), md.UserId, md.AuthId, pushUpdates)

		sync_client.BroadcastChannelUpdates(context.Background(), channel.GetChannelId(), pushUpdates, md.UserId)
	}).(*mtproto.Updates), nil
}
