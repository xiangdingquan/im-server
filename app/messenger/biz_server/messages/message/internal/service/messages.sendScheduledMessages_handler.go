package service

import (
	"context"

	"open.chat/app/messenger/msg/msgpb"
	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesSendScheduledMessages(ctx context.Context, request *mtproto.TLMessagesSendScheduledMessages) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.sendScheduledMessages - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	peer := model.FromInputPeer2(md.UserId, request.Peer)
	switch peer.PeerType {
	case model.PEER_SELF:
		peer.PeerType = model.PEER_USER
	case model.PEER_USER:
	case model.PEER_CHAT:
	case model.PEER_CHANNEL:
	default:
		err := mtproto.ErrPeerIdInvalid
		log.Errorf("messages.sendScheduledMessages - error: %v", err)
		return nil, err
	}

	if len(request.Id) != 1 {
		err := mtproto.ErrMessageIdInvalid
		log.Errorf("messages.sendScheduledMessages - error: %v", err)
		return nil, err
	}

	msgBoxList := s.MessageFacade.GetScheduledMessageListByIdList(ctx, md.UserId, peer, request.Id)
	if len(msgBoxList) == 0 {
		err := mtproto.ErrMessageIdInvalid
		log.Errorf("messages.sendScheduledMessages - error: %v", err)
		return nil, err
	}

	var (
		resultUpdates *mtproto.Updates
		err           error
	)

	resultUpdates, err = s.MsgFacade.SendMessage(ctx, md.UserId, md.AuthId, peer, &msgpb.OutboxMessage{
		NoWebpage:    false,
		Background:   false,
		RandomId:     msgBoxList[0].RandomId,
		Message:      msgBoxList[0].ToMessage(md.UserId),
		ScheduleDate: nil,
	})
	if err != nil {
		log.Errorf("messages.sendScheduledMessages - error: %v", err)
		return nil, err
	}

	return model.WrapperGoFunc(resultUpdates, func() {
		sync_client.PushUpdates(context.Background(), md.UserId, model.MakeUpdatesByUpdates(mtproto.MakeTLUpdateDeleteScheduledMessages(&mtproto.Update{
			Peer_PEER: peer.ToPeer(),
			Messages:  request.Id,
		}).To_Update()))
	}).(*mtproto.Updates), nil
}
