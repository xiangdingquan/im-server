package logic

import (
	"context"
	"time"

	"open.chat/app/pkg/code"

	"open.chat/app/messenger/biz_server/auth/internal/core"
	"open.chat/app/messenger/biz_server/auth/internal/model"
	"open.chat/app/smscode"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

type AuthLogic struct {
	*core.AuthCore
	smscode.VerifyCodeInterface
}

func NewAuthSignLogic(core *core.AuthCore, code *code.SmsVerifyCodeConfig) *AuthLogic {
	return &AuthLogic{
		AuthCore:            core,
		VerifyCodeInterface: smscode.New(code),
	}
}

func (m *AuthLogic) DoAuthSendCode(
	ctx context.Context,
	authKeyId int64,
	sessionId int64,
	phoneNumber string,
	phoneRegistered,
	allowFlashCall,
	currentNumber bool,
	apiId int32,
	apiHash string,
	cb func(codeData *model.PhoneCodeTransaction) error) (codeData *model.PhoneCodeTransaction, err error) {

	sentCodeType, nextCodeType := model.MakeCodeType(phoneRegistered, allowFlashCall, currentNumber)
	if codeData, err = m.AuthCore.CreatePhoneCode(ctx,
		authKeyId,
		sessionId,
		phoneNumber,
		phoneRegistered,
		sentCodeType,
		nextCodeType,
		model.CodeStateSend); err != nil {
		return
	}

	if cb != nil {
		if err = cb(codeData); err != nil {
			return
		}
	}

	// after sendSms success, save codeData
	m.AuthCore.UpdatePhoneCodeData(ctx, authKeyId, phoneNumber, codeData.PhoneCodeHash, codeData)

	return
}

func (m *AuthLogic) DoAuthReSendCode(ctx context.Context,
	authKeyId int64,
	phoneNumber, phoneCodeHash string,
	cb func(codeData *model.PhoneCodeTransaction) error) (codeData *model.PhoneCodeTransaction, err error) {
	if codeData, err = m.AuthCore.GetPhoneCode(ctx, authKeyId, phoneNumber, phoneCodeHash); err != nil {
		return
	}

	if codeData.State != model.CodeStateSent && codeData.State != model.CodeStateSignIn {
		err = mtproto.ErrInternelServerError
		return
	}

	now := int32(time.Now().Unix())
	if now > codeData.PhoneCodeExpired {
		err = mtproto.ErrPhoneCodeExpired
		return
	}

	if cb != nil {
		err = cb(codeData)
		if err != nil {
			return
		}
	}

	m.AuthCore.UpdatePhoneCodeData(context.Background(), authKeyId, phoneNumber, codeData.PhoneCodeHash, codeData)

	return
}

func (m *AuthLogic) DoAuthCancelCode(ctx context.Context, authKeyId int64, phoneNumber, phoneCodeHash string) error {
	return m.AuthCore.DeletePhoneCode(ctx, authKeyId, phoneNumber, phoneCodeHash)
}

func (m *AuthLogic) DoAuthSignIn(ctx context.Context,
	authKeyId int64,
	phoneNumber,
	phoneCode,
	phoneCodeHash string,
	cb func(codeData *model.PhoneCodeTransaction) error) (codeData *model.PhoneCodeTransaction, err error) {
	if codeData, err = m.AuthCore.GetPhoneCode(ctx, authKeyId, phoneNumber, phoneCodeHash); err != nil {
		return
	}

	if codeData.State != model.CodeStateSent && codeData.State != model.CodeStateSignIn {
		log.Errorf("error - invalid codeData state: %v", codeData)
		err = mtproto.ErrInternelServerError
		return
	}

	now := int32(time.Now().Unix())
	if now > codeData.PhoneCodeExpired {
		err = mtproto.ErrPhoneCodeExpired
		return
	}

	if cb != nil {
		if err = cb(codeData); err != nil {
			err = mtproto.ErrPhoneCodeInvalid
			return
		}
	}

	if codeData.PhoneNumberRegistered {
		codeData.State = model.CodeStateOk
	} else {
		codeData.State = model.CodeStateSignIn
	}

	err = m.AuthCore.UpdatePhoneCodeData(ctx, authKeyId, phoneNumber, phoneCodeHash, codeData)
	return
}

func (m *AuthLogic) DoAuthSignUp(ctx context.Context, authKeyId int64, phoneNumber string, phoneCode *string, phoneCodeHash string) (codeData *model.PhoneCodeTransaction, err error) {
	if codeData, err = m.AuthCore.GetPhoneCode(ctx, authKeyId, phoneNumber, phoneCodeHash); err != nil {
		return
	}
	if codeData.State != model.CodeStateSignIn && codeData.State != model.CodeStateDeleted && codeData.State != model.CodeStateSignUp {
		err = mtproto.ErrInputRequestInvalid
		log.Errorf("invalid code state(%d) - err: %v", codeData.State, err)
		return
	}

	now := int32(time.Now().Unix())
	if now > codeData.PhoneCodeExpired {
		err = mtproto.ErrPhoneCodeExpired
		return
	}

	if phoneCode != nil && *phoneCode != codeData.PhoneCodeExtraData {
		if err = m.VerifyCodeInterface.VerifySmsCode(ctx,
			codeData.PhoneCodeHash,
			*phoneCode,
			codeData.PhoneCodeExtraData); err != nil {
			return
		}
	}

	codeData.State = model.CodeStateOk
	return
}
