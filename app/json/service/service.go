package service

import (
	"context"
	"fmt"
	"strings"

	"open.chat/app/json/helper"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

type (
	// Services .
	Service struct {
		//methods []methodHandlers
		methods []map[string]helper.MethodHandler
	}
)

var (
	service *Service = nil
	New              = instance
)

func instance() *Service {
	if service == nil {
		service = new(Service)
	}
	return service
}

func (s *Service) AppendServices(ms map[string]helper.MethodHandler) {
	s.methods = append(s.methods, ms)
}

// TestJSON .
func (s *Service) TestJSON(uID uint32, customMethod string, strJSON string) (string, error) { //136817694
	ret, err := s.Handle(context.Background(), &grpc_util.RpcMetadata{UserId: int32(uID)}, customMethod, &mtproto.DataJSON{
		Data: strJSON,
	})
	return ret.GetData(), err
}

func (s *Service) webRequest(ctx context.Context, md *grpc_util.RpcMetadata, api []string, request *helper.DataJSON) *helper.ResultJSON {
	body := &struct {
		UID    uint32      `json:"uid"`
		AuthId int64       `json:"authId"`
		Data   interface{} `json:"data"`
	}{
		UID:    uint32(md.UserId),
		AuthId: md.AuthId,
	}
	err := request.GetJSONData(&body.Data)
	if err != nil {
		return &helper.ResultJSON{Code: -1, Msg: "post json data is wrong"}
	}
	result := &helper.ResultJSON{}
	result.Code, result.Msg, err = helper.WebInterface(strings.Join(api, "/"), body, &result.Data)
	if err != nil {
		result.Code = 404
		result.Msg = "web server request fail"
	}
	return result
}

// Handle .
func (s *Service) Handle(ctx context.Context, md *grpc_util.RpcMetadata, method string, request *mtproto.DataJSON) (*mtproto.DataJSON, error) {
	log.Infof("request:%s[%s]", method, request.Data)
	if request.Data == "" {
		request.Data = "{}"
	}
	api := strings.Split(method, ".")
	if len(api) > 1 && api[0] == "web" {
		data := s.webRequest(ctx, md, api[1:], &helper.DataJSON{DataJSON: request})
		if data == nil {
			data = &helper.ResultJSON{Code: 0, Msg: "error"}
		}
		return data.ToDataJSON().DataJSON, nil
	} else {
		for _, m := range s.methods {
			fn, ok := m[method]
			if !ok {
				continue
			}
			data, err := fn(ctx, md, &helper.DataJSON{DataJSON: request})
			if err == nil {
				log.Infof("response:[%s]", data.ToDataJSON().GetData())
			} else {
				log.Infof("response:[%s]", err.Error())
				if data == nil {
					data = &helper.ResultJSON{Code: 0, Msg: "error", Data: err.Error()}
				}
			}
			return data.ToDataJSON().DataJSON, err
		}
		log.Errorf("method not exist:%s,%s", method, request.Data)
		return nil, fmt.Errorf("method not exist")
	}
}
