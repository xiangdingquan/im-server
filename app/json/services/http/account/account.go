package account

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"math/rand"
	msg_facade "open.chat/app/messenger/msg/facade"
	"open.chat/app/service/auth_session/authsessionpb"
	authsession_client "open.chat/app/service/auth_session/client"
	"open.chat/pkg/log"
	"strings"
	"time"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"open.chat/app/json/consts"
	"open.chat/app/json/helper"
	svc "open.chat/app/json/service"
	"open.chat/app/json/service/http"
	"open.chat/app/sysconfig"
	"open.chat/mtproto"

	types "github.com/gogo/protobuf/types"
	"open.chat/app/messenger/biz_server/account"
	"open.chat/app/messenger/biz_server/upload"

	"open.chat/app/json/services/http/account/core"
	banned_facade "open.chat/app/service/biz_service/banned/facade"
	user_client "open.chat/app/service/biz_service/user/client"
	idgen "open.chat/app/service/idgen/client"

	"open.chat/app/messenger/biz_server/auth"
	dfs_facade "open.chat/app/service/dfs/facade"

	"open.chat/app/messenger/biz_server/photos"
	username_facade "open.chat/app/service/biz_service/username/facade"
	"open.chat/model"
	"open.chat/pkg/phonenumber"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type cls struct {
	*core.BannedIpCore
	mtproto.RPCAuthServer
	mtproto.RPCAccountServer
	mtproto.RPCPhotosServer
	mtproto.RPCUploadServer
	user_client.UserFacade
	banned_facade.BannedFacade
	dfs_facade.DfsFacade
	username_facade.UsernameFacade
	msg_facade.MsgFacade
	AuthSessionRpcClient authsessionpb.RPCSessionClient
}

// New .
func New(s *svc.Service, rg *bm.RouterGroup) {
	service := &cls{
		BannedIpCore:     core.New(nil),
		RPCAuthServer:    auth.New(),
		RPCAccountServer: account.New(),
		RPCPhotosServer:  photos.New(),
		RPCUploadServer:  upload.New(),
	}
	var (
		ac struct {
			Wardenclient *warden.ClientConfig
		}
		err error
	)
	service.UserFacade, err = user_client.NewUserFacade("local")
	checkErr(err)
	service.UsernameFacade, err = username_facade.NewUsernameFacade("local")
	checkErr(err)
	service.BannedFacade, err = banned_facade.NewBannedFacade("local")
	checkErr(err)
	service.DfsFacade, err = dfs_facade.NewDfsFacade("dfs")
	checkErr(err)
	service.MsgFacade, err = msg_facade.NewMsgFacade("emsg")
	checkErr(err)
	checkErr(paladin.Get("application.toml").UnmarshalTOML(&ac))
	service.AuthSessionRpcClient, err = authsession_client.New(ac.Wardenclient)
	checkErr(err)
	http.RegisterAccount(service, rg)
}

