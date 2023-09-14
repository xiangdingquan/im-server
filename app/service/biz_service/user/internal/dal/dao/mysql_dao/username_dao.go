package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type UsernameDAO struct {
	db *sqlx.DB
}

func NewUsernameDAO(db *sqlx.DB) *UsernameDAO {
	return &UsernameDAO{db}
}

func (dao *UsernameDAO) Insert(ctx context.Context, do *dataobject.UsernameDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into username(peer_type, peer_id, username) values (:peer_type, :peer_id, :username)"
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

func (dao *UsernameDAO) InsertTx(tx *sqlx.Tx, do *dataobject.UsernameDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into username(peer_type, peer_id, username) values (:peer_type, :peer_id, :username)"
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

func (dao *UsernameDAO) SelectList(ctx context.Context, nameList []string) (rList []dataobject.UsernameDO, err error) {
	var (
		query = "select peer_type, peer_id, username from username where username in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, nameList)
	if err != nil {
		log.Errorf("sqlx.In in SelectList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.UsernameDO
	for rows.Next() {
		v := dataobject.UsernameDO{}
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

func (dao *UsernameDAO) SelectByPeer(ctx context.Context, peer_type int8, peer_id int32) (rValue *dataobject.UsernameDO, err error) {
	var (
		query = "select peer_type, peer_id, username from username where peer_type = ? and peer_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, peer_type, peer_id)

	if err != nil {
		log.Errorf("queryx in SelectByPeer(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UsernameDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByPeer(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *UsernameDAO) SelectByUserId(ctx context.Context, peer_id int32) (rValue *dataobject.UsernameDO, err error) {
	var (
		query = "select peer_type, peer_id, username from username where peer_type = 2 and peer_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, peer_id)

	if err != nil {
		log.Errorf("queryx in SelectByUserId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UsernameDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByUserId(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *UsernameDAO) SelectByChannelId(ctx context.Context, peer_id int32) (rValue *dataobject.UsernameDO, err error) {
	var (
		query = "select peer_type, peer_id, username from username where peer_type = 4 and peer_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, peer_id)

	if err != nil {
		log.Errorf("queryx in SelectByChannelId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UsernameDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByChannelId(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *UsernameDAO) SelectByUsername(ctx context.Context, username string) (rValue *dataobject.UsernameDO, err error) {
	var (
		query = "select peer_type, peer_id, username from username where username = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, username)

	if err != nil {
		log.Errorf("queryx in SelectByUsername(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UsernameDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByUsername(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *UsernameDAO) UpdateUsername(ctx context.Context, username string, peer_type int8, peer_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update username set username = ? where peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, username, peer_type, peer_id)

	if err != nil {
		log.Errorf("exec in UpdateUsername(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateUsername(_), error: %v", err)
	}

	return
}

func (dao *UsernameDAO) UpdateUsernameTx(tx *sqlx.Tx, username string, peer_type int8, peer_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update username set username = ? where peer_type = ? and peer_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, username, peer_type, peer_id)

	if err != nil {
		log.Errorf("exec in UpdateUsername(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateUsername(_), error: %v", err)
	}

	return
}

func (dao *UsernameDAO) SearchByQueryNotIdList(ctx context.Context, q2 string, id_list []int32, limit int32) (rList []dataobject.UsernameDO, err error) {
	var (
		query = "select peer_type, peer_id from username where username like ? and peer_id not in (?) limit ?"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, q2, id_list, limit)
	if err != nil {
		log.Errorf("sqlx.In in SearchByQueryNotIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SearchByQueryNotIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.UsernameDO
	for rows.Next() {
		v := dataobject.UsernameDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SearchByQueryNotIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
