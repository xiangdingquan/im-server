package mysql_dao

import (
	"context"

	"open.chat/app/bots/gif/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type GiphyDatasDAO struct {
	db *sqlx.DB
}

func NewGiphyDatasDAO(db *sqlx.DB) *GiphyDatasDAO {
	return &GiphyDatasDAO{db}
}

func (dao *GiphyDatasDAO) SelectAll(ctx context.Context) (rList []dataobject.GiphyDatasDO, err error) {
	var (
		query = "select giphy_id, document_id, photo_id from giphy_datas"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query)

	if err != nil {
		log.Errorf("queryx in SelectAll(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.GiphyDatasDO
	for rows.Next() {
		v := dataobject.GiphyDatasDO{}
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

func (dao *GiphyDatasDAO) SelectById(ctx context.Context) (rValue *dataobject.GiphyDatasDO, err error) {
	var (
		query = "select giphy_id, document_id, photo_id from giphy_datas"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query)

	if err != nil {
		log.Errorf("queryx in SelectById(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.GiphyDatasDO{}
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
