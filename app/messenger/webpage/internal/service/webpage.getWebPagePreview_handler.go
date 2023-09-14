package service

import (
	"context"
	"strings"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) GetWebPagePreview(ctx context.Context, request *mtproto.TLMessagesGetWebPagePreview) (*mtproto.WebPage, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("messages.getWebPagePreview - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	url := strings.TrimSpace(request.Message)
	if url == "" {
		err := mtproto.ErrMessageEmpty
		log.Errorf("messages.getWebPagePreview - error: %v", err)
		return nil, err
	}

	webPage, _ := s.Dao.GetCacheWebPage(ctx, url)
	if webPage == nil {
		webPage = mtproto.MakeTLWebPageEmpty(&mtproto.WebPage{
			Id: 0,
		}).To_WebPage()
	}

	log.Debugf("messages.getWebPagePreview - reply %s", webPage.DebugString())
	return webPage, nil
}
