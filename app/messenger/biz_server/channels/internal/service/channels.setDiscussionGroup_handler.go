package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsSetDiscussionGroup(ctx context.Context, request *mtproto.TLChannelsSetDiscussionGroup) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.setDiscussionGroup - metadata: %s, request: %s", md.DebugString(), request.DebugString())
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("channels.setDiscussionGroup - error: %v", err)
		return nil, err
	}

	var groupId int32
	switch request.GetGroup().GetPredicateName() {
	case mtproto.Predicate_inputChannel:
		groupId = request.Group.ChannelId
	case mtproto.Predicate_inputChannelEmpty:
		groupId = 0
	}
	err := s.ChannelFacade.SetDiscussionGroup(ctx, md.UserId, request.GetBroadcast().GetChannelId(), groupId)
	if err != nil {
		log.Errorf("channels.setDiscussionGroup - error: %v", err)
		return nil, err
	}

	log.Debugf("channels.setDiscussionGroup - reply: {true}")
	return mtproto.BoolFalse, nil
}
