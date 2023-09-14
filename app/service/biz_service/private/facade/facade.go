package private_facade

import (
	"context"
	"fmt"

	"open.chat/model"
	"open.chat/mtproto"
)

type PrivateFacade interface {
	// draft
	SaveDraftMessage(ctx context.Context, userId, peerId int32, message *mtproto.DraftMessage) error
	ClearDraftMessage(ctx context.Context, userId, peerId int32) error
	GetAllDrafts(ctx context.Context, userId int32) ([]int32, []*mtproto.DraftMessage, error)
	ClearAllDrafts(ctx context.Context, userId int32) error

	// dialog
	MarkDialogUnread(ctx context.Context, userId, peerId int32, unreadMark bool) error
	ToggleDialogPin(ctx context.Context, userId, peerId int32, pinned bool) (int32, error)
	GetDialogUnreadMarkList(ctx context.Context, userId int32) []*mtproto.DialogPeer
	GetDialogsCount(ctx context.Context, userId int32, excludePinned bool, folderId int32) int
	GetDialogsByOffsetId(ctx context.Context, userId int32, excludePinned bool, offsetId, limit int32) (dialogs model.DialogExtList)
	GetDialogsByOffsetDate(ctx context.Context, userId int32, excludePinned bool, offsetDate, limit int32) model.DialogExtList
	GetDialogs(ctx context.Context, userId int32, excludePinned bool, folderId int32) model.DialogExtList
	GetDialogsByIdList(ctx context.Context, userId int32, idList []int32) model.DialogExtList
	GetPinnedDialogs(ctx context.Context, userId, folderId int32) (dialogs model.DialogExtList)
	ReorderPinnedDialogs(ctx context.Context, userId int32, force bool, folderId int32, idList []int32) error

	GetDialogFolder(ctx context.Context, userId, folderId int32) (*model.DialogExt, error)
	GetPinnedDialogPeers(ctx context.Context, userId, folderId int32) model.DialogPinnedExtList
	EditPeerFoldersByIdList(ctx context.Context, userId, folderId int32, idList []int32) (map[int32]int64, error)

	//
	GetTopMessage(ctx context.Context, userId, peerId int32) int32
	UpdateReadInbox(ctx context.Context, userId, peerId int32, readInboxId int32)
	UpdateReadOutbox(ctx context.Context, userId, peerId int32, readOutboxId int32)
	GetReadInboxMaxId(ctx context.Context, userId int32, peerId int32) int32

	//
	GetUserPinnedMessage(ctx context.Context, userId int32, peerId int32) (pinnedMsgId int32)
	UpdateUserPinnedMessage(ctx context.Context, userId int32, peerId, pinnedMsgId int32) error

	InsertOrUpdateDialogFilter(ctx context.Context, userId, id int32, dialogFilter *mtproto.DialogFilter) error
	DeleteDialogFilter(ctx context.Context, userId, id int32) error
	UpdateDialogFiltersOrder(ctx context.Context, userId int32, order []int32) error
	GetDialogFilters(ctx context.Context, userId int32) model.DialogFilterExtList
}

type Instance func() PrivateFacade

var instances = make(map[string]Instance)

func Register(name string, inst Instance) {
	if inst == nil {
		panic("register instance is nil")
	}
	if _, ok := instances[name]; ok {
		panic("register called twice for instance " + name)
	}
	instances[name] = inst
}

func NewPrivateFacade(name string) (inst PrivateFacade, err error) {
	instanceFunc, ok := instances[name]
	if !ok {
		err = fmt.Errorf("unknown adapter name %q (forgot to import?)", name)
		return
	}
	inst = instanceFunc()
	return
}
