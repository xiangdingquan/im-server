package core

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gogo/protobuf/types"
	"open.chat/app/service/biz_service/chat/dal/dataobject"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

// ////////////////////////////////////////////////////////////////////////////////////
func makeDialog(dialogDO *dataobject.ChatParticipantsDO) *mtproto.Dialog {
	dialog := mtproto.MakeTLDialog(&mtproto.Dialog{
		Pinned:              util.Int8ToBool(dialogDO.IsPinned),
		UnreadMark:          util.Int8ToBool(dialogDO.UnreadMark),
		Peer:                model.MakePeerChat(dialogDO.ChatId),
		TopMessage:          dialogDO.TopMessage,
		ReadInboxMaxId:      dialogDO.ReadInboxMaxId,
		ReadOutboxMaxId:     dialogDO.ReadOutboxMaxId,
		UnreadCount:         dialogDO.UnreadCount,
		UnreadMentionsCount: dialogDO.UnreadMentionsCount,
		NotifySettings:      nil,
		Pts:                 nil,
		Draft:               nil,
		FolderId:            nil,
	}).To_Dialog()

	// draft message.
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

	// NotifySettings
	dialog.NotifySettings = mtproto.MakeTLPeerNotifySettings(&mtproto.PeerNotifySettings{}).To_PeerNotifySettings()

	// folder_id
	if dialogDO.FolderId != 0 {
		dialog.FolderId = &types.Int32Value{Value: dialogDO.FolderId}
	}

	return dialog
}

func makeDialogList(i model.DialogExtList, dialogDOList []dataobject.ChatParticipantsDO) (dialogs model.DialogExtList) {
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
func (m *ChatCore) UpdateReadInbox(ctx context.Context, userId, peerId int32, readInboxId int32) {
	m.ChatParticipantsDAO.UpdateUnreadByPeer(ctx, readInboxId, userId, peerId)
}

func (m *ChatCore) UpdateReadOutbox(ctx context.Context, userId, peerId int32, readOutboxId int32) {
	m.ChatParticipantsDAO.UpdateReadOutboxMaxIdByPeer(ctx, readOutboxId, userId, peerId)
}

func (m *ChatCore) GetTopMessage(ctx context.Context, userId int32, peerId int32) int32 {
	var (
		topMessage int32
	)

	do, _ := m.ChatParticipantsDAO.SelectByParticipantId(ctx, peerId, userId)

	if do != nil {
		topMessage = do.TopMessage
	}

	return topMessage
}

func (m *ChatCore) GetReadInboxMaxId(ctx context.Context, userId, peerId int32) int32 {
	var (
		readInboxMaxId int32
	)

	do, _ := m.ChatParticipantsDAO.SelectByParticipantId(ctx, peerId, userId)

	if do != nil {
		readInboxMaxId = do.ReadInboxMaxId
	}

	return readInboxMaxId
}

func (m *ChatCore) ResetTopMessage(ctx context.Context, userId int32, peerId int32) {
	m.ChatParticipantsDAO.UpdateCustomMap(ctx, map[string]interface{}{
		"top_message":         0,
		"is_pinned":           0,
		"order_pinned":        0,
		"folder_pinned":       0,
		"folder_order_pinned": 0,
		"unread_count":        0,
		"unread_mark":         0,
	}, userId, peerId)
}

func (m *ChatCore) MarkDialogUnread(ctx context.Context, userId int32, peerId int32, unreadMark bool) error {
	_, err := m.ChatParticipantsDAO.UpdateMarkDialogUnread(ctx, util.BoolToInt8(unreadMark), userId, peerId)
	return err
}

func (m *ChatCore) ToggleDialogPin(ctx context.Context, userId, peerId int32, pinned bool) (folderId int32, err error) {
	if dialogs, err2 := m.ChatParticipantsDAO.SelectListByChatIdList(ctx, userId, []int32{peerId}); err2 != nil || len(dialogs) == 0 {
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
			// pinned
			cMap["is_pinned"] = 1
			cMap["order_pinned"] = orderPinned
		} else {
			// cancel pinned
			cMap["is_pinned"] = 0
			cMap["order_pinned"] = 0
		}
	} else {
		if pinned {
			// pinned
			cMap["folder_pinned"] = 1
			cMap["folder_order_pinned"] = orderPinned
		} else {
			// cancel pinned
			cMap["folder_pinned"] = 0
			cMap["folder_order_pinned"] = 0
		}
	}

	_, err = m.ChatParticipantsDAO.UpdateCustomMap(ctx, cMap, userId, peerId)
	return
}

