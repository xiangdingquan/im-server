package services

import (
	svc "open.chat/app/json/service"
	"open.chat/app/json/services/handler/accounts"
	"open.chat/app/json/services/handler/blogs"
	"open.chat/app/json/services/handler/call"
	"open.chat/app/json/services/handler/channel"
	"open.chat/app/json/services/handler/chats"
	"open.chat/app/json/services/handler/discover"
	"open.chat/app/json/services/handler/messages"
	"open.chat/app/json/services/handler/redpacket"
	"open.chat/app/json/services/handler/remittance"
	"open.chat/app/json/services/handler/system"
	"open.chat/app/json/services/handler/users"
	"open.chat/app/json/services/handler/wallet"
)

// New .
func RegistHandler(s *svc.Service) {
	system.New(s)
	call.New(s)
	discover.New(s)
	redpacket.New(s)
	wallet.New(s)
	users.New(s)
	channel.New(s)
	chats.New(s)
	remittance.New(s)
	blogs.New(s)
	accounts.New(s)
	messages.New(s)
}
