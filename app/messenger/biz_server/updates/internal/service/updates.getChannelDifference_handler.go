package service

import (
	"context"

	"github.com/gogo/protobuf/types"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

const (
	channelPollTimeout = int32(30)
	maxLimit           = int32(100)
)

func (s *Service) UpdatesGetChannelDifference(ctx context.Context, request *mtproto.TLUpdatesGetChannelDifference) (*mtproto.Updates_ChannelDifference, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("updates.getChannelDifference#3173d78 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	difference, err := s.getChannelDifference(ctx, md, request)
	if err != nil {
		log.Errorf("updates.getChannelDifference#3173d78 - error: %v", err)
		return nil, err
	}

	log.Debugf("updates.getChannelDifference#3173d78 - reply: %s", logger.JsonDebugData(difference))
	return difference, nil
}

func (s *Service) getChannelDifference(
	ctx context.Context,
	md *grpc_util.RpcMetadata,
	request *mtproto.TLUpdatesGetChannelDifference) (*mtproto.Updates_ChannelDifference, error) {

	channelId, err := getChannelId(ctx, request.Channel)
	if err != nil {
		return nil, err
	}

	channel, err := s.ChannelCore.GetMutableChannel(ctx, channelId, md.UserId)
	if err != nil {
		log.Errorf("updates.getChannelDifference#3173d78 - error: %v", err)
		return nil, err
	}

	participant := channel.GetImmutableChannelParticipant(md.UserId)
	if participant != nil && participant.IsKicked() {
		log.Errorf("updates.getChannelDifference#3173d78 - error: invalid participant(%d)", md.UserId)
		return nil, mtproto.ErrChannelPrivate
	}

	if participant != nil && participant.AvailableMinPts > request.Pts {
		request.Pts = participant.AvailableMinPts
	}

	switch request.Filter.PredicateName {
	case mtproto.Predicate_channelMessagesFilterEmpty:
	case mtproto.Predicate_channelMessagesFilter:
	default:
		log.Errorf("invalid filter - %v", request.Filter)
	}

	var difference *mtproto.Updates_ChannelDifference

	if request.Pts >= channel.Channel.Pts {
		difference = makeChannelDifferenceEmpty(channel.Channel.Pts)
	} else {
		difference = mtproto.MakeTLUpdatesChannelDifference(&mtproto.Updates_ChannelDifference{
			Timeout:      &types.Int32Value{Value: channelPollTimeout},
			Pts:          request.Pts,
			NewMessages:  []*mtproto.Message{},
			OtherUpdates: []*mtproto.Update{},
			Chats:        []*mtproto.Chat{},
			Users:        []*mtproto.User{},
		}).To_Updates_ChannelDifference()

		limit := request.Limit
		if limit > maxLimit {
			limit = maxLimit
		}

		updateList := s.ChannelCore.GetChannelUpdateListByGtPts(ctx, channelId, md.UserId, request.Pts, limit)
		if len(updateList) < int(limit) {
			difference.Final = true
		} else {
			difference.Final = false
		}

		for _, update := range updateList {
			if update.Pts_INT32 > difference.Pts {
				difference.Pts = update.Pts_INT32
			}
			switch update.PredicateName {
			case mtproto.Predicate_updateNewChannelMessage:
				newMessage := update.GetMessage_MESSAGE()
				if newMessage.PredicateName == mtproto.Predicate_messageEmpty {
					continue
				}
				difference.NewMessages = append(difference.NewMessages, newMessage)
			case mtproto.Predicate_updateDeleteChannelMessages,
				mtproto.Predicate_updateEditChannelMessage,
				mtproto.Predicate_updateChannelWebPage:
				difference.OtherUpdates = append(difference.OtherUpdates, update)
			default:
				log.Errorf("bug: %v", update)
			}
		}

		if len(difference.NewMessages) > 0 {
			userIdList, _, _ := model.PickAllIdListByMessages(difference.NewMessages)
			difference.Users = s.UserFacade.GetUserListByIdList(ctx, md.UserId, userIdList)
			difference.Chats = []*mtproto.Chat{channel.ToUnsafeChat(md.UserId)}
		}
	}
	return difference, nil
}

func getChannelId(ctx context.Context, channel *mtproto.InputChannel) (int32, error) {
	if channel == nil {
		return 0, mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_BAD_REQUEST)
	}

	if channel.PredicateName == mtproto.Predicate_inputChannelEmpty {
		return 0, mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_INPUT_CHANNEL_EMPTY)
	}

	return channel.ChannelId, nil
}

func makeChannelDifferenceEmpty(pts int32) *mtproto.Updates_ChannelDifference {
	return mtproto.MakeTLUpdatesChannelDifferenceEmpty(&mtproto.Updates_ChannelDifference{
		Final:   true,
		Pts:     pts,
		Timeout: &types.Int32Value{Value: channelPollTimeout},
	}).To_Updates_ChannelDifference()
}
