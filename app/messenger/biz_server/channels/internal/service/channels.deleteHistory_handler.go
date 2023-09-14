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

func (s *Service) ChannelsDeleteHistory(ctx context.Context, request *mtproto.TLChannelsDeleteHistory) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.deleteHistory - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("channels.deleteHistory - error: %v", err)
		return nil, err
	}

	if !IsInputChannel(request.Channel) {
		err := mtproto.ErrChannelInvalid
		log.Errorf("channels.deleteHistory - error: %v", err)
		return nil, err
	}

	_, err := s.ChannelFacade.DeleteHistory(ctx, request.Channel.ChannelId, md.UserId, request.MaxId)
	if err != nil {
		log.Errorf("channels.deleteHistory - error: %v", err)
		return nil, err
	}

	log.Debugf("channels.deleteHistory - reply: {true}")
	return model.WrapperGoFunc(mtproto.BoolTrue, func() {
		sync_client.SyncUpdatesNotMe(context.Background(), md.UserId, md.AuthId, mtproto.MakeTLUpdates(&mtproto.Updates{
			Updates: []*mtproto.Update{mtproto.MakeTLUpdateChannelAvailableMessages(&mtproto.Update{
				ChannelId:      request.Channel.ChannelId,
				AvailableMinId: request.MaxId,
			}).To_Update()},
			Users: []*mtproto.User{},
			Chats: []*mtproto.Chat{},
			Date:  int32(time.Now().Unix()),
			Seq:   0,
		}).To_Updates())
	}).(*mtproto.Bool), nil
}
