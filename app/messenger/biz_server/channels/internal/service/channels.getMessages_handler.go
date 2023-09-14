package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsGetMessages(ctx context.Context, request *mtproto.TLChannelsGetMessages) (*mtproto.Messages_Messages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("channels.getMessages - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if !IsInputChannel(request.Channel) {
		err := mtproto.ErrChannelInvalid
		log.Errorf("channels.getMessages - error: %v", err)
		return nil, err
	}

	var (
		channelId = request.Channel.ChannelId
		idList    []int32
	)

	channel, err := s.ChannelFacade.GetMutableChannel(ctx, channelId, md.UserId)
	if err != nil {
		log.Errorf("channels.getMessages - error: %v", err)
		return nil, err
	}

	switch request.Constructor {
	case mtproto.CRC32_channels_getMessages_93d7b347:
		idList = request.Id_VECTORINT32
	case mtproto.CRC32_channels_getMessages_ad8c9a23:
		for _, id := range request.Id_VECTORINPUTMESSAGE {
			switch id.PredicateName {
			case mtproto.Predicate_inputMessageID:
				idList = append(idList, id.Id)
			case mtproto.Predicate_inputMessageReplyTo:
				idList = append(idList, id.Id)
			case mtproto.Predicate_inputMessagePinned:
				idList = append(idList, channel.Channel.PinnedMsgId)
			}
		}
	default:
		err := mtproto.ErrInputConstructorInvalid
		log.Errorf("channels.getMessages - error: %v", err)
		return nil, err
	}

	if len(idList) == 0 {
		err := mtproto.ErrMessageIdsEmpty
		log.Errorf("channels.getMessages - error: %v", err)
		return nil, err
	}

	boxMsgList := s.MessageFacade.GetChannelMessageList(ctx, md.UserId, channelId, idList)
	messagesMessages := mtproto.MakeTLMessagesChannelMessages(&mtproto.Messages_Messages{
		Messages: nil,
		Chats:    []*mtproto.Chat{channel.ToUnsafeChat(md.UserId)},
		Users:    nil,
		Pts:      channel.Channel.Pts,
		Count:    int32(len(boxMsgList)),
	}).To_Messages_Messages()
	messagesMessages.Messages, messagesMessages.Users, _ = boxMsgList.ToMessagesPeersList(ctx, md.UserId, s.UserFacade, nil, nil)

	log.Debugf("channels.getMessages - reply: %s", messagesMessages.DebugString())
	return messagesMessages, nil
}
