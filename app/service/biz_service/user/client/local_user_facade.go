package user_client

import (
	"context"

	"open.chat/app/service/biz_service/user/internal/core"
	"open.chat/app/service/biz_service/user/internal/dao"
	"open.chat/model"
	"open.chat/mtproto"
)

type localUserFacade struct {
	*core.UserCore
}

func localUserFacadeInstance() UserFacade {
	return &localUserFacade{
		UserCore: core.New(dao.New()),
	}
}

// ///////////////////////////////////////////////////////////////////////////
func (c *localUserFacade) CreateNewUser(ctx context.Context, keyId int64, phoneNumber, countryCode, firstName, lastName string) (*mtproto.User, error) {
	return c.UserCore.CreateNewUser(ctx, keyId, phoneNumber, countryCode, firstName, lastName)
}

func (c *localUserFacade) DeleteUser(ctx context.Context, userId int32, reason string) (bool, error) {
	return c.UserCore.DeleteUser(ctx, userId, reason)
}

// ///////////////////////////////////////////////////////////////////////////
func (c *localUserFacade) SetAccountDaysTTL(ctx context.Context, userId int32, ttl int32) error {
	return c.UserCore.SetAccountDaysTTL(ctx, userId, ttl)
}

func (c *localUserFacade) GetAccountDaysTTL(ctx context.Context, userId int32) (int32, error) {
	return c.UserCore.GetAccountDaysTTL(ctx, userId)
}

// ///////////////////////////////////////////////////////////////////////////
func (c *localUserFacade) GetNotifySettings(
	ctx context.Context,
	userId int32,
	peer *model.PeerUtil) (*mtproto.PeerNotifySettings, error) {

	return c.UserCore.GetNotifySettings(ctx, userId, peer)
}

func (c *localUserFacade) GetNotifySettingsByPeerList(
	ctx context.Context,
	userId int32,
	userIdList, chatIdList, channelIdList []int32) ([]map[int32]*mtproto.PeerNotifySettings, error) {

	return c.UserCore.GetNotifySettingsByPeerList(ctx, userId, userIdList, chatIdList, channelIdList)
}

func (c *localUserFacade) SetNotifySettings(ctx context.Context, userId int32, peer *model.PeerUtil, settings *mtproto.PeerNotifySettings) error {
	return c.UserCore.SetNotifySettings(ctx, userId, peer, settings)
}

func (c *localUserFacade) ResetNotifySettings(ctx context.Context, userId int32) error {
	return c.UserCore.ResetNotifySettings(ctx, userId)
}

func (c *localUserFacade) GetAllNotifySettings(ctx context.Context, userId int32) (settings map[int64]*mtproto.PeerNotifySettings, err error) {
	return c.UserCore.GetAllNotifySettings(ctx, userId)
}

// ///////////////////////////////////////////////////////////////////////////
func (c *localUserFacade) GetCacheUserPhotos(ctx context.Context, userId int32) (*model.UserPhotos, error) {
	return c.UserCore.GetCacheUserPhotos(ctx, userId)
}

func (c *localUserFacade) PutCacheUserPhotos(ctx context.Context, userId int32, photos *model.UserPhotos) error {
	return c.UserCore.PutCacheUserPhotos(ctx, userId, photos)
}

// ///////////////////////////////////////////////////////////////////////////
func (c *localUserFacade) GetPrivacy(ctx context.Context, userId int32, keyType int) ([]*mtproto.PrivacyRule, error) {
	return c.UserCore.GetPrivacy(ctx, userId, keyType)
}

func (c *localUserFacade) SetPrivacy(ctx context.Context, userId int32, keyType int, rules []*mtproto.PrivacyRule) error {
	return c.UserCore.SetPrivacy(ctx, userId, keyType, rules)
}

func (c *localUserFacade) CheckPrivacy(ctx context.Context, keyType int, selfId, peerId int32, isContact bool) bool {
	return c.UserCore.CheckPrivacy(ctx, keyType, selfId, peerId, isContact)
}

// ///////////////////////////////////////////////////////////////////////////
func (c *localUserFacade) GetUserSelf(ctx context.Context, id int32) (*mtproto.User, error) {
	return c.UserCore.GetUserSelf(ctx, id)
}

func (c *localUserFacade) GetUserById(ctx context.Context, selfUserId, userId int32) (*mtproto.User, error) {
	return c.UserCore.GetUserById(ctx, selfUserId, userId)
}

func (c *localUserFacade) GetUserListByIdList(ctx context.Context, selfUserId int32, userIdList []int32) []*mtproto.User {
	return c.UserCore.GetUserListByIdList(ctx, selfUserId, userIdList)
}

func (c *localUserFacade) GetUserByToken(ctx context.Context, token string) (*mtproto.User, error) {
	return c.UserCore.GetUserByToken(ctx, token)
}

