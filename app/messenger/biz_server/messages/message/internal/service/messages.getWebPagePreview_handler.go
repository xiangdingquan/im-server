package service

import (
	"context"
	"strings"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetWebPagePreview(ctx context.Context, request *mtproto.TLMessagesGetWebPagePreview) (*mtproto.MessageMedia, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getWebPagePreview - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.getWebPagePreview - error: %v", err)
		return nil, err
	}

	url := strings.TrimSpace(request.Message)
	if url == "" {
		err := mtproto.ErrMessageEmpty
		log.Errorf("messages.getWebPagePreview - error: %v", err)
		return nil, err
	}

	ctx, _ = grpc_util.RpcMetadataToOutgoing(ctx, md)
	webpage, err := s.RPCWebPageClient.GetPendingWebPagePreview(ctx, request)
	if err != nil {
		log.Errorf("messages.getWebPagePreview - error: %v", err)
		return nil, err
	}
	media := mtproto.MakeTLMessageMediaWebPage(&mtproto.MessageMedia{
		Webpage: webpage,
	})

	log.Debugf("messages.getWebPagePreview - reply: %s", media.DebugString())
	return media.To_MessageMedia(), nil
}
