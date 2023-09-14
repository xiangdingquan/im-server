package dao

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-kratos/kratos/pkg/net/rpc/warden"

	"open.chat/app/interface/botway/botapi"
	"open.chat/app/interface/botway/internal/dal/dao/mysql_dao"
	"open.chat/app/pkg/env2"
	"open.chat/app/pkg/mysql_util"
	"open.chat/app/service/auth_session/authsessionpb"
	authsession_client "open.chat/app/service/auth_session/client"
	idgen "open.chat/app/service/idgen/client"
	status_client "open.chat/app/service/status/client"
	"open.chat/mtproto"
	"open.chat/pkg/cache"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/grpc_util/client"
	"open.chat/pkg/log"
)

var (
	idMap2 = map[string]string{
		"/mtproto.RPCAccount":            env2.MessengerBizServerAccountId,
		"/mtproto.RPCAuth":               env2.MessengerBizServerAuthId,
		"/mtproto.RPCBots":               env2.MessengerBizServerBotsId,
		"/mtproto.RPCChannels":           env2.MessengerBizServerChannelsId,
		"/mtproto.RPCContacts":           env2.MessengerBizServerContactsId,
		"/mtproto.RPCFolders":            env2.MessengerBizServerFoldersId,
		"/mtproto.RPCHelp":               env2.MessengerBizServerHelpId,
		"/mtproto.RPCLangpack":           env2.MessengerBizServerLangpackId,
		"/mtproto.RPCMessagesBot":        env2.MessengerBizServerMessagesBotId,
		"/mtproto.RPCMessagesChat":       env2.MessengerBizServerMessagesChatId,
		"/mtproto.RPCMessagesDialog":     env2.MessengerBizServerMessagesDialogId,
		"/mtproto.RPCMessagesMessage":    env2.MessengerBizServerMessagesMessageId,
		"/mtproto.RPCMessagesSecretchat": env2.MessengerBizServerMessagesSecretchatId,
		"/mtproto.RPCMessagesSticker":    env2.MessengerBizServerMessagesStickerId,
		"/mtproto.RPCPayments":           env2.MessengerBizServerPaymentsId,
		"/mtproto.RPCPhone":              env2.MessengerBizServerPhoneId,
		"/mtproto.RPCPhotos":             env2.MessengerBizServerPhotosId,
		"/mtproto.RPCStats":              env2.MessengerBizServerStatsId,
		"/mtproto.RPCStickers":           env2.MessengerBizServerStickersId,
		"/mtproto.RPCUpdates":            env2.MessengerBizServerUpdatesId,
		"/mtproto.RPCUpload":             env2.MessengerBizServerUploadId,
		"/mtproto.RPCUsers":              env2.MessengerBizServerUsersId,
		"/mtproto.RPCWallet":             env2.MessengerBizServerWalletId,
		"/mtproto.RPCBlogs":              env2.MessengerBizServerBlogsId,
	}

	idMap = map[string]string{
		"/mtproto.RPCAccount":            env2.MessengerBizServerId,
		"/mtproto.RPCAuth":               env2.MessengerBizServerId,
		"/mtproto.RPCBots":               env2.MessengerBizServerId,
		"/mtproto.RPCChannels":           env2.MessengerBizServerId,
		"/mtproto.RPCContacts":           env2.MessengerBizServerId,
		"/mtproto.RPCFolders":            env2.MessengerBizServerId,
		"/mtproto.RPCHelp":               env2.MessengerBizServerId,
		"/mtproto.RPCLangpack":           env2.MessengerBizServerId,
		"/mtproto.RPCMessagesBot":        env2.MessengerBizServerId,
		"/mtproto.RPCMessagesChat":       env2.MessengerBizServerId,
		"/mtproto.RPCMessagesDialog":     env2.MessengerBizServerId,
		"/mtproto.RPCMessagesMessage":    env2.MessengerBizServerId,
		"/mtproto.RPCMessagesSecretchat": env2.MessengerBizServerId,
		"/mtproto.RPCMessagesSticker":    env2.MessengerBizServerId,
		"/mtproto.RPCPayments":           env2.MessengerBizServerId,
		"/mtproto.RPCPhone":              env2.MessengerBizServerId,
		"/mtproto.RPCPhotos":             env2.MessengerBizServerId,
		"/mtproto.RPCStats":              env2.MessengerBizServerId,
		"/mtproto.RPCStickers":           env2.MessengerBizServerId,
		"/mtproto.RPCUpdates":            env2.MessengerBizServerId,
		"/mtproto.RPCUpload":             env2.MessengerBizServerId,
		"/mtproto.RPCUsers":              env2.MessengerBizServerId,
		"/mtproto.RPCWallet":             env2.MessengerBizServerId,
		"/mtproto.RPCBlogs":              env2.MessengerBizServerId,
	}
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.BotUpdatesDAO
}

