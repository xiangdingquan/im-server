package service

import (
	"context"
	"encoding/binary"
	"time"

	"open.chat/app/messenger/biz_server/auth/internal/model"
	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/app/service/auth_session/authsessionpb"
	model2 "open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AuthAcceptLoginToken(ctx context.Context, request *mtproto.TLAuthAcceptLoginToken) (*mtproto.Authorization, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("auth.acceptLoginToken - metadata: %s, request: %s", md.DebugString(), request.DebugString())
	if len(request.Token) != 24 {
		err := mtproto.ErrAuthTokenInvalid
		log.Errorf("auth.acceptLoginToken - error: %v", err)
		return nil, err
	}

	var (
		keyId = int64(binary.BigEndian.Uint64(request.Token))
	)

	qrCode, err := s.AuthCore.GetQRCode(ctx, keyId)
	if err != nil || qrCode == nil {
		err := mtproto.ErrAuthTokenExpired
		log.Errorf("auth.acceptLoginToken - error: %v", err)
		return nil, err
	}

	log.Debugf("auth.acceptLoginToken - error: %#v", qrCode)

	if !qrCode.CheckByToken(request.Token) {
		err := mtproto.ErrAuthTokenInvalid
		log.Errorf("auth.acceptLoginToken - error: %v", err)
		return nil, err
	}

	if qrCode.ExpireAt >= time.Now().Unix() {
	}

	switch qrCode.State {
	case model.QRCodeStateNew:
	case model.QRCodeStateAccepted, model.QRCodeStateSuccess:
		err := mtproto.ErrAuthTokenAccepted
		log.Errorf("auth.acceptLoginToken - error: %v", err)
		return nil, err
	default:
		err := mtproto.ErrAuthTokenInvalid
		log.Errorf("auth.acceptLoginToken - error: %v", err)
		return nil, err
	}

	s.AuthCore.UpdateQRCode(ctx, keyId, map[string]interface{}{
		"user_id": md.UserId,
		"state":   model.QRCodeStateAccepted,
	})

	user, err := s.UserFacade.GetUserSelf(ctx, md.UserId)
	if err != nil {
		log.Errorf("auth.acceptLoginToken - error: %v", err)
		return nil, err
	}

	hash, err := s.AuthSessionRpcClient.SessionBindAuthKeyUser(ctx, &authsessionpb.TLSessionBindAuthKeyUser{
		AuthKeyId: qrCode.AuthKeyId,
		UserId:    user.GetId(),
	})
	if err != nil {
		log.Errorf("auth.acceptLoginToken - error: %v", err)
		return nil, err
	}
	authorization, err := s.AuthSessionRpcClient.SessionGetAuthorization(ctx, &authsessionpb.TLSessionGetAuthorization{
		AuthKeyId: qrCode.AuthKeyId,
	})
	if err != nil {
		log.Errorf("auth.acceptLoginToken - error: %v", err)
		return nil, err
	}

	authorization.DateCreated = int32(time.Now().Unix())
	authorization.DateActive = authorization.DateCreated
	authorization.Hash = hash.V

	log.Debugf("auth.acceptLoginToken - result: %v", authorization.DebugString())
	return model2.WrapperGoFunc(authorization, func() {
		sync_client.SyncUpdatesMe(ctx,
			user.Id,
			qrCode.AuthKeyId,
			qrCode.SessionId,
			qrCode.ServerId,
			mtproto.MakeTLUpdateShort(&mtproto.Updates{
				Update: mtproto.MakeTLUpdateLoginToken(nil).To_Update(),
				Date:   int32(time.Now().Unix()),
			}).To_Updates())
	}).(*mtproto.Authorization), nil
}
