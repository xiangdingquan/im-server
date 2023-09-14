package core

import (
	"open.chat/app/service/biz_service/report/internal/dao"
)

type ReportCore struct {
	*dao.Dao
}

func New(dao *dao.Dao) *ReportCore {
	return &ReportCore{Dao: dao}
}
