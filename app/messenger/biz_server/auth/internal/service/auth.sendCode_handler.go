package service

import (
	"context"
	"encoding/json"
	"open.chat/app/messenger/biz_server/auth/internal/logic"
	"open.chat/app/messenger/biz_server/auth/internal/model"
	idgen "open.chat/app/service/idgen/client"
	status_client "open.chat/app/service/status/client"
	"open.chat/app/sysconfig"
	model2 "open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) authSendCode(ctx context.Context, authKeyId, sessionId int64, ipAddr string, request *mtproto.TLAuthSendCode) (reply *mtproto.Auth_SentCode, err error) {
	switch request.Constructor {
	case mtproto.CRC32_auth_sendCode_ccfd70cf:
		request.Settings = &mtproto.CodeSettings{
			AllowFlashcall: request.AllowFlashcall,
			CurrentNumber:  mtproto.FromBool(request.CurrentNumber),
		}
	case mtproto.CRC32_auth_sendCode_a677244f:
	case mtproto.CRC32_auth_sendCode_86aef0ec:
		request.Settings = &mtproto.CodeSettings{
			AllowFlashcall: request.AllowFlashcall,
			CurrentNumber:  mtproto.FromBool(request.CurrentNumber),
		}
	default:
		err = mtproto.ErrTypeConstructorInvalid
		log.Errorf("invalid constructor request: %s", request.DebugString())
		return
	}

	if err = s.AuthCore.CheckApiIdAndHash(request.ApiId, request.ApiHash); err != nil {
		log.Errorf("invalid api: {api_id: %d, api_hash: %s}", request.ApiId, request.ApiHash)
		return
	}

	req := sendCodeReq{}
	err = json.Unmarshal([]byte(request.PhoneNumber), &req)
	var phoneNumber string
	useJson := err == nil
	needPhoneCode := sysconfig.GetConfig2Bool(ctx, sysconfig.ConfigKeysRegisterNeedPhoneCode, false, 0)
	if !useJson || needPhoneCode || req.IsSignup || len(req.PhoneNumber) != 0 {
		if useJson {
			if req.IsSignup {
				if !needPhoneCode && req.PhoneNumber == "" {
					req.PhoneNumber, _ = s.AuthCore.GetCachePhoneNumber(ctx, authKeyId)
					if req.PhoneNumber == "" {
						req.PhoneNumber, _ = idgen.GetNextPhoneNumber(ctx, "86100")
						s.AuthCore.PutCachePhoneNumber(ctx, authKeyId, req.PhoneNumber)
					}
				}
			} else {
				if len(req.PhoneNumber) == 0 && len(req.UserName) == 0 {
					err = mtproto.ErrUsernameInvalid
					return
				}
			}
			request.PhoneNumber = req.PhoneNumber
		}
		phoneNumber, err = checkPhoneNumberInvalid(request.PhoneNumber)
		if phoneNumber == "" || err != nil {
			if !useJson || req.IsSignup {
				log.Errorf("check phone_number(%s) error - %v", request.PhoneNumber, err)
				err = mtproto.ErrPhoneNumberInvalid
				return
			}
		}
	}

	if useJson {
		if len(req.Password) == 0 {
			log.Errorf("need input password")
			err = mtproto.ErrUserPasswordNeeded
			return
		}

		var userPassword string
		if len(phoneNumber) > 0 {
			var ok bool
			ok, userPassword, err = s.UserFacade.GetPasswordByPhone(ctx, phoneNumber)
			if req.IsSignup {
				if ok {
					err = mtproto.ErrPhoneNumberOccupied
					return
				}
			} else {
				if !ok {
					log.Errorf("get user by phone number fail - %s => %v", phoneNumber, err)
					err = mtproto.ErrPhoneNumberUnoccupied
					return
				}
			}
		}

		if len(req.UserName) > 0 {
			if !model2.CheckUsernameInvalid(req.UserName) {
				log.Errorf("user name error - %s", req.UserName)
				err = mtproto.ErrUsernameInvalid
				return
			}
			var number string
			number, userPassword, err = s.UserFacade.GetPhoneAndPassword(ctx, req.UserName)
			if req.IsSignup {
				if number != "" {
					err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USERNAME_OCCUPIED)
					return
				}
			} else {
				if number == "" {
					log.Errorf("check user name fail - %s => %v", req.UserName, err)
					err = mtproto.ErrUserNameNotExist
					return
				}
				phoneNumber = number
			}
		}

		if req.IsSignup {
			err = s.checkIpLimit(ctx, ipAddr)
			if err != nil {
				log.Errorf("auth.sendCode - error: %v", err)
				return
			}

			if sysconfig.GetConfig2Bool(ctx, sysconfig.ConfigKeysRegisterNeedInviter, false, 0) {
				if req.Inviter == 0 {
					err = mtproto.ErrInviteCodeInvalid
					return
				} else {
					var inviter *mtproto.User
					inviter, err = s.UserFacade.GetUserSelf(ctx, int32(req.Inviter))
					if inviter == nil || err != nil {
						err = mtproto.ErrInviteCodeInvalid
						return
					}
				}
			}
		} else {
			//if userPassword == "" {
			//	log.Errorf("this user unset password - %s => %v", phoneNumber, err)
			//	err = mtproto.ErrUserPasswordNotSet
			//	return
			//}
			//
			//if req.Password != userPassword {
			//	log.Errorf("check password fail %s == %s", req.Password, userPassword)
			//	err = mtproto.ErrPasswordVerifyInvalid
			//	return
			//}
			err = s.checkPassword(ctx, phoneNumber, userPassword, req.Password)
			if err != nil {
				log.Errorf("check password err: %v", err)
				return
			}
		}
	} else {
		err = s.checkPhoneNumberLogin(ctx, phoneNumber)
		if err != nil {
			log.Errorf("check phone number login err: %v", err)
			return
		}
	}

	banned := s.BannedFacade.CheckPhoneNumberBanned(ctx, phoneNumber)
	if banned {
		log.Warnf("{phone_number: %s} banned: %v", phoneNumber, err)
		err = mtproto.ErrPhoneNumberBanned
		return
	}

	if s.AuthCore.IsBannedIp(ctx, ipAddr) {
		err = mtproto.ErrIpAddressBanned
		return
	}

	actionType := logic.GetActionType(request)
	if err = s.AuthCore.CheckCanDoAction(ctx, authKeyId, phoneNumber, actionType); err != nil {
		log.Warnf("check can do action - %s: %v", phoneNumber, err)
		return
	}

	var (
		phoneRegistered bool
		user            *mtproto.User
	)

	if user, err = s.UserFacade.GetUserSelfByPhoneNumber(ctx, phoneNumber); err != nil {
		log.Errorf("checkPhoneNumberExist error: %v", err)
		return
	}

	phoneRegistered = user != nil
	if phoneRegistered {
		if !s.AuthCore.IsBindIp(ctx, user.Id, ipAddr) {
			err = mtproto.ErrUserBindedIpAddress
			return
		}
	}

	codeData, err2 := s.DoAuthSendCode(ctx,
		authKeyId,
		sessionId,
		phoneNumber,
		phoneRegistered,
		request.Settings.AllowFlashcall,
		request.Settings.CurrentNumber,
		request.ApiId,
		request.ApiHash,
		func(codeData2 *model.PhoneCodeTransaction) error {
			if codeData2.State == model.CodeStateSent {
				log.Debugf("codeSent")
				return nil
			}
			if useJson && (!req.IsSignup || !needPhoneCode) {
				log.Debugf("not check phone code:%d:%s", codeData2.AuthKeyId, codeData2.PhoneNumber)
				codeData2.PhoneCode = "88888"
				codeData2.SentCodeType = model.CodeTypeSms
				codeData2.PhoneCodeExtraData = codeData2.PhoneCode
			} else if phoneRegistered && status_client.CheckUserOnline(ctx, user.Id) {
				log.Debugf("user online")
				codeData2.SentCodeType = model.CodeTypeApp
				codeData2.PhoneCodeExtraData = codeData2.PhoneCode
				langType := model2.GetLangType(ctx, s.AuthSessionRpcClient)
				go func() {
					s.pushSignInMessage(context.Background(), user.Id, codeData2.PhoneCode, langType)
				}()
			} else {
				log.Debugf("send code by sms")
				if extraData, err := s.VerifyCodeInterface.SendSmsVerifyCode(
					context.Background(),
					phoneNumber,
					codeData2.PhoneCode,
					codeData2.PhoneCodeHash,
					model2.GetLangType(ctx, s.AuthSessionRpcClient)); err != nil {
					return err
				} else {
					codeData2.SentCodeType = model.CodeTypeSms
					codeData2.PhoneCodeExtraData = extraData
				}
			}

			codeData2.NextCodeType = model.CodeTypeSms
			codeData2.State = model.CodeStateSent

			return nil
		})

	if err2 != nil {
		log.Error(err2.Error())
		err = err2
		return
	}
	reply = codeData.ToAuthSentCode()
	return
}

