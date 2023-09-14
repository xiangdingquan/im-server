package system

import (
	"context"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"open.chat/app/service/auth_session/authsessionpb"
	authsession_client "open.chat/app/service/auth_session/client"
	"open.chat/app/sysconfig"
	"open.chat/model"
	"time"

	"open.chat/app/json/helper"
	svc "open.chat/app/json/service"
	"open.chat/pkg/grpc_util"

	"open.chat/app/json/service/handler"
	user_client "open.chat/app/service/biz_service/user/client"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type cls struct {
	user_client.UserFacade
	AuthSessionRpcClient authsessionpb.RPCSessionClient
}

// New .
func New(s *svc.Service) {
	c := new(cls)
	var err error
	c.UserFacade, err = user_client.NewUserFacade("local")
	checkErr(err)

	var ac struct {
		Wardenclient *warden.ClientConfig
	}
	checkErr(paladin.Get("application.toml").UnmarshalTOML(&ac))
	c.AuthSessionRpcClient, err = authsession_client.New(ac.Wardenclient)
	checkErr(err)

	s.AppendServices(handler.RegisterSystem(c))
}

func (s *cls) SyncTime(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TSyncTime) *helper.ResultJSON {
	var data = struct {
		ClientTime uint64 `json:"client_time"`
		ServerTime uint64 `json:"server_time"`
	}{
		ClientTime: r.ClientTime,
		ServerTime: uint64(time.Now().UTC().UnixNano() / 1e6),
	}
	return &helper.ResultJSON{Code: 200, Msg: "success", Data: data}
}

func (s *cls) RelayData(ctx context.Context, md *grpc_util.RpcMetadata, r *helper.TrelayData) *helper.ResultJSON {
	r.From = (uint32)(md.UserId)
	if r.PushUpdates(ctx, nil) != nil {
		return &helper.ResultJSON{Code: -1, Msg: "notice message fail"}
	}
	return &helper.ResultJSON{Code: 200, Msg: "relay success"}
}

func (s *cls) SendSmsCode(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TSendSmsCode) *helper.ResultJSON {
	me, err := s.UserFacade.GetUserSelf(ctx, md.UserId)
	if err != nil || me == nil {
		return &helper.ResultJSON{Code: -1, Msg: "get self info fail"}
	}

	code, err := helper.SendVerifyCode(ctx, r.Type, me.GetPhone().GetValue(), model.GetLangType(ctx, s.AuthSessionRpcClient))
	if err != nil {
		return &helper.ResultJSON{Code: -2, Msg: "send verify code fail"}
	}

	if code == "" {
		return &helper.ResultJSON{Code: -3, Msg: "send verify code fail"}
	}

	return &helper.ResultJSON{Code: 200, Msg: "success"}
}

func (s *cls) VerifySmsCode(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TVerifySmsCode) *helper.ResultJSON {
	me, err := s.UserFacade.GetUserSelf(ctx, md.UserId)
	if err != nil || me == nil {
		return &helper.ResultJSON{Code: -1, Msg: "get self info fail"}
	}

	if helper.VerifyCode(ctx, r.Type, me.GetPhone().GetValue(), r.Code) != nil {
		return &helper.ResultJSON{Code: 400, Msg: "verify Code is wrong"}
	}

	return &helper.ResultJSON{Code: 200, Msg: "success"}

}

func (s *cls) GetOssConfig(ctx context.Context, md *grpc_util.RpcMetadata) *helper.ResultJSON {
	return &helper.ResultJSON{Code: 200, Msg: "success", Data: struct {
		BucketName      string `json:"bucketName"`
		Endpoint        string `json:"endpoint"`
		AccessKeyId     string `json:"accessKeyId"`
		AccessKeySecret string `json:"accessKeySecret"`
	}{
		BucketName:      sysconfig.GetConfig2String(ctx, sysconfig.ConfigKeyOssBucketName, "", 0),
		Endpoint:        sysconfig.GetConfig2String(ctx, sysconfig.ConfigKeyOssEndpoint, "", 0),
		AccessKeyId:     sysconfig.GetConfig2String(ctx, sysconfig.ConfigKeyOssAccessKeyId, "", 0),
		AccessKeySecret: sysconfig.GetConfig2String(ctx, sysconfig.ConfigKeyOssAccessKeySecret, "", 0),
	}}
}
