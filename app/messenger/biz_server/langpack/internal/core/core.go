package core

import (
	"open.chat/app/messenger/biz_server/langpack/internal/dao"
)

type LangPackCore struct {
	*dao.Dao
}

func New(dao *dao.Dao) *LangPackCore {
	return &LangPackCore{dao}
}
