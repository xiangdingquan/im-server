package service

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/net/rpc/warden"

	admin_log_client "open.chat/app/job/admin_log/client"
	msg_facade "open.chat/app/messenger/msg/facade"
	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/app/service/auth_session/authsessionpb"
	authsession_client "open.chat/app/service/auth_session/client"
	channel_facade "open.chat/app/service/biz_service/channel/facade"
	chat_facade "open.chat/app/service/biz_service/chat/facade"
	message_facade "open.chat/app/service/biz_service/message/facade"
	report_facade "open.chat/app/service/biz_service/report/facade"
	user_client "open.chat/app/service/biz_service/user/client"
	username_facade "open.chat/app/service/biz_service/username/facade"
)

type Service struct {
	user_client.UserFacade
	username_facade.UsernameFacade
	msg_facade.MsgFacade
	report_facade.ReportFacade
	chat_facade.ChatFacade
	channel_facade.ChannelFacade
	message_facade.MessageFacade

	*admin_log_client.AdminLogClient
	authsessionpb.RPCSessionClient
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

	// s.Dao = dao.New()
	s.UserFacade, err = user_client.NewUserFacade("local")
	checkErr(err)
	s.UsernameFacade, err = username_facade.NewUsernameFacade("local")
	checkErr(err)
	s.MsgFacade, err = msg_facade.NewMsgFacade("emsg")
	checkErr(err)
	s.ReportFacade, err = report_facade.NewReportFacade("local")
	checkErr(err)
	s.ChatFacade, err = chat_facade.NewChatFacade("local")
	checkErr(err)
	s.ChannelFacade, err = channel_facade.NewChannelFacade("local")
	checkErr(err)
	s.MessageFacade, err = message_facade.NewMessageFacade("local")
	checkErr(err)
	s.RPCSessionClient, err = authsession_client.New(ac.Wardenclient)
	checkErr(err)

	s.AdminLogClient = admin_log_client.New()
	sync_client.New()

	return s
}

func (s *Service) Close() {

}