func (m *ChatCore) GetDialogUnreadMarkList(ctx context.Context, userId int32) []*mtproto.DialogPeer {
	var (
		peerList []*mtproto.DialogPeer
	)

	idList, _ := m.ChatParticipantsDAO.SelectMarkDialogUnreadList(ctx, userId)
	for _, id := range idList {
		peerList = append(peerList, mtproto.MakeTLDialogPeer(&mtproto.DialogPeer{
			Peer: model.MakePeerChat(id),
		}).To_DialogPeer())
	}

	return peerList
}

func (m *ChatCore) GetDialogsCount(ctx context.Context, userId int32, excludePinned bool, folderId int32) int {
	if excludePinned {
		where := fmt.Sprintf("user_id = %d AND is_pinned = 0 AND folderId = %d AND top_message > 0 AND (state=0 OR state=2)", userId, folderId)
		c1 := m.CommonDAO.CalcSizeByWhere(ctx, "chat_participants", where)
		return c1
	} else {
		where := fmt.Sprintf("user_id = %d AND folderId = %d AND top_message > 0 AND (state=0 OR state=2)", userId, folderId)
		c1 := m.CommonDAO.CalcSizeByWhere(ctx, "chat_participants", where)
		return c1
	}
}

func (m *ChatCore) GetDialogsByOffsetId(ctx context.Context, userId int32, excludePinned bool, offsetId int32, limit int32) (dialogs model.DialogExtList) {
	var doList []dataobject.ChatParticipantsDO
	if excludePinned {
		doList, _ = m.ChatParticipantsDAO.SelectExcludePinnedByOffsetId(ctx, userId, offsetId, limit)
	} else {
		doList, _ = m.ChatParticipantsDAO.SelectByOffsetId(ctx, userId, offsetId, limit)
	}

	for i := 0; i < len(doList); i++ {
		doList[i].UnreadMentionsCount = int32(m.CommonDAO.CalcSize(ctx, "messages", map[string]interface{}{
			"user_id":   userId,
			"peer_type": model.PEER_CHAT,
			"peer_id":   doList[i].ChatId,
			"mentioned": 1,
			"deleted":   0,
		}))
	}
	dialogs = makeDialogList(dialogs, doList)

	return
}

func (m *ChatCore) GetDialogsByIdList(ctx context.Context, userId int32, idList []int32) (dialogs model.DialogExtList) {
	if len(idList) == 0 {
		return
	}
	doList, _ := m.ChatParticipantsDAO.SelectListByChatIdList(ctx, userId, idList)
	for i := 0; i < len(doList); i++ {
		doList[i].UnreadMentionsCount = int32(m.CommonDAO.CalcSize(ctx, "messages", map[string]interface{}{
			"user_id":   userId,
			"peer_type": model.PEER_CHAT,
			"peer_id":   doList[i].ChatId,
			"mentioned": 1,
			"deleted":   0,
		}))
	}

	for _, id := range idList {
		found := false
		for i := 0; i < len(doList); i++ {
			if doList[i].ChatId == id {
				found = true
				break
			}
		}
		if !found {
			doList = append(doList, dataobject.ChatParticipantsDO{
				UserId:     userId,
				ChatId:     id,
				TopMessage: 0,
			})
		}
	}

	for i := 0; i < len(doList); i++ {
		doList[i].UnreadMentionsCount = int32(m.CommonDAO.CalcSize(ctx, "messages", map[string]interface{}{
			"user_id":   userId,
			"peer_type": model.PEER_CHAT,
			"peer_id":   doList[i].ChatId,
			"mentioned": 1,
			"deleted":   0,
		}))
	}
	dialogs = makeDialogList(dialogs, doList)

	return
}

func (m *ChatCore) GetPinnedDialogs(ctx context.Context, userId, folderId int32) (dialogs model.DialogExtList) {
	doList, _ := m.ChatParticipantsDAO.SelectPinnedDialogs(ctx, userId, folderId)

	for i := 0; i < len(doList); i++ {
		doList[i].UnreadMentionsCount = int32(m.CommonDAO.CalcSize(ctx, "messages", map[string]interface{}{
			"user_id":   userId,
			"peer_type": model.PEER_CHAT,
			"peer_id":   doList[i].ChatId,
			"mentioned": 1,
			"deleted":   0,
		}))
	}
	dialogs = makeDialogList(dialogs, doList)
	return
}

