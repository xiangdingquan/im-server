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

func (s *Service) MessagesSetTyping(ctx context.Context, request *mtproto.TLMessagesSetTyping) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.setTyping#a3825e50 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		peer = model.FromInputPeer2(md.UserId, request.GetPeer())
		date = int32(time.Now().Unix())
	)

	switch peer.PeerType {
	case model.PEER_USER:
		updates := mtproto.MakeTLUpdateShort(&mtproto.Updates{
			Update: mtproto.MakeTLUpdateUserTyping(&mtproto.Update{
				UserId: md.UserId,
				Action: request.Action,
			}).To_Update(),
			Date: date,
		}).To_Updates()
		sync_client.PushUpdates(ctx, peer.PeerId, updates)
	case model.PEER_CHAT:
		updates := mtproto.MakeTLUpdateShort(&mtproto.Updates{
			Update: mtproto.MakeTLUpdateChatUserTyping(&mtproto.Update{
				ChatId: peer.PeerId,
				UserId: md.UserId,
				Action: request.Action,
			}).To_Update(),
			Date: date,
		}).To_Updates()
		sync_client.BroadcastChatUpdates(ctx, peer.PeerId, updates, md.UserId)
	case model.PEER_CHANNEL:
		updates := mtproto.MakeTLUpdateShort(&mtproto.Updates{
			Update: mtproto.MakeTLUpdateChannelUserTyping(&mtproto.Update{
				ChannelId: peer.PeerId,
				UserId:    md.UserId,
				Action:    request.Action,
			}).To_Update(),
			Date: date,
		}).To_Updates()
		sync_client.BroadcastChannelUpdates(ctx, peer.PeerId, updates, md.UserId)
	}

	log.Debugf("messages.setTyping#a3825e50 - reply: {true}")
	return mtproto.ToBool(true), nil
}
