package service

import (
	"context"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetPollVotes(ctx context.Context, request *mtproto.TLMessagesGetPollVotes) (*mtproto.Messages_VotesList, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getPollVotes - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		err    error
		pollId int64
		peer   = model.FromInputPeer2(md.UserId, request.Peer)
		boxMsg *model.MessageBox
		limit  = request.Limit
	)

	if limit > 50 {
		limit = 50
	}

	if md.IsBot {
		err = mtproto.ErrBotMethodInvalid
		log.Errorf("messages.getPollVotes - error: %v", err)
		return nil, err
	}

	switch peer.PeerType {
	case model.PEER_SELF, model.PEER_USER, model.PEER_CHAT:
		if boxMsg, err = s.MessageFacade.GetUserMessage(ctx, md.UserId, request.Id); err != nil {
			log.Errorf("messages.getPollVotes - error: %v", err)
			err = mtproto.ErrMessageIdInvalid
			return nil, err
		}
	case model.PEER_CHANNEL:
		if boxMsg, err = s.MessageFacade.GetChannelMessage(ctx, md.UserId, peer.PeerId, request.Id); err != nil {
			log.Errorf("messages.getPollVotes - error: %v", err)
			err = mtproto.ErrMessageIdInvalid
			return nil, err
		}
	default:
		err = mtproto.ErrPeerIdInvalid
		log.Errorf("messages.getPollVotes - error: %v", err)
		return nil, err
	}

	pollId, err = model.GetPollIdByMessage(boxMsg.Message.GetMedia())
	if err != nil {
		log.Errorf("messages.getPollVotes - error: %v", err)
		return nil, err
	}

	var (
		option string
		offset string
	)

	if len(request.GetOption()) > 0 {
		option = hack.String(request.Option)
	}
	offset = request.GetOffset().GetValue()

	votersList, err := s.PollFacade.GetPollVoters(ctx, md.UserId, pollId, option, offset, limit)
	if err != nil {
		log.Errorf("getPollVoters - error: %v", err)
		return nil, err
	}

	voters := make([]int32, 0, len(votersList.Votes))
	for _, v := range votersList.Votes {
		voters = append(voters, v.UserId)
	}

	votersList.Users = s.UserFacade.GetUserListByIdList(ctx, md.UserId, voters)
	if votersList.Users == nil {
		votersList.Users = []*mtproto.User{}
	}

	log.Debugf("messages.getPollVotes - result: %s", votersList.DebugString())
	return votersList, nil
}
