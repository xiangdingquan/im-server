package mysql_dao

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"open.chat/app/service/biz_service/channel/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type ChannelsDAO struct {
	db *sqlx.DB
}

func NewChannelsDAO(db *sqlx.DB) *ChannelsDAO {
	return &ChannelsDAO{db}
}

func (dao *ChannelsDAO) Insert(ctx context.Context, do *dataobject.ChannelsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into channels(creator_user_id, access_hash, secret_key_id, random_id, top_message, read_outbox_max_id, date2, pts, participants_count, admins_count, title, about, link, broadcast, megagroup, democracy, signatures, migrated_from_chat_id, `date`) values (:creator_user_id, :access_hash, :secret_key_id, :random_id, :top_message, :read_outbox_max_id, :date2, :pts, :participants_count, :admins_count, :title, :about, :link, :broadcast, :megagroup, :democracy, :signatures, :migrated_from_chat_id, :date)"
		r     sql.Result
	)

	r, err = dao.db.NamedExec(ctx, query, do)
	if err != nil {
		log.Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

func (dao *ChannelsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.ChannelsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into channels(creator_user_id, access_hash, secret_key_id, random_id, top_message, read_outbox_max_id, date2, pts, participants_count, admins_count, title, about, link, `lat`, `long`, `accuracy_radius`, `address`, `has_link`, `has_geo`, slowmode_enabled, broadcast, megagroup, democracy, signatures, migrated_from_chat_id, `date`) values (:creator_user_id, :access_hash, :secret_key_id, :random_id, :top_message, :read_outbox_max_id, :date2, :pts, :participants_count, :admins_count, :title, :about, :link, :lat, :long, :accuracy_radius, :address, :has_link, :has_geo, :slowmode_enabled, :broadcast, :megagroup, :democracy, :signatures, :migrated_from_chat_id, :date)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		log.Errorf("namedExec in Insert(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
	if err != nil {
		log.Errorf("lastInsertId in Insert(%v)_error: %v", do, err)
		return
	}
	rowsAffected, err = r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Insert(%v)_error: %v", do, err)
	}

	return
}

func (dao *ChannelsDAO) SelectByOffsetDate(ctx context.Context, idList []int32, date2 int32, limit int32) (rList []dataobject.ChannelsDO, err error) {
	var (
		query = "select id, creator_user_id, access_hash, secret_key_id, random_id, top_message, pinned_msg_id, read_outbox_max_id, date2, pts, participants_count, admins_count, kicked_count, banned_count, title, about, photo, public, username, link, `lat`, `long`, `accuracy_radius`, `address`, `has_geo`, `slowmode_enabled`, `slowmode_seconds`, broadcast, megagroup, democracy, signatures, admins_enabled, default_banned_rights, migrated_from_chat_id, pre_history_hidden, has_link, linked_chat_id, deactivated, version, deleted, `date` from channels where id in (?) and deleted = 0 and date2 < ? order by date2 desc limit ?"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, idList, date2, limit)
	if err != nil {
		log.Errorf("sqlx.In in SelectByOffsetDate(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectByOffsetDate(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChannelsDO
	for rows.Next() {
		v := dataobject.ChannelsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByOffsetDate(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ChannelsDAO) Select(ctx context.Context, id int32) (rValue *dataobject.ChannelsDO, err error) {
	var (
		query = "select id, creator_user_id, access_hash, secret_key_id, random_id, top_message, pinned_msg_id, read_outbox_max_id, date2, pts, participants_count, admins_count, kicked_count, banned_count, title, about, photo, public, username, link, `lat`, `long`, `accuracy_radius`, `address`, `has_geo`, `slowmode_enabled`, `slowmode_seconds`, broadcast, megagroup, democracy, signatures, admins_enabled, default_banned_rights, migrated_from_chat_id, pre_history_hidden, has_link, linked_chat_id, deactivated, version, deleted, `date` from channels where id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, id)

	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.ChannelsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in Select(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *ChannelsDAO) SelectLastCreator(ctx context.Context, creator_user_id int32) (rValue *dataobject.ChannelsDO, err error) {
	var (
		query = "select id, creator_user_id, access_hash, secret_key_id, random_id, top_message, pinned_msg_id, read_outbox_max_id, date2, pts, participants_count, admins_count, kicked_count, banned_count, title, about, photo, public, username, link, `lat`, `long`, `accuracy_radius`, `address`, `has_geo`, `slowmode_enabled`, `slowmode_seconds`, broadcast, megagroup, democracy, signatures, admins_enabled, default_banned_rights, migrated_from_chat_id, pre_history_hidden, has_link, linked_chat_id, deactivated, version, deleted, `date` from channels where creator_user_id = ? order by `date` desc limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, creator_user_id)

	if err != nil {
		log.Errorf("queryx in SelectLastCreator(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.ChannelsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectLastCreator(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *ChannelsDAO) SelectByTitle(ctx context.Context, title string) (rList []int32, err error) {
	var query = "select id from channels where title = ? and deleted = 0 and username != '' limit 5"
	err = dao.db.Select(ctx, &rList, query, title)

	if err != nil {
		log.Errorf("select in SelectByTitle(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateTitle(ctx context.Context, title string, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set title = ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, title, date, id)

	if err != nil {
		log.Errorf("exec in UpdateTitle(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateTitle(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateTitleTx(tx *sqlx.Tx, title string, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set title = ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, title, date, id)

	if err != nil {
		log.Errorf("exec in UpdateTitle(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateTitle(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateAbout(ctx context.Context, about string, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set about = ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, about, date, id)

	if err != nil {
		log.Errorf("exec in UpdateAbout(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateAbout(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateAboutTx(tx *sqlx.Tx, about string, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set about = ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, about, date, id)

	if err != nil {
		log.Errorf("exec in UpdateAbout(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateAbout(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateNotice(ctx context.Context, notice string, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set notice = ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, notice, date, id)

	if err != nil {
		log.Errorf("exec in UpdateNotice(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateNotice(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateNoticeTx(tx *sqlx.Tx, notice string, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set notice = ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, notice, date, id)

	if err != nil {
		log.Errorf("exec in UpdateNotice(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateNotice(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateLink(ctx context.Context, link string, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set link = ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, link, date, id)

	if err != nil {
		log.Errorf("exec in UpdateLink(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateLink(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateLinkTx(tx *sqlx.Tx, link string, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set link = ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, link, date, id)

	if err != nil {
		log.Errorf("exec in UpdateLink(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateLink(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) SelectByLink(ctx context.Context, link string) (rValue *dataobject.ChannelsDO, err error) {
	var (
		query = "select id, creator_user_id, access_hash, secret_key_id, random_id, top_message, pinned_msg_id, read_outbox_max_id, date2, pts, participants_count, admins_count, kicked_count, banned_count, title, about, photo, public, username, link, `lat`, `long`, `accuracy_radius`, `address`, `has_geo`, `slowmode_enabled`, `slowmode_seconds`, broadcast, megagroup, democracy, signatures, admins_enabled, default_banned_rights, migrated_from_chat_id, pre_history_hidden, has_link, linked_chat_id, deactivated, version, deleted, `date` from channels where deactivated = 0 and link = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, link)

	if err != nil {
		log.Errorf("queryx in SelectByLink(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.ChannelsDO{}
	if rows.Next() {
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectByLink(_), error: %v", err)
			return
		} else {
			rValue = do
		}
	}

	return
}

func (dao *ChannelsDAO) SelectByIdList(ctx context.Context, idList []int32) (rList []dataobject.ChannelsDO, err error) {
	var (
		query = "select id, creator_user_id, access_hash, secret_key_id, random_id, top_message, pinned_msg_id, read_outbox_max_id, date2, pts, participants_count, admins_count, kicked_count, banned_count, title, about, photo, public, username, link, `lat`, `long`, `accuracy_radius`, `address`, `has_geo`, `slowmode_enabled`, `slowmode_seconds`, broadcast, megagroup, democracy, signatures, admins_enabled, default_banned_rights, migrated_from_chat_id, pre_history_hidden, has_link, linked_chat_id, deactivated, version, deleted, `date` from channels where deleted = 0 and id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	query, a, err = sqlx.In(query, idList)
	if err != nil {
		log.Errorf("sqlx.In in SelectByIdList(_), error: %v", err)
		return
	}
	rows, err = dao.db.Query(ctx, query, a...)

	if err != nil {
		log.Errorf("queryx in SelectByIdList(_), error: %v", err)
		return
	}

	defer rows.Close()

	var values []dataobject.ChannelsDO
	for rows.Next() {
		v := dataobject.ChannelsDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByIdList(_), error: %v", err)
			return
		}
		values = append(values, v)
	}
	rList = values

	return
}

func (dao *ChannelsDAO) UpdateParticipantCount(ctx context.Context, participants_count int32, admins_count int32, kicked_count int32, banned_count int32, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set participants_count = participants_count + ?, admins_count = admins_count + ?, kicked_count = kicked_count + ?, banned_count = banned_count + ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, participants_count, admins_count, kicked_count, banned_count, date, id)

	if err != nil {
		log.Errorf("exec in UpdateParticipantCount(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateParticipantCount(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateParticipantCountTx(tx *sqlx.Tx, participants_count int32, admins_count int32, kicked_count int32, banned_count int32, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set participants_count = participants_count + ?, admins_count = admins_count + ?, kicked_count = kicked_count + ?, banned_count = banned_count + ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, participants_count, admins_count, kicked_count, banned_count, date, id)

	if err != nil {
		log.Errorf("exec in UpdateParticipantCount(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateParticipantCount(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdatePhoto(ctx context.Context, photo string, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set photo = ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, photo, date, id)

	if err != nil {
		log.Errorf("exec in UpdatePhoto(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdatePhoto(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdatePhotoTx(tx *sqlx.Tx, photo string, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set photo = ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, photo, date, id)

	if err != nil {
		log.Errorf("exec in UpdatePhoto(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdatePhoto(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateTopMessage(ctx context.Context, top_message int32, pts int32, date2 int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set top_message = ?, pts = ?, date2 = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, top_message, pts, date2, id)

	if err != nil {
		log.Errorf("exec in UpdateTopMessage(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateTopMessage(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateTopMessageTx(tx *sqlx.Tx, top_message int32, pts int32, date2 int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set top_message = ?, pts = ?, date2 = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, top_message, pts, date2, id)

	if err != nil {
		log.Errorf("exec in UpdateTopMessage(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateTopMessage(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateAdminsEnabled(ctx context.Context, admins_enabled int8, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set admins_enabled = ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, admins_enabled, date, id)

	if err != nil {
		log.Errorf("exec in UpdateAdminsEnabled(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateAdminsEnabled(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateAdminsEnabledTx(tx *sqlx.Tx, admins_enabled int8, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set admins_enabled = ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, admins_enabled, date, id)

	if err != nil {
		log.Errorf("exec in UpdateAdminsEnabled(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateAdminsEnabled(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateDefaultBannedRights(ctx context.Context, default_banned_rights int32, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set default_banned_rights = ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, default_banned_rights, date, id)

	if err != nil {
		log.Errorf("exec in UpdateDefaultBannedRights(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateDefaultBannedRights(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateDefaultBannedRightsTx(tx *sqlx.Tx, default_banned_rights int32, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set default_banned_rights = ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, default_banned_rights, date, id)

	if err != nil {
		log.Errorf("exec in UpdateDefaultBannedRights(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateDefaultBannedRights(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateVersion(ctx context.Context, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, date, id)

	if err != nil {
		log.Errorf("exec in UpdateVersion(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateVersion(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateVersionTx(tx *sqlx.Tx, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, date, id)

	if err != nil {
		log.Errorf("exec in UpdateVersion(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateVersion(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateDemocracy(ctx context.Context, democracy int8, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set democracy = ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, democracy, date, id)

	if err != nil {
		log.Errorf("exec in UpdateDemocracy(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateDemocracy(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateDemocracyTx(tx *sqlx.Tx, democracy int8, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set democracy = ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, democracy, date, id)

	if err != nil {
		log.Errorf("exec in UpdateDemocracy(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateDemocracy(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateSignatures(ctx context.Context, signatures int8, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set signatures = ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, signatures, date, id)

	if err != nil {
		log.Errorf("exec in UpdateSignatures(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateSignatures(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateSignaturesTx(tx *sqlx.Tx, signatures int8, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set signatures = ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, signatures, date, id)

	if err != nil {
		log.Errorf("exec in UpdateSignatures(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateSignatures(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateUsername(ctx context.Context, username string, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set username = ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, username, date, id)

	if err != nil {
		log.Errorf("exec in UpdateUsername(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateUsername(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateUsernameTx(tx *sqlx.Tx, username string, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set username = ?, `date` = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, username, date, id)

	if err != nil {
		log.Errorf("exec in UpdateUsername(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateUsername(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateLinkedChatId(ctx context.Context, linked_chat_id int32, has_link int8, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set linked_chat_id = ?, has_link = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, linked_chat_id, has_link, id)

	if err != nil {
		log.Errorf("exec in UpdateLinkedChatId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateLinkedChatId(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateLinkedChatIdTx(tx *sqlx.Tx, linked_chat_id int32, has_link int8, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set linked_chat_id = ?, has_link = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, linked_chat_id, has_link, id)

	if err != nil {
		log.Errorf("exec in UpdateLinkedChatId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateLinkedChatId(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) ClearLinkedChatId(ctx context.Context, linked_chat_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set linked_chat_id = 0, has_link = 0 where linked_chat_id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, linked_chat_id)

	if err != nil {
		log.Errorf("exec in ClearLinkedChatId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in ClearLinkedChatId(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) ClearLinkedChatIdTx(tx *sqlx.Tx, linked_chat_id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set linked_chat_id = 0, has_link = 0 where linked_chat_id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, linked_chat_id)

	if err != nil {
		log.Errorf("exec in ClearLinkedChatId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in ClearLinkedChatId(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) Update(ctx context.Context, cMap map[string]interface{}, id int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update channels set %s where id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, id)

	rResult, err = dao.db.Exec(ctx, query, aValues...)

	if err != nil {
		log.Errorf("exec in Update(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Update(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateTx(tx *sqlx.Tx, cMap map[string]interface{}, id int32) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update channels set %s where id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, id)

	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		log.Errorf("exec in Update(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Update(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) Delete(ctx context.Context, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set deactivated = 1, deleted = 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, id)

	if err != nil {
		log.Errorf("exec in Delete(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Delete(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) DeleteTx(tx *sqlx.Tx, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set deactivated = 1, deleted = 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, id)

	if err != nil {
		log.Errorf("exec in Delete(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in Delete(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateReadOutboxMaxId(ctx context.Context, read_outbox_max_id int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set read_outbox_max_id = ? where id = ? and read_outbox_max_id < ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, read_outbox_max_id, id, read_outbox_max_id)

	if err != nil {
		log.Errorf("exec in UpdateReadOutboxMaxId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateReadOutboxMaxId(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateReadOutboxMaxIdTx(tx *sqlx.Tx, read_outbox_max_id int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set read_outbox_max_id = ? where id = ? and read_outbox_max_id < ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, read_outbox_max_id, id, read_outbox_max_id)

	if err != nil {
		log.Errorf("exec in UpdateReadOutboxMaxId(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateReadOutboxMaxId(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) SelectReadOutboxMaxId(ctx context.Context, id int32) (rValue int32, err error) {
	var query = "select read_outbox_max_id from channels where id = ?"
	err = dao.db.Get(ctx, &rValue, query, id)

	if err != nil {
		log.Errorf("get in SelectReadOutboxMaxId(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateGeoAddress(ctx context.Context, hasGeo int8, lat, long float64, radius int32, address string, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set `has_geo` = ?, `lat` = ?, `long` = ?, `accuracy_radius` = ?, `address` = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, hasGeo, lat, long, radius, address, id)

	if err != nil {
		log.Errorf("exec in UpdateGeoAddress(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateGeoAddress(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateGeoAddressTx(tx *sqlx.Tx, hasGeo int8, lat, long float64, radius int32, address string, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set `has_geo` = ?, `lat` = ?, `long` = ?, `accuracy_radius` = ?, `address` = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, hasGeo, lat, long, radius, id)

	if err != nil {
		log.Errorf("exec in UpdateGeoAddressTx(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateGeoAddressTx(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateSlowMode(ctx context.Context, slowMode int8, seconds int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set `slowmode_enabled` = ?, `slowmode_seconds` = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, slowMode, id)

	if err != nil {
		log.Errorf("exec in UpdateSlowMode(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateSlowMode(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) UpdateSlowModeTx(tx *sqlx.Tx, slowMode int8, seconds int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update channels set `slowmode_enabled` = ?, `slowmode_seconds` = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, slowMode, seconds, id)

	if err != nil {
		log.Errorf("exec in UpdateSlowModeTx(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateSlowModeTx(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) SelectPublicChannels(ctx context.Context, offset int32, limit int32) (rList []int32, err error) {
	var query = "select id from channels where username != '' order by date desc,participants_count desc limit ?, ?"
	err = dao.db.Select(ctx, &rList, query, offset, limit)

	if err != nil {
		log.Errorf("select in SelectPublicChannels(_), error: %v", err)
	}

	return
}

func (dao *ChannelsDAO) SelectChannelBannedKeywords(ctx context.Context, id uint32) (keywords []string, err error) {
	var (
		query = "SELECT IFNULL(`banned_keyword`,'[]') banned_keyword FROM `channels` WHERE `id` = ?"
	)

	var keyword string
	err = dao.db.Get(ctx, &keyword, query, id)

	if err != nil {
		log.Errorf("queryx in SelectChannelBannedKeywords(_), error: %v", err)
		return
	}

	json.Unmarshal([]byte(keyword), &keywords)

	return
}
