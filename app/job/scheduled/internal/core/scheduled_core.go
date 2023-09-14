package core

import (
	"context"

	"open.chat/app/job/scheduled/internal/dao"
	"open.chat/pkg/log"
)

type ScheduledCore struct {
	*dao.Dao
}

func New(d *dao.Dao) *ScheduledCore {
	if d == nil {
		d = dao.New()
	}
	return &ScheduledCore{d}
}

func (m *ScheduledCore) GetUserLangCode(ctx context.Context, user_id int32) (rValue string) {
	var (
		query = "select auth_key_id from auth_users where user_id = ? and deleted = 0 order by updated_at desc limit 1"
	)
	rValue = "en"
	var auth_key_id int64 = 0
	err := m.Get(ctx, &auth_key_id, query, user_id)

	if err != nil {
		log.Errorf("get in SelectAuthKey(_), error: %v", err)
		return
	}

	if err == nil && auth_key_id > 0 {
		var query = "select system_lang_code from auths where auth_key_id = ? limit 1"
		err = m.Get(ctx, &rValue, query, auth_key_id)
		if err != nil {
			log.Errorf("get in SelectAuthKey(_), error: %v", err)
			return
		}
	}
	return
}
