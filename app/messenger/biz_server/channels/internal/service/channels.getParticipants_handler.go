package service

import (
	"context"
	"github.com/pkg/errors"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

type filterData struct {
	contacts map[int32]bool
}

func (s *Service) ChannelsGetParticipants(ctx context.Context, request *mtproto.TLChannelsGetParticipants) (*mtproto.Channels_ChannelParticipants, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.getParticipants - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		err          error
		channel      *model.MutableChannel
		channelId    int32
		participants = make([]*mtproto.ChannelParticipant, 0)
		userIdList   model.IDList
	)

	if request.Limit > 200 {
		request.Limit = 200
	}

	switch request.Filter.PredicateName {
	case mtproto.Predicate_channelParticipantsRecent:
	case mtproto.Predicate_channelParticipantsAdmins:
	case mtproto.Predicate_channelParticipantsKicked:
	case mtproto.Predicate_channelParticipantsBots:
	case mtproto.Predicate_channelParticipantsBanned:
	case mtproto.Predicate_channelParticipantsSearch:
	case mtproto.Predicate_channelParticipantsContacts:
	//case mtproto.Predicate_channelParticipantsMentions:
	//	request.Filter.Q_STRING = request.Filter.GetQ_FLAGSTRING().GetValue()
	default:
		log.Errorf("channels.getParticipants#123e05e9 - channel's filter invalid: %v", request.DebugString())
		return nil, mtproto.ErrInputRequestInvalid
	}

	if !IsInputChannel(request.Channel) {
		err := mtproto.ErrChannelInvalid
		log.Errorf("channels.getParticipant - error: %v", err)
		return nil, err
	} else {
		channelId = request.Channel.ChannelId
	}

	if channel, err = s.ChannelFacade.GetMutableChannel(ctx, channelId, md.UserId); err != nil {
		log.Errorf("channels.getParticipant - error: %v", err)
		return nil, err
	}

	me := channel.GetImmutableChannelParticipant(md.UserId)

	if channel.Channel.IsBroadcast() {
		if me == nil || (me.IsStateOk() && !me.IsCreatorOrAdmin()) {
			err = mtproto.ErrChatAdminRequired
			log.Errorf("channels.getParticipant - error: %v", err)
			return nil, err
		}
	} else if channel.Channel.IsMegagroup() {

	}

	var (
		count int32
	)

	fData, err := s.prepareFilterData(ctx, md.UserId, request.Filter)
	if err != nil {
		log.Errorf("channels.getParticipant, prepareFilterData failed - error: %v", err)
		return nil, err
	}

	channel.FetchAndWalk(func() []*model.ImmutableChannelParticipant {
		var (
			participants []*model.ImmutableChannelParticipant
		)
		count, participants = s.ChannelFacade.GetChannelParticipants(ctx, channel.Channel, request.Filter, request.Offset, request.Limit)
		return participants
	}, func(participant *model.ImmutableChannelParticipant) {
		participant, err := s.filterParticipant(participant, request.Filter, fData)
		if err != nil || participant == nil {
			return
		}

		if participant.IsChatMemberNormal() {
			userIdList.AddIfNot(participant.UserId)
		} else if participant.IsChatMemberCreator() {
			userIdList.AddIfNot(participant.UserId)
		} else if participant.IsChatMemberAdmin() {
			userIdList.AddIfNot(participant.UserId, participant.InviterId, participant.PromotedBy)
		} else if participant.IsChatMemberBanned() {
			userIdList.AddIfNot(participant.UserId, participant.InviterId, participant.KickedBy)
		}
		participants = append(
			participants,
			model.GetFirstValue(participant.ToUnsafeChannelParticipant(md.UserId)).(*mtproto.ChannelParticipant))
	})

	channelParticipants := mtproto.MakeTLChannelsChannelParticipants(&mtproto.Channels_ChannelParticipants{
		Count:        count,
		Participants: participants,
		Users:        s.UserFacade.GetUserListByIdList(ctx, md.UserId, userIdList),
	}).To_Channels_ChannelParticipants()

	log.Debugf("channels.getParticipants - reply: {%s}", channelParticipants.DebugString())
	return channelParticipants, nil
}

func (s *Service) prepareFilterData(ctx context.Context, myUID int32, filter *mtproto.ChannelParticipantsFilter) (*filterData, error) {
	switch filter.GetPredicateName() {
	case mtproto.Predicate_channelParticipantsContacts:
		m := make(map[int32]bool)
		l := s.UserFacade.GetContactList(ctx, myUID)
		for _, c := range l {
			m[c.UserId] = true
		}
		out := &filterData{contacts: m}
		return out, nil
	default:
		return nil, nil
	}
}

func (s *Service) filterParticipant(participant *model.ImmutableChannelParticipant, filter *mtproto.ChannelParticipantsFilter, data *filterData) (*model.ImmutableChannelParticipant, error) {
	if filter.GetPredicateName() != mtproto.Predicate_channelParticipantsContacts {
		return participant, nil
	}

	//log.Debugf("channels.getParticipants, participant:%d, contacts:%v", participant.UserId, data.contacts)
	if data == nil || data.contacts == nil {
		log.Errorf("channels.getParticipants, filter data invalid")
		return nil, errors.New("filter data invalid")
	}

	if _, ok := data.contacts[participant.UserId]; ok {
		return participant, nil
	} else {
		return nil, nil
	}
}
