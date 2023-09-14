package service

import (
	"context"

	"open.chat/app/service/auth_session/authsessionpb"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) SessionGetLangCode(ctx context.Context, request *authsessionpb.TLSessionGetLangCode) (*mtproto.String, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("session.getLangCode - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	keyData, err := s.AuthSessionCore.Dao.GetAuthKey(ctx, request.GetAuthKeyId())
	if err != nil {
		log.Errorf("session.getLangCode - error: %v", err)
		return nil, err
	}

	langCode := s.AuthSessionCore.GetLangCode(ctx, keyData.PermAuthKeyId)
	log.Debugf("langCode = %s", langCode)
	reply := &mtproto.TLString{Data2: &mtproto.String{
		V: langCode,
	}}

	log.Debugf("session.getLangCode - reply: {%s}", reply.DebugString())
	return reply.To_String(), nil
}
