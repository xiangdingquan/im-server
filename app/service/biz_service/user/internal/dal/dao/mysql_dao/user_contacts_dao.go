package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type UserContactsDAO struct {
	db *sqlx.DB
}

func NewUserContactsDAO(db *sqlx.DB) *UserContactsDAO {
	return &UserContactsDAO{db}
}

func (dao *UserContactsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.UserContactsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_contacts(owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, date2) values (:owner_user_id, :contact_user_id, :contact_phone, :contact_first_name, :contact_last_name, :mutual, :date2) on duplicate key update contact_phone = values(contact_phone), contact_first_name = values(contact_first_name), contact_last_name = values(contact_last_name), mutual = values(mutual), date2 = values(date2), is_deleted = 0"
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

func (dao *UserContactsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.UserContactsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into user_contacts(owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, date2) values (:owner_user_id, :contact_user_id, :contact_phone, :contact_first_name, :contact_last_name, :mutual, :date2) on duplicate key update contact_phone = values(contact_phone), contact_first_name = values(contact_first_name), contact_last_name = values(contact_last_name), mutual = values(mutual), date2 = values(date2), is_deleted = 0"
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

func (dao *UserContactsDAO) SelectContact(ctx context.Context, owner_user_id int32, contact_user_id int32) (rValue *dataobject.UserContactsDO, err error) {
	var (
		query = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where is_deleted = 0 and owner_user_id = ? and contact_user_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, owner_user_id, contact_user_id)

	if err != nil {
		log.Errorf("queryx in SelectContact(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UserContactsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectContact(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *UserContactsDAO) SelectByContactId(ctx context.Context, owner_user_id int32, contact_user_id int32) (rValue *dataobject.UserContactsDO, err error) {
	var (
		query = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where owner_user_id = ? and contact_user_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, owner_user_id, contact_user_id)

	if err != nil {
		log.Errorf("queryx in SelectByContactId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UserContactsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByContactId(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *UserContactsDAO) SelectListByPhoneList(ctx context.Context, owner_user_id int32, phoneList []string) (rList []dataobject.UserContactsDO, err error) {
	var (
		query = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where is_deleted = 0 and owner_user_id = ? and contact_phone in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, owner_user_id, phoneList)
	if err != nil {
		log.Errorf("sqlx.In in SelectListByPhoneList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectListByPhoneList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.UserContactsDO
	for rows.Next() {
		v := dataobject.UserContactsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectListByPhoneList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *UserContactsDAO) SelectAllUserContacts(ctx context.Context, owner_user_id int32) (rList []dataobject.UserContactsDO, err error) {
	var (
		query = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where owner_user_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, owner_user_id)

	if err != nil {
		log.Errorf("queryx in SelectAllUserContacts(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.UserContactsDO
	for rows.Next() {
		v := dataobject.UserContactsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectAllUserContacts(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *UserContactsDAO) SelectUserContacts(ctx context.Context, owner_user_id int32) (rList []dataobject.UserContactsDO, err error) {
	var (
		query = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where owner_user_id = ? and is_deleted = 0 order by contact_user_id asc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, owner_user_id)

	if err != nil {
		log.Errorf("queryx in SelectUserContacts(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.UserContactsDO
	for rows.Next() {
		v := dataobject.UserContactsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectUserContacts(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *UserContactsDAO) SelectListByIdList(ctx context.Context, owner_user_id int32, id_list []int32) (rList []dataobject.UserContactsDO, err error) {
	var (
		query = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where owner_user_id = ? and contact_user_id in (?) and is_deleted = 0"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, owner_user_id, id_list)
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

	var values []dataobject.UserContactsDO
	for rows.Next() {
		v := dataobject.UserContactsDO{}
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

func (dao *UserContactsDAO) SelectListByOwnerListAndContactList(ctx context.Context, idList1 []int32, idList2 []int32) (rList []dataobject.UserContactsDO, err error) {
	var (
		query = "select id, owner_user_id, contact_user_id, contact_phone, contact_first_name, contact_last_name, mutual, is_deleted from user_contacts where owner_user_id in (?) and contact_user_id in (?) and is_deleted = 0"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, idList1, idList2)
	if err != nil {
		log.Errorf("sqlx.In in SelectListByOwnerListAndContactList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectListByOwnerListAndContactList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.UserContactsDO
	for rows.Next() {
		v := dataobject.UserContactsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectListByOwnerListAndContactList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *UserContactsDAO) UpdateContactNameById(ctx context.Context, contact_first_name string, contact_last_name string, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set contact_first_name = ?, contact_last_name = ?, is_deleted = 0 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, contact_first_name, contact_last_name, id)

	if err != nil {
		log.Errorf("exec in UpdateContactNameById(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateContactNameById(_), error: %v", err)
	}

	return
}

func (dao *UserContactsDAO) UpdateContactNameByIdTx(tx *sqlx.Tx, contact_first_name string, contact_last_name string, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set contact_first_name = ?, contact_last_name = ?, is_deleted = 0 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, contact_first_name, contact_last_name, id)

	if err != nil {
		log.Errorf("exec in UpdateContactNameById(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateContactNameById(_), error: %v", err)
	}

	return
}

func (dao *UserContactsDAO) UpdateContactName(ctx context.Context, contact_first_name string, contact_last_name string, owner_user_id int32, contact_user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set contact_first_name = ?, contact_last_name = ?, is_deleted = 0 where contact_user_id != 0 and (owner_user_id = ? and contact_user_id = ?)"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, contact_first_name, contact_last_name, owner_user_id, contact_user_id)

	if err != nil {
		log.Errorf("exec in UpdateContactName(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateContactName(_), error: %v", err)
	}

	return
}

func (dao *UserContactsDAO) UpdateContactNameTx(tx *sqlx.Tx, contact_first_name string, contact_last_name string, owner_user_id int32, contact_user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set contact_first_name = ?, contact_last_name = ?, is_deleted = 0 where contact_user_id != 0 and (owner_user_id = ? and contact_user_id = ?)"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, contact_first_name, contact_last_name, owner_user_id, contact_user_id)

	if err != nil {
		log.Errorf("exec in UpdateContactName(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateContactName(_), error: %v", err)
	}

	return
}

func (dao *UserContactsDAO) UpdateMutual(ctx context.Context, mutual int8, owner_user_id int32, contact_user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set mutual = ? where contact_user_id != 0 and (owner_user_id = ? and contact_user_id = ?)"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, mutual, owner_user_id, contact_user_id)

	if err != nil {
		log.Errorf("exec in UpdateMutual(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateMutual(_), error: %v", err)
	}

	return
}

func (dao *UserContactsDAO) UpdateMutualTx(tx *sqlx.Tx, mutual int8, owner_user_id int32, contact_user_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set mutual = ? where contact_user_id != 0 and (owner_user_id = ? and contact_user_id = ?)"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, mutual, owner_user_id, contact_user_id)

	if err != nil {
		log.Errorf("exec in UpdateMutual(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateMutual(_), error: %v", err)
	}

	return
}

func (dao *UserContactsDAO) DeleteContacts(ctx context.Context, owner_user_id int32, id_list []int32) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set is_deleted = 1, mutual = 0 where contact_user_id != 0 and (owner_user_id = ? and contact_user_id in (?))"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, owner_user_id, id_list)
	if err != nil {
		log.Errorf("sqlx.In in DeleteContacts(_), error: %v", err)
		return
	}
	rResult, err = dao.db.Exec(ctx, query, a...)

	if err != nil {
		log.Errorf("exec in DeleteContacts(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DeleteContacts(_), error: %v", err)
	}

	return
}

func (dao *UserContactsDAO) DeleteContactsTx(tx *sqlx.Tx, owner_user_id int32, id_list []int32) (rowsAffected int64, err error) {
	var (
		query   = "update user_contacts set is_deleted = 1, mutual = 0 where contact_user_id != 0 and (owner_user_id = ? and contact_user_id in (?))"
		a       []interface{}
		rResult sql.Result
	)

	query, a, err = sqlx.In(query, owner_user_id, id_list)
	if err != nil {
		log.Errorf("sqlx.In in DeleteContacts(_), error: %v", err)
		return
	}
	rResult, err = tx.Exec(query, a...)

	if err != nil {
		log.Errorf("exec in DeleteContacts(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DeleteContacts(_), error: %v", err)
	}

	return
}
