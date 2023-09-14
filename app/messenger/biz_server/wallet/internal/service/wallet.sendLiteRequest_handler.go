package service

import (
	"context"
	"fmt"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"
)

func (s *Service) WalletSendLiteRequest(ctx context.Context, request *mtproto.TLWalletSendLiteRequest) (*mtproto.Wallet_LiteResponse, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("wallet.sendLiteRequest - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	return nil, fmt.Errorf("wallet.sendLiteRequest - not imp WalletSendLiteRequest")
}