func (m *ChatCore) GetDialogs(ctx context.Context, userId int32, excludePinned bool, folderId int32) (dialogs model.DialogExtList) {
	var doList []dataobject.ChatParticipantsDO
	if excludePinned {
		if folderId == 0 {
			doList, _ = m.ChatParticipantsDAO.SelectExcludePinnedDialogs(ctx, userId, "is_pinned", folderId)
		} else {
			doList, _ = m.ChatParticipantsDAO.SelectExcludePinnedDialogs(ctx, userId, "folder_pinned", folderId)
		}
	} else {
		doList, _ = m.ChatParticipantsDAO.SelectDialogs(ctx, userId, folderId)
	}
	for i := 0; i < len(doList); i++ {
		doList[i].UnreadMentionsCount = int32(m.CommonDAO.CalcSize(ctx, "messages", map[string]interface{}{
			"user_id":   userId,
			"peer_type": model.PEER_CHAT,
			"peer_id":   doList[i].ChatId,
			"mentioned": 1,
			"deleted":   0,
		}))
	}
	return makeDialogList(dialogs, doList)
}

func (m *ChatCore) GetDialogsByOffsetDate(ctx context.Context, userId int32, excludePinned bool, offsetDate, limit int32) model.DialogExtList {
	log.Warnf("not impl GetDialogsByOffsetDate")
	return nil
}

func (m *ChatCore) ReorderPinnedDialogs(ctx context.Context, userId int32, force bool, folderId int32, idList []int32) (err error) {
	orderPinned := time.Now().Unix()

	if folderId == 0 {
		if force {
			_, err = m.ChatParticipantsDAO.UpdateUnPinnedList(ctx, userId, idList)
		}

		for _, id := range idList {
			_, err = m.ChatParticipantsDAO.UpdateCustomMap(ctx, map[string]interface{}{
				"is_pinned":    1,
				"order_pinned": orderPinned << 32,
			}, userId, id)
			orderPinned -= 1
		}
	} else {
		if force {
			_, err = m.ChatParticipantsDAO.UpdateFolderUnPinnedList(ctx, userId, idList)
		}

		for _, id := range idList {
			_, err = m.ChatParticipantsDAO.UpdateCustomMap(ctx, map[string]interface{}{
				"folder_pinned":       1,
				"folder_order_pinned": orderPinned << 32,
			}, userId, id)
			orderPinned -= 1
		}
	}

	return
}

func (m *ChatCore) GetDialogFolder(ctx context.Context, userId, folderId int32) (*model.DialogExt, error) {
	dialogDOList, _ := m.ChatParticipantsDAO.SelectDialogs(ctx, userId, folderId)
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
			dialogExt.Dialog.Peer = model.MakePeerChat(dialogDOList[i].ChatId)
			dialogExt.Dialog.TopMessage = dialogDOList[i].TopMessage
		}
		if dialogDOList[i].UnreadCount > 0 {
			dialogExt.Dialog.UnreadMutedPeersCount += 1
			dialogExt.Dialog.UnreadMutedMessagesCount += dialogDOList[i].UnreadCount
		} else if dialogDOList[i].UnreadMark == 1 {
			// if unread_mark then 1
			dialogExt.Dialog.UnreadMutedPeersCount += 1
			dialogExt.Dialog.UnreadMutedMessagesCount += 1
		}
	}
	return dialogExt, nil
}

func (m *ChatCore) GetPinnedDialogPeers(ctx context.Context, userId, folderId int32) (dialogs model.DialogPinnedExtList) {
	doList, _ := m.ChatParticipantsDAO.SelectPinnedDialogs(ctx, userId, folderId)

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
			PeerUtil: model.PeerUtil{PeerType: model.PEER_USER, PeerId: doList[i].ChatId},
		})
	}
	return
}

func (m *ChatCore) EditPeerFoldersByIdList(ctx context.Context, userId, folderId int32, idList []int32) (map[int32]int64, error) {
	needPinnedDialogs := make(map[int32]int64)
	pinnedDOList, _ := m.ChatParticipantsDAO.SelectListByChatIdList(ctx, userId, idList)
	for i := 0; i < len(pinnedDOList); i++ {
		if folderId == 0 {
			if pinnedDOList[i].IsPinned == 1 {
				needPinnedDialogs[pinnedDOList[i].ChatId] = pinnedDOList[i].OrderPinned
			}
		} else {
			if pinnedDOList[i].FolderPinned == 1 {
				needPinnedDialogs[pinnedDOList[i].ChatId] = pinnedDOList[i].FolderOrderPinned
			}
		}
	}

	m.ChatParticipantsDAO.UpdateFolderId(ctx, folderId, userId, idList)
	return needPinnedDialogs, nil
}
