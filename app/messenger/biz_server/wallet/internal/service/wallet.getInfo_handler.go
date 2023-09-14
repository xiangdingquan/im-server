package service

import (
	"context"

	"open.chat/pkg/log"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
)

// wallet.getInfo#86f1ee4f = wallet.Info;
func (s *Service) WalletGetInfo(ctx context.Context, request *mtproto.TLWalletGetInfo) (*mtproto.Wallet_Info, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("wallet.getInfo - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		wallet *mtproto.Wallet_Info
		err    error
	)

	wallet, err = s.WalletFacade.GetInfo(ctx, md.GetUserId())
	if err != nil {
		err = mtproto.ErrButtonTypeInvalid
		log.Errorf("blogs.reward: %v", err)
		return nil, err
	}

	return wallet, nil
}
