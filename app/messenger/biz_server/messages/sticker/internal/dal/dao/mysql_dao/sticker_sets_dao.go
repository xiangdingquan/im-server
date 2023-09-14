package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/messenger/biz_server/messages/sticker/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type StickerSetsDAO struct {
	db *sqlx.DB
}

func NewStickerSetsDAO(db *sqlx.DB) *StickerSetsDAO {
	return &StickerSetsDAO{db}
}

func (dao *StickerSetsDAO) Insert(ctx context.Context, do *dataobject.StickerSetsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into sticker_sets(sticker_set_id, access_hash) values (:sticker_set_id, :access_hash)"
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

func (dao *StickerSetsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.StickerSetsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into sticker_sets(sticker_set_id, access_hash) values (:sticker_set_id, :access_hash)"
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

func (dao *StickerSetsDAO) SelectAll(ctx context.Context) (rList []dataobject.StickerSetsDO, err error) {
	var (
		query = "select sticker_set_id, access_hash, title, short_name, count, hash, official, masks, archived, animated, installed_date, thumb, thumb_dc_id from sticker_sets where hash > 0 and short_name not in ('AnimatedEmojies')"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query)

	if err != nil {
		log.Errorf("queryx in SelectAll(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.StickerSetsDO
	for rows.Next() {
		v := dataobject.StickerSetsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectAll(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *StickerSetsDAO) SelectByID(ctx context.Context, sticker_set_id int64, access_hash int64) (rValue *dataobject.StickerSetsDO, err error) {
	var (
		query = "select sticker_set_id, access_hash, title, short_name, count, hash, official, masks, archived, animated, installed_date, thumb, thumb_dc_id from sticker_sets where sticker_set_id = ? and access_hash = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, sticker_set_id, access_hash)

	if err != nil {
		log.Errorf("queryx in SelectByID(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.StickerSetsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByID(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *StickerSetsDAO) SelectByShortName(ctx context.Context, short_name string) (rValue *dataobject.StickerSetsDO, err error) {
	var (
		query = "select sticker_set_id, access_hash, title, short_name, count, hash, official, masks, archived, animated, installed_date, thumb, thumb_dc_id from sticker_sets where short_name = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, short_name)

	if err != nil {
		log.Errorf("queryx in SelectByShortName(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.StickerSetsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByShortName(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}
