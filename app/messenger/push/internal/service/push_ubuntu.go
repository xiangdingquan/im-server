package service

import (
	"context"

	"open.chat/pkg/log"
)

func (s *Service) onPushUbuntu(ctx context.Context, pushAuthKeyId int64, token string, secret []byte, pushMsg *PushMessage) error {
	log.Warnf("not impl")

	return nil
}
