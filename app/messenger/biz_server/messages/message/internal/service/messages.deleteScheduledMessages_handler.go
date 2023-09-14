package service

import (
	"context"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesDeleteScheduledMessages(ctx context.Context, request *mtproto.TLMessagesDeleteScheduledMessages) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.deleteScheduledMessages - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	peer := model.FromInputPeer2(md.UserId, request.Peer)
	switch peer.PeerType {
	case model.PEER_SELF:
	case model.PEER_USER:
	case model.PEER_CHAT:
	case model.PEER_CHANNEL:
	default:
		err := mtproto.ErrPeerIdInvalid
		log.Errorf("messages.deleteScheduledMessages - error: %v", err)
		return nil, err
	}

	if len(request.Id) > 0 {
		err := s.MessageFacade.DeleteScheduledMessageList(ctx, md.UserId, peer, request.Id)
		if err != nil {
			log.Errorf("messages.deleteScheduledMessages - error: %v", err)
			return nil, err
		}
	} else {
		log.Warnf("request.id is empty")
		return model.MakeEmptyUpdates(), nil
	}

	updateDeleteScheduledMessages := mtproto.MakeTLUpdateDeleteScheduledMessages(&mtproto.Update{
		Peer_PEER: peer.ToPeer(),
		Messages:  request.Id,
	}).To_Update()
	resultUpdates := model.MakeUpdatesByUpdates(updateDeleteScheduledMessages)

	log.Debugf("messages.deleteScheduledMessages - reply: %s", resultUpdates.DebugString())
	return model.WrapperGoFunc(resultUpdates, func() {
		sync_client.SyncUpdatesNotMe(context.Background(), md.UserId, md.AuthId, model.MakeUpdatesByUpdates(updateDeleteScheduledMessages))
	}).(*mtproto.Updates), nil
}
