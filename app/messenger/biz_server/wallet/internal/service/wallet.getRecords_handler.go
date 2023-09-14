package service

import (
	"context"

	"open.chat/model"
	"open.chat/pkg/log"

	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
)

// wallet.getRecords#42d1204f flags:# type:flags.0?wallet.RecordType date:flags.1?int offset:int limit:int = wallet.Records;
func (s *Service) WalletGetRecords(ctx context.Context, request *mtproto.TLWalletGetRecords) (*mtproto.Wallet_Records, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("wallet.getRecords - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		reply          *mtproto.Wallet_Records
		err            error
		recordTypeUtil *model.WalletRecordTypeUtil
	)

	if request.GetType() != nil {
		recordTypeUtil = model.FromWalletRecordTypeUtil(request.GetType())
	}

	reply, err = s.WalletFacade.GetRecords(ctx, md.GetUserId(), recordTypeUtil, request.GetDate().GetValue(), request.GetOffset(), request.GetLimit())
	if err != nil {
		err = mtproto.ErrButtonTypeInvalid
		log.Errorf("blogs.reward: %v", err)
		return nil, err
	}

	return reply, nil
}
