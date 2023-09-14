package service

import (
	"context"

	"open.chat/app/messenger/push/internal/dao/jpush"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) onPushJPush(ctx context.Context, pushAuthKeyId int64, token string, secret []byte, pushMsg *PushMessage) error {
	if s.jpushClient == nil {
		return nil
	}

	titel := "通知"
	if len(pushMsg.LocArgs) > 0 {
		titel = pushMsg.LocArgs[0]
	}

	payload := &jpush.Payload{
		Platform: "all",
		Audience: jpush.Audience{
			RegistrationId: []string{token},
		},
		Notification: jpush.Notification{
			Android: jpush.Android{
				Title:       titel,
				Alert:       pushMsg.Message,
				Sound:       pushMsg.Sound,
				BadgeAddNum: 1,
			},
		},
	}
	log.Infof("jpush:[%s][%s]", token, logger.JsonDebugData(payload))
	response, err := s.jpushClient.Push(payload)
	if err != nil {
		log.Errorf("jpush send response(%+v) error(%v)\n", response, err)
	} else {
		log.Infof("jpush recv:%s", logger.JsonDebugData(response))
	}
	return nil
}
