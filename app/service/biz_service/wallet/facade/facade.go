package blog_facade

import (
	"context"
	"fmt"

	"open.chat/model"
	"open.chat/mtproto"
)

type WalletFacade interface {
	GetInfo(ctx context.Context, fromUserId int32) (*mtproto.Wallet_Info, error)
	CheckPassword(ctx context.Context, fromUserId int32, password string) (bool, error)
	//CanReward(ctx context.Context, fromUserId int32, userId int32, blogId int32) (bool, error)
	RewardBlog(ctx context.Context, fromUserId int32, userId int32, blogId int32, amount float64) (bool, error)

	GetRecords(ctx context.Context, fromUserId int32, recordType *model.WalletRecordTypeUtil, date int32, offset int32, limit int32) (*mtproto.Wallet_Records, error)
}

type Instance func() WalletFacade

var instances = make(map[string]Instance)

func Register(name string, inst Instance) {
	if inst == nil {
		panic("register instance is nil")
	}
	if _, ok := instances[name]; ok {
		panic("register called twice for instance " + name)
	}
	instances[name] = inst
}

func NewWalletFacade(name string) (inst WalletFacade, err error) {
	instanceFunc, ok := instances[name]
	if !ok {
		err = fmt.Errorf("unknown adapter name %q (forgot to import?)", name)
		return
	}
	inst = instanceFunc()
	return
}
