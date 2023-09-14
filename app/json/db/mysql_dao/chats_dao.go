package mysqldao

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

// ChatsDAO .
type ChatsDAO struct {
	db *sqlx.DB
}

// NewChatsDAO .
func NewChatsDAO(db *sqlx.DB) *ChatsDAO {
	return &ChatsDAO{db}
}

// SelectChatByKey .
func (dao *ChatsDAO) SelectChatBannedRights(ctx context.Context, id uint32) (rights uint32, err error) {
	var (
		query = "SELECT `banned_rights_ex` FROM `chats` WHERE `id` = ?"
	)

	err = dao.db.Get(ctx, &rights, query, id)

	if err != nil {
		log.Errorf("queryx in SelectChatBannedRights(_), error: %v", err)
	}

	return
}

func (dao *ChatsDAO) UpdateChatBannedRights(ctx context.Context, id, rights uint32) (err error) {
	var (
		query string = "UPDATE `chats` SET `banned_rights_ex` = ? WHERE `id` = ?"
		r     sql.Result
	)

	r, err = dao.db.Exec(ctx, query, rights, id)
	if err != nil {
		log.Errorf("Exec in UpdateChatBannedRights()_error: %v", err)
		return
	}

	rowsAffected, err := r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateChatBannedRights()_error: %v", err)
		return
	}

	if rowsAffected < 1 {
		err = errors.New("update fail")
	}

	return
}

// SelectChatBannedKeywords .
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

func (dao *ChatsDAO) UpdateChatKeywords(ctx context.Context, id uint32, keywords []string) (err error) {
	var (
		query string = "UPDATE `chats` SET `banned_keyword` = ? WHERE `id` = ?"
		r     sql.Result
	)

	v, err := json.Marshal(keywords)
	if err != nil {
		return
	}
	keyword := string(v)
	r, err = dao.db.Exec(ctx, query, keyword, id)
	if err != nil {
		log.Errorf("Exec in UpdateChatKeywords()_error: %v", err)
		return
	}

	rowsAffected, err := r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateChatKeywords()_error: %v", err)
		return
	}

	if rowsAffected < 1 {
		err = errors.New("update fail")
	}

	return
}

// SelectChannelBannedRights .
func (dao *ChatsDAO) SelectChannelBannedRights(ctx context.Context, id uint32) (rights uint32, err error) {
	var (
		query = "SELECT `banned_rights_ex` FROM `channels` WHERE `id` = ?"
	)

	err = dao.db.Get(ctx, &rights, query, id)

	if err != nil {
		log.Errorf("queryx in SelectChannelBannedRights(_), error: %v", err)
	}

	return
}

func (dao *ChatsDAO) UpdateChannelBannedRights(ctx context.Context, id, rights uint32) (err error) {
	var (
		query string = "UPDATE `channels` SET `banned_rights_ex` = ? WHERE `id` = ?"
		r     sql.Result
	)

	r, err = dao.db.Exec(ctx, query, rights, id)
	if err != nil {
		log.Errorf("Exec in UpdateChannelBannedRights()_error: %v", err)
		return
	}

	rowsAffected, err := r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateChannelBannedRights()_error: %v", err)
		return
	}

	if rowsAffected < 1 {
		err = errors.New("update fail")
	}

	return
}

// SelectChannelBannedKeyword .
func (dao *ChatsDAO) SelectChannelBannedKeywords(ctx context.Context, id uint32) (keywords []string, err error) {
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

func (dao *ChatsDAO) UpdateChannelKeywords(ctx context.Context, id uint32, keywords []string) (err error) {
	var (
		query string = "UPDATE `channels` SET `banned_keyword` = ? WHERE `id` = ?"
		r     sql.Result
	)

	v, err := json.Marshal(keywords)
	if err != nil {
		return
	}

	keyword := string(v)
	r, err = dao.db.Exec(ctx, query, keyword, id)
	if err != nil {
		log.Errorf("Exec in UpdateChannelKeywords()_error: %v", err)
		return
	}

	rowsAffected, err := r.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateChannelKeywords()_error: %v", err)
		return
	}

	if rowsAffected < 1 {
		err = errors.New("update fail")
	}

	return
}

func (dao *ChatsDAO) ClearChatLink(ctx context.Context, date int32, id int32) (err error) {
	var (
		query   = "update `chats` set `link` = '', `date` = ?, `version` = `version` + 1 where `id` = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, date, id)

	if err != nil {
		log.Errorf("exec in ClearChatLink(_), error: %v", err)
		return
	}

	rowsAffected, err := rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in ClearChatLink(_), error: %v", err)
		return
	}

	if rowsAffected < 1 {
		err = errors.New("update fail")
	}

	return
}

func (dao *ChatsDAO) ClearChannelLink(ctx context.Context, date int32, id int32) (err error) {
	var (
		query   = "update `channels` set `link` = '', `date` = ?, `version` = `version` + 1 where `id` = ?"
		rResult sql.Result
	)
	rResult, err = dao.db.Exec(ctx, query, date, id)

	if err != nil {
		log.Errorf("exec in ClearChannelLink(_), error: %v", err)
		return
	}

	rowsAffected, err := rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in ClearChannelLink(_), error: %v", err)
		return
	}

	if rowsAffected < 1 {
		err = errors.New("update fail")
	}

	return
}

func (dao *ChatsDAO) UpdateChatNickname(ctx context.Context, chatId, uid uint32, nickname string) (err error) {
	var (
		query   = "update chat_participants set nickname=? where chat_id=? and user_id=?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, nickname, chatId, uid)

	if err != nil {
		log.Errorf("exec in UpdateChatNickname(%d, %d, %s), error: %v", chatId, uid, nickname, err)
		return
	}

	rowsAffected, err := rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateChatNickname(%d, %d, %s), error: %v", chatId, uid, nickname, err)
		return
	}

	if rowsAffected != 1 {
		log.Errorf("rowsAffected should be 1, but got %d, UpdateChatNickname(%d, %d, %s), error: %v", rowsAffected, chatId, uid, nickname, err)
		err = errors.New("update fail")
	}

	return
}

func (dao *ChatsDAO) UpdateChannelNickname(ctx context.Context, chatId, uid uint32, nickname string) (err error) {
	var (
		query   = "update channel_participants set nickname=? where channel_id=? and user_id=?"
		rResult sql.Result
	)

	rResult, err = dao.db.Exec(ctx, query, nickname, chatId, uid)

	if err != nil {
		log.Errorf("exec in UpdateChannelNickname(%d, %d, %s), error: %v", chatId, uid, nickname, err)
		return
	}

	rowsAffected, err := rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateChannelNickname(%d, %d, %s), error: %v", chatId, uid, nickname, err)
		return
	}

	if rowsAffected != 1 {
		log.Errorf("rowsAffected should be 1, but got %d, UpdateChannelNickname(%d, %d, %s), error: %v", rowsAffected, chatId, uid, nickname, err)
		err = errors.New("update fail")
	}

	return
}
