package service

import (
	"context"
	"crypto/tls"
	"math/rand"
	"strings"
	"time"

	"github.com/go-kratos/kratos/pkg/conf/env"
	"github.com/go-kratos/kratos/pkg/conf/paladin"

	"open.chat/app/admin/api_server/internal/dao"
	"open.chat/app/messenger/msg/facade"
	"open.chat/app/service/dfs/facade"
	_ "open.chat/app/service/dfs/facade/dfs"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/http_client"
	"open.chat/pkg/log"
)

type Service struct {
	c *Config
	*dao.Dao
	msg_facade.MsgFacade
	Authorization string
	dfs_facade.DfsFacade
}

func New() (s *Service) {
	var (
		ac  = &Config{}
		err error
	)

	if err = paladin.Get("application.toml").UnmarshalTOML(&ac); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}

	s = &Service{
		c:             ac,
		Authorization: ac.Key + ":" + ac.Secret,
	}

	s.Dao = dao.New(ac.WardenClient)
	if s.MsgFacade, err = msg_facade.NewMsgFacade("emsg"); err != nil {
		panic(err)
	}
	s.DfsFacade, err = dfs_facade.NewDfsFacade("dfs")
	if err != nil {
		panic(err)
	}

	return s
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context) (err error) {
	return nil
}

// Close close the resource.
func (s *Service) Close() {
}

func (s *Service) Invoke2(ctx context.Context, authId int64, id int32, r mtproto.TLObject) (mtproto.TLObject, error) {
	var (
		err       error
		rpcResult mtproto.TLObject
	)

	rpcMetadata := &grpc_util.RpcMetadata{
		ServerId:    env.Hostname,
		ClientAddr:  env.Hostname,
		AuthId:      authId,
		SessionId:   0,
		ReceiveTime: time.Now().Unix(),
		UserId:      id,
		ClientMsgId: 0,
		IsBot:       false,
		Layer:       111,
		IsAdmin:     true,
	}

	log.Debugf("rpc_request: {%s}", model.TLObjectToJson(r))

	rpcResult, err = s.Dao.Invoke(ctx, rpcMetadata, r)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	log.Debugf("rpc_result: {%s}", model.TLObjectToJson(rpcResult))
	return rpcResult, nil
}

func (s *Service) Invoke(ctx context.Context, id int32, r mtproto.TLObject) (mtproto.TLObject, error) {
	return s.Invoke2(ctx, 0, id, r)
}

func (s *Service) Auth(authorization string) bool {
	return authorization == s.Authorization
}

func urlToInputFile(
	fileUrl string,
	cb func(fileId int64, filePart int32, bytes []byte) error) (*mtproto.InputFile, error) {
	var (
		fileUpload []byte
		err        error
	)

	if strings.HasSuffix(fileUrl, "https://") {
		fileUpload, err = http_client.Get(fileUrl).
			SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
			SetTimeout(5*time.Second, 60*time.Second).
			Bytes()
	} else {
		fileUpload, err = http_client.Get(fileUrl).
			SetTimeout(5*time.Second, 60*time.Second).
			Bytes()
	}
	if err != nil {
		log.Errorf("getFileError: %v", err)
	}
	file := mtproto.MakeTLInputFile(&mtproto.InputFile{
		Id:          rand.Int63(),
		Parts:       0,
		Name:        fileUrl,
		Md5Checksum: "",
	}).To_InputFile()

	for i := 0; i < len(fileUpload); i = i + 512*1024 {
		if i+512*1024 > len(fileUpload) {
			file.Parts = int32(i + 1)
			cb(file.Id, int32(i), fileUpload[i:len(fileUpload)-i])
		} else {
			cb(file.Id, int32(i), fileUpload[i:i+512*1024])
		}
	}

	return file, nil
}
