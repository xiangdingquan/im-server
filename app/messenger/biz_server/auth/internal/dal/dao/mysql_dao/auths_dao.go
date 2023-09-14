package mysql_dao

import (
	"context"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type AuthsDAO struct {
	db *sqlx.DB
}

func NewAuthsDAO(db *sqlx.DB) *AuthsDAO {
	return &AuthsDAO{db}
}

func (a *AuthsDAO) SelectUserListByIp(ctx context.Context, ip string) (rList []int32, err error) {
	var (
		sql  = "SELECT user_id FROM auths l inner join auth_users r on(l.auth_key_id = r.auth_key_id) WHERE client_ip =? GROUP BY user_id"
		rows *sqlx.Rows
	)
	rows, err = a.db.Query(ctx, sql, ip)

	if err != nil {
		log.Errorf("queryx in GetUserListByIp(_), error: %v", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var userId int32
		err = rows.Scan(&userId)
		if err != nil {
			log.Errorf("scan in GetUserListByIp(_), error: %v", err)
			return
		}
		rList = append(rList, userId)
	}

	return
}
