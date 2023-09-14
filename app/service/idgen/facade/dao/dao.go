package dao

import (
	"context"
	"database/sql"

	"open.chat/app/pkg/mysql_util"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type Mysql struct {
	*sqlx.DB
}

func newMysqlDao() *Mysql {
	db := mysql_util.GetSingletonSqlxDB()
	return &Mysql{
		DB: db,
	}
}

func (d *Mysql) Close() error {
	return d.DB.Close()
}

func (d *Mysql) Ping(ctx context.Context) (err error) {
	return d.DB.Ping(ctx)
}

// Dao dao.
type Dao struct {
	*Mysql
}

// New new a dao and return.
func New() (dao *Dao) {
	dao = &Dao{
		Mysql: newMysqlDao(),
	}
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.Mysql.Close()
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	return d.Mysql.Ping(ctx)
}

func (dao *Dao) GetNextSeqTx(tx *sqlx.Tx, prefix string) (seq int64, err error) {
	var (
		query   = "SELECT `seq` FROM `phone_number_seqs` WHERE prefix = ? FOR UPDATE"
		rows    *sqlx.Rows
		rResult sql.Result
	)

	rows, err = tx.Query(query, prefix)
	if err != nil {
		log.Errorf("exec in GetNextSeqTx(_), error: %v", err)
		return
	}
	if rows.Next() {
		rows.Scan(&seq)
	}
	rows.Close()

	query = "UPDATE `phone_number_seqs` SET `seq` = `seq` + 1 WHERE prefix = ?"
	rResult, err = tx.Exec(query, prefix)
	if err != nil {
		log.Errorf("exec in Update seq(_), error: %v", err)
		return
	}
	_, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Update seq(_), error: %v", err)
	}

	return
}
