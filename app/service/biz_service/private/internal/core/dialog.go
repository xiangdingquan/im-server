package core

import (
	"context"
	"encoding/json"
	"fmt"

	"time"

	"github.com/gogo/protobuf/types"
	"open.chat/app/service/biz_service/private/internal/dal/dataobject"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

// ////////////////////////////////////////////////////////////////////////////////////
func makeDialog(dialogDO *dataobject.ConversationsDO) *mtproto.Dialog {
	dialog := mtproto.MakeTLDialog(&mtproto.Dialog{
		Pinned:              util.Int8ToBool(dialogDO.IsPinned),
		UnreadMark:          util.Int8ToBool(dialogDO.UnreadMark),
		Peer:                model.MakePeerUser(dialogDO.PeerId),
		TopMessage:          dialogDO.TopMessage,
		ReadInboxMaxId:      dialogDO.ReadInboxMaxId,
		ReadOutboxMaxId:     dialogDO.ReadOutboxMaxId,
		UnreadCount:         dialogDO.UnreadCount,
		UnreadMentionsCount: 0,
		NotifySettings:      nil,
		Pts:                 nil,
		Draft:               nil,
		FolderId:            nil,
	}).To_Dialog()

	if dialogDO.DraftType == 2 {
		draft := &mtproto.DraftMessage{}
		err := json.Unmarshal([]byte(dialogDO.DraftMessageData), &draft)
		if err == nil {
			dialog.Draft = draft
		} else {
			dialog.Draft = mtproto.MakeTLDraftMessageEmpty(draft).To_DraftMessage()
		}
	} else if dialogDO.DraftType == 1 {
		dialog.Draft = mtproto.MakeTLDraftMessageEmpty(nil).To_DraftMessage()
	}

	dialog.NotifySettings = mtproto.MakeTLPeerNotifySettings(&mtproto.PeerNotifySettings{}).To_PeerNotifySettings()

	if dialogDO.FolderId != 0 {
		dialog.FolderId = &types.Int32Value{Value: dialogDO.FolderId}
	}

	return dialog
}

func makeDialogList(i model.DialogExtList, dialogDOList []dataobject.ConversationsDO) (dialogs model.DialogExtList) {
	if len(i) == 0 {
		dialogs = []*model.DialogExt{}
	} else {
		dialogs = i
	}

	for i := 0; i < len(dialogDOList); i++ {
		dialogs = append(dialogs, &model.DialogExt{Order: int64(dialogDOList[i].Date2), Dialog: makeDialog(&dialogDOList[i])})
	}
	return
}

// ////////////////////////////////////////////////////////////////////////////////////
func (m *PrivateCore) UpdateReadInbox(ctx context.Context, userId, peerId int32, readInboxId int32) {
	m.ConversationsDAO.UpdateUnreadByPeer(ctx, readInboxId, userId, peerId)
}

func (m *PrivateCore) UpdateReadOutbox(ctx context.Context, userId, peerId int32, readOutboxId int32) {
	m.ConversationsDAO.UpdateReadOutboxMaxIdByPeer(ctx, readOutboxId, userId, peerId)
}

func (m *PrivateCore) GetTopMessage(ctx context.Context, userId int32, peerId int32) int32 {
	var (
		topMessage int32
	)

	do, _ := m.ConversationsDAO.SelectByPeer(ctx, userId, peerId)
	if do != nil {
		topMessage = do.TopMessage
	}

	return topMessage
}

func (m *PrivateCore) GetReadInboxMaxId(ctx context.Context, userId, peerId int32) int32 {
	var (
		readInboxMaxId int32
	)

	do, _ := m.ConversationsDAO.SelectByPeer(ctx, userId, peerId)
	if do != nil {
		readInboxMaxId = do.ReadInboxMaxId
	}

	return readInboxMaxId
}

func (m *PrivateCore) ResetTopMessage(ctx context.Context, userId int32, peerId int32) {
	m.ConversationsDAO.UpdateCustomMap(ctx, map[string]interface{}{
		"top_message":         0,
		"is_pinned":           0,
		"order_pinned":        0,
		"folder_pinned":       0,
		"folder_order_pinned": 0,
		"unread_count":        0,
		"unread_mark":         0,
	}, userId, peerId)
}

func (m *PrivateCore) MarkDialogUnread(ctx context.Context, userId int32, peerId int32, unreadMark bool) error {
	_, err := m.ConversationsDAO.UpdateMarkDialogUnread(ctx, util.BoolToInt8(unreadMark), userId, peerId)
	return err
}

func (m *PrivateCore) ToggleDialogPin(ctx context.Context, userId, peerId int32, pinned bool) (folderId int32, err error) {
	if dialogs, err2 := m.ConversationsDAO.SelectListByPeerList(ctx, userId, []int32{peerId}); err2 != nil || len(dialogs) == 0 {
		err = mtproto.ErrPeerIdInvalid
		return
	} else {
		folderId = dialogs[0].FolderId
	}

	var (
		orderPinned = time.Now().Unix() << 32
		cMap        = make(map[string]interface{})
	)

	if folderId == 0 {
		if pinned {
			cMap["is_pinned"] = 1
			cMap["order_pinned"] = orderPinned
		} else {
			cMap["is_pinned"] = 0
			cMap["order_pinned"] = 0
		}
	} else {
		if pinned {
			cMap["folder_pinned"] = 1
			cMap["folder_order_pinned"] = orderPinned
		} else {
			cMap["folder_pinned"] = 0
			cMap["folder_order_pinned"] = 0
		}
	}

	_, err = m.ConversationsDAO.UpdateCustomMap(ctx, cMap, userId, peerId)
	return
}

func (m *PrivateCore) GetDialogUnreadMarkList(ctx context.Context, userId int32) []*mtproto.DialogPeer {
	var (
		peerList []*mtproto.DialogPeer
	)

	idList, _ := m.ConversationsDAO.SelectMarkDialogUnreadList(ctx, userId)
	for _, id := range idList {
		peerList = append(peerList, mtproto.MakeTLDialogPeer(&mtproto.DialogPeer{
			Peer: model.MakePeerUser(id),
		}).To_DialogPeer())
	}

	return peerList
}

func (m *PrivateCore) GetDialogsCount(ctx context.Context, userId int32, excludePinned bool, folderId int32) int {
	if excludePinned {
		where := fmt.Sprintf("user_id = %d AND is_pinned = 0 AND folderId = %d AND top_message > 0", userId, folderId)
		c1 := m.CommonDAO.CalcSizeByWhere(ctx, "conversations", where)
		return c1
	} else {
		where := fmt.Sprintf("user_id = %d AND folderId = %d AND top_message > 0 AND deleted = 0", userId, folderId)
		c1 := m.CommonDAO.CalcSizeByWhere(ctx, "conversations", where)
		return c1
	}
}

func (m *PrivateCore) GetDialogsByOffsetId(ctx context.Context, userId int32, excludePinned bool, offsetId int32, limit int32) (dialogs model.DialogExtList) {
	var doList []dataobject.ConversationsDO
	if excludePinned {
		doList, _ = m.ConversationsDAO.SelectExcludePinnedByOffsetId(ctx, userId, offsetId, limit)
	} else {
		doList, _ = m.ConversationsDAO.SelectByOffsetId(ctx, userId, offsetId, limit)
	}
	dialogs = makeDialogList(dialogs, doList)

	return
}

func (m *PrivateCore) GetDialogsByIdList(ctx context.Context, userId int32, idList []int32) (dialogs model.DialogExtList) {
	if len(idList) == 0 {
		return
	}
	doList, _ := m.ConversationsDAO.SelectListByPeerList(ctx, userId, idList)
	for _, id := range idList {
		found := false
		for i := 0; i < len(doList); i++ {
			if doList[i].PeerId == id {
				found = true
				break
			}
		}
		if !found {
			doList = append(doList, dataobject.ConversationsDO{
				UserId:     userId,
				PeerId:     id,
				TopMessage: 0,
			})
		}
	}
	dialogs = makeDialogList(dialogs, doList)

	return
}

func (m *PrivateCore) GetPinnedDialogs(ctx context.Context, userId, folderId int32) (dialogs model.DialogExtList) {
	doList, _ := m.ConversationsDAO.SelectPinnedDialogs(ctx, userId, folderId)
	dialogs = makeDialogList(dialogs, doList)
	return
}

func (m *PrivateCore) GetDialogs(ctx context.Context, userId int32, excludePinned bool, folderId int32) (dialogs model.DialogExtList) {
	var doList []dataobject.ConversationsDO
	if excludePinned {
		if folderId == 0 {
			doList, _ = m.ConversationsDAO.SelectExcludePinnedDialogs(ctx, userId, "is_pinned", folderId)
			log.Debugf("getDialogs - params(%d, %v, %d), %v", userId, excludePinned, folderId, doList)
		} else {
			doList, _ = m.ConversationsDAO.SelectExcludePinnedDialogs(ctx, userId, "folder_pinned", folderId)
			log.Debugf("getDialogs - %v", doList)
		}
	} else {
		doList, _ = m.ConversationsDAO.SelectDialogs(ctx, userId, folderId)
		log.Debugf("getDialogs - %v", doList)
	}
	return makeDialogList(dialogs, doList)
}

func (m *PrivateCore) GetDialogsByOffsetDate(ctx context.Context, userId int32, excludePinned bool, offsetDate, limit int32) model.DialogExtList {
	log.Warnf("not impl GetDialogsByOffsetDate")
	return nil
}

func (m *PrivateCore) ReorderPinnedDialogs(ctx context.Context, userId int32, force bool, folderId int32, idList []int32) (err error) {
	orderPinned := time.Now().Unix()

	if folderId == 0 {
		if force {
			_, err = m.ConversationsDAO.UpdateUnPinnedList(ctx, userId, idList)
		}

		for _, id := range idList {
			_, err = m.ConversationsDAO.UpdateCustomMap(ctx, map[string]interface{}{
				"is_pinned":    1,
				"order_pinned": orderPinned << 32,
			}, userId, id)
			orderPinned -= 1
		}
	} else {
		if force {
			_, err = m.ConversationsDAO.UpdateFolderUnPinnedList(ctx, userId, idList)
		}

		for _, id := range idList {
			_, err = m.ConversationsDAO.UpdateCustomMap(ctx, map[string]interface{}{
				"folder_pinned":       1,
				"folder_order_pinned": orderPinned << 32,
			}, userId, id)
			orderPinned -= 1
		}
	}

	return
}

func (m *PrivateCore) GetDialogFolder(ctx context.Context, userId, folderId int32) (*model.DialogExt, error) {
	dialogDOList, _ := m.ConversationsDAO.SelectDialogs(ctx, userId, folderId)
	if len(dialogDOList) == 0 {
		return nil, nil
	}

	dialogExt := &model.DialogExt{
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

	for i := 0; i < len(dialogDOList); i++ {
		order := dialogDOList[i].FolderOrderPinned
		if order == 0 {
			order = int64(dialogDOList[i].TopMessage)
		}
		if order > dialogExt.Order {
			dialogExt.Order = order
			dialogExt.Dialog.Peer = model.MakePeerUser(dialogDOList[i].PeerId)
			dialogExt.Dialog.TopMessage = dialogDOList[i].TopMessage
		}
		if dialogDOList[i].UnreadCount > 0 {
			dialogExt.Dialog.UnreadMutedPeersCount += 1
			dialogExt.Dialog.UnreadMutedMessagesCount += dialogDOList[i].UnreadCount
		} else if dialogDOList[i].UnreadMark == 1 {
			dialogExt.Dialog.UnreadMutedPeersCount += 1
			dialogExt.Dialog.UnreadMutedMessagesCount += 1
		}
	}
	return dialogExt, nil
}

func (m *PrivateCore) GetPinnedDialogPeers(ctx context.Context, userId, folderId int32) (dialogs model.DialogPinnedExtList) {
	doList, _ := m.ConversationsDAO.SelectPinnedDialogs(ctx, userId, folderId)

	var order int64 = 0
	dialogs = make([]model.DialogPinnedExt, 0, len(doList))
	for i := 0; i < len(doList); i++ {
		if folderId == 0 {
			order = doList[i].OrderPinned
		} else {
			order = doList[i].FolderOrderPinned
		}
		dialogs = append(dialogs, model.DialogPinnedExt{
			Order:    order,
			PeerUtil: model.PeerUtil{PeerType: model.PEER_USER, PeerId: doList[i].PeerId},
		})
	}
	return
}

func (m *PrivateCore) EditPeerFoldersByIdList(ctx context.Context, userId, folderId int32, idList []int32) (map[int32]int64, error) {
	needPinnedDialogs := make(map[int32]int64)
	pinnedDOList, _ := m.ConversationsDAO.SelectListByPeerList(ctx, userId, idList)
	for i := 0; i < len(pinnedDOList); i++ {
		if folderId == 0 {
			if pinnedDOList[i].IsPinned == 1 {
				needPinnedDialogs[pinnedDOList[i].PeerId] = pinnedDOList[i].OrderPinned
			}
		} else {
			if pinnedDOList[i].FolderPinned == 1 {
				needPinnedDialogs[pinnedDOList[i].PeerId] = pinnedDOList[i].FolderOrderPinned
			}
		}
	}

	m.ConversationsDAO.UpdateFolderId(ctx, folderId, userId, idList)
	return needPinnedDialogs, nil
}

func (m *PrivateCore) UpdatePinnedMessage(ctx context.Context, userId int32, peerId, pinnedMsgId int32) error {
	_, err := m.ConversationsDAO.UpdateCustomMap(ctx, map[string]interface{}{
		"pinned_msg_id": pinnedMsgId,
	}, userId, peerId)
	return err
}

func (m *PrivateCore) GetPinnedMessage(ctx context.Context, userId int32, peerId int32) (pinnedMsgId int32) {
	dDO, _ := m.ConversationsDAO.SelectByPeer(ctx, userId, peerId)
	if dDO != nil {
		pinnedMsgId = dDO.PinnedMsgId
	}
	return
}
