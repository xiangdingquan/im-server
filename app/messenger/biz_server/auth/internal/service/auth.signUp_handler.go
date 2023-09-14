package service

import (
	"context"
	"encoding/json"
	"math/rand"
	"open.chat/app/sysconfig"

	"open.chat/app/json/helper"
	"open.chat/app/messenger/biz_server/auth/internal/logic"
	"open.chat/app/messenger/biz_server/auth/internal/model"
	"open.chat/app/pkg/env2"
	"open.chat/app/service/auth_session/authsessionpb"
	model2 "open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/crypto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/phonenumber"
)

func (s *Service) AuthSignUp(ctx context.Context, request *mtproto.TLAuthSignUp) (reply *mtproto.Auth_Authorization, err error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("auth.signUp - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	logic.DoLogAuthAction(s.AuthCore, md, request.PhoneNumber, logic.GetActionType(request), "auth.signUp")

	var phoneCode *string = nil
	switch request.Constructor {
	case mtproto.CRC32_auth_signUp_1b067634:
		if request.PhoneCode == "" {
			err = mtproto.ErrPhoneCodeEmpty
			log.Errorf("auth.signUp - error: %v", err)
			return
		} else {
			phoneCode = &request.PhoneCode
		}
	case mtproto.CRC32_auth_signUp_80eee427:
		phoneCode = nil
	default:
		err = mtproto.ErrTypeConstructorInvalid
		log.Errorf("auth.signUp - error: %v", err)
		return
	}
	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("auth.signUp - error: %v", err)
		return
	}

	req := sendCodeReq{}
	err = json.Unmarshal([]byte(request.PhoneNumber), &req)
	useJson := err == nil
	if useJson {
		if len(req.PhoneNumber) == 0 {
			req.PhoneNumber, _ = s.AuthCore.GetCachePhoneNumber(ctx, md.AuthId)
		}
		request.PhoneNumber = req.PhoneNumber
	}

	if request.PhoneNumber == "" {
		log.Errorf("check phone_number error - empty")
		err = mtproto.ErrPhoneNumberInvalid
		return
	}

	pNumber, err := phonenumber.MakePhoneNumberHelper(request.PhoneNumber, "")
	if err != nil {
		log.Errorf("check phone_number error - %v", err)
		err = mtproto.ErrPhoneNumberInvalid
		return
	}
	phoneNumber := pNumber.GetNormalizeDigits()

	if request.PhoneCodeHash == "" {
		log.Errorf("check phone_code_hash error - empty")
		err = mtproto.ErrPhoneCodeHashEmpty
		return
	}

	if request.FirstName == "" {
		log.Errorf("check first_name error - empty")
		err = mtproto.ErrFirstNameInvalid
		return
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	var (
		codeData *model.PhoneCodeTransaction
	)
	codeData, err = s.DoAuthSignUp(ctx, md.AuthId, phoneNumber, phoneCode, request.PhoneCodeHash)
	if err != nil {
		log.Error(err.Error())
		return
	}

	var (
		user *mtproto.User
	)

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
		return nil, err
	}

	var (
		firstName      = request.FirstName
		lastName       = request.LastName
		username       string
		password       string
		predefinedUser *mtproto.PredefinedUser
	)

	if useJson {
		username = req.UserName
		password = req.Password
	}

	if env2.PredefinedUser {
		predefinedUser, _ = s.UserFacade.GetPredefinedUser(ctx, phoneNumber)
		if predefinedUser == nil {
			log.Errorf("check predefinedUsers error - %v", err)
			err = mtproto.ErrPhoneNumberInvalid
			return nil, err
		}
		firstName = predefinedUser.GetFirstName().GetValue()
		lastName = predefinedUser.GetLastName().GetValue()
		username = predefinedUser.GetUsername().GetValue()
	}

	// Create new user
	if user, err = s.UserFacade.CreateNewUser(ctx,
		key.AuthKeyId(),
		phoneNumber,
		pNumber.GetRegionCode(),
		firstName,
		lastName); err != nil {

		log.Errorf("createNewUser error: %v", err)
		return
	}

	if username != "" {
		s.UserFacade.UpdateUsername(ctx, user.Id, username)
		s.UsernameFacade.UpdateUsername(ctx, model2.PEER_USER, user.Id, username)
	}

	var inviter *mtproto.User = nil
	if useJson && req.Inviter > 0 {
		inviter, _ = s.UserFacade.GetUserSelf(ctx, int32(req.Inviter))
		if inviter != nil {
			s.UserFacade.UpdateUserInviter(ctx, user.Id, int32(req.Inviter))
		}
	}

	if password != "" {
		s.UserFacade.UpdateUserPassword(ctx, user.Id, password)
	}

	if env2.PredefinedUser {
		s.UserFacade.PredefinedBindRegisteredUserId(ctx, phoneNumber, user.Id)
		if username != "" {
			s.UserFacade.UpdateVerified(ctx, user.Id, predefinedUser.Verified)
		}
	}

	user.Self = true

	// bind auth_key and user_id
	_, err = s.AuthSessionRpcClient.SessionBindAuthKeyUser(ctx, &authsessionpb.TLSessionBindAuthKeyUser{
		AuthKeyId: md.AuthId,
		UserId:    user.GetId(),
	})
	if err != nil {
		log.Errorf("bindAuthKeyUser error: %v", err)
		err = mtproto.ErrInternelServerError
	}

	reply = mtproto.MakeTLAuthAuthorization(&mtproto.Auth_Authorization{
		User: user,
	}).To_Auth_Authorization()

	langType := model2.GetLangType(ctx, s.AuthSessionRpcClient)

	go func() {
		ctx = context.Background()
		s.DeletePhoneCode(ctx, md.AuthId, phoneNumber, request.PhoneCodeHash)
		s.pushSignInMessage(ctx, user.GetId(), codeData.PhoneCode, langType)
		s.addCustomerServiceAsFriend(ctx, user.GetId(), firstName, lastName, langType)
		s.onContactSignUp(ctx, md.AuthId, user.GetId(), phoneNumber)
		if inviter != nil {
			addInviterAsFriend := sysconfig.GetConfig2Bool(ctx, sysconfig.ConfigKeysRegisterAddInviterAsFriend, false, 0)
			log.Debugf("addInviterAsFriend: %b", addInviterAsFriend)
			if addInviterAsFriend {
				//有邀请人添加为好友
				s.UserFacade.AddContact(ctx, user.Id, int32(req.Inviter), false, inviter.GetFirstName().GetValue(), inviter.GetLastName().GetValue())
				//通知邀请人加为好友
				s.UserFacade.AddContact(ctx, int32(req.Inviter), user.Id, false, firstName, lastName)
				helper.MakeSender(uint32(user.Id), 0, 3, 101).PushMessage(ctx, uint32(inviter.GetId()), nil)
			}
		}
	}()

	log.Debugf("auth.signUp#1b067634 - reply: %s\n", reply.DebugString())
	return
}

