package service

import (
	"context"
	"math/rand"
	"time"

	"open.chat/app/admin/api_server/api"
	"open.chat/model"
	"open.chat/mtproto"
)

func (s *Service) PushServiceNotifications(ctx context.Context, i *api.PushServiceNotificationRequest) (err error) {
	var fromId int32 = 777000
	outBox := mtproto.MakeTLMessage(&mtproto.Message{
		Out:             true,
		Date:            int32(time.Now().Unix()),
		FromId_FLAGPEER: model.MakePeerUser(fromId),
		ToId:            model.MakePeerUser(0),
		Message:         i.Text,
		Entities:        i.Entities,
	}).To_Message()

	for _, id := range i.PushIdList {
		outBox.PeerId.UserId = id
		outBox.ToId.UserId = id
		s.MsgFacade.PushUserMessage(ctx, 1, fromId, id, rand.Int63(), outBox)
	}

	return nil
}
