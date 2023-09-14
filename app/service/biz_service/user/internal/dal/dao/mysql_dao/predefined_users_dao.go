package mysql_dao

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type PredefinedUsersDAO struct {
	db *sqlx.DB
}

func NewPredefinedUsersDAO(db *sqlx.DB) *PredefinedUsersDAO {
	return &PredefinedUsersDAO{db}
}

func (dao *PredefinedUsersDAO) Insert(ctx context.Context, do *dataobject.PredefinedUsersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into predefined_users(first_name, last_name, username, phone, code, verified) values (:first_name, :last_name, :username, :phone, :code, :verified)"
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

func (dao *PredefinedUsersDAO) InsertTx(tx *sqlx.Tx, do *dataobject.PredefinedUsersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into predefined_users(first_name, last_name, username, phone, code, verified) values (:first_name, :last_name, :username, :phone, :code, :verified)"
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

func (dao *PredefinedUsersDAO) SelectByPhone(ctx context.Context, phone string) (rValue *dataobject.PredefinedUsersDO, err error) {
	var (
		query = "select id, phone, first_name, last_name, username, code, verified, registered_user_id from predefined_users where phone = ? and deleted = 0 limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, phone)

	if err != nil {
		log.Errorf("queryx in SelectByPhone(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.PredefinedUsersDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByPhone(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *PredefinedUsersDAO) SelectPredefinedUsersAll(ctx context.Context) (rList []dataobject.PredefinedUsersDO, err error) {
	var (
		query = "select id, phone, first_name, last_name, username, code, verified, registered_user_id from predefined_users where deleted = 0 order by username asc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query)

	if err != nil {
		log.Errorf("queryx in SelectPredefinedUsersAll(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.PredefinedUsersDO
	for rows.Next() {
		v := dataobject.PredefinedUsersDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectPredefinedUsersAll(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *PredefinedUsersDAO) Delete(ctx context.Context, phone string) (rowsAffected int64, err error) {
	var (
		query   = "update predefined_users set deleted = 0 where phone = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, phone)

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

func (dao *PredefinedUsersDAO) DeleteTx(tx *sqlx.Tx, phone string) (rowsAffected int64, err error) {
	var (
		query   = "update predefined_users set deleted = 0 where phone = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, phone)

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

func (dao *PredefinedUsersDAO) Update(ctx context.Context, cMap map[string]interface{}, phone string) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update predefined_users set %s where phone = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, phone)

	rResult, err = dao.db.Exec(ctx, query, aValues...)

	if err != nil {
		log.Errorf("exec in Update(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Update(_), error: %v", err)
	}

	return
}

func (dao *PredefinedUsersDAO) UpdateTx(tx *sqlx.Tx, cMap map[string]interface{}, phone string) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update predefined_users set %s where phone = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, phone)

	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		log.Errorf("exec in Update(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Update(_), error: %v", err)
	}

	return
}
