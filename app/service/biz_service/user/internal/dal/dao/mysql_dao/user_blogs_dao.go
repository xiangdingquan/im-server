package mysql_dao

import (
	"database/sql"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

// UserBlogsDAO .
type UserBlogsDAO struct {
	db *sqlx.DB
}

// NewAvcallsDAO .
func NewUserBlogsDAO(db *sqlx.DB) *UserBlogsDAO {
	return &UserBlogsDAO{db}
}

// InsertTx .
func (dao *UserBlogsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.UserBlogsDO) (lastInsertID, rowsAffected int64, err error) {
	var (
		query string = "INSERT INTO `blogs`(user_id, `date`) VALUES (:user_id, :date)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		log.Errorf("namedExec in InsertTx(%v), error: %v", do, err)
		return
	}

	lastInsertID, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertTx(%v)_error: %v", do, err)
		return
	}

	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertTx(%v)_error: %v", do, err)
	}

	return
}
