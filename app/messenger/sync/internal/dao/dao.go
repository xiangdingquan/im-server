package dao

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"open.chat/app/messenger/sync/internal/dal/dao/mysql_dao"
	"open.chat/app/messenger/sync/internal/dal/dataobject"
	idgen "open.chat/app/service/idgen/client"
	status_client "open.chat/app/service/status/client"
	"open.chat/mtproto"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/util"
)

const (
	seqUpdatesNgenId = "seq_updates_ngen_"
	botUpdatesNgenId = "bot_updates_ngen_"
)

const (
	PTS_UPDATE_TYPE_UNKNOWN = 0

	PTS_UPDATE_NEW_MESSAGE           = 1
	PTS_UPDATE_DELETE_MESSAGES       = 2
	PTS_UPDATE_READ_HISTORY_OUTBOX   = 3
	PTS_UPDATE_READ_HISTORY_INBOX    = 4
	PTS_UPDATE_WEBPAGE               = 5
	PTS_UPDATE_READ_MESSAGE_CONTENTS = 6
	PTS_UPDATE_EDIT_MESSAGE          = 7

	PTS_UPDATE_NEW_ENCRYPTED_MESSAGE = 8

	PTS_UPDATE_NEW_CHANNEL_MESSAGE     = 9
	PTS_UPDATE_DELETE_CHANNEL_MESSAGES = 10
	PTS_UPDATE_EDIT_CHANNEL_MESSAGE    = 11
	PTS_UPDATE_EDIT_CHANNEL_WEBPAGE    = 12

	PTS_UPDATE_NEW_BLOG     = 21
	PTS_UPDATE_DELETE_BLOG  = 22
	PTS_UPDATE_BLOG_FOLLOW  = 23
	PTS_UPDATE_BLOG_COMMENT = 24
	PTS_UPDATE_BLOG_LIKE    = 25
)

type Mysql struct {
	*sqlx.DB
	*mysql_dao.AuthSeqUpdatesDAO
	*mysql_dao.UserQtsUpdatesDAO
	*mysql_dao.UserPtsUpdatesDAO
	*mysql_dao.ChannelPtsUpdatesDAO
	*mysql_dao.BotUpdatesDAO
	*mysql_dao.BlogPtsUpdatesDAO
	*sqlx.CommonDAO
}