func (c *localUserFacade) GetUserName(ctx context.Context, userId int32) string {
	return c.UserCore.GetUserName(ctx, userId)
}

func (c *localUserFacade) CheckPhoneNumberExist(ctx context.Context, phoneNumber string) (bool, error) {
	return c.UserCore.CheckPhoneNumberExist(ctx, phoneNumber)
}

func (c *localUserFacade) GetUserByPhoneNumber(ctx context.Context, selfId int32, phoneNumber string) (*mtproto.User, error) {
	return c.UserCore.GetUserByPhoneNumber(ctx, selfId, phoneNumber)
}

func (c *localUserFacade) GetUserSelfByPhoneNumber(ctx context.Context, phoneNumber string) (*mtproto.User, error) {
	return c.UserCore.GetUserSelfByPhoneNumber(ctx, phoneNumber)
}

func (c *localUserFacade) GetUserListByPhoneNumberList(ctx context.Context, selfId int32, phoneNumberList []string) []*mtproto.User {
	return c.UserCore.GetUserListByPhoneNumberList(ctx, selfId, phoneNumberList)
}

func (c *localUserFacade) GetCountryCodeByUser(ctx context.Context, userId int32) (string, error) {
	return c.UserCore.GetCountryCodeByUser(ctx, userId)
}

func (c *localUserFacade) GetUserByUsername(ctx context.Context, selfId int32, username string) (*mtproto.User, error) {
	return c.UserCore.GetUserByUsername(ctx, selfId, username)
}

func (c *localUserFacade) GetPasswordByPhone(ctx context.Context, phoneNumber string) (bool, string, error) {
	return c.UserCore.GetPasswordByPhone(ctx, phoneNumber)
}

func (c *localUserFacade) GetPhoneAndPassword(ctx context.Context, username string) (string, string, error) {
	return c.UserCore.GetPhoneAndPassword(ctx, username)
}

func (c *localUserFacade) GetPasswordById(ctx context.Context, selfId int32) (string, error) {
	return c.UserCore.GetPasswordById(ctx, selfId)
}

func (c *localUserFacade) GetPhoneById(ctx context.Context, selfId int32) (string, error) {
	return c.UserCore.GetPhoneById(ctx, selfId)
}

func (c *localUserFacade) UpdateUsername(ctx context.Context, id int32, username string) (bool, error) {
	return c.UserCore.UpdateUsername(ctx, id, username)
}

func (c *localUserFacade) UpdateUserPassword(ctx context.Context, id int32, password string) (bool, error) {
	return c.UserCore.UpdateUserPassword(ctx, id, password)
}

func (c *localUserFacade) UpdateUserInviter(ctx context.Context, id int32, inviter int32) (bool, error) {
	return c.UserCore.UpdateUserInviter(ctx, id, inviter)
}

func (c *localUserFacade) UpdateUserInfoExt(ctx context.Context, selfId int32, gender int32, birth string, country, countryCode, province, city, cityCode string) (bool, error) {
	return c.UserCore.UpdateUserInfoExt(ctx, selfId, gender, birth, country, countryCode, province, city, cityCode)
}

// ///////////////////////////////////////////////////////////////////////////
// Account
func (c *localUserFacade) CheckSessionPasswordNeeded(ctx context.Context, userId int32) (bool, error) {
	return c.UserCore.CheckSessionPasswordNeeded(ctx, userId)
}

func (c *localUserFacade) CheckRecoverCode(ctx context.Context, userId int32, code string) error {
	return c.UserCore.CheckRecoverCode(ctx, userId, code)
}

func (c *localUserFacade) GetPasswordRecovery(ctx context.Context, userId int32) (*mtproto.Auth_PasswordRecovery, error) {
	return nil, nil
}

// ///////////////////////////////////////////////////////////////////////////
// userFull
func (c *localUserFacade) GetFullUser(ctx context.Context, selfId, userId int32) (*mtproto.UserFull, error) {
	return c.UserCore.GetFullUser(ctx, selfId, userId)
}

// ///////////////////////////////////////////////////////////////////////////
// wallPaper
func (c *localUserFacade) GetWallPaperList(ctx context.Context) ([]*mtproto.WallPaper, error) {
	return c.UserCore.GetWallPaperList(ctx)
}

/////////////////////////////////////////////////////////////////////////////

func (c *localUserFacade) Report(ctx context.Context, userId, reportType, peerType, peerId, messageSenderUserId, messageId, reason int32, text string) (bool, error) {
	return c.UserCore.Report(ctx, userId, reportType, peerType, peerId, messageSenderUserId, messageId, reason, text)
}