func (s *Service) AuthSendCode(ctx context.Context, request *mtproto.TLAuthSendCode) (reply *mtproto.Auth_SentCode, err error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("auth.sendCode - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("auth.sendCode - error: %v", err)
		return
	}

	reply, err = s.authSendCode(ctx, md.AuthId, md.SessionId, md.ClientAddr, request)
	if err != nil {
		log.Errorf("auth.sendCode - error: {%v}", err)
		return
	}

	logic.DoLogAuthAction(s.AuthCore, md, request.PhoneNumber, logic.GetActionType(request), "auth.sendCode#86aef0ec")

	log.Debugf("auth.sendCode - reply: %s", reply.DebugString())
	return
}

func (s *Service) checkIpLimit(ctx context.Context, ip string) error {
	setting := sysconfig.GetConfig2Uint32(ctx, sysconfig.ConfigKeyRegisterLimitOfIp, 1, 0)
	if setting == 0 {
		log.Debugf("checkIpLimit - no limit")
		return nil
	}

	l, err := s.AuthCore.GetUserListByIp(ctx, ip)
	if err != nil {
		log.Errorf("checkIpLimit - GetUserListByIp - error: {%v}", err)
		return err
	}

	c := uint32(len(l))
	log.Debugf("checkIpLimit - ip: %s, actually count: %d, system setting: %d", ip, c, setting)
	if c >= setting {
		log.Errorf("checkIpLimit - exceed limit - ip: %s, actually count: %d, system setting: %d", ip, c, setting)
		return mtproto.ErrPhoneNumberAppSignupForbidden
	}

	return nil
}

