package service

import (
	"context"
	"strings"
	"time"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/app/messenger/webpage/internal/service/webpage"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

const (
	cacheWebPageTimeout = 24 * 24 * 60
)

func (s *Service) GetPendingWebPagePreview(ctx context.Context, request *mtproto.TLMessagesGetWebPagePreview) (*mtproto.WebPage, error) {
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
		webPage = mtproto.MakeTLWebPagePending(&mtproto.WebPage{
			Id:   idgen.GetUUID(),
			Date: int32(time.Now().Unix()),
		}).To_WebPage()

		s.Dao.PutCacheWebPage(ctx, url, webPage, cacheWebPageTimeout)

		go func() {
			ctx2 := context.Background()

			webPage2 := webpage.GetWebPagePreview(url)
			s.Dao.PutCacheWebPage(ctx2, url, webPage2, cacheWebPageTimeout)

			pts := int32(idgen.NextChannelPtsId(ctx2, md.UserId))
			sync_client.SyncUpdatesMe(context.Background(),
				md.UserId,
				md.AuthId,
				md.SessionId,
				md.ServerId,
				model.MakeUpdatesByUpdates(mtproto.MakeTLUpdateWebPage(&mtproto.Update{
					Pts_INT32: pts,
					PtsCount:  1,
					Webpage:   webPage2,
				}).To_Update()))
		}()
	}

	log.Debugf("messages.getWebPagePreview - reply %s", webPage.DebugString())
	return webPage, nil
}
