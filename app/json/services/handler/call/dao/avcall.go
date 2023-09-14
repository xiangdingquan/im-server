package dao

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"open.chat/app/infra/databus/pkg/cache/redis"
	"open.chat/app/json/db/dbo"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

// TAvCallAndRecord .
type TAvCallAndRecord struct {
	CallID      uint32    `db:"call_id" json:"callId"`
	ChannelName string    `db:"channel_name" json:"channelName"` //通道名
	ChatID      uint32    `db:"chat_id" json:"chatId"`           //ChatId
	From        uint32    `db:"owner_uid" json:"from"`           //创建者
	Members     []uint32  `db:"-" json:"to"`                     //成员
	CreateTime  time.Time `db:"created_at" json:"-"`             //创建通话时间
	CreateAt    uint32    `db:"-" json:"createAt"`               //创建通话时间
	CloseAt     uint32    `db:"close_at" json:"closeAt"`         //是否已取消通话
	IsMeetingAV bool      `db:"is_meet" json:"isMeetingAV"`      //会话是否为多人
	IsVideo     bool      `db:"is_video" json:"isVideo"`         //是否开启视频
	EnterAt     uint32    `db:"enter_at" json:"enterAt"`         //开始通话时间
	LeaveAt     uint32    `db:"leave_at" json:"leaveAt"`         //结束通话时间
	MemberInfo  string    `db:"member_uids" json:"-"`            //成员
	//IsTimeOut   bool      `db:"-" json:"isTimeOut"`              //是否已超时
}

const (
	avCallTimeout     = 30 // salt timeout
	cacheavCallPrefix = "avcall_invite"
)

func genCacheKey(uid uint32) string {
	return fmt.Sprintf("%s_%d", cacheavCallPrefix, uid)
}

// Create .
func (d *Dao) Create(ctx context.Context, channelName string, chatID, ownerUID uint32, to []uint32, isMeetbool, isVideo bool) (uint32, error) {
	members, err := json.Marshal(to)
	if err != nil {
		return 0, nil
	}
	tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		callDO := &dbo.AvcallDO{
			ChannelName: channelName,
			ChatID:      chatID,
			OwnerUID:    ownerUID,
			MemberInfo:  string(members),
			//StartAt:    uint32(time.Now().Unix()),
			IsVideo: isVideo,
			IsMeet:  isMeetbool,
		}
		callID, _, err := d.AvcallsDAO.InsertTx(tx, callDO)
		if err != nil {
			result.Err = err
			return
		}

		recordDO := &dbo.AvcallRecordDO{
			CallID:  uint32(callID),
			UserID:  ownerUID,
			IsRead:  true,
			EnterAt: uint32(time.Now().Unix()),
		}
		_, _, err = d.AvcallsRecordsDAO.InsertTx(tx, recordDO)
		if err != nil {
			result.Err = err
			return
		}

		for _, uid := range to {
			err = d.PutCacheAvCall(ctx, uid, uint32(callID))
			if err != nil {
				result.Err = err
				return
			}
		}
		result.Data = (uint32)(callID)
	})

	if tR.Err != nil {
		return 0, tR.Err
	}

	return tR.Data.(uint32), nil
}

func (d *Dao) Cancel(ctx context.Context, cdo *dbo.AvcallDO) error {
	tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		cdo.CloseAt = uint32(time.Now().Unix())
		rowsAffected, err := d.AvcallsDAO.UpdateWithID(ctx, map[string]interface{}{
			"close_at": cdo.CloseAt,
		}, cdo.ID)
		if rowsAffected == 0 {
			result.Err = err
			return
		}

		for _, uid := range cdo.Members {
			d.DelCacheAvCall(ctx, uid)
		}

		result.Data = cdo.ID
	})

	return tR.Err
}

func (d *Dao) Start(ctx context.Context, uid uint32, cdo *dbo.AvcallDO) error {
	tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		if cdo.StartAt == 0 {
			cdo.StartAt = uint32(time.Now().Unix())
			rowsAffected, err := d.AvcallsDAO.UpdateWithIDTx(tx, map[string]interface{}{
				"start_at": cdo.StartAt,
			}, cdo.ID)
			if rowsAffected == 0 {
				result.Err = err
				return
			}
			result.Data = (uint32)(cdo.ID)
		}

		recordDO := &dbo.AvcallRecordDO{
			CallID:  cdo.ID,
			UserID:  uid,
			IsRead:  true,
			EnterAt: cdo.StartAt,
		}
		rID, _, err := d.AvcallsRecordsDAO.InsertTx(tx, recordDO)
		if err != nil {
			result.Err = err
			return
		}

		d.DelCacheAvCall(ctx, uid)

		result.Data = (uint32)(rID)
	})

	return tR.Err
}

