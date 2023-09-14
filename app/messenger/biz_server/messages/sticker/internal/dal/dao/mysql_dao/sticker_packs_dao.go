package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/messenger/biz_server/messages/sticker/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type StickerPacksDAO struct {
	db *sqlx.DB
}

func NewStickerPacksDAO(db *sqlx.DB) *StickerPacksDAO {
	return &StickerPacksDAO{db}
}

func (dao *StickerPacksDAO) Insert(ctx context.Context, do *dataobject.StickerPacksDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into sticker_packs(sticker_set_id, emoticon, document_id) values (:sticker_set_id, :emoticon, :document_id)"
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

func (dao *StickerPacksDAO) InsertTx(tx *sqlx.Tx, do *dataobject.StickerPacksDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into sticker_packs(sticker_set_id, emoticon, document_id) values (:sticker_set_id, :emoticon, :document_id)"
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

func (dao *StickerPacksDAO) SelectBySetID(ctx context.Context, sticker_set_id int64) (rList []dataobject.StickerPacksDO, err error) {
	var (
		query = "select sticker_set_id, emoticon, document_id from sticker_packs where sticker_set_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, sticker_set_id)

	if err != nil {
		log.Errorf("queryx in SelectBySetID(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.StickerPacksDO
	for rows.Next() {
		v := dataobject.StickerPacksDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectBySetID(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