func (s *cls) SignUp(ctx context.Context, r *http.TAccountSignUp) *helper.ResultJSON {
	var (
		userName    string
		password    string
		nickName    string
		phoneNumber string
		phone       string
		countryCode string
		err         error
	)

	needPhoneCode := sysconfig.GetConfig2Bool(ctx, sysconfig.ConfigKeysRegisterNeedPhoneCode, false, 0)
	//log.Debugf("account.signUp needPhoneCode:%b", needPhoneCode)

	if needPhoneCode {
		phoneNumber = strings.Trim(r.PhoneNumber, " ")
		if phoneNumber == "" {
			if needPhoneCode {
				return &helper.ResultJSON{Code: -2, Msg: "please post phone number"}
			}
			r.PhoneNumber, _ = idgen.GetNextPhoneNumber(ctx, "86100")
		}

		pNumber, err := phonenumber.MakePhoneNumberHelper(r.PhoneNumber, "")
		if err != nil {
			return &helper.ResultJSON{Code: -6, Msg: "phone number invalid"}
		}

		phone = pNumber.GetNormalizeDigits()
		if ok, _ := s.UserFacade.CheckPhoneNumberExist(ctx, phone); ok {
			return &helper.ResultJSON{Code: -9, Msg: "phone number is exist"}
		}

		phoneCode := strings.Trim(r.PhoneCode, " ")
		if phoneCode == "" {
			return &helper.ResultJSON{Code: -11, Msg: "please post phone code"}
		}
		if helper.VerifyCode(ctx, consts.SmsCodeType_RegistAccount, phone, phoneCode) != nil {
			return &helper.ResultJSON{Code: -12, Msg: "phone code verification fail"}
		}

		countryCode = pNumber.GetRegionCode()
	} else {
		userName = strings.Trim(r.UserName, " ")
		if userName == "" {
			return &helper.ResultJSON{Code: -1, Msg: "please post user name"}
		}

		password = strings.Trim(r.Password, " ")
		if password == "" {
			return &helper.ResultJSON{Code: -3, Msg: "please post password"}
		}

		if !model.CheckUsernameInvalid(userName) {
			return &helper.ResultJSON{Code: -5, Msg: "user name invalid"}
		}

		checked, err := s.UsernameFacade.CheckUsername(ctx, userName)
		if err != nil {
			return &helper.ResultJSON{Code: -7, Msg: "system error"}
		}

		if checked == model.UsernameExisted {
			return &helper.ResultJSON{Code: -8, Msg: "user name is exist"}
		}

		phone, _ = idgen.GetNextPhoneNumber(ctx, "86100")
	}

	//log.Debugf("accounts.signUp, needPhoneCode:%b, phone:%s", needPhoneCode, phone)

	nickName = strings.Trim(r.NickName, " ")
	if nickName == "" {
		return &helper.ResultJSON{Code: -4, Msg: "please post nick name"}
	}

	needInviter := sysconfig.GetConfig2Bool(ctx, sysconfig.ConfigKeysRegisterNeedInviter, false, 0)
	if r.Inviter == 0 {
		if needInviter {
			return &helper.ResultJSON{Code: -9, Msg: "please post invite code"}
		}
	}

	var inviter *mtproto.User = nil
	if r.Inviter != 0 {
		inviter, err = s.UserFacade.GetUserSelf(ctx, int32(r.Inviter))
		if inviter == nil || err != nil {
			if needInviter {
				return &helper.ResultJSON{Code: -10, Msg: "invite code is invalid"}
			}
		}
	}

	user, err := s.UserFacade.CreateNewUser(ctx, 0, phone, countryCode, nickName, "")
	if user == nil || err != nil {
		return &helper.ResultJSON{Code: -13, Msg: "create user fail"}
	}
	s.UserFacade.UpdateUserPassword(ctx, user.Id, password)
	s.UserFacade.UpdateUsername(ctx, user.Id, userName)
	s.UsernameFacade.UpdateUsername(ctx, model.PEER_USER, user.Id, userName)
	if inviter != nil {
		s.UserFacade.UpdateUserInviter(ctx, user.Id, inviter.GetId())
		//有邀请人添加为好友
		addInviterAsFriend := sysconfig.GetConfig2Bool(ctx, sysconfig.ConfigKeysRegisterAddInviterAsFriend, false, 0)
		//log.Debugf("account.signUp addInviterAsFriend:%b", addInviterAsFriend)
		if addInviterAsFriend {
			s.UserFacade.AddContact(ctx, user.Id, inviter.GetId(), false, inviter.GetFirstName().GetValue(), inviter.GetLastName().GetValue())
			s.UserFacade.AddContact(ctx, inviter.GetId(), user.Id, false, nickName, "")
			helper.MakeSender(uint32(user.Id), 0, 3, 101).PushMessage(ctx, uint32(inviter.GetId()), nil)
		}
	}

	addCustomerServiceAsFriend := func(ctx context.Context, uid int32, firstName, lastName string) {
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

		pushMsg := func(ctx context.Context, userId, csId int32) {
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

				return model.Localize(ctx, s.AuthSessionRpcClient, model.LocalizationWords{
					model.LocalizationEN:      j.EN,
					model.LocalizationCN:      j.CN,
					model.LocalizationDefault: j.EN,
				}), nil
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
		pushMsg(ctx, uid, csId)
	}
	addCustomerServiceAsFriend(ctx, user.Id, nickName, "")

	//pUser, err := s.UserFacade.GetUserSelf(ctx, user.GetId())
	//if pUser == nil || err != nil {
	//	return &helper.ResultJSON{Code: -14, Msg: "invite code is invalid"}
	//}

	data := struct {
		Uid int32 `json:"userId"`
	}{
		Uid: user.GetId(),
	}
	return &helper.ResultJSON{Code: 0, Msg: "success", Data: data}
}

func (s *cls) UpdateUsername(ctx context.Context, r *http.TAccountUpdateUsername) *helper.ResultJSON {
	ctx, err := helper.DefaultMetadata(ctx, r.Uid, 0)
	if err != nil {
		return &helper.ResultJSON{Code: -1, Msg: err.Error()}
	}
	req := &mtproto.TLAccountUpdateUsername{
		Username: r.Username,
	}
	_, err = s.RPCAccountServer.AccountUpdateUsername(ctx, req)
	if err != nil {
		return &helper.ResultJSON{Code: -2, Msg: err.Error()}
	}
	return &helper.ResultJSON{Code: 0, Msg: "success"}
}

func (s *cls) UpdateProfile(ctx context.Context, r *http.TAccountUpdateProfile) *helper.ResultJSON {
	ctx, err := helper.DefaultMetadata(ctx, r.Uid, 0)
	if err != nil {
		return &helper.ResultJSON{Code: -1, Msg: err.Error()}
	}
	req := &mtproto.TLAccountUpdateProfile{}
	if len(r.About) != 0 {
		req.About = &types.StringValue{Value: r.About}
	} else {
		req.FirstName = &types.StringValue{Value: r.FirstName}
		req.LastName = &types.StringValue{Value: r.LastName}
	}
	_, err = s.RPCAccountServer.AccountUpdateProfile(ctx, req)
	if err != nil {
		return &helper.ResultJSON{Code: -2, Msg: err.Error()}
	}
	return &helper.ResultJSON{Code: 0, Msg: "success"}
}

func (s *cls) UpdatePhoto(ctx context.Context, r *http.TAccountUpdatePhoto) *helper.ResultJSON {
	creatorId := rand.Int63()
	ctx, err := helper.DefaultMetadata(ctx, r.Uid, creatorId)
	if err != nil {
		return &helper.ResultJSON{Code: -3, Msg: err.Error()}
	}
	req := &mtproto.TLPhotosUploadProfilePhoto{}
	req.File, err = helper.UrlToInputFile(r.Photo, func(fileId int64, filePart int32, bytes []byte) error {
		return s.DfsFacade.WriteFilePartData(ctx, creatorId, fileId, filePart, bytes)
	})

	if err != nil {
		return &helper.ResultJSON{Code: -1, Msg: "save file fail"}
	}

	res, err := s.RPCPhotosServer.PhotosUploadProfilePhoto(ctx, req)
	if err != nil {
		return &helper.ResultJSON{Code: -2, Msg: "upload photo fail"}
	}

	res2, err := s.RPCPhotosServer.PhotosUpdateProfilePhotoF0BB5152(ctx, &mtproto.TLPhotosUpdateProfilePhotoF0BB5152{
		Id: mtproto.MakeTLInputPhoto(&mtproto.InputPhoto{
			Id: res.GetPhoto().GetId(),
		}).To_InputPhoto(),
	})
	if err != nil {
		return &helper.ResultJSON{Code: -3, Msg: "update photo fail"}
	}

	data := struct {
		PhotoId int64 `json:"photoId"`
	}{
		PhotoId: res2.GetPhotoId(),
	}
	return &helper.ResultJSON{Code: 0, Msg: "success", Data: data}
}

func (s *cls) GetPhoto(ctx context.Context, r *http.TAccountGetPhoto) *[]byte {
	ctx, err := helper.DefaultMetadata(ctx, r.Uid, 0)
	if err != nil {
		return nil
	}

	user, err := s.UserFacade.GetUserSelf(ctx, int32(r.Uid))
	if err != nil || user == nil {
		return nil
	}

	userPhoto := user.GetPhoto()
	if userPhoto == nil {
		return nil
	}

	var fl *mtproto.FileLocation
	if r.BigPhoto && userPhoto.GetPhotoBig() != nil {
		fl = userPhoto.GetPhotoBig()
	} else if userPhoto.GetPhotoSmall() != nil {
		fl = userPhoto.GetPhotoSmall()
	}
	if fl == nil {
		return nil
	}

	limit := 1024 * 100
	gf := &mtproto.TLUploadGetFile{
		Location: mtproto.MakeTLInputPeerPhotoFileLocation(&mtproto.InputFileLocation{
			VolumeId: fl.VolumeId,
			LocalId:  fl.LocalId,
			Peer:     mtproto.MakeTLInputPeerSelf(nil).To_InputPeer(),
		}).To_InputFileLocation(),
		Offset: 0,
		Limit:  int32(limit),
	}
	data := make([]byte, 0)
	uf, err := s.RPCUploadServer.UploadGetFile(ctx, gf)
	if err != nil {
		return &data
	}
	data = append(data, uf.Bytes...)
	for len(uf.Bytes) >= limit {
		gf.Offset += int32(limit)
		uf, err = s.RPCUploadServer.UploadGetFile(ctx, gf)
		if err != nil {
			return &data
		}
		data = append(data, uf.Bytes...)
	}
	return &data
}

func (s *cls) ToggleBan(ctx context.Context, r *http.TAccountBan) *helper.ResultJSON {
	ctx, err := helper.DefaultMetadata(ctx, r.Uid, 0)
	if err != nil {
		return &helper.ResultJSON{Code: -1, Msg: err.Error()}
	}

	getPhone := func(uid int32) string {
		phone, err := s.UserFacade.GetPhoneById(ctx, uid)
		if err != nil {
			log.Errorf("account.toggleBan: getPhone: %v", err)
			return ""
		}
		return phone
	}

	phone := getPhone(int32(r.Uid))
	//log.Debugf("ToggleBan uid: %d phone: %s", r.Uid, phone)
	if phone == "" {
		return &helper.ResultJSON{Code: -3, Msg: "get user phone number fail"}
	}

	if r.Expires != 0 {
		s.BannedFacade.Ban(ctx, phone, int32(r.Expires), r.Reason)
	} else {
		s.BannedFacade.UnBan(ctx, phone)
	}

	return &helper.ResultJSON{Code: 0, Msg: "success"}
}

func (s *cls) CreateVirtual(ctx context.Context, r *http.TCreateVirtual) *helper.ResultJSON {
	nickName := strings.Trim(r.NickName, " ")
	if nickName == "" {
		return &helper.ResultJSON{Code: -2, Msg: "please post nick name"}
	}

	phoneNumber, _ := idgen.GetNextPhoneNumber(ctx, "86101")
	pNumber, err := phonenumber.MakePhoneNumberHelper(phoneNumber, "")
	if err != nil {
		return &helper.ResultJSON{Code: -3, Msg: "phone number invalid"}
	}

	phone := pNumber.GetNormalizeDigits()
	if ok, _ := s.UserFacade.CheckPhoneNumberExist(ctx, phone); ok {
		return &helper.ResultJSON{Code: -4, Msg: "phone number is exist"}
	}

	user, err := s.UserFacade.CreateNewUser(ctx, 0, phone, pNumber.GetRegionCode(), nickName, "")
	if user == nil || err != nil {
		return &helper.ResultJSON{Code: -5, Msg: "create user fail"}
	}

	data := struct {
		Uid int32 `json:"userId"`
	}{
		Uid: user.GetId(),
	}
	return &helper.ResultJSON{Code: 0, Msg: "success", Data: data}
}
