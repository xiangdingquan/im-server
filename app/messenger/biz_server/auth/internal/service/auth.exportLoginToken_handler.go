package service

import (
	"context"
	"time"

	"open.chat/app/messenger/biz_server/auth/internal/model"
	"open.chat/mtproto"
	"open.chat/pkg/crypto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

const (
	qrCodeTimeout = 60 // salt timeout
)

func (s *Service) AuthExportLoginToken(ctx context.Context, request *mtproto.TLAuthExportLoginToken) (*mtproto.Auth_LoginToken, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("auth.exportLoginToken - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	qrCode, err := s.AuthCore.GetQRCode(ctx, md.AuthId)
	if err != nil {
		log.Errorf("getQRCode - error: %v", err)
		return nil, err
	} else if qrCode == nil {
		qrCode = &model.QRCodeTransaction{
			AuthKeyId: md.AuthId,
			SessionId: md.SessionId,
			ServerId:  md.ServerId,
			ApiId:     request.ApiId,
			ApiHash:   request.ApiHash,
			CodeHash:  crypto.GenerateStringNonce(16),
			ExpireAt:  time.Now().Unix() + qrCodeTimeout,
			UserId:    0,
			State:     model.QRCodeStateNew,
		}
		log.Debugf("putQRCode - %#v", qrCode)
		if err = s.AuthCore.PutQRCode(ctx, md.AuthId, qrCode, qrCodeTimeout+2); err != nil {
			log.Errorf("putQRCode - error: %v", err)
			return nil, err
		}
	} else {
		log.Debugf("putQRCode - %#v", qrCode)
	}

	var qrLoginToken *mtproto.Auth_LoginToken

	switch qrCode.State {
	case model.QRCodeStateAccepted, model.QRCodeStateSuccess:
		if sessionPasswordNeeded, err := s.AccountFacade.CheckSessionPasswordNeeded(ctx, qrCode.UserId); sessionPasswordNeeded {
			log.Infof("auth.exportLoginToken - registered, next step auth.checkPassword: %v", err)
			err = mtproto.ErrSessionPasswordNeeded
			return nil, err
		}

		user, err := s.UserFacade.GetUserSelf(ctx, qrCode.UserId)
		if err != nil {
			log.Errorf("auth.exportLoginToken - error: %v", err)
			return nil, err
		}
		qrLoginToken = mtproto.MakeTLAuthLoginTokenSuccess(&mtproto.Auth_LoginToken{
			Authorization: mtproto.MakeTLAuthAuthorization(&mtproto.Auth_Authorization{
				TmpSessions: nil,
				User:        user,
			}).To_Auth_Authorization(),
		}).To_Auth_LoginToken()
	default:
		qrLoginToken = mtproto.MakeTLAuthLoginToken(&mtproto.Auth_LoginToken{
			Expires: int32(qrCode.ExpireAt),
			Token:   qrCode.Token(),
		}).To_Auth_LoginToken()
	}

	log.Debugf("auth.exportLoginToken - result: %s", qrLoginToken.DebugString())
	return qrLoginToken, nil
}
