package pusher

import (
	"context"
	"math/rand"
	"time"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"open.chat/app/json/helper"
	svc "open.chat/app/json/service"
	"open.chat/app/json/service/http"
	"open.chat/mtproto"
	"open.chat/pkg/log"

	"open.chat/app/messenger/biz_server/account"
	msg_facade "open.chat/app/messenger/msg/facade"
	"open.chat/model"
)

const (
	system_uid int32 = 777000
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type cls struct {
	msg_facade.MsgFacade
	srv mtproto.RPCAccountServer
}

// New .
func New(s *svc.Service, rg *bm.RouterGroup) {
	var err error
	service := &cls{
		srv: account.New(),
	}
	service.MsgFacade, err = msg_facade.NewMsgFacade("emsg")
	checkErr(err)
	http.RegisterPusher(service, rg)
}

func (s *cls) send(ctx context.Context, isSend bool, r *http.TPusherMessage) *helper.ResultJSON {
	if len(r.Uids) == 0 {
		return &helper.ResultJSON{Code: -1, Msg: "please input uids"}
	}
	if len(r.Msg) == 0 {
		return &helper.ResultJSON{Code: -2, Msg: "please input message content"}
	}

	go func(uids []uint32, msg string) {
		var pushType int32 = 1
		if !isSend {
			pushType = 0
		}
		for _, uid := range uids {
			message := mtproto.MakeTLMessage(&mtproto.Message{
				Out:             true,
				Date:            int32(time.Now().Unix()),
				FromId_FLAGPEER: model.MakePeerUser(system_uid),
				ToId:            model.MakePeerUser(int32(uid)),
				Message:         msg,
				Entities:        []*mtproto.MessageEntity{},
			}).To_Message()
			err := s.MsgFacade.PushUserMessage(ctx, pushType, system_uid, int32(uid), rand.Int63(), message)
			if err != nil {
				log.Errorf("PushUserMessage:%s", err.Error())
			}
		}
	}(r.Uids, r.Msg)
	return &helper.ResultJSON{Code: 0, Msg: "success"}
}

func (s *cls) Messages(ctx context.Context, r *http.TPusherMessage) *helper.ResultJSON {
	return s.send(ctx, false, r)
}

func (s *cls) Notification(ctx context.Context, r *http.TPusherMessage) *helper.ResultJSON {
	return s.send(ctx, true, r)
}
