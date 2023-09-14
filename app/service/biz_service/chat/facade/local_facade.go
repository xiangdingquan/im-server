package chat_facade

import (
	"context"

	"open.chat/app/service/biz_service/chat/core"
	"open.chat/app/service/biz_service/chat/dao"
	"open.chat/model"
	"open.chat/mtproto"
)

type localChatFacade struct {
	*core.ChatCore
}

func New() ChatFacade {
	return &localChatFacade{
		ChatCore: core.New(dao.New()),
	}
}

func (c *localChatFacade) SaveDraftMessage(ctx context.Context, userId, peerId int32, message *mtproto.DraftMessage) error {
	return c.ChatCore.SaveDraftMessage(ctx, userId, peerId, message)
}

func (c *localChatFacade) ClearDraftMessage(ctx context.Context, userId, peerId int32) error {
	return c.ChatCore.ClearDraftMessage(ctx, userId, peerId)
}

func (c *localChatFacade) GetAllDrafts(ctx context.Context, userId int32) ([]int32, []*mtproto.DraftMessage, error) {
	return c.ChatCore.GetAllDrafts(ctx, userId)
}

func (c *localChatFacade) ClearAllDrafts(ctx context.Context, userId int32) error {
	return c.ChatCore.ClearAllDrafts(ctx, userId)
}

func (c *localChatFacade) MarkDialogUnread(ctx context.Context, userId, peerId int32, unreadMark bool) error {
	return c.ChatCore.MarkDialogUnread(ctx, userId, peerId, unreadMark)
}

func (c *localChatFacade) ToggleDialogPin(ctx context.Context, userId, peerId int32, pinned bool) (int32, error) {
	return c.ChatCore.ToggleDialogPin(ctx, userId, peerId, pinned)
}

func (c *localChatFacade) GetDialogUnreadMarkList(ctx context.Context, userId int32) []*mtproto.DialogPeer {
	return c.ChatCore.GetDialogUnreadMarkList(ctx, userId)
}

func (c *localChatFacade) GetDialogsCount(ctx context.Context, userId int32, excludePinned bool, folderId int32) int {
	return c.ChatCore.GetDialogsCount(ctx, userId, excludePinned, folderId)
}

func (c *localChatFacade) GetDialogsByOffsetDate(ctx context.Context, userId int32, excludePinned bool, offsetDate, limit int32) model.DialogExtList {
	return c.ChatCore.GetDialogsByOffsetDate(ctx, userId, excludePinned, offsetDate, limit)
}

func (c *localChatFacade) GetDialogsByOffsetId(ctx context.Context, userId int32, excludePinned bool, offsetId, limit int32) (dialogs model.DialogExtList) {
	return c.ChatCore.GetDialogsByOffsetId(ctx, userId, excludePinned, offsetId, limit)
}

func (c *localChatFacade) GetDialogs(ctx context.Context, userId int32, excludePinned bool, folderId int32) model.DialogExtList {
	return c.ChatCore.GetDialogs(ctx, userId, excludePinned, folderId)
}

func (c *localChatFacade) GetDialogsByIdList(ctx context.Context, userId int32, idList []int32) model.DialogExtList {
	return c.ChatCore.GetDialogsByIdList(ctx, userId, idList)
}

func (c *localChatFacade) GetDialogFolder(ctx context.Context, userId, folderId int32) (*model.DialogExt, error) {
	return c.ChatCore.GetDialogFolder(ctx, userId, folderId)
}

func (c *localChatFacade) GetPinnedDialogs(ctx context.Context, userId, folderId int32) (dialogs model.DialogExtList) {
	return c.ChatCore.GetPinnedDialogs(ctx, userId, folderId)
}

func (c *localChatFacade) ReorderPinnedDialogs(ctx context.Context, userId int32, force bool, folderId int32, idList []int32) error {
	return c.ChatCore.ReorderPinnedDialogs(ctx, userId, force, folderId, idList)
}

func (c *localChatFacade) GetMutableChat(ctx context.Context, chatId int32, id ...int32) (chat *model.MutableChat, err error) {
	return c.ChatCore.GetMutableChat(ctx, chatId, id...)
}

func (c *localChatFacade) GetMutableChatListByIdList(ctx context.Context, selfUserId int32, chatId ...int32) (chats []*model.MutableChat, err error) {
	return c.ChatCore.GetMutableChatListByIdList(ctx, selfUserId, chatId...)
}

func (c *localChatFacade) GetChatListByIdList(ctx context.Context, selfUserId int32, idList []int32) (chats []*mtproto.Chat) {
	return c.ChatCore.GetChatListByIdList(ctx, selfUserId, idList)
}

func (c *localChatFacade) GetChatBySelfId(ctx context.Context, selfUserId, chatId int32) (chat *mtproto.Chat) {
	return c.ChatCore.GetChatBySelfId(ctx, selfUserId, chatId)
}

func (c *localChatFacade) GetAllChats(ctx context.Context, selfUserId int32) (chats []*mtproto.Chat) {
	return c.ChatCore.GetAllChats(ctx, selfUserId)
}

func (c *localChatFacade) GetPinnedDialogPeers(ctx context.Context, userId, folderId int32) model.DialogPinnedExtList {
	return c.ChatCore.GetPinnedDialogPeers(ctx, userId, folderId)
}

func (c *localChatFacade) EditPeerFoldersByIdList(ctx context.Context, userId, folderId int32, idList []int32) (map[int32]int64, error) {
	return c.ChatCore.EditPeerFoldersByIdList(ctx, userId, folderId, idList)
}

func (c *localChatFacade) GetTopMessage(ctx context.Context, userId, peerId int32) int32 {
	return c.ChatCore.GetTopMessage(ctx, userId, peerId)
}

