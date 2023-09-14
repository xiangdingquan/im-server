package core

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/database/sqlx"
)

func (m *UserCore) CheckPhoneNumberExist(ctx context.Context, phoneNumber string) (bool, error) {
	return sqlx.CheckExists(ctx, m.Dao.DB, "users", map[string]interface{}{
		"phone": phoneNumber,
	})
}

func (m *UserCore) GetUserByPhoneNumber(ctx context.Context, selfId int32, phoneNumber string) (*mtproto.User, error) {
	do, err := m.UsersDAO.SelectByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		return nil, err
	} else if do == nil {
		return nil, nil
	}

	users := m.getMutableUsersBySelfId(ctx, selfId, do.Id)
	return users.ToUnsafeUser(selfId, do.Id)
}

func (m *UserCore) GetUserSelfByPhoneNumber(ctx context.Context, phoneNumber string) (*mtproto.User, error) {
	do, err := m.UsersDAO.SelectByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		return nil, err
	} else if do == nil {
		return nil, nil
	}

	users := m.GetMutableUsers(ctx, do.Id)
	return users.ToUnsafeUserSelf(do.Id)
}

func (m *UserCore) GetUserListByPhoneNumberList(ctx context.Context, selfId int32, phoneNumberList []string) []*mtproto.User {
	usersDOList, _ := m.UsersDAO.SelectUsersByPhoneList(ctx, phoneNumberList)
	if len(usersDOList) == 0 {
		return []*mtproto.User{}
	}

	idList := make([]int32, 0, len(usersDOList)+1)
	idList = append(idList, selfId)
	for i := 0; i < len(usersDOList); i++ {
		idList = append(idList, usersDOList[i].Id)
	}

	return m.GetUserListByIdList(ctx, selfId, idList)
}

func (m *UserCore) GetUserSelf(ctx context.Context, id int32) (*mtproto.User, error) {
	users := m.GetMutableUsers(ctx, id)
	return users.ToUnsafeUserSelf(id)
}

func (m *UserCore) GetUserById(ctx context.Context, selfUserId, userId int32) (*mtproto.User, error) {
	users := m.getMutableUsersBySelfId(ctx, selfUserId, userId)
	return users.ToUnsafeUser(selfUserId, userId)
}

func (m *UserCore) GetUserListByIdList(ctx context.Context, selfUserId int32, userIdList []int32) (users []*mtproto.User) {
	if len(userIdList) == 0 {
		users = []*mtproto.User{}
		return
	}

	mutableUsers := m.getMutableUsersBySelfId(ctx, selfUserId, userIdList...)
	return mutableUsers.GetUsersByIdList(selfUserId, userIdList)
}

func (m *UserCore) GetUserByToken(ctx context.Context, token string) (*mtproto.User, error) {
	botId, err := m.BotsDAO.SelectByToken(ctx, token)
	if err != nil || botId == 0 {
		err = mtproto.ErrTokenInvalid
		return nil, err
	}

	users := m.GetMutableUsers(ctx, botId)
	return users.ToUnsafeUserSelf(botId)
}

func (m *UserCore) GetCountryCodeByUser(ctx context.Context, userId int32) (string, error) {
	if do, err := m.UsersDAO.SelectCountryCode(ctx, userId); err != nil {
		return "", err
	} else if do == nil {
		return "", nil
	} else {
		return do.CountryCode, nil
	}
}

func (m *UserCore) GetUserByUsername(ctx context.Context, selfId int32, username string) (*mtproto.User, error) {
	if do, err := m.UsersDAO.SelectByUsername(ctx, username); err != nil {
		return nil, err
	} else if do == nil {
		return nil, nil
	} else {
		return m.GetUserById(ctx, selfId, do.Id)
	}
}

func (m *UserCore) GetPasswordByPhone(ctx context.Context, phoneNumber string) (bool, string, error) {
	if do, err := m.UsersDAO.SelectByPhoneNumber(ctx, phoneNumber); do == nil || err != nil {
		return false, "", err
	} else {
		uDo, _ := m.UsersDAO.SelectById(ctx, do.Id)
		if uDo == nil || uDo.Deleted == 1 {
			return false, "", err
		}
		return true, uDo.Password, nil
	}
}

func (m *UserCore) GetPhoneAndPassword(ctx context.Context, username string) (string, string, error) {
	if do, err := m.UsersDAO.SelectByUsername(ctx, username); do == nil || err != nil {
		return "", "", err
	} else {
		uDo, _ := m.UsersDAO.SelectById(ctx, do.Id)
		if uDo == nil || uDo.Deleted == 1 {
			return "", "", err
		}
		return uDo.Phone, uDo.Password, nil
	}
}

func (m *UserCore) GetPasswordById(ctx context.Context, selfId int32) (string, error) {
	uDo, err := m.UsersDAO.SelectById(ctx, selfId)
	if err != nil {
		return "", err
	}

	if uDo == nil || uDo.Deleted == 1 {
		return "", err
	}

	return uDo.Password, nil
}

func (m *UserCore) GetPhoneById(ctx context.Context, selfId int32) (string, error) {
	uDo, err := m.UsersDAO.SelectById(ctx, selfId)
	if err != nil {
		return "", err
	}

	if uDo == nil || uDo.Deleted == 1 {
		return "", err
	}

	return uDo.Phone, nil
}

func (m *UserCore) DeleteUser(ctx context.Context, userId int32, reason string) (bool, error) {
	affected, err := m.UsersDAO.Delete(ctx, reason, userId)
	return affected == 1, err
}

func (m *UserCore) CheckUserAccessHash(userId int32, accessHash int64) bool {
	return true
}

func (m *UserCore) SearchChannelParticipants(ctx context.Context, selfId int32, channelId int32, q string) []*mtproto.User {
	q = "%" + q + "%"
	doList, _ := m.UsersDAO.QueryChannelParticipants(ctx, channelId, q, q, q)
	idList := make([]int32, 0, len(doList))
	for i := 0; i < len(doList); i++ {
		idList = append(idList, doList[i].Id)
	}

	return m.GetUserListByIdList(ctx, selfId, idList)
}

func (m *UserCore) GetUserName(ctx context.Context, userId int32) string {
	uDO, _ := m.UsersDAO.SelectById(ctx, userId)
	if uDO == nil || uDO.Deleted == 1 {
		return "Deleted Account"
	}

	if uDO.FirstName == "" && uDO.LastName == "" {
		return ""
	} else if uDO.FirstName == "" {
		return uDO.LastName
	} else if uDO.LastName == "" {
		return uDO.FirstName
	} else {
		return uDO.FirstName + " " + uDO.LastName
	}
}

func (m *UserCore) GetUserPassword(ctx context.Context, userId int32) string {
	uDO, _ := m.UsersDAO.SelectById(ctx, userId)
	if uDO == nil || uDO.Deleted == 1 {
		return "Deleted_Account"
	}

	return uDO.Password
}

func (m *UserCore) IsInternal(ctx context.Context, userId int32) bool {
	uDO, _ := m.UsersDAO.SelectById(ctx, userId)
	if uDO == nil || uDO.Deleted == 1 {
		return false
	}
	return uDO.IsInternal != 0
}

func (m *UserCore) GetCustomerServiceList(ctx context.Context) []int32 {
	doList, _ := m.UsersDAO.SelectCustomerServiceList(ctx)
	idList := make([]int32, 0, len(doList))
	for i := 0; i < len(doList); i++ {
		idList = append(idList, doList[i].Id)
	}
	return idList
}
