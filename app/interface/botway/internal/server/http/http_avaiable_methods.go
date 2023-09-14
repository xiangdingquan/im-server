package http

import (
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"open.chat/app/interface/botway/botapi"
)

func getMe(c *bm.Context) {
	req := new(botapi.GetMe2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.GetMe(c, token, req)
	})
}

func sendMessage(c *bm.Context) {
	req := new(botapi.SendMessage2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SendMessage(c, token, req)
	})
}

func forwardMessage(c *bm.Context) {
	req := new(botapi.ForwardMessage2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.ForwardMessage(c, token, req)
	})
}

func sendPhoto(c *bm.Context) {
	req := new(botapi.SendPhoto2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SendPhoto(c, token, req)
	})
}

func sendAudio(c *bm.Context) {
	req := new(botapi.SendAudio2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SendAudio(c, token, req)
	})
}

func sendDocument(c *bm.Context) {
	req := new(botapi.SendDocument2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SendDocument(c, token, req)
	})
}

func sendVideo(c *bm.Context) {
	req := new(botapi.SendVideo2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SendVideo(c, token, req)
	})
}

func sendAnimation(c *bm.Context) {
	req := new(botapi.SendAnimation2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SendAnimation(c, token, req)
	})
}

func sendVoice(c *bm.Context) {
	req := new(botapi.SendVoice2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SendVoice(c, token, req)
	})
}

func sendVideoNote(c *bm.Context) {
	req := new(botapi.SendVideoNote2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SendVideoNote(c, token, req)
	})
}

func sendMediaGroup(c *bm.Context) {
	req := new(botapi.SendMediaGroup2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SendMediaGroup(c, token, req)
	})

}

func sendLocation(c *bm.Context) {
	req := new(botapi.SendLocation2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SendLocation(c, token, req)
	})
}

func editMessageLiveLocation(c *bm.Context) {
	req := new(botapi.EditMessageLiveLocation2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.EditMessageLiveLocation(c, token, req)
	})
}

func stopMessageLiveLocation(c *bm.Context) {
	req := new(botapi.StopMessageLiveLocation2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.StopMessageLiveLocation(c, token, req)
	})
}

func sendVenue(c *bm.Context) {
	req := new(botapi.SendVenue2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SendVenue(c, token, req)
	})
}

func sendContact(c *bm.Context) {
	req := new(botapi.SendContact2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SendContact(c, token, req)
	})
}

func sendPoll(c *bm.Context) {
	req := new(botapi.SendPoll2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SendPoll(c, token, req)
	})
}

func sendChatAction(c *bm.Context) {
	req := new(botapi.SendChatAction2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SendChatAction(c, token, req)
	})
}

func getUserProfilePhotos(c *bm.Context) {
	req := new(botapi.GetUserProfilePhotos2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.GetUserProfilePhotos(c, token, req)
	})
}

func getFile(c *bm.Context) {
	req := new(botapi.GetFile2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.GetFile(c, token, req)
	})
}

func kickChatMember(c *bm.Context) {
	req := new(botapi.KickChatMember2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.KickChatMember(c, token, req)
	})
}

func unbanChatMember(c *bm.Context) {
	req := new(botapi.UnbanChatMember2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.UnbanChatMember(c, token, req)
	})
}

func restrictChatMember(c *bm.Context) {
	req := new(botapi.RestrictChatMember2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.RestrictChatMember(c, token, req)
	})
}

func promoteChatMember(c *bm.Context) {
	req := new(botapi.PromoteChatMember2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.PromoteChatMember(c, token, req)
	})
}

func setChatPermissions(c *bm.Context) {
	req := new(botapi.SetChatPermissions2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SetChatPermissions(c, token, req)
	})
}

func exportChatInviteLink(c *bm.Context) {
	req := new(botapi.ExportChatInviteLink2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.ExportChatInviteLink(c, token, req)
	})
}

func setChatPhoto(c *bm.Context) {
	req := new(botapi.SetChatPhoto2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SetChatPhoto(c, token, req)
	})
}

func deleteChatPhoto(c *bm.Context) {
	req := new(botapi.DeleteChatPhoto2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.DeleteChatPhoto(c, token, req)
	})
}

func setChatTitle(c *bm.Context) {
	req := new(botapi.SetChatTitle2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SetChatTitle(c, token, req)
	})
}

func setChatDescription(c *bm.Context) {
	req := new(botapi.SetChatDescription2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SetChatDescription(c, token, req)
	})
}

func pinChatMessage(c *bm.Context) {
	req := new(botapi.PinChatMessage2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.PinChatMessage(c, token, req)
	})
}

func unpinChatMessage(c *bm.Context) {
	req := new(botapi.UnpinChatMessage2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.UnpinChatMessage(c, token, req)
	})
}

func leaveChat(c *bm.Context) {
	req := new(botapi.LeaveChat2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.LeaveChat(c, token, req)
	})
}

func getChat(c *bm.Context) {
	req := new(botapi.GetChat2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.GetChat(c, token, req)
	})
}

func getChatAdministrators(c *bm.Context) {
	req := new(botapi.GetChatAdministrators2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.GetChatAdministrators(c, token, req)
	})
}

func getChatMembersCount(c *bm.Context) {
	req := new(botapi.GetChatMembersCount2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.GetChatMembersCount(c, token, req)
	})
}

func getChatMember(c *bm.Context) {
	req := new(botapi.GetChatMember2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.GetChatMember(c, token, req)
	})
}

func setChatStickerSet(c *bm.Context) {
	req := new(botapi.SendContact2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.SendContact(c, token, req)
	})
}

func deleteChatStickerSet(c *bm.Context) {
	req := new(botapi.DeleteChatStickerSet2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.DeleteChatStickerSet(c, token, req)
	})
}

func answerCallbackQuery(c *bm.Context) {
	req := new(botapi.AnswerCallbackQuery2)
	botHandlerHelper(c, req, func(c *bm.Context, token string) (interface{}, error) {
		return svc.AnswerCallbackQuery(c, token, req)
	})
}
