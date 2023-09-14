package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) getPeerDialogs(ctx context.Context, selfUserId int32, peers []*model.PeerUtil) (model.DialogExtList, error) {
	var (
		peerUsers    []int32
		peerChats    []int32
		peerChannels []int32
	)

	for _, peer := range peers {
		switch peer.PeerType {
		case model.PEER_SELF, model.PEER_USER:
			peerUsers = append(peerUsers, peer.PeerId)
		case model.PEER_CHAT:
			peerChats = append(peerChats, peer.PeerId)
		case model.PEER_CHANNEL:
			peerChannels = append(peerChannels, peer.PeerId)
		}
	}

	// 1. load private dialogs
	privateDialogs := s.PrivateFacade.GetDialogsByIdList(ctx, selfUserId, peerUsers)

	// 2. load chat dialogs
	chatDialogs := s.ChatFacade.GetDialogsByIdList(ctx, selfUserId, peerChats)

	// 3. load channel dialogs
	channelDialogs := s.ChannelFacade.GetDialogsByIdList(ctx, selfUserId, peerChannels)

	return mergeLimitDialogs(0, privateDialogs, chatDialogs, channelDialogs), nil
}

func (s *Service) getPeerDialogFolder(ctx context.Context, selfUserId, folderId int32) (model.DialogExtList, error) {
	var (
		hasDialogFolder = false
		dialogFolder    = &model.DialogExt{
			Dialog: mtproto.MakeTLDialogFolder(&mtproto.Dialog{
				Pinned: false,
				Folder: mtproto.MakeTLFolder(&mtproto.Folder{
					AutofillNewBroadcasts:     false,
					AutofillPublicGroups:      false,
					AutofillNewCorrespondents: false,
					Id:                        folderId,
					Title:                     "Archived Chats",
					Photo:                     nil,
				}).To_Folder(),
				Peer:                       nil,
				TopMessage:                 0,
				UnreadMutedPeersCount:      0,
				UnreadUnmutedPeersCount:    0,
				UnreadMutedMessagesCount:   0,
				UnreadUnmutedMessagesCount: 0,
			}).To_Dialog(),
			Order: -1,
		}
	)

	// 1. load private dialogs
	if privateDialog, _ := s.PrivateFacade.GetDialogFolder(ctx, selfUserId, folderId); privateDialog != nil {
		if privateDialog.Order > dialogFolder.Order {
			hasDialogFolder = true
			dialogFolder.Order = privateDialog.Order
			dialogFolder.Dialog.Peer = privateDialog.Peer
			dialogFolder.Dialog.TopMessage = privateDialog.TopMessage
		}
		dialogFolder.UnreadUnmutedPeersCount += privateDialog.UnreadUnmutedPeersCount
		dialogFolder.UnreadUnmutedMessagesCount += privateDialog.UnreadUnmutedMessagesCount
	}

	// 2. load chat dialogs
	if chatDialog, _ := s.ChatFacade.GetDialogFolder(ctx, selfUserId, folderId); chatDialog != nil {
		if chatDialog.Order > dialogFolder.Order {
			hasDialogFolder = true
			dialogFolder.Order = chatDialog.Order
			dialogFolder.Peer = chatDialog.Peer
			dialogFolder.TopMessage = chatDialog.TopMessage
		}
		dialogFolder.UnreadUnmutedPeersCount += chatDialog.UnreadUnmutedPeersCount
		dialogFolder.UnreadUnmutedMessagesCount += chatDialog.UnreadUnmutedMessagesCount
	}

	// 3. load channel dialogs
	if channelDialog, _ := s.ChannelFacade.GetDialogFolder(ctx, selfUserId, folderId); channelDialog != nil {
		if channelDialog.Order > dialogFolder.Order {
			hasDialogFolder = true
			dialogFolder.Order = channelDialog.Order
			dialogFolder.Peer = channelDialog.Peer
			dialogFolder.TopMessage = channelDialog.TopMessage
		}
		dialogFolder.UnreadUnmutedPeersCount += channelDialog.UnreadUnmutedPeersCount
		dialogFolder.UnreadUnmutedMessagesCount += channelDialog.UnreadUnmutedMessagesCount
	}

	if !hasDialogFolder {
		return []*model.DialogExt{}, nil
	} else {
		return []*model.DialogExt{dialogFolder}, nil
	}
}