func (s *Service) onContactSignUp(ctx context.Context, authKeyId int64, userId int32, phone string) {
	log.Debugf("onContactSignUp - {phone: %s}")
	importers := s.UserFacade.GetImportersByPhone(ctx, phone)
	for _, c := range importers {
		log.Debugf("importer: %v", c)
		v := s.AccountFacade.GetSettingValueString(ctx, int32(c.ClientId), "contactSignUpNotification", "")
		if v == "true" {
			s.MsgFacade.PushUserMessage(ctx, 1, userId, int32(c.ClientId), rand.Int63(),
				model2.MakeContactSignUpMessage(userId, int32(c.ClientId)))
		} else {
			log.Debugf("not setting contactSignUpNotification")
		}
	}
	s.UserFacade.DeleteImportersByPhone(ctx, phone)
}

func (s *Service) addCustomerServiceAsFriend(ctx context.Context, uid int32, firstName, lastName string, langType model2.LangType) {
	getCustomerServiceId := func() int32 {
		l := s.UserFacade.GetCustomerServiceList(ctx)
		if len(l) == 0 {
			return 0
		}
		return l[uid%int32(len(l))]
	}
	csId := getCustomerServiceId()
	if csId == 0 {
		log.Errorf("addCustomerServiceAsFriend error - no customer service")
		return
	}

	cs, err := s.UserFacade.GetUserSelf(ctx, csId)
	if err != nil {
		log.Errorf("addCustomerServiceAsFriend error - get customer service info failed - %v", err)
		return
	}

	s.UserFacade.AddContact(ctx, uid, csId, false, cs.GetFirstName().GetValue(), cs.GetLastName().GetValue())
	s.UserFacade.AddContact(ctx, csId, uid, false, firstName, lastName)

	s.pushCustomerServiceMessageForRegister(ctx, uid, csId, langType)
}
