package service

import (
	"context"

	"github.com/gogo/protobuf/types"
	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesToggleDialogPin(ctx context.Context, request *mtproto.TLMessagesToggleDialogPin) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.toggleDialogPin - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		err          error
		peer         *model.PeerUtil
		users        model.MutableUsers
		chat         *model.MutableChat
		channel      *model.MutableChannel
		flagFolderId *types.Int32Value
	)

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.toggleDialogPin - error: %v", err)
		return nil, err
	}

	switch request.Constructor {
	case mtproto.CRC32_messages_toggleDialogPin_3289be6a:
		peer = model.FromInputPeer2(md.UserId, request.Peer_INPUTPEER)
	case mtproto.CRC32_messages_toggleDialogPin_a731e257:
		switch request.Peer_INPUTDIALOGPEER.GetPredicateName() {
		case mtproto.Predicate_inputDialogPeer:
			peer = model.FromInputPeer2(md.UserId, request.Peer_INPUTDIALOGPEER.Peer)
		case mtproto.Predicate_inputDialogPeerFolder:
			log.Warnf("client not send inputDialogPeerFolder: %v", request.Peer_INPUTDIALOGPEER)
			return mtproto.BoolFalse, nil
		default:
			err = mtproto.ErrInputConstructorInvalid
			log.Errorf("messages.toggleDialogPin - error: %v", err)
			return nil, err
		}
	default:
		err = mtproto.ErrInputConstructorInvalid
		log.Errorf("messages.toggleDialogPin - error: %v", err)
		return nil, err
	}

	switch peer.PeerType {
	case model.PEER_SELF, model.PEER_USER:
		users = s.UserFacade.GetMutableUsers(ctx, md.UserId, peer.PeerId)
		if !users.CheckExistUser(peer.PeerId) {
			log.Errorf("messages.toggleDialogPin - not exist: (%d, %d)", peer.PeerId)
			err = mtproto.ErrPeerIdInvalid
			return nil, err
		}

		folderId, _ := s.PrivateFacade.ToggleDialogPin(ctx, md.UserId, peer.PeerId, request.Pinned)
		if folderId == 1 {
			flagFolderId = &types.Int32Value{Value: folderId}
		}
	case model.PEER_CHAT:
		if chat, err = s.ChatFacade.GetMutableChat(ctx, peer.PeerId, md.UserId); err != nil {
			log.Errorf("messages.toggleDialogPin - error: %v", err)
			err = mtproto.ErrPeerIdInvalid
			return nil, err
		} else if chat == nil {
			log.Errorf("messages.toggleDialogPin - chat not exist: (%d, %d)", peer.PeerId)
			err = mtproto.ErrPeerIdInvalid
			return nil, err
		}

		folderId, _ := s.ChatFacade.ToggleDialogPin(ctx, md.UserId, peer.PeerId, request.Pinned)
		if folderId == 1 {
			flagFolderId = &types.Int32Value{Value: folderId}
		}
	case model.PEER_CHANNEL:
		if channel, err = s.ChannelFacade.GetMutableChannel(ctx, peer.PeerId, md.UserId); err != nil {
			log.Errorf("messages.toggleDialogPin - error: %v", err)
			err = mtproto.ErrPeerIdInvalid
			return nil, err
		} else if channel == nil {
			log.Errorf("messages.toggleDialogPin - chat not exist: (%d, %d)", peer.PeerId)
			err = mtproto.ErrPeerIdInvalid
			return nil, err
		}

		folderId, _ := s.ChannelFacade.ToggleDialogPin(ctx, md.UserId, peer.PeerId, request.Pinned)
		if folderId == 1 {
			flagFolderId = &types.Int32Value{Value: folderId}
		}
	default:
		err := mtproto.ErrPeerIdInvalid
		log.Errorf("messages.toggleDialogPin - error: %v", err)
		return nil, err
	}

	log.Debugf("messages.toggleDialogPin - reply {true}")
	return model.WrapperGoFunc(mtproto.BoolTrue, func() {
		syncUpdates := model.NewUpdatesLogic(md.UserId)
		updateDialogPinned := mtproto.MakeTLUpdateDialogPinned(&mtproto.Update{
			Pinned:   request.GetPinned(),
			FolderId: flagFolderId,
			Peer_DIALOGPEER: mtproto.MakeTLDialogPeer(&mtproto.DialogPeer{
				Peer: peer.ToPeer(),
			}).To_DialogPeer(),
		}).To_Update()

		syncUpdates.AddUpdate(updateDialogPinned)

		switch peer.PeerType {
		case model.PEER_SELF, model.PEER_USER:
			syncUpdates.AddUser(model.GetFirstValue(users.ToUnsafeUser(md.UserId, peer.PeerId)).(*mtproto.User))
		case model.PEER_CHAT:
			syncUpdates.AddChat(chat.ToUnsafeChat(peer.PeerId))
		case model.PEER_CHANNEL:
			syncUpdates.AddChat(channel.ToUnsafeChat(peer.PeerId))
		}
		sync_client.SyncUpdatesNotMe(context.Background(), md.UserId, md.AuthId, syncUpdates.ToUpdates())
	}).(*mtproto.Bool), nil
}
