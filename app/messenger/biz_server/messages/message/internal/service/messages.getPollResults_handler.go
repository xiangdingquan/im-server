package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) MessagesGetPollResults(ctx context.Context, request *mtproto.TLMessagesGetPollResults) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getPollResults - metadata: %s, request: %s", logger.JsonDebugData(md), logger.JsonDebugData(request))

	var (
		err    error
		pollId int64
		peer   = model.FromInputPeer2(md.UserId, request.Peer)
		boxMsg *model.MessageBox
	)

	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("messages.getPollResults - error: %v", err)
		return nil, err
	}

	switch peer.PeerType {
	case model.PEER_SELF, model.PEER_USER, model.PEER_CHAT:
		if boxMsg, err = s.MessageFacade.GetUserMessage(ctx, md.UserId, request.MsgId); err != nil {
			log.Errorf("messages.getPollResults - error: %v", err)
			err = mtproto.ErrMessageIdInvalid
			return nil, err
		}
	case model.PEER_CHANNEL:
		if boxMsg, err = s.MessageFacade.GetChannelMessage(ctx, md.UserId, peer.PeerId, request.MsgId); err != nil {
			log.Errorf("messages.getPollResults - error: %v", err)
			err = mtproto.ErrMessageIdInvalid
			return nil, err
		}
	default:
		err = mtproto.ErrPeerIdInvalid
		log.Errorf("messages.getPollResults - error: %v", err)
		return nil, err
	}

	pollId, err = model.GetPollIdByMessage(boxMsg.Message.GetMedia())
	if err != nil {
		log.Errorf("messages.getPollResults - error: %v", err)
		return nil, err
	}

	mediaPoll, err := s.PollFacade.GetMediaPoll(ctx, md.UserId, pollId)
	if err != nil {
		log.Errorf("messages.getPollResults - error: %v", err)
		return nil, err
	}

	result := model.MakeUpdatesByUpdates(mtproto.MakeTLUpdateMessagePoll(&mtproto.Update{
		PollId:  mediaPoll.Poll.Id,
		Poll:    nil,
		Results: mediaPoll.Results,
	}).To_Update())

	log.Debugf("messages.getPollResults - reply %s", result.DebugString())
	return result, nil
}
