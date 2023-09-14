package none

import (
	"context"
	"open.chat/model"

	"open.chat/app/pkg/code"
	"open.chat/mtproto"
)

func New(c *code.SmsVerifyCodeConfig) *noneVerifyCode {
	return &noneVerifyCode{
		code: c,
	}
}

type noneVerifyCode struct {
	code *code.SmsVerifyCodeConfig
}

func (m *noneVerifyCode) SendSmsVerifyCode(ctx context.Context, phoneNumber, code, codeHash string, langCode model.LangType) (string, error) {
	return code, nil
}

func (m *noneVerifyCode) VerifySmsCode(ctx context.Context, codeHash, code, extraData string) error {
	if code != "54321" {
		return mtproto.ErrPhoneCodeInvalid
	}
	return nil
}
