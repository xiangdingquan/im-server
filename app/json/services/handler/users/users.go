package users

import (
	"context"

	"open.chat/app/json/helper"
	svc "open.chat/app/json/service"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/phonenumber"

	"open.chat/app/json/service/handler"
	user_client "open.chat/app/service/biz_service/user/client"
	"open.chat/model"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type cls struct {
	user_client.UserFacade
}

// New .
func New(s *svc.Service) {
	c := &cls{}
	var err error
	c.UserFacade, err = user_client.NewUserFacade("local")
	checkErr(err)
	s.AppendServices(handler.RegisterUsers(c))
}

func (s *cls) makePhoneNumber(ctx context.Context, selfId int32, Phone string) string {
	region, _ := s.UserFacade.GetCountryCodeByUser(ctx, selfId)
	phoneNumber, err2 := phonenumber.MakePhoneNumberHelper(Phone, region)
	if err2 != nil {
		phoneNumber, err2 = phonenumber.MakePhoneNumberHelper(Phone, "")
		if err2 != nil {
			return ""
		}
	}
	return phoneNumber.GetNormalizeDigits()
}

func (s *cls) Info(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TUsersInfo) *helper.ResultJSON {
	var uids []int32 = make([]int32, 0)
	if len(r.Uids) == 0 {
		uids = append(uids, md.UserId)
		//return &helper.ResultJSON{Code: -1, Msg: "please post user id list"}
	}
	if len(r.Uids) > 500 {
		return &helper.ResultJSON{Code: -2, Msg: "user id count too much"}
	}

	//hasSelf := false
	for _, uid := range r.Uids {
		//	if uid == uint32(md.UserId) {
		//		hasSelf = true
		//	}
		//uids[i] = int32(uid)
		uids = append(uids, int32(uid))
	}
	//if !hasSelf {
	//	//uids = append(uids, md.UserId)
	//}

	users := s.GetMutableUsers(ctx, uids...)
	list := make([]*model.ImmutableUser, 0)
	for _, uid := range uids {
		if user, ok := users.GetImmutableUser(uid); ok {
			list = append(list, user)
		}
	}

	if len(list) == 0 {
		return &helper.ResultJSON{Code: 400, Msg: "get users info fail"}
	}

	var data = make([]struct {
		Uid       uint32 `json:"uId"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Username  string `json:"userName"`
		//CanSendMessage bool   `json:"canSendMessage"`
		//CanAddContact  bool   `json:"canAddContact"`
		IsInternal bool `json:"isInternal"`
	}, len(list))
	for i, user := range list {
		//isMeContact := s.UserFacade.CheckContact(ctx, md.UserId, user.User.Id)
		//isUserContact := s.UserFacade.CheckContact(ctx, user.User.Id, md.UserId)
		data[i].Uid = uint32(user.User.Id)
		data[i].FirstName = user.User.FirstName
		data[i].LastName = user.User.LastName
		data[i].Username = user.User.Username
		data[i].IsInternal = user.User.IsInternal
		//data[i].CanSendMessage = user.User.IsInternal || isMeContact || isUserContact
		//data[i].CanAddContact = !isMeContact && (user.User.IsInternal || isUserContact)
	}
	return &helper.ResultJSON{Code: 200, Msg: "success", Data: data}
}

func (s *cls) SearchByPhone(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TUsersSearchByPhone) *helper.ResultJSON {
	selfId := md.UserId
	phone := s.makePhoneNumber(ctx, selfId, r.Phone)
	if phone == "" {
		return &helper.ResultJSON{Code: -1, Msg: "phone number is invalid"}
	}
	users, err := s.UserFacade.GetUserByPhoneNumber(ctx, selfId, phone)
	if users == nil || err != nil {
		return &helper.ResultJSON{Code: -2, Msg: "search user fail"}
	}
	isContact, _ := s.UserFacade.GetContactAndMutual(ctx, users.Id, selfId)
	if !s.UserFacade.CheckPrivacy(ctx, model.ADDED_BY_PHONE, users.Id, selfId, isContact) {
		return &helper.ResultJSON{Code: -3, Msg: "you don't have permission to view this user"}
	}
	var data = struct {
		Uid uint32 `json:"uId"`
	}{
		Uid: uint32(users.Id),
	}
	return &helper.ResultJSON{Code: 200, Msg: "success", Data: data}
}

func (s *cls) QueryPrivacySettings(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TUsersInfo) *helper.ResultJSON {
	if len(r.Uids) == 0 {
		return &helper.ResultJSON{Code: -1, Msg: "please post user id list"}
	}
	if len(r.Uids) > 500 {
		return &helper.ResultJSON{Code: -2, Msg: "user id count too much"}
	}
	var data = make([]struct {
		Uid         uint32 `json:"uId"`
		IsBlacklist bool   `json:"isBlacklist"`
	}, len(r.Uids))

	for i, uid := range r.Uids {
		data[i].Uid = uid
		data[i].IsBlacklist = s.UserFacade.IsBlockedByUser(ctx, int32(uid), md.UserId)
	}

	return &helper.ResultJSON{Code: 200, Msg: "success", Data: data}
}

func (s *cls) SetUserInfoExt(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TSetUserInfoExt) *helper.ResultJSON {
	ok, err := s.UpdateUserInfoExt(ctx, md.UserId, r.Gender, r.Birth, r.Country, r.CountryCode, r.Province, r.City, r.CityCode)
	if err != nil || !ok {
		return &helper.ResultJSON{Code: -1, Msg: "update user info ext failed"}
	}
	return &helper.ResultJSON{Code: 200, Msg: "success"}
}

func (s *cls) GetUserInfoExt(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TGetUserInfoExt) *helper.ResultJSON {
	uid := r.UserId
	if uid == 0 {
		uid = md.UserId
	}
	users := s.GetMutableUsers(ctx, uid)
	user, ok := users[uid]
	if !ok {
		return &helper.ResultJSON{Code: -1, Msg: "user not found"}
	}

	var data = struct {
		Gender      int32  `json:"gender"`
		Birth       string `json:"birth"`
		Country     string `json:"country"`
		CountryCode string `json:"countryCode"`
		Province    string `json:"province"`
		City        string `json:"city"`
		CityCode    string `json:"cityCode"`
	}{
		Gender:      user.User.Gender,
		Birth:       user.User.Birth,
		Country:     user.User.Country,
		CountryCode: user.User.CountryCode,
		Province:    user.User.Province,
		City:        user.User.City,
		CityCode:    user.User.CityCode,
	}

	return &helper.ResultJSON{Code: 200, Msg: "success", Data: data}
}
