package dao

import (
	"context"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	relay_client "open.chat/app/interface/relay/client"
	"open.chat/app/messenger/biz_server/phone/internal/dal/dao/mysql_dao"
	"open.chat/app/messenger/biz_server/phone/internal/dal/dataobject"
	"open.chat/app/pkg/mysql_util"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/util"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.PhoneCallDebugsDAO
	*mysql_dao.PhoneCallRatingsDAO
}

func newMysqlDao(c *sqlx.Config) *Mysql {
	db := mysql_util.GetSingletonSqlxDB()
	return &Mysql{
		DB:                  db,
		PhoneCallDebugsDAO:  mysql_dao.NewPhoneCallDebugsDAO(db),
		PhoneCallRatingsDAO: mysql_dao.NewPhoneCallRatingsDAO(db),
	}
}

func (d *Mysql) Close() error {
	return d.DB.Close()
}

func (d *Mysql) Ping(ctx context.Context) (err error) {
	return d.DB.Ping(ctx)
}

// Dao dao.
type Dao struct {
	*Mysql
	*Redis
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func New() (dao *Dao) {
	var (
		dc struct {
			Mysql *sqlx.Config
		}
	)
	checkErr(paladin.Get("mysql.toml").UnmarshalTOML(&dc))
	dao = &Dao{
		Mysql: newMysqlDao(dc.Mysql),
		Redis: newRedisDao(),
	}

	relay_client.New()
	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.Redis.Close()
	d.Mysql.Close()
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	if err = d.Redis.Ping(ctx); err != nil {
		return
	}
	return d.Mysql.Ping(ctx)
}

// Ping ping the resource.
func (d *Dao) SetCallDebug(ctx context.Context, callId int64, userId int32, authId int64, debugData string) (err error) {
	_, _, err = d.PhoneCallDebugsDAO.InsertIgnore(ctx, &dataobject.PhoneCallDebugsDO{
		CallId:               callId,
		ParticipantId:        userId,
		ParticipantAuthKeyId: authId,
		DebugData:            debugData,
	})
	return
}

// Ping ping the resource.
func (d *Dao) SetCallRating(ctx context.Context, callId int64, userId int32, authId int64, initiative bool, rating int32, comment string) (err error) {
	_, _, err = d.PhoneCallRatingsDAO.InsertIgnore(ctx, &dataobject.PhoneCallRatingsDO{
		CallId:               callId,
		ParticipantId:        userId,
		ParticipantAuthKeyId: authId,
		UserInitiative:       util.BoolToInt8(initiative),
		Rating:               rating,
		Comment:              comment,
	})
	return
}
