package service

import (
	"context"
	"strconv"
	"time"

	"open.chat/app/messenger/push/internal/dao/apns2"
	"open.chat/pkg/crypto"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

var (
	bundID = "bh.tg.ios"
)

func (s *Service) onPushAPNS(ctx context.Context, pushAuthKeyId int64, token string, secret []byte, pushMsg *PushMessage) error {
	if s.apnsClient == nil {
		return nil
	}

	titel := "push"
	if len(pushMsg.LocArgs) > 0 {
		titel = pushMsg.LocArgs[0]
	}
	aps := apns2.Aps{
		Alert: apns2.Alert{
			Title: titel,
			Body:  pushMsg.Message,
		},
		Badge:          int(pushMsg.Badge),
		Sound:          pushMsg.Sound,
		MutableContent: 1,
	}

	payload := &apns2.Payload{
		Aps:    aps,
		TaskID: crypto.RandomString(32),
		FromId: strconv.Itoa(int(pushMsg.Custom.FromId)),
		MsgId:  strconv.Itoa(int(pushMsg.Custom.MsgId)),
	}
	log.Infof("apns:[%s][%s]", token, logger.JsonDebugData(payload))
	resp, err := s.apnsClient.Push(token, payload, time.Now().Unix())
	if err != nil {
		log.Errorf("error: %s", err.Error())
	} else {
		log.Infof("recv:%s", logger.JsonDebugData(resp))
	}
	return nil
}
