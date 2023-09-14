package dao

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/go-kratos/kratos/pkg/net/rpc/warden"

	"open.chat/app/pkg/env2"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/grpc_util/client"
)

var (
	idMap2 = map[string]string{
		"/mtproto.RPCAccount":            env2.MessengerBizServerAccountId,
		"/mtproto.RPCAuth":               env2.MessengerBizServerAuthId,
		"/mtproto.RPCBots":               env2.MessengerBizServerBotsId,
		"/mtproto.RPCChannels":           env2.MessengerBizServerChannelsId,
		"/mtproto.RPCContacts":           env2.MessengerBizServerContactsId,
		"/mtproto.RPCHelp":               env2.MessengerBizServerHelpId,
		"/mtproto.RPCLangpack":           env2.MessengerBizServerLangpackId,
		"/mtproto.RPCMessagesBot":        env2.MessengerBizServerMessagesBotId,
		"/mtproto.RPCMessagesChat":       env2.MessengerBizServerMessagesChatId,
		"/mtproto.RPCMessagesDialog":     env2.MessengerBizServerMessagesDialogId,
		"/mtproto.RPCMessagesMessage":    env2.MessengerBizServerMessagesMessageId,
		"/mtproto.RPCMessagesSecretchat": env2.MessengerBizServerMessagesSecretchatId,
		"/mtproto.RPCMessagesSticker":    env2.MessengerBizServerMessagesStickerId,
		"/mtproto.RPCPayments":           env2.MessengerBizServerPaymentsId,
		"/mtproto.RPCPhoneId":            env2.MessengerBizServerPhoneId,
		"/mtproto.RPCPhotos":             env2.MessengerBizServerPhotosId,
		"/mtproto.RPCStickers":           env2.MessengerBizServerStickersId,
		"/mtproto.RPCUpdates":            env2.MessengerBizServerUpdatesId,
		"/mtproto.RPCUpload":             env2.MessengerBizServerUploadId,
		"/mtproto.RPCUsers":              env2.MessengerBizServerUsersId,
		"/mtproto.RPCFolders":            env2.MessengerBizServerFoldersId,
	}

	idMap = map[string]string{
		"/mtproto.RPCAccount":            env2.MessengerBizServerId,
		"/mtproto.RPCAuth":               env2.MessengerBizServerId,
		"/mtproto.RPCBots":               env2.MessengerBizServerId,
		"/mtproto.RPCChannels":           env2.MessengerBizServerId,
		"/mtproto.RPCContacts":           env2.MessengerBizServerId,
		"/mtproto.RPCHelp":               env2.MessengerBizServerId,
		"/mtproto.RPCLangpack":           env2.MessengerBizServerId,
		"/mtproto.RPCMessagesBot":        env2.MessengerBizServerId,
		"/mtproto.RPCMessagesChat":       env2.MessengerBizServerId,
		"/mtproto.RPCMessagesDialog":     env2.MessengerBizServerId,
		"/mtproto.RPCMessagesMessage":    env2.MessengerBizServerId,
		"/mtproto.RPCMessagesSecretchat": env2.MessengerBizServerId,
		"/mtproto.RPCMessagesSticker":    env2.MessengerBizServerId,
		"/mtproto.RPCPayments":           env2.MessengerBizServerId,
		"/mtproto.RPCPhoneId":            env2.MessengerBizServerId,
		"/mtproto.RPCPhotos":             env2.MessengerBizServerId,
		"/mtproto.RPCStickers":           env2.MessengerBizServerId,
		"/mtproto.RPCUpdates":            env2.MessengerBizServerId,
		"/mtproto.RPCUpload":             env2.MessengerBizServerId,
		"/mtproto.RPCUsers":              env2.MessengerBizServerId,
		"/mtproto.RPCFolders":            env2.MessengerBizServerId,
	}
)

type Dao struct {
	BizClients map[string]*client.RPCClient
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func New(c *warden.ClientConfig) (dao *Dao) {
	dao = &Dao{}

	dao.BizClients = newBizClients(c)
	return
}

// Close close the resource.
func (d *Dao) Close() {
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	return nil
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

func (d *Dao) Invoke(ctx context.Context, rpcMetaData *grpc_util.RpcMetadata, object mtproto.TLObject) (mtproto.TLObject, error) {
	c, err := d.GetRpcClientByRequest(object)
	if err != nil {
		return nil, err
	}
	return c.InvokeContext(ctx, rpcMetaData, object)
}
