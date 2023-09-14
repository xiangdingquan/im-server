package mysql_dao

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"open.chat/app/service/biz_service/poll/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type PollsDAO struct {
	db *sqlx.DB
}

func NewPollsDAO(db *sqlx.DB) *PollsDAO {
	return &PollsDAO{db}
}

func (dao *PollsDAO) Insert(ctx context.Context, do *dataobject.PollsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into polls(poll_id, creator, question, closed, multiple_choice, public_voters, quiz, text0, option0, correct_answer0, text1, option1, correct_answer1, text2, option2, correct_answer2, text3, option3, correct_answer3, text4, option4, correct_answer4, text5, option5, correct_answer5, text6, option6, correct_answer6, text7, option7, correct_answer7, text8, option8, correct_answer8, text9, option9, correct_answer9, date2) values (:poll_id, :creator, :question, :closed, :multiple_choice, :public_voters, :quiz, :text0, :option0, :correct_answer0, :text1, :option1, :correct_answer1, :text2, :option2, :correct_answer2, :text3, :option3, :correct_answer3, :text4, :option4, :correct_answer4, :text5, :option5, :correct_answer5, :text6, :option6, :correct_answer6, :text7, :option7, :correct_answer7, :text8, :option8, :correct_answer8, :text9, :option9, :correct_answer9, :date2)"
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

func (dao *PollsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.PollsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into polls(poll_id, creator, question, closed, multiple_choice, public_voters, quiz, text0, option0, correct_answer0, text1, option1, correct_answer1, text2, option2, correct_answer2, text3, option3, correct_answer3, text4, option4, correct_answer4, text5, option5, correct_answer5, text6, option6, correct_answer6, text7, option7, correct_answer7, text8, option8, correct_answer8, text9, option9, correct_answer9, date2) values (:poll_id, :creator, :question, :closed, :multiple_choice, :public_voters, :quiz, :text0, :option0, :correct_answer0, :text1, :option1, :correct_answer1, :text2, :option2, :correct_answer2, :text3, :option3, :correct_answer3, :text4, :option4, :correct_answer4, :text5, :option5, :correct_answer5, :text6, :option6, :correct_answer6, :text7, :option7, :correct_answer7, :text8, :option8, :correct_answer8, :text9, :option9, :correct_answer9, :date2)"
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

func (dao *PollsDAO) Select(ctx context.Context, id int64) (rValue *dataobject.PollsDO, err error) {
	var (
		query = "select id, poll_id, creator, question, closed, multiple_choice, public_voters, quiz, text0, option0, correct_answer0, voters0, text1, option1, correct_answer1, voters1, text2, option2, correct_answer2, voters2, text3, option3, correct_answer3, voters3, text4, option4, correct_answer4, voters4, text5, correct_answer5, option5, voters5, text6, option6, correct_answer6, voters6, text7, option7, correct_answer7, voters7, text8, option8, correct_answer8, voters8, text9, option9, correct_answer9, voters9 from polls where id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, id)

	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.PollsDO{}
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

func (dao *PollsDAO) Update(ctx context.Context, cMap map[string]interface{}, id int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update polls set %s where id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, id)

	rResult, err = dao.db.Exec(ctx, query, aValues...)

	if err != nil {
		log.Errorf("exec in Update(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Update(_), error: %v", err)
	}

	return
}

func (dao *PollsDAO) UpdateTx(tx *sqlx.Tx, cMap map[string]interface{}, id int64) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update polls set %s where id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, id)

	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		log.Errorf("exec in Update(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Update(_), error: %v", err)
	}

	return
}
