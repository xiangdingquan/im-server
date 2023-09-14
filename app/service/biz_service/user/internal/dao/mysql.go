package dao

import (
	"context"

	"open.chat/app/pkg/mysql_util"
	"open.chat/app/service/biz_service/user/internal/dal/dao/mysql_dao"
	"open.chat/pkg/database/sqlx"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.UsersDAO
	*mysql_dao.UserPresencesDAO
	*mysql_dao.UserContactsDAO
	*mysql_dao.UserPasswordsDAO
	*mysql_dao.UserBlocksDAO
	*mysql_dao.BotsDAO
	*mysql_dao.BotCommandsDAO
	*mysql_dao.UserPrivaciesDAO
	*mysql_dao.UsernameDAO
	*mysql_dao.UserNotifySettingsDAO
	*mysql_dao.WallPapersDAO
	*mysql_dao.ReportsDAO
	*mysql_dao.UnregisteredContactsDAO
	*mysql_dao.PopularContactsDAO
	*mysql_dao.ImportedContactsDAO
	*mysql_dao.PhoneBooksDAO
	*mysql_dao.PredefinedUsersDAO
	*mysql_dao.UserPeerSettingsDAO
	*sqlx.CommonDAO
	*mysql_dao.UserWalletDAO
	*mysql_dao.UserBlogsDAO
}

func newMysqlDao() *Mysql {
	db := mysql_util.GetSingletonSqlxDB()
	return &Mysql{
		DB:                      db,
		UsersDAO:                mysql_dao.NewUsersDAO(db),
		UserPresencesDAO:        mysql_dao.NewUserPresencesDAO(db),
		UserContactsDAO:         mysql_dao.NewUserContactsDAO(db),
		UserPasswordsDAO:        mysql_dao.NewUserPasswordsDAO(db),
		UserBlocksDAO:           mysql_dao.NewUserBlocksDAO(db),
		BotsDAO:                 mysql_dao.NewBotsDAO(db),
		BotCommandsDAO:          mysql_dao.NewBotCommandsDAO(db),
		UserPrivaciesDAO:        mysql_dao.NewUserPrivaciesDAO(db),
		UsernameDAO:             mysql_dao.NewUsernameDAO(db),
		UserNotifySettingsDAO:   mysql_dao.NewUserNotifySettingsDAO(db),
		WallPapersDAO:           mysql_dao.NewWallPapersDAO(db),
		ReportsDAO:              mysql_dao.NewReportsDAO(db),
		UnregisteredContactsDAO: mysql_dao.NewUnregisteredContactsDAO(db),
		PopularContactsDAO:      mysql_dao.NewPopularContactsDAO(db),
		ImportedContactsDAO:     mysql_dao.NewImportedContactsDAO(db),
		PhoneBooksDAO:           mysql_dao.NewPhoneBooksDAO(db),
		PredefinedUsersDAO:      mysql_dao.NewPredefinedUsersDAO(db),
		UserPeerSettingsDAO:     mysql_dao.NewUserPeerSettingsDAO(db),
		CommonDAO:               sqlx.NewCommonDAO(db),
		UserWalletDAO:           mysql_dao.NewUserWalletDAO(db),
		UserBlogsDAO:            mysql_dao.NewUserBlogsDAO(db),
	}
}

func (d *Mysql) Close() error {
	return d.DB.Close()
}

func (d *Mysql) Ping(ctx context.Context) (err error) {
	return d.DB.Ping(ctx)
}
