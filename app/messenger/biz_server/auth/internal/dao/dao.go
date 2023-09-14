package dao

import (
	"context"
	"flag"
	"net"

	"open.chat/pkg/log"

	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	"github.com/oschwald/geoip2-golang"

	"open.chat/app/messenger/biz_server/auth/internal/dal/dao/mysql_dao"
	"open.chat/app/pkg/mysql_util"
	"open.chat/app/service/auth_session/authsessionpb"
	authsession_client "open.chat/app/service/auth_session/client"
	idgen "open.chat/app/service/idgen/client"
	status_client "open.chat/app/service/status/client"
	"open.chat/pkg/database/sqlx"
)

var (
	mmdb string
)

func init() {
	flag.StringVar(&mmdb, "mmdb", "./GeoLite2-City.mmdb", "mmdb")
}

type Mysql struct {
	*sqlx.DB
	*mysql_dao.AuthOpLogsDAO
	*mysql_dao.BannedIpDAO
	*mysql_dao.AuthsDAO
	*mysql_dao.UserBindIpsDAO
}

func newMysqlDao() *Mysql {
	db := mysql_util.GetSingletonSqlxDB()
	return &Mysql{
		DB:             db,
		AuthOpLogsDAO:  mysql_dao.NewAuthOpLogsDAO(db),
		BannedIpDAO:    mysql_dao.NewBannedIpDAO(db),
		AuthsDAO:       mysql_dao.NewAuthsDAO(db),
		UserBindIpsDAO: mysql_dao.NewUserBindIpsDAO(db),
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
	AuthSessionRpcClient authsessionpb.RPCSessionClient
	MMDB                 *geoip2.Reader
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// New new a dao and return.
func New(c *warden.ClientConfig) (dao *Dao) {
	dao = &Dao{
		Mysql: newMysqlDao(),
		Redis: newRedisDao(),
	}

	var err error
	dao.AuthSessionRpcClient, err = authsession_client.New(c)
	checkErr(err)
	dao.MMDB, err = geoip2.Open(mmdb)
	checkErr(err)
	status_client.New()
	idgen.NewSeqIDGen()

	return
}

// Close close the resource.
func (d *Dao) Close() {
	d.MMDB.Close()
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

func (d *Dao) GetCountryAndRegionByIp(ip string) (string, string) {
	r, err := d.MMDB.City(net.ParseIP(ip))
	if err != nil {
		log.Errorf("getCountryAndRegionByIp - error: %v", err)
		return "", ""
	}

	return r.City.Names["en"] + ", " + r.Country.Names["en"], r.Country.IsoCode
}
