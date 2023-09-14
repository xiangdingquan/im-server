package service

import (
	"context"

	"open.chat/app/bots/botpb"
	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetBotCallbackAnswer(ctx context.Context, request *mtproto.TLMessagesGetBotCallbackAnswer) (*mtproto.Messages_BotCallbackAnswer, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getBotCallbackAnswer - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.getBotCallbackAnswer - error: %v", err)
		return nil, err
	}

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	// 400	CHANNEL_INVALID	The provided channel is invalid
	// 400	DATA_INVALID	Encrypted data invalid
	// 400	MESSAGE_ID_INVALID	The provided message id is invalid
	// 400	PEER_ID_INVALID	The provided peer id is invalid
	// -503	Timeout	Timeout while fetching data
	//
	peer := model.FromInputPeer2(md.UserId, request.Peer)

	peerMsg, err := s.MessageFacade.GetPeerUserMessage(ctx, md.UserId, request.MsgId, peer.PeerId)
	if err != nil {
		log.Errorf("messages.getBotCallbackAnswer - error: %v", err)
		err = mtproto.ErrMessageIdInvalid
		return nil, err
	}

	switch peer.PeerType {
	case model.PEER_USER:
		if model.IsBotFather(peer.PeerId) {
			return s.Dao.BotsClient.MessagesGetBotCallbackAnswer(ctx, &botpb.GetBotCallbackAnswer{
				UserId:    md.UserId,
				AuthKeyId: md.AuthId,
				IsGame:    request.Game,
				BotId:     peer.PeerId,
				Message:   peerMsg.Message,
				Data:      hack.String(request.Data),
			})
		} else {
			queryId, err := s.Dao.PutCacheRpcMetadata(ctx, md)
			if err != nil {
				log.Errorf("messages.getBotCallbackAnswer - error: %v", err)
				return nil, err
			}
			go func() {
				ctx2 := context.Background()
				pushUpdates := model.MakeUpdatesByUpdates(mtproto.MakeTLUpdateBotCallbackQuery(&mtproto.Update{
					QueryId:         queryId,
					UserId:          md.UserId,
					Peer_PEER:       model.MakePeerUser(md.UserId),
					MsgId_INT32:     peerMsg.MessageId,
					ChatInstance:    0,
					Data_FLAGBYTES:  request.Data,
					GameShortName:   nil,
					Message_MESSAGE: peerMsg.Message,
				}).To_Update())
				pushUpdates.Users = s.UserFacade.GetUserListByIdList(ctx2, peer.PeerId, []int32{md.UserId, peer.PeerId})
				sync_client.PushUpdates(ctx2, peer.PeerId, pushUpdates)
			}()
			err = mtproto.ErrNotReturnClient
			log.Errorf("messages.getBotCallbackAnswer - error: %v", err)
			return nil, err
		}

	case model.PEER_CHAT:
	case model.PEER_CHANNEL:
	default:
		err := mtproto.ErrPeerIdInvalid
		log.Errorf("messages.getBotCallbackAnswer - error: %v", err)
		return nil, err
	}

	err = mtproto.ErrNotReturnClient
	log.Errorf("messages.getBotCallbackAnswer - error: %v", err)
	return nil, err
}
