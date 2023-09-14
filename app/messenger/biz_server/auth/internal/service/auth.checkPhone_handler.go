package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AuthCheckPhone(ctx context.Context, request *mtproto.TLAuthCheckPhone) (*mtproto.Auth_CheckedPhone, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("auth.checkPhone#6fe51dfb - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	phoneRegistered, _ := s.UserFacade.CheckPhoneNumberExist(ctx, request.PhoneNumber)
	reply := &mtproto.Auth_CheckedPhone{
		PredicateName:   mtproto.Predicate_auth_checkedPhone,
		PhoneRegistered: mtproto.ToBool(phoneRegistered),
	}

	log.Debugf("auth.checkPhone#6fe51dfb - reply: %s", reply.DebugString())
	return reply, nil
}
