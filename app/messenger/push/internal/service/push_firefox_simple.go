package service

import (
	"context"
)

func (s *Service) onPushDeprecated(ctx context.Context, pushAuthKeyId int64, token string, secret []byte, pushMsg *PushMessage) error {
	s.onPushJPush(ctx, pushAuthKeyId, token, secret, pushMsg)
	return nil
}
