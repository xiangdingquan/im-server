package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type BotCommandsDAO struct {
	db *sqlx.DB
}

func NewBotCommandsDAO(db *sqlx.DB) *BotCommandsDAO {
	return &BotCommandsDAO{db}
}

func (dao *BotCommandsDAO) InsertBulk(ctx context.Context, doList []*dataobject.BotCommandsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, doList)
	if err != nil {
		log.Errorf("namedExec in InsertBulk(%v), error: %v", doList, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertBulk(%v)_error: %v", doList, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertBulk(%v)_error: %v", doList, err)
	}

	return
}

func (dao *BotCommandsDAO) InsertBulkTx(tx *sqlx.Tx, doList []*dataobject.BotCommandsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, doList)
	if err != nil {
		log.Errorf("namedExec in InsertBulk(%v), error: %v", doList, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertBulk(%v)_error: %v", doList, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertBulk(%v)_error: %v", doList, err)
	}

	return
}

func (dao *BotCommandsDAO) Delete(ctx context.Context, bot_id int32) (rowsAffected int64, err error) {
	var (
		query   = "delete from bot_commands where bot_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, bot_id)

	if err != nil {
		log.Errorf("exec in Delete(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Delete(_), error: %v", err)
	}

	return
}

func (dao *BotCommandsDAO) DeleteTx(tx *sqlx.Tx, bot_id int32) (rowsAffected int64, err error) {
	var (
		query   = "delete from bot_commands where bot_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, bot_id)

	if err != nil {
		log.Errorf("exec in Delete(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Delete(_), error: %v", err)
	}

	return
}

func (dao *BotCommandsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.BotCommandsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)"
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

func (dao *BotCommandsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.BotCommandsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into bot_commands(bot_id, command, description) values (:bot_id, :command, :description)"
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

func (dao *BotCommandsDAO) SelectList(ctx context.Context, bot_id int32) (rList []dataobject.BotCommandsDO, err error) {
	var (
		query = "select id, bot_id, command, description from bot_commands where bot_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, bot_id)

	if err != nil {
		log.Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.BotCommandsDO
	for rows.Next() {
		v := dataobject.BotCommandsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *BotCommandsDAO) SelectListByIdList(ctx context.Context, id_list []int32) (rList []dataobject.BotCommandsDO, err error) {
	var (
		query = "select id, bot_id, command, description from bot_commands where bot_id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, id_list)
	if err != nil {
		log.Errorf("sqlx.In in SelectListByIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectListByIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.BotCommandsDO
	for rows.Next() {
		v := dataobject.BotCommandsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectListByIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
