package service

import (
	"context"
	"math/rand"

	"github.com/gogo/protobuf/types"

	"open.chat/app/bots/botpb"
	"open.chat/app/service/media/client"
	"open.chat/mtproto"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) MessagesGetInlineBotResults(ctx context.Context, r *botpb.GetInlineBotResults) (*mtproto.Messages_BotResults, error) {
	log.Debugf("messages.getInlineBotResults - request: {%s}", logger.JsonDebugData(r))

	if r.Offset != "" {
		return mtproto.MakeTLMessagesBotResults(&mtproto.Messages_BotResults{
			Gallery:    true,
			QueryId:    rand.Int63(),
			NextOffset: nil,
			SwitchPm:   nil,
			Results:    []*mtproto.BotInlineResult{},
			CacheTime:  3600,
			Users:      []*mtproto.User{},
		}).To_Messages_BotResults(), nil
	}

	gifDOList, err := s.Dao.GiphyDatasDAO.SelectAll(ctx)
	if err != nil {
		log.Errorf("messages.getInlineBotResults - error: %v", err)
		return nil, err
	}

	mediaResults := make([]*mtproto.BotInlineResult, 0, len(gifDOList))
	for i := 0; i < len(gifDOList); i++ {
		photo := media_client.GetPhoto(gifDOList[i].PhotoId)
		gif, _ := media_client.GetDocumentByIdList([]int64{gifDOList[i].DocumentId})
		if photo == nil || len(gif) == 0 {
			continue
		}

		mediaResults = append(mediaResults, mtproto.MakeTLBotInlineMediaResult(&mtproto.BotInlineResult{
			Id:          gifDOList[i].GiphyId,
			Type:        "gif",
			Photo:       photo,
			Document:    gif[0],
			Title:       nil,
			Description: nil,
			SendMessage: mtproto.MakeTLBotInlineMessageMediaAuto(&mtproto.BotInlineMessage{
				Message:     "",
				Entities:    nil,
				ReplyMarkup: nil,
			}).To_BotInlineMessage(),
		}).To_BotInlineResult())
	}

	botResults := mtproto.MakeTLMessagesBotResults(&mtproto.Messages_BotResults{
		Gallery:    true,
		QueryId:    rand.Int63(),
		NextOffset: &types.StringValue{Value: "1"},
		SwitchPm:   nil,
		Results:    mediaResults,
		CacheTime:  3600,
		Users:      []*mtproto.User{},
	}).To_Messages_BotResults()

	log.Debugf("messages.getInlineBotResults - result: %s", botResults.DebugString())
	return botResults, nil
}

func (s *Service) MessagesGetBotCallbackAnswer(ctx context.Context, r *botpb.GetBotCallbackAnswer) (*mtproto.Messages_BotCallbackAnswer, error) {
	log.Debugf("messages.getBotCallbackAnswer - request: {%s}", logger.JsonDebugData(r))

	err := mtproto.ErrMethodNotImpl

	log.Errorf("messages.getBotCallbackAnswer - reply: %v", err)
	return nil, err
}

func (s *Service) MessagesQueryInlineBotResult(ctx context.Context, r *botpb.QueryInlineBotResult) (*mtproto.BotInlineResult, error) {
	log.Debugf("messages.queryInlineBotResult - request: {%s}", logger.JsonDebugData(r))

	giphyDO, err := s.Dao.GiphyDatasDAO.SelectById(ctx)
	if err != nil {
		log.Errorf("messages.queryInlineBotResult - error: %v", err)
		return nil, err
	} else if giphyDO == nil {
		err := mtproto.ErrQueryIdEmpty
		log.Errorf("messages.queryInlineBotResult - error: %v", err)
		return nil, err
	}

	photo := media_client.GetPhoto(giphyDO.PhotoId)
	gif, _ := media_client.GetDocumentByIdList([]int64{giphyDO.DocumentId})
	if photo == nil || len(gif) == 0 {
		err := mtproto.ErrQueryIdEmpty
		log.Errorf("messages.queryInlineBotResult - error: %v", err)
		return nil, err
	}

	botInlineResult := mtproto.MakeTLBotInlineMediaResult(&mtproto.BotInlineResult{
		Id:          giphyDO.GiphyId,
		Type:        "gif",
		Photo:       photo,
		Document:    gif[0],
		Title:       nil,
		Description: nil,
		SendMessage: mtproto.MakeTLBotInlineMessageMediaAuto(&mtproto.BotInlineMessage{
			Message:     "",
			Entities:    nil,
			ReplyMarkup: nil,
		}).To_BotInlineMessage(),
	}).To_BotInlineResult()

	log.Debugf("messages.queryInlineBotResult - result: %s", botInlineResult.Description)
	return nil, err
}
