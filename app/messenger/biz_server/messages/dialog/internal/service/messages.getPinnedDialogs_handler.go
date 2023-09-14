package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) getPinnedDialogs(ctx context.Context, selfUserId, folderId int32) (model.DialogExtList, error) {
	// 1. load private dialogs
	privateDialogs := s.PrivateFacade.GetPinnedDialogs(ctx, selfUserId, folderId)

	// 2. load chat dialogs
	chatDialogs := s.ChatFacade.GetPinnedDialogs(ctx, selfUserId, folderId)

	// 3. load channel dialogs
	channelDialogs := s.ChannelFacade.GetPinnedDialogs(ctx, selfUserId, folderId)

	return mergeLimitDialogs(0, privateDialogs, chatDialogs, channelDialogs), nil
}

func (s *Service) MessagesGetPinnedDialogs(ctx context.Context, request *mtproto.TLMessagesGetPinnedDialogs) (*mtproto.Messages_PeerDialogs, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getPinnedDialogs - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.getPinnedDialogs - error: %v", err)
		return nil, err
	}

	if request.FolderId != 0 && request.FolderId != 1 {
		err := mtproto.ErrFolderIdInvalid
		log.Errorf("messages.getPinnedDialogs - error: %v", err)
		return nil, err
	}

	state, err := s.UpdatesFacade.GetState(ctx, md.AuthId, md.UserId)
	if err != nil {
		log.Errorf("messages.getPinnedDialogs - getState error: %v", err)
		return nil, err
	}

	var dialogs model.DialogExtList

	if request.FolderId == 0 {
		if dialogFolder, err := s.getPeerDialogFolder(ctx, md.UserId, 1); err != nil {
			log.Errorf("messages.getPinnedDialogs - getPinnedDialogs error: %v", err)
			return nil, err
		} else if len(dialogFolder) > 0 {
			dialogs = append(dialogs, dialogFolder...)
		}
	}

	pinnedDialogs, err := s.getPinnedDialogs(ctx, md.UserId, request.FolderId)
	if err != nil {
		log.Errorf("messages.getPinnedDialogs - getPinnedDialogs error: %v", err)
		return nil, err
	}
	dialogs = append(dialogs, pinnedDialogs...)

	dialogsData, err := s.ToMessageDialogs(ctx, md.UserId, dialogs)
	if err != nil {
		log.Errorf("messages.getPinnedDialogs#e254d64e - getPeerDialogs error: %v", err)
		return nil, err
	}

	peerDialogs := mtproto.MakeTLMessagesPeerDialogs(&mtproto.Messages_PeerDialogs{
		Dialogs:  dialogsData.Dialogs,
		Messages: dialogsData.Messages,
		Users:    dialogsData.Users,
		Chats:    dialogsData.Chats,
		State:    state,
	}).To_Messages_PeerDialogs()

	log.Debugf("messages.getPinnedDialogs - reply: %s", peerDialogs.DebugString())
	return peerDialogs, nil
}
