package core

import (
	"context"
	"encoding/json"
	"time"

	"open.chat/app/service/biz_service/private/internal/dal/dataobject"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
)

func (m *PrivateCore) InsertOrUpdateDialogFilter(ctx context.Context, userId, id int32, dialogFilter *mtproto.DialogFilter) error {
	_, _, err := m.DialogFiltersDAO.InsertOrUpdate(ctx, &dataobject.DialogFiltersDO{
		UserId:         userId,
		DialogFilterId: id,
		DialogFilter:   hack.String(model.TLObjectToJson(dialogFilter)),
		OrderValue:     0,
		Deleted:        0,
	})
	return err
}

func (m *PrivateCore) DeleteDialogFilter(ctx context.Context, userId, id int32) error {
	_, err := m.DialogFiltersDAO.Clear(ctx, userId, id)
	return err
}

func (m *PrivateCore) UpdateDialogFiltersOrder(ctx context.Context, userId int32, order []int32) error {
	var (
		err error
		now = time.Now().Unix()
	)
	for _, id := range order {
		if _, err = m.DialogFiltersDAO.UpdateOrder(ctx, now, userId, id); err != nil {
			return err
		}
		now--
	}
	return nil
}

func (m *PrivateCore) GetDialogFilters(ctx context.Context, userId int32) model.DialogFilterExtList {
	doList, _ := m.DialogFiltersDAO.SelectList(ctx, userId)
	dialogFilters := make([]*model.DialogFilterExt, 0, len(doList))
	for i := 0; i < len(doList); i++ {
		dialogFilter := &model.DialogFilterExt{
			Id:           doList[i].DialogFilterId,
			DialogFilter: mtproto.MakeTLDialogFilter(nil).To_DialogFilter(),
			Order:        doList[i].OrderValue,
		}
		if err := json.Unmarshal(hack.Bytes(doList[i].DialogFilter), dialogFilter.DialogFilter); err != nil {
			log.Errorf("json.Unmarshal(%v) - error: %v", doList[i], err)
			continue
		}
		dialogFilters = append(dialogFilters, dialogFilter)
	}
	return dialogFilters
}
