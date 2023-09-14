package core

import (
	"context"

	"open.chat/pkg/util"
)

func (m *UserCore) UpdateFirstAndLastName(ctx context.Context, id int32, firstName, lastName string) (bool, error) {
	rowsAffected, err := m.UsersDAO.UpdateUser(ctx, map[string]interface{}{
		"first_name": firstName,
		"last_name":  lastName,
	}, id)
	return rowsAffected != 0, err
}

func (m *UserCore) UpdateAbout(ctx context.Context, id int32, about string) (bool, error) {
	rowsAffected, err := m.UsersDAO.UpdateUser(ctx, map[string]interface{}{
		"about": about,
	}, id)
	return rowsAffected != 0, err
}

func (m *UserCore) GetFirstAndLastName(ctx context.Context, id int32) (string, string, error) {
	do, err := m.UsersDAO.SelectById(ctx, id)
	if err != nil {
		return "", "", err
	} else {
		if do == nil {
			return "", "", err
		} else {
			return do.FirstName, do.LastName, nil
		}
	}
}

func (m *UserCore) UpdateUsername(ctx context.Context, id int32, username string) (bool, error) {
	rowsAffected, err := m.UsersDAO.UpdateUser(ctx, map[string]interface{}{
		"username": username,
	}, id)
	return rowsAffected != 0, err
}

func (m *UserCore) UpdateUserPassword(ctx context.Context, id int32, password string) (bool, error) {
	raw, md5 := m.formatPassword(password)
	rowsAffected, err := m.UsersDAO.UpdateUser(ctx, map[string]interface{}{
		"password":     md5,
		"raw_password": raw,
	}, id)
	return rowsAffected != 0, err
}

func (m *UserCore) UpdateUserInviter(ctx context.Context, id int32, inviter int32) (bool, error) {
	do, err := m.UsersDAO.SelectById(ctx, id)
	if do == nil || err != nil {
		return false, err
	}
	rowsAffected, err := m.UsersDAO.UpdateUser(ctx, map[string]interface{}{
		"inviter_uid": inviter,
		"channel_id":  do.ChannelId,
	}, id)
	return rowsAffected != 0, err
}

func (m *UserCore) UpdateVerified(ctx context.Context, id int32, verified bool) (bool, error) {
	rowsAffected, err := m.UsersDAO.UpdateUser(ctx, map[string]interface{}{
		"verified": util.BoolToInt8(verified),
	}, id)
	return rowsAffected != 0, err
}

func (m *UserCore) UpdateUserInfoExt(ctx context.Context, id int32, gender int32, birth string, country, countryCode, province, city, cityCode string) (bool, error) {
	rowsAffected, err := m.UsersDAO.UpdateUser(ctx, map[string]interface{}{
		"gender":       gender,
		"birth":        birth,
		"country":      country,
		"country_code": countryCode,
		"province":     province,
		"city":         city,
		"city_code":    cityCode,
	}, id)
	return rowsAffected != 0, err
}
