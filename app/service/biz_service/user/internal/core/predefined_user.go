package core

import (
	"context"

	"github.com/gogo/protobuf/types"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/mtproto"
	"open.chat/pkg/util"
)

func makePredefinedUser(do *dataobject.PredefinedUsersDO) *mtproto.PredefinedUser {
	pUser := mtproto.MakeTLPredefinedUser(&mtproto.PredefinedUser{
		Phone:            do.Phone,
		FirstName:        &types.StringValue{Value: do.FirstName},
		LastName:         nil,
		Username:         nil,
		Code:             do.Code,
		Verified:         util.Int8ToBool(do.Verified),
		RegisteredUserId: nil,
		Banned:           false,
	}).To_PredefinedUser()

	if do.LastName != "" {
		pUser.LastName = &types.StringValue{Value: do.LastName}
	}

	if do.Username != "" {
		pUser.Username = &types.StringValue{Value: do.Username}
	}

	if do.RegisteredUserId != 0 {
		pUser.RegisteredUserId = &types.Int32Value{Value: do.RegisteredUserId}
	}

	return pUser
}

func (m *UserCore) CreateNewPredefinedUser(ctx context.Context, phoneNumber, firstName, lastName, username, code string, verified bool) (*mtproto.PredefinedUser, error) {
	do := &dataobject.PredefinedUsersDO{
		Phone:     phoneNumber,
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
		Code:      code,
		Verified:  util.BoolToInt8(verified),
	}

	if _, _, err := m.PredefinedUsersDAO.Insert(ctx, do); err != nil {
		err = mtproto.ErrPhoneNumberOccupied
		return nil, err
	}

	return makePredefinedUser(do), nil
}

func (m *UserCore) GetPredefinedUser(ctx context.Context, phone string) (*mtproto.PredefinedUser, error) {
	do, err := m.PredefinedUsersDAO.SelectByPhone(ctx, phone)
	if do == nil {
		err = mtproto.ErrPhoneNumberUnoccupied
		return nil, err
	}

	return makePredefinedUser(do), nil
}

func (m *UserCore) GetPredefinedUserList(ctx context.Context) []*mtproto.PredefinedUser {
	doList, _ := m.PredefinedUsersDAO.SelectPredefinedUsersAll(ctx)

	users := make([]*mtproto.PredefinedUser, 0, len(doList))
	for i := 0; i < len(doList); i++ {
		users = append(users, makePredefinedUser(&doList[i]))
	}

	return users
}

func (m *UserCore) UpdatePredefinedFirstAndLastName(ctx context.Context, phone, firstName, lastName string) (*mtproto.PredefinedUser, error) {
	do, err := m.PredefinedUsersDAO.SelectByPhone(ctx, phone)
	if do == nil {
		err = mtproto.ErrUserIdInvalid
		return nil, err
	}
	do.FirstName = firstName
	do.LastName = lastName

	if _, err = m.PredefinedUsersDAO.Update(ctx, map[string]interface{}{
		"first_name": do.FirstName,
		"last_name":  do.LastName,
	}, phone); err != nil {
		return nil, err
	}

	return makePredefinedUser(do), nil
}

func (m *UserCore) UpdatePredefinedVerified(ctx context.Context, phone string, verified bool) (*mtproto.PredefinedUser, error) {
	do, err := m.PredefinedUsersDAO.SelectByPhone(ctx, phone)
	if do == nil {
		err = mtproto.ErrUserIdInvalid
		return nil, err
	}
	do.Verified = util.BoolToInt8(verified)

	if _, err = m.PredefinedUsersDAO.Update(ctx, map[string]interface{}{
		"verified": do.Verified,
	}, phone); err != nil {
		return nil, err
	}

	return makePredefinedUser(do), nil
}

func (m *UserCore) UpdatePredefinedUsername(ctx context.Context, phone, username string) (*mtproto.PredefinedUser, error) {
	do, err := m.PredefinedUsersDAO.SelectByPhone(ctx, phone)
	if do == nil {
		err = mtproto.ErrUserIdInvalid
		return nil, err
	}
	do.Username = username

	if _, err = m.PredefinedUsersDAO.Update(ctx, map[string]interface{}{
		"username": do.Username,
	}, phone); err != nil {
		return nil, err
	}

	return makePredefinedUser(do), nil
}

func (m *UserCore) UpdatePredefinedCode(ctx context.Context, phone, code string) (*mtproto.PredefinedUser, error) {
	do, err := m.PredefinedUsersDAO.SelectByPhone(ctx, phone)
	if do == nil {
		err = mtproto.ErrUserIdInvalid
		return nil, err
	}
	do.Code = code

	if _, err = m.PredefinedUsersDAO.Update(ctx, map[string]interface{}{
		"code": do.Code,
	}, phone); err != nil {
		return nil, err
	}

	return makePredefinedUser(do), nil
}

func (m *UserCore) PredefinedBindRegisteredUserId(ctx context.Context, phone string, registeredUserId int32) bool {
	m.PredefinedUsersDAO.Update(ctx, map[string]interface{}{
		"registered_user_id": registeredUserId,
	}, phone)

	return true
}
