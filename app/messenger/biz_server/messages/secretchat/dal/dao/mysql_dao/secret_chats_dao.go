package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/messenger/biz_server/messages/secretchat/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type SecretChatsDAO struct {
	db *sqlx.DB
}

func NewSecretChatsDAO(db *sqlx.DB) *SecretChatsDAO {
	return &SecretChatsDAO{db}
}

func (dao *SecretChatsDAO) Insert(ctx context.Context, do *dataobject.SecretChatsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into secret_chats(id, access_hash, admin_id, participant_id, admin_auth_key_id, random_id, g_a, state, `date`) values (:random_id, :access_hash, :admin_id, :participant_id, :admin_auth_key_id, :random_id, :g_a, :state, :date)"
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

func (dao *SecretChatsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.SecretChatsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into secret_chats(access_hash, admin_id, participant_id, admin_auth_key_id, random_id, g_a, state, `date`) values (:access_hash, :admin_id, :participant_id, :admin_auth_key_id, :random_id, :g_a, :state, :date)"
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

func (dao *SecretChatsDAO) Select(ctx context.Context, id int32) (rValue *dataobject.SecretChatsDO, err error) {
	var (
		query = "select id, access_hash, admin_id, participant_id, admin_auth_key_id, participant_auth_key_id, random_id, g_a, g_b, key_fingerprint, state, `date` from secret_chats where id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, id)

	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.SecretChatsDO{}
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

func (dao *SecretChatsDAO) UpdateGB(ctx context.Context, participant_auth_key_id int64, g_b string, key_fingerprint int64, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update secret_chats set participant_auth_key_id = ?, g_b = ?, key_fingerprint = ?, state = 2 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, participant_auth_key_id, g_b, key_fingerprint, id)

	if err != nil {
		log.Errorf("exec in UpdateGB(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateGB(_), error: %v", err)
	}

	return
}

func (dao *SecretChatsDAO) UpdateGBTx(tx *sqlx.Tx, participant_auth_key_id int64, g_b string, key_fingerprint int64, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update secret_chats set participant_auth_key_id = ?, g_b = ?, key_fingerprint = ?, state = 2 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, participant_auth_key_id, g_b, key_fingerprint, id)

	if err != nil {
		log.Errorf("exec in UpdateGB(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateGB(_), error: %v", err)
	}

	return
}

func (dao *SecretChatsDAO) UpdateState(ctx context.Context, state int8, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update secret_chats set state = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, state, id)

	if err != nil {
		log.Errorf("exec in UpdateState(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateState(_), error: %v", err)
	}

	return
}

func (dao *SecretChatsDAO) UpdateStateTx(tx *sqlx.Tx, state int8, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update secret_chats set state = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, state, id)

	if err != nil {
		log.Errorf("exec in UpdateState(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateState(_), error: %v", err)
	}

	return
}

func (dao *SecretChatsDAO) SelectRequested(ctx context.Context, selfId int32) (rList []*dataobject.SecretChatsDO, err error) {
	var (
		query = "select id, access_hash, admin_id, participant_id, admin_auth_key_id, participant_auth_key_id, random_id, g_a, g_b, key_fingerprint, state, `date` from secret_chats where participant_id = ? and state=1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, selfId)

	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		do := &dataobject.SecretChatsDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectRequested(%d), error: %v", selfId, err)
		} else {
			rList = append(rList, do)
		}
	}

	return
}
