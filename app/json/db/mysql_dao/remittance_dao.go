package mysqldao

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"open.chat/app/json/db/dbo"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

const remittanceColumns = "id, chat_id, payer_uid, payee_uid, type, description, amount, status, create_date, created_at, updated_at"

type RemittanceDao struct {
	db *sqlx.DB
}

func NewRemittanceDao(db *sqlx.DB) *RemittanceDao {
	return &RemittanceDao{db: db}
}

func (dao *RemittanceDao) InsertTx(tx *sqlx.Tx, do *dbo.RemittanceDO) (lastInsertID, rowsAffected int64, err error) {
	var (
		query = "INSERT INTO `remittances` (chat_id, payer_uid, payee_uid, type, description, amount, status, create_date) VALUES" +
			"(:chat_id, :payer_uid, :payee_uid, :type, :description, :amount, :status, :create_date)"
		r sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		log.Errorf("namedExec in InsertTx(%v), error: %v", do, err)
		return
	}

	lastInsertID, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertID in InsertTx(%v), error: %v", do, err)
		return
	}

	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in InsertTx(%v), error: %v", do, err)
		return
	}

	return
}

func (dao *RemittanceDao) SelectByID(ctx context.Context, remittanceID uint32) (rValue *dbo.RemittanceDO, err error) {
	var (
		query = "SELECT " + remittanceColumns + " FROM remittances WHERE id=?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, remittanceID)

	if err != nil {
		log.Errorf("queryx in SelectByID(%d), error: %v", remittanceID, err)
		return
	}

	defer rows.Close()

	do := &dbo.RemittanceDO{}
	if rows.Next() {
		if err = rows.StructScan(do); err != nil {
			log.Errorf("structScan in SelectByID(%d), error: %v", remittanceID, err)
		} else {
			dao.fillDBO(do)

			rValue = do
		}
	}

	return
}

func (dao *RemittanceDao) UpdateRemittanceStatusTx(tx *sqlx.Tx, remittanceID uint32, status uint8) error {
	var (
		query = "UPDATE remittances SET status=? WHERE id=? and status=0"
	)

	r, err := tx.Exec(query, status, remittanceID)
	if err != nil {
		log.Errorf("Exec in UpdateRemittanceStatus(%d, %s), error: %v", remittanceID, status, err)
		return err
	}

	rA, err := r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected get in UpdateRemittanceStatus(%d, %s), error: %v", remittanceID, status, err)
		return err
	}

	if rA != 1 {
		log.Errorf("rowsAffected check in UpdateRemittanceStatus(%d, %s), rowsAffected: %d, error: %v", remittanceID, status, rA, err)
		return errors.New("update status failed")
	}

	return nil
}

func (dao *RemittanceDao) SelectByPayer(ctx context.Context, payerUID uint32, fromId uint32, limit uint32) ([]*dbo.RemittanceDO, error) {
	var (
		r   *sqlx.Rows
		err error
	)

	funcSign := fmt.Sprintf("SelectByPayer(%d, %d, %d)", payerUID, fromId, limit)

	if fromId == 0 {
		query := "select " + remittanceColumns + " FROM remittances where payer_uid=? and deleted=0 order by id desc limit ?"
		r, err = dao.db.Query(ctx, query, payerUID, limit)
	} else {
		query := "select " + remittanceColumns + " FROM remittances where payer_uid=? and deleted=0 and id<? order by id desc limit ?"
		r, err = dao.db.Query(ctx, query, payerUID, fromId, limit)
	}

	if err != nil {
		log.Errorf("queryx in SelectByPayer(%d, %d, %d), error: %v", payerUID, fromId, limit, err)
		return nil, err
	}

	out := dao.rowsToList(r, funcSign)
	return out, nil
}

func (dao *RemittanceDao) SelectByPayee(ctx context.Context, payeeUID uint32, fromId uint32, limit uint32) ([]*dbo.RemittanceDO, error) {
	var (
		r   *sqlx.Rows
		err error
	)

	funcSign := fmt.Sprintf("SelectByPayee(%d, %d, %d)", payeeUID, fromId, limit)

	if fromId == 0 {
		query := "select " + remittanceColumns + " FROM remittances where payee_uid=? and deleted=0 order by id desc limit ?"
		r, err = dao.db.Query(ctx, query, payeeUID, limit)
	} else {
		query := "select " + remittanceColumns + " FROM remittances where payee_uid=? and deleted=0 and id<? order by id desc limit ?"
		r, err = dao.db.Query(ctx, query, payeeUID, fromId, limit)
	}

	if err != nil {
		log.Errorf("queryx in %s, error: %v", funcSign, err)
		return nil, err
	}

	out := dao.rowsToList(r, funcSign)
	return out, nil
}

func (dao *RemittanceDao) rowsToList(r *sqlx.Rows, funcSign string) []*dbo.RemittanceDO {
	out := make([]*dbo.RemittanceDO, 0)
	for r.Next() {
		v := dbo.RemittanceDO{}
		if err := r.StructScan(&v); err != nil {
			log.Errorf("structScan in %s, error: %v", funcSign, err)
		} else {
			dao.fillDBO(&v)
			out = append(out, &v)
		}
	}
	return out
}

func (dao *RemittanceDao) fillDBO(o *dbo.RemittanceDO) {
	o.RemittedAt = uint32(o.CreateTime.Unix())
	switch o.Status {
	case 1:
		o.ReceivedAt = uint32(o.UpdateTime.Unix())
	case 2:
		o.RefundedAt = uint32(o.UpdateTime.Unix())
	}
}

func (dao *RemittanceDao) SelectTimeoutList(ctx context.Context, date int32) ([]*dbo.RemittanceDO, error) {
	rows, err := dao.db.Query(ctx, "SELECT "+remittanceColumns+" FROM remittances WHERE status = 0 AND create_date <= ?", date)
	if err != nil {
		log.Errorf("queryx in SelectTimeoutList(%d), error: %v", date, err)
		return []*dbo.RemittanceDO{}, err
	}
	defer rows.Close()
	return dao.rowsToList(rows, fmt.Sprintf("SelectTimeoutList(%d)", date)), err
}
