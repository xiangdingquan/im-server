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

func (s *Service) MessagesSaveDraft(ctx context.Context, request *mtproto.TLMessagesSaveDraft) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.saveDraft - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		peer                = model.FromInputPeer2(md.UserId, request.Peer)
		draft               *mtproto.DraftMessage
		isDraftMessageEmpty = true
		date                = int32(time.Now().Unix())
		users               model.MutableUsers
		chat                *model.MutableChat
		channel             *model.MutableChannel
		err                 error
	)

	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("messages.saveDraft - error: %v", err)
		return nil, err
	}

	if request.NoWebpage == true {
		isDraftMessageEmpty = false
	} else if request.ReplyToMsgId != nil {
		isDraftMessageEmpty = false
	} else if request.Message != "" {
		isDraftMessageEmpty = false
	} else if request.Entities != nil {
		isDraftMessageEmpty = false
	}

	if isDraftMessageEmpty {
		draft = model.MakeDraftMessageEmpty(date)
	} else {
		draft = mtproto.MakeTLDraftMessage(&mtproto.DraftMessage{
			NoWebpage:    request.GetNoWebpage(),
			ReplyToMsgId: request.GetReplyToMsgId(),
			Message:      request.GetMessage(),
			Entities:     request.GetEntities(),
			Date_INT32:   int32(time.Now().Unix()),
		}).To_DraftMessage()
	}

	switch peer.PeerType {
	case model.PEER_SELF, model.PEER_USER:
		users = s.UserFacade.GetMutableUsers(ctx, md.UserId, peer.PeerId)
		if !users.CheckExistUser(peer.PeerId) {
			log.Errorf("messages.saveDraft - not exist: (%d, %d)", peer.PeerId)
			err = mtproto.ErrPeerIdInvalid
			return nil, err
		}

		if isDraftMessageEmpty {
			s.PrivateFacade.ClearDraftMessage(ctx, md.UserId, peer.PeerId)
		} else {
			s.PrivateFacade.SaveDraftMessage(ctx, md.UserId, peer.PeerId, draft)
		}
	case model.PEER_CHAT:
		if chat, err = s.ChatFacade.GetMutableChat(ctx, peer.PeerId, md.UserId); err != nil {
			log.Errorf("messages.saveDraft - error: %v", err)
			err = mtproto.ErrPeerIdInvalid
			return nil, err
		} else if chat == nil {
			log.Errorf("messages.saveDraft - chat not exist: (%d, %d)", peer.PeerId)
			err = mtproto.ErrPeerIdInvalid
			return nil, err
		}

		if isDraftMessageEmpty {
			s.ChatFacade.ClearDraftMessage(ctx, md.UserId, peer.PeerId)
		} else {
			s.ChatFacade.SaveDraftMessage(ctx, md.UserId, peer.PeerId, draft)
		}
	case model.PEER_CHANNEL:
		if channel, err = s.ChannelFacade.GetMutableChannel(ctx, peer.PeerId, md.UserId); err != nil {
			log.Errorf("messages.saveDraft - error: %v", err)
			err = mtproto.ErrPeerIdInvalid
			return nil, err
		} else if channel == nil {
			log.Errorf("messages.saveDraft - chat not exist: (%d, %d)", peer.PeerId)
			err = mtproto.ErrPeerIdInvalid
			return nil, err
		}

		if isDraftMessageEmpty {
			s.ChannelFacade.ClearDraftMessage(ctx, md.UserId, peer.PeerId)
		} else {
			s.ChannelFacade.SaveDraftMessage(ctx, md.UserId, peer.PeerId, draft)
		}
	default:
		err = mtproto.ErrPeerIdInvalid
		log.Errorf("messages.saveDraft - error: %v", err)
		return nil, err
	}

	log.Debugf("messages.saveDraft#bc39e14b - reply: {true}")
	return model.WrapperGoFunc(mtproto.BoolTrue, func() {
		syncUpdates := model.NewUpdatesLogic(md.UserId)
		updateDraftMessage := mtproto.MakeTLUpdateDraftMessage(&mtproto.Update{
			Peer_PEER: peer.ToPeer(),
			Draft:     draft,
		})
		syncUpdates.AddUpdate(updateDraftMessage.To_Update())

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
