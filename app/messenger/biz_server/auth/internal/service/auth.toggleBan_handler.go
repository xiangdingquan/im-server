package service

import (
	"context"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AuthToggleBan(ctx context.Context, request *mtproto.TLAuthToggleBan) (*mtproto.PredefinedUser, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.toggleBan - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// check
	if !md.IsAdmin {
		err := mtproto.ErrApiServerNeeded
		log.Errorf("account.toggleBan - error: %v", err)
		return nil, err
	}

	if request.Phone == "" {
		err := mtproto.ErrPhoneNumberInvalid
		log.Errorf("account.toggleBan - error: %v", err)
		return nil, err
	}

	user, err := s.UserFacade.GetUserSelfByPhoneNumber(ctx, request.Phone)
	if err != nil {
		log.Errorf("account.toggleBan - error: %v", err)
		err := mtproto.ErrInternelServerError
		return nil, err
	} else if user == nil {
		log.Warnf("account.toggleBan - user not registered: %s", request.Phone)
	}

	if request.GetExpires() != nil {
		s.BannedFacade.Ban(ctx, request.Phone, request.GetExpires().GetValue(), request.GetReason().GetValue())
		if user != nil {
			tKeyIdList, err := s.RPCSessionClient.SessionResetAuthorization(ctx, &authsessionpb.TLSessionResetAuthorization{
				UserId:    user.Id,
				AuthKeyId: 0,
				Hash:      0,
			})

			if err != nil {
				log.Errorf("account.resetAuthorization#df77f3bc - error: %v", err)
				return nil, err
			}

			for _, id := range tKeyIdList.Datas {
				upds := mtproto.MakeTLUpdateAccountResetAuthorization(&mtproto.Updates{
					UserId:    user.Id,
					AuthKeyId: id,
				}).To_Updates()

				sync_client.SyncUpdatesMe(ctx, user.Id, id, 0, "", upds)
			}

			s.RPCSessionClient.SessionUnbindAuthKeyUser(ctx, &authsessionpb.TLSessionUnbindAuthKeyUser{
				AuthKeyId: 0,
				UserId:    user.Id,
			})
		}
	} else {
		s.BannedFacade.UnBan(ctx, request.Phone)
	}

	predefinedUser, err := s.UserFacade.GetPredefinedUser(ctx, request.Phone)
	if err != nil {
		log.Errorf("account.toggleBan - error: {%v}", err)
		return nil, err
	}

	log.Debugf("account.toggleBan - reply: {%s}", predefinedUser.DebugString())
	return predefinedUser, nil
}
