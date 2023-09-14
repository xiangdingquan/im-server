package helper

import (
	"context"
	"errors"
	"fmt"
	"open.chat/model"

	"open.chat/app/infra/databus/pkg/cache/redis"
	"open.chat/app/json/consts"
	"open.chat/app/pkg/redis_util"
	"open.chat/app/smscode"
	"open.chat/pkg/log"
	"open.chat/pkg/random2"
)

type (
	Sms struct {
		redis *redis_util.Redis
		sms   smscode.VerifyCodeInterface
	}
)

const (
	sendCodeTimeout     = 300 // salt timeout
	cacheSendCodePrefix = "Sms_Send_Code"
)

func genCacheKey(smsType consts.SmsCodeType, phoneNumber string) string {
	return fmt.Sprintf("%s_%s_%d", cacheSendCodePrefix, phoneNumber, smsType)
}

var s *Sms

func (d *Sms) PutCacheSmsCode(ctx context.Context, smsType consts.SmsCodeType, phoneNumber string, code string) (err error) {
	cacheKey := genCacheKey(smsType, phoneNumber)
	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()
	if _, err = conn.Do("SETEX", cacheKey, sendCodeTimeout, code); err != nil {
		log.Errorf("conn.SETEX(%s) error(%v)", cacheKey, err)
	}
	return
}

func (d *Sms) GetCacheSmsCode(ctx context.Context, smsType consts.SmsCodeType, phoneNumber string) (code string, err error) {
	cacheKey := genCacheKey(smsType, phoneNumber)
	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()
	v, err := redis.String(conn.Do("GET", cacheKey))
	if err != nil {
		if err != redis.ErrNil {
			log.Errorf("conn.Do(GET %s) error(%v)", cacheKey, err)
		} else {
			err = nil
		}
	} else {
		code = v
	}
	return
}

func (d *Sms) DelCacheSmsCode(ctx context.Context, smsType consts.SmsCodeType, phoneNumber string) (err error) {
	cacheKey := genCacheKey(smsType, phoneNumber)
	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()
	if _, err = conn.Do("DEL", cacheKey); err != nil {
		log.Errorf("conn.DEL(%s) error(%v)", cacheKey, err)
	}
	return
}

func GetCacheCode(ctx context.Context, smsType consts.SmsCodeType, phoneNumber string) (string, error) {
	if s == nil {
		s = &Sms{
			redis: redis_util.GetSingletonRedis(),
			sms:   smscode.New(nil),
		}
	}
	code, err := s.GetCacheSmsCode(ctx, smsType, phoneNumber)
	if err != nil {
		return "", err
	}
	return code, nil
}

// SendVerifyCode .
func SendVerifyCode(ctx context.Context, smsType consts.SmsCodeType, phoneNumber string, langType model.LangType) (string, error) {
	if s == nil {
		s = &Sms{
			redis: redis_util.GetSingletonRedis(),
			sms:   smscode.New(nil),
		}
	}

	code := random2.RandomNumeric(5)
	code, err := s.sms.SendSmsVerifyCode(ctx, phoneNumber, code, "", langType)
	if err != nil {
		return "", errors.New("send code fail")
	}

	err = s.PutCacheSmsCode(ctx, smsType, phoneNumber, code)
	if err != nil {
		return "", errors.New("verify code save fail")
	}

	return code, nil
}

// VerifyCode .
func VerifyCode(ctx context.Context, smsType consts.SmsCodeType, phoneNumber string, code string) error {
	if s == nil {
		s = &Sms{
			redis: redis_util.GetSingletonRedis(),
			sms:   smscode.New(nil),
		}
	}

	old, err := GetCacheCode(ctx, smsType, phoneNumber)
	if err != nil {
		return err
	}

	if code != old && s.sms.VerifySmsCode(ctx, "", code, old) != nil {
		return errors.New("code is invalid")
	}
	s.DelCacheSmsCode(ctx, smsType, phoneNumber)
	return nil
}
