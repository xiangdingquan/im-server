package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"open.chat/app/messenger/push/internal/dao/apns2"
	"open.chat/pkg/crypto"
	"open.chat/pkg/log"
)

func (s *Service) onPushAPNSVoIP(ctx context.Context, pushAuthKeyId int64, token string, secret []byte, pushMsg *PushMessage) error {
	log.Warnf("onPushAPNSVoIP not impl")

	if s.apnsClient == nil {
		return nil
	}

	authKey := crypto.NewAuthKey(0, secret)
	pBase64Val := encodeNotificationMessageData(authKey, pushMsg)
	if pBase64Val == "" {
		return fmt.Errorf("aesIgeEncrypt message error")
	}

	aps := apns2.Aps{
		Alert: apns2.Alert{
			Title: "Your fruit",
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
		P:      pBase64Val,
	}
	bs, _ := json.Marshal(payload)
	fmt.Printf("payload(%s)", bs)
	resp, err := s.apnsClient.Push(token, payload, time.Now().Unix())
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("StatusCode:", resp.StatusCode, "ApnsID:", resp.ApnsID, "Reason:", resp.Reason)

	return nil
}
