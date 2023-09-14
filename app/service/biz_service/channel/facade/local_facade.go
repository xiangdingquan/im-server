package channel_facade

import (
	"context"
	"math"
	"time"

	"github.com/gogo/protobuf/types"
	"open.chat/app/service/biz_service/channel/core"
	"open.chat/app/service/biz_service/channel/dao"
	"open.chat/model"
	"open.chat/mtproto"
)

type localChannelFacade struct {
	*core.ChannelCore
}

func New() ChannelFacade {
	return &localChannelFacade{
		ChannelCore: core.New(dao.New()),
	}
}

func (c *localChannelFacade) SaveDraftMessage(ctx context.Context, userId, peerId int32, message *mtproto.DraftMessage) error {
	return c.ChannelCore.SaveDraftMessage(ctx, userId, peerId, message)
}

func (c *localChannelFacade) ClearDraftMessage(ctx context.Context, userId, peerId int32) error {
	return c.ChannelCore.ClearDraftMessage(ctx, userId, peerId)
}

func (c *localChannelFacade) GetAllDrafts(ctx context.Context, userId int32) ([]int32, []*mtproto.DraftMessage, error) {
	return c.ChannelCore.GetAllDrafts(ctx, userId)
}

func (c *localChannelFacade) ClearAllDrafts(ctx context.Context, userId int32) error {
	return c.ChannelCore.ClearAllDrafts(ctx, userId)
}

func (c *localChannelFacade) MarkDialogUnread(ctx context.Context, userId, peerId int32, unreadMark bool) error {
	return c.ChannelCore.MarkDialogUnread(ctx, userId, peerId, unreadMark)
}

func (c *localChannelFacade) ToggleDialogPin(ctx context.Context, userId, peerId int32, pinned bool) (int32, error) {
	return c.ChannelCore.ToggleDialogPin(ctx, userId, peerId, pinned)
}

func (c *localChannelFacade) GetDialogUnreadMarkList(ctx context.Context, userId int32) []*mtproto.DialogPeer {
	return c.ChannelCore.GetDialogUnreadMarkList(ctx, userId)
}

func (c *localChannelFacade) GetDialogsCount(ctx context.Context, userId int32, excludePinned bool, folderId int32) int {
	return c.ChannelCore.GetDialogsCount(ctx, userId, excludePinned, folderId)
}

func (c *localChannelFacade) GetDialogsByOffsetDate(ctx context.Context, userId int32, excludePinned bool, offsetDate, limit int32) model.DialogExtList {
	return c.ChannelCore.GetDialogsByOffsetDate(ctx, userId, excludePinned, offsetDate, limit)
}

func (c *localChannelFacade) GetDialogsByOffsetId(ctx context.Context, userId int32, excludePinned bool, offsetId, limit int32) (dialogs model.DialogExtList) {
	return c.ChannelCore.GetDialogsByOffsetId(ctx, userId, excludePinned, offsetId, limit)
}

func (c *localChannelFacade) GetDialogs(ctx context.Context, userId int32, excludePinned bool, folderId int32) model.DialogExtList {
	return c.ChannelCore.GetDialogs(ctx, userId, excludePinned, folderId)
}

func (c *localChannelFacade) GetDialogsByIdList(ctx context.Context, userId int32, idList []int32) model.DialogExtList {
	return c.ChannelCore.GetDialogsByIdList(ctx, userId, idList)
}

func (c *localChannelFacade) GetPinnedDialogs(ctx context.Context, userId, folderId int32) (dialogs model.DialogExtList) {
	return c.ChannelCore.GetPinnedDialogs(ctx, userId, folderId)
}

func (c *localChannelFacade) ReorderPinnedDialogs(ctx context.Context, userId int32, force bool, folderId int32, idList []int32) error {
	return c.ChannelCore.ReorderPinnedDialogs(ctx, userId, force, folderId, idList)
}

func (c *localChannelFacade) GetDialogFolder(ctx context.Context, userId, folderId int32) (*model.DialogExt, error) {
	return c.ChannelCore.GetDialogFolder(ctx, userId, folderId)
}

