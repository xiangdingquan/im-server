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

func (s *Service) MessagesEditChatDefaultBannedRights(ctx context.Context, request *mtproto.TLMessagesEditChatDefaultBannedRights) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.editChatDefaultBannedRights#a5866b41 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		peer  = model.FromInputPeer2(md.UserId, request.Peer)
		reply *mtproto.Updates
		date  = int32(time.Now().Unix())
	)

	switch peer.PeerType {
	case model.PEER_CHAT:
		chat, err := s.ChatFacade.EditChatDefaultBannedRights(ctx, peer.PeerId, md.UserId, request.BannedRights)
		if err != nil {
			log.Errorf("messages.editChatDefaultBannedRights - error: %v", err)
			return nil, err
		}

		// push updates
		go func() {
			defaultBannedUpdates := mtproto.MakeTLUpdateShort(&mtproto.Updates{
				Update: mtproto.MakeTLUpdateChatDefaultBannedRights(&mtproto.Update{
					Peer_PEER:           model.MakePeerChat(peer.PeerId),
					DefaultBannedRights: chat.Chat.DefaultBannedRights,
					Version:             chat.Chat.Version,
				}).To_Update(),
				Date: date,
			}).To_Updates()

			sync_client.SyncUpdatesNotMe(context.Background(), md.UserId, md.AuthId, defaultBannedUpdates)
			sync_client.BroadcastChatUpdates(context.Background(), peer.PeerId, defaultBannedUpdates, md.UserId)
		}()

		reply = mtproto.MakeTLUpdates(&mtproto.Updates{
			Updates: []*mtproto.Update{},
			Users:   []*mtproto.User{},
			Chats:   []*mtproto.Chat{chat.ToUnsafeChat(md.UserId)},
			Date:    int32(time.Now().Unix()),
			Seq:     0,
		}).To_Updates()
	case model.PEER_CHANNEL:
		channel, err := s.ChannelFacade.EditChatDefaultBannedRights(ctx, peer.PeerId, md.UserId, request.BannedRights)
		if err != nil {
			log.Errorf("messages.editChatDefaultBannedRights - error: %v", err)
			return nil, err
		}

		// push updates
		go func() {
			defaultBannedUpdates := mtproto.MakeTLUpdateShort(&mtproto.Updates{
				Update: mtproto.MakeTLUpdateChatDefaultBannedRights(&mtproto.Update{
					Peer_PEER:           model.MakePeerChannel(peer.PeerId),
					DefaultBannedRights: channel.Channel.DefaultBannedRights.ToChatBannedRights(),
					Version:             0,
				}).To_Update(),
				Date: date,
			}).To_Updates()

			sync_client.SyncUpdatesNotMe(context.Background(), md.UserId, md.AuthId, defaultBannedUpdates)
			sync_client.BroadcastChannelUpdates(context.Background(), peer.PeerId, defaultBannedUpdates, md.UserId)
		}()

		reply = mtproto.MakeTLUpdates(&mtproto.Updates{
			Updates: []*mtproto.Update{},
			Users:   []*mtproto.User{},
			Chats:   []*mtproto.Chat{channel.ToUnsafeChat(md.UserId)},
			Date:    int32(time.Now().Unix()),
			Seq:     0,
		}).To_Updates()
	default:
		err := mtproto.ErrPeerIdInvalid
		log.Errorf("invalid peer type: {%v}")
		return nil, err
	}

	log.Debugf("messages.editChatDefaultBannedRights - reply: {%s}", reply.DebugString())
	return reply, nil
}
