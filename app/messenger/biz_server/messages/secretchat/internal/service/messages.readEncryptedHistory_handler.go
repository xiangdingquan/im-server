package service

import (
	"context"
	"time"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesReadEncryptedHistory(ctx context.Context, request *mtproto.TLMessagesReadEncryptedHistory) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.readEncryptedHistory#7f4b690a - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if request.Peer == nil {
		log.Errorf("peer is nil")
		return nil, mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
	}

	encryptChatData, err := s.SecretChatCore.MakeSecretChatData(ctx, request.Peer.ChatId, request.Peer.AccessHash)
	if err != nil {
		log.Errorf("acceptEncryption error: {%v}", err)
		return nil, err
	}

	go func() {
		ctx2 := context.Background()

		peerId, _ := encryptChatData.GetSecretChatPeerId(md.UserId)
		pushUpdates := model.NewUpdatesLogic(peerId)
		updateEncryptedMessagesRead := &mtproto.TLUpdateEncryptedMessagesRead{Data2: &mtproto.Update{
			ChatId:  encryptChatData.Id,
			MaxDate: request.MaxDate,
			Date:    int32(time.Now().Unix()),
		}}

		pushUpdates.AddUpdate(updateEncryptedMessagesRead.To_Update())
		users := s.UserFacade.GetUserListByIdList(ctx2, peerId, []int32{encryptChatData.AdminId, encryptChatData.ParticipantId})
		pushUpdates.AddUsers(users)
		sync_client.PushUpdates(ctx2, peerId, pushUpdates.ToUpdates())
	}()

	log.Debug("messages.readEncryptedHistory#7f4b690a - reply {true}")
	return mtproto.ToBool(true), nil
}
