package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type UserBlocksDAO struct {
	db *sqlx.DB
}

func NewUserBlocksDAO(db *sqlx.DB) *UserBlocksDAO {
	return &UserBlocksDAO{db}
}

func (dao *UserBlocksDAO) InsertOrUpdate(ctx context.Context, do *dataobject.UserBlocksDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_blocks(user_id, block_id, `date`) values (:user_id, :block_id, :date) on duplicate key update `date` = values(`date`), deleted = 0"
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

func (dao *UserBlocksDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.UserBlocksDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_blocks(user_id, block_id, `date`) values (:user_id, :block_id, :date) on duplicate key update `date` = values(`date`), deleted = 0"
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

func (dao *UserBlocksDAO) SelectList(ctx context.Context, user_id int32, limit int32) (rList []dataobject.UserBlocksDO, err error) {
	var (
		query = "select user_id, block_id, `date` from user_blocks where user_id = ? and deleted = 0 order by id asc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, limit)

	if err != nil {
		log.Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.UserBlocksDO
	for rows.Next() {
		v := dataobject.UserBlocksDO{}
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

func (dao *UserBlocksDAO) SelectListByIdList(ctx context.Context, user_id int32, idList []int32) (rList []int32, err error) {
	var (
		query = "select block_id from user_blocks where user_id = ? and block_id in (?) and deleted = 0"
		a     []interface{}
	)
	query, a, err = sqlx.In(query, user_id, idList)
	if err != nil {
		log.Errorf("sqlx.In in SelectListByIdList(_), error: %v", err)
		return
	}

	err = dao.db.Select(ctx, &rList, query, a...)

	if err != nil {
		log.Errorf("select in SelectListByIdList(_), error: %v", err)
	}

	return
}

func (dao *UserBlocksDAO) Select(ctx context.Context, user_id int32, block_id int32) (rValue *dataobject.UserBlocksDO, err error) {
	var (
		query = "select user_id, block_id, `date` from user_blocks where user_id = ? and block_id = ? and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id, block_id)

	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UserBlocksDO{}
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

func (dao *UserBlocksDAO) Delete(ctx context.Context, user_id int32, block_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update user_blocks set deleted = 1, `date` = 0 where user_id = ? and block_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id, block_id)

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

func (dao *UserBlocksDAO) DeleteTx(tx *sqlx.Tx, user_id int32, block_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update user_blocks set deleted = 1, `date` = 0 where user_id = ? and block_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id, block_id)

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
