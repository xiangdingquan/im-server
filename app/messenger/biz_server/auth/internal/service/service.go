package service

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"open.chat/app/sysconfig"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"open.chat/app/json/helper"
	"open.chat/app/json/services/handler/call"
	"open.chat/app/json/services/handler/redpacket"
	"open.chat/app/json/services/handler/remittance"
	"open.chat/app/messenger/biz_server/auth/internal/core"
	"open.chat/app/messenger/biz_server/auth/internal/dao"
	"open.chat/app/messenger/biz_server/auth/internal/logic"
	msg_facade "open.chat/app/messenger/msg/facade"
	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/app/pkg/code"
	"open.chat/app/pkg/env2"
	"open.chat/app/service/auth_session/authsessionpb"
	authsession_client "open.chat/app/service/auth_session/client"
	account_facade "open.chat/app/service/biz_service/account/facade"
	banned_facade "open.chat/app/service/biz_service/banned/facade"
	user_client "open.chat/app/service/biz_service/user/client"
	username_facade "open.chat/app/service/biz_service/username/facade"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
	"open.chat/pkg/phonenumber"
	"time"
)

type Service struct {
	authsessionpb.RPCSessionClient
	user_client.UserFacade
	banned_facade.BannedFacade
	account_facade.AccountFacade
	msg_facade.MsgFacade
	username_facade.UsernameFacade
	// *core.AuthCore
	*logic.AuthLogic
}

