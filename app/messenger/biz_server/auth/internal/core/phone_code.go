package core

import (
	"context"
	"time"

	"open.chat/app/messenger/biz_server/auth/internal/dal/dataobject"
	"open.chat/app/messenger/biz_server/auth/internal/model"
	"open.chat/mtproto"
	"open.chat/pkg/crypto"
	"open.chat/pkg/log"
	"open.chat/pkg/random2"
)

// for sendCode
func (c *AuthCore) CreatePhoneCode(ctx context.Context,
	authKeyId int64,
	sessionId int64,
	phoneNumber string,
	phoneNumberRegistered bool,
	sendCodeType, nextCodeType, state int) (codeData *model.PhoneCodeTransaction, err error) {

	newCodeData := func() *model.PhoneCodeTransaction {
		return &model.PhoneCodeTransaction{
			AuthKeyId:             authKeyId,
			PhoneNumber:           phoneNumber,
			SessionId:             sessionId,
			PhoneNumberRegistered: phoneNumberRegistered,
			PhoneCode:             random2.RandomNumeric(5),
			PhoneCodeHash:         crypto.GenerateStringNonce(16),
			PhoneCodeExpired:      int32(time.Now().Unix() + 3*60),
			SentCodeType:          sendCodeType,
			FlashCallPattern:      "*",
			NextCodeType:          nextCodeType,
			State:                 model.CodeStateSend, // model.CodeStateSent
		}
	}

	if codeData, err = c.Dao.GetCachePhoneCode(ctx, authKeyId, phoneNumber); err != nil {
		log.Errorf("getCachePhoneCode - error: %v", err)
		err = mtproto.ErrInternelServerError
		return
	}
	if codeData == nil {
		codeData = newCodeData()
	} else if sessionId != codeData.SessionId {
		codeData.State = model.CodeStateSend
		codeData.SessionId = sessionId
	}

	return
}

func (c *AuthCore) GetPhoneCode(ctx context.Context,
	authKeyId int64,
	phoneNumber, phoneCodeHash string) (codeData *model.PhoneCodeTransaction, err error) {

	if codeData, err = c.GetCachePhoneCode(ctx, authKeyId, phoneNumber); err != nil {
		log.Errorf("getPhoneCode by {authKeyId: %d, phoneNumber: %s} error - %v", authKeyId, phoneNumber, err)
		err = mtproto.ErrPhoneCodeExpired
		return
	} else if codeData == nil {
		log.Errorf("getPhoneCode by {authKeyId: %d, phoneNumber: %s} error - %v", authKeyId, phoneNumber, err)
		err = mtproto.ErrPhoneCodeExpired
		return
	} else if codeData.PhoneCodeHash != phoneCodeHash {
		log.Errorf("getPhoneCode by {authKeyId: %d, phoneNumber: %s, phoneCodeHash: %s} error - invalid phoneCodeHash",
			authKeyId,
			phoneNumber,
			phoneCodeHash)
		err = mtproto.ErrPhoneCodeInvalid
	}
	return
}

func (c *AuthCore) DeletePhoneCode(ctx context.Context, authKeyId int64, phoneNumber, phoneCodeHash string) error {
	return c.DeleteCachePhoneCode(ctx, authKeyId, phoneNumber)
}

func (c *AuthCore) UpdatePhoneCodeData(ctx context.Context,
	authKeyId int64,
	phoneNumber, phoneCodeHash string,
	codeData *model.PhoneCodeTransaction) error {
	return c.PutCachePhoneCode(ctx, authKeyId, phoneNumber, codeData)
}

func (c *AuthCore) CheckCanDoAction(ctx context.Context,
	authKeyId int64,
	phoneNumber string,
	actionType int) error {
	return nil
}

func (c *AuthCore) LogAuthAction(ctx context.Context,
	authKeyId, msgId int64,
	clientIp string,
	phoneNumber string,
	actionType int, log string) {

	do := &dataobject.AuthOpLogsDO{
		AuthKeyId: authKeyId,
		Ip:        clientIp,
		OpType:    int32(actionType),
		LogText:   log,
	}
	c.AuthOpLogsDAO.Insert(ctx, do)
}
