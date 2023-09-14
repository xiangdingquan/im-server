package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountSetAccountTTL(ctx context.Context, request *mtproto.TLAccountSetAccountTTL) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.setAccountTTL - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var err error

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("account.setAccountTTL - error: %v", err)
		return nil, err
	}

	// Check ttl
	ttl := request.GetTtl().GetDays()
	switch ttl {
	case 30:
	case 90:
	case 180:
	case 365:
	default:
		err = mtproto.ErrTtlDaysInvalid
		log.Errorf("account.setAccountTTL - error: %v", err)
		return nil, err
	}

	if err = s.UserFacade.SetAccountDaysTTL(ctx, md.UserId, ttl); err != nil {
		log.Errorf("account.setAccountTTL - error: %v", err)
		return mtproto.ToBool(false), nil
	}

	log.Debugf("account.setAccountTTL - reply: {true}")
	return mtproto.ToBool(true), nil
}