func (c *localChannelFacade) GetPinnedDialogPeers(ctx context.Context, userId, folderId int32) model.DialogPinnedExtList {
	return c.ChannelCore.GetPinnedDialogPeers(ctx, userId, folderId)
}

func (c *localChannelFacade) EditPeerFoldersByIdList(ctx context.Context, userId, folderId int32, idList []int32) (map[int32]int64, error) {
	return c.ChannelCore.EditPeerFoldersByIdList(ctx, userId, folderId, idList)
}

func (c *localChannelFacade) SetDiscussionGroup(ctx context.Context, userId, broadcastId, groupId int32) error {
	return c.ChannelCore.SetDiscussionGroup(ctx, userId, broadcastId, groupId)
}

func (c *localChannelFacade) SearchChannelByTitle(ctx context.Context, q string) []int32 {
	return c.ChannelCore.SearchChannelByTitle(ctx, q)
}

func (c *localChannelFacade) GetChannelListByIdList(ctx context.Context, userId int32, id ...int32) []*mtproto.Chat {
	return c.ChannelCore.GetChannelListByIdList(ctx, userId, id)
}

func (c *localChannelFacade) GetChannelById(ctx context.Context, userId, id int32) *mtproto.Chat {
	return c.ChannelCore.GetChannelById(ctx, userId, id)
}

func (c *localChannelFacade) GetChannelParticipantIdList(ctx context.Context, channelId int32) []int32 {
	return c.ChannelCore.GetChannelParticipantIdList(ctx, channelId)
}

func (c *localChannelFacade) GetChannelAdminParticipantIdList(ctx context.Context, channelId int32) []int32 {
	return c.ChannelCore.GetChannelAdminParticipantIdList(ctx, channelId)
}

func (c *localChannelFacade) GetUsersChannelIdList(ctx context.Context, id []int32) map[int32][]int32 {
	return c.ChannelCore.GetUsersChannelIdList(ctx, id)
}

func (c *localChannelFacade) GetAllChannels(ctx context.Context, userId int32) model.MutableChannels {
	return c.ChannelCore.GetAllChannels(ctx, userId)
}

