package service

import (
	"context"

	"open.chat/app/messenger/msg/msgpb"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesReadMessageContents(ctx context.Context, request *mtproto.TLMessagesReadMessageContents) (*mtproto.Messages_AffectedMessages, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.readMessageContents - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	messages := s.MessageFacade.GetUserMessageList(ctx, md.UserId, request.GetId())
	if len(messages) == 0 {
		log.Warnf("missing messages")
		return mtproto.MakeTLMessagesAffectedMessages(&mtproto.Messages_AffectedMessages{
			Pts:      int32(idgen.CurrentPtsId(ctx, md.UserId)),
			PtsCount: 0,
		}).To_Messages_AffectedMessages(), nil
	}

	peer := model.MakeUserPeerUtil(messages[0].SendUserId)
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
			log.Warnf("content has readed")
		}
	}

	affected, err := s.MsgFacade.ReadMessageContents(ctx, md.UserId, md.AuthId, peer, contents)
	if err != nil {
		log.Errorf("messages.readMessageContents - %v", err)
	}

	log.Debugf("messages.readMessageContents - reply: {%s}", affected.DebugString())
	return affected, nil
}
