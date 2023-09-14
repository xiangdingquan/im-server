package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) WalletGetKeySecretSalt(ctx context.Context, request *mtproto.TLWalletGetKeySecretSalt) (*mtproto.Wallet_KeySecretSalt, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("wallet.getKeySecretSalt - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return nil, fmt.Errorf("wallet.getKeySecretSalt - not imp WalletGetKeySecretSalt")
}