func (c *localChannelFacade) CheckParticipantIsExist(ctx context.Context, userId int32, chatIdList []int32) bool {
	chatIds := c.GetUsersChannelIdList(ctx, []int32{userId})
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

func (c *localChannelFacade) GetMutableChannel(ctx context.Context, channelId int32, id ...int32) (chat *model.MutableChannel, err error) {
	return c.ChannelCore.GetMutableChannel(ctx, channelId, id...)
}

func (c *localChannelFacade) GetMutableChannelByLink(ctx context.Context, link string, id ...int32) (*model.MutableChannel, error) {
	return c.ChannelCore.GetMutableChannelByLink(ctx, link, id...)
}

func (c *localChannelFacade) GetTopMessage(ctx context.Context, userId, peerId int32) int32 {
	return c.ChannelCore.GetTopMessage(ctx, userId, peerId)
}

func (c *localChannelFacade) UpdateReadInbox(ctx context.Context, userId, peerId int32, readInboxId int32) {
	c.ChannelCore.UpdateReadInbox(ctx, userId, peerId, readInboxId)
}

func (c *localChannelFacade) UpdateReadOutbox(ctx context.Context, userId, peerId int32, readOutboxId int32) {
	c.ChannelCore.UpdateReadOutbox(ctx, userId, peerId, readOutboxId)
}

func (c *localChannelFacade) ReadOutboxHistory(ctx context.Context, channelId, userId, maxId int32) bool {
	return c.ChannelCore.ReadOutboxHistory(ctx, channelId, userId, maxId)
}

func (c *localChannelFacade) GetReadInboxMaxId(ctx context.Context, userId, peerId int32) int32 {
	return c.ChannelCore.GetReadInboxMaxId(ctx, userId, peerId)
}

func (c *localChannelFacade) CreateChannel(ctx context.Context, creatorId int32, secretKeyId int64, isBroadcast bool, title, about, notice string, geo *mtproto.InputGeoPoint, address *types.StringValue, randomId int64) (*model.MutableChannel, error) {
	return c.ChannelCore.CreateChannel(ctx, creatorId, secretKeyId, isBroadcast, title, about, notice, geo, address, randomId)
}

func (c *localChannelFacade) InviteToChannel(ctx context.Context, channelId, inviterId int32, id ...int32) (*model.MutableChannel, []int32, error) {
	return c.ChannelCore.InviteToChannel(ctx, channelId, inviterId, id...)
}

func (c *localChannelFacade) JoinChannel(ctx context.Context, channelId, joinId int32, force bool) (*model.MutableChannel, error) {
	return c.ChannelCore.JoinChannel(ctx, channelId, joinId, force)
}

func (c *localChannelFacade) LeaveChannel(ctx context.Context, channelId, userId int32) (*model.MutableChannel, error) {
	return c.ChannelCore.LeaveChannel(ctx, channelId, userId)
}

func (c *localChannelFacade) EditTitle(ctx context.Context, channelId, editUserId int32, title string) (*model.MutableChannel, error) {
	return c.ChannelCore.EditTitle(ctx, channelId, editUserId, title)
}

func (c *localChannelFacade) EditAbout(ctx context.Context, channelId, aboutUserId int32, about string) (*model.MutableChannel, error) {
	return c.ChannelCore.EditAbout(ctx, channelId, aboutUserId, about)
}

func (c *localChannelFacade) EditNotice(ctx context.Context, channelId, noticeUserId int32, notice string) (*model.MutableChannel, error) {
	return c.ChannelCore.EditNotice(ctx, channelId, noticeUserId, notice)
}

func (c *localChannelFacade) EditPhoto(ctx context.Context, channelId, editUserId int32, photo *mtproto.Photo) (*model.MutableChannel, error) {
	return c.ChannelCore.EditPhoto(ctx, channelId, editUserId, photo)
}

func (c *localChannelFacade) EditAdminRights(ctx context.Context, channelId, operatorId, editChannelAdminsId, adminRights int32, rank string) (*model.MutableChannel, bool, error) {
	return c.ChannelCore.EditAdminRights(ctx, channelId, operatorId, editChannelAdminsId, model.ChatAdminRights(adminRights), rank)
}

func (c *localChannelFacade) EditBanned(ctx context.Context, channelId, operatorId, bannedUserId int32, bannedRights model.ChatBannedRights) (*model.MutableChannel, bool, error) {
	return c.ChannelCore.EditBanned(ctx, channelId, operatorId, bannedUserId, bannedRights)
}

func (c *localChannelFacade) ExportChannelInvite(ctx context.Context, channelId, operatorId int32) (*model.MutableChannel, error) {
	return c.ChannelCore.ExportChannelInvite(ctx, channelId, operatorId)
}

func (c *localChannelFacade) ToggleSignatures(ctx context.Context, channelId, operatorId int32, enabled bool) (*model.MutableChannel, error) {
	return c.ChannelCore.ToggleSignatures(ctx, channelId, operatorId, enabled)
}

func (c *localChannelFacade) ToggleInvites(ctx context.Context, channelId, operatorId int32, enabled bool) (*model.MutableChannel, error) {
	return c.ChannelCore.ToggleInvites(ctx, channelId, operatorId, enabled)
}

func (c *localChannelFacade) UpdateUsername(ctx context.Context, channelId, operatorId int32, username string) (*model.MutableChannel, error) {
	return c.ChannelCore.UpdateUsername(ctx, channelId, operatorId, username)
}

func (c *localChannelFacade) TogglePreHistoryHidden(ctx context.Context, channelId, operatorId int32, enabled bool) (*model.MutableChannel, error) {
	return c.ChannelCore.TogglePreHistoryHidden(ctx, channelId, operatorId, enabled)
}

func (c *localChannelFacade) DeleteChannel(ctx context.Context, channelId, operatorId int32) (*model.MutableChannel, error) {
	return c.ChannelCore.DeleteChannel(ctx, channelId, operatorId)
}

func (c *localChannelFacade) EditChatDefaultBannedRights(ctx context.Context, channelId, operatorId int32, bannedRights *mtproto.ChatBannedRights) (*model.MutableChannel, error) {
	return c.ChannelCore.EditChatDefaultBannedRights(ctx, channelId, operatorId, bannedRights)
}

func (c *localChannelFacade) ToggleSlowMode(ctx context.Context, channelId, operatorId int32, seconds int32) (*model.MutableChannel, error) {
	return c.ChannelCore.ToggleSlowMode(ctx, channelId, operatorId, seconds)
}

func (c *localChannelFacade) DeleteHistory(ctx context.Context, channelId, operatorId, maxId int32) (*model.MutableChannel, error) {
	return c.ChannelCore.DeleteMyHistory(ctx, channelId, operatorId, maxId)
}

func (c *localChannelFacade) UpdatePinnedMessage(ctx context.Context, channelId, operatorId, id int32) (*model.MutableChannel, error) {
	return c.ChannelCore.UpdatePinnedMessage(ctx, channelId, operatorId, id)
}

func (c *localChannelFacade) EditLocation(ctx context.Context, channelId, operatorId int32, geo *mtproto.InputGeoPoint, address string) (bool, error) {
	return c.ChannelCore.EditLocation(ctx, channelId, operatorId, geo, address)
}

func (c *localChannelFacade) GetAdminedPublicChannels(ctx context.Context, userId int32, byLocation, checkLimit bool) (model.MutableChannels, error) {
	return c.ChannelCore.GetAdminedPublicChannels(ctx, userId, byLocation, checkLimit)
}

func (c *localChannelFacade) GetAdminLogs(ctx context.Context, channelId int32, q string, eventsFilter int32, admins []int32, maxId, minId int64, limit int32) []*mtproto.ChannelAdminLogEvent {
	return c.ChannelCore.GetAdminLogs(ctx, channelId, q, eventsFilter, admins, maxId, minId, limit)
}

func (c *localChannelFacade) GetLeftChannelList(ctx context.Context, userId, offset int32) (model.MutableChannels, error) {
	return c.ChannelCore.GetLeftChannelList(ctx, userId, offset)
}

func (c *localChannelFacade) GetMyAdminChannelList(ctx context.Context, userId int32) model.MutableChannels {
	return c.ChannelCore.GetMyAdminChannelList(ctx, userId)
}

func (c *localChannelFacade) MigrateFromChat(ctx context.Context, userId int32, secretKeyId int64, chat *model.MutableChat) (*model.MutableChannel, error) {
	return c.ChannelCore.MigrateFromChat(ctx, secretKeyId, chat)
}

func (c *localChannelFacade) GetChannelParticipantRecentList(ctx context.Context, channel *model.ImmutableChannel, offset, limit, hash int32) []*model.ImmutableChannelParticipant {
	return c.ChannelCore.GetChannelParticipantRecentList(ctx, channel, offset, limit, hash)
}

func (c *localChannelFacade) GetChannelParticipantAdminList(ctx context.Context, channel *model.ImmutableChannel, offset, limit, hash int32) []*model.ImmutableChannelParticipant {
	return c.ChannelCore.GetChannelParticipantAdminList(ctx, channel, offset, limit, hash)
}

func (c *localChannelFacade) GetChannelParticipantKickedList(ctx context.Context, channel *model.ImmutableChannel, q string, offset, limit, hash int32) []*model.ImmutableChannelParticipant {
	return c.ChannelCore.GetChannelParticipantKickedList(ctx, channel, q, offset, limit, hash)
}

func (c *localChannelFacade) GetChannelParticipantBotList(ctx context.Context, channel *model.ImmutableChannel, offset, limit, hash int32) []*model.ImmutableChannelParticipant {
	return c.ChannelCore.GetChannelParticipantBotList(ctx, channel, offset, limit, hash)
}

func (c *localChannelFacade) GetChannelParticipantBannedList(ctx context.Context, channel *model.ImmutableChannel, q string, offset, limit, hash int32) []*model.ImmutableChannelParticipant {
	return c.ChannelCore.GetChannelParticipantBannedList(ctx, channel, q, offset, limit, hash)
}

func (c *localChannelFacade) GetChannelParticipantListBySearch(ctx context.Context, channel *model.ImmutableChannel, q string, offset, limit, hash int32) []*model.ImmutableChannelParticipant {
	return c.ChannelCore.GetChannelParticipantListBySearch(ctx, channel, q, offset, limit, hash)
}

func (c *localChannelFacade) GetChannelParticipantContactList(ctx context.Context, channel *model.ImmutableChannel, q string, offset, limit, hash int32) []*model.ImmutableChannelParticipant {
	return c.ChannelCore.GetChannelParticipantContactList(ctx, channel, q, offset, limit, hash)
}

func (c *localChannelFacade) GetChannelParticipants(ctx context.Context, channel *model.ImmutableChannel, filter *mtproto.ChannelParticipantsFilter, offset, limit int32) (int32, []*model.ImmutableChannelParticipant) {
	return c.ChannelCore.GetChannelParticipants(ctx, channel, filter, offset, limit)
}

func (c *localChannelFacade) GetParticipantCounts(ctx context.Context, id int32) (int32, int32, int32, int32) {
	return c.ChannelCore.GetParticipantCounts(ctx, id)
}

func (c *localChannelFacade) GetChannelMessagesViews(ctx context.Context, channelId int32, idList []int32, increment bool) []int32 {
	return c.ChannelCore.GetChannelMessagesViews(ctx, channelId, idList, increment)
}

func (c *localChannelFacade) IncrementChannelMessagesViews(ctx context.Context, channelId int32, idList []int32) {
	c.ChannelCore.IncrementChannelMessagesViews(ctx, channelId, idList)
}

func (c *localChannelFacade) GetChannelMessagesViews2(ctx context.Context, channelId int32, idList []int32, increment bool) []*mtproto.MessageViews {
	return c.ChannelCore.GetChannelMessagesViews2(ctx, channelId, idList, increment)
}

func (c *localChannelFacade) SearchPublicChannels(ctx context.Context, offset int32, limit int32) []int32 {
	return c.ChannelCore.SearchPublicChannel(ctx, offset, limit)
}

func (c *localChannelFacade) SetLocated(ctx context.Context, lat, long float64, uid int32, has_expiration bool, expriation int32) error {
	_, err := c.ChannelCore.GeoAdd(ctx, lat, long, uid)
	if err != nil {
		return err
	}

	// pass 0x7fffffff to disable expiry, 0 to make the current geolocation private

	if !has_expiration {
		return nil
	}

	if expriation == 0 {
		_, err = c.ChannelCore.GeoDel(ctx, uid)
	} else if 0 < expriation && expriation < math.MaxInt32 {
		_, err = c.ChannelCore.SetGeoExpiration(ctx, uid, expriation)
	} else {
		_, err = c.ChannelCore.DelGeoExpiration(ctx, uid)
	}

	return err
}

func (c *localChannelFacade) SearchNearBy(ctx context.Context, lat, long float64, radius int32, limit int32) ([]*model.NearByUser, error) {
	l, err := c.ChannelCore.GeoRadius(ctx, lat, long, radius, limit)
	if err != nil {
		return nil, err
	}

	uidList := make([]int32, len(l))
	for i, v := range l {
		uidList[i] = v.Id
	}

	if len(uidList) == 0 {
		return nil, nil
	}

	m, err := c.ChannelCore.GetLocationExpirations(ctx, uidList)
	if err != nil {
		return nil, err
	}

	for _, v := range l {
		if e, ok := m[v.Id]; ok {
			v.Expiration = e
		} else {
			v.Expiration = math.MaxInt32
		}
	}

	return l, nil
}

func (c *localChannelFacade) ClearExpiredLocation(ctx context.Context) error {
	ids, err := c.ChannelCore.GetExpiredLocation(ctx, int32(time.Now().UTC().Unix()))
	if err != nil {
		return err
	}

	for _, id := range ids {
		_, _ = c.ChannelCore.GeoDel(ctx, id)
	}

	return nil
}

func init() {
	Register("local", New)
}

func (c *localChannelFacade) GetFilterKeywords(ctx context.Context, id uint32) (keywords []string, err error) {
	return c.ChannelCore.GetFilterKeywords(ctx, id)
}
