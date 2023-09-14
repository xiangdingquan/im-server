package core

import (
	"context"
	"encoding/json"
	"time"

	"open.chat/pkg/database/sqlx"

	"github.com/gogo/protobuf/types"

	"fmt"

	"open.chat/app/service/biz_service/channel/dal/dataobject"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (m *ChannelCore) makeDialog(ctx context.Context, userId int32, dialogDO *dataobject.ChannelParticipantsExtDO) *mtproto.Dialog {
	dialog := mtproto.MakeTLDialog(&mtproto.Dialog{
		Pinned:              util.Int8ToBool(dialogDO.IsPinned),
		UnreadMark:          util.Int8ToBool(dialogDO.UnreadMark),
		Peer:                model.MakePeerChannel(dialogDO.ChannelId),
		TopMessage:          dialogDO.TopMessage,
		ReadInboxMaxId:      dialogDO.ReadInboxMaxId,
		ReadOutboxMaxId:     dialogDO.ReadOutboxMaxId,
		UnreadCount:         m.GetUnreadCount(ctx, dialogDO.ChannelId, userId),
		UnreadMentionsCount: m.GetChannelUnreadMentionsCount(ctx, userId, dialogDO.ChannelId, dialogDO.ReadInboxMaxId),
		NotifySettings:      nil,
		Pts:                 &types.Int32Value{Value: dialogDO.Pts},
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

func (m *ChannelCore) makeUnknownDialog(channelId, topMessage, pts int32) *mtproto.Dialog {
	return mtproto.MakeTLDialog(&mtproto.Dialog{
		Pinned:              false,
		UnreadMark:          false,
		Peer:                model.MakePeerChannel(channelId),
		TopMessage:          topMessage,
		ReadInboxMaxId:      0,
		ReadOutboxMaxId:     0,
		UnreadCount:         0,
		UnreadMentionsCount: 0,
		NotifySettings:      model.MakeDefaultPeerNotifySettings(model.PEER_CHANNEL),
		Pts:                 &types.Int32Value{Value: pts},
		Draft:               nil,
		FolderId:            nil,
	}).To_Dialog()
}

func (m *ChannelCore) makeDialogList(ctx context.Context, userId int32, i model.DialogExtList, dialogDOList []dataobject.ChannelParticipantsExtDO) (dialogs model.DialogExtList) {
	if len(i) == 0 {
		dialogs = []*model.DialogExt{}
	} else {
		dialogs = i
	}

	for i := 0; i < len(dialogDOList); i++ {
		dialogs = append(dialogs, &model.DialogExt{
			Order:          int64(dialogDOList[i].Date2),
			Dialog:         m.makeDialog(ctx, userId, &dialogDOList[i]),
			AvailableMinId: dialogDOList[i].AvailableMinId,
			Date:           dialogDOList[i].Date,
		})
	}
	return
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (m *ChannelCore) GetDialogUnreadMarkList(ctx context.Context, userId int32) []*mtproto.DialogPeer {
	var (
		peerList []*mtproto.DialogPeer
	)

	idList, _ := m.ChannelParticipantsDAO.SelectMarkDialogUnreadList(ctx, userId)
	for _, id := range idList {
		peerList = append(peerList, mtproto.MakeTLDialogPeer(&mtproto.DialogPeer{
			Peer: model.MakePeerChannel(id),
		}).To_DialogPeer())
	}

	return peerList
}

func (m *ChannelCore) GetDialogsCount(ctx context.Context, userId int32, excludePinned bool, folderId int32) int {
	if excludePinned {
		where := fmt.Sprintf("user_id = %d AND is_pinned = 0 AND folderId = %d AND (state=0 OR state=2)", userId, folderId)
		c1 := m.CommonDAO.CalcSizeByWhere(ctx, "channel_participants", where)
		return c1
	} else {
		where := fmt.Sprintf("user_id = %d AND folderId = %d AND (state=0 OR state=2)", userId, folderId)
		c1 := m.CommonDAO.CalcSizeByWhere(ctx, "channel_participants", where)
		return c1
	}
}

func (m *ChannelCore) GetDialogs(ctx context.Context, userId int32, excludePinned bool, folderId int32) (dialogs model.DialogExtList) {
	var doList []dataobject.ChannelParticipantsExtDO

	if excludePinned {
		if folderId == 0 {
			doList, _ = m.ChannelParticipantsDAO.SelectExcludePinnedDialogs(ctx, userId, "is_pinned", folderId)
		} else {
			doList, _ = m.ChannelParticipantsDAO.SelectExcludePinnedDialogs(ctx, userId, "folder_pinned", folderId)
		}
	} else {
		doList, _ = m.ChannelParticipantsDAO.SelectDialogs(ctx, userId, folderId)
	}

	return m.makeDialogList(ctx, userId, dialogs, doList)
}

func (m *ChannelCore) GetDialogsByIdList(ctx context.Context, userId int32, idList []int32) (dialogs model.DialogExtList) {
	if len(idList) == 0 {
		return
	}
	doList, _ := m.ChannelParticipantsDAO.SelectExtListByChannelIdList(ctx, userId, idList)
	for _, id := range idList {
		found := -1
		for i := 0; i < len(doList); i++ {
			if doList[i].ChannelId == id {
				found = i
				break
			}
		}

		if found == -1 {
			// m.GetTopMessage()
			if foundDO, _ := m.ChannelsDAO.Select(ctx, id); foundDO != nil {
				dialogs = append(dialogs, &model.DialogExt{
					Order:          int64(foundDO.Date2),
					Dialog:         m.makeUnknownDialog(id, foundDO.TopMessage, foundDO.Pts),
					AvailableMinId: 0,
					Date:           foundDO.Date,
				})
			}
		} else {
			// doList[found].UnreadMentionsCount = m.chann

			dialogs = append(dialogs, &model.DialogExt{
				Order:          int64(doList[found].Date2),
				Dialog:         m.makeDialog(ctx, userId, &doList[found]),
				AvailableMinId: doList[found].AvailableMinId,
				Date:           doList[found].Date,
			})
		}
	}

	return
}

func (m *ChannelCore) ToggleDialogPin(ctx context.Context, userId, peerId int32, pinned bool) (folderId int32, err error) {
	if dialogs, err2 := m.ChannelParticipantsDAO.SelectListByChannelIdList(ctx, userId, []int32{peerId}); err2 != nil || len(dialogs) == 0 {
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

	_, err = m.ChannelParticipantsDAO.Update(ctx, cMap, userId, peerId)
	return
}

func (m *ChannelCore) GetPinnedDialogs(ctx context.Context, userId, folderId int32) (dialogs model.DialogExtList) {
	doList, _ := m.ChannelParticipantsDAO.SelectPinnedDialogs(ctx, userId, folderId)
	dialogs = m.makeDialogList(ctx, userId, dialogs, doList)
	return
}

func (m *ChannelCore) ReorderPinnedDialogs(ctx context.Context, userId int32, force bool, folderId int32, idList []int32) (err error) {
	orderPinned := time.Now().Unix()

	if folderId == 0 {
		if force {
			_, err = m.ChannelParticipantsDAO.UpdateUnPinnedList(ctx, userId, idList)
		}

		for _, id := range idList {
			_, err = m.ChannelParticipantsDAO.Update(ctx, map[string]interface{}{
				"is_pinned":    1,
				"order_pinned": orderPinned << 32,
			}, userId, id)
			orderPinned -= 1
		}
	} else {
		if force {
			_, err = m.ChannelParticipantsDAO.UpdateFolderUnPinnedList(ctx, userId, idList)
		}

		for _, id := range idList {
			_, err = m.ChannelParticipantsDAO.Update(ctx, map[string]interface{}{
				"folder_pinned":       1,
				"folder_order_pinned": orderPinned << 32,
			}, userId, id)
			orderPinned -= 1
		}
	}

	return
}

func (m *ChannelCore) GetDialogFolder(ctx context.Context, userId, folderId int32) (*model.DialogExt, error) {
	dialogDOList, _ := m.ChannelParticipantsDAO.SelectDialogs(ctx, userId, folderId)
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
			dialogExt.Dialog.Peer = model.MakePeerChannel(dialogDOList[i].ChannelId)
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

func (m *ChannelCore) GetPinnedDialogPeers(ctx context.Context, userId, folderId int32) (dialogs model.DialogPinnedExtList) {
	doList, _ := m.ChannelParticipantsDAO.SelectPinnedDialogs(ctx, userId, folderId)

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
			PeerUtil: *model.MakeChannelPeerUtil(doList[i].ChannelId),
		})
	}
	return
}

func (m *ChannelCore) MarkDialogUnread(ctx context.Context, userId int32, peerId int32, unreadMark bool) error {
	_, err := m.ChannelParticipantsDAO.UpdateMarkDialogUnread(ctx, util.BoolToInt8(unreadMark), userId, peerId)
	return err
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (m *ChannelCore) GetDialogsByOffsetDate(ctx context.Context, userId int32, excludePinned bool, offsetDate, limit int32) model.DialogExtList {
	log.Warnf("not impl GetDialogsByOffsetDate")
	return []*model.DialogExt{}
}

func (m *ChannelCore) GetDialogsByOffsetId(ctx context.Context, userId int32, excludePinned bool, offsetDate, limit int32) model.DialogExtList {
	log.Warnf("not impl GetDialogsByOffsetId")
	return []*model.DialogExt{}
}

func (m *ChannelCore) EditPeerFoldersByIdList(ctx context.Context, userId, folderId int32, idList []int32) (map[int32]int64, error) {

	needPinnedDialogs := make(map[int32]int64)
	if len(idList) == 0 {
		return needPinnedDialogs, nil
	}

	pinnedDOList, _ := m.ChannelParticipantsDAO.SelectListByChannelIdList(ctx, userId, idList)
	for i := 0; i < len(pinnedDOList); i++ {
		if folderId == 0 {
			if pinnedDOList[i].IsPinned == 1 {
				needPinnedDialogs[pinnedDOList[i].ChannelId] = pinnedDOList[i].OrderPinned
			}
		} else {
			if pinnedDOList[i].FolderPinned == 1 {
				needPinnedDialogs[pinnedDOList[i].ChannelId] = pinnedDOList[i].FolderOrderPinned
			}
		}
	}

	m.ChannelParticipantsDAO.UpdateFolderId(ctx, folderId, userId, idList)
	return needPinnedDialogs, nil
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (m *ChannelCore) GetTopMessage(ctx context.Context, userId int32, peerId int32) int32 {
	var (
		topMessage int32
	)

	do, _ := m.ChannelParticipantsDAO.SelectExtByChannelIdUserId(ctx, peerId, userId)
	if do != nil {
		topMessage = do.TopMessage
	}

	return topMessage
}

func (m *ChannelCore) GetReadInboxMaxId(ctx context.Context, userId int32, peerId int32) int32 {
	var (
		readInboxMaxId int32
	)

	do, _ := m.ChannelParticipantsDAO.SelectExtByChannelIdUserId(ctx, peerId, userId)
	if do != nil {
		readInboxMaxId = do.ReadInboxMaxId
	}

	return readInboxMaxId
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (m *ChannelCore) UpdateReadInbox(ctx context.Context, userId, peerId int32, readInboxId int32) {
	log.Warnf("not impl UpdateReadInbox")
}

func (m *ChannelCore) UpdateReadOutbox(ctx context.Context, userId, peerId int32, readOutboxId int32) {
	m.ChannelParticipantsDAO.UpdateReadInboxMaxId(ctx, readOutboxId, userId, peerId)
}

func (m *ChannelCore) ReadOutboxHistory(ctx context.Context, channelId, userId, maxId int32) bool {
	sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		m.ChannelParticipantsDAO.UpdateReadInboxMaxIdTx(tx, maxId, userId, channelId)
		m.ChannelsDAO.UpdateReadOutboxMaxIdTx(tx, maxId, channelId)
	})
	return true
}

func (m *ChannelCore) GetUnreadCount(ctx context.Context, channelId, userId int32) int32 {
	return m.ChannelParticipantsDAO.SelectUnreadCountByChannelIdUserId(ctx, channelId, userId)
}

func (m *ChannelCore) GetReadOutboxMaxId(ctx context.Context, channelId int32) int32 {
	rValue, _ := m.ChannelsDAO.SelectReadOutboxMaxId(ctx, channelId)
	return rValue
}
