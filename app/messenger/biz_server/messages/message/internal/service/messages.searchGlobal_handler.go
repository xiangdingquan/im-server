package service

import (
	"context"
	"math"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesSearchGlobal(ctx context.Context, request *mtproto.TLMessagesSearchGlobal) (*mtproto.Messages_Messages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.searchGlobal - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.searchGlobal - error: %v", err)
		return nil, err
	}

	if request.Q == "" {
		err := mtproto.ErrSearchQueryEmpty
		log.Errorf("messages.searchGlobal - error: %v", err)
		return nil, err
	}

	var (
		offsetId = request.OffsetId
		limit    = request.Limit
	)

	if offsetId == 0 {
		offsetId = math.MaxInt32
	}

	if limit > 50 {
		limit = 50
	}

	messages := mtproto.MakeTLMessagesMessages(&mtproto.Messages_Messages{
		Messages: []*mtproto.Message{},
		Chats:    []*mtproto.Chat{},
		Users:    []*mtproto.User{},
	}).To_Messages_Messages()

	boxList := s.MessageFacade.SearchGlobal(ctx, md.UserId, request.Q, offsetId, limit)
	messages.Messages, messages.Users, messages.Chats = boxList.ToMessagesPeersList(ctx, md.UserId, s.UserFacade, s.ChatFacade, s.ChannelFacade)

	log.Debugf("messages.searchGlobal - reply: %s", messages.DebugString())
	return messages, nil
}
