package service

import (
	"context"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesSetEncryptedTyping(ctx context.Context, request *mtproto.TLMessagesSetEncryptedTyping) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.setEncryptedTyping - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if request.Peer == nil {
		log.Errorf("peer is nil")
		return nil, mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
	}

	encryptChatData, err := s.SecretChatCore.MakeSecretChatData(ctx, request.Peer.ChatId, request.Peer.AccessHash)
	if err != nil {
		log.Errorf("messages.setEncryptedTyping - reply: {%v}", err)
		return mtproto.ToBool(false), nil
	}

	go func() {
		ctx2 := context.Background()

		pushUserId, _ := encryptChatData.GetSecretChatPeerId(md.UserId)

		typing := &mtproto.TLUpdateEncryptedChatTyping{Data2: &mtproto.Update{
			ChatId: md.UserId,
		}}
		pushUpdates := model.NewUpdatesLogicByUpdate(pushUserId, typing.To_Update())
		sync_client.PushUpdates(ctx2, pushUserId, pushUpdates.ToUpdates())
	}()

	log.Debugf("messages.setEncryptedTyping - reply: {true}")
	return mtproto.ToBool(true), nil
}
