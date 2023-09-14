package service

import (
	"context"
	"open.chat/app/interface/botway/botapi"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

func (s *Service) SendGame(ctx context.Context, token string, r *botapi.SendGame2) (*botapi.Message, error) {
	log.Warnf("sendGame - method not impl")
	return nil, mtproto.ErrMethodNotImpl
}

func (s *Service) SetGameScore(ctx context.Context, token string, r *botapi.SetGameScore2) (*botapi.Message, error) {
	log.Warnf("setGameScore - method not impl")
	return nil, mtproto.ErrMethodNotImpl
}

func (s *Service) GetGameHighScores(ctx context.Context, token string, r *botapi.GetGameHighScores2) ([]*botapi.GameHighScore, error) {
	log.Warnf("getGameHighScores - method not impl")
	return nil, mtproto.ErrMethodNotImpl
}
