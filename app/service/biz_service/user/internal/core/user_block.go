package core

import (
	"context"
	"time"

	"math"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/mtproto"
)

func (m *UserCore) IsBlockedByUser(ctx context.Context, selfUserId, id int32) bool {
	do, _ := m.UserBlocksDAO.Select(ctx, selfUserId, id)
	return do != nil
}

func (m *UserCore) BlockUser(ctx context.Context, userId, blockId int32) bool {
	var (
		date = time.Now().Unix()
		do   = &dataobject.UserBlocksDO{
			UserId:  userId,
			BlockId: blockId,
			Date:    int32(date),
		}
		blockedList = map[int32]int64{blockId: date}
	)
	if _, _, err := m.UserBlocksDAO.InsertOrUpdate(ctx, do); err != nil {
		return false
	}

	doList, err := m.UserBlocksDAO.SelectList(ctx, userId, math.MaxInt32)
	if err != nil {
		return false
	}

	for i := 0; i < len(doList); i++ {
		blockedList[doList[i].BlockId] = int64(doList[i].Date)
	}
	delete(blockedList, blockId)
	return true
}

func (m *UserCore) UnBlockUser(ctx context.Context, userId, blockId int32) bool {
	_, err := m.UserBlocksDAO.Delete(ctx, userId, blockId)
	if err != nil {
		return false
	}
	return true
}

func (m *UserCore) GetBlockedList(ctx context.Context, userId, offset, limit int32) []*mtproto.ContactBlocked {
	doList, _ := m.UserBlocksDAO.SelectList(ctx, userId, limit)
	bockedList := make([]*mtproto.ContactBlocked, 0, len(doList))
	for _, do := range doList {
		blocked := mtproto.MakeTLContactBlocked(&mtproto.ContactBlocked{
			UserId: do.BlockId,
			Date:   do.Date,
		})
		bockedList = append(bockedList, blocked.To_ContactBlocked())
	}
	return bockedList
}

func (m *UserCore) CheckBlockUser(ctx context.Context, selfUserId, id int32) bool {
	blockedList := make(map[int32]int64)

	doList, err := m.UserBlocksDAO.SelectList(ctx, selfUserId, math.MaxInt16)
	if err != nil {
		return false
	}

	for i := 0; i < len(doList); i++ {
		blockedList[doList[i].BlockId] = int64(doList[i].Date)
	}
	if _, ok := blockedList[id]; !ok {
		return false
	}
	return true
}

func (m *UserCore) CheckBlockUserList(ctx context.Context, selfUserId int32, idList []int32) []int32 {
	var blockedList []int32
	if len(idList) > 0 {
		blockedList, _ = m.UserBlocksDAO.SelectListByIdList(ctx, selfUserId, idList)
	}

	if blockedList == nil {
		blockedList = []int32{}
	}

	return blockedList
}