type sendCodeReq struct {
	IsSignup    bool   `json:"isSignup,omitempty"`
	Inviter     uint32 `json:"inviteCode,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	UserName    string `json:"userName,omitempty"`
	Password    string `json:"password"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func New() *Service {
	var (
		ac struct {
			Wardenclient   *warden.ClientConfig
			PredefinedUser bool
			MeSmsUrl       string
			Code           *code.SmsVerifyCodeConfig
			Redpacket      *redpacket.RedpacketConfig
			Remittance     *remittance.Config
			Agora          *call.ThirdAgoraConfig
			Imapi          *helper.HttpApiConfig
		}
		err error
	)

	checkErr(paladin.Get("application.toml").UnmarshalTOML(&ac))
	log.Debugf("ac: Wardenclient: %#v, PredefinedUser: %v, code: %v",
		ac.Wardenclient,
		ac.PredefinedUser,
		ac.Code)

	s := new(Service)
	// s.Dao = dao.New()
	if ac.PredefinedUser {
		env2.PredefinedUser = ac.PredefinedUser
	}
	//s.MeSmsUrl = ac.MeSmsUrl
	//if s.MeSmsUrl != "" {
	//	env2.SMS_CODE_NAME = "ME"
	//}

	// log.Debugf("%v", ac.Wardenclient)
	s.RPCSessionClient, err = authsession_client.New(ac.Wardenclient)
	checkErr(err)

	s.AuthLogic = logic.NewAuthSignLogic(core.New(dao.New(ac.Wardenclient)), ac.Code)
	redpacket.G_RedpacketCfg = ac.Redpacket
	remittance.G_RemittanceCfg = ac.Remittance
	call.G_AgoraToken = ac.Agora
	helper.HttpApi = ac.Imapi

	s.UserFacade, err = user_client.NewUserFacade("local")
	checkErr(err)
	s.UsernameFacade, err = username_facade.NewUsernameFacade("local")
	checkErr(err)
	s.BannedFacade, err = banned_facade.NewBannedFacade("local")
	checkErr(err)
	s.AccountFacade, err = account_facade.NewAccountFacade("local")
	checkErr(err)
	s.MsgFacade, err = msg_facade.NewMsgFacade("emsg")
	checkErr(err)

	sync_client.New()

	return s
}

// Close close the resource.
func (s *Service) Close() {
	// s.Dao.Close()
}

func checkPhoneNumberInvalid(phone string) (string, error) {
	// 3. check number
	// 3.1. empty
	if phone == "" {
		// log.Errorf("check phone_number error - empty")
		return "", mtproto.ErrPhoneNumberInvalid
	}

	// 3.2. check phone_number
	// 客户端发送的手机号格式为: "+86 111 1111 1111"，归一化
	// We need getRegionCode from phone_number
	pNumber, err := phonenumber.MakePhoneNumberHelper(phone, "")
	if err != nil {
		// log.Errorf("check phone_number error - %v", err)
		// err = mtproto.ErrPhoneNumberInvalid
		return "", mtproto.ErrPhoneNumberInvalid
	}

	return pNumber.GetNormalizeDigits(), nil
}

const signInMessageTpl1 = `Login code: %s. Do not give this code to anyone, even if they say they are from %s!

This code can be used to log in to your %s account. We never ask it for anything else.

If you didn't request this code by trying to log in on another device, simply ignore this message.

【Warning!】This software is for developers to test and debug, not for public use! If you are not a developer or a stranger, please do not download and install it! Thank you!`

const signInMessageTpl = ``

func (s *Service) getSignInMessageTpl(langType model.LangType) string {
	return model.LocalizationWords{
		model.LocalizationEN:      signInMessageTpl1,
		model.LocalizationCN:      signInMessageTpl,
		model.LocalizationDefault: signInMessageTpl1,
	}[langType]
}

func (s *Service) pushSignInMessage(ctx context.Context, signInUserId int32, code string, langType model.LangType) {
	var fromId int32 = 777000
	message := mtproto.MakeTLMessage(&mtproto.Message{
		Out:             true,
		Date:            int32(time.Now().Unix()),
		FromId_FLAGPEER: model.MakePeerUser(fromId),
		ToId:            model.MakePeerUser(signInUserId),
		Message:         fmt.Sprintf(s.getSignInMessageTpl(langType), code, env2.MY_APP_NAME),
		Entities: []*mtproto.MessageEntity{
			mtproto.MakeTLMessageEntityBold(&mtproto.MessageEntity{
				Offset: 4,
				Length: 3,
			}).To_MessageEntity(),
			mtproto.MakeTLMessageEntityBold(&mtproto.MessageEntity{
				Offset: 23,
				Length: 2,
			}).To_MessageEntity(),
		},
	}).To_Message()
	randomId := rand.Int63()
	s.MsgFacade.PushUserMessage(ctx, 1, fromId, signInUserId, randomId, message)
}

func (s *Service) pushCustomerServiceMessageForRegister(ctx context.Context, userId, csId int32, langType model.LangType) {
	messageLocalized := func() (string, error) {
		msg := sysconfig.GetConfig2String(ctx, sysconfig.ConfigKeyCustomerServiceMessageForRegister, "{}", 0)
		type csm struct {
			EN string `json:"en"`
			CN string `json:"cn"`
		}
		j := &csm{}
		err := json.Unmarshal([]byte(msg), j)
		if err != nil {
			log.Errorf("pushCustomerServiceMessageForRegister error - %v", err)
			return "", err
		}

		return model.LocalizationWords{
			model.LocalizationEN:      j.EN,
			model.LocalizationCN:      j.CN,
			model.LocalizationDefault: j.EN,
		}[langType], nil
	}
	m, err := messageLocalized()
	if err != nil {
		log.Errorf("pushCustomerServiceMessageForRegister error - %v", err)
		return
	}

	fromId := csId
	message := mtproto.MakeTLMessage(&mtproto.Message{
		Out:             true,
		Date:            int32(time.Now().Unix()),
		FromId_FLAGPEER: model.MakePeerUser(fromId),
		ToId:            model.MakePeerUser(userId),
		Message:         m,
		Entities:        []*mtproto.MessageEntity{},
	}).To_Message()
	randomId := rand.Int63()
	s.MsgFacade.PushUserMessage(ctx, 1, fromId, userId, randomId, message)
}

func (s *Service) pushPasswordFloodMessage(ctx context.Context, uid int32) {
	m := model.Localize(ctx, s.AuthSessionRpcClient, model.LocalizationWords{
		model.LocalizationEN:      "You have tried to log in too many times. Please try again later.",
		model.LocalizationCN:      "您已经尝试登录太多次，请稍后再试。",
		model.LocalizationDefault: "You have tried to log in too many times. Please try again later.",
	})
	message := mtproto.MakeTLMessage(&mtproto.Message{
		Out:             true,
		Date:            int32(time.Now().Unix()),
		FromId_FLAGPEER: model.MakePeerUser(777000),
		ToId:            model.MakePeerUser(uid),
		Message:         m,
		Entities:        []*mtproto.MessageEntity{},
	}).To_Message()
	randomId := rand.Int63()
	s.MsgFacade.PushUserMessage(ctx, 1, 777000, uid, randomId, message)
}
