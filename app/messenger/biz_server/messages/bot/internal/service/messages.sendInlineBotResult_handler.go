package service

import (
	"context"
	"time"

	"github.com/gogo/protobuf/types"
	"open.chat/app/messenger/msg/msgpb"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesSendInlineBotResult(ctx context.Context, request *mtproto.TLMessagesSendInlineBotResult) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.sendInlineBotResult - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	// 400	CHAT_ADMIN_REQUIRED	You must be an admin in this chat to do this
	// 403	CHAT_SEND_INLINE_FORBIDDEN	You can't send inline messages in this group
	// 403	CHAT_SEND_MEDIA_FORBIDDEN	You can't send media in this chat
	// 403	CHAT_WRITE_FORBIDDEN	You can't write in this chat
	// 400	INLINE_RESULT_EXPIRED	The inline query expired
	// 400	PEER_ID_INVALID	The provided peer id is invalid
	// 400	QUERY_ID_EMPTY	The query ID is empty
	// 400	RESULT_ID_EMPTY	Result ID empty
	// 400	USER_BANNED_IN_CHANNEL	You're banned from sending messages in supergroups/channels
	// 400	WEBPAGE_CURL_FAILED	Failure while fetching the webpage with cURL
	// 400	WEBPAGE_MEDIA_EMPTY	Webpage media empty
	// 400	YOU_BLOCKED_USER	You blocked this user

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.sendInlineBotResult - error: %v", err)
		return nil, err
	}

	// peer
	var (
		err  error
		peer = model.FromInputPeer2(md.UserId, request.Peer)
	)

	switch peer.PeerType {
	case model.PEER_SELF:
		peer.PeerType = model.PEER_USER
	case model.PEER_USER:
	case model.PEER_CHAT:
	case model.PEER_CHANNEL:
	default:
		log.Errorf("invalid peer: %v", request.Peer)
		err = mtproto.ErrPeerIdInvalid
		return nil, err
	}

	cacheInlineBotResult, err := s.Dao.GetCacheInlineBotResults(ctx, request.QueryId, request.Id)
	if err != nil {
		log.Errorf("messages.sendInlineBotResult - error: %v", err)
		return nil, err
	}

	outMessage := mtproto.MakeTLMessage(&mtproto.Message{
		Out:               true,
		Mentioned:         false,
		MediaUnread:       false,
		Silent:            request.Silent,
		Post:              false,
		FromScheduled:     false,
		Legacy:            false,
		EditHide:          false,
		Id:                0,
		FromId_FLAGPEER:   model.MakePeerUser(md.UserId),
		ToId:              peer.ToPeer(),
		FwdFrom:           nil,
		ViaBotId:          &types.Int32Value{Value: cacheInlineBotResult.BotId},
		ReplyTo:           nil,
		Date:              int32(time.Now().Unix()),
		Message:           "",
		Media:             nil,
		ReplyMarkup:       nil,
		Entities:          nil,
		Views:             nil,
		Forwards:          nil,
		EditDate:          nil,
		PostAuthor:        nil,
		GroupedId:         nil,
		RestrictionReason: nil,
	}).To_Message()

	if request.ReplyToMsgId != nil {
		outMessage.ReplyTo = mtproto.MakeTLMessageReplyHeader(
			&mtproto.MessageReplyHeader{
				ReplyToMsgId: request.ReplyToMsgId.GetValue(),
			}).To_MessageReplyHeader()
	}

	outMessage.Media = mtproto.MakeTLMessageMediaDocument(&mtproto.MessageMedia{
		Document:   cacheInlineBotResult.BotInlineResult.Document,
		TtlSeconds: nil,
	}).To_MessageMedia()

	meUpdates, err := s.MsgFacade.SendMessage(ctx, md.UserId, md.AuthId, peer, &msgpb.OutboxMessage{
		NoWebpage:    true,
		Background:   request.Background,
		RandomId:     request.RandomId,
		Message:      outMessage,
		ScheduleDate: request.ScheduleDate,
	})

	if err != nil {
		log.Errorf("messages.sendInlineBotResult - error: %v", err)
		return nil, err
	}

	return model.WrapperGoFunc(meUpdates, func() {
		if request.ClearDraft {
			s.doClearDraft(context.Background(), md.UserId, md.AuthId, peer)
		}

		s.UserFacade.UpdateUserStatus(context.Background(), md.UserId, time.Now().Unix())
	}).(*mtproto.Updates), nil
}

func (s *Service) doClearDraft(ctx context.Context, userId int32, authKeyId int64, peer *model.PeerUtil) {
	switch peer.PeerType {
	case model.PEER_USER:
		s.PrivateFacade.ClearDraftMessage(ctx, userId, peer.PeerId)
	case model.PEER_CHAT:
		s.ChatFacade.ClearDraftMessage(ctx, userId, peer.PeerId)
	case model.PEER_CHANNEL:
		s.ChannelFacade.ClearDraftMessage(ctx, userId, peer.PeerId)
	default:
	}
}
