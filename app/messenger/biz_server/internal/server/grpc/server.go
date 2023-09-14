package grpc

import (
	"google.golang.org/grpc"
	"open.chat/app/messenger/biz_server/account"
	"open.chat/app/messenger/biz_server/auth"
	"open.chat/app/messenger/biz_server/biz"
	"open.chat/app/messenger/biz_server/blogs"
	"open.chat/app/messenger/biz_server/bots"
	"open.chat/app/messenger/biz_server/channels"
	"open.chat/app/messenger/biz_server/contacts"
	"open.chat/app/messenger/biz_server/folders"
	"open.chat/app/messenger/biz_server/help"
	"open.chat/app/messenger/biz_server/langpack"
	"open.chat/app/messenger/biz_server/messages/bot"
	"open.chat/app/messenger/biz_server/messages/chat"
	"open.chat/app/messenger/biz_server/messages/dialog"
	"open.chat/app/messenger/biz_server/messages/message"
	"open.chat/app/messenger/biz_server/messages/secretchat"
	"open.chat/app/messenger/biz_server/messages/sticker"
	"open.chat/app/messenger/biz_server/payments"
	"open.chat/app/messenger/biz_server/phone"
	"open.chat/app/messenger/biz_server/photos"
	"open.chat/app/messenger/biz_server/stats"
	"open.chat/app/messenger/biz_server/stickers"
	"open.chat/app/messenger/biz_server/updates"
	"open.chat/app/messenger/biz_server/upload"
	"open.chat/app/messenger/biz_server/users"
	"open.chat/app/messenger/biz_server/wallet"
	"open.chat/app/pkg/env2"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util/server"
)

// New new a grpc server.
func New() *server.RPCServer {
	return server.NewRpcServer(env2.MessengerBizServerId, func(s *grpc.Server) {
		mtproto.RegisterRPCAccountServer(s, account.New())
		mtproto.RegisterRPCAuthServer(s, auth.New())
		mtproto.RegisterRPCBizServer(s, biz.New())
		mtproto.RegisterRPCBlogsServer(s, blogs.New())
		mtproto.RegisterRPCBotsServer(s, bots.New())
		mtproto.RegisterRPCChannelsServer(s, channels.New())
		mtproto.RegisterRPCContactsServer(s, contacts.New())
		mtproto.RegisterRPCHelpServer(s, help.New())
		mtproto.RegisterRPCLangpackServer(s, langpack.New())
		mtproto.RegisterRPCMessagesBotServer(s, bot.New())
		mtproto.RegisterRPCMessagesChatServer(s, chat.New())
		mtproto.RegisterRPCMessagesDialogServer(s, dialog.New())
		mtproto.RegisterRPCMessagesMessageServer(s, message.New())
		mtproto.RegisterRPCMessagesSecretChatServer(s, secretchat.New())
		mtproto.RegisterRPCMessagesStickerServer(s, sticker.New())
		mtproto.RegisterRPCPaymentsServer(s, payments.New())
		mtproto.RegisterRPCPhoneServer(s, phone.New())
		mtproto.RegisterRPCPhotosServer(s, photos.New())
		mtproto.RegisterRPCStickersServer(s, stickers.New())
		mtproto.RegisterRPCUpdatesServer(s, updates.New())
		mtproto.RegisterRPCUploadServer(s, upload.New())
		mtproto.RegisterRPCUsersServer(s, users.New())
		mtproto.RegisterRPCFoldersServer(s, folders.New())
		mtproto.RegisterRPCWalletServer(s, wallet.New())
		mtproto.RegisterRPCStatsServer(s, stats.New())
	})
}