func newMysqlDao(c *sqlx.Config) *Mysql {
	db := sqlx.NewMySQL(c)
	return &Mysql{
		DB:                   db,
		AuthSeqUpdatesDAO:    mysql_dao.NewAuthSeqUpdatesDAO(db),
		UserQtsUpdatesDAO:    mysql_dao.NewUserQtsUpdatesDAO(db),
		UserPtsUpdatesDAO:    mysql_dao.NewUserPtsUpdatesDAO(db),
		ChannelPtsUpdatesDAO: mysql_dao.NewChannelPtsUpdatesDAO(db),
		BotUpdatesDAO:        mysql_dao.NewBotUpdatesDAO(db),
		BlogPtsUpdatesDAO:    mysql_dao.NewBlogPtsUpdatesDAO(db),
		CommonDAO:            sqlx.NewCommonDAO(db),
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

	// init status
	status_client.New()
	idgen.NewSeqIDGen()

	dao = &Dao{
		Mysql: newMysqlDao(dc.Mysql),
	}
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

func (d *Dao) AddQtsToUpdatesQueue(ctx context.Context, userId, qts, updateType int32, updateData []byte) int32 {
	do := &dataobject.UserQtsUpdatesDO{
		UserId:     userId,
		UpdateType: updateType,
		UpdateData: updateData,
		Date2:      int32(time.Now().Unix()),
		Qts:        qts,
	}

	i, _, _ := d.UserQtsUpdatesDAO.Insert(ctx, do)
	return int32(i)
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

func getUpdateType(update *mtproto.Update) int8 {
	switch update.PredicateName {
	case mtproto.Predicate_updateNewMessage:
		return PTS_UPDATE_NEW_MESSAGE
	case mtproto.Predicate_updateDeleteMessages:
		return PTS_UPDATE_DELETE_MESSAGES
	case mtproto.Predicate_updateReadHistoryOutbox:
		return PTS_UPDATE_READ_HISTORY_OUTBOX
	case mtproto.Predicate_updateReadHistoryInbox:
		return PTS_UPDATE_READ_HISTORY_INBOX
	case mtproto.Predicate_updateWebPage:
		return PTS_UPDATE_WEBPAGE
	case mtproto.Predicate_updateReadMessagesContents:
		return PTS_UPDATE_READ_MESSAGE_CONTENTS
	case mtproto.Predicate_updateEditMessage:
		return PTS_UPDATE_EDIT_MESSAGE

	case mtproto.Predicate_updateNewEncryptedMessage:
		return PTS_UPDATE_NEW_ENCRYPTED_MESSAGE

	case mtproto.Predicate_updateNewChannelMessage:
		return PTS_UPDATE_NEW_CHANNEL_MESSAGE
	case mtproto.Predicate_updateDeleteChannelMessages:
		return PTS_UPDATE_DELETE_CHANNEL_MESSAGES
	case mtproto.Predicate_updateEditChannelMessage:
		return PTS_UPDATE_EDIT_CHANNEL_MESSAGE
	case mtproto.Predicate_updateChannelWebPage:
		return PTS_UPDATE_EDIT_CHANNEL_WEBPAGE

	case mtproto.Predicate_updateNewBlog:
		return PTS_UPDATE_NEW_BLOG
	case mtproto.Predicate_updateDeleteBlog:
		return PTS_UPDATE_DELETE_BLOG
	case mtproto.Predicate_updateBlogFollow:
		return PTS_UPDATE_BLOG_FOLLOW
	case mtproto.Predicate_updateBlogComment:
		return PTS_UPDATE_BLOG_COMMENT
	case mtproto.Predicate_updateBlogLike:
		return PTS_UPDATE_BLOG_LIKE
	}
	return PTS_UPDATE_TYPE_UNKNOWN
}

func (d *Dao) AddToPtsQueue(ctx context.Context, userId, pts, ptsCount int32, update *mtproto.Update) int32 {
	updateData, _ := json.Marshal(update)

	do := &dataobject.UserPtsUpdatesDO{
		UserId:     userId,
		Pts:        pts,
		PtsCount:   ptsCount,
		UpdateType: getUpdateType(update),
		UpdateData: string(updateData),
		Date2:      int32(time.Now().Unix()),
	}

	i, _, _ := d.UserPtsUpdatesDAO.Insert(ctx, do)
	return int32(i)
}

func (d *Dao) AddToChannelPtsQueue(ctx context.Context, channelId, pts, ptsCount int32, update *mtproto.Update) int32 {
	updateData, _ := json.Marshal(update)

	do := &dataobject.ChannelPtsUpdatesDO{
		ChannelId:  channelId,
		Pts:        pts,
		PtsCount:   ptsCount,
		UpdateType: getUpdateType(update),
		UpdateData: string(updateData),
		Date2:      int32(time.Now().Unix()),
	}

	i, _, _ := d.ChannelPtsUpdatesDAO.Insert(ctx, do)
	return int32(i)
}

func (d *Dao) NextSeqId(ctx context.Context, key int64) (seq int64) {
	seq, _ = idgen.GetNextSeqID(ctx, seqUpdatesNgenId+util.Int64ToString(key))
	return
}

func (d *Dao) AddToBotUpdateQueue(ctx context.Context, botId int32, update *mtproto.Update) int32 {
	updateId := int32(d.NextUpdateId(ctx, botId))
	updateData, _ := json.Marshal(update)

	do := &dataobject.BotUpdatesDO{
		BotId:      botId,
		UpdateId:   updateId,
		UpdateType: getUpdateType(update),
		UpdateData: string(updateData),
		Date2:      time.Now().Unix(),
	}

	if _, _, err := d.BotUpdatesDAO.Insert(ctx, do); err != nil {
		_ = err
	}
	return updateId
}

func (d *Dao) NextUpdateId(ctx context.Context, key int32) (seq int64) {
	seq, _ = idgen.GetNextSeqID(ctx, botUpdatesNgenId+util.Int32ToString(key))
	return
}

func (d *Dao) AddToBlogPtsQueue(ctx context.Context, userId, pts, ptsCount int32, update *mtproto.Update) int32 {
	updateData, _ := json.Marshal(update)
	do := &dataobject.BlogPtsUpdatesDO{
		UserId:     userId,
		Pts:        pts,
		PtsCount:   ptsCount,
		UpdateType: getUpdateType(update),
		UpdateData: string(updateData),
		Date:       int32(time.Now().Unix()),
	}
	i, _, _ := d.BlogPtsUpdatesDAO.Insert(ctx, do)
	return int32(i)
}
