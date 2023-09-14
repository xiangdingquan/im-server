package service

import (
	"context"
	"time"

	"github.com/gogo/protobuf/types"

	"open.chat/app/messenger/msg/msgpb"
	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/app/sysconfig"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesEditMessage(ctx context.Context, request *mtproto.TLMessagesEditMessage) (reply *mtproto.Updates, err error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.editMessage#5d1b8dd - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		hasBot       = md.IsBot
		isPoll       bool
		peer         = model.FromInputPeer2(md.UserId, request.Peer)
		editMessages model.MessageBoxList
	)

	switch peer.PeerType {
	case model.PEER_SELF, model.PEER_USER, model.PEER_CHAT:
		if peer.PeerType == model.PEER_SELF {
			peer.PeerType = model.PEER_USER
		}
		if peer.PeerType == model.PEER_USER {
			if !md.IsBot {
				hasBot = s.UserFacade.IsBot(ctx, peer.PeerId)
			}
		}
		editMessages = s.MessageFacade.GetUserMessageList(ctx, md.UserId, []int32{request.Id})
	case model.PEER_CHANNEL:
		editMessages = s.MessageFacade.GetChannelMessageList(ctx, md.UserId, peer.PeerId, []int32{request.Id})
	default:
		log.Errorf("invalid peer: %v", request.Peer)
		err = mtproto.ErrPeerIdInvalid
		return
	}

	if len(editMessages) != 1 {
		err = mtproto.ErrMessageEmpty
		log.Errorf("messages.editMessage - emptyMessage(%d)", request.Id)
		return
	}

	if editMessages[0].SelfUserId != md.UserId {
		err = mtproto.ErrMessageAuthorRequired
		log.Errorf("messages.editMessage - emptyMessage(%d)", request.Id)
		return
	}

	outMessage := editMessages[0].Message
	outMessage.EditDate = &types.Int32Value{Value: int32(time.Now().Unix())}

	if request.Entities != nil {
		outMessage.Entities = request.Entities
	}

	if request.ReplyMarkup != nil {
		outMessage.ReplyMarkup = request.ReplyMarkup
	}

	if request.Media != nil {
		outMessage.Media, err = s.makeMediaByInputMedia(ctx, md.UserId, md.AuthId, peer, request.Media)
		if err != nil {
			if err.Error() == string(sysconfig.ConfigKeysBanImages) {
				if request.Message != nil {
					request.Message.Value = request.Message.Value + "[违规发图片,被屏蔽]"
				} else {
					request.Message = &types.StringValue{
						Value: "[违规发图片,被屏蔽]",
					}
				}
			} else {
				log.Errorf("messages.editMessage - media error: %v", err)
				return
			}
		}
		if outMessage.Media != nil && outMessage.Media.PredicateName == mtproto.Predicate_messageMediaPoll {
			isPoll = true
		}
	}

	if request.Message != nil {
		if request.Message.Value == "" {
			err = mtproto.ErrMessageEmpty
			log.Errorf("message empty: %v", err)
			return
		}
		outMessage.Message = request.Message.Value
		outMessage.Entities = nil
		outMessage, _ = s.fixMessageEntities(ctx, md.UserId, peer, request.NoWebpage, outMessage, hasBot)
	}

	if isPoll {
		var (
			pollId    int64
			mediaPoll *model.MediaPoll
		)
		pollId, err = model.GetPollIdByMessage(outMessage.GetMedia())
		if err != nil {
			log.Errorf("messages.editMessage - media error: %v", err)
			return
		}
		mediaPoll, err = s.PollFacade.CloseMediaPoll(ctx, md.UserId, pollId)
		if err != nil {
			log.Errorf("messages.editMessage - media error: %v", err)
			return
		}

		reply = model.MakeUpdatesByUpdates(mediaPoll.ToUpdateMessagePoll())

		go func() {
			pushUpdates := reply
			sync_client.SyncUpdatesNotMe(context.Background(), md.UserId, md.AuthId, pushUpdates)

			switch peer.PeerType {
			case model.PEER_USER:
				sync_client.PushUpdates(context.Background(), peer.PeerId, pushUpdates)
			case model.PEER_CHAT:
				sync_client.BroadcastChatUpdates(context.Background(), peer.PeerId, pushUpdates, md.UserId)
			case model.PEER_CHANNEL:
				sync_client.BroadcastChannelUpdates(context.Background(), peer.PeerId, pushUpdates, md.UserId)
			}
		}()
	} else {
		outboxMsg := &msgpb.OutboxMessage{
			NoWebpage:    request.NoWebpage,
			Message:      outMessage,
			ScheduleDate: request.ScheduleDate,
		}
		reply, err = s.MsgFacade.EditMessage(ctx, md.UserId, md.AuthId, peer, outboxMsg)
	}

	log.Debugf("messages.editMessage#5d1b8dd - reply: %s", reply.DebugString())
	return
}
