package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/media/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type PhotoDatasDAO struct {
	db *sqlx.DB
}

func NewPhotoDatasDAO(db *sqlx.DB) *PhotoDatasDAO {
	return &PhotoDatasDAO{db}
}

func (dao *PhotoDatasDAO) Insert(ctx context.Context, do *dataobject.PhotoDatasDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into photo_datas(photo_id, photo_type, dc_id, volume_id, local_id, access_hash, width, height, file_size, file_path, ext) values (:photo_id, :photo_type, :dc_id, :volume_id, :local_id, :access_hash, :width, :height, :file_size, :file_path, :ext)"
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

func (dao *PhotoDatasDAO) InsertTx(tx *sqlx.Tx, do *dataobject.PhotoDatasDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into photo_datas(photo_id, photo_type, dc_id, volume_id, local_id, access_hash, width, height, file_size, file_path, ext) values (:photo_id, :photo_type, :dc_id, :volume_id, :local_id, :access_hash, :width, :height, :file_size, :file_path, :ext)"
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

func (dao *PhotoDatasDAO) SelectByFileLocation(ctx context.Context, volume_id int64, local_id int32, access_hash int64) (rValue *dataobject.PhotoDatasDO, err error) {
	var (
		query = "select id, photo_id, photo_type, dc_id, volume_id, local_id, access_hash, width, height, file_size, file_path, ext from photo_datas where dc_id = 2 and volume_id = ? and local_id = ? and access_hash = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, volume_id, local_id, access_hash)

	if err != nil {
		log.Errorf("queryx in SelectByFileLocation(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.PhotoDatasDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByFileLocation(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *PhotoDatasDAO) SelectAccessHash(ctx context.Context, volume_id int64, local_id int32) (rValue int64, err error) {
	var query = "select access_hash from photo_datas where dc_id = 2 and volume_id = ? and local_id = ? limit 1"
	err = dao.db.Get(ctx, &rValue, query, volume_id, local_id)

	if err != nil {
		log.Errorf("get in SelectAccessHash(_), error: %v", err)
	}

	return
}

func (dao *PhotoDatasDAO) SelectListByPhotoId(ctx context.Context, photo_id int64) (rList []dataobject.PhotoDatasDO, err error) {
	var (
		query = "select id, photo_id, photo_type, dc_id, volume_id, local_id, access_hash, width, height, file_size, file_path, ext from photo_datas where photo_id = ? order by local_id asc"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, photo_id)

	if err != nil {
		log.Errorf("queryx in SelectListByPhotoId(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.PhotoDatasDO
	for rows.Next() {
		v := dataobject.PhotoDatasDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectListByPhotoId(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