func (s *Service) MessagesGetPeerDialogs(ctx context.Context, request *mtproto.TLMessagesGetPeerDialogs) (*mtproto.Messages_PeerDialogs, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getPeerDialogs - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.getPeerDialogs - error: %v", err)
		return nil, err
	}

	var (
		peers    []*model.PeerUtil
		folderId int32 = -1
	)
	switch request.Constructor {
	case mtproto.CRC32_messages_getPeerDialogs_2d9776b9:
		for _, peer := range request.Peers_VECTORINPUTPEER {
			p := model.FromInputPeer2(md.UserId, peer)
			switch p.PeerType {
			case model.PEER_SELF:
			case model.PEER_USER:
			case model.PEER_CHAT:
			case model.PEER_CHANNEL:
			default:
				err := mtproto.ErrPeerIdInvalid
				log.Errorf("messages.getPeerDialogs - getPeerDialogs error: %v", err)
				return nil, err
			}
			peers = append(peers, p)
		}
	case mtproto.CRC32_messages_getPeerDialogs_e470bcfd:
		for _, peer := range request.Peers_VECTORINPUTDIALOGPEER {
			switch peer.PredicateName {
			case mtproto.Predicate_inputDialogPeer:
				p := model.FromInputPeer2(md.UserId, peer.Peer)
				switch p.PeerType {
				case model.PEER_SELF:
				case model.PEER_USER:
				case model.PEER_CHAT:
				case model.PEER_CHANNEL:
				default:
					err := mtproto.ErrPeerIdInvalid
					log.Errorf("messages.getPeerDialogs - getPeerDialogs error: %v", err)
					return nil, err
				}
				peers = append(peers, p)
			case mtproto.Predicate_inputDialogPeerFolder:
				if folderId == -1 {
					folderId = peer.FolderId
				} else {
					err := mtproto.ErrFolderIdInvalid
					log.Errorf("messages.getPeerDialogs - getPeerDialogs error: %v", err)
					return nil, err
				}
			default:
				err := mtproto.ErrPeerIdInvalid
				log.Errorf("messages.getPeerDialogs - getPeerDialogs error: %v", err)
				return nil, err
			}
		}
	default:
		err := mtproto.ErrInputConstructorInvalid
		log.Errorf("messages.getPeerDialogs - error: %v", err)
		return nil, err
	}

	if folderId == 0 {
		err := mtproto.ErrFolderIdInvalid
		log.Errorf("messages.getPeerDialogs - getPeerDialogs error: %v", err)
		return nil, err
	}

	var (
		err     error
		dialogs model.DialogExtList
		state   *mtproto.Updates_State
	)

	state, err = s.UpdatesFacade.GetState(ctx, md.AuthId, md.UserId)
	if err != nil {
		log.Errorf("messages.getPeerDialogs - getState error: %v", err)
		return nil, err
	}

	if folderId == 1 {
		dialogs, err = s.getPeerDialogFolder(ctx, md.UserId, folderId)
		if err != nil {
			log.Errorf("messages.getPeerDialogs - getPeerDialogs error: %v", err)
			return nil, err
		}
	} else {
		dialogs, err = s.getPeerDialogs(ctx, md.UserId, peers)
		if err != nil {
			log.Errorf("messages.getPeerDialogs - getPeerDialogs error: %v", err)
			return nil, err
		}
	}

	dialogsData, err := s.ToMessageDialogs(ctx, md.UserId, dialogs)
	if err != nil {
		log.Errorf("messages.getPeerDialogs - getPeerDialogs error: %v", err)
		return nil, err
	}

	peerDialogs := mtproto.MakeTLMessagesPeerDialogs(&mtproto.Messages_PeerDialogs{
		Dialogs:  dialogsData.Dialogs,
		Messages: dialogsData.Messages,
		Users:    dialogsData.Users,
		Chats:    dialogsData.Chats,
		State:    state,
	}).To_Messages_PeerDialogs()

	log.Debugf("messages.getPeerDialogs - reply: %s", peerDialogs.DebugString())
	return peerDialogs, nil
}