func (d *Dao) Stop(ctx context.Context, uid uint32, cdo *dbo.AvcallDO) error {
	tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		now := uint32(time.Now().Unix())
		if !cdo.IsMeet || uid == cdo.OwnerUID { //通话结束
			cdo.CloseAt = now
			rowsAffected, err := d.AvcallsDAO.UpdateWithIDTx(tx, map[string]interface{}{
				"close_at": now,
			}, cdo.ID)
			if rowsAffected == 0 {
				result.Err = err
				return
			}
			result.Data = (uint32)(cdo.ID)

			_, err = d.AvcallsRecordsDAO.UpdateWithCallIDTx(tx, map[string]interface{}{
				"leave_at": now,
			}, cdo.ID)
			if err != nil {
				result.Err = err
				return
			}
		} else { //成员离开
			rowsAffected, err := d.AvcallsRecordsDAO.UpdateWithCallAndUserTx(tx, map[string]interface{}{
				"leave_at": now,
			}, cdo.ID, uid)
			if rowsAffected == 0 {
				result.Err = err
				return
			}
		}
	})

	return tR.Err
}

// QueryMeRecords . type 0:单聊呼出和接听 1单聊呼出 2单聊接听 3会议
func (d *Dao) QueryMeRecords(ctx context.Context, uID uint32, _type int8, count, page uint32) (rList []TAvCallAndRecord, err error) {
	if page == 0 {
		page = 1
	}
	var startPs = count * (page - 1)
	query := "SELECT l.call_id, r.channel_name, r.chat_id, r.owner_uid, IFNULL(r.`member_uids`,'[]') member_uids, r.created_at, r.close_at, r.is_meet, r.is_video, l.enter_at, l.leave_at FROM `avcall_records` l LEFT JOIN `avcalls` r ON(r.id = l.call_id) WHERE l.user_id = ? AND l.deleted = 0"
	switch _type {
	case 1:
		query += " AND r.owner_uid = l.user_id"
	case 2:
		query += " AND r.owner_uid <> l.user_id"
	case 3:
		query += " AND r.is_meet = 1"
	default:
	}
	query += " ORDER BY l.id DESC LIMIT ?,?"
	var rows *sqlx.Rows
	rows, err = d.Query(ctx, query, uID, startPs, count)

	if err != nil {
		log.Errorf("queryx in QueryMeRecords(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []TAvCallAndRecord
	for rows.Next() {
		v := TAvCallAndRecord{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in QueryMeRecords(_), error: %v", err)
		}
		v.CreateAt = (uint32)(v.CreateTime.Unix())
		json.Unmarshal([]byte(v.MemberInfo), &v.Members)
		values = append(values, v)
	}
	rList = values
	return
}

// QueryMeOffline .
func (d *Dao) QueryMeOffline(ctx context.Context, uID uint32, timeOut uint32) (rList []TAvCallAndRecord, err error) {
	avCallID, err := d.GetCacheAvCall(ctx, uID)
	if avCallID == 0 {
		return make([]TAvCallAndRecord, 0), err
	}

	cdo, err := d.AvcallsDAO.SelectByCallID(ctx, avCallID)
	if cdo == nil {
		return make([]TAvCallAndRecord, 0), err
	}

	rList = []TAvCallAndRecord{
		{
			CallID:      cdo.ID,
			ChannelName: cdo.ChannelName,
			ChatID:      cdo.ChatID,
			CreateTime:  cdo.CreateTime,
			CreateAt:    cdo.CreateAt,
			CloseAt:     cdo.CloseAt,
			IsMeetingAV: cdo.IsMeet,
			IsVideo:     cdo.IsVideo,
			EnterAt:     0,
			LeaveAt:     cdo.CloseAt,
			From:        cdo.OwnerUID,
			MemberInfo:  cdo.MemberInfo,
			Members:     cdo.Members,
		},
	}
	return
}

func (d *Redis) PutCacheAvCall(ctx context.Context, uid uint32, avCallID uint32) (err error) {
	cacheKey := genCacheKey(uid)
	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()
	if _, err = conn.Do("SETEX", cacheKey, avCallTimeout, avCallID); err != nil {
		log.Errorf("conn.SETEX(%s) error(%v)", cacheKey, err)
	}
	return
}

func (d *Redis) GetCacheAvCall(ctx context.Context, uid uint32) (avCallID uint32, err error) {
	cacheKey := genCacheKey(uid)
	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()
	v, err := redis.Uint64(conn.Do("GET", cacheKey))
	if err != nil {
		if err != redis.ErrNil {
			log.Errorf("conn.Do(GET %s) error(%v)", cacheKey, err)
		} else {
			err = nil
		}
	} else {
		avCallID = uint32(v)
	}
	return
}

func (d *Redis) DelCacheAvCall(ctx context.Context, uid uint32) (err error) {
	cacheKey := genCacheKey(uid)
	conn := d.redis.Redis.Get(ctx)
	defer conn.Close()
	if _, err = conn.Do("DEL", cacheKey); err != nil {
		log.Errorf("conn.DEL(%s) error(%v)", cacheKey, err)
	}
	return
}
