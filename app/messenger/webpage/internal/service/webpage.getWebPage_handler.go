package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) GetWebPage(ctx context.Context, request *mtproto.TLMessagesGetWebPage) (*mtproto.WebPage, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getWebPage - metadata: %s, request: %s", md.DebugString(), request.DebugString())
	webPage := mtproto.MakeTLWebPageEmpty(&mtproto.WebPage{
		Id: 0,
	}).To_WebPage()

	log.Debugf("messages.getWebPage - reply %s", webPage.DebugString())
	return webPage, nil
}
