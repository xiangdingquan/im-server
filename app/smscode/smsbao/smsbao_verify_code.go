package smsbao

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/url"
	"open.chat/model"
	"reflect"
	"regexp"
	"strings"

	"open.chat/app/pkg/code"
	"open.chat/app/sysconfig"
	"open.chat/mtproto"
	"open.chat/pkg/http_client"
	"open.chat/pkg/log"
)

const (
	smsHost = "https://api.smsbao.com/"
)

type smsBaoVerifyCode struct {
	*code.SmsVerifyCodeConfig
}

// New ...
func New(c *code.SmsVerifyCodeConfig) *smsBaoVerifyCode {
	return &smsBaoVerifyCode{
		SmsVerifyCodeConfig: c,
	}
}

// InArray 查找字符是否在数组中 ...
func (m *smsBaoVerifyCode) InArray(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}

	return false
}

func (m *smsBaoVerifyCode) GetFilterNumbers() []string {
	noSmsMobileNumbers := []string{
		"14733331111",
		"14733332222",
		"14733333333",
		"14733334444",
		"14733335555",
		"13123456788",
		"13123456789",
	}
	return noSmsMobileNumbers
}

func (m *smsBaoVerifyCode) SendSmsVerifyCode(ctx context.Context, phoneNumber, code, codeHash string, langType model.LangType) (string, error) {
	finalAPIURL := ""
	finalPhone := ""
	pat := "^(\\+?0?86[\\- ]?)?1[3-9]\\d{9}$"
	if ok, _ := regexp.MatchString(pat, phoneNumber); ok {
		log.Infof("send_sms [%s] via china", phoneNumber)
		finalAPIURL = smsHost + "sms?"
		length := len(phoneNumber)
		finalPhone = phoneNumber[length-11 : length]
		filterList := m.GetFilterNumbers()
		for _, number := range filterList {
			if strings.HasSuffix(finalPhone, number) {
				return "54321", nil
			}
		}
	} else { //国际短信
		log.Infof("send_sms [%s] via oversea", phoneNumber)
		finalAPIURL = smsHost + "wsms?"
		finalPhone = url.QueryEscape(phoneNumber)
	}

	statusStr := map[string]string{
		"0":  "短信发送成功",
		"-1": "参数不全",
		"-2": "服务器空间不支持,请确认支持curl或者fsocket，联系您的空间商解决或者更换空间！",
		"30": "密码错误",
		"40": "账号不存在",
		"41": "余额不足",
		"42": "帐户已过期",
		"43": "IP地址限制",
		"50": "内容含有敏感词",
	}
	pass := md5.Sum([]byte(m.Password))
	passMd5 := hex.EncodeToString(pass[:])
	content := url.QueryEscape(m.getRawContent(ctx, code, langType))
	sendurl := fmt.Sprintf("%su=%s&p=%s&m=%s&c=%s", finalAPIURL, m.UserName, passMd5, finalPhone, content)
	log.Infof("%s", sendurl)
	req := http_client.Get(sendurl)
	str, err := req.String()
	log.Infof("%s[%s]", str, statusStr[str])
	return code, err
}

func (m *smsBaoVerifyCode) VerifySmsCode(ctx context.Context, codeHash, code, extraData string) error {
	if len(code) != 5 {
		log.Errorf("verifySmsCode - len(code) != 5")
		return mtproto.ErrPhoneCodeInvalid
	}

	if code != extraData {
		verifyCode := sysconfig.GetConfig2String(ctx, sysconfig.ConfigKeysCommonSmsCode, "", 0)
		if len(verifyCode) == 0 || code != verifyCode {
			log.Errorf("verifySmsCode - code invalid")
			return mtproto.ErrPhoneCodeInvalid
		}
	}

	return nil

}

func (m *smsBaoVerifyCode) getRawContent(ctx context.Context, smsCode string, langType model.LangType) string {
	log.Debugf("getRawContent - smsCode: %s, langCode: %d", smsCode, langType)
	return model.LocalizationWords{
		model.LocalizationCN:      "【" + m.Signed + "】您的验证码为：" + smsCode,
		model.LocalizationEN:      "[" + m.Signed + "] " + smsCode + " is your login code.",
		model.LocalizationDefault: "[" + m.Signed + "] " + smsCode + " is your login code.",
	}[langType]
}
