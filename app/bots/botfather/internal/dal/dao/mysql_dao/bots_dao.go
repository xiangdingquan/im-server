package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/bots/botfather/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type BotsDAO struct {
	db *sqlx.DB
}

func NewBotsDAO(db *sqlx.DB) *BotsDAO {
	return &BotsDAO{db}
}

func (dao *BotsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.BotsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into bots(bot_id, bot_type, creator_user_id, token, description) values (:bot_id, :bot_type, :creator_user_id, :token, :description)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		log.Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", do, err)
	}

	return
}

func (dao *BotsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.BotsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into bots(bot_id, bot_type, creator_user_id, token, description) values (:bot_id, :bot_type, :creator_user_id, :token, :description)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		log.Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", do, err)
	}

	return
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

func (dao *BotsDAO) UpdateDescription(ctx context.Context, description string, bot_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update bots set description = ? where bot_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, description, bot_id)

	if err != nil {
		log.Errorf("exec in UpdateDescription(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateDescription(_), error: %v", err)
	}

	return
}

func (dao *BotsDAO) UpdateDescriptionTx(tx *sqlx.Tx, description string, bot_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update bots set description = ? where bot_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, description, bot_id)

	if err != nil {
		log.Errorf("exec in UpdateDescription(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateDescription(_), error: %v", err)
	}

	return
}

func (dao *BotsDAO) UpdateToken(ctx context.Context, token string, bot_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update bots set token = ? where bot_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, token, bot_id)

	if err != nil {
		log.Errorf("exec in UpdateToken(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateToken(_), error: %v", err)
	}

	return
}

func (dao *BotsDAO) UpdateTokenTx(tx *sqlx.Tx, token string, bot_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update bots set token = ? where bot_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, token, bot_id)

	if err != nil {
		log.Errorf("exec in UpdateToken(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateToken(_), error: %v", err)
	}

	return
}
