package service

import (
	"context"
	"open.chat/app/interface/botway/botapi"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

func (s *Service) AnswerInlineQuery(ctx context.Context, token string, req *botapi.AnswerInlineQuery2) (bool, error) {
	log.Warnf("answerInlineQuery - method not impl")
	return false, mtproto.ErrMethodNotImpl
}
