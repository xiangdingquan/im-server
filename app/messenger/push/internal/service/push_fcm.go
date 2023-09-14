package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"open.chat/app/messenger/push/internal/dao/fcm"
	"open.chat/pkg/crypto"
)

func (s *Service) onPushFCM(ctx context.Context, pushAuthKeyId int64, token string, secret []byte, pushMsg *PushMessage) error {
	if s.fcmClient == nil {
		return nil
	}

	authKey := crypto.NewAuthKey(0, secret)
	pBase64Val := encodeNotificationMessageData(authKey, pushMsg)
	if pBase64Val == "" {
		return fmt.Errorf("aesIgeEncrypt message error")
	}

	message := &fcm.Message{
		DryRun:         false,
		Data:           map[string]string{"p": pBase64Val},
		To:             token,
		Priority:       fcm.PriorityHigh,
		DelayWhileIdle: true,
		CollapseKey:    "",
		TimeToLive:     int(time.Hour.Seconds()),
		Android:        fcm.Android{Priority: fcm.PriorityHigh},
	}
	response, err := s.fcmClient.Send(message)
	msgB, _ := json.Marshal(message)
	fmt.Printf("msg(%s)\n", msgB)
	if err != nil {
		fmt.Printf("fcm send response(%+v) error(%v)\n", response, err)
	} else {
		fmt.Println("Status Code   :", response.StatusCode)
		fmt.Println("Success       :", response.Success)
		fmt.Println("Fail          :", response.Fail)
		fmt.Println("Canonical_ids :", response.CanonicalIDs)
		fmt.Println("Topic MsgId   :", response.MsgID)
	}

	return nil
}
