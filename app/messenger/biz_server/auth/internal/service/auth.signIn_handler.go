package service

import (
	"context"
	"encoding/json"

	"open.chat/app/messenger/biz_server/auth/internal/logic"
	"open.chat/app/messenger/biz_server/auth/internal/model"
	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/app/pkg/env2"
	"open.chat/app/service/auth_session/authsessionpb"
	model2 "open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/crypto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/phonenumber"
)

func (s *Service) AuthSignIn(ctx context.Context, request *mtproto.TLAuthSignIn) (reply *mtproto.Auth_Authorization, err error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("auth.signIn#bcd51581 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("auth.sendCode - error: %v", err)
		return
	}

	if request.PhoneCode == "" || request.PhoneCodeHash == "" {
		err = mtproto.ErrPhoneCodeEmpty
		log.Error(err.Error())
		return nil, err
	}

	req := sendCodeReq{}
	err = json.Unmarshal([]byte(request.PhoneNumber), &req)
	var phoneNumber string
	useJson := err == nil
	if useJson {
		if len(req.PhoneNumber) == 0 {
			req.PhoneNumber, _ = s.AuthCore.GetCachePhoneNumber(ctx, md.AuthId)
		}
		request.PhoneNumber = req.PhoneNumber
	}

	phoneNumber, err = checkPhoneNumberInvalid(request.PhoneNumber)
	if (!useJson || req.IsSignup) && err != nil {
		log.Errorf("check phone_number(%s) error - %v", request.PhoneNumber, err)
		err = mtproto.ErrPhoneNumberInvalid
		return
	}

	if useJson {
		if len(req.Password) == 0 {
			log.Errorf("need input password")
			err = mtproto.ErrUserPasswordNeeded
			return
		}

		if phoneNumber == "" {
			if !model2.CheckUsernameInvalid(req.UserName) {
				log.Errorf("user name error - %s", req.UserName)
				err = mtproto.ErrUsernameInvalid
				return
			} else {
				phoneNumber, _, err = s.UserFacade.GetPhoneAndPassword(ctx, req.UserName)
				if req.IsSignup {
					if phoneNumber != "" {
						err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USERNAME_OCCUPIED)
						return
					}
				} else {
					if phoneNumber == "" {
						log.Errorf("check user name fail - %s => %v", req.UserName, err)
						err = mtproto.ErrUserNameNotExist
						return
					}
				}
			}
		}
	}

	actionType := logic.GetActionType(request)
	if err = s.AuthCore.CheckCanDoAction(ctx, md.AuthId, phoneNumber, actionType); err != nil {
		log.Warnf("check can do action - %s: %v", phoneNumber, err)
		return
	}

	codeData, err2 := s.DoAuthSignIn(ctx, md.AuthId, phoneNumber, request.PhoneCode, request.PhoneCodeHash, func(codeData2 *model.PhoneCodeTransaction) error {
		if request.PhoneCode == codeData2.PhoneCodeExtraData {
			return nil
		}
		return s.VerifyCodeInterface.VerifySmsCode(ctx, codeData2.PhoneCodeHash, request.PhoneCode, codeData2.PhoneCodeExtraData)
	})

	if err2 != nil {
		log.Error(err2.Error())
		err = err2
		return
	}

	logic.DoLogAuthAction(s.AuthCore, md, phoneNumber, logic.GetActionType(request), "auth.signIn#bcd51581")
	if !codeData.PhoneNumberRegistered {
		if !env2.PredefinedUser2 {
			if md.Layer >= 104 {
				reply = mtproto.MakeTLAuthAuthorizationSignUpRequired(&mtproto.Auth_Authorization{}).To_Auth_Authorization()
				log.Debugf("auth.signIn - reply: %s\n", reply.DebugString())
				return
			} else {
				log.Warnf("auth.signIn - not registered, next step auth.signIn, %v", err)
				err = mtproto.ErrPhoneNumberUnoccupied
				return
			}
		} else {
			predefinedUser, err2 := s.UserFacade.GetPredefinedUser(ctx, phoneNumber)
			if err2 != nil {
				log.Warnf("auth.signIn - not registered, next step auth.signIn, %v", err2)
				err = mtproto.ErrPhoneNumberUnoccupied
				return
			}

			key := crypto.CreateAuthKey()
			_, err = s.AuthSessionRpcClient.SessionSetAuthKey(ctx, &authsessionpb.TLSessionSetAuthKey{
				AuthKey: &authsessionpb.AuthKeyInfo{
					AuthKeyId:          key.AuthKeyId(),
					AuthKey:            key.AuthKey(),
					AuthKeyType:        model2.AuthKeyTypePerm,
					PermAuthKeyId:      key.AuthKeyId(),
					TempAuthKeyId:      0,
					MediaTempAuthKeyId: 0,
				},
				FutureSalt: nil,
			})

			if err != nil {
				log.Errorf("create user secret key error")
				err = mtproto.ErrPhoneNumberUnoccupied
				return
			}

			pNumber, _ := phonenumber.MakePhoneNumberHelper(phoneNumber, "")
			_, err = s.UserFacade.CreateNewUser(ctx,
				key.AuthKeyId(),
				phoneNumber,
				pNumber.GetRegionCode(),
				predefinedUser.GetFirstName().GetValue(),
				predefinedUser.GetLastName().GetValue())
			if err != nil {
				log.Errorf("create user error")
				err = mtproto.ErrPhoneNumberUnoccupied
				return
			}

			codeData.PhoneNumberRegistered = true
		}
	}

	var (
		user *mtproto.User
	)

	user, err = s.UserFacade.GetUserSelfByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		log.Errorf("user(%s) is err - %v", phoneNumber, err)
		return
	} else if user == nil {
		log.Errorf("user(%s) is nil", phoneNumber)
		err = mtproto.ErrInternelServerError
		return
	}

	s.AuthSessionRpcClient.SessionBindAuthKeyUser(ctx, &authsessionpb.TLSessionBindAuthKeyUser{
		AuthKeyId: md.AuthId,
		UserId:    user.GetId(),
	})

	// 暂时屏蔽，二步验证的时候再打开
	//if sessionPasswordNeeded, err := s.AccountFacade.CheckSessionPasswordNeeded(ctx, user.Id); sessionPasswordNeeded {
	//	log.Infof("auth.signIn - registered, next step auth.checkPassword: %v", err)
	//	err = mtproto.ErrSessionPasswordNeeded
	//	return nil, err
	//}

	s.DeletePhoneCode(context.Background(), md.AuthId, phoneNumber, request.PhoneCodeHash)

	go func() {
		region, _ := s.Dao.GetCountryAndRegionByIp(md.ClientAddr)
		signInN := model2.MakeSignInServiceNotification(user, md.AuthId, md.Client, region, md.ClientAddr)
		pushUpdates := model2.MakeUpdatesByUpdates(signInN)

		sync_client.SyncUpdatesNotMe(context.Background(), user.GetId(), md.AuthId, pushUpdates)
	}()

	reply = mtproto.MakeTLAuthAuthorization(&mtproto.Auth_Authorization{
		User: user,
	}).To_Auth_Authorization()

	log.Debugf("auth.signIn - reply: %s\n", reply.DebugString())
	return
}
