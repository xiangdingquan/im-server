package http

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"open.chat/app/interface/botway/botapi"
)

func sendGame(c *bm.Context) {
	req := new(botapi.SendGame2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SendGame(c, token, req)
	})
}

func setGameScore(c *bm.Context) {
	req := new(botapi.SetGameScore2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SetGameScore(c, token, req)
	})
}

func getGameHighScores(c *bm.Context) {
	req := new(botapi.GetGameHighScores2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.GetGameHighScores(c, token, req)
	})
}
