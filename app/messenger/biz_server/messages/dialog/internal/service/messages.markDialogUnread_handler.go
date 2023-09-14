package service

import (
	"context"

	"time"

	sync_client "open.chat/app/messenger/sync/client"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesMarkDialogUnread(ctx context.Context, request *mtproto.TLMessagesMarkDialogUnread) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.markDialogUnread - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		err          error
		peer         *model.PeerUtil
		users        model.MutableUsers
		chat         *model.MutableChat
		channel      *model.MutableChannel
		readInboxId  int32
		readOutboxId int32
	)

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.markDialogUnread - error: %v", err)
		return nil, err
	}

	switch request.Peer.GetPredicateName() {
	case mtproto.Predicate_inputDialogPeer:
		peer = model.FromInputPeer2(md.UserId, request.Peer.Peer)
	case mtproto.Predicate_inputDialogPeerFolder:
		log.Warnf("client not send inputDialogPeerFolder: %v", request.Peer)
		return mtproto.BoolFalse, nil
	default:
		err = mtproto.ErrInputConstructorInvalid
		log.Errorf("messages.markDialogUnread - error: %v", err)
		return nil, err
	}

	switch peer.PeerType {
	case model.PEER_SELF, model.PEER_USER:
		users = s.UserFacade.GetMutableUsers(ctx, md.UserId, peer.PeerId)
		if !users.CheckExistUser(peer.PeerId) {
			log.Errorf("messages.markDialogUnread - not exist: (%d, %d)", peer.PeerId)
			err = mtproto.ErrPeerIdInvalid
			return nil, err
		}
		s.PrivateFacade.MarkDialogUnread(ctx, md.UserId, peer.PeerId, request.Unread)
	case model.PEER_CHAT:
		if chat, err = s.ChatFacade.GetMutableChat(ctx, peer.PeerId, md.UserId); err != nil {
			log.Errorf("messages.markDialogUnread - error: %v", err)
			err = mtproto.ErrPeerIdInvalid
			return nil, err
		} else if chat == nil {
			log.Errorf("messages.markDialogUnread - chat not exist: (%d, %d)", peer.PeerId)
			err = mtproto.ErrPeerIdInvalid
			return nil, err
		}

		s.ChatFacade.MarkDialogUnread(ctx, md.UserId, peer.PeerId, request.Unread)
	case model.PEER_CHANNEL:
		if channel, err = s.ChannelFacade.GetMutableChannel(ctx, peer.PeerId, md.UserId); err != nil {
			log.Errorf("messages.markDialogUnread - error: %v", err)
			err = mtproto.ErrPeerIdInvalid
			return nil, err
		} else if channel == nil {
			log.Errorf("messages.markDialogUnread - chat not exist: (%d, %d)", peer.PeerId)
			err = mtproto.ErrPeerIdInvalid
			return nil, err
		}
		s.ChannelFacade.MarkDialogUnread(ctx, md.UserId, peer.PeerId, request.Unread)
	default:
		err := mtproto.ErrPeerIdInvalid
		log.Errorf("messages.markDialogUnread - error: %v", err)
		return nil, err
	}

	log.Debugf("messages.markDialogUnread - reply {true}")
	return model.WrapperGoFunc(mtproto.BoolTrue, func() {
		syncUpdates := model.NewUpdatesLogic(md.UserId)
		updateDialogUnreadMark := mtproto.MakeTLUpdateDialogUnreadMark(&mtproto.Update{
			Unread: request.Unread,
			Peer_DIALOGPEER: mtproto.MakeTLDialogPeer(&mtproto.DialogPeer{
				Peer: peer.ToPeer(),
			}).To_DialogPeer(),
		}).To_Update()

		syncUpdates.AddUpdate(updateDialogUnreadMark)

		switch peer.PeerType {
		case model.PEER_SELF, model.PEER_USER:
			if u, _ := users.ToUnsafeUser(md.UserId, peer.PeerId); u != nil {
				syncUpdates.AddUser(u)
			}
		case model.PEER_CHAT:
			syncUpdates.AddChat(chat.ToUnsafeChat(peer.PeerId))
		case model.PEER_CHANNEL:
			syncUpdates.AddChat(channel.ToUnsafeChat(peer.PeerId))
		}
		sync_client.SyncUpdatesNotMe(context.Background(), md.UserId, md.AuthId, syncUpdates.ToUpdates())
		if !request.Unread {
			var updateReadInbox *mtproto.Update
			switch peer.PeerType {
			case model.PEER_SELF, model.PEER_USER:
				// 读收件箱
				readInboxId = s.PrivateFacade.GetTopMessage(ctx, md.UserId, peer.PeerId)
				if readInboxId > 0 {
					s.PrivateFacade.UpdateReadInbox(ctx, md.UserId, peer.PeerId, readInboxId)
				}
				updateReadInbox = mtproto.MakeTLUpdateReadHistoryInbox(&mtproto.Update{
					FolderId:         nil,
					Peer_PEER:        model.MakePeerUser(peer.PeerId),
					MaxId:            readInboxId,
					StillUnreadCount: 0,
					Pts_INT32:        int32(idgen.NextPtsId(context.Background(), md.UserId)),
					PtsCount:         1,
				}).To_Update()
			case model.PEER_CHAT:
				// 读收件箱
				readInboxId = s.ChatFacade.GetTopMessage(ctx, md.UserId, peer.PeerId)
				if readInboxId > 0 {
					s.ChatFacade.UpdateReadInbox(ctx, md.UserId, peer.PeerId, readInboxId)
				}
				updateReadInbox = mtproto.MakeTLUpdateReadHistoryInbox(&mtproto.Update{
					FolderId:         nil,
					Peer_PEER:        model.MakePeerChat(peer.PeerId),
					MaxId:            readInboxId,
					StillUnreadCount: 0,
					Pts_INT32:        int32(idgen.NextPtsId(context.Background(), md.UserId)),
					PtsCount:         1,
				}).To_Update()
			case model.PEER_CHANNEL:
				// 读收件箱
				readInboxId = s.ChannelFacade.GetTopMessage(ctx, md.UserId, peer.PeerId)
				if readInboxId > 0 {
					s.ChannelFacade.UpdateReadInbox(ctx, md.UserId, peer.PeerId, readInboxId)
				}
				updateReadInbox = mtproto.MakeTLUpdateReadChannelInbox(&mtproto.Update{
					FolderId:         nil,
					ChannelId:        peer.PeerId,
					MaxId:            readInboxId,
					StillUnreadCount: 0,
					Pts_INT32:        0,
				}).To_Update()
			}

			updatesReadInbox := mtproto.MakeTLUpdates(&mtproto.Updates{
				Updates: []*mtproto.Update{updateReadInbox},
				Users:   []*mtproto.User{},
				Chats:   []*mtproto.Chat{},
				Date:    int32(time.Now().Unix()),
				Seq:     0,
			}).To_Updates()
			sync_client.SyncUpdatesNotMe(context.Background(), md.UserId, md.AuthId, updatesReadInbox)
		}

		// updateReadOutbox
		if !request.Unread {
			var updateReadOutbox *mtproto.Update
			switch peer.PeerType {
			case model.PEER_SELF, model.PEER_USER:
				if md.UserId == peer.PeerId {
					return
				}

				// 发件箱已读
				readOutboxId = s.PrivateFacade.GetTopMessage(ctx, peer.PeerId, md.UserId)
				if readOutboxId > 0 {
					s.PrivateFacade.UpdateReadOutbox(ctx, peer.PeerId, peer.PeerId, readOutboxId)
				}
				updateReadOutbox = mtproto.MakeTLUpdateReadHistoryOutbox(&mtproto.Update{
					Peer_PEER: model.MakePeerUser(md.UserId),
					MaxId:     readOutboxId,
					Pts_INT32: int32(idgen.NextPtsId(context.Background(), peer.PeerId)),
					PtsCount:  1,
				}).To_Update()

				updatesReadOutbox := mtproto.MakeTLUpdates(&mtproto.Updates{
					Updates: []*mtproto.Update{updateReadOutbox},
					Users:   []*mtproto.User{},
					Chats:   []*mtproto.Chat{},
					Date:    int32(time.Now().Unix()),
					Seq:     0,
				}).To_Updates()
				sync_client.PushUpdates(context.Background(), md.UserId, updatesReadOutbox)
			case model.PEER_CHAT:
				usersUpdatesList := model.MakeUserUpdatesList(0)
				chat.Walk(func(userId int32, participant *model.ImmutableChatParticipant) (err error) {
					// 发件箱已读
					readOutboxId = s.PrivateFacade.GetTopMessage(ctx, peer.PeerId, md.UserId)
					if readOutboxId > 0 {
						s.PrivateFacade.UpdateReadOutbox(ctx, peer.PeerId, peer.PeerId, readOutboxId)
					}
					updateReadOutbox = mtproto.MakeTLUpdateReadHistoryOutbox(&mtproto.Update{
						Peer_PEER: model.MakePeerUser(md.UserId),
						MaxId:     readOutboxId,
						Pts_INT32: int32(idgen.NextPtsId(context.Background(), peer.PeerId)),
						PtsCount:  1,
					}).To_Update()

					usersUpdatesList.Add(userId, mtproto.MakeTLUpdates(&mtproto.Updates{
						Updates: []*mtproto.Update{updateReadOutbox},
						Users:   []*mtproto.User{},
						Chats:   []*mtproto.Chat{},
						Date:    int32(time.Now().Unix()),
						Seq:     0,
					}).To_Updates())
					return
				})
				sync_client.PushUsersUpdates(context.Background(), usersUpdatesList...)
			case model.PEER_CHANNEL:
				// 发件箱已读
				readOutboxId = s.ChannelFacade.GetTopMessage(ctx, peer.PeerId, md.UserId)
				if readOutboxId > 0 {
					s.ChannelFacade.UpdateReadOutbox(ctx, peer.PeerId, peer.PeerId, readOutboxId)
				}
				updateReadOutbox = mtproto.MakeTLUpdateReadChannelOutbox(&mtproto.Update{
					ChannelId: peer.PeerId,
					MaxId:     readOutboxId,
				}).To_Update()

				updatesReadOutbox := mtproto.MakeTLUpdates(&mtproto.Updates{
					Updates: []*mtproto.Update{updateReadOutbox},
					Users:   []*mtproto.User{},
					Chats:   []*mtproto.Chat{},
					Date:    int32(time.Now().Unix()),
					Seq:     0,
				}).To_Updates()
				sync_client.BroadcastChannelUpdates(context.Background(), peer.PeerId, updatesReadOutbox)
			}
		}
	}).(*mtproto.Bool), nil
}
