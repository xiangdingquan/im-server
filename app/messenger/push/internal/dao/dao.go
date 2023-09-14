package dao

import (
	"context"
	"encoding/base64"
	"time"

	"github.com/go-kratos/kratos/pkg/conf/paladin"

	"open.chat/app/messenger/push/internal/dal/dao/mysql_dao"
	"open.chat/app/messenger/push/internal/dal/dataobject"
	idgen "open.chat/app/service/idgen/client"
	status_client "open.chat/app/service/status/client"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

const (
	seqUpdatesNgenId = "seq_updates_ngen_"
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.DevicesDAO
	*mysql_dao.AuthSeqUpdatesDAO
}

func newMysqlDao(c *sqlx.Config) *Mysql {
	db := sqlx.NewMySQL(c)
	return &Mysql{
		DB:                db,
		DevicesDAO:        mysql_dao.NewDevicesDAO(db),
		AuthSeqUpdatesDAO: mysql_dao.NewAuthSeqUpdatesDAO(db),
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
	// *fcm.Client
	*Mysql
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// New new a dao and return.
func New() (dao *Dao) {
	var (
		dc struct {
			Mysql *sqlx.Config
		}
	)
	checkErr(paladin.Get("mysql.toml").UnmarshalTOML(&dc))

	dao = &Dao{
		Mysql: newMysqlDao(dc.Mysql),
	}

	// init status
	status_client.New()
	idgen.NewSeqIDGen()

	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.Mysql.Close()
}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	return d.Mysql.Ping(ctx)
}

func (d *Dao) WalkDevices(ctx context.Context,
	userId int32,
	excludes []int64,
	pushData interface{},
	cb func(ctx context.Context, authKeyId int64, tokenType int8, token string, secret []byte, pushData interface{})) error {

	doList, err := d.DevicesDAO.SelectListByUser(ctx, userId)
	if err != nil {
		return err
	}

	for i := 0; i < len(doList); i++ {
		if ok, _ := util.Contains(doList[i].AuthKeyId, excludes); !ok {
			var secret []byte

			if len(doList[i].Secret) > 0 {
				secret, err = base64.RawStdEncoding.DecodeString(doList[i].Secret)
				if err != nil {
					log.Errorf("invalid secret - device: %d, error: %v", doList[i], err)
					continue
				}
			}

			cb(ctx, doList[i].AuthKeyId, doList[i].TokenType, doList[i].Token, secret, pushData)
		}
	}
	return nil
}

func (d *Dao) NextSeqId(ctx context.Context, key int64) (seq int64) {
	seq, _ = idgen.GetNextSeqID(ctx, seqUpdatesNgenId+util.Int64ToString(key))
	return
}

func (d *Dao) AddSeqToUpdatesQueue(ctx context.Context, authId int64, userId, updateType int32, updateData []byte) int32 {
	seq := int32(d.NextSeqId(ctx, authId))
	do := &dataobject.AuthSeqUpdatesDO{
		AuthId:     authId,
		UserId:     userId,
		UpdateType: updateType,
		UpdateData: updateData,
		Date2:      int32(time.Now().Unix()),
		Seq:        seq,
	}

	i, _, _ := d.AuthSeqUpdatesDAO.Insert(ctx, do)
	return int32(i)
}

func (d *Dao) GetUserLangCode(ctx context.Context, user_id int32) (rValue string) {
	var (
		query = "select auth_key_id from auth_users where user_id = ? and deleted = 0 order by updated_at desc limit 1"
	)
	rValue = "en"
	var auth_key_id int64 = 0
	err := d.Get(ctx, &auth_key_id, query, user_id)

	if err != nil {
		log.Errorf("get in GetUserLangCode(_), error: %v", err)
		return
	}

	if err == nil && auth_key_id > 0 {
		var query = "select system_lang_code from auths where auth_key_id = ? limit 1"
		err = d.Get(ctx, &rValue, query, auth_key_id)
		if err != nil {
			log.Errorf("get in GetUserLangCode(_), error: %v", err)
			return
		}
	}
	return
}
