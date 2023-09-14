package service

import (
	"context"
	"open.chat/app/interface/botway/botapi"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

func (s *Service) SendInvoice(ctx context.Context, token string, req *botapi.SendInvoice2) (*botapi.Message, error) {
	log.Warnf("sendInvoice - method not impl")
	return nil, mtproto.ErrMethodNotImpl
}

func (s *Service) AnswerShippingQuery(ctx context.Context, token string, req *botapi.AnswerShippingQuery2) (bool, error) {
	log.Warnf("answerShippingQuery - method not impl")
	return false, mtproto.ErrMethodNotImpl
}

func (s *Service) AnswerPreCheckoutQuery(ctx context.Context, token string, req *botapi.AnswerPreCheckoutQuery2) (bool, error) {
	log.Warnf("answerPreCheckoutQuery - method not impl")
	return false, mtproto.ErrMethodNotImpl
}