func (c *localUserFacade) ReportIdList(ctx context.Context, userId, reportType, peerType, peerId, messageSenderUserId int32, messageIdList []int32, reason int32, text string) (bool, error) {
	return c.UserCore.ReportIdList(ctx, userId, reportType, peerType, peerId, messageSenderUserId, messageIdList, reason, text)
}

// ///////////////////////////////////////////////////////////////////////////
func (c *localUserFacade) UpdateAbout(ctx context.Context, userId int32, about string) (bool, error) {
	return c.UserCore.UpdateAbout(ctx, userId, about)
}

func (c *localUserFacade) UpdateFirstAndLastName(ctx context.Context, userId int32, firstName, lastName string) (bool, error) {
	return c.UserCore.UpdateFirstAndLastName(ctx, userId, firstName, lastName)
}

func (c *localUserFacade) UpdateVerified(ctx context.Context, userId int32, verified bool) (bool, error) {
	return c.UserCore.UpdateVerified(ctx, userId, verified)
}

func (c *localUserFacade) GetFirstAndLastName(ctx context.Context, userId int32) (string, string, error) {
	return c.UserCore.GetFirstAndLastName(ctx, userId)
}

func (c *localUserFacade) UpdateUserStatus(ctx context.Context, userId int32, lastSeenAt int64) error {
	return c.UserCore.UpdateUserStatus(ctx, userId, lastSeenAt)
}

func (c *localUserFacade) GetContactUserIdList(ctx context.Context, userId int32) []int32 {
	return c.UserCore.GetContactUserIdList(ctx, userId)
}

func (c *localUserFacade) GetMutualContactUserIdList(ctx context.Context, mutual bool, userId int32) []int32 {
	return c.UserCore.GetMutualContactUserIdList(ctx, mutual, userId)
}

func (c *localUserFacade) IsBlockedByUser(ctx context.Context, selfUserId, id int32) bool {
	return c.UserCore.IsBlockedByUser(ctx, selfUserId, id)
}

func (c *localUserFacade) CheckBlockUserList(ctx context.Context, selfUserId int32, idList []int32) []int32 {
	return c.UserCore.CheckBlockUserList(ctx, selfUserId, idList)
}

func (c *localUserFacade) GetUserStatus2(ctx context.Context, selfId, userId int32, isContact, isBlocked bool) *mtproto.UserStatus {
	return c.UserCore.GetUserStatus2(ctx, selfId, userId, isContact, isBlocked)
}

func (c *localUserFacade) GetLastSeenList(ctx context.Context, id []int32) map[int32]int64 {
	return c.UserCore.GetLastSeenList(ctx, id)
}

// ///////////////////////////////////////////////////////////////////////////
func (c *localUserFacade) CheckUserAccessHash(ctx context.Context, userId int32, hash int64) bool {
	return c.UserCore.CheckUserAccessHash(userId, hash)
}

// ///////////////////////////////////////////////////////////////////////////
func (c *localUserFacade) BlockUser(ctx context.Context, userId, blockId int32) bool {
	return c.UserCore.BlockUser(ctx, userId, blockId)
}

func (c *localUserFacade) UnBlockUser(ctx context.Context, userId, blockId int32) bool {
	return c.UserCore.UnBlockUser(ctx, userId, blockId)
}

func (c *localUserFacade) DeleteContact(ctx context.Context, userId, deleteId int32, mutual bool) bool {
	return c.UserCore.DeleteContact(ctx, userId, deleteId, mutual)
}

func (c *localUserFacade) GetContactLink(ctx context.Context, userId, contactId int32) (myLink, foreignLink *mtproto.ContactLink) {
	return c.UserCore.GetContactLink(ctx, userId, contactId)
}

func (c *localUserFacade) GetBlockedList(ctx context.Context, userId, offset, limit int32) []*mtproto.ContactBlocked {
	return c.UserCore.GetBlockedList(ctx, userId, offset, limit)
}

func (c *localUserFacade) GetContactList(ctx context.Context, userId int32) []*mtproto.Contact {
	return c.UserCore.GetContactList(ctx, userId)
}

func (c *localUserFacade) GetAllContactList(ctx context.Context, userId int32) []*mtproto.Contact {
	return c.UserCore.GetAllContactList(ctx, userId)
}

func (c *localUserFacade) SearchContacts(ctx context.Context, userId int32, q string, limit int32) ([]int32, []int32) {
	return c.UserCore.SearchContacts(ctx, userId, q, limit)
}

func (c *localUserFacade) GetContactAndMutual(ctx context.Context, selfUserId, id int32) (bool, bool) {
	return c.UserCore.GetContactAndMutual(ctx, selfUserId, id)
}

func (c *localUserFacade) BackupPhoneBooks(ctx context.Context, authKeyId int64, contacts []*mtproto.InputContact) {
	c.UserCore.BackupPhoneBooks(ctx, authKeyId, contacts)
}

