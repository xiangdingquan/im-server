package blog_facade

import (
	"context"

	"open.chat/app/service/biz_service/wallet/internal/core"
	"open.chat/app/service/biz_service/wallet/internal/dao"
	"open.chat/model"
	"open.chat/mtproto"
)

type localWalletFacade struct {
	*core.WalletCore
}

func New() WalletFacade {
	return &localWalletFacade{
		WalletCore: core.New(dao.New()),
	}
}

func (b *localWalletFacade) GetInfo(ctx context.Context, fromUserId int32) (*mtproto.Wallet_Info, error) {
	return b.WalletCore.GetInfo(ctx, fromUserId)
}

func (b *localWalletFacade) CheckPassword(ctx context.Context, fromUserId int32, password string) (bool, error) {
	return true, nil
}

func (b *localWalletFacade) RewardBlog(ctx context.Context, fromUserId int32, userId int32, blogId int32, amount float64) (bool, error) {
	return b.WalletCore.RewardBlog(ctx, fromUserId, userId, blogId, amount)
}

func (b *localWalletFacade) GetRecords(ctx context.Context, fromUserId int32, recordType *model.WalletRecordTypeUtil, date int32, offset int32, limit int32) (records *mtproto.Wallet_Records, err error) {
	if recordType == nil || recordType.WalletRecordType == model.WalletRecordType_Invalid {
		records, err = b.WalletCore.GetAllRecords(ctx, fromUserId, date, offset, limit)
	} else {
		records, err = b.WalletCore.GetRecordsByType(ctx, fromUserId, recordType.WalletRecordType, date, offset, limit)
	}
	return records, err
}

func init() {
	Register("local", New)
}
