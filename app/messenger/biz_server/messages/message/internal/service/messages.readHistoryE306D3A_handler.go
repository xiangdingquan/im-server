package service

import (
	"context"

	sync_client "open.chat/app/messenger/sync/client"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) readHistory(ctx context.Context, md *grpc_util.RpcMetadata, peer *model.PeerUtil, maxId, offset int32) (*mtproto.Messages_AffectedMessages, error) {
	var (
		pts, ptsCount  int32
		chat           *model.MutableChat
		isTopMessage   bool
		err            error
		readInboxMaxId int32 = -1
	)

	switch peer.PeerType {
	case model.PEER_USER:
		if maxId == 0 {
			maxId = s.PrivateFacade.GetTopMessage(ctx, md.UserId, peer.PeerId)
		}
		readInboxMaxId = s.PrivateFacade.GetReadInboxMaxId(ctx, md.UserId, peer.PeerId)
		s.PrivateFacade.UpdateReadInbox(ctx, md.UserId, peer.PeerId, maxId)
	case model.PEER_CHAT:
		chat, err = s.ChatFacade.GetMutableChat(ctx, peer.PeerId, md.UserId)
		if err != nil {
			log.Errorf("not found chat_id: %d", peer.PeerId)
			return nil, err
		}

		p := chat.GetImmutableChatParticipant(md.UserId)
		if p == nil {
			err = mtproto.ErrUserNotParticipant
			log.Errorf("{chat_id: %d, user_id: %d} not existed: %v", md.UserId, peer.PeerId, err)
			return nil, err
		}

		if maxId == 0 {
			isTopMessage = true
			maxId = p.Dialog.TopMessage
		}
		readInboxMaxId = s.ChatFacade.GetReadInboxMaxId(ctx, md.UserId, peer.PeerId)
		s.ChatFacade.UpdateReadInbox(ctx, md.UserId, peer.PeerId, maxId)
	}
	if readInboxMaxId >= 0 {
		s.MessageFacade.ReadEphemeralMsgByBetween(ctx, md.UserId, peer, readInboxMaxId, maxId)
	}
	//
	pts = int32(idgen.NextPtsId(ctx, md.UserId))
	ptsCount = 1

	go func() {
		ctx2 := context.Background()
		syncUpdatesHelper := model.MakeUpdatesHelper(mtproto.MakeTLUpdateReadHistoryInbox(&mtproto.Update{
			Peer_PEER: peer.ToPeer(),
			MaxId:     maxId,
			Pts_INT32: pts,
			PtsCount:  ptsCount,
		}).To_Update())

		err := sync_client.SyncUpdatesNotMe(ctx2,
			md.UserId,
			md.AuthId,
			syncUpdatesHelper.ToSyncNotMeUpdates(ctx2, md.UserId, nil, nil, nil))
		if err != nil {
			log.Error(err.Error())
		}

		switch peer.PeerType {
		case model.PEER_USER:
			outboxMaxMessageId := s.MessageFacade.GetPeerUserMessageId(ctx2, md.UserId, maxId, peer.PeerId)
			s.PrivateFacade.UpdateReadOutbox(ctx2, peer.PeerId, md.UserId, outboxMaxMessageId)

			pushUpdatesHelper := model.MakeUpdatesHelper(mtproto.MakeTLUpdateReadHistoryOutbox(&mtproto.Update{
				Peer_PEER: model.MakePeerUser(md.UserId),
				MaxId:     outboxMaxMessageId,
				Pts_INT32: int32(idgen.NextPtsId(ctx, peer.PeerId)),
				PtsCount:  1,
			}).To_Update())

			sync_client.PushUpdates(ctx2,
				peer.PeerId,
				pushUpdatesHelper.ToPushUpdates(ctx2, peer.PeerId, nil, nil, nil))
		case model.PEER_CHAT:
			chat.Walk(func(userId int32, participant *model.ImmutableChatParticipant) error {
				if userId == md.UserId {
					return nil
				}

				outboxMaxMessageId := participant.Dialog.TopMessage
				if !isTopMessage {
					outboxMaxMessageId = s.MessageFacade.GetPeerUserMessageId(ctx2, md.UserId, maxId, userId)
				}
				s.ChatFacade.UpdateReadOutbox(ctx2, userId, peer.PeerId, outboxMaxMessageId)

				pushUpdatesHelper := model.MakeUpdatesHelper(mtproto.MakeTLUpdateReadHistoryOutbox(&mtproto.Update{
					Peer_PEER: model.MakePeerChat(peer.PeerId),
					MaxId:     outboxMaxMessageId,
					Pts_INT32: int32(idgen.NextPtsId(ctx, userId)),
					PtsCount:  1,
				}).To_Update())

				sync_client.PushUpdates(ctx2,
					userId,
					pushUpdatesHelper.ToPushUpdates(ctx2, userId, nil, nil, nil))
				return nil
			})
		}
	}()

	affected := mtproto.MakeTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
		Pts:      pts,
		PtsCount: ptsCount,
	})

	return affected.To_Messages_AffectedMessages(), nil
}

func (s *Service) MessagesReadHistoryE306D3A(ctx context.Context, request *mtproto.TLMessagesReadHistoryE306D3A) (*mtproto.Messages_AffectedMessages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.readHistory#e306d3a - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		err  error
		peer = model.FromInputPeer2(md.UserId, request.Peer)
	)

	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("messages.readHistory - error: %v", err)
		return nil, err
	}

	switch peer.PeerType {
	case model.PEER_SELF:
		peer.PeerType = model.PEER_USER
	case model.PEER_USER:
	case model.PEER_CHAT:
	default:
		log.Errorf("invalid peer: %v", request.Peer)
		err = mtproto.ErrPeerIdInvalid
		return nil, err
	}

	affected, err := s.readHistory(ctx, md, peer, request.MaxId, 0)
	if err != nil {
		log.Errorf("messages.readHistory - readHistory error: %v", err)
		return nil, err
	}

	log.Debugf("messages.readHistory - reply: {%s}", affected.DebugString())
	return affected, nil
}
