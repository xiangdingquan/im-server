package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/messenger/biz_server/phone/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type PhoneCallRatingsDAO struct {
	db *sqlx.DB
}

func NewPhoneCallRatingsDAO(db *sqlx.DB) *PhoneCallRatingsDAO {
	return &PhoneCallRatingsDAO{db}
}

func (dao *PhoneCallRatingsDAO) InsertIgnore(ctx context.Context, do *dataobject.PhoneCallRatingsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into phone_call_ratings(call_id, participant_id, participant_auth_key_id, user_initiative, rating, `comment`) values (:call_id, :participant_id, :participant_auth_key_id, :user_initiative, :rating, :comment)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		log.Errorf("namedExec in InsertIgnore(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertIgnore(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertIgnore(%v)_error: %v", do, err)
	}

	return
}

func (dao *PhoneCallRatingsDAO) InsertIgnoreTx(tx *sqlx.Tx, do *dataobject.PhoneCallRatingsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into phone_call_ratings(call_id, participant_id, participant_auth_key_id, user_initiative, rating, `comment`) values (:call_id, :participant_id, :participant_auth_key_id, :user_initiative, :rating, :comment)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		log.Errorf("namedExec in InsertIgnore(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertIgnore(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertIgnore(%v)_error: %v", do, err)
	}

	return
}
