package dao

import (
	"context"
	"time"

	"open.chat/app/service/biz_service/wallet/internal/dal/dataobject"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/pkg/database/sqlx"
)

// Dao dao.
type Dao struct {
	*Mysql
	*Redis
}

// New new a dao and return.
func New() (dao *Dao) {
	dao = &Dao{
		Mysql: newMysqlDao(),
		Redis: newRedisDao(),
	}

	idgen.NewUUID()
	idgen.NewSeqIDGen()

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

func (d *Dao) IncBalance(ctx context.Context, uid int32, type_ int8, money float64, related int32, remarks string) (int32, error) {
	tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		//增加余额
		_, err := d.WalletDAO.IncreaseBalanceTx(tx, uid, money)
		if err != nil {
			result.Err = err
			return
		}
		//增加日志
		wrDo := &dataobject.WalletRecordDO{
			UID:     uid,
			Type:    int8(type_),
			Amount:  money,
			Related: related,
			Remarks: remarks,
			Date:    int32(time.Now().Unix()),
		}
		rid, _, err := d.WalletRecordDAO.InsertTx(tx, wrDo)
		if err != nil {
			result.Err = err
			return
		}

		result.Data = rid
	})

	if tR.Err != nil {
		return 0, tR.Err
	}

	return int32(tR.Data.(int64)), tR.Err
}

func (d *Dao) IncBalanceTx(tx *sqlx.Tx, uid int32, type_ int8, money float64, related int32, remarks string) (int32, error) {
	//增加余额
	_, err := d.WalletDAO.IncreaseBalanceTx(tx, uid, money)
	if err != nil {
		return 0, err
	}
	//增加日志
	wrDo := &dataobject.WalletRecordDO{
		UID:     uid,
		Type:    int8(type_),
		Amount:  money,
		Related: related,
		Remarks: remarks,
		Date:    int32(time.Now().Unix()),
	}
	rid, _, err := d.WalletRecordDAO.InsertTx(tx, wrDo)
	if err != nil {
		return 0, err
	}

	return (int32)(rid), err
}

func (d *Dao) DecBalanceTx(tx *sqlx.Tx, uid int32, type_ int8, money float64, related int32, remarks string) (int32, error) {
	return d.IncBalanceTx(tx, uid, type_, -money, related, remarks)
}
