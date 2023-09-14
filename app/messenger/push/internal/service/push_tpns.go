package service

import (
	"context"

	"open.chat/app/messenger/push/internal/dao/tpns"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) onPushTPNS(ctx context.Context, pushAuthKeyId int64, token string, secret []byte, pushMsg *PushMessage) error {
	if s.tpnsClient == nil {
		return nil
	}

	titel := "通知"
	if len(pushMsg.LocArgs) > 0 {
		titel = pushMsg.LocArgs[0]
	}

	payload := &tpns.Payload{
		AudienceType: "token",
		TokenList:    []string{token},
		MessageType:  "notify",
		Message: tpns.Message{
			Title:   titel,
			Content: pushMsg.Message,
		},
	}
	log.Infof("tpns:[%s][%s]", token, logger.JsonDebugData(payload))
	response, err := s.tpnsClient.Push(payload)
	if err != nil {
		log.Errorf("tpns send response(%+v) error(%v)\n", response, err)
	} else {
		log.Infof("tpns recv:%s", logger.JsonDebugData(response))
	}
	return nil
}
