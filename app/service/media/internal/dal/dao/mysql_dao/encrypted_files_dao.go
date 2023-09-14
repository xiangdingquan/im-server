package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/media/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type EncryptedFilesDAO struct {
	db *sqlx.DB
}

func NewEncryptedFilesDAO(db *sqlx.DB) *EncryptedFilesDAO {
	return &EncryptedFilesDAO{db}
}

func (dao *EncryptedFilesDAO) Insert(ctx context.Context, do *dataobject.EncryptedFilesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into encrypted_files(encrypted_file_id, access_hash, dc_id, file_size, key_fingerprint, md5_checksum, file_path) values (:encrypted_file_id, :access_hash, :dc_id, :file_size, :key_fingerprint, :md5_checksum, :file_path)"
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

func (dao *EncryptedFilesDAO) InsertTx(tx *sqlx.Tx, do *dataobject.EncryptedFilesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into encrypted_files(encrypted_file_id, access_hash, dc_id, file_size, key_fingerprint, md5_checksum, file_path) values (:encrypted_file_id, :access_hash, :dc_id, :file_size, :key_fingerprint, :md5_checksum, :file_path)"
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

func (dao *EncryptedFilesDAO) SelectByFileLocation(ctx context.Context, encrypted_file_id int64, access_hash int64) (rValue *dataobject.EncryptedFilesDO, err error) {
	var (
		query = "select id, encrypted_file_id, access_hash, dc_id, file_size, key_fingerprint, md5_checksum, file_path from encrypted_files where dc_id = 2 and encrypted_file_id = ? and access_hash = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, encrypted_file_id, access_hash)

	if err != nil {
		log.Errorf("queryx in SelectByFileLocation(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.EncryptedFilesDO{}
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

func (dao *EncryptedFilesDAO) SelectByIdList(ctx context.Context, idList []int64) (rList []dataobject.EncryptedFilesDO, err error) {
	var (
		query = "select id, encrypted_file_id, access_hash, dc_id, file_size, key_fingerprint, md5_checksum, file_path from encrypted_files where encrypted_file_id in (?)"
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

	var values []dataobject.EncryptedFilesDO
	for rows.Next() {
		v := dataobject.EncryptedFilesDO{}
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
