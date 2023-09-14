package service

import (
	"context"
	"math"

	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/math2"
)

func (s *Service) MessagesSearch(ctx context.Context, request *mtproto.TLMessagesSearch) (*mtproto.Messages_Messages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.search - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.search - error: %v", err)
		return nil, err
	}

	var (
		messages *mtproto.Messages_Messages
		offsetId       = request.OffsetId
		limit          = request.Limit
		minId    int32 = 0
	)

	if offsetId == 0 {
		offsetId = math.MaxInt32
	}

	if limit > 50 {
		limit = 50
	}

	peer := model.FromInputPeer2(md.UserId, request.Peer)
	switch peer.PeerType {
	case model.PEER_EMPTY, model.PEER_SELF, model.PEER_USER, model.PEER_CHAT:
		messages = mtproto.MakeTLMessagesMessages(&mtproto.Messages_Messages{
			Messages: []*mtproto.Message{},
			Chats:    []*mtproto.Chat{},
			Users:    []*mtproto.User{},
		}).To_Messages_Messages()
	case model.PEER_CHANNEL:
		minId = 0
		var channel *model.MutableChannel
		var err error
		if channel, err = s.ChannelFacade.GetMutableChannel(ctx, peer.PeerId, md.UserId); err != nil {
			err = mtproto.ErrPeerIdInvalid
			log.Errorf("messages.getHistory - error: %v", err)
			return nil, err
		}

		me := channel.GetImmutableChannelParticipant(md.UserId)
		if me != nil {
			minId = math2.Int32Max(me.AvailableMinId, request.MinId)
		}

		messages = mtproto.MakeTLMessagesChannelMessages(&mtproto.Messages_Messages{
			Messages: []*mtproto.Message{},
			Chats:    []*mtproto.Chat{},
			Users:    []*mtproto.User{},
		}).To_Messages_Messages()
	default:
	}
	fType := model.FromMessagesFilter(request.Filter)
	switch fType {
	case model.FilterPhotos:
		log.Debugf("messages.search - invalid filter: %s", request.DebugString())
	case model.FilterVideo:
		log.Debugf("messages.search - invalid filter: %s", request.DebugString())
	case model.FilterPhotoVideo:
		boxList := s.MessageFacade.SearchByMediaType(ctx, md.UserId, peer, model.MEDIA_PHOTOVIDEO, minId, offsetId, limit)
		messages.Messages, messages.Users, messages.Chats = boxList.ToMessagesPeersList(ctx, md.UserId, s.UserFacade, s.ChatFacade, s.ChannelFacade)
	case model.FilterDocument:
		boxList := s.MessageFacade.SearchByMediaType(ctx, md.UserId, peer, model.MEDIA_FILE, minId, offsetId, limit)
		messages.Messages, messages.Users, messages.Chats = boxList.ToMessagesPeersList(ctx, md.UserId, s.UserFacade, s.ChatFacade, s.ChannelFacade)
	case model.FilterUrl:
		boxList := s.MessageFacade.SearchByMediaType(ctx, md.UserId, peer, model.MEDIA_URL, minId, offsetId, limit)
		messages.Messages, messages.Users, messages.Chats = boxList.ToMessagesPeersList(ctx, md.UserId, s.UserFacade, s.ChatFacade, s.ChannelFacade)
	case model.FilterGif:
		boxList := s.MessageFacade.SearchByMediaType(ctx, md.UserId, peer, model.MEDIA_GIF, minId, offsetId, limit)
		messages.Messages, messages.Users, messages.Chats = boxList.ToMessagesPeersList(ctx, md.UserId, s.UserFacade, s.ChatFacade, s.ChannelFacade)
	case model.FilterVoice:
		log.Debugf("messages.search - invalid filter: %s", request.DebugString())
	case model.FilterMusic:
		boxList := s.MessageFacade.SearchByMediaType(ctx, md.UserId, peer, model.MEDIA_MUSIC, minId, offsetId, limit)
		messages.Messages, messages.Users, messages.Chats = boxList.ToMessagesPeersList(ctx, md.UserId, s.UserFacade, s.ChatFacade, s.ChannelFacade)
	case model.FilterChatPhotos:
	case model.FilterPhoneCalls:
		boxList := s.MessageFacade.SearchByMediaType(ctx, md.UserId, peer, model.MEDIA_PHONE_CALL, minId, offsetId, limit)
		messages.Messages, messages.Users, messages.Chats = boxList.ToMessagesPeersList(ctx, md.UserId, s.UserFacade, s.ChatFacade, s.ChannelFacade)
	case model.FilterRoundVoice:
		boxList := s.MessageFacade.SearchByMediaType(ctx, md.UserId, peer, model.MEDIA_AUDIO, minId, offsetId, limit)
		messages.Messages, messages.Users, messages.Chats = boxList.ToMessagesPeersList(ctx, md.UserId, s.UserFacade, s.ChatFacade, s.ChannelFacade)
	case model.FilterRoundVideo:
		log.Debugf("messages.search - invalid filter: %s", request.DebugString())
	case model.FilterMyMentions:
		log.Debugf("messages.search - invalid filter: %s", request.DebugString())
	case model.FilterGeo:
		log.Debugf("messages.search - invalid filter: %s", request.DebugString())
	case model.FilterContacts:
		log.Debugf("messages.search - invalid filter: %s", request.DebugString())
	case model.FilterEmpty:
		if request.Q == "" {
			err := mtproto.ErrSearchQueryEmpty
			log.Errorf("messages.search - error: %v", err)
			return nil, err
		}

		boxList := s.MessageFacade.Search(ctx, md.UserId, peer, request.Q, minId, offsetId, limit)
		messages.Messages, messages.Users, messages.Chats = boxList.ToMessagesPeersList(ctx, md.UserId, s.UserFacade, s.ChatFacade, s.ChannelFacade)
	default:
	}
	log.Debugf("messages.search#39e9ea0 - reply: %s", messages.DebugString())
	return messages, nil
}
