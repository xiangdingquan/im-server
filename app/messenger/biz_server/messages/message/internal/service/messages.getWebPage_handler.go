package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) MessagesGetWebPage(ctx context.Context, request *mtproto.TLMessagesGetWebPage) (*mtproto.WebPage, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getWebPage#32ca8f91 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.IsBot {
		err := mtproto.ErrBotMethodInvalid
		log.Errorf("messages.getWebPagePreview - error: %v", err)
		return nil, err
	}

	ctx, _ = grpc_util.RpcMetadataToOutgoing(ctx, md)
	webpage, err := s.RPCWebPageClient.GetWebPage(ctx, request)
	if err != nil {
		log.Errorf("messages.getWebPagePreview - error: %v", err)
		return nil, err
	}

	log.Debugf("messages.getWebPagePreview#25223e24 - reply: %s\n", webpage.DebugString())
	return webpage, nil
}
