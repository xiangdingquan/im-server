package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/auth_session/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type AuthsDAO struct {
	db *sqlx.DB
}

func NewAuthsDAO(db *sqlx.DB) *AuthsDAO {
	return &AuthsDAO{db}
}

func (dao *AuthsDAO) InsertOrUpdate(ctx context.Context, do *dataobject.AuthsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auths(auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, proxy, params, client_ip) values (:auth_key_id, :layer, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :proxy, :params, :client_ip) on duplicate key update layer = values(layer), system_version = values(system_version), app_version = values(app_version), system_lang_code = values(system_lang_code), lang_pack = values(lang_pack), lang_code = values(lang_code), proxy = values(proxy), params = values(params), client_ip = values(client_ip)"
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

func (dao *AuthsDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.AuthsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into auths(auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, proxy, params, client_ip) values (:auth_key_id, :layer, :api_id, :device_model, :system_version, :app_version, :system_lang_code, :lang_pack, :lang_code, :proxy, :params, :client_ip) on duplicate key update layer = values(layer), system_version = values(system_version), app_version = values(app_version), system_lang_code = values(system_lang_code), lang_pack = values(lang_pack), lang_code = values(lang_code), proxy = values(proxy), params = values(params), client_ip = values(client_ip)"
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

func (dao *AuthsDAO) SelectSessions(ctx context.Context, idList []int64) (rList []dataobject.AuthsDO, err error) {
	var (
		query = "select auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, client_ip from auths where auth_key_id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, idList)
	if err != nil {
		log.Errorf("sqlx.In in SelectSessions(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectSessions(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.AuthsDO
	for rows.Next() {
		v := dataobject.AuthsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectSessions(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *AuthsDAO) SelectByAuthKeyId(ctx context.Context, auth_key_id int64) (rValue *dataobject.AuthsDO, err error) {
	var (
		query = "select auth_key_id, layer, api_id, device_model, system_version, app_version, system_lang_code, lang_pack, lang_code, client_ip from auths where auth_key_id = ? and deleted = 0 limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, auth_key_id)

	if err != nil {
		log.Errorf("queryx in SelectByAuthKeyId(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.AuthsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByAuthKeyId(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *AuthsDAO) SelectLayer(ctx context.Context, auth_key_id int64) (rValue int32, err error) {
	var query = "select layer from auths where auth_key_id = ? limit 1"
	err = dao.db.Get(ctx, &rValue, query, auth_key_id)

	if err != nil {
		log.Errorf("get in SelectLayer(_), error: %v", err)
	}

	return
}

func (dao *AuthsDAO) SelectLangCode(ctx context.Context, auth_key_id int64) (langCode string, systemLangCode string, err error) {
	var query = "select lang_code, system_lang_code from auths where auth_key_id = ? limit 1"
	var rValue = struct {
		LangCode       string `db:"lang_code"`
		SystemLangCode string `db:"system_lang_code"`
	}{}
	err = dao.db.Get(ctx, &rValue, query, auth_key_id)

	if err != nil {
		log.Errorf("get in SelectLangCode(_), error: %v", err)
	}
	langCode = rValue.LangCode
	systemLangCode = rValue.SystemLangCode
	return
}
