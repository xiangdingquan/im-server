package report_facade

import (
	"context"

	"open.chat/app/service/biz_service/report/internal/core"
	"open.chat/app/service/biz_service/report/internal/dao"
)

type localReportFacade struct {
	*core.ReportCore
}

func New() ReportFacade {
	return &localReportFacade{
		ReportCore: core.New(dao.New()),
	}
}

func (c *localReportFacade) Report(ctx context.Context, userId, reportType, peerType, peerId, messageSenderUserId, messageId, reason int32, text string) (bool, error) {
	return c.ReportCore.Report(ctx, userId, reportType, peerType, peerId, messageSenderUserId, messageId, reason, text)
}

func (c *localReportFacade) ReportIdList(ctx context.Context, userId, reportType, peerType, peerId, messageSenderUserId int32, messageIdList []int32, reason int32, text string) (bool, error) {
	return c.ReportCore.ReportIdList(ctx, userId, reportType, peerType, peerId, messageSenderUserId, messageIdList, reason, text)
}

func init() {
	Register("local", New)
}
