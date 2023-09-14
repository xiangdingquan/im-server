package channel_facade

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/types"
	"open.chat/model"
	"open.chat/mtproto"
)

type ChannelFacade interface {
	// draft
	SaveDraftMessage(ctx context.Context, userId, peerId int32, message *mtproto.DraftMessage) error
	ClearDraftMessage(ctx context.Context, userId, peerId int32) error
	GetAllDrafts(ctx context.Context, userId int32) ([]int32, []*mtproto.DraftMessage, error)
	ClearAllDrafts(ctx context.Context, userId int32) error

	// dialog
	MarkDialogUnread(ctx context.Context, userId, channelId int32, unreadMark bool) error
	ToggleDialogPin(ctx context.Context, userId, channelId int32, pinned bool) (int32, error)
	GetDialogUnreadMarkList(ctx context.Context, userId int32) []*mtproto.DialogPeer
	GetDialogsByOffsetDate(ctx context.Context, userId int32, excludePinned bool, offsetDate, limit int32) model.DialogExtList
	GetDialogs(ctx context.Context, userId int32, excludePinned bool, folderId int32) model.DialogExtList
	GetDialogsByIdList(ctx context.Context, userId int32, idList []int32) model.DialogExtList
	GetDialogsCount(ctx context.Context, userId int32, excludePinned bool, folderId int32) int
	GetPinnedDialogs(ctx context.Context, userId, folderId int32) (dialogs model.DialogExtList)
	ReorderPinnedDialogs(ctx context.Context, userId int32, force bool, folderId int32, idList []int32) error

	GetDialogFolder(ctx context.Context, userId, folderId int32) (*model.DialogExt, error)
	GetPinnedDialogPeers(ctx context.Context, userId, folderId int32) model.DialogPinnedExtList
	EditPeerFoldersByIdList(ctx context.Context, userId, folderId int32, idList []int32) (map[int32]int64, error)

	// message
	SearchChannelByTitle(ctx context.Context, q string) []int32
	GetChannelListByIdList(ctx context.Context, userId int32, id ...int32) []*mtproto.Chat
	GetChannelById(ctx context.Context, userId, id int32) *mtproto.Chat
	GetChannelParticipantIdList(ctx context.Context, channelId int32) []int32
	GetChannelAdminParticipantIdList(ctx context.Context, channelId int32) []int32
	GetUsersChannelIdList(ctx context.Context, id []int32) map[int32][]int32
	CheckParticipantIsExist(ctx context.Context, userId int32, chatIdList []int32) bool
	GetAllChannels(ctx context.Context, userId int32) model.MutableChannels

	GetMutableChannel(ctx context.Context, channelId int32, id ...int32) (chat *model.MutableChannel, err error)
	GetMutableChannelByLink(ctx context.Context, link string, id ...int32) (*model.MutableChannel, error)

	GetTopMessage(ctx context.Context, userId, peerId int32) int32
	UpdateReadInbox(ctx context.Context, userId, peerId int32, readInboxId int32)
	UpdateReadOutbox(ctx context.Context, userId, peerId int32, readOutboxId int32)
	ReadOutboxHistory(ctx context.Context, channelId, userId, maxId int32) bool
	GetReadInboxMaxId(ctx context.Context, userId int32, peerId int32) int32

	// logic
	CreateChannel(ctx context.Context, creatorId int32, secretKeyId int64, isBroadcast bool, title, about, notice string, geo *mtproto.InputGeoPoint, address *types.StringValue, randomId int64) (*model.MutableChannel, error)
	InviteToChannel(ctx context.Context, channelId, inviterId int32, id ...int32) (*model.MutableChannel, []int32, error)
	JoinChannel(ctx context.Context, channelId, joinId int32, force bool) (*model.MutableChannel, error)
	LeaveChannel(ctx context.Context, channelId, userId int32) (*model.MutableChannel, error)
	EditTitle(ctx context.Context, channelId, editUserId int32, title string) (*model.MutableChannel, error)
	EditAbout(ctx context.Context, channelId, aboutUserId int32, about string) (*model.MutableChannel, error)
	EditNotice(ctx context.Context, channelId, noticeUserId int32, notice string) (*model.MutableChannel, error)
	EditPhoto(ctx context.Context, channelId, editUserId int32, photo *mtproto.Photo) (*model.MutableChannel, error)
	EditAdminRights(ctx context.Context, channelId, operatorId, editChannelAdminsId, adminRights int32, rank string) (*model.MutableChannel, bool, error)
	EditBanned(ctx context.Context, channelId, operatorId, bannedUserId int32, bannedRights model.ChatBannedRights) (*model.MutableChannel, bool, error)
	ExportChannelInvite(ctx context.Context, channelId, operatorId int32) (*model.MutableChannel, error)
	ToggleSignatures(ctx context.Context, channelId, operatorId int32, enabled bool) (*model.MutableChannel, error)
	ToggleInvites(ctx context.Context, channelId, operatorId int32, enabled bool) (*model.MutableChannel, error)
	UpdateUsername(ctx context.Context, channelId, operatorId int32, username string) (*model.MutableChannel, error)
	TogglePreHistoryHidden(ctx context.Context, channelId, operatorId int32, enabled bool) (*model.MutableChannel, error)
	DeleteChannel(ctx context.Context, channelId, operatorId int32) (*model.MutableChannel, error)
	EditChatDefaultBannedRights(ctx context.Context, channelId, operatorId int32, bannedRights *mtproto.ChatBannedRights) (*model.MutableChannel, error)
	ToggleSlowMode(ctx context.Context, channelId, operatorId int32, seconds int32) (*model.MutableChannel, error)
	DeleteHistory(ctx context.Context, channelId, operatorId, maxId int32) (*model.MutableChannel, error)
	UpdatePinnedMessage(ctx context.Context, channelId, operatorId, id int32) (*model.MutableChannel, error)
	EditLocation(ctx context.Context, channelId, operatorId int32, geo *mtproto.InputGeoPoint, address string) (bool, error)

	SetDiscussionGroup(ctx context.Context, userId, broadcastId, groupId int32) error
	// channel
	GetAdminLogs(ctx context.Context, channelId int32, q string, eventsFilter int32, admins []int32, maxId, minId int64, limit int32) []*mtproto.ChannelAdminLogEvent
	GetAdminedPublicChannels(ctx context.Context, userId int32, byLocation, checkLimit bool) (model.MutableChannels, error)
	GetLeftChannelList(ctx context.Context, userId, offset int32) (model.MutableChannels, error)
	GetMyAdminChannelList(ctx context.Context, userId int32) model.MutableChannels

	MigrateFromChat(ctx context.Context, userId int32, secretKeyId int64, chat *model.MutableChat) (*model.MutableChannel, error)

	GetChannelParticipantRecentList(ctx context.Context, channel *model.ImmutableChannel, offset, limit, hash int32) []*model.ImmutableChannelParticipant
	GetChannelParticipantAdminList(ctx context.Context, channel *model.ImmutableChannel, offset, limit, hash int32) []*model.ImmutableChannelParticipant
	GetChannelParticipantKickedList(ctx context.Context, channel *model.ImmutableChannel, q string, offset, limit, hash int32) []*model.ImmutableChannelParticipant
	GetChannelParticipantBotList(ctx context.Context, channel *model.ImmutableChannel, offset, limit, hash int32) []*model.ImmutableChannelParticipant
	GetChannelParticipantBannedList(ctx context.Context, channel *model.ImmutableChannel, q string, offset, limit, hash int32) []*model.ImmutableChannelParticipant
	GetChannelParticipantListBySearch(ctx context.Context, channel *model.ImmutableChannel, q string, offset, limit, hash int32) []*model.ImmutableChannelParticipant
	GetChannelParticipantContactList(ctx context.Context, channel *model.ImmutableChannel, q string, offset, limit, hash int32) []*model.ImmutableChannelParticipant
	// GetAvailableMinId(ctx context.Context, userId, channelId int32) int32
	GetChannelParticipants(ctx context.Context, channel *model.ImmutableChannel, filter *mtproto.ChannelParticipantsFilter, offset, limit int32) (int32, []*model.ImmutableChannelParticipant)

	GetParticipantCounts(ctx context.Context, id int32) (int32, int32, int32, int32)

	GetChannelMessagesViews(ctx context.Context, channelId int32, idList []int32, increment bool) []int32
	IncrementChannelMessagesViews(ctx context.Context, channelId int32, idList []int32)
	GetChannelMessagesViews2(ctx context.Context, channelId int32, idList []int32, increment bool) []*mtproto.MessageViews

	SearchPublicChannels(ctx context.Context, offset int32, limit int32) []int32
	SetLocated(ctx context.Context, lat, long float64, uid int32, has_expiration bool, expriation int32) error
	SearchNearBy(ctx context.Context, lat, long float64, radius int32, limit int32) ([]*model.NearByUser, error)
	ClearExpiredLocation(ctx context.Context) error

	GetFilterKeywords(ctx context.Context, id uint32) ([]string, error)
}
type Instance func() ChannelFacade

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

func NewChannelFacade(name string) (inst ChannelFacade, err error) {
	instanceFunc, ok := instances[name]
	if !ok {
		err = fmt.Errorf("unknown adapter name %q (forgot to import?)", name)
		return
	}
	inst = instanceFunc()
	return
}
