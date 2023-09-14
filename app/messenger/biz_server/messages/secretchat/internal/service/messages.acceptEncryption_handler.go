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

func (s *Service) MessagesAcceptEncryption(ctx context.Context, request *mtproto.TLMessagesAcceptEncryption) (*mtproto.EncryptedChat, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.acceptEncryption - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if request.Peer == nil {
		log.Errorf("peer is nil")
		return nil, mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
	}

	encryptChatData, err := s.SecretChatCore.MakeSecretChatData(ctx, request.Peer.ChatId, request.Peer.AccessHash)
	if err != nil {
		log.Errorf("acceptEncryption error: {%v}", err)
		return nil, mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
	}

	err = encryptChatData.DoAcceptEncryption(ctx, md.UserId, md.AuthId, request.GB, request.KeyFingerprint)
	if err != nil {
		log.Errorf("acceptEncryption error: {%v}", err)
		return nil, err
	}

	go func() {
		ctx2 := context.Background()

		pushUpdates := model.NewUpdatesLogic(encryptChatData.AdminId)

		pushChat, err := encryptChatData.ToEncryptedChat(encryptChatData.AdminId)
		if err != nil {
			return
		}

		updateRequestedEncryption := &mtproto.TLUpdateEncryption{Data2: &mtproto.Update{
			Date: int32(time.Now().Unix()),
			Chat: pushChat,
		}}
		pushUpdates.AddUpdate(updateRequestedEncryption.To_Update())

		users := s.UserFacade.GetUserListByIdList(ctx2, encryptChatData.AdminId, []int32{encryptChatData.AdminId, encryptChatData.ParticipantId})
		pushUpdates.AddUsers(users)
		sync_client.PushUpdates(ctx2, encryptChatData.AdminId, pushUpdates.ToUpdates())
	}()

	replyEncryptedChat, err := encryptChatData.ToEncryptedChat(md.UserId)
	if err != nil {
		log.Errorf("acceptEncryption error: {%v}", err)
		return nil, err
	}

	log.Debugf("messages.acceptEncryption - reply: %s", replyEncryptedChat.DebugString())
	return replyEncryptedChat, nil
}
