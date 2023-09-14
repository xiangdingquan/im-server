package mysql_dao

import (
	"context"
	"open.chat/app/service/biz_service/auth/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type AuthsDAO struct {
	db *sqlx.DB
}

func NewAuthsDAO(db *sqlx.DB) *AuthsDAO {
	return &AuthsDAO{db: db}
}

func (dao *AuthsDAO) SelectByAuthKeyId(ctx context.Context, authKeyId int64) (rValue *dataobject.AuthsDo, err error) {
	var (
		query = "select api_id,device_model,system_version from auths where auth_key_id=?"
		rows  *sqlx.Rows
	)

	rows, err = dao.db.Query(ctx, query, authKeyId)

	if err != nil {
		log.Errorf("queryx in SelectByAuthKeyId(%d), error: %v", authKeyId, err)
		return
	}

	defer rows.Close()

	do := &dataobject.AuthsDo{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByAuthKeyId(%d), error: %v", authKeyId, err)
			return
		} else {
			rValue = do
		}
	}

	return
}
