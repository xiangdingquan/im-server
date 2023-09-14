package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) getDialogs(ctx context.Context, selfUserId int32, excludePinned bool, folderId, offsetDate, offsetId int32, offsetPeer *model.PeerUtil, limit, hash int32) (model.DialogExtList, int32, error) {
	// 1. calc and load peer and chat dialogs
	peerUserDialogs := s.PrivateFacade.GetDialogs(ctx, selfUserId, excludePinned, folderId)
	//log.Debugf("getDialogs - %v", peerUserDialogs)
	peerChatDialogs := s.ChatFacade.GetDialogs(ctx, selfUserId, excludePinned, folderId)
	//log.Debugf("getDialogs - %v", peerChatDialogs)
	peerChannelDialogs := s.ChannelFacade.GetDialogs(ctx, selfUserId, excludePinned, folderId)
	//log.Debugf("getDialogs - %v", peerChannelDialogs)

	dialogs := mergeLimitDialogs(0, peerUserDialogs, peerChatDialogs, peerChannelDialogs)
	//log.Debugf("getDialogs - %s", logger.JsonDebugData(dialogs))
	if (offsetPeer.PeerType == model.PEER_EMPTY) && (offsetId == 0 || offsetId == 2147483647 || offsetDate == 0 || offsetDate == 2147483647) {
		if len(dialogs) >= int(limit) {
			dialogs = dialogs[:limit]
		}
		return dialogs, int32(len(peerUserDialogs) + len(peerChatDialogs) + len(peerChannelDialogs)), nil
	}
	offset := 0
	for i, dialog := range dialogs {
		if dialog.TopMessage == offsetId && int32(dialog.Order) == offsetDate {
			bFind := false
			switch offsetPeer.PeerType {
			case model.PEER_EMPTY:
				bFind = true
			case model.PEER_SELF, model.PEER_USER:
				bFind = dialog.Peer.UserId == offsetPeer.PeerId
			case model.PEER_CHAT:
				bFind = dialog.Peer.ChatId == offsetPeer.PeerId
			case model.PEER_CHANNEL:
				bFind = dialog.Peer.ChannelId == offsetPeer.PeerId
			}
			if bFind {
				offset = i + 1
				break
			}
		}
	}
	if offset > 0 && len(dialogs) > 1 {
		if offset+int(limit) >= len(dialogs) {
			dialogs = dialogs[offset:]
		} else {
			dialogs = dialogs[offset : offset+int(limit)]
		}
	} else {
		dialogs = dialogs[:0]
	}
	return dialogs, int32(len(peerUserDialogs) + len(peerChatDialogs) + len(peerChannelDialogs)), nil
}

func (s *Service) MessagesGetDialogs(ctx context.Context, request *mtproto.TLMessagesGetDialogs) (*mtproto.Messages_Dialogs, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getDialogs - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.getDialogs - error: %v", err)
		return nil, err
	}

	var (
		peer     = model.FromInputPeer2(md.UserId, request.OffsetPeer)
		folderId = request.GetFolderId().GetValue()
		limit    = request.Limit
	)

	if limit > 50 {
		limit = 50
	}

	switch peer.PeerType {
	case model.PEER_EMPTY:
	case model.PEER_SELF:
	case model.PEER_USER:
	case model.PEER_CHAT:
	case model.PEER_CHANNEL:
	default:
		err := mtproto.ErrInputConstructorInvalid
		log.Errorf("messages.getDialogs - error: %v", err)
		return nil, err
	}

	dialogs, count, err := s.getDialogs(ctx, md.UserId, request.ExcludePinned, folderId, request.OffsetDate, request.OffsetId, peer, limit, request.Hash)
	if err != nil {
		log.Errorf("messages.getDialogs#b098aee6 - getDialogs error: %v", err)
		return nil, err
	}

	log.Debugf("dialog: %s", logger.JsonDebugData(dialogs))

	dialogsData, err := s.ToMessageDialogs(ctx, md.UserId, dialogs)
	if err != nil {
		log.Errorf("messages.getDialogs#b098aee6 - getDialogs error: %v", err)
		return nil, err
	}

	messageDialogs := mtproto.MakeTLMessagesDialogsSlice(&mtproto.Messages_Dialogs{
		Dialogs:  dialogsData.Dialogs,
		Messages: dialogsData.Messages,
		Chats:    dialogsData.Chats,
		Users:    dialogsData.Users,
		Count:    count,
	})
	log.Debugf("messages.getDialogs#b098aee6 - reply: %s", messageDialogs.DebugString())
	return messageDialogs.To_Messages_Dialogs(), nil
}
