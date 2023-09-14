package service

import (
	"context"
	"time"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) AccountGetTmpPassword(ctx context.Context, request *mtproto.TLAccountGetTmpPassword) (*mtproto.Account_TmpPassword, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("account.getTmpPassword#4a82327e - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	// Check password_hash invalid, android source code
	// byte[] hash = new byte[currentPassword.current_salt.length * 2 + passwordBytes.length];
	// System.arraycopy(currentPassword.current_salt, 0, hash, 0, currentPassword.current_salt.length);
	// System.arraycopy(passwordBytes, 0, hash, currentPassword.current_salt.length, passwordBytes.length);
	// System.arraycopy(currentPassword.current_salt, 0, hash, hash.length - currentPassword.current_salt.length, currentPassword.current_salt.length);

	// account.tmpPassword#db64fd34 tmp_password:bytes valid_until:int = account.TmpPassword;
	tmpPassword := mtproto.MakeTLAccountTmpPassword(&mtproto.Account_TmpPassword{
		TmpPassword: []byte("01234567899876543210"),
		ValidUntil:  int32(time.Now().Unix()) + request.Period,
	})

	log.Debugf("account.getTmpPassword#4a82327e - reply: %s", tmpPassword.DebugString())
	return tmpPassword.To_Account_TmpPassword(), nil
}
