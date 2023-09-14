package mysql_dao

import (
	"database/sql"
	"open.chat/app/service/biz_service/blog/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type BlogTopicMappingsDAO struct {
	db *sqlx.DB
}

func NewBlogTopicMappingsDAO(db *sqlx.DB) *BlogTopicMappingsDAO {
	return &BlogTopicMappingsDAO{db}
}

func (dao *BlogTopicMappingsDAO) InsertTx(tx *sqlx.Tx, doList []*dataobject.BlogTopicMappingsDo) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into blog_topic_mappings(topic_id, moment_id) values (:topic_id, :moment_id)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, doList)
	if err != nil {
		log.Errorf("namedExec in InsertTx(_), error: %v", err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertTx(_)_error: %v", err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertTx(_)_error: %v", err)
	}

	return
}
