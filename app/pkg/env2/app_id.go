package env2

import (
	"flag"
)

const (
	//infra
	InfraDatabusId = "infra.databus"

	// interface
	InterfaceGatewayId        = "interface.gateway"
	InterfaceBotwayId         = "interface.botway"
	InterfaceSessionSessionId = "interface.session.session"
	InterfaceSessionPushId    = "interface.session.push"
	InterfaceRelayId          = "interface.relay"

	// bots
	BotsBotFatherId = "bots.botfather"
	BotsGifId       = "bots.gif"

	// job
	JobPush = "job.push"

	////////////////////////////////////////////////////////////////////////////////////////
	// messenger.biz_server
	MessengerBizServerId = "messenger.biz_server"

	// biz_servers
	MessengerBizServerAccountId            = "messenger.biz_server.account"
	MessengerBizServerAuthId               = "messenger.biz_server.auth"
	MessengerBizServerBotsId               = "messenger.biz_server.bots"
	MessengerBizServerChannelsId           = "messenger.biz_server.channels"
	MessengerBizServerContactsId           = "messenger.biz_server.contacts"
	MessengerBizServerHelpId               = "messenger.biz_server.help"
	MessengerBizServerLangpackId           = "messenger.biz_server.langpack"
	MessengerBizServerMessagesBotId        = "messenger.biz_server.messages.bot"
	MessengerBizServerMessagesChatId       = "messenger.biz_server.messages.chat"
	MessengerBizServerMessagesDialogId     = "messenger.biz_server.messages.dialog"
	MessengerBizServerMessagesMessageId    = "messenger.biz_server.messages.message"
	MessengerBizServerMessagesSecretchatId = "messenger.biz_server.messages.secretchat"
	MessengerBizServerMessagesStickerId    = "messenger.biz_server.messages.sticker"
	MessengerBizServerPaymentsId           = "messenger.biz_server.payments"
	MessengerBizServerPhoneId              = "messenger.biz_server.phone"
	MessengerBizServerPhotosId             = "messenger.biz_server.photos"
	MessengerBizServerStickersId           = "messenger.biz_server.stickers"
	MessengerBizServerUpdatesId            = "messenger.biz_server.updates"
	MessengerBizServerUploadId             = "messenger.biz_server.upload"
	MessengerBizServerUsersId              = "messenger.biz_server.users"
	MessengerBizServerFoldersId            = "messenger.biz_server.folders"
	MessengerBizServerWalletId             = "messenger.biz_server.wallet"
	MessengerBizServerStatsId              = "messenger.biz_server.stats"
	MessengerBizServerBlogsId              = "messenger.biz_server.blogs"

	// ...
	MessengerInboxId   = "messenger.inbox"
	MessengerOutboxId  = "messenger.outbox"
	MessengerSyncId    = "messenger.sync"
	MessengerWebPageId = "messenger.webpage"

	// service
	ServiceAuthSessionId = "service.auth_session"
	ServiceMediaId       = "service.media"

	// infra
	InfraConfigId = "infra.config"
)

var (
	SMS_CODE_NAME = ""
	MY_APP_NAME   = "" //Gochat
	MY_WEB_SITE   = "" //gochat8.com
	T_ME          = "" //https://jk.gochat8.com

	PredefinedUser = false

	// predefined2 - auto register
	PredefinedUser2 = false

	Server2 = false
)

func init() {
	flag.StringVar(&SMS_CODE_NAME, "code", "", "code")
	flag.StringVar(&MY_APP_NAME, "app_name", "", "app_name")
	flag.StringVar(&MY_WEB_SITE, "site_name", "", "site_name")
	flag.StringVar(&T_ME, "t_me", "", "t_me")
	flag.BoolVar(&PredefinedUser, "predefined", false, "predefined")
	flag.BoolVar(&PredefinedUser2, "predefined2", false, "predefined2")
	flag.BoolVar(&Server2, "server2", false, "server2")
}
