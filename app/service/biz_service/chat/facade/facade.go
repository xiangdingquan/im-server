package chat_facade

import (
	"context"
	"fmt"

	"open.chat/model"
	"open.chat/mtproto"
)

type ChatFacade interface {
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
	GetDialogsByOffsetDate(ctx context.Context, userId int32, excludePinned bool, offsetDate, limit int32) model.DialogExtList
	GetDialogsByOffsetId(ctx context.Context, userId int32, excludePinned bool, offsetId, limit int32) (dialogs model.DialogExtList)
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
	GetReadInboxMaxId(ctx context.Context, userId, peerId int32) int32

	// message
	GetMutableChat(ctx context.Context, chatId int32, id ...int32) (chat *model.MutableChat, err error)
	GetChatListByIdList(ctx context.Context, selfUserId int32, idList []int32) (chats []*mtproto.Chat)
	GetChatBySelfId(ctx context.Context, selfUserId, chatId int32) (chat *mtproto.Chat)
	GetAllChats(ctx context.Context, selfUserId int32) (chats []*mtproto.Chat)

	CreateChat2(ctx context.Context, creatorId int32, userIdList []int32, title string) (chat *model.MutableChat, err error)
	DeleteChatUser(ctx context.Context, chatId, operatorId, deleteUserId int32) (*model.MutableChat, error)
	EditChatTitle(ctx context.Context, chatId, editUserId int32, title string) (*model.MutableChat, error)
	EditChatAbout(ctx context.Context, chatId, editUserId int32, about string) (*model.MutableChat, error)
	EditChatNotice(ctx context.Context, chatId, editUserId int32, notice string) (*model.MutableChat, error)
	EditChatPhoto(ctx context.Context, chatId, editUserId int32, chatPhoto *mtproto.Photo) (*model.MutableChat, error)
	EditChatAdmin(ctx context.Context, chatId, operatorId, editChatAdminId int32, isAdmin bool) (*model.MutableChat, error)
	EditChatDefaultBannedRights(ctx context.Context, chatId, operatorId int32, bannedRights *mtproto.ChatBannedRights) (*model.MutableChat, error)
	ToggleChatAdmins(ctx context.Context, chatId, operatorId int32, adminsEnabled bool) (*model.MutableChat, error)
	ExportChatInvite(ctx context.Context, chatId, inviteUserId int32) (string, error)
	AddChatUser(ctx context.Context, chatId, inviterId, userId int32) (*model.MutableChat, error)
	GetMutableChatByLink(ctx context.Context, link string, id ...int32) (chat *model.MutableChat, err error)

	MigratedToChannel(ctx context.Context, chat *model.MutableChat, id int32, accessHash int64) error

	GetChatParticipantIdList(ctx context.Context, chatId int32) ([]int32, error)

	UpdateChatPinnedMessage(ctx context.Context, userId, chatId int32, chatPinnedList map[int32]int32) (*model.MutableChat, error)
	UpdateUnChatPinnedMessage(ctx context.Context, userId, chatId int32) (*model.MutableChat, error)

	GetUsersChatIdList(ctx context.Context, id []int32) map[int32][]int32
	GetMyChatList(ctx context.Context, userId int32, isCreator bool) []*mtproto.Chat
	CheckParticipantIsExist(ctx context.Context, userId int32, chatIdList []int32) bool

	GetFilterKeywords(ctx context.Context, id uint32) ([]string, error)
}

type Instance func() ChatFacade

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

func NewChatFacade(name string) (inst ChatFacade, err error) {
	instanceFunc, ok := instances[name]
	if !ok {
		err = fmt.Errorf("unknown adapter name %q (forgot to import?)", name)
		return
	}
	inst = instanceFunc()
	return
}
