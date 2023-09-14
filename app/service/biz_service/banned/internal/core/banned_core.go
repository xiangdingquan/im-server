package core

import (
	"context"
	"time"

	"open.chat/app/service/biz_service/banned/internal/dal/dataobject"
	"open.chat/app/service/biz_service/banned/internal/dao"
)

type BannedCore struct {
	*dao.Dao
}

func New(dao *dao.Dao) *BannedCore {
	return &BannedCore{dao}
}

func (m *BannedCore) CheckPhoneNumberBanned(ctx context.Context, phoneNumber string) bool {
	do, _ := m.BannedDAO.CheckBannedByPhone(ctx, phoneNumber)
	var ban = false
	if do != nil {
		ban = time.Now().Unix() > do.BannedTime+do.Expires
	}
	return ban
}

func (m *BannedCore) GetBannedByPhoneList(ctx context.Context, phoneList []string) map[string]bool {
	bMap := make(map[string]bool, len(phoneList))
	pList, _ := m.BannedDAO.SelectPhoneList(ctx, phoneList)
	for _, p := range pList {
		bMap[p] = true
	}
	return bMap
}

func (m *BannedCore) Ban(ctx context.Context, phoneNumber string, expires int32, reason string) bool {
	m.BannedDAO.InsertOrUpdate(ctx, &dataobject.BannedDO{
		Phone:        phoneNumber,
		BannedTime:   time.Now().Unix(),
		Expires:      int64(expires),
		BannedReason: reason,
		Log:          "ban",
		State:        1,
	})
	return true
}

func (m *BannedCore) UnBan(ctx context.Context, phoneNumber string) bool {
	m.BannedDAO.Update(ctx, time.Now().Unix(), "unBan", 0, phoneNumber)
	return false
}
