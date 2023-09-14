package mysqldao

import (
	"context"
	"database/sql"
	"open.chat/app/json/db/dbo"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

// RedPacketRecordsDAO .
type RedPacketRecordsDAO struct {
	db *sqlx.DB
}

// NewRedPacketRecordsDAO .
func NewRedPacketRecordsDAO(db *sqlx.DB) *RedPacketRecordsDAO {
	return &RedPacketRecordsDAO{db}
}

// InsertTx .
func (dao *RedPacketRecordsDAO) InsertTx(tx *sqlx.Tx, do *dbo.TRedPacketRecordDo) (lastInsertID, rowsAffected int64, err error) {
	var (
		query string = "INSERT INTO `red_packet_records`(red_packet_id, user_id, price) VALUES (:red_packet_id, :user_id, :price)"
		r     sql.Result
	)
	//添加红包记录
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

// Select .
func (dao *RedPacketRecordsDAO) Select(ctx context.Context, redpacketID, userID uint32) (rValue *dbo.TRedPacketRecordDo, err error) {
	var (
		query = "SELECT id, red_packet_id, user_id, price, created_at FROM `red_packet_records` WHERE red_packet_id = ? AND user_id = ? AND deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, redpacketID, userID)
	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dbo.TRedPacketRecordDo{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in Select(_), error: %v", err)
		}
		do.CreateAt = (uint32)(do.CreateTime.Unix())
		rValue = do
	}

	return
}

// SelectsByRid .
func (dao *RedPacketRecordsDAO) SelectsByRid(ctx context.Context, redpacketID uint32) (rList []dbo.TRedPacketRecordDo, err error) {
	var (
		query = "SELECT id, red_packet_id, user_id, price, created_at FROM `red_packet_records` WHERE red_packet_id = ? AND deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, redpacketID)

	if err != nil {
		log.Errorf("queryx in SelectsByRid(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dbo.TRedPacketRecordDo
	for rows.Next() {
		v := dbo.TRedPacketRecordDo{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectsByRid(_), error: %v", err)
		}
		v.CreateAt = (uint32)(v.CreateTime.Unix())
		values = append(values, v)
	}
	rList = values
	return
}

// SelectsByRid .
func (dao *RedPacketRecordsDAO) SelectsByRids(ctx context.Context, redPacketIDs []uint32) (rList []dbo.TRedPacketRecordDo, err error) {
	var (
		query = "SELECT id, red_packet_id, user_id, price, created_at FROM `red_packet_records` WHERE red_packet_id IN(?) AND deleted = 0"
		rows  *sqlx.Rows
	)

	query, args, err := sqlx.In(query, redPacketIDs)
	if err != nil {
		log.Error("sqlx.In in SelectsByRids(_), error: %v", err)
		return
	}

	rows, err = dao.db.Query(ctx, query, args...)

	if err != nil {
		log.Errorf("queryx in SelectsByRids(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dbo.TRedPacketRecordDo
	for rows.Next() {
		v := dbo.TRedPacketRecordDo{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectsByRids(_), error: %v", err)
		}
		v.CreateAt = (uint32)(v.CreateTime.Unix())
		values = append(values, v)
	}
	rList = values
	return
}

// SelectsByUid .
func (dao *RedPacketRecordsDAO) SelectsByUid(ctx context.Context, uID uint32) (rList []dbo.TRedPacketRecordDo, err error) {
	var (
		query = "SELECT id, red_packet_id, user_id, price, created_at FROM `red_packet_records` WHERE user_id = ? AND deleted = 0"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, uID)

	if err != nil {
		log.Errorf("queryx in SelectsByUid(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dbo.TRedPacketRecordDo
	for rows.Next() {
		v := dbo.TRedPacketRecordDo{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectsByUid(_), error: %v", err)
		}
		v.CreateAt = (uint32)(v.CreateTime.Unix())
		values = append(values, v)
	}
	rList = values
	return
}

func (dao *RedPacketRecordsDAO) Statistics(ctx context.Context, uID uint32, year uint32) (rValue *dbo.TRedPacketStatistics, err error) {
	rValue = &dbo.TRedPacketStatistics{}

	if year == 0 {
		query := `SELECT
					IFNULL(sum(price),0) total_price, count(*) cnt, count(if(t1.price=t2.mp,1,NULL)) top_price_count
				FROM 
					red_packet_records as t1
				JOIN
					(SELECT
						MAX(price) mp, red_packet_id 
					FROM
						red_packet_records
					WHERE
						red_packet_id in (SELECT red_packet_id from red_packet_records WHERE user_id=? AND deleted=0)
					GROUP BY
						red_packet_id)
					AS
						t2
				ON
					t1.red_packet_id = t2.red_packet_id
				WHERE
					user_id=? AND deleted=0`
		err = dao.db.Get(ctx, rValue, query, uID, uID)
	} else {
		query := `SELECT
					IFNULL(sum(price),0) total_price, count(*) cnt, count(if(t1.price=t2.mp,1,NULL)) top_price_count
				FROM 
					red_packet_records as t1
				JOIN
					(SELECT
						MAX(price) mp, red_packet_id 
					FROM
						red_packet_records
					WHERE
						red_packet_id in (select red_packet_id from red_packet_records where user_id=? AND deleted=0 AND year(created_at)=?)
					GROUP BY
						red_packet_id)
					AS
						t2
				ON
					t1.red_packet_id = t2.red_packet_id
				WHERE
					user_id=? AND deleted=0 AND year(created_at)=?`
		err = dao.db.Get(ctx, rValue, query, uID, year, uID, year)
	}
	return
}