func (c *localUserFacade) ImportContacts(ctx context.Context, userId int32, contacts []*mtproto.InputContact) ([]*mtproto.ImportedContact, []*mtproto.PopularContact, []int32) {
	return c.UserCore.ImportContacts(ctx, userId, contacts)
}

func (c *localUserFacade) AddContact(ctx context.Context, selfUserId, contactId int32, addPhonePrivacyException bool, contactFirstName, contactLastName string) (bool, error) {
	return c.UserCore.AddContact(ctx, selfUserId, contactId, addPhonePrivacyException, contactFirstName, contactLastName)
}

func (c *localUserFacade) CheckContact(ctx context.Context, selfUserId, id int32) bool {
	r, _ := c.UserCore.GetContactAndMutual(ctx, selfUserId, id)
	return r
}

func (c *localUserFacade) GetImportersByPhone(ctx context.Context, phone string) []*mtproto.InputContact {
	return c.UserCore.GetImportersByPhone(ctx, phone)
}

func (c *localUserFacade) DeleteImportersByPhone(ctx context.Context, phone string) {
	c.UserCore.DeleteImportersByPhone(ctx, phone)
}

func (c *localUserFacade) SearchChannelParticipants(ctx context.Context, selfId int32, channelId int32, q string) []*mtproto.User {
	return c.UserCore.SearchChannelParticipants(ctx, selfId, channelId, q)
}

func (c *localUserFacade) SetBotCommands(ctx context.Context, botId int32, commands []*mtproto.BotCommand) error {
	return c.UserCore.SetBotCommands(ctx, botId, commands)
}

func (c *localUserFacade) IsBot(ctx context.Context, id int32) bool {
	return c.UserCore.IsBot(ctx, id)
}

func (c *localUserFacade) GetMutableUsers(ctx context.Context, idList ...int32) model.MutableUsers {
	return c.UserCore.GetMutableUsers(ctx, idList...)
}

func (c *localUserFacade) CreateNewPredefinedUser(ctx context.Context, phone, firstName, lastName, username, code string, verified bool) (*mtproto.PredefinedUser, error) {
	return c.UserCore.CreateNewPredefinedUser(ctx, phone, firstName, lastName, username, code, verified)
}

func (c *localUserFacade) GetPredefinedUser(ctx context.Context, phone string) (*mtproto.PredefinedUser, error) {
	return c.UserCore.GetPredefinedUser(ctx, phone)
}

func (c *localUserFacade) GetPredefinedUserList(ctx context.Context) []*mtproto.PredefinedUser {
	return c.UserCore.GetPredefinedUserList(ctx)
}

func (c *localUserFacade) UpdatePredefinedFirstAndLastName(ctx context.Context, phone, firstName, lastName string) (*mtproto.PredefinedUser, error) {
	return c.UserCore.UpdatePredefinedFirstAndLastName(ctx, phone, firstName, lastName)
}

func (c *localUserFacade) UpdatePredefinedVerified(ctx context.Context, phone string, verified bool) (*mtproto.PredefinedUser, error) {
	return c.UserCore.UpdatePredefinedVerified(ctx, phone, verified)
}

func (c *localUserFacade) UpdatePredefinedUsername(ctx context.Context, phone, username string) (*mtproto.PredefinedUser, error) {
	return c.UserCore.UpdatePredefinedUsername(ctx, phone, username)
}

func (c *localUserFacade) UpdatePredefinedCode(ctx context.Context, phone, code string) (*mtproto.PredefinedUser, error) {
	return c.UserCore.UpdatePredefinedCode(ctx, phone, code)
}

func (c *localUserFacade) PredefinedBindRegisteredUserId(ctx context.Context, phone string, registeredUserId int32) bool {
	return c.UserCore.PredefinedBindRegisteredUserId(ctx, phone, registeredUserId)
}

func (c *localUserFacade) AddPeerSettings(ctx context.Context, selfId int32, peer *model.PeerUtil, settings *mtproto.PeerSettings) error {
	return c.UserCore.AddPeerSettings(ctx, selfId, peer, settings)
}

func (c *localUserFacade) GetPeerSettings(ctx context.Context, selfId int32, peer *model.PeerUtil) (*mtproto.PeerSettings, error) {
	return c.UserCore.GetPeerSettings(ctx, selfId, peer)
}

func (c *localUserFacade) DeletePeerSettings(ctx context.Context, selfId int32, peer *model.PeerUtil) error {
	return c.UserCore.DeletePeerSettings(ctx, selfId, peer)
}

func (c *localUserFacade) GetCustomServiceList(ctx context.Context) []int32 {
	return c.UserCore.GetCustomerServiceList(ctx)
}

func init() {
	Register("local", localUserFacadeInstance)
}