func newMysqlDao() *Mysql {
	db := mysql_util.GetSingletonSqlxDB()
	return &Mysql{
		DB:            db,
		BotUpdatesDAO: mysql_dao.NewBotUpdatesDAO(db),
	}
}

func (d *Mysql) Close() error {
	return d.DB.Close()
}

func (d *Mysql) Ping(ctx context.Context) (err error) {
	return d.DB.Ping(ctx)
}

type Dao struct {
	BizClients           map[string]*client.RPCClient
	AuthSessionRpcClient authsessionpb.RPCSessionClient
	cache                *cache.LRUCache
	*Mysql
	*Redis
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func New(c *warden.ClientConfig) (dao *Dao) {
	idgen.NewUUID()

	err := status_client.New()
	if err != nil {
		panic(err)
	}

	dao = &Dao{
		Mysql: newMysqlDao(),
		Redis: newRedisDao(),
	}

	dao.AuthSessionRpcClient, err = authsession_client.New(c)
	if err != nil {
		panic(err)
	}

	dao.BizClients = newBizClients(c)
	dao.cache = cache.NewLRUCache(1024 * 1024 * 1024)

	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.Redis.Close()
	d.Mysql.Close()
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	if err = d.Redis.Ping(ctx); err != nil {
		return
	}
	return d.Mysql.Ping(ctx)
}

func (d *Dao) GetUserByToken(ctx context.Context, token string) (me *botapi.User, err error) {
	me2 := new(botapi.User)
	err = d.Mysql.DB.QueryRow(ctx,
		"SELECT id, is_bot, first_name, last_name, username FROM users WHERE id IN (SELECT bot_id FROM `bots` WHERE token = ?)",
		token).StructScan(me2)
	if err != nil && err != sqlx.ErrNoRows {
		log.Error("d.GetDemo.Query error(%v)", err)
		return
	}
	me = me2
	return
}

func (d *Dao) GetUser(ctx context.Context, userId int32) (user *botapi.User, err error) {
	me2 := new(botapi.User)
	err = d.Mysql.DB.QueryRow(ctx,
		"SELECT id, is_bot, first_name, last_name, username FROM users WHERE id  = ?", userId).StructScan(me2)
	if err != nil && err != sqlx.ErrNoRows {
		log.Error("d.GetUser.Query error(%v)", err)
		return
	}
	user = me2
	return
}

func newBizClients(config *warden.ClientConfig) (bizClients map[string]*client.RPCClient) {
	var (
		clients   = make(map[string]*client.RPCClient)
		registers = mtproto.GetRPCContextRegisters()
	)

	for k, v := range idMap {
		c, err := client.NewRPCClient(v, config)
		if err != nil {
			panic(err)
		}
		clients[k] = c
	}

	bizClients = make(map[string]*client.RPCClient)

	for m, ctx := range registers {
		for k, _ := range idMap {
			if strings.HasPrefix(ctx.Method, k) {
				bizClients[m] = clients[k]
				break
			}
		}
	}

	return
}

func (d *Dao) GetRpcClientByRequest(t interface{}) (*client.RPCClient, error) {
	rt := reflect.TypeOf(t)
	if rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	if c, ok := d.BizClients[rt.Name()]; ok {
		return c, nil
	}
	return nil, fmt.Errorf("not found method: %s", rt.Name())
}

func (d *Dao) Invoke(rpcMetaData *grpc_util.RpcMetadata, object mtproto.TLObject) (mtproto.TLObject, error) {
	c, err := d.GetRpcClientByRequest(object)
	if err != nil {
		return nil, err
	}
	return c.Invoke(rpcMetaData, object)
}
