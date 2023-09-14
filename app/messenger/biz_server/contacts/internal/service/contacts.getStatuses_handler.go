package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
	"open.chat/pkg/logger"
)

func (s *Service) ContactsGetStatuses(ctx context.Context, request *mtproto.TLContactsGetStatuses) (*mtproto.Vector_ContactStatus, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("contacts.getStatuses#c4a353ee - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	cList := s.UserFacade.GetContactList(ctx, md.UserId)

	statusList := &mtproto.Vector_ContactStatus{
		Datas: make([]*mtproto.ContactStatus, 0, len(cList)),
	}

	for _, c := range cList {
		status := s.UserFacade.GetUserStatus2(ctx, md.UserId, c.UserId, true, false)
		if status != nil {
			contactStatus := mtproto.MakeTLContactStatus(&mtproto.ContactStatus{
				UserId: c.UserId,
				Status: status,
			})
			statusList.Datas = append(statusList.Datas, contactStatus.To_ContactStatus())
		}
	}

	log.Debugf("contacts.getStatuses#c4a353ee - reply: {%s}", logger.JsonDebugData(statusList))
	return statusList, nil
}
