package mysqldao

import (
	"context"
	"database/sql"
	"errors"

	"open.chat/app/json/db/dbo"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

// RedPacketDAO .
type RedPacketDAO struct {
	db *sqlx.DB
}

// NewAvcallsDAO .
func NewRedPacketDAO(db *sqlx.DB) *RedPacketDAO {
	return &RedPacketDAO{db}
}

// InsertTx .
func (dao *RedPacketDAO) InsertTx(tx *sqlx.Tx, do *dbo.RedPacketDO) (lastInsertID, rowsAffected int64, err error) {
	var (
		query string = "INSERT INTO `red_packets`(chat_id, owner_uid, type, title, price, total_price, total_count, remain_price, remain_count, create_date) VALUES (:chat_id, :owner_uid, :type, :title, :price, :total_price, :total_count, :remain_price, :remain_count, :create_date)"
		r     sql.Result
	)

	//生成红包
	r, err = tx.NamedExec(query, do)
	if err != nil {
		log.Errorf("namedExec in InsertTx(%v), error: %v", do, err)
		return
	}

	lastInsertID, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in InsertTx(%v)_error: %v", do, err)
		return
	}

	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertTx(%v)_error: %v", do, err)
	}

	return
}

// SelectByID .
func (dao *RedPacketDAO) SelectByID(ctx context.Context, redpacketID uint32) (rValue *dbo.RedPacketDO, err error) {
	var (
		query = "SELECT id, chat_id, owner_uid, type, title, price, total_price, total_count, remain_price, remain_count, create_date, completed, created_at FROM `red_packets` WHERE id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, redpacketID)

	if err != nil {
		log.Errorf("queryx in SelectByCallID(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dbo.RedPacketDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByCallID(_), error: %v", err)
		} else {
			do.CreateAt = (uint32)(do.CreateTime.Unix())
			rValue = do
		}
	}

	return
}

// DecreaseOne .减少一个红包
func (dao *RedPacketDAO) DecreaseOne(ctx context.Context, redPacketID uint32, price float64) (rowsAffected int64, err error) {
	var (
		query string = "UPDATE `red_packets` SET remain_price = remain_price - ?,remain_count = remain_count - 1, completed=IF(remain_count=0, 1, completed) WHERE `id` = ? AND `completed` = 0"
		r     sql.Result
	)

	r, err = dao.db.Exec(ctx, query, price, redPacketID)
	if err != nil {
		log.Errorf("Exec in DecreaseOne()_error: %v", err)
		return
	}

	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DecreaseOne()_error: %v", err)
		return
	}

	if rowsAffected < 1 {
		err = errors.New("decrease a redpacket fail")
	}

	return
}

// DecreaseOneTx .减少一个红包
func (dao *RedPacketDAO) DecreaseOneTx(tx *sqlx.Tx, redPacketID uint32, price float64) (rowsAffected int64, err error) {
	var (
		query string = "UPDATE `red_packets` SET remain_price = remain_price - ?,remain_count = remain_count - 1, completed=IF(remain_count=0, 1, completed) WHERE `id` = ? AND `completed` = 0"
		r     sql.Result
	)

	r, err = tx.Exec(query, price, redPacketID)
	if err != nil {
		log.Errorf("Exec in DecreaseOne()_error: %v", err)
		return
	}

	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DecreaseOne()_error: %v", err)
		return
	}

	if rowsAffected < 1 {
		err = errors.New("decrease a redpacket fail")
	}

	return
}

// SetCompleted .设置完成
func (dao *RedPacketDAO) SetCompleted(ctx context.Context, redPacketID uint32) (rowsAffected int64, err error) {
	var (
		query string = "UPDATE `red_packets` SET completed = 1 WHERE id = ? AND completed = 0"
		r     sql.Result
	)

	r, err = dao.db.Exec(ctx, query, redPacketID)
	if err != nil {
		log.Errorf("Exec in SetCompleted()_error: %v", err)
		return
	}

	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in SetCompleted()_error: %v", err)
		return
	}

	return
}

// SetCompletedTx .设置完成
func (dao *RedPacketDAO) SetCompletedTx(tx *sqlx.Tx, redPacketID uint32) (rowsAffected int64, err error) {
	var (
		query string = "UPDATE `red_packets` SET completed = 1 WHERE id = ? AND completed = 0"
		r     sql.Result
	)

	r, err = tx.Exec(query, redPacketID)
	if err != nil {
		log.Errorf("Exec in SetCompletedTx()_error: %v", err)
		return
	}

	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in SetCompletedTx()_error: %v", err)
		return
	}

	return
}

// SelectTimeoutList .查找过期的红包
func (dao *RedPacketDAO) SelectTimeoutList(ctx context.Context, date int32) (rValue []*dbo.RedPacketDO, err error) {
	var (
		query = "SELECT id, chat_id, owner_uid, type, title, price, total_price, total_count, remain_price, remain_count, create_date, completed, created_at FROM `red_packets` WHERE completed = 0 AND create_date <= ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, date)

	if err != nil {
		log.Errorf("queryx in SelectTimeoutList(_), error: %v", err)
		return
	}

	defer rows.Close()

	rValue = make([]*dbo.RedPacketDO, 0)
	for rows.Next() {
		do := &dbo.RedPacketDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectTimeoutList(_), error: %v", err)
		} else {
			do.CreateAt = (uint32)(do.CreateTime.Unix())
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *RedPacketDAO) Statistics(ctx context.Context, uID uint32, year uint32) (rValue *dbo.TRedPacketStatistics, err error) {
	rValue = &dbo.TRedPacketStatistics{}

	if year == 0 {
		query := "select IFNULL(sum(total_price),0) total_price, count(*) cnt, 0 top_price_count from `red_packets` where owner_uid=? and deleted=0"
		err = dao.db.Get(ctx, rValue, query, uID)
	} else {
		query := "select IFNULL(sum(total_price),0) total_price, count(*) cnt, 0 top_price_count from `red_packets` where owner_uid=? and deleted=0 and year(created_at)=?"
		err = dao.db.Get(ctx, rValue, query, uID, year)
	}
	return
}
