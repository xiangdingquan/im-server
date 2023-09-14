package service

import (
	"context"

	"open.chat/app/admin/api_server/api"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

func (s *Service) SendVerifyCode(ctx context.Context, i *api.SendVerifyCode) (r *api.CodeResult, err error) {
	log.Debugf("sendVerifyCode - request: %s", i.DebugString())
	r = &api.CodeResult{
		Code: i.Code,
	}
	log.Debugf("sendVerifyCode - reply: %s", r.DebugString())

	return
}

func (s *Service) VerifyCode(ctx context.Context, i *api.VerifyCode) (r *api.VoidRsp, err error) {
	log.Debugf("verifyCode - request: %s", i.DebugString())

	var (
		result mtproto.TLObject
		rUser  *mtproto.PredefinedUser
		ok     bool
	)

	req := &mtproto.TLUsersGetPredefinedUser{
		Phone: i.ExtraData,
	}

	result, err = s.Invoke(ctx, 777000, req)
	if err != nil {
		log.Errorf("verifyCode - error: %v", err)
		return
	} else if rUser, ok = result.(*mtproto.PredefinedUser); !ok {
		log.Errorf("verifyCode - error: invalid Vector_PredefinedUser type")
		err = mtproto.ErrInternelServerError
		return
	}

	if rUser.GetCode() != i.Code {
		err = mtproto.ErrPhoneCodeInvalid
		log.Errorf("verifyCode - error: %v", err)
		return
	}

	r = api.GVoidRsp
	log.Debugf("verifyCode - reply: %s", r.DebugString())
	return
}
