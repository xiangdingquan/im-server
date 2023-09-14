package http

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"open.chat/app/interface/botway/botapi"
)

func sendSticker(c *bm.Context) {
	req := new(botapi.SendSticker2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SendSticker(c, token, req)
	})
}

func getStickerSet(c *bm.Context) {
	req := new(botapi.GetStickerSet2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.GetStickerSet(c, token, req)
	})
}

func uploadStickerFile(c *bm.Context) {
	req := new(botapi.UploadStickerFile2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.UploadStickerFile(c, token, req)
	})
}

func createNewStickerSet(c *bm.Context) {
	req := new(botapi.CreateNewStickerSet2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.CreateNewStickerSet(c, token, req)
	})
}

func addStickerToSet(c *bm.Context) {
	req := new(botapi.AddStickerToSet2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.AddStickerToSet(c, token, req)
	})
}

func setStickerPositionInSet(c *bm.Context) {
	req := new(botapi.SetStickerPositionInSet2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SetStickerPositionInSet(c, token, req)
	})
}

func deleteStickerFromSet(c *bm.Context) {
	req := new(botapi.DeleteStickerFromSet2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.DeleteStickerFromSet(c, token, req)
	})
}
