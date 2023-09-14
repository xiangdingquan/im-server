package private_facade

import (
	"context"

	"open.chat/app/service/biz_service/private/internal/core"
	"open.chat/app/service/biz_service/private/internal/dao"
	"open.chat/model"
	"open.chat/mtproto"
)

type localPrivateFacade struct {
	*core.PrivateCore
}

func New() PrivateFacade {
	return &localPrivateFacade{
		PrivateCore: core.New(dao.New()),
	}
}

// /////////////////////////////////////////////////////////////////////////////////////////////
// draft
func (c *localPrivateFacade) SaveDraftMessage(ctx context.Context, userId, peerId int32, draft *mtproto.DraftMessage) error {
	return c.PrivateCore.SaveDraftMessage(ctx, userId, peerId, draft)
}

func (c *localPrivateFacade) ClearDraftMessage(ctx context.Context, userId, peerId int32) error {
	return c.PrivateCore.ClearDraftMessage(ctx, userId, peerId)
}

func (c *localPrivateFacade) GetAllDrafts(ctx context.Context, userId int32) ([]int32, []*mtproto.DraftMessage, error) {
	return c.PrivateCore.GetAllDrafts(ctx, userId)
}

func (c *localPrivateFacade) ClearAllDrafts(ctx context.Context, userId int32) error {
	return c.PrivateCore.ClearAllDrafts(ctx, userId)
}

// /////////////////////////////////////////////////////////////////////////////////////////////
// dialog
func (c *localPrivateFacade) MarkDialogUnread(ctx context.Context, userId int32, peerId int32, unreadMark bool) error {
	return c.PrivateCore.MarkDialogUnread(ctx, userId, peerId, unreadMark)
}

func (c *localPrivateFacade) ToggleDialogPin(ctx context.Context, userId, peerId int32, pinned bool) (int32, error) {
	return c.PrivateCore.ToggleDialogPin(ctx, userId, peerId, pinned)
}

func (c *localPrivateFacade) GetDialogUnreadMarkList(ctx context.Context, userId int32) []*mtproto.DialogPeer {
	return c.PrivateCore.GetDialogUnreadMarkList(ctx, userId)
}

func (c *localPrivateFacade) GetDialogsCount(ctx context.Context, userId int32, excludePinned bool, folderId int32) int {
	return c.PrivateCore.GetDialogsCount(ctx, userId, excludePinned, folderId)
}

func (c *localPrivateFacade) GetDialogsByOffsetId(ctx context.Context, userId int32, excludePinned bool, offsetId int32, limit int32) (dialogs model.DialogExtList) {
	return c.PrivateCore.GetDialogsByOffsetId(ctx, userId, excludePinned, offsetId, limit)
}

func (c *localPrivateFacade) GetDialogsByIdList(ctx context.Context, userId int32, idList []int32) model.DialogExtList {
	return c.PrivateCore.GetDialogsByIdList(ctx, userId, idList)
}

func (c *localPrivateFacade) GetPinnedDialogs(ctx context.Context, userId, folderId int32) (dialogs model.DialogExtList) {
	return c.PrivateCore.GetPinnedDialogs(ctx, userId, folderId)
}

func (c *localPrivateFacade) GetDialogs(ctx context.Context, userId int32, excludePinned bool, folderId int32) model.DialogExtList {
	return c.PrivateCore.GetDialogs(ctx, userId, excludePinned, folderId)
}

func (c *localPrivateFacade) GetDialogsByOffsetDate(ctx context.Context, userId int32, excludePinned bool, offsetDate, limit int32) model.DialogExtList {
	return c.PrivateCore.GetDialogsByOffsetDate(ctx, userId, excludePinned, offsetDate, limit)
}

func (c *localPrivateFacade) ReorderPinnedDialogs(ctx context.Context, userId int32, force bool, folderId int32, idList []int32) error {
	return c.PrivateCore.ReorderPinnedDialogs(ctx, userId, force, folderId, idList)
}

func (c *localPrivateFacade) GetDialogFolder(ctx context.Context, userId, folderId int32) (*model.DialogExt, error) {
	return c.PrivateCore.GetDialogFolder(ctx, userId, folderId)
}

func (c *localPrivateFacade) GetPinnedDialogPeers(ctx context.Context, userId, folderId int32) model.DialogPinnedExtList {
	return c.PrivateCore.GetPinnedDialogPeers(ctx, userId, folderId)
}

func (c *localPrivateFacade) EditPeerFoldersByIdList(ctx context.Context, userId, folderId int32, idList []int32) (map[int32]int64, error) {
	return c.PrivateCore.EditPeerFoldersByIdList(ctx, userId, folderId, idList)
}

func (c *localPrivateFacade) GetTopMessage(ctx context.Context, userId, peerId int32) int32 {
	return c.PrivateCore.GetTopMessage(ctx, userId, peerId)
}

func (c *localPrivateFacade) UpdateReadInbox(ctx context.Context, userId, peerId int32, readInboxId int32) {
	c.PrivateCore.UpdateReadInbox(ctx, userId, peerId, readInboxId)
}

func (c *localPrivateFacade) UpdateReadOutbox(ctx context.Context, userId, peerId int32, readOutboxId int32) {
	c.PrivateCore.UpdateReadOutbox(ctx, userId, peerId, readOutboxId)
}

func (c *localPrivateFacade) GetReadInboxMaxId(ctx context.Context, userId, peerId int32) int32 {
	return c.PrivateCore.GetReadInboxMaxId(ctx, userId, peerId)
}

func (c *localPrivateFacade) GetUserPinnedMessage(ctx context.Context, userId int32, peerId int32) (pinnedMsgId int32) {
	pinnedMsgId = c.PrivateCore.GetPinnedMessage(ctx, userId, peerId)
	return
}

func (c *localPrivateFacade) UpdateUserPinnedMessage(ctx context.Context, userId int32, peerId, pinnedMsgId int32) error {
	return c.PrivateCore.UpdatePinnedMessage(ctx, userId, peerId, pinnedMsgId)
}

func (c *localPrivateFacade) InsertOrUpdateDialogFilter(ctx context.Context, userId, id int32, dialogFilter *mtproto.DialogFilter) error {
	return c.PrivateCore.InsertOrUpdateDialogFilter(ctx, userId, id, dialogFilter)
}

func (c *localPrivateFacade) DeleteDialogFilter(ctx context.Context, userId, id int32) error {
	return c.PrivateCore.DeleteDialogFilter(ctx, userId, id)
}

func (c *localPrivateFacade) UpdateDialogFiltersOrder(ctx context.Context, userId int32, order []int32) error {
	return c.PrivateCore.UpdateDialogFiltersOrder(ctx, userId, order)
}

func (c *localPrivateFacade) GetDialogFilters(ctx context.Context, userId int32) model.DialogFilterExtList {
	return c.PrivateCore.GetDialogFilters(ctx, userId)
}

func init() {
	Register("local", New)
}
