package service

import (
	"context"
	"open.chat/app/interface/botway/botapi"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

func (s *Service) SetPassportDataErrors(ctx context.Context, token string, req *botapi.SetPassportDataErrors2) (bool, error) {
	log.Warnf("setPassportDataErrors - method not impl")
	return false, mtproto.ErrMethodNotImpl
}
