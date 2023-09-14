package mysqldao

import "open.chat/pkg/database/sqlx"

type RemittanceRecordDao struct {
	db *sqlx.DB
}

func NewRemittanceRecordDao(db *sqlx.DB) *RemittanceRecordDao {
	return &RemittanceRecordDao{db: db}
}
