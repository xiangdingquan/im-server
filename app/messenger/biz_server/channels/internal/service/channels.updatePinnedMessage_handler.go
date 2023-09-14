package service

import (
	"context"
	"math/rand"

	"open.chat/app/messenger/msg/msgpb"
	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsUpdatePinnedMessage(ctx context.Context, request *mtproto.TLChannelsUpdatePinnedMessage) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.updatePinnedMessage - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if !IsInputChannel(request.Channel) {
		err := mtproto.ErrChannelInvalid
		log.Errorf("channels.updatePinnedMessage - error: %v", err)
		return nil, err
	}

	var (
		err       error
		channelId = request.Channel.ChannelId
		meUpdates *mtproto.Updates
	)

	channel, err := s.ChannelFacade.UpdatePinnedMessage(ctx, channelId, md.UserId, request.Id)
	if err != nil {
		log.Errorf("channels.updatePinnedMessage - error: %v", err)
		return nil, err
	}

	if request.Id > 0 {
		meUpdates, err = s.MsgFacade.SendMessage(ctx,
			md.UserId,
			md.AuthId,
			model.MakeChannelPeerUtil(channelId),
			&msgpb.OutboxMessage{
				NoWebpage:    true,
				Background:   false,
				RandomId:     rand.Int63(),
				Message:      channel.MakeMessageService(md.UserId, request.Silent, request.Id, model.MakeMessageActionPinMessage()),
				ScheduleDate: nil,
			})
		if err != nil {
			log.Errorf("channels.updatePinnedMessage - error: %v", err)
			return nil, err
		}
	} else {
		meUpdates = model.MakeUpdatesByUpdates(
			mtproto.MakeTLUpdateChannelPinnedMessage(&mtproto.Update{
				ChannelId: channelId,
				Id_INT32:  request.Id,
			}).To_Update(),
			mtproto.MakeTLUpdatePinnedChannelMessages(&mtproto.Update{
				Pinned:    true,
				ChannelId: channelId,
				Messages:  []int32{request.Id},
			}).To_Update(),
		)

		go func() {
			pushUpdates := meUpdates

			sync_client.SyncUpdatesNotMe(context.Background(), md.UserId, md.AuthId, pushUpdates)
			sync_client.BroadcastChannelUpdates(context.Background(), channelId, pushUpdates, md.UserId)
		}()
	}

	log.Debugf("channels.updatePinnedMessage - reply: %s", meUpdates.DebugString())
	return meUpdates, nil
}
