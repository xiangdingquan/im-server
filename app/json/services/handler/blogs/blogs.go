package blogs

import (
	"context"
	"encoding/json"
	"open.chat/app/json/helper"
	svc "open.chat/app/json/service"
	"open.chat/app/json/service/handler"
	blog_client "open.chat/app/service/biz_service/blog/facade"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

type (
	cls struct {
		blog_client.BlogFacade
	}
)

func New(s *svc.Service) {
	service := &cls{}
	var err error
	service.BlogFacade, err = blog_client.NewBlogFacade("local")
	helper.CheckErr(err)
	s.AppendServices(handler.RegisterBlogs(service))
}

func (s *cls) SetPrivacy(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TBlogsSetPrivacy) *helper.ResultJSON {
	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("blogs.setPrivacy - error: %v", err)
		return &helper.ResultJSON{Code: err.Code(), Msg: err.Message()}
	}

	b, err := json.Marshal(r.Privacy.Rules)
	if err != nil {
		log.Errorf("blogs.setPrivacy, rules not json, error: %v", err)
		return &helper.ResultJSON{Code: -1, Msg: "rules not json"}
	}

	err = s.BlogFacade.SetPrivacy(ctx, md.UserId, r.Privacy.Key, string(b))
	if err != nil {
		log.Errorf("blogs.setPrivacy, save to db failed, error: %v", err)
		return &helper.ResultJSON{Code: -2, Msg: "set failed"}
	}

	return &helper.ResultJSON{Code: 200, Msg: "success"}
}

func (s *cls) ModifyPrivacyUsers(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TBlogsModifyPrivacyUsers) *helper.ResultJSON {
	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("blogs.setPrivacy - error: %v", err)
		return &helper.ResultJSON{Code: err.Code(), Msg: err.Message()}
	}

	err := s.BlogFacade.ModifyPrivacyUsers(ctx, md.UserId, r.Key, r.Users, r.IsAdding)
	if err != nil {
		log.Errorf("blogs.modifyPrivacyUsers error: %v", err)
		return &helper.ResultJSON{Code: -1, Msg: "modify failed"}
	}

	return &helper.ResultJSON{Code: 200, Msg: "success"}
}

func (s *cls) GetAllPrivacy(ctx context.Context, md *grpc_util.RpcMetadata) *helper.ResultJSON {
	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("blogs.setPrivacy - error: %v", err)
		return &helper.ResultJSON{Code: err.Code(), Msg: err.Message()}
	}

	m, err := s.BlogFacade.GetUserPrivacy(ctx, md.UserId)
	if err != nil {
		log.Errorf("blogs.getAllPrivacy(%d), get from facade, error: %v", md.UserId, err)
		return &helper.ResultJSON{Code: -1, Msg: "get failed"}
	}

	out := handler.TBlogsUserPrivacyList{}
	for k, v := range m {
		var rL []handler.TBlogsUserPrivacyRule
		err := json.Unmarshal([]byte(v), &rL)
		if err != nil {
			log.Errorf("blogs.getAllPrivacy(%d), unmarshal (%s), error: %v", md.UserId, v, err)
			return &helper.ResultJSON{Code: -2, Msg: "unmarshal failed"}
		}

		out.PrivacyList = append(out.PrivacyList, handler.TBlogsUserPrivacy{
			Key:   k,
			Rules: rL,
		})
	}

	return &helper.ResultJSON{Code: 200, Msg: "success", Data: out}
}
