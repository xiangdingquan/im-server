package mysql_dao

import (
	"context"
	"database/sql"
	"open.chat/app/service/biz_service/blog/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type BlogUserPrivaciesDAO struct {
	db *sqlx.DB
}

func NewBlogUserPrivaciesDAO(db *sqlx.DB) *BlogUserPrivaciesDAO {
	return &BlogUserPrivaciesDAO{db}
}

func (dao *BlogUserPrivaciesDAO) InsertOrUpdate(ctx context.Context, do *dataobject.BlogUserPrivaciesDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "INSERT INTO blog_user_privacies (user_id, key_type, rules) values (:user_id, :key_type, :rules) " +
			"ON DUPLICATE KEY UPDATE rules=values(rules)"
		r sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		log.Errorf("namedExec in BlogUserPrivaciesDAO.InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in BlogUserPrivaciesDAO.InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in BlogUserPrivaciesDAO.InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	return
}

func (dao *BlogUserPrivaciesDAO) SelectPrivacy(ctx context.Context, userId int32, key int8) (rValue *dataobject.BlogUserPrivaciesDO, err error) {
	var (
		query = "select id, user_id, key_type, rules from blog_user_privacies where user_id=? and key_type=?"
	)

	err = dao.db.Get(ctx, rValue, query, userId, key)
	if err != nil {
		log.Errorf("Get in BlogUserPrivaciesDAO.SelectPrivacy(%d, %d), error: %v", userId, key, err)
	}

	return
}

func (dao *BlogUserPrivaciesDAO) SelectUserPrivacy(ctx context.Context, userId int32) (rList []dataobject.BlogUserPrivaciesDO, err error) {
	var (
		query = "select id, user_id, key_type, rules from blog_user_privacies where user_id=?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, userId)

	if err != nil {
		log.Errorf("queryx in BlogUserPrivaciesDAO.SelectUserPrivacy(%d), error: %v", userId, err)
		return
	}

	defer rows.Close()

	var values []dataobject.BlogUserPrivaciesDO
	for rows.Next() {
		v := dataobject.BlogUserPrivaciesDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in BlogUserPrivaciesDAO.SelectUserPrivacy(%d), error: %v", userId, err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
