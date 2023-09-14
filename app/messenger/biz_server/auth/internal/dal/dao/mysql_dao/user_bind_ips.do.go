package mysql_dao

import (
	"context"
	"encoding/json"

	"open.chat/app/messenger/biz_server/auth/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

// BannedIpDAO .
type UserBindIpsDAO struct {
	db *sqlx.DB
}

// NewBannedIpDAO .
func NewUserBindIpsDAO(db *sqlx.DB) *UserBindIpsDAO {
	return &UserBindIpsDAO{db}
}

// Select .
func (dao *UserBindIpsDAO) Select(ctx context.Context, userId int32) (rValue *dataobject.UserBindIps, err error) {
	var (
		query = "SELECT id, user_id, ip_addrs FROM `user_bind_ips` WHERE user_id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, userId)

	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.UserBindIps{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in Select(_), error: %v", err)
		} else {
			json.Unmarshal([]byte(do.IpAddrs), &do.IpList)
			rValue = do
		}
	}

	return
}
