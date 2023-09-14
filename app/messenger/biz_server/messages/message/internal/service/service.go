package service

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"open.chat/app/messenger/webpage/webpagepb"
	"open.chat/app/pkg/env2"
	"open.chat/app/service/auth_session/authsessionpb"
	authsession_client "open.chat/app/service/auth_session/client"
	"open.chat/pkg/grpc_util/client"

	msg_facade "open.chat/app/messenger/msg/facade"
	sync_client "open.chat/app/messenger/sync/client"
	auth_facade "open.chat/app/service/biz_service/auth/facade"
	channel_facade "open.chat/app/service/biz_service/channel/facade"
	chat_facade "open.chat/app/service/biz_service/chat/facade"
	message_facade "open.chat/app/service/biz_service/message/facade"
	poll_facade "open.chat/app/service/biz_service/poll/facade"
	private_facade "open.chat/app/service/biz_service/private/facade"
	user_client "open.chat/app/service/biz_service/user/client"
	username_facade "open.chat/app/service/biz_service/username/facade"
)

type Service struct {
	user_client.UserFacade
	username_facade.UsernameFacade
	msg_facade.MsgFacade
	private_facade.PrivateFacade
	message_facade.MessageFacade
	chat_facade.ChatFacade
	channel_facade.ChannelFacade
	poll_facade.PollFacade
	webpagepb.RPCWebPageClient
	auth_facade.AuthFacade
	AuthSessionRpcClient authsessionpb.RPCSessionClient
}

type sendContactByPhone struct {
	Uid uint32 `json:"userId,omitempty"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func New() *Service {
	var (
		ac struct {
			Wardenclient *warden.ClientConfig
		}
		err error
		s   = new(Service)
	)

	checkErr(paladin.Get("application.toml").UnmarshalTOML(&ac))

	s.UserFacade, err = user_client.NewUserFacade("local")
	checkErr(err)
	s.UsernameFacade, err = username_facade.NewUsernameFacade("local")
	checkErr(err)
	s.MsgFacade, err = msg_facade.NewMsgFacade("emsg")
	checkErr(err)
	s.MessageFacade, err = message_facade.NewMessageFacade("local")
	checkErr(err)
	s.PrivateFacade, err = private_facade.NewPrivateFacade("local")
	checkErr(err)
	s.ChatFacade, err = chat_facade.NewChatFacade("local")
	checkErr(err)
	s.ChannelFacade, err = channel_facade.NewChannelFacade("local")
	checkErr(err)
	s.PollFacade, err = poll_facade.NewPollFacade("local")
	checkErr(err)
	s.AuthFacade, err = auth_facade.NewAuthFacade("local")
	checkErr(err)
	s.AuthSessionRpcClient, err = authsession_client.New(ac.Wardenclient)
	checkErr(err)

	if conn, err2 := client.NewClient(env2.MessengerWebPageId, ac.Wardenclient); err2 != nil {
		checkErr(err2)
	} else {
		s.RPCWebPageClient = webpagepb.NewRPCWebPageClient(conn)
	}

	sync_client.New()
	return s
}

func (s *Service) Close() {

}
