package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/private/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type DialogFiltersDAO struct {
	db *sqlx.DB
}

func NewDialogFiltersDAO(db *sqlx.DB) *DialogFiltersDAO {
	return &DialogFiltersDAO{db}
}

func (dao *DialogFiltersDAO) InsertOrUpdate(ctx context.Context, do *dataobject.DialogFiltersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_filters(user_id, dialog_filter_id, dialog_filter, order_value) values (:user_id, :dialog_filter_id, :dialog_filter, :order_value) on duplicate key update dialog_filter = values(dialog_filter), order_value = values(order_value), deleted = 0"
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

func (dao *DialogFiltersDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.DialogFiltersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into dialog_filters(user_id, dialog_filter_id, dialog_filter, order_value) values (:user_id, :dialog_filter_id, :dialog_filter, :order_value) on duplicate key update dialog_filter = values(dialog_filter), order_value = values(order_value), deleted = 0"
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

func (dao *DialogFiltersDAO) SelectList(ctx context.Context, user_id int32) (rList []dataobject.DialogFiltersDO, err error) {
	var (
		query = "select user_id, dialog_filter_id, dialog_filter from dialog_filters where user_id = ? and deleted = 0 order by order_value desc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, user_id)

	if err != nil {
		log.Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.DialogFiltersDO
	for rows.Next() {
		v := dataobject.DialogFiltersDO{}
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

func (dao *DialogFiltersDAO) UpdateOrder(ctx context.Context, order_value int64, user_id int32, dialog_filter_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_filters set order_value = ? where user_id = ? and dialog_filter_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, order_value, user_id, dialog_filter_id)

	if err != nil {
		log.Errorf("exec in UpdateOrder(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateOrder(_), error: %v", err)
	}

	return
}

func (dao *DialogFiltersDAO) UpdateOrderTx(tx *sqlx.Tx, order_value int64, user_id int32, dialog_filter_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_filters set order_value = ? where user_id = ? and dialog_filter_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, order_value, user_id, dialog_filter_id)

	if err != nil {
		log.Errorf("exec in UpdateOrder(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateOrder(_), error: %v", err)
	}

	return
}

func (dao *DialogFiltersDAO) Clear(ctx context.Context, user_id int32, dialog_filter_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_filters set deleted = 1, dialog_filter = '', order_value = 0 where user_id = ? and dialog_filter_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id, dialog_filter_id)

	if err != nil {
		log.Errorf("exec in Clear(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Clear(_), error: %v", err)
	}

	return
}

func (dao *DialogFiltersDAO) ClearTx(tx *sqlx.Tx, user_id int32, dialog_filter_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update dialog_filters set deleted = 1, dialog_filter = '', order_value = 0 where user_id = ? and dialog_filter_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id, dialog_filter_id)

	if err != nil {
		log.Errorf("exec in Clear(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Clear(_), error: %v", err)
	}

	return
}
