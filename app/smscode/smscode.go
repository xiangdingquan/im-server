package smscode

import (
	"context"
	"open.chat/model"

	"open.chat/app/pkg/code"
	"open.chat/app/smscode/none"
	"open.chat/app/smscode/smsbao"
)

type VerifyCodeInterface interface {
	SendSmsVerifyCode(ctx context.Context, phoneNumber, code, codeHash string, langType model.LangType) (string, error)
	VerifySmsCode(ctx context.Context, codeHash, code, extraData string) error
}

var g_code VerifyCodeInterface = nil

func New(c *code.SmsVerifyCodeConfig) VerifyCodeInterface {
	if g_code == nil {
		g_code = newVerifyCode(c)
	}
	return g_code
}

func newVerifyCode(c *code.SmsVerifyCodeConfig) VerifyCodeInterface {
	if c == nil {
		c = new(code.SmsVerifyCodeConfig)
	}

	switch c.Name {
	case "smsbao":
		return smsbao.New(c)
	case "none":
		return none.New(c)
	}
	return none.New(c)
}
