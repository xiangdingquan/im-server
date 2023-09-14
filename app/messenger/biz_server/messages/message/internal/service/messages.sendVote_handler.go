package service

import (
	"context"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) MessagesSendVote(ctx context.Context, request *mtproto.TLMessagesSendVote) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.sendVote - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		err     error
		pollId  int64
		peer    = model.FromInputPeer2(md.UserId, request.Peer)
		options = make([]string, 0, len(request.Options))
		msgList model.MessageBoxList
	)

	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("messages.sendVote - error: %v", err)
		return nil, err
	}

	switch peer.PeerType {
	case model.PEER_SELF:
		peer.PeerType = model.PEER_USER
	case model.PEER_USER:
	case model.PEER_CHAT:
	case model.PEER_CHANNEL:
	default:
		err = mtproto.ErrPeerIdInvalid
		log.Errorf("messages.sendVote - error: %v", err)
		return nil, err
	}

	for _, v := range request.Options {
		options = append(options, hack.String(v))
	}

	if peer.PeerType == model.PEER_CHANNEL {
		msgList = s.MessageFacade.GetChannelMessageList(ctx, md.UserId, peer.PeerId, []int32{request.MsgId})
	} else {
		msgList = s.MessageFacade.GetUserMessageList(ctx, md.UserId, []int32{request.MsgId})
	}

	if len(msgList) == 0 {
		err = mtproto.ErrMessageIdInvalid
		log.Errorf("messages.sendVote - error: %v", err)
		return nil, err
	}

	pollId, err = model.GetPollIdByMessage(msgList[0].Message.GetMedia())
	if err != nil {
		log.Errorf("messages.sendVote - error: %v", err)
		return nil, err
	}

	mediaPoll, err := s.PollFacade.SendVote(ctx, md.UserId, pollId, options)
	if err != nil {
		log.Errorf("messages.sendVote - error: %v", err)
		return nil, err
	}
	log.Debugf("mediaPoll - %v", mediaPoll.DebugString())

	result := model.MakeUpdatesByUpdates(mtproto.MakeTLUpdateMessagePoll(&mtproto.Update{
		PollId:  mediaPoll.Poll.Id,
		Poll:    mediaPoll.Poll,
		Results: mediaPoll.Results,
	}).To_Update())

	log.Debugf("messages.sendVote - reply %s", result.DebugString())
	return model.WrapperGoFunc(result, func() {
		go func() {
			var (
				pushUpdates = result
			)

			sync_client.SyncUpdatesNotMe(context.Background(), md.UserId, md.AuthId, pushUpdates)

			switch peer.PeerId {
			case model.PEER_USER:
				sync_client.PushUpdates(context.Background(), peer.PeerId, pushUpdates)
			case model.PEER_CHAT:
				sync_client.BroadcastChatUpdates(context.Background(), peer.PeerId, pushUpdates, md.UserId)
			case model.PEER_CHANNEL:
				sync_client.BroadcastChannelUpdates(context.Background(), peer.PeerId, pushUpdates, md.UserId)
			}
		}()
	}).(*mtproto.Updates), nil
}