func (c *localChatFacade) UpdateReadInbox(ctx context.Context, userId, peerId int32, readInboxId int32) {
	c.ChatCore.UpdateReadInbox(ctx, userId, peerId, readInboxId)
}

func (c *localChatFacade) UpdateReadOutbox(ctx context.Context, userId, peerId int32, readOutboxId int32) {
	c.ChatCore.UpdateReadOutbox(ctx, userId, peerId, readOutboxId)
}

func (c *localChatFacade) GetReadInboxMaxId(ctx context.Context, userId, peerId int32) int32 {
	return c.ChatCore.GetReadInboxMaxId(ctx, userId, peerId)
}

func (c *localChatFacade) CreateChat2(ctx context.Context, creatorId int32, userIdList []int32, title string) (chat *model.MutableChat, err error) {
	return c.ChatCore.CreateChat2(ctx, creatorId, userIdList, title)
}

func (c *localChatFacade) DeleteChatUser(ctx context.Context, chatId, operatorId, deleteUserId int32) (*model.MutableChat, error) {
	return c.ChatCore.DeleteChatUser(ctx, chatId, operatorId, deleteUserId)
}

func (c *localChatFacade) EditChatTitle(ctx context.Context, chatId, editUserId int32, title string) (*model.MutableChat, error) {
	return c.ChatCore.EditChatTitle(ctx, chatId, editUserId, title)
}

func (c *localChatFacade) EditChatAbout(ctx context.Context, chatId, editUserId int32, about string) (*model.MutableChat, error) {
	return c.ChatCore.EditChatAbout(ctx, chatId, editUserId, about)
}

func (c *localChatFacade) EditChatNotice(ctx context.Context, chatId, editUserId int32, notice string) (*model.MutableChat, error) {
	return c.ChatCore.EditChatNotice(ctx, chatId, editUserId, notice)
}

func (c *localChatFacade) EditChatPhoto(ctx context.Context, chatId, editUserId int32, chatPhoto *mtproto.Photo) (*model.MutableChat, error) {
	return c.ChatCore.EditChatPhoto(ctx, chatId, editUserId, chatPhoto)
}

func (c *localChatFacade) EditChatAdmin(ctx context.Context, chatId, operatorId, editChatAdminId int32, isAdmin bool) (*model.MutableChat, error) {
	return c.ChatCore.EditChatAdmin(ctx, chatId, operatorId, editChatAdminId, isAdmin)
}

func (c *localChatFacade) EditChatDefaultBannedRights(ctx context.Context, chatId, operatorId int32, bannedRights *mtproto.ChatBannedRights) (*model.MutableChat, error) {
	return c.ChatCore.EditChatDefaultBannedRights(ctx, chatId, operatorId, bannedRights)
}

func (c *localChatFacade) ToggleChatAdmins(ctx context.Context, chatId, operatorId int32, adminsEnabled bool) (*model.MutableChat, error) {
	return c.ChatCore.ToggleChatAdmins(ctx, chatId, operatorId, adminsEnabled)
}

func (c *localChatFacade) ExportChatInvite(ctx context.Context, chatId, inviteUserId int32) (string, error) {
	return c.ChatCore.ExportChatInvite(ctx, chatId, inviteUserId)
}

func (c *localChatFacade) AddChatUser(ctx context.Context, chatId, inviterId, userId int32) (*model.MutableChat, error) {
	return c.ChatCore.AddChatUser(ctx, chatId, inviterId, userId)
}

func (c *localChatFacade) GetMutableChatByLink(ctx context.Context, link string, id ...int32) (chat *model.MutableChat, err error) {
	return c.ChatCore.GetMutableChatByLink(ctx, link, id...)
}

func (c *localChatFacade) MigratedToChannel(ctx context.Context, chat *model.MutableChat, id int32, accessHash int64) error {
	return c.ChatCore.MigratedToChannel(ctx, chat, id, accessHash)
}

func (c *localChatFacade) GetChatParticipantIdList(ctx context.Context, chatId int32) ([]int32, error) {
	return c.ChatCore.GetChatParticipantIdList(ctx, chatId)
}

func (c *localChatFacade) UpdateChatPinnedMessage(ctx context.Context, userId, chatId int32, chatPinnedList map[int32]int32) (*model.MutableChat, error) {
	return c.ChatCore.UpdatePinnedMessage(ctx, userId, chatId, chatPinnedList)
}

func (c *localChatFacade) UpdateUnChatPinnedMessage(ctx context.Context, userId, chatId int32) (*model.MutableChat, error) {
	return c.ChatCore.UpdateUnPinnedMessage(ctx, userId, chatId)
}

func (c *localChatFacade) GetUsersChatIdList(ctx context.Context, id []int32) map[int32][]int32 {
	return c.ChatCore.GetUsersChatIdList(ctx, id)
}

func (c *localChatFacade) GetMyChatList(ctx context.Context, userId int32, isCreator bool) []*mtproto.Chat {
	return c.ChatCore.GetMyChatList(ctx, userId, isCreator)
}

func (c *localChatFacade) CheckParticipantIsExist(ctx context.Context, userId int32, chatIdList []int32) bool {
	chatIds := c.GetUsersChatIdList(ctx, []int32{userId})
	if len(chatIds) == 0 {
		return false
	} else if ids, ok := chatIds[userId]; !ok {
		return false
	} else {
		for _, id := range chatIdList {
			for _, id2 := range ids {
				if id == id2 {
					return true
				}
			}
		}
	}
	return false
}

func (c *localChatFacade) GetFilterKeywords(ctx context.Context, id uint32) (keywords []string, err error) {
	return c.ChatCore.GetFilterKeywords(ctx, id)
}

func init() {
	Register("local", New)
}
