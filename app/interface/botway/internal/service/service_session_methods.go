package service

import (
	"context"
	"open.chat/app/interface/gateway/egatepb"
	"open.chat/mtproto"
)

func (s *Service) ReceiveData(ctx context.Context, r *egatepb.SessionRawData) (reply *mtproto.Bool, err error) {
	return mtproto.ToBool(true), nil
}
