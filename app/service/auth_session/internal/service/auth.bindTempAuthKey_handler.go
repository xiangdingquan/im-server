package service

import (
	"context"
	"encoding/base64"
	"fmt"

	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/crypto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AuthBindTempAuthKey(ctx context.Context, request *mtproto.TLAuthBindTempAuthKey) (reply *mtproto.Bool, err error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("auth.bindTempAuthKey - metadata: %s, request: %s", md.DebugString(), request.DebugString())
	keyData, err := s.AuthSessionCore.Dao.GetAuthKey(ctx, request.GetPermAuthKeyId())
	if err != nil {
		log.Errorf("auth.bindTempAuthKey - error: %v", err)
		do, _ := s.AuthKeysDAO.SelectByAuthKeyId(ctx, request.GetPermAuthKeyId())
		if do == nil {
			err := fmt.Errorf("not find key - keyId = %d", request.GetPermAuthKeyId())
			return nil, err
		}
		authKey, err := base64.RawStdEncoding.DecodeString(do.Body)
		if err != nil {
			log.Errorf("read keyData error - keyId = %d, %v", request.GetPermAuthKeyId(), err)
			return nil, err
		}
		keyInfo := &authsessionpb.AuthKeyInfo{
			AuthKeyId:          request.GetPermAuthKeyId(),
			AuthKey:            authKey,
			AuthKeyType:        model.AuthKeyTypePerm,
			PermAuthKeyId:      request.GetPermAuthKeyId(),
			TempAuthKeyId:      0,
			MediaTempAuthKeyId: 0,
		}

		keyData = &model.AuthKeyData{
			AuthKeyId:          keyInfo.AuthKeyId,
			AuthKeyType:        int(keyInfo.AuthKeyType),
			AuthKey:            keyInfo.AuthKey,
			PermAuthKeyId:      keyInfo.PermAuthKeyId,
			TempAuthKeyId:      keyInfo.TempAuthKeyId,
			MediaTempAuthKeyId: keyInfo.MediaTempAuthKeyId,
		}
		s.Dao.PutAuthKey(ctx, request.GetPermAuthKeyId(), keyData, 0)
	}

	permAuthKey := crypto.NewAuthKey(request.PermAuthKeyId, keyData.AuthKey)
	innerData, err := permAuthKey.AesIgeDecryptV1(request.EncryptedMessage[8:8+16], request.EncryptedMessage[8+16:])
	if err != nil {
		log.Errorf("auth.bindTempAuthKey - error: %v", err)
		return nil, err
	}

	dbuf := mtproto.NewDecodeBuf(innerData[32:])
	o := dbuf.Object()
	if dbuf.GetError() != nil {
		log.Errorf("auth.bindTempAuthKey - error: %v", dbuf.GetError())
		return nil, dbuf.GetError()
	} else if bindAuthKeyInner, ok := o.(*mtproto.TLBindAuthKeyInner); !ok {
		log.Errorf("auth.bindTempAuthKey - invalid innerData")
		return nil, mtproto.ErrInternelServerError
	} else {
		log.Debugf("auth.bindTempAuthKey - bind_auth_key_inner: %s", bindAuthKeyInner.DebugString())
		tempKeyData, err := s.AuthSessionCore.Dao.GetAuthKey(ctx, bindAuthKeyInner.GetTempAuthKeyId())
		if err != nil {
			log.Errorf("auth.bindTempAuthKey - invalid innerData")
			return nil, mtproto.ErrInternelServerError
		}

		s.AuthSessionCore.Dao.UnsafeBindKeyId(ctx,
			bindAuthKeyInner.GetPermAuthKeyId(),
			tempKeyData.AuthKeyType,
			bindAuthKeyInner.GetTempAuthKeyId())
		s.AuthSessionCore.Dao.UnsafeBindKeyId(ctx,
			bindAuthKeyInner.GetTempAuthKeyId(),
			model.AuthKeyTypePerm,
			bindAuthKeyInner.GetPermAuthKeyId())
	}

	log.Debugf("session.bindTempAuthKey - reply: {true}")
	return mtproto.BoolTrue, nil
}
