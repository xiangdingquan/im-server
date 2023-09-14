package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/messenger/sync/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type BotUpdatesDAO struct {
	db *sqlx.DB
}

func NewBotUpdatesDAO(db *sqlx.DB) *BotUpdatesDAO {
	return &BotUpdatesDAO{db}
}

func (dao *BotUpdatesDAO) Insert(ctx context.Context, do *dataobject.BotUpdatesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into bot_updates(bot_id, update_id, update_type, update_data, date2) values (:bot_id, :update_id, :update_type, :update_data, :date2)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		log.Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

func (dao *BotUpdatesDAO) InsertTx(tx *sqlx.Tx, do *dataobject.BotUpdatesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into bot_updates(bot_id, update_id, update_type, update_data, date2) values (:bot_id, :update_id, :update_type, :update_data, :date2)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		log.Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

func (dao *BotUpdatesDAO) SelectByLastUpdateId(ctx context.Context, bot_id int32) (rValue *dataobject.BotUpdatesDO, err error) {
	var (
		query = "select bot_id, update_id, update_type, update_data, date2 from bot_updates where bot_id = ? order by update_id desc limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, bot_id)

	if err != nil {
		log.Errorf("queryx in SelectByLastUpdateId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.BotUpdatesDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByLastUpdateId(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *BotUpdatesDAO) SelectByGtUpdateId(ctx context.Context, bot_id int32, update_id int32, limit int32) (rList []dataobject.BotUpdatesDO, err error) {
	var (
		query = "select bot_id, update_id, update_type, update_data, date2 from bot_updates where bot_id = ? and update_id > ? order by update_id asc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, bot_id, update_id, limit)

	if err != nil {
		log.Errorf("queryx in SelectByGtUpdateId(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.BotUpdatesDO
	for rows.Next() {
		v := dataobject.BotUpdatesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByGtUpdateId(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
