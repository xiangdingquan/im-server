package service

import (
	"context"

	"open.chat/mtproto"
	"open.chat/pkg/crypto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountGetPassword(ctx context.Context, request *mtproto.TLAccountGetPassword) (*mtproto.Account_Password, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.getPassword#548a30f5 - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	if md.Layer == 74 {
		// 	account.noPassword#96dabc18 new_salt:bytes email_unconfirmed_pattern:string = account.Password;
		password := mtproto.MakeTLAccountNoPassword(&mtproto.Account_Password{
			NewSalt:                        crypto.RandomBytes(8),
			EmailUnconfirmedPattern_STRING: "",
		})
		log.Debugf("account.getPassword#548a30f5 - reply: %s", password.DebugString())
		return password.To_Account_Password(), nil
	} else {
		// mtproto.CRC32_account_password
		password, err := s.AccountFacade.GetPassword(ctx, md.UserId)
		if err != nil {
			log.Errorf("account.getPassword#548a30f5 - error: %v", err)
			return nil, err
		}

		log.Debugf("account.getPassword#548a30f5 - reply: %s", password.DebugString())
		return password, nil
	}
}
