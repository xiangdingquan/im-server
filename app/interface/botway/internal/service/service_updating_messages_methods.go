package service

import (
	"context"
	"github.com/gogo/protobuf/types"
	"open.chat/app/interface/botway/botapi"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

func (s *Service) EditMessageText(ctx context.Context, token string, req *botapi.EditMessageText2) (message *botapi.Message, err error) {
	var (
		resp mtproto.TLObject
		me   *mtproto.Updates
		ok   bool
	)

	i := &mtproto.TLMessagesEditMessage{
		NoWebpage:    req.DisableWebPagePreview,
		Peer:         MakePeer(req.ChatId),
		Id:           req.MessageId,
		Message:      &types.StringValue{Value: req.Text},
		ReplyMarkup:  nil,
		Entities:     nil,
		ScheduleDate: nil,
	}

	if req.ReplyMarkup != nil {
		i.ReplyMarkup = encodeToReplyMarkup(&botapi.ReplyMarkup{InlineKeyboard: req.ReplyMarkup.InlineKeyboard})
	}

	log.Debugf("send: %s", i.DebugString())
	resp, err = s.Invoke(ctx, token, i)
	if err != nil {
		return
	}

	if me, ok = resp.(*mtproto.Updates); !ok {
		err = mtproto.ErrInternelServerError
		log.Errorf("invalid error")
		return
	}

	message = ToMessage(me)
	return
}

func (s *Service) EditMessageCaption(ctx context.Context, token string, req *botapi.EditMessageCaption2) (*botapi.Message, error) {
	log.Warnf("editMessageCaption - method not impl")
	return nil, mtproto.ErrMethodNotImpl
}

func (s *Service) EditMessageMedia(ctx context.Context, token string, req *botapi.EditMessageMedia2) (*botapi.Message, error) {
	log.Warnf("editMessageMedia - method not impl")
	return nil, mtproto.ErrMethodNotImpl
}

func (s *Service) EditMessageReplyMarkup(ctx context.Context, token string, req *botapi.EditMessageReplyMarkup2) (*botapi.Message, error) {
	log.Warnf("editMessageReplyMarkup - method not impl")
	return nil, mtproto.ErrMethodNotImpl
}

func (s *Service) StopPoll(ctx context.Context, token string, req *botapi.StopPoll2) (*botapi.Poll, error) {
	log.Warnf("stopPoll - method not impl")
	return nil, mtproto.ErrMethodNotImpl
}

func (s *Service) DeleteMessage(ctx context.Context, token string, req *botapi.DeleteMessage2) (bool, error) {
	log.Warnf("deleteMessage - method not impl")
	return false, mtproto.ErrMethodNotImpl
}
