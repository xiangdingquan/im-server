package mysql_dao

import (
	"context"
	"database/sql"
	"encoding/json"

	"open.chat/app/service/biz_service/chat/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type ChatsDAO struct {
	db *sqlx.DB
}

func NewChatsDAO(db *sqlx.DB) *ChatsDAO {
	return &ChatsDAO{db}
}

func (dao *ChatsDAO) Insert(ctx context.Context, do *dataobject.ChatsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chats(creator_user_id, access_hash, random_id, participant_count, title, about, `date`) values (:creator_user_id, :access_hash, :random_id, :participant_count, :title, :about, :date)"
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

func (dao *ChatsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.ChatsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into chats(creator_user_id, access_hash, random_id, participant_count, title, about, `date`) values (:creator_user_id, :access_hash, :random_id, :participant_count, :title, :about, :date)"
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

func (dao *ChatsDAO) Select(ctx context.Context, id int32) (rValue *dataobject.ChatsDO, err error) {
	var (
		query = "select id, creator_user_id, access_hash, participant_count, title, about, photo, link, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, id)

	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.ChatsDO{}
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

func (dao *ChatsDAO) SelectLastCreator(ctx context.Context, creator_user_id int32) (rValue *dataobject.ChatsDO, err error) {
	var (
		query = "select id, creator_user_id, access_hash, participant_count, title, about, photo, link, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where creator_user_id = ? order by `date` desc limit 1"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, creator_user_id)

	if err != nil {
		log.Errorf("queryx in SelectLastCreator(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.ChatsDO{}
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

func (dao *ChatsDAO) UpdateTitle(ctx context.Context, title string, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chats set title = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, title, id)

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

func (dao *ChatsDAO) UpdateTitleTx(tx *sqlx.Tx, title string, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chats set title = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, title, id)

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

func (dao *ChatsDAO) UpdateAbout(ctx context.Context, about string, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chats set about = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, about, id)

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

func (dao *ChatsDAO) UpdateAboutTx(tx *sqlx.Tx, about string, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chats set about = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, about, id)

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

func (dao *ChatsDAO) UpdateNotice(ctx context.Context, notice string, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chats set notice = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, notice, id)

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

func (dao *ChatsDAO) UpdateNoticeTx(tx *sqlx.Tx, notice string, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chats set notice = ? where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, notice, id)

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

func (dao *ChatsDAO) SelectByIdList(ctx context.Context, idList []int32) (rList []dataobject.ChatsDO, err error) {
	var (
		query = "select id, creator_user_id, access_hash, participant_count, title, about, photo, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where id in (?)"
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

	var values []dataobject.ChatsDO
	for rows.Next() {
		v := dataobject.ChatsDO{}
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

func (dao *ChatsDAO) UpdateParticipantCount(ctx context.Context, participant_count int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chats set participant_count = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, participant_count, id)

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

func (dao *ChatsDAO) UpdateParticipantCountTx(tx *sqlx.Tx, participant_count int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chats set participant_count = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, participant_count, id)

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

func (dao *ChatsDAO) UpdatePhoto(ctx context.Context, photo string, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chats set photo = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, photo, id)

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

func (dao *ChatsDAO) UpdatePhotoTx(tx *sqlx.Tx, photo string, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chats set photo = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, photo, id)

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

func (dao *ChatsDAO) UpdateAdminsEnabled(ctx context.Context, admins_enabled int8, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chats set admins_enabled = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, admins_enabled, id)

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

func (dao *ChatsDAO) UpdateAdminsEnabledTx(tx *sqlx.Tx, admins_enabled int8, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chats set admins_enabled = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, admins_enabled, id)

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

func (dao *ChatsDAO) UpdateDefaultBannedRights(ctx context.Context, default_banned_rights int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chats set default_banned_rights = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, default_banned_rights, id)

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

func (dao *ChatsDAO) UpdateDefaultBannedRightsTx(tx *sqlx.Tx, default_banned_rights int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chats set default_banned_rights = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, default_banned_rights, id)

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

func (dao *ChatsDAO) UpdateVersion(ctx context.Context, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chats set version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, id)

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

func (dao *ChatsDAO) UpdateVersionTx(tx *sqlx.Tx, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chats set version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, id)

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

func (dao *ChatsDAO) UpdateDeactivated(ctx context.Context, deactivated int8, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chats set deactivated = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, deactivated, id)

	if err != nil {
		log.Errorf("exec in UpdateDeactivated(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateDeactivated(_), error: %v", err)
	}

	return
}

func (dao *ChatsDAO) UpdateDeactivatedTx(tx *sqlx.Tx, deactivated int8, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chats set deactivated = ?, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, deactivated, id)

	if err != nil {
		log.Errorf("exec in UpdateDeactivated(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateDeactivated(_), error: %v", err)
	}

	return
}

func (dao *ChatsDAO) SelectByLink(ctx context.Context, link string) (rValue *dataobject.ChatsDO, err error) {
	var (
		query = "select id, creator_user_id, access_hash, participant_count, title, about, photo, link, admins_enabled, default_banned_rights, migrated_to_id, migrated_to_access_hash, deactivated, version, `date` from chats where deactivated = 0 and link = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, link)

	if err != nil {
		log.Errorf("queryx in SelectByLink(_), error: %v", err)
		return
	}

	defer rows.Close()

	do := &dataobject.ChatsDO{}
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

func (dao *ChatsDAO) UpdateLink(ctx context.Context, link string, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chats set link = ?, `date` = ?, version = version + 1 where id = ?"
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

func (dao *ChatsDAO) UpdateLinkTx(tx *sqlx.Tx, link string, date int32, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chats set link = ?, `date` = ?, version = version + 1 where id = ?"
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

func (dao *ChatsDAO) UpdateMigratedTo(ctx context.Context, migrated_to_id int32, migrated_to_access_hash int64, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chats set migrated_to_id = ?, migrated_to_access_hash = ?, participant_count = 0, deactivated = 1, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, migrated_to_id, migrated_to_access_hash, id)

	if err != nil {
		log.Errorf("exec in UpdateMigratedTo(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateMigratedTo(_), error: %v", err)
	}

	return
}

func (dao *ChatsDAO) UpdateMigratedToTx(tx *sqlx.Tx, migrated_to_id int32, migrated_to_access_hash int64, id int32) (rowsAffected int64, err error) {
	var (
		query   = "update chats set migrated_to_id = ?, migrated_to_access_hash = ?, participant_count = 0, deactivated = 1, version = version + 1 where id = ?"
		rResult sql.Result
	)
	rResult, err = tx.Exec(query, migrated_to_id, migrated_to_access_hash, id)

	if err != nil {
		log.Errorf("exec in UpdateMigratedTo(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateMigratedTo(_), error: %v", err)
	}

	return
}

func (dao *ChatsDAO) SelectChatBannedKeywords(ctx context.Context, id uint32) (keywords []string, err error) {
	var (
		query = "SELECT IFNULL(`banned_keyword`,'[]') banned_keyword FROM `chats` WHERE `id` = ?"
	)

	var keyword string
	err = dao.db.Get(ctx, &keyword, query, id)

	if err != nil {
		log.Errorf("queryx in SelectChatBannedKeywords(_), error: %v", err)
		return
	}

	json.Unmarshal([]byte(keyword), &keywords)

	return
}
