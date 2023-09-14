package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/media/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type DocumentsDAO struct {
	db *sqlx.DB
}

func NewDocumentsDAO(db *sqlx.DB) *DocumentsDAO {
	return &DocumentsDAO{db}
}

func (dao *DocumentsDAO) Insert(ctx context.Context, do *dataobject.DocumentsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into documents(document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, attributes) values (:document_id, :access_hash, :dc_id, :file_path, :file_size, :uploaded_file_name, :ext, :mime_type, :thumb_id, :attributes)"
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

func (dao *DocumentsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.DocumentsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into documents(document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, attributes) values (:document_id, :access_hash, :dc_id, :file_path, :file_size, :uploaded_file_name, :ext, :mime_type, :thumb_id, :attributes)"
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

func (dao *DocumentsDAO) SelectByFileLocation(ctx context.Context, document_id int64, access_hash int64, version int32) (rValue *dataobject.DocumentsDO, err error) {
	var (
		query = "select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, attributes, version from documents where dc_id = 2 and document_id = ? and access_hash = ? and version = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, document_id, access_hash, version)

	if err != nil {
		log.Errorf("queryx in SelectByFileLocation(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.DocumentsDO{}
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

func (dao *DocumentsDAO) SelectByIdList(ctx context.Context, idList []int64) (rList []dataobject.DocumentsDO, err error) {
	var (
		query = "select id, document_id, access_hash, dc_id, file_path, file_size, uploaded_file_name, ext, mime_type, thumb_id, attributes, version from documents where document_id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, idList)
	if err != nil {
		log.Errorf("sqlx.In in SelectByIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectByIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.DocumentsDO
	for rows.Next() {
		v := dataobject.DocumentsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
