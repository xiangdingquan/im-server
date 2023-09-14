package http

import (
	"net/http"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"

	"open.chat/app/interface/botway/internal/service"
	"open.chat/pkg/log"
)

var (
	svc *service.Service
)

func New(s *service.Service) (engine *bm.Engine) {
	var (
		hc struct {
			Server *bm.ServerConfig
		}
	)
	if err := paladin.Get("http.toml").UnmarshalTOML(&hc); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}
	svc = s
	engine = bm.DefaultServer(hc.Server)
	initRouter(engine)
	if err := engine.Start(); err != nil {
		panic(err)
	}
	return
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/bot")
	{
		g.GET("/getUpdates/:token", getUpdates)
		g.POST("/getUpdates/:token", getUpdates)
		g.GET("/setWebhook/:token", setWebhook)
		g.POST("/setWebhook/:token", setWebhook)
		g.GET("/deleteWebhook/:token", deleteWebhook)
		g.POST("/deleteWebhook/:token", deleteWebhook)
		g.GET("/getWebhookInfo/:token", getWebhookInfo)
		g.POST("/getWebhookInfo/:token", getWebhookInfo)
		g.GET("/getMe/:token", getMe)
		g.POST("/getMe/:token", getMe)
		g.GET("/sendMessage/:token", sendMessage)
		g.POST("/sendMessage/:token", sendMessage)
		g.GET("/forwardMessage/:token", forwardMessage)
		g.POST("/forwardMessage/:token", forwardMessage)
		g.GET("/sendPhoto/:token", sendPhoto)
		g.POST("/sendPhoto/:token", sendPhoto)
		g.GET("/sendAudio/:token", sendAudio)
		g.POST("/sendAudio/:token", sendAudio)
		g.GET("/sendDocument/:token", sendDocument)
		g.POST("/sendDocument/:token", sendDocument)
		g.GET("/sendVideo/:token", sendVideo)
		g.POST("/sendVideo/:token", sendVideo)
		g.GET("/sendAnimation/:token", sendAnimation)
		g.POST("/sendAnimation/:token", sendAnimation)
		g.GET("/sendVoice/:token", sendVoice)
		g.POST("/sendVoice/:token", sendVoice)
		g.GET("/sendVideoNote/:token", sendVideoNote)
		g.POST("/sendVideoNote/:token", sendVideoNote)
		g.GET("/sendMediaGroup/:token", sendMediaGroup)
		g.POST("/sendMediaGroup/:token", sendMediaGroup)
		g.GET("/sendLocation/:token", sendLocation)
		g.POST("/sendLocatio/:tokenn", sendLocation)
		g.GET("/editMessageLiveLocation/:token", editMessageLiveLocation)
		g.POST("/editMessageLiveLocation/:token", editMessageLiveLocation)
		g.GET("/stopMessageLiveLocation/:token", stopMessageLiveLocation)
		g.POST("/stopMessageLiveLocation/:token", stopMessageLiveLocation)
		g.GET("/sendVenue/:token", sendVenue)
		g.POST("/sendVenue/:token", sendVenue)
		g.GET("/sendContact/:token", sendContact)
		g.POST("/sendContact/:token", sendContact)
		g.GET("/sendPoll/:token", sendPoll)
		g.POST("/sendPoll/:token", sendPoll)
		g.GET("/sendChatAction/:token", sendChatAction)
		g.POST("/sendChatAction/:token", sendChatAction)
		g.GET("/getUserProfilePhotos/:token", getUserProfilePhotos)
		g.POST("/getUserProfilePhotos/:token", getUserProfilePhotos)
		g.GET("/getFile/:token", getFile)
		g.POST("/getFile/:token", getFile)
		g.GET("/kickChatMember/:token", kickChatMember)
		g.POST("/kickChatMember/:token", kickChatMember)
		g.GET("/unbanChatMember/:token", unbanChatMember)
		g.POST("/unbanChatMember/:token", unbanChatMember)
		g.GET("/restrictChatMember/:token", restrictChatMember)
		g.POST("/restrictChatMember/:token", restrictChatMember)
		g.GET("/promoteChatMember/:token", promoteChatMember)
		g.POST("/promoteChatMember/:token", promoteChatMember)
		g.GET("/setChatPermissions/:token", setChatPermissions)
		g.POST("/setChatPermissions/:token", setChatPermissions)
		g.GET("/exportChatInviteLink/:token", exportChatInviteLink)
		g.POST("/exportChatInviteLink/:token", exportChatInviteLink)
		g.GET("/setChatPhoto/:token", setChatPhoto)
		g.POST("/setChatPhoto/:token", setChatPhoto)
		g.GET("/deleteChatPhoto/:token", deleteChatPhoto)
		g.POST("/deleteChatPhoto/:token", deleteChatPhoto)
		g.GET("/setChatTitle/:token", setChatTitle)
		g.POST("/setChatTitle/:token", setChatTitle)
		g.GET("/setChatDescription/:token", setChatDescription)
		g.POST("/setChatDescription/:token", setChatDescription)
		g.GET("/pinChatMessage/:token", pinChatMessage)
		g.POST("/pinChatMessage/:token", pinChatMessage)
		g.GET("/unpinChatMessage/:token", unpinChatMessage)
		g.POST("/unpinChatMessage/:token", unpinChatMessage)
		g.GET("/leaveChat/:token", leaveChat)
		g.POST("/leaveChat/:token", leaveChat)
		g.GET("/getChat/:token", getChat)
		g.POST("/getChat/:token", getChat)
		g.GET("/getChatAdministrators/:token", getChatAdministrators)
		g.POST("/getChatAdministrators/:token", getChatAdministrators)
		g.GET("/getChatMembersCount/:token", getChatMembersCount)
		g.POST("/getChatMembersCount/:token", getChatMembersCount)
		g.GET("/getChatMember/:token", getChatMember)
		g.POST("/getChatMember/:token", getChatMember)
		g.GET("/setChatStickerSet/:token", setChatStickerSet)
		g.POST("/setChatStickerSet/:token", setChatStickerSet)
		g.GET("/deleteChatStickerSet/:token", deleteChatStickerSet)
		g.POST("/deleteChatStickerSet/:token", deleteChatStickerSet)
		g.GET("/answerCallbackQuery/:token", answerCallbackQuery)
		g.POST("/answerCallbackQuery/:token", answerCallbackQuery)

		g.GET("/editMessageText/:token", editMessageText)
		g.POST("/editMessageText/:token", editMessageText)
		g.GET("/editMessageCaption/:token", editMessageCaption)
		g.POST("/editMessageCaption/:token", editMessageCaption)
		g.GET("/editMessageMedia/:token", editMessageMedia)
		g.POST("/editMessageMedia/:token", editMessageMedia)
		g.GET("/editMessageReplyMarkup/:token", editMessageReplyMarkup)
		g.POST("/editMessageReplyMarkup/:token", editMessageReplyMarkup)
		g.GET("/stopPoll/:token", stopPoll)
		g.POST("/stopPoll/:token", stopPoll)
		g.GET("/deleteMessage/:token", deleteMessage)
		g.POST("/deleteMessage/:token", deleteMessage)

		g.GET("/sendSticker/:token", sendSticker)
		g.POST("/sendSticker/:token", sendSticker)
		g.GET("/getStickerSet/:token", getStickerSet)
		g.POST("/getStickerSet/:token", getStickerSet)
		g.GET("/uploadStickerFile/:token", uploadStickerFile)
		g.POST("/uploadStickerFile/:token", uploadStickerFile)
		g.GET("/createNewStickerSet/:token", createNewStickerSet)
		g.POST("/createNewStickerSet/:token", createNewStickerSet)
		g.GET("/addStickerToSet/:token", addStickerToSet)
		g.POST("/addStickerToSet/:token", addStickerToSet)
		g.GET("/setStickerPositionInSet/:token", setStickerPositionInSet)
		g.POST("/setStickerPositionInSet/:token", setStickerPositionInSet)
		g.GET("/deleteStickerFromSet/:token", deleteStickerFromSet)
		g.POST("/deleteStickerFromSet/:token", deleteStickerFromSet)

		g.GET("/answerInlineQuery/:token", answerInlineQuery)
		g.POST("/answerInlineQuery/:token", answerInlineQuery)

		g.GET("/sendInvoice/:token", sendInvoice)
		g.POST("/sendInvoice/:token", sendInvoice)
		g.GET("/answerShippingQuery/:token", answerShippingQuery)
		g.POST("/answerShippingQuery/:token", answerShippingQuery)
		g.GET("/answerPreCheckoutQuery/:token", answerPreCheckoutQuery)
		g.POST("/answerPreCheckoutQuery/:token", answerPreCheckoutQuery)

		g.GET("/setPassportDataErrors/:token", setPassportDataErrors)
		g.POST("/setPassportDataErrors/:token", setPassportDataErrors)

		g.GET("/sendGame/:token", sendGame)
		g.POST("/sendGame/:token", sendGame)
		g.GET("/setGameScore/:token", setGameScore)
		g.POST("/setGameScore/:token", setGameScore)
		g.GET("/getGameHighScores/:token", getGameHighScores)
		g.POST("/getGameHighScores/:token", getGameHighScores)
	}
}

func ping(ctx *bm.Context) {
	if err := svc.Ping(ctx); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}
