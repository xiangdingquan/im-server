package mysql_dao

import (
	"context"

	"open.chat/app/messenger/biz_server/help/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type AppConfigsDAO struct {
	db *sqlx.DB
}

func NewAppConfigsDAO(db *sqlx.DB) *AppConfigsDAO {
	return &AppConfigsDAO{db}
}

func (dao *AppConfigsDAO) SelectList(ctx context.Context) (rList []dataobject.AppConfigsDO, err error) {
	var (
		query = "select key2, type2, value2 from app_configs where deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query)

	if err != nil {
		log.Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.AppConfigsDO
	for rows.Next() {
		v := dataobject.AppConfigsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}
