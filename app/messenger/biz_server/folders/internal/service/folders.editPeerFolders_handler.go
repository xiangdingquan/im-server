package service

import (
	"context"
	"sort"

	"github.com/gogo/protobuf/types"
	sync_client "open.chat/app/messenger/sync/client"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) FoldersEditPeerFolders(ctx context.Context, request *mtproto.TLFoldersEditPeerFolders) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("folders.editPeerFolders - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		err                   error
		folderId              int32 = -1
		peerUsers             []int32
		peerChats             []int32
		peerChannels          []int32
		needPinnedDialogs     model.DialogPinnedExtList
		needPeerPinnedDialogs map[int32]int64
	)

	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("folders.editPeerFolders - error: %v", err)
		return nil, err
	}

	// check FOLDER_ID_INVALID
	for _, folderPeer := range request.FolderPeers {
		if folderPeer.FolderId != 0 && folderPeer.FolderId != 1 {
			err = mtproto.ErrFolderIdInvalid
			log.Errorf("folders.editPeerFolders - error: %v", err)
			return nil, err
		}
		if folderId == -1 {
			folderId = folderPeer.FolderId
		}

		if folderPeer.FolderId != folderId {
			err = mtproto.ErrFolderIdInvalid
			log.Errorf("folders.editPeerFolders - error: %v", err)
			return nil, err
		}

		peer := model.FromInputPeer2(md.UserId, folderPeer.Peer)
		switch peer.PeerType {
		case model.PEER_SELF, model.PEER_USER:
			peerUsers = append(peerUsers, peer.PeerId)
		case model.PEER_CHAT:
			peerChats = append(peerChats, peer.PeerId)
		case model.PEER_CHANNEL:
			peerChannels = append(peerChannels, peer.PeerId)
		default:
			err := mtproto.ErrInputConstructorInvalid
			log.Errorf("folders.editPeerFolders - error: %v", err)
			return nil, err
		}
	}

	// updateFolderPeers
	updateFolderPeers := mtproto.MakeTLUpdateFolderPeers(&mtproto.Update{
		FolderPeers: make([]*mtproto.FolderPeer, 0, len(request.FolderPeers)),
	}).To_Update()

	// private
	needPeerPinnedDialogs, _ = s.PrivateFacade.EditPeerFoldersByIdList(ctx, md.UserId, folderId, peerUsers)
	for _, id := range peerUsers {
		updateFolderPeers.FolderPeers = append(updateFolderPeers.FolderPeers, mtproto.MakeTLFolderPeer(&mtproto.FolderPeer{
			Peer:     model.MakePeerUser(id),
			FolderId: folderId,
		}).To_FolderPeer())
	}

	for k, v := range needPeerPinnedDialogs {
		needPinnedDialogs = needPinnedDialogs.Add(model.PEER_USER, k, v)
	}

	// chat
	needPeerPinnedDialogs, _ = s.ChatFacade.EditPeerFoldersByIdList(ctx, md.UserId, folderId, peerChats)
	for _, id := range peerChats {
		updateFolderPeers.FolderPeers = append(updateFolderPeers.FolderPeers, mtproto.MakeTLFolderPeer(&mtproto.FolderPeer{
			Peer:     model.MakePeerChat(id),
			FolderId: folderId,
		}).To_FolderPeer())
	}

	for k, v := range needPeerPinnedDialogs {
		needPinnedDialogs = needPinnedDialogs.Add(model.PEER_CHAT, k, v)
	}

	// channel
	needPeerPinnedDialogs, _ = s.ChannelFacade.EditPeerFoldersByIdList(ctx, md.UserId, folderId, peerChannels)
	for _, id := range peerChannels {
		updateFolderPeers.FolderPeers = append(updateFolderPeers.FolderPeers, mtproto.MakeTLFolderPeer(&mtproto.FolderPeer{
			Peer:     model.MakePeerChannel(id),
			FolderId: folderId,
		}).To_FolderPeer())
	}
	updateFolderPeers.Pts_INT32 = int32(idgen.NextPtsId(ctx, md.UserId))
	updateFolderPeers.PtsCount = 1

	for k, v := range needPeerPinnedDialogs {
		needPinnedDialogs = needPinnedDialogs.Add(model.PEER_CHANNEL, k, v)
	}

	// updatePinnedDialogs
	updatePinnedDialogs := mtproto.MakeTLUpdatePinnedDialogs(&mtproto.Update{
		FolderId:                   nil,
		Order_FLAGVECTORDIALOGPEER: []*mtproto.DialogPeer{},
	}).To_Update()

	log.Debugf("needPinnedDialogs - %v", needPinnedDialogs)
	if len(needPinnedDialogs) > 0 {
		var (
			pinnedDialogs model.DialogPinnedExtList
		)

		pinnedDialogs = append(pinnedDialogs, s.PrivateFacade.GetPinnedDialogPeers(ctx, md.UserId, folderId)...)
		pinnedDialogs = append(pinnedDialogs, s.ChatFacade.GetPinnedDialogPeers(ctx, md.UserId, folderId)...)
		pinnedDialogs = append(pinnedDialogs, s.ChannelFacade.GetPinnedDialogPeers(ctx, md.UserId, folderId)...)

		c := sort.Reverse(pinnedDialogs)
		sort.Sort(c)

		if folderId == 1 {
			updatePinnedDialogs.FolderId = &types.Int32Value{Value: folderId}
		} else {
			// mark
			for i := 5; i < len(pinnedDialogs); i++ {
				switch pinnedDialogs[i].PeerType {
				case model.PEER_USER:
					s.PrivateFacade.ToggleDialogPin(ctx, md.UserId, pinnedDialogs[i].PeerId, false)
				case model.PEER_CHAT:
					s.ChatFacade.ToggleDialogPin(ctx, md.UserId, pinnedDialogs[i].PeerId, false)
				case model.PEER_CHANNEL:
					s.ChannelFacade.ToggleDialogPin(ctx, md.UserId, pinnedDialogs[i].PeerId, false)
				}
			}

			// cut
			if len(pinnedDialogs) > 5 {
				pinnedDialogs = pinnedDialogs[:5]
			}
		}

		for _, d := range pinnedDialogs {
			switch d.PeerType {
			case model.PEER_USER:
				updatePinnedDialogs.Order_FLAGVECTORDIALOGPEER = append(updatePinnedDialogs.Order_FLAGVECTORDIALOGPEER, mtproto.MakeTLDialogPeer(&mtproto.DialogPeer{
					FolderId: folderId,
					Peer:     model.MakePeerUser(d.PeerId),
				}).To_DialogPeer())
				peerUsers = append(peerUsers, d.PeerId)
			case model.PEER_CHAT:
				updatePinnedDialogs.Order_FLAGVECTORDIALOGPEER = append(updatePinnedDialogs.Order_FLAGVECTORDIALOGPEER, mtproto.MakeTLDialogPeer(&mtproto.DialogPeer{
					FolderId: folderId,
					Peer:     model.MakePeerChat(d.PeerId),
				}).To_DialogPeer())
				peerChats = append(peerChats, d.PeerId)
			case model.PEER_CHANNEL:
				updatePinnedDialogs.Order_FLAGVECTORDIALOGPEER = append(updatePinnedDialogs.Order_FLAGVECTORDIALOGPEER, mtproto.MakeTLDialogPeer(&mtproto.DialogPeer{
					FolderId: folderId,
					Peer:     model.MakePeerChannel(d.PeerId),
				}).To_DialogPeer())
				peerChannels = append(peerChannels, d.PeerId)
			}
		}
	}

	// make updates
	updates := model.NewUpdatesLogic(md.UserId)

	// updates
	updates.AddUpdate(updateFolderPeers)
	updates.AddUpdate(updatePinnedDialogs)

	// user
	users := s.UserFacade.GetUserListByIdList(ctx, md.UserId, peerUsers)
	updates.AddUsers(users)

	// chat
	chats := s.ChatFacade.GetChatListByIdList(ctx, md.UserId, peerChats)
	updates.AddChats(chats)

	// channel
	updates.AddChats(s.ChannelFacade.GetChannelListByIdList(ctx, md.UserId, peerChannels...))

	reply := updates.ToUpdates()
	log.Debugf("folders.editPeerFolders - reply: %s", reply.DebugString())
	return model.WrapperGoFunc(reply, func() {
		sync_client.SyncUpdatesNotMe(context.Background(), md.UserId, md.AuthId, reply)
	}).(*mtproto.Updates), nil
}
