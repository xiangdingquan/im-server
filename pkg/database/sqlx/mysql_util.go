package sqlx

import (
	"fmt"
	"strings"

	"context"
	"errors"

	"open.chat/pkg/log"
)

type CommonDAO struct {
	db *DB
}

func NewCommonDAO(db *DB) *CommonDAO {
	return &CommonDAO{db}
}

// 检查是否存在
func (dao *CommonDAO) CheckExists(ctx context.Context, table string, params map[string]interface{}) bool {
	if len(params) == 0 {
		log.Errorf("CheckExists - [%s] error: params empty!", table)
		return false
	}

	names := make([]string, 0, len(params))
	for k, _ := range params {
		names = append(names, k+" = :"+k)
		// log.Info("k: ", k, ", v: ", v)
	}
	sql := fmt.Sprintf("SELECT 1 FROM %s WHERE %s LIMIT 1", table, strings.Join(names, " AND "))
	// log.Info("checkExists - sql: ", sql, ", params: ", params)
	rows, err := dao.db.NamedQuery(ctx, sql, params)
	if err != nil {
		log.Errorf("CheckExists - [%s] error: %s", table, err)
		return false
	}

	defer rows.Close()

	if rows.Next() {
		return true
	} else {
		return false
	}
}

func (dao *CommonDAO) CalcSize(ctx context.Context, table string, params map[string]interface{}) int {
	if len(params) == 0 {
		log.Errorf("calcSize - [%s] error: params empty!", table)
		return 0
	}

	aValues := make([]interface{}, 0, len(params))
	names := make([]string, 0, len(params))
	for k, v := range params {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
		// log.Info("k: ", k, ", v: ", v)
	}
	sql := fmt.Sprintf("SELECT count(id) FROM %s WHERE %s", table, strings.Join(names, " AND "))

	var count int
	err := dao.db.Get(ctx, &count, sql, aValues...)
	if err != nil {
		log.Errorf("calcSize - [%s] error: %s", sql, err)
		return 0
	}
	return count
}

func (dao *CommonDAO) CalcSizeByWhere(ctx context.Context, table, where string) int {
	sql := fmt.Sprintf("SELECT count(id) FROM %s WHERE %s", table, where)

	var count int
	err := dao.db.Get(ctx, &count, sql)
	if err != nil {
		log.Errorf("calcSize - [%s] error: %s", sql, err)
		return 0
	}
	return count
}

// 检查是否存在
func CheckExists(ctx context.Context, db *DB, table string, params map[string]interface{}) (bool, error) {
	_ = ctx

	if len(params) == 0 {
		log.Errorf("checkExists - [%s] error: params empty!", table)
		return false, errors.New("params empty")
	}

	names := make([]string, 0, len(params))
	for k, _ := range params {
		names = append(names, k+" = :"+k)
		// log.Info("k: ", k, ", v: ", v)
	}
	sql := fmt.Sprintf("SELECT 1 FROM %s WHERE %s LIMIT 1", table, strings.Join(names, " AND "))
	// log.Info("checkExists - sql: ", sql, ", params: ", params)
	rows, err := db.NamedQuery(ctx, sql, params)
	if err != nil {
		log.Errorf("checkExists - [%s] error: %s", table, err)
		return false, err
	}

	defer rows.Close()

	if rows.Next() {
		return true, nil
	} else {
		return false, nil
	}
}
