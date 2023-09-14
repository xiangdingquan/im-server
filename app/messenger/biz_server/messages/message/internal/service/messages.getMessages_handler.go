package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetMessages(ctx context.Context, request *mtproto.TLMessagesGetMessages) (*mtproto.Messages_Messages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	//log.Debugf("messages.getMessages#63c66506 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		idList []int32
	)

	switch request.Constructor {
	case mtproto.CRC32_messages_getMessages_4222fa74:
		idList = request.Id_VECTORINT32
	case mtproto.CRC32_messages_getMessages_63c66506:
		for _, id := range request.Id_VECTORINPUTMESSAGE {
			switch id.PredicateName {
			case mtproto.Predicate_inputMessageID:
				idList = append(idList, id.Id)
			default:
				err := mtproto.ErrInputConstructorInvalid
				log.Errorf("messages.getMessages - error: %v", err)
				return nil, err
			}
		}
	default:
		err := mtproto.ErrInputConstructorInvalid
		log.Errorf("messages.getMessages - error: %v", err)
		return nil, err
	}

	messages := s.MessageFacade.GetUserMessageList(ctx, md.UserId, idList)
	messagesMessages := mtproto.MakeTLMessagesMessages(&mtproto.Messages_Messages{
		Count: 0,
	}).To_Messages_Messages()
	messagesMessages.Messages, messagesMessages.Users, messagesMessages.Chats = messages.ToMessagesPeersList(ctx, md.UserId, s.UserFacade, s.ChatFacade, s.ChannelFacade)

	//log.Debugf("messages.getMessages - reply: %s", messagesMessages.DebugString())
	return messagesMessages, nil
}
