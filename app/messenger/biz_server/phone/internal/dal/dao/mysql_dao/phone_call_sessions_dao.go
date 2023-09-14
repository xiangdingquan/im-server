package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/messenger/biz_server/phone/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type PhoneCallSessionsDAO struct {
	db *sqlx.DB
}

func NewPhoneCallSessionsDAO(db *sqlx.DB) *PhoneCallSessionsDAO {
	return &PhoneCallSessionsDAO{db}
}

func (dao *PhoneCallSessionsDAO) InsertOrGetId(ctx context.Context, do *dataobject.PhoneCallSessionsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into phone_call_sessions(access_hash, admin_id, participant_id, admin_auth_key_id, participant_auth_key_id, random_id, admin_protocol, participant_protocol, g_a_hash, g_a, g_b, key_fingerprint, connections, admin_debug_data, participant_debug_data, `date`, state) values (:access_hash, :admin_id, :participant_id, :admin_auth_key_id, :participant_auth_key_id, :random_id, :admin_protocol, '', :g_a_hash, '', '', :key_fingerprint, '', '', '', :date, 0) on duplicate key update id = last_insert_id(id)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		log.Errorf("namedExec in InsertOrGetId(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertOrGetId(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertOrGetId(%v)_error: %v", do, err)
	}

	return
}

func (dao *PhoneCallSessionsDAO) InsertOrGetIdTx(tx *sqlx.Tx, do *dataobject.PhoneCallSessionsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into phone_call_sessions(access_hash, admin_id, participant_id, admin_auth_key_id, participant_auth_key_id, random_id, admin_protocol, participant_protocol, g_a_hash, g_a, g_b, key_fingerprint, connections, admin_debug_data, participant_debug_data, `date`, state) values (:access_hash, :admin_id, :participant_id, :admin_auth_key_id, :participant_auth_key_id, :random_id, :admin_protocol, '', :g_a_hash, '', '', :key_fingerprint, '', '', '', :date, 0) on duplicate key update id = last_insert_id(id)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		log.Errorf("namedExec in InsertOrGetId(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertOrGetId(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertOrGetId(%v)_error: %v", do, err)
	}

	return
}

func (dao *PhoneCallSessionsDAO) Select(ctx context.Context, id int64) (rValue *dataobject.PhoneCallSessionsDO, err error) {
	var (
		query = "select id, access_hash, admin_id, participant_id, admin_auth_key_id, participant_auth_key_id, random_id, admin_protocol, participant_protocol, g_a_hash, g_a, g_b, key_fingerprint, connections, admin_debug_data, participant_debug_data, `date`, state from phone_call_sessions where id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, id)

	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.PhoneCallSessionsDO{}
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

func (dao *PhoneCallSessionsDAO) UpdateParticipant(ctx context.Context, g_b string, participant_auth_key_id int64, participant_protocol string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update phone_call_sessions set g_b = ?, participant_auth_key_id = ?, participant_protocol = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, g_b, participant_auth_key_id, participant_protocol, id)

	if err != nil {
		log.Errorf("exec in UpdateParticipant(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateParticipant(_), error: %v", err)
	}

	return
}

func (dao *PhoneCallSessionsDAO) UpdateParticipantTx(tx *sqlx.Tx, g_b string, participant_auth_key_id int64, participant_protocol string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update phone_call_sessions set g_b = ?, participant_auth_key_id = ?, participant_protocol = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, g_b, participant_auth_key_id, participant_protocol, id)

	if err != nil {
		log.Errorf("exec in UpdateParticipant(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateParticipant(_), error: %v", err)
	}

	return
}

func (dao *PhoneCallSessionsDAO) UpdateAdminDebugData(ctx context.Context, admin_debug_data string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update phone_call_sessions set admin_debug_data = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, admin_debug_data, id)

	if err != nil {
		log.Errorf("exec in UpdateAdminDebugData(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateAdminDebugData(_), error: %v", err)
	}

	return
}

func (dao *PhoneCallSessionsDAO) UpdateAdminDebugDataTx(tx *sqlx.Tx, admin_debug_data string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update phone_call_sessions set admin_debug_data = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, admin_debug_data, id)

	if err != nil {
		log.Errorf("exec in UpdateAdminDebugData(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateAdminDebugData(_), error: %v", err)
	}

	return
}

func (dao *PhoneCallSessionsDAO) UpdateParticipantDebugData(ctx context.Context, participant_debug_data string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update phone_call_sessions set participant_debug_data = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, participant_debug_data, id)

	if err != nil {
		log.Errorf("exec in UpdateParticipantDebugData(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateParticipantDebugData(_), error: %v", err)
	}

	return
}

func (dao *PhoneCallSessionsDAO) UpdateParticipantDebugDataTx(tx *sqlx.Tx, participant_debug_data string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update phone_call_sessions set participant_debug_data = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, participant_debug_data, id)

	if err != nil {
		log.Errorf("exec in UpdateParticipantDebugData(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateParticipantDebugData(_), error: %v", err)
	}

	return
}

func (dao *PhoneCallSessionsDAO) UpdateAdminCallRating(ctx context.Context, admin_rating int32, admin_comment string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update phone_call_sessions set admin_rating = ?, admin_comment = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, admin_rating, admin_comment, id)

	if err != nil {
		log.Errorf("exec in UpdateAdminCallRating(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateAdminCallRating(_), error: %v", err)
	}

	return
}

func (dao *PhoneCallSessionsDAO) UpdateAdminCallRatingTx(tx *sqlx.Tx, admin_rating int32, admin_comment string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update phone_call_sessions set admin_rating = ?, admin_comment = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, admin_rating, admin_comment, id)

	if err != nil {
		log.Errorf("exec in UpdateAdminCallRating(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateAdminCallRating(_), error: %v", err)
	}

	return
}

func (dao *PhoneCallSessionsDAO) UpdateParticipantCallRating(ctx context.Context, participant_rating int32, participant_comment string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update phone_call_sessions set participant_rating = ?, participant_comment = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, participant_rating, participant_comment, id)

	if err != nil {
		log.Errorf("exec in UpdateParticipantCallRating(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateParticipantCallRating(_), error: %v", err)
	}

	return
}

func (dao *PhoneCallSessionsDAO) UpdateParticipantCallRatingTx(tx *sqlx.Tx, participant_rating int32, participant_comment string, id int64) (rowsAffected int64, err error) {
	var (
		query   = "update phone_call_sessions set participant_rating = ?, participant_comment = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, participant_rating, participant_comment, id)

	if err != nil {
		log.Errorf("exec in UpdateParticipantCallRating(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateParticipantCallRating(_), error: %v", err)
	}

	return
}
