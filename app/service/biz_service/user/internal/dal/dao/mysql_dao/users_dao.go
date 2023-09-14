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

type UsersDAO struct {
	db *sqlx.DB
}

func NewUsersDAO(db *sqlx.DB) *UsersDAO {
	return &UsersDAO{db}
}

func (dao *UsersDAO) Insert(ctx context.Context, do *dataobject.UsersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into users(user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, about, is_bot, is_virtual) values (:user_type, :access_hash, :secret_key_id, :first_name, :last_name, :username, :phone, :country_code, :verified, :about, :is_bot, :is_virtual)"
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

func (dao *UsersDAO) InsertTx(tx *sqlx.Tx, do *dataobject.UsersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into users(user_type, access_hash, secret_key_id, first_name, last_name, username, phone, country_code, verified, about, is_bot, is_virtual) values (:user_type, :access_hash, :secret_key_id, :first_name, :last_name, :username, :phone, :country_code, :verified, :about, :is_bot, :is_virtual)"
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

func (dao *UsersDAO) SelectByPhoneNumber(ctx context.Context, phone string) (rValue *dataobject.UsersDO, err error) {
	var (
		query = "select id, user_type, access_hash, secret_key_id, first_name, channel_id, inviter_uid, last_name, username, password, phone, photos, country_code, verified, about, is_bot, is_internal, is_virtual, deleted from users where phone = ? limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, phone)

	if err != nil {
		log.Errorf("queryx in SelectByPhoneNumber(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UsersDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByPhoneNumber(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *UsersDAO) SelectById(ctx context.Context, id int32) (rValue *dataobject.UsersDO, err error) {
	var (
		query = "select id, user_type, access_hash, secret_key_id, channel_id, inviter_uid, first_name, last_name, username, password, phone, photos, country_code, verified, about, is_bot, is_internal, is_virtual, deleted from users where id = ? limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, id)

	if err != nil {
		log.Errorf("queryx in SelectById(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UsersDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectById(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *UsersDAO) SelectUsersByIdList(ctx context.Context, id_list []int32) (rList []dataobject.UsersDO, err error) {
	var (
		query = "select id, user_type, access_hash, secret_key_id, channel_id, inviter_uid, first_name, last_name, username, password, phone, photos, country_code, country, province, city, city_code, gender, birth, verified, about, is_bot, is_internal, is_virtual, deleted from users where id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, id_list)
	if err != nil {
		log.Errorf("sqlx.In in SelectUsersByIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectUsersByIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.UsersDO
	for rows.Next() {
		v := dataobject.UsersDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectUsersByIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *UsersDAO) SelectUsersByPhoneList(ctx context.Context, phoneList []string) (rList []dataobject.UsersDO, err error) {
	var (
		query = "select id, user_type, access_hash, secret_key_id, channel_id, inviter_uid, first_name, last_name, username, password, phone, photos, country_code, verified, about, is_bot, is_internal, is_virtual, deleted from users where phone in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, phoneList)
	if err != nil {
		log.Errorf("sqlx.In in SelectUsersByPhoneList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectUsersByPhoneList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.UsersDO
	for rows.Next() {
		v := dataobject.UsersDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectUsersByPhoneList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *UsersDAO) SelectByQueryString(ctx context.Context, username string, first_name string, last_name string, phone string) (rList []dataobject.UsersDO, err error) {
	var (
		query = "select id, user_type, access_hash, secret_key_id, channel_id, inviter_uid, first_name, last_name, username, password, phone, photos, country_code, verified, about, is_bot, is_internal, is_virtual, deleted from users where username = ? or first_name = ? or last_name = ? or phone = ? limit 20"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, username, first_name, last_name, phone)

	if err != nil {
		log.Errorf("queryx in SelectByQueryString(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.UsersDO
	for rows.Next() {
		v := dataobject.UsersDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByQueryString(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *UsersDAO) SearchByQueryNotIdList(ctx context.Context, q2 string, id_list []int32, limit int32) (rList []dataobject.UsersDO, err error) {
	var (
		query = "select id from users where username like ? and id not in (?) limit ?"
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

	var values []dataobject.UsersDO
	for rows.Next() {
		v := dataobject.UsersDO{}
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

func (dao *UsersDAO) Delete(ctx context.Context, delete_reason string, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update users set deleted = 1, delete_reason = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, delete_reason, id)

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

func (dao *UsersDAO) DeleteTx(tx *sqlx.Tx, delete_reason string, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update users set deleted = 1, delete_reason = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, delete_reason, id)

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

func (dao *UsersDAO) UpdateUsername(ctx context.Context, username string, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update users set username = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, username, id)

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

func (dao *UsersDAO) UpdateUsernameTx(tx *sqlx.Tx, username string, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update users set username = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, username, id)

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

func (dao *UsersDAO) UpdateFirstAndLastName(ctx context.Context, first_name string, last_name string, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update users set first_name = ?, last_name = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, first_name, last_name, id)

	if err != nil {
		log.Errorf("exec in UpdateFirstAndLastName(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateFirstAndLastName(_), error: %v", err)
	}

	return
}

func (dao *UsersDAO) UpdateFirstAndLastNameTx(tx *sqlx.Tx, first_name string, last_name string, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update users set first_name = ?, last_name = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, first_name, last_name, id)

	if err != nil {
		log.Errorf("exec in UpdateFirstAndLastName(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateFirstAndLastName(_), error: %v", err)
	}

	return
}

func (dao *UsersDAO) UpdateAbout(ctx context.Context, about string, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update users set about = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, about, id)

	if err != nil {
		log.Errorf("exec in UpdateAbout(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateAbout(_), error: %v", err)
	}

	return
}

func (dao *UsersDAO) UpdateAboutTx(tx *sqlx.Tx, about string, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update users set about = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, about, id)

	if err != nil {
		log.Errorf("exec in UpdateAbout(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateAbout(_), error: %v", err)
	}

	return
}

func (dao *UsersDAO) UpdateProfile(ctx context.Context, first_name string, last_name string, about string, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update users set first_name = ?, last_name = ?, about = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, first_name, last_name, about, id)

	if err != nil {
		log.Errorf("exec in UpdateProfile(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateProfile(_), error: %v", err)
	}

	return
}

func (dao *UsersDAO) UpdateProfileTx(tx *sqlx.Tx, first_name string, last_name string, about string, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update users set first_name = ?, last_name = ?, about = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, first_name, last_name, about, id)

	if err != nil {
		log.Errorf("exec in UpdateProfile(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateProfile(_), error: %v", err)
	}

	return
}

func (dao *UsersDAO) SelectByUsername(ctx context.Context, username string) (rValue *dataobject.UsersDO, err error) {
	var (
		query = "select id from users where username = ? limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, username)

	if err != nil {
		log.Errorf("queryx in SelectByUsername(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UsersDO{}
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

func (dao *UsersDAO) SelectAccountDaysTTL(ctx context.Context, id int32) (rValue *dataobject.UsersDO, err error) {
	var (
		query = "select account_days_ttl from users where id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, id)

	if err != nil {
		log.Errorf("queryx in SelectAccountDaysTTL(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UsersDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectAccountDaysTTL(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *UsersDAO) UpdateAccountDaysTTL(ctx context.Context, account_days_ttl int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update users set account_days_ttl = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, account_days_ttl, id)

	if err != nil {
		log.Errorf("exec in UpdateAccountDaysTTL(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateAccountDaysTTL(_), error: %v", err)
	}

	return
}

func (dao *UsersDAO) UpdateAccountDaysTTLTx(tx *sqlx.Tx, account_days_ttl int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update users set account_days_ttl = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, account_days_ttl, id)

	if err != nil {
		log.Errorf("exec in UpdateAccountDaysTTL(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateAccountDaysTTL(_), error: %v", err)
	}

	return
}

func (dao *UsersDAO) SelectProfilePhotos(ctx context.Context, id int32) (rValue *dataobject.UsersDO, err error) {
	var (
		query = "select photos from users where id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, id)

	if err != nil {
		log.Errorf("queryx in SelectProfilePhotos(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UsersDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectProfilePhotos(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *UsersDAO) SelectCountryCode(ctx context.Context, id int32) (rValue *dataobject.UsersDO, err error) {
	var (
		query = "select country_code from users where id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, id)

	if err != nil {
		log.Errorf("queryx in SelectCountryCode(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UsersDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectCountryCode(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *UsersDAO) UpdateProfilePhotos(ctx context.Context, photos string, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update users set photos = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, photos, id)

	if err != nil {
		log.Errorf("exec in UpdateProfilePhotos(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateProfilePhotos(_), error: %v", err)
	}

	return
}

func (dao *UsersDAO) UpdateProfilePhotosTx(tx *sqlx.Tx, photos string, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update users set photos = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, photos, id)

	if err != nil {
		log.Errorf("exec in UpdateProfilePhotos(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateProfilePhotos(_), error: %v", err)
	}

	return
}

func (dao *UsersDAO) UpdateUser(ctx context.Context, cMap map[string]interface{}, id int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update users set %s where id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, id)

	rResult, err = dao.db.Exec(ctx, query, aValues...)

	if err != nil {
		log.Errorf("exec in UpdateUser(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateUser(_), error: %v", err)
	}

	return
}

func (dao *UsersDAO) UpdateUserTx(tx *sqlx.Tx, cMap map[string]interface{}, id int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update users set %s where id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, id)

	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		log.Errorf("exec in UpdateUser(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateUser(_), error: %v", err)
	}

	return
}

func (dao *UsersDAO) QueryChannelParticipants(ctx context.Context, channelId int32, q1 string, q2 string, q3 string) (rList []dataobject.UsersDO, err error) {
	var (
		query = "select id, user_type, access_hash, secret_key_id, channel_id, inviter_uid, first_name, last_name, username, password, phone, photos, country_code, verified, about, is_bot, is_internal, deleted from users where id in (select user_id from channel_participants where channel_id = ? and state = 0) and (first_name like ? or last_name like ? or username like ?)"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, channelId, q1, q2, q3)

	if err != nil {
		log.Errorf("queryx in QueryChannelParticipants(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.UsersDO
	for rows.Next() {
		v := dataobject.UsersDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in QueryChannelParticipants(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *UsersDAO) SelectCustomerServiceList(ctx context.Context) (rList []dataobject.UsersDO, err error) {
	var (
		query = "select id, user_type, access_hash, secret_key_id, channel_id, inviter_uid, first_name, last_name, username, password, phone, photos, country_code, verified, about, is_bot, is_internal, deleted from users where is_customer_service = 1 and deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query)

	if err != nil {
		log.Errorf("queryx in SelectCustomerServiceList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.UsersDO
	for rows.Next() {
		v := dataobject.UsersDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectCustomerServiceList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
