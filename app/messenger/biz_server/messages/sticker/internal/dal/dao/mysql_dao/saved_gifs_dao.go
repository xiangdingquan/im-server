package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/messenger/biz_server/messages/sticker/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type SavedGifsDAO struct {
	db *sqlx.DB
}

func NewSavedGifsDAO(db *sqlx.DB) *SavedGifsDAO {
	return &SavedGifsDAO{db}
}

func (dao *SavedGifsDAO) InsertIgnore(ctx context.Context, do *dataobject.SavedGifsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into saved_gifs(user_id, gif_id) values (:user_id, :gif_id)"
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

func (dao *SavedGifsDAO) InsertIgnoreTx(tx *sqlx.Tx, do *dataobject.SavedGifsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert ignore into saved_gifs(user_id, gif_id) values (:user_id, :gif_id)"
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

func (dao *SavedGifsDAO) SelectAll(ctx context.Context, user_id int32) (rList []int64, err error) {
	var query = "select gif_id from saved_gifs where user_id = ?"
	err = dao.db.Select(ctx, &rList, query, user_id)

	if err != nil {
		log.Errorf("select in SelectAll(_), error: %v", err)
	}

	return
}

func (dao *SavedGifsDAO) Delete(ctx context.Context, user_id int32, gif_id int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from saved_gifs where user_id = ? and gif_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, user_id, gif_id)

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

func (dao *SavedGifsDAO) DeleteTx(tx *sqlx.Tx, user_id int32, gif_id int64) (rowsAffected int64, err error) {
	var (
		query   = "delete from saved_gifs where user_id = ? and gif_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, user_id, gif_id)

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
