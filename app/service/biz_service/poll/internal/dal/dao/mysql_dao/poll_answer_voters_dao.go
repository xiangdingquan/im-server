package mysql_dao

import (
	"context"
	"database/sql"

	"open.chat/app/service/biz_service/poll/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type PollAnswerVotersDAO struct {
	db *sqlx.DB
}

func NewPollAnswerVotersDAO(db *sqlx.DB) *PollAnswerVotersDAO {
	return &PollAnswerVotersDAO{db}
}

func (dao *PollAnswerVotersDAO) Insert(ctx context.Context, do *dataobject.PollAnswerVotersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into poll_answer_voters(poll_id, vote_user_id, options, option0, option1, option2, option3, option4, option5, option6, option7, option8, option9, date2) values (:poll_id, :vote_user_id, :options, :option0, :option1, :option2, :option3, :option4, :option5, :option6, :option7, :option8, :option9, :date2)"
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

func (dao *PollAnswerVotersDAO) InsertTx(tx *sqlx.Tx, do *dataobject.PollAnswerVotersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into poll_answer_voters(poll_id, vote_user_id, options, option0, option1, option2, option3, option4, option5, option6, option7, option8, option9, date2) values (:poll_id, :vote_user_id, :options, :option0, :option1, :option2, :option3, :option4, :option5, :option6, :option7, :option8, :option9, :date2)"
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

func (dao *PollAnswerVotersDAO) InsertOrUpdate(ctx context.Context, do *dataobject.PollAnswerVotersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into poll_answer_voters(poll_id, vote_user_id, options, option0, option1, option2, option3, option4, option5, option6, option7, option8, option9, date2, deleted) values (:poll_id, :vote_user_id, :options, :option0, :option1, :option2, :option3, :option4, :option5, :option6, :option7, :option8, :option9, :date2, :deleted) on duplicate key update options = values(options), option0 = values(option0), option1 = values(option1), option2 = values(option2), option3 = values(option3), option4 = values(option4), option5 = values(option5), option6 = values(option6), option7 = values(option7), option8 = values(option8), option9 = values(option9), date2 = values(date2), deleted = values(deleted)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		log.Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", do, err)
	}

	return
}

func (dao *PollAnswerVotersDAO) InsertOrUpdateTx(tx *sqlx.Tx, do *dataobject.PollAnswerVotersDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into poll_answer_voters(poll_id, vote_user_id, options, option0, option1, option2, option3, option4, option5, option6, option7, option8, option9, date2, deleted) values (:poll_id, :vote_user_id, :options, :option0, :option1, :option2, :option3, :option4, :option5, :option6, :option7, :option8, :option9, :date2, :deleted) on duplicate key update options = values(options), option0 = values(option0), option1 = values(option1), option2 = values(option2), option3 = values(option3), option4 = values(option4), option5 = values(option5), option6 = values(option6), option7 = values(option7), option8 = values(option8), option9 = values(option9), date2 = values(date2), deleted = values(deleted)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		log.Errorf("namedExec in InsertOrUpdate(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertOrUpdate(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertOrUpdate(%v)_error: %v", do, err)
	}

	return
}

func (dao *PollAnswerVotersDAO) Select(ctx context.Context, poll_id int64, vote_user_id int32) (rValue *dataobject.PollAnswerVotersDO, err error) {
	var (
		query = "select poll_id, vote_user_id, options, options, option0, option1, option2, option3, option4, option5, option6, option7, option8, option9 from poll_answer_voters where poll_id = ? and vote_user_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, poll_id, vote_user_id)

	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.PollAnswerVotersDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in Select(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *PollAnswerVotersDAO) SelectRecentVoters(ctx context.Context, poll_id int64, limit int32) (rList []int32, err error) {
	var query = "select vote_user_id from poll_answer_voters where poll_id = ? and deleted = 0 order by date2 desc limit ?"
	err = dao.db.Select(ctx, &rList, query, poll_id, limit)

	if err != nil {
		log.Errorf("select in SelectRecentVoters(_), error: %v", err)
	}

	return
}

func (dao *PollAnswerVotersDAO) SelectVoters(ctx context.Context, poll_id int64, date2 int64, limit int32) (rList []dataobject.PollAnswerVotersDO, err error) {
	var (
		query = "select vote_user_id, date2 from poll_answer_voters where poll_id = ? and deleted = 0 and date2 < ? order by date2 desc limit ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, poll_id, date2, limit)

	if err != nil {
		log.Errorf("queryx in SelectVoters(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.PollAnswerVotersDO
	for rows.Next() {
		v := dataobject.PollAnswerVotersDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectVoters(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
