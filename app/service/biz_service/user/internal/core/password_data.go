package core

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

const (
	kStatePasswordNone             = 0
	kStateNoRecoveryPassword       = 1
	kStateEmailUnconfirmedPassword = 2
	kStateConfirmedPassword        = 3
)

func (m *UserCore) CheckRecoverCode(ctx context.Context, userId int32, code string) error {
	if do, err := m.UserPasswordsDAO.SelectCode(ctx, userId); err != nil {
		return err
	} else if do == nil {
		err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_INTERNAL_SERVER_ERROR)
		log.Errorf("%v: not found user_password row, user_id - %d", err, userId)
		return err
	} else {
		if do.Code != code {
			err := mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_CODE_INVALID)
			log.Errorf("%s: userId - %d, code - %s", err, userId, code)
			return err
		}
	}

	return nil
}

func (m *UserCore) CheckSessionPasswordNeeded(ctx context.Context, userId int32) (bool, error) {
	if do, err := m.UserPasswordsDAO.SelectByUserId(ctx, userId); err != nil {
		return false, err
	} else if do != nil {
		return do.State == kStateNoRecoveryPassword || do.State == kStateConfirmedPassword, nil
	}

	return false, nil
}

func (m *UserCore) formatPassword(password string) (rawPassword, md5Password string) {
	if len(password) == 32 {
		rawPassword = ""
		md5Password = password
	} else {
		pass := md5.Sum([]byte(password))
		md5Password = hex.EncodeToString(pass[:])
		rawPassword = password
	}

	log.Debugf("formatPassword: rawPassword - %s, md5Password - %s", rawPassword, md5Password)
	return
}
