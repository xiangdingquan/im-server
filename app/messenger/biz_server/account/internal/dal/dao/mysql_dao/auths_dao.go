package mysql_dao

import (
	"open.chat/pkg/database/sqlx"
)

type AuthsDAO struct {
	db *sqlx.DB
}

func NewAuthsDAO(db *sqlx.DB) *AuthsDAO {
	return &AuthsDAO{db}
}
