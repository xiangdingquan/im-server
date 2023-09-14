package mysqldao

import (
	"context"
	"database/sql"
	"open.chat/app/json/db/dbo"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type MessageReactionDAO struct {
	db *sqlx.DB
}

func NewMessageReactionDAO(db *sqlx.DB) *MessageReactionDAO {
	return &MessageReactionDAO{db}
}

func (dao *MessageReactionDAO) Insert(ctx context.Context, do *dbo.MessageReactionDo) (lastInsertID, rowsAffected int64, err error) {
	var (
		query = `INSERT INTO message_reactions (type, chat_id, message_id, user_id, reaction_id) VALUES
				(:type, :chat_id, :message_id, :user_id, :reaction_id)
				ON DUPLICATE KEY UPDATE reaction_id=VALUES(reaction_id)`
		r sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		log.Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertID, err = r.LastInsertId()
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

func (dao *MessageReactionDAO) SelectReaction(ctx context.Context, chatType int8, chatId int64, messageIds []int32) (rList []*dbo.MessageReactionDo, err error) {
	var (
		query = "SELECT type,chat_id,message_id,user_id,reaction_id FROM message_reactions where type=? and chat_id=? and message_id in (?) and reaction_id!=0"
		rows  *sqlx.Rows
	)

	query, args, err := sqlx.In(query, chatType, chatId, messageIds)
	if err != nil {
		log.Error("sqlx.In in SelectReaction(%d, %d, _), error: %v", chatType, chatId, err)
		return
	}

	rows, err = dao.db.Query(ctx, query, args...)
	if err != nil {
		log.Error("queryx in SelectReaction(%d, %d, _), error: %v", chatType, chatId, err)
		return
	}

	defer rows.Close()

	l := make([]*dbo.MessageReactionDo, 0)
	for rows.Next() {
		v := &dbo.MessageReactionDo{}
		err = rows.StructScan(v)
		if err != nil {
			log.Error("sqlx.In in SelectReaction(%d, %d, _), error: %v", chatType, chatId, err)
			return
		}
		l = append(l, v)
	}
	rList = l
	return
}
