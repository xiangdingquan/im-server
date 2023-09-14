package mysql_dao

import (
	"context"
	"fmt"

	"open.chat/app/service/biz_service/poll/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

func (dao *PollAnswerVotersDAO) SelectOptionVoters(ctx context.Context, poll_id int64, option string, date2 int64, limit int32) (rList []dataobject.PollAnswerVotersDO, err error) {
	var (
		query = fmt.Sprintf("select vote_user_id from poll_answer_voters where poll_id = ? and %s = 1 and deleted = 0 and date2 < ? order by date2 desc limit ?", option)
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
