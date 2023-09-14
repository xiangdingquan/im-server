package service

import (
	"context"

	"open.chat/app/messenger/msg/msgpb"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) ChannelsReadMessageContents(ctx context.Context, request *mtproto.TLChannelsReadMessageContents) (*mtproto.Bool, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Infof("channels.readMessageContents - metadata: %s, request: %s", md.DebugString(), request.DebugString())
	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("channels.readMessageContents - error: %v", err)
		return nil, err
	}

	if !IsInputChannel(request.Channel) {
		err := mtproto.ErrChannelInvalid
		log.Errorf("channels.readMessageContents - error: %v", err)
		return nil, err
	}

	messages := s.MessageFacade.GetChannelMessageList(ctx, md.UserId, request.Channel.ChannelId, request.GetId())
	if len(messages) == 0 {
		log.Warnf("missing messages")
		return mtproto.BoolTrue, nil
	}

	contents := make([]*msgpb.ContentMessage, 0, len(messages))
	for _, m := range messages {
		if m.Message.GetMentioned() {
			contents = append(contents, &msgpb.ContentMessage{
				Id:          m.MessageId,
				IsMentioned: true,
			})
		} else if m.Message.GetMediaUnread() {
			contents = append(contents, &msgpb.ContentMessage{
				Id:          m.MessageId,
				IsMentioned: false,
			})
		} else {
			log.Warnf("channels.readMessageContents - content has read")
		}
	}

	if len(contents) == 0 {
		log.Warnf("missing messages")
		return mtproto.BoolTrue, nil
	}

	_, err := s.MsgFacade.ReadMessageContents(ctx,
		md.UserId,
		md.AuthId,
		model.MakeChannelPeerUtil(request.Channel.ChannelId),
		contents)

	if err != nil {
		log.Errorf("channels.readMessageContents - %v", err)
	}

	log.Debugf("channels.readMessageContents - reply: {true}")
	return mtproto.BoolTrue, nil
}
