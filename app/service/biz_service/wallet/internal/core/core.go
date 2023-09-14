package core

import (
	"context"
	"database/sql"
	"time"

	"github.com/gogo/protobuf/types"
	"open.chat/app/service/biz_service/wallet/internal/dao"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type WalletCore struct {
	*dao.Dao
}

func New(dao *dao.Dao) *WalletCore {
	return &WalletCore{
		Dao: dao,
	}
}

func (m *WalletCore) GetInfo(ctx context.Context, userId int32) (*mtproto.Wallet_Info, error) {
	wdo, err := m.WalletDAO.SelectByUser(ctx, userId)
	if err != nil {
		return nil, mtproto.ErrInternelServerError
	}
	if wdo == nil {
		return nil, mtproto.ErrButtonTypeInvalid
	}
	info := mtproto.MakeTLWalletInfo(&mtproto.Wallet_Info{
		Address:     wdo.Address,
		Balance:     wdo.Balance,
		HasPassword: len(wdo.Password) > 0,
	}).To_Wallet_Info()
	return info, nil
}

func (m *WalletCore) recordReward(tx *sqlx.Tx, do interface{}) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into blog_rewards(user_id, target_uid, blog_id, amount, date) values (:user_id, :target_uid, :blog_id, :amount, :date)"
		r     sql.Result
	)
	r, err = tx.NamedExec(query, do)
	if err != nil {
		log.Errorf("namedExec in recordReward(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in recordReward(%v)_error: %v", do, err)
		return
	}

	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in recordReward(%v)_error: %v", do, err)
	}
	return
}

func (m *WalletCore) RewardBlog(ctx context.Context, fromUserId int32, userId int32, blogId int32, amount float64) (bool, error) {
	var (
		now = time.Now().Unix()
	)
	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		brDo := map[string]interface{}{
			"user_id":    fromUserId,
			"target_uid": userId,
			"blog_id":    blogId,
			"amount":     amount,
			"date":       now,
		}
		rid, _, err := m.recordReward(tx, brDo)
		if err != nil {
			result.Err = err
			return
		}
		result.Data = (uint32)(rid)

		//减少余额
		_, err = m.DecBalanceTx(tx, fromUserId, model.WalletRecordType_BlogReward, amount, blogId, "打赏用户")
		if err != nil {
			result.Err = err
			return
		}

		//增加余额
		_, err = m.IncBalanceTx(tx, fromUserId, model.WalletRecordType_BlogGotReward, amount, blogId, "获得打赏")
		if err != nil {
			result.Err = err
			return
		}
	})

	if tR.Err != nil {
		return false, tR.Err
	}
	return true, nil
}

func (m *WalletCore) GetAllRecords(ctx context.Context, fromUserId int32, date int32, offset int32, limit int32) (*mtproto.Wallet_Records, error) {
	records := mtproto.MakeTLWalletRecordsNotModified(nil).To_Wallet_Records()
	wrdos, err := m.WalletRecordDAO.SelectByUser(ctx, fromUserId, date, offset, limit)
	if err != nil {
		return records, mtproto.ErrInternelServerError
	}
	records = records.To_WalletRecords().To_Wallet_Records()
	records.Count = int32(m.WalletRecordDAO.SelectCountByUser(ctx, fromUserId, date))
	for _, wr := range wrdos {
		recordType := &model.WalletRecordTypeUtil{
			WalletRecordType: wr.Type,
		}
		switch wr.Type {
		case model.WalletRecordType_Modify:
		case model.WalletRecordType_Recharge, model.WalletRecordType_Withdrawal, model.WalletRecordType_WithdrawFail:
		case model.WalletRecordType_TransferIn, model.WalletRecordType_TransferOut:
		case model.WalletRecordType_CreateRedPacket, model.WalletRecordType_GetRedPacket, model.WalletRecordType_GivebackRedPacket:
		case model.WalletRecordType_BlogReward, model.WalletRecordType_BlogGotReward:
		case model.WalletRecordType_RemitRemittance, model.WalletRecordType_ReceiveRemittance,
			model.WalletRecordType_RefundRemittanceByUser, model.WalletRecordType_RefundRemittanceBySystem:
		}
		record := mtproto.MakeTLWalletRecord(&mtproto.Wallet_Record{
			Id:     wr.ID,
			Type:   recordType.ToWalletRecordType(),
			Amount: wr.Amount,
			Date:   wr.Date,
		}).To_Wallet_Record()
		if len(wr.Remarks) > 0 {
			record.Remark = &types.StringValue{
				Value: wr.Remarks,
			}
		}
		records.Records = append(records.Records, record)
	}
	return records, nil
}

func (m *WalletCore) GetRecordsByType(ctx context.Context, fromUserId int32, recordType int8, date int32, offset int32, limit int32) (*mtproto.Wallet_Records, error) {
	records := mtproto.MakeTLWalletRecordsNotModified(nil).To_Wallet_Records()
	wrdos, err := m.WalletRecordDAO.SelectByType(ctx, fromUserId, recordType, date, offset, limit)
	if err != nil {
		return records, mtproto.ErrInternelServerError
	}
	records = records.To_WalletRecords().To_Wallet_Records()
	records.Count = int32(m.WalletRecordDAO.SelectCountByType(ctx, fromUserId, recordType, date))
	for _, wr := range wrdos {
		recordType := &model.WalletRecordTypeUtil{
			WalletRecordType: wr.Type,
		}
		switch wr.Type {
		case model.WalletRecordType_Modify:
		case model.WalletRecordType_Recharge, model.WalletRecordType_Withdrawal, model.WalletRecordType_WithdrawFail:
		case model.WalletRecordType_TransferIn, model.WalletRecordType_TransferOut:
		case model.WalletRecordType_CreateRedPacket, model.WalletRecordType_GetRedPacket, model.WalletRecordType_GivebackRedPacket:
		case model.WalletRecordType_BlogReward, model.WalletRecordType_BlogGotReward:
		case model.WalletRecordType_RemitRemittance, model.WalletRecordType_ReceiveRemittance,
			model.WalletRecordType_RefundRemittanceByUser, model.WalletRecordType_RefundRemittanceBySystem:
		}
		record := mtproto.MakeTLWalletRecord(&mtproto.Wallet_Record{
			Id:     wr.ID,
			Type:   recordType.ToWalletRecordType(),
			Amount: wr.Amount,
			Date:   wr.Date,
		}).To_Wallet_Record()
		if len(wr.Remarks) > 0 {
			record.Remark = &types.StringValue{
				Value: wr.Remarks,
			}
		}
		records.Records = append(records.Records, record)
	}
	return records, nil
}
