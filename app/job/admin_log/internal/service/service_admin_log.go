package service

import (
	"context"
	"time"

	"open.chat/app/job/admin_log/internal/dal/dataobject"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/hack"
)

func (s *Service) onChannelAdminEventAction(ctx context.Context, aId, channelId int32, action *mtproto.ChannelAdminLogEventAction) error {
	_, _, err := s.ChannelAdminLogsDAO.Insert(ctx, &dataobject.ChannelAdminLogsDO{
		UserId:    aId,
		ChannelId: channelId,
		Event:     int32(model.FromChannelAdminLogEventAction(action)),
		EventData: hack.String(model.TLObjectToJson(action)),
		Query:     "",
		Date2:     int32(time.Now().Unix()),
	})
	return err
}
