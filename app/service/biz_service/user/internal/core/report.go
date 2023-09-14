package core

import (
	"context"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
)

func (m *UserCore) Report(ctx context.Context, userId, reportType, peerType, peerId, messageSenderUserId, messageId, reason int32, text string) (bool, error) {
	do := &dataobject.ReportsDO{
		UserId:              userId,
		ReportType:          int8(reportType),
		PeerType:            int8(peerType),
		PeerId:              peerId,
		MessageSenderUserId: messageSenderUserId,
		MessageId:           messageId,
		Reason:              int8(reason),
		Text:                text,
	}
	id, _, err := m.ReportsDAO.Insert(ctx, do)
	return id > 0, err
}

func (m *UserCore) ReportIdList(ctx context.Context, userId, reportType, peerType, peerId, messageSenderUserId int32, messageIdList []int32, reason int32, text string) (bool, error) {
	bulkDOList := make([]*dataobject.ReportsDO, 0, len(messageIdList))

	for _, id := range messageIdList {
		bulkDOList = append(bulkDOList, &dataobject.ReportsDO{
			UserId:              userId,
			ReportType:          int8(reportType),
			PeerType:            int8(peerType),
			PeerId:              peerId,
			MessageSenderUserId: messageSenderUserId,
			MessageId:           id,
			Reason:              int8(reason),
			Text:                text,
		})
	}

	lastInsertId, _, err := m.ReportsDAO.InsertBulk(ctx, bulkDOList)
	return lastInsertId > 0, err
}
