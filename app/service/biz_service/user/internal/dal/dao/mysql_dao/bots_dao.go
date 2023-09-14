package mysql_dao

import (
	"context"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type BotsDAO struct {
	db *sqlx.DB
}

func NewBotsDAO(db *sqlx.DB) *BotsDAO {
	return &BotsDAO{db}
}

func (dao *BotsDAO) Select(ctx context.Context, bot_id int32) (rValue *dataobject.BotsDO, err error) {
	var (
		query = "select id, bot_id, bot_type, creator_user_id, token, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder from bots where bot_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, bot_id)

	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.BotsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in Select(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *BotsDAO) SelectByToken(ctx context.Context, token string) (rValue int32, err error) {
	var query = "select bot_id from bots where token = ?"
	err = dao.db.Get(ctx, &rValue, query, token)

	if err != nil {
		log.Errorf("get in SelectByToken(_), error: %v", err)
	}

	return
}

func (dao *BotsDAO) SelectByIdList(ctx context.Context, id_list []int32) (rList []dataobject.BotsDO, err error) {
	var (
		query = "select id, bot_id, bot_type, creator_user_id, token, description, bot_chat_history, bot_nochats, bot_inline_geo, bot_info_version, bot_inline_placeholder from bots where bot_id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, id_list)
	if err != nil {
		log.Errorf("sqlx.In in SelectByIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectByIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.BotsDO
	for rows.Next() {
		v := dataobject.BotsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
