package user_client

import (
	"context"
	"fmt"

	"open.chat/model"
	"open.chat/mtproto"
)

type UserFacade interface {
	// daysTTL
	SetAccountDaysTTL(ctx context.Context, userId int32, ttl int32) error
	GetAccountDaysTTL(ctx context.Context, userId int32) (int32, error)

	// photo
	GetCacheUserPhotos(ctx context.Context, userId int32) (*model.UserPhotos, error)
	PutCacheUserPhotos(ctx context.Context, userId int32, photos *model.UserPhotos) error

	// notify_setting
	GetNotifySettings(ctx context.Context, userId int32, peer *model.PeerUtil) (*mtproto.PeerNotifySettings, error)
	SetNotifySettings(ctx context.Context, userId int32, peer *model.PeerUtil, settings *mtproto.PeerNotifySettings) error
	ResetNotifySettings(ctx context.Context, userId int32) error
	GetNotifySettingsByPeerList(ctx context.Context, userId int32, userIdList, chatIdList, channelIdList []int32) ([]map[int32]*mtproto.PeerNotifySettings, error)
	GetAllNotifySettings(ctx context.Context, userId int32) (settings map[int64]*mtproto.PeerNotifySettings, err error)

	// privacy
	GetPrivacy(ctx context.Context, userId int32, keyType int) ([]*mtproto.PrivacyRule, error)
	SetPrivacy(ctx context.Context, userId int32, keyType int, rules []*mtproto.PrivacyRule) error
	CheckPrivacy(ctx context.Context, keyType int, selfId, peerId int32, isContact bool) bool

	// user
	CreateNewUser(ctx context.Context, keyId int64, phoneNumber, countryCode, firstName, lastName string) (*mtproto.User, error)
	DeleteUser(ctx context.Context, userId int32, reason string) (bool, error)
	GetUserSelf(ctx context.Context, id int32) (*mtproto.User, error)
	GetUserById(ctx context.Context, selfUserId, userId int32) (*mtproto.User, error)
	GetUserListByIdList(ctx context.Context, selfUserId int32, userIdList []int32) []*mtproto.User
	GetUserName(ctx context.Context, userId int32) string
	GetUserByToken(ctx context.Context, token string) (*mtproto.User, error)

	CheckPhoneNumberExist(ctx context.Context, phoneNumber string) (bool, error)
	GetUserByPhoneNumber(ctx context.Context, selfId int32, phoneNumber string) (*mtproto.User, error)
	GetUserSelfByPhoneNumber(ctx context.Context, phoneNumber string) (*mtproto.User, error)
	GetUserListByPhoneNumberList(ctx context.Context, selfId int32, phoneNumberList []string) []*mtproto.User

	GetCountryCodeByUser(ctx context.Context, userId int32) (string, error)
	GetUserByUsername(ctx context.Context, selfId int32, username string) (*mtproto.User, error)
	GetPasswordByPhone(ctx context.Context, phoneNumber string) (bool, string, error)
	GetPhoneAndPassword(ctx context.Context, username string) (string, string, error)
	GetPasswordById(ctx context.Context, selfId int32) (string, error)
	GetPhoneById(ctx context.Context, selfId int32) (string, error)
	UpdateUsername(ctx context.Context, id int32, username string) (bool, error)
	UpdateUserPassword(ctx context.Context, id int32, password string) (bool, error)
	UpdateUserInviter(ctx context.Context, id int32, inviter int32) (bool, error)
	UpdateUserInfoExt(ctx context.Context, selfId int32, gender int32, birth string, country, countryCode, province, city, cityCode string) (bool, error)

	// Account
	CheckSessionPasswordNeeded(ctx context.Context, userId int32) (bool, error)
	CheckRecoverCode(ctx context.Context, userId int32, code string) error
	GetPasswordRecovery(ctx context.Context, userId int32) (*mtproto.Auth_PasswordRecovery, error)

	// UserFull
	GetFullUser(ctx context.Context, selfId, userId int32) (*mtproto.UserFull, error)

	// wallPaper
	GetWallPaperList(ctx context.Context) ([]*mtproto.WallPaper, error)

	// report
	Report(ctx context.Context, userId, reportType, peerType, peerId, messageSenderUserId, messageId, reason int32, text string) (bool, error)
	ReportIdList(ctx context.Context, userId, reportType, peerType, peerId, messageSenderUserId int32, messageIdList []int32, reason int32, text string) (bool, error)

	// profile
	UpdateAbout(ctx context.Context, userId int32, about string) (bool, error)
	UpdateFirstAndLastName(ctx context.Context, userId int32, firstName, lastName string) (bool, error)
	GetFirstAndLastName(ctx context.Context, userId int32) (string, string, error)
	UpdateVerified(ctx context.Context, userId int32, verified bool) (bool, error)

	//
	UpdateUserStatus(ctx context.Context, userId int32, lastSeenAt int64) error
	GetContactUserIdList(ctx context.Context, userId int32) []int32
	GetMutualContactUserIdList(ctx context.Context, mutual bool, userId int32) []int32
	GetUserStatus2(ctx context.Context, selfId, userId int32, isContact, isBlocked bool) *mtproto.UserStatus
	GetLastSeenList(ctx context.Context, id []int32) map[int32]int64

	//
	CheckUserAccessHash(ctx context.Context, userId int32, hash int64) bool

	//
	BlockUser(ctx context.Context, userId, blockId int32) bool
	UnBlockUser(ctx context.Context, userId, blockId int32) bool
	IsBlockedByUser(ctx context.Context, selfUserId, id int32) bool
	CheckBlockUserList(ctx context.Context, selfUserId int32, idList []int32) []int32

	DeleteContact(ctx context.Context, userId, deleteId int32, mutual bool) bool
	GetContactLink(ctx context.Context, userId, contactId int32) (myLink, foreignLink *mtproto.ContactLink)
	GetBlockedList(ctx context.Context, userId, offset, limit int32) []*mtproto.ContactBlocked
	GetContactList(ctx context.Context, userId int32) []*mtproto.Contact
	GetAllContactList(ctx context.Context, userId int32) []*mtproto.Contact
	SearchContacts(ctx context.Context, userId int32, q string, limit int32) ([]int32, []int32)
	BackupPhoneBooks(ctx context.Context, authKeyId int64, contacts []*mtproto.InputContact)
	ImportContacts(ctx context.Context, userId int32, contacts []*mtproto.InputContact) ([]*mtproto.ImportedContact, []*mtproto.PopularContact, []int32)
	GetContactAndMutual(ctx context.Context, selfUserId, id int32) (bool, bool)
	AddContact(ctx context.Context, selfUserId, contactId int32, addPhonePrivacyException bool, contactFirstName, contactLastName string) (bool, error)
	CheckContact(ctx context.Context, selfUserId, id int32) bool
	GetImportersByPhone(ctx context.Context, phone string) []*mtproto.InputContact
	DeleteImportersByPhone(ctx context.Context, phone string)

	// search
	SearchChannelParticipants(ctx context.Context, selfId int32, channelId int32, q string) []*mtproto.User

	// bots
	// IsBot(userId int32) bool
	SetBotCommands(ctx context.Context, botId int32, commands []*mtproto.BotCommand) error
	IsBot(ctx context.Context, id int32) bool

	//
	GetMutableUsers(ctx context.Context, idList ...int32) model.MutableUsers
	// GetImmutableUsers(ctx context.Context, idList ...int32) (map[int32]*mtproto.User, error)

	CreateNewPredefinedUser(ctx context.Context, phone, firstName, lastName, username, code string, verified bool) (*mtproto.PredefinedUser, error)
	GetPredefinedUser(ctx context.Context, phone string) (*mtproto.PredefinedUser, error)
	GetPredefinedUserList(ctx context.Context) []*mtproto.PredefinedUser
	UpdatePredefinedFirstAndLastName(ctx context.Context, phone, firstName, lastName string) (*mtproto.PredefinedUser, error)
	UpdatePredefinedVerified(ctx context.Context, phone string, verified bool) (*mtproto.PredefinedUser, error)
	UpdatePredefinedUsername(ctx context.Context, phone, username string) (*mtproto.PredefinedUser, error)
	UpdatePredefinedCode(ctx context.Context, phone, code string) (*mtproto.PredefinedUser, error)
	PredefinedBindRegisteredUserId(ctx context.Context, phone string, registeredUserId int32) bool

	AddPeerSettings(ctx context.Context, selfId int32, peer *model.PeerUtil, settings *mtproto.PeerSettings) error
	GetPeerSettings(ctx context.Context, selfId int32, peer *model.PeerUtil) (*mtproto.PeerSettings, error)
	DeletePeerSettings(ctx context.Context, selfId int32, peer *model.PeerUtil) error

	// GetCustomerServiceList Customer service id list
	GetCustomerServiceList(ctx context.Context) []int32
}

type Instance func() UserFacade

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

func NewUserFacade(name string) (inst UserFacade, err error) {
	instanceFunc, ok := instances[name]
	if !ok {
		err = fmt.Errorf("unknown adapter name %q (forgot to import?)", name)
		return
	}
	inst = instanceFunc()
	return
}