func (s *Service) checkPassword(ctx context.Context, phoneNumber string, userPassword, reqPassword string) (err error) {
	if userPassword == "" {
		log.Errorf("this user unset password - %s => %v", phoneNumber, err)
		err = mtproto.ErrUserPasswordNotSet
		return
	}

	user, err := s.UserFacade.GetUserSelfByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		log.Errorf("checkPassword error: %v", err)
		return
	}

	uid := user.Id

	flooded, err := s.AuthCore.IsPasswordFlood(ctx, uid)
	if err != nil {
		return err
	}

	if flooded {
		return mtproto.ErrPhonePasswordFlood
	}

	if reqPassword == userPassword {
		s.AuthCore.CleanPasswordFloodCount(ctx, uid)
		return nil
	}

	log.Errorf("check password fail %s == %s", reqPassword, userPassword)

	limit := sysconfig.GetConfig2Int32(ctx, sysconfig.ConfigKeyPasswordFloodLimit, 8, 0)
	interval := sysconfig.GetConfig2Int32(ctx, sysconfig.ConfigKeyPasswordFloodInterval, 3600, 0)
	flooded, err = s.AuthCore.IncPFCount(ctx, uid, interval, 86400, limit)
	if err != nil {
		return err
	}
	log.Debugf("check password - flooded: %v", flooded)
	if flooded {
		s.pushPasswordFloodMessage(ctx, uid)
		return mtproto.ErrPhonePasswordFlood
	} else {
		return mtproto.ErrPasswordVerifyInvalid
	}
}

func (s *Service) checkPhoneNumberLogin(ctx context.Context, phoneNumber string) error {
	user, _ := s.UserFacade.GetUserSelfByPhoneNumber(ctx, phoneNumber)
	// 用户已注册， 返回OK
	if user != nil {
		return nil
	}

	// 未注册需要判断是否开启了邀请码
	if sysconfig.GetConfig2Bool(ctx, sysconfig.ConfigKeysRegisterNeedInviter, false, 0) {
		return mtproto.ErrInviteCodeInvalid
	}

	return nil
}
