package mysql_dao

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"open.chat/app/service/biz_service/blog/internal/dal/dataobject"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

type BlogMomentsDAO struct {
	db *sqlx.DB
}

func NewBlogMomentsDAO(db *sqlx.DB) *BlogMomentsDAO {
	return &BlogMomentsDAO{db}
}

func (dao *BlogMomentsDAO) Insert(ctx context.Context, do *dataobject.BlogMomentsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into blog_moments(user_id, blog_id, text, entities, video, photos, mention_uids, share_type, member_uids, has_geo, lat, `long`, address, `date`) values (:user_id, :blog_id, :text, :entities, :video, :photos, :mention_uids, :share_type, :member_uids, :has_geo, :lat, :long, :address, :date)"
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

func (dao *BlogMomentsDAO) InsertTx(tx *sqlx.Tx, do *dataobject.BlogMomentsDO) (lastInsertId, rowsAffected int64, err error) {
	var (
		query = "insert into blog_moments(user_id, blog_id, text, entities, video, photos, mention_uids, share_type, member_uids, has_geo, lat, `long`, address, `date`, topics) values (:user_id, :blog_id, :text, :entities, :video, :photos, :mention_uids, :share_type, :member_uids, :has_geo, :lat, :long, :address, :date, :topics)"
		r     sql.Result
	)

	r, err = tx.NamedExec(query, do)
	if err != nil {
		log.Errorf("namedExec in InsertTx(%v), error: %v", do, err)
		return
	}

	lastInsertId, err = r.LastInsertId()
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

func (dao *BlogMomentsDAO) Select(ctx context.Context, id int32) (rValue *dataobject.BlogMomentsDO, err error) {
	var (
		query = "select id, user_id, blog_id, text, entities, video, photos, mention_uids, share_type, member_uids, has_geo, lat, `long`, address, likes, commits, `date`, topics, sort, `deleted` from blog_moments where id = ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, id)

	if err != nil {
		log.Errorf("queryx in Select(_), error: %v", err)
		return
	}

	defer rows.Close()

	if rows.Next() {
		do := &dataobject.BlogMomentsDO{}
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

func (dao *BlogMomentsDAO) Update(ctx context.Context, id int32, cMap map[string]interface{}) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update blog_moments set %s where id = ?", strings.Join(names, ", "))
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

func (dao *BlogMomentsDAO) UpdateTx(tx *sqlx.Tx, id int32, cMap map[string]interface{}) (rowsAffected int64, err error) {
	names := make([]string, 0, len(cMap))
	aValues := make([]interface{}, 0, len(cMap))
	for k, v := range cMap {
		names = append(names, k+" = ?")
		aValues = append(aValues, v)
	}

	var (
		query   = fmt.Sprintf("update blog_moments set %s where id = ?", strings.Join(names, ", "))
		rResult sql.Result
	)

	aValues = append(aValues, id)
	rResult, err = tx.Exec(query, aValues...)

	if err != nil {
		log.Errorf("exec in UpdateTx(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in UpdateTx(_), error: %v", err)
	}

	return
}

func (dao *BlogMomentsDAO) SelectList(ctx context.Context, ids []int32) (rValue []*dataobject.BlogMomentsDO, err error) {
	var (
		query = "select id, user_id, blog_id, text, entities, video, photos, mention_uids, share_type, member_uids, has_geo, lat, `long`, address, likes, commits, `date`, topics, sort, `deleted` from blog_moments where id in (?) and deleted = 0"
		a     []interface{}
		rows  *sqlx.Rows
	)
	rValue = make([]*dataobject.BlogMomentsDO, 0)
	if len(ids) == 0 {
		return rValue, nil
	}
	query, a, err = sqlx.In(query, ids)
	if err != nil {
		log.Errorf("sqlx.In in SelectList(_), error: %v", err)
		return
	}

	rows, err = dao.db.Query(ctx, query, a...)
	if err != nil {
		log.Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		do := &dataobject.BlogMomentsDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectList(_), error: %v", err)
			return
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *BlogMomentsDAO) SelectListIgnoreDeletedFlag(ctx context.Context, ids []int32) (rValue []*dataobject.BlogMomentsDO, err error) {
	var (
		query = "select id, user_id, blog_id, text, entities, video, photos, mention_uids, share_type, member_uids, has_geo, lat, `long`, address, likes, commits, `date`, topics, sort, `deleted` from blog_moments where id in (?)"
		a     []interface{}
		rows  *sqlx.Rows
	)
	rValue = make([]*dataobject.BlogMomentsDO, 0)
	if len(ids) == 0 {
		return rValue, nil
	}
	query, a, err = sqlx.In(query, ids)
	if err != nil {
		log.Errorf("sqlx.In in SelectList(_), error: %v", err)
		return
	}

	rows, err = dao.db.Query(ctx, query, a...)
	if err != nil {
		log.Errorf("queryx in SelectList(_), error: %v", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		do := &dataobject.BlogMomentsDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectList(_), error: %v", err)
			return
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *BlogMomentsDAO) SelectForwardByPublic(ctx context.Context, fromUserId int32, min_id int32, count int32) (rValue []*dataobject.BlogMomentsDO, err error) {
	var (
		query = "select id, user_id, blog_id, text, entities, video, photos, mention_uids, share_type, member_uids, has_geo, lat, `long`, address, likes, commits, `date`, topics, sort, `deleted` from blog_moments bm where id > ? and sort=0 and share_type = 0 and deleted = 0 and not exists (select id from blog_moment_deletes WHERE user_id = ? and blog_id = bm.id and deleted = 0) ORDER BY id ASC limit 0, ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, min_id, fromUserId, count)

	if err != nil {
		log.Errorf("queryx in SelectForwardByPublic(_), error: %v", err)
		return
	}

	defer rows.Close()

	rValue = make([]*dataobject.BlogMomentsDO, 0)
	for rows.Next() {
		do := &dataobject.BlogMomentsDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectForwardByPublic(_), error: %v", err)
			return
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *BlogMomentsDAO) SelectBackwardByPublic(ctx context.Context, fromUserId int32, min_id int32, count int32) (rValue []*dataobject.BlogMomentsDO, err error) {
	var (
		query = "select id, user_id, blog_id, text, entities, video, photos, mention_uids, share_type, member_uids, has_geo, lat, `long`, address, likes, commits, `date`, topics, sort, `deleted` from blog_moments bm where id < ? and sort=0 and share_type = 0 and deleted = 0 and not exists (select id from blog_moment_deletes WHERE user_id = ? and blog_id = bm.id and deleted = 0) ORDER BY id DESC limit 0, ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, min_id, fromUserId, count)

	if err != nil {
		log.Errorf("queryx in SelectBackwardByPublic(_), error: %v", err)
		return
	}

	defer rows.Close()

	rValue = make([]*dataobject.BlogMomentsDO, 0)
	for rows.Next() {
		do := &dataobject.BlogMomentsDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectBackwardByPublic(_), error: %v", err)
			return
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *BlogMomentsDAO) SelectSortedByPublic(ctx context.Context, fromUserId int32) (rValue []*dataobject.BlogMomentsDO, err error) {
	var (
		query = "select id, user_id, blog_id, text, entities, video, photos, mention_uids, share_type, member_uids, has_geo, lat, `long`, address, likes, commits, `date`, topics, sort, `deleted` from blog_moments bm where sort>0 and share_type = 0 and deleted = 0 and not exists (select id from blog_moment_deletes WHERE user_id = ? and blog_id = bm.id and deleted = 0) ORDER BY sort DESC"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, fromUserId)

	if err != nil {
		log.Errorf("queryx in SelectSortedByPublic(_), error: %v", err)
		return
	}

	defer rows.Close()

	rValue = make([]*dataobject.BlogMomentsDO, 0)
	for rows.Next() {
		do := &dataobject.BlogMomentsDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectSortedByPublic(_), error: %v", err)
			return
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *BlogMomentsDAO) SelectForwardBySelf(ctx context.Context, fromUserId int32, min_id int32, count int32) (rValue []*dataobject.BlogMomentsDO, err error) {
	var (
		query = "select id, user_id, blog_id, text, entities, video, photos, mention_uids, share_type, member_uids, has_geo, lat, `long`, address, likes, commits, `date`, topics, sort, `deleted` from blog_moments where id > ? and user_id = ? and deleted = 0 ORDER BY id ASC limit 0, ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, min_id, fromUserId, count)

	if err != nil {
		log.Errorf("queryx in SelectForwardBySelf(_), error: %v", err)
		return
	}

	defer rows.Close()

	rValue = make([]*dataobject.BlogMomentsDO, 0)
	for rows.Next() {
		do := &dataobject.BlogMomentsDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectForwardBySelf(_), error: %v", err)
			return
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *BlogMomentsDAO) SelectBackwardBySelf(ctx context.Context, fromUserId int32, min_id int32, count int32) (rValue []*dataobject.BlogMomentsDO, err error) {
	var (
		query = "select id, user_id, blog_id, text, entities, video, photos, mention_uids, share_type, member_uids, has_geo, lat, `long`, address, likes, commits, `date`, topics, sort, `deleted` from blog_moments where id < ? and user_id = ? and deleted = 0 ORDER BY id DESC limit 0, ?"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, min_id, fromUserId, count)

	if err != nil {
		log.Errorf("queryx in SelectBackwardBySelf(_), error: %v", err)
		return
	}

	defer rows.Close()

	rValue = make([]*dataobject.BlogMomentsDO, 0)
	for rows.Next() {
		do := &dataobject.BlogMomentsDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectBackwardBySelf(_), error: %v", err)
			return
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *BlogMomentsDAO) SelectForwardByUser(ctx context.Context, fromUserId int32, uid int32, visibles []int8, min_id int32, count int32, minDate uint32) (rValue []*dataobject.BlogMomentsDO, err error) {
	var (
		query = "select id, user_id, blog_id, text, entities, video, photos, mention_uids, share_type, member_uids, has_geo, lat, `long`, address, likes, commits, `date`, topics, sort, `deleted` from blog_moments bm where id > ? and user_id = ? and share_type in (?) and `date` > ? and deleted = 0 and not exists (select id from blog_moment_deletes WHERE user_id = ? and blog_id = bm.id and deleted = 0) ORDER BY id ASC limit 0, ?"
		a     []interface{}
		rows  *sqlx.Rows
	)
	rValue = make([]*dataobject.BlogMomentsDO, 0)
	query, a, err = sqlx.In(query, min_id, uid, visibles, minDate, fromUserId, count)
	if err != nil {
		log.Errorf("sqlx.In in SelectForwardByUsers(_), error: %v", err)
		return
	}

	rows, err = dao.db.Query(ctx, query, a...)
	if err != nil {
		log.Errorf("queryx in SelectForwardByUsers(_), error: %v", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		do := &dataobject.BlogMomentsDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectForwardByUsers(_), error: %v", err)
			return
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *BlogMomentsDAO) SelectBackwardByUser(ctx context.Context, fromUserId int32, uid int32, visibles []int8, min_id int32, count int32, minDate uint32) (rValue []*dataobject.BlogMomentsDO, err error) {
	var (
		query = "select id, user_id, blog_id, text, entities, video, photos, mention_uids, share_type, member_uids, has_geo, lat, `long`, address, likes, commits, `date`, topics, sort, `deleted` from blog_moments bm where id < ? and user_id = ? and share_type in (?) and `date` > ? and deleted = 0 and not exists (select id from blog_moment_deletes WHERE user_id = ? and blog_id = bm.id and deleted = 0) ORDER BY id DESC limit 0, ?"
		a     []interface{}
		rows  *sqlx.Rows
	)
	rValue = make([]*dataobject.BlogMomentsDO, 0)
	query, a, err = sqlx.In(query, min_id, uid, visibles, minDate, fromUserId, count)
	if err != nil {
		log.Errorf("sqlx.In in SelectForwardByUsers(_), error: %v", err)
		return
	}

	rows, err = dao.db.Query(ctx, query, a...)
	if err != nil {
		log.Errorf("queryx in SelectBackwardByUsers(_), error: %v", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		do := &dataobject.BlogMomentsDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectBackwardByUsers(_), error: %v", err)
			return
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *BlogMomentsDAO) SelectForwardByUsers(ctx context.Context, fromUserId int32, uids []int32, min_id int32, count int32) (rValue []*dataobject.BlogMomentsDO, err error) {
	var (
		query = "select id, user_id, blog_id, text, entities, video, photos, mention_uids, share_type, member_uids, has_geo, lat, `long`, address, likes, commits, `date`, topics, sort, `deleted` from blog_moments bm where id > ? and user_id in (?) and deleted = 0 and not exists (select id from blog_moment_deletes WHERE user_id = ? and blog_id = bm.id and deleted = 0) ORDER BY id ASC limit 0, ?"
		a     []interface{}
		rows  *sqlx.Rows
	)
	rValue = make([]*dataobject.BlogMomentsDO, 0)
	if len(uids) == 0 {
		return rValue, nil
	}

	query, a, err = sqlx.In(query, min_id, uids, fromUserId, count)
	if err != nil {
		log.Errorf("sqlx.In in SelectForwardByUsers(_), error: %v", err)
		return
	}

	rows, err = dao.db.Query(ctx, query, a...)
	if err != nil {
		log.Errorf("queryx in SelectForwardByUsers(_), error: %v", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		do := &dataobject.BlogMomentsDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectForwardByUsers(_), error: %v", err)
			return
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *BlogMomentsDAO) SelectBackwardByUsers(ctx context.Context, fromUserId int32, uids []int32, min_id int32, count int32) (rValue []*dataobject.BlogMomentsDO, err error) {
	var (
		query = "select id, user_id, blog_id, text, entities, video, photos, mention_uids, share_type, member_uids, has_geo, lat, `long`, address, likes, commits, `date`, topics, sort, `deleted` from blog_moments bm where id < ? and user_id in (?) and deleted = 0 and not exists (select id from blog_moment_deletes WHERE user_id = ? and blog_id = bm.id and deleted = 0) ORDER BY id DESC limit 0, ?"
		a     []interface{}
		rows  *sqlx.Rows
	)
	rValue = make([]*dataobject.BlogMomentsDO, 0)
	if len(uids) == 0 {
		return rValue, nil
	}

	query, a, err = sqlx.In(query, min_id, uids, fromUserId, count)
	if err != nil {
		log.Errorf("sqlx.In in SelectBackwardByUsers(_), error: %v", err)
		return
	}

	rows, err = dao.db.Query(ctx, query, a...)
	if err != nil {
		log.Errorf("queryx in SelectBackwardByUsers(_), error: %v", err)
		return
	}

	defer rows.Close()

	for rows.Next() {
		do := &dataobject.BlogMomentsDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectBackwardByUsers(_), error: %v", err)
			return
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *BlogMomentsDAO) SelectForwardByTopic(ctx context.Context, fromUserId int32, min_id int32, count int32, topicId int32, excludeLiked bool) (rValue []*dataobject.BlogMomentsDO, err error) {
	var (
		query = ""
		rows  *sqlx.Rows
	)
	if excludeLiked {
		query = "select id, user_id, blog_id, text, entities, video, photos, mention_uids, share_type, member_uids, has_geo, lat, `long`, address, likes, commits, `date`, topics, sort, `deleted` " +
			"from blog_moments bm where id<? and share_type = 0 and deleted = 0 " +
			"and exists (select moment_id from blog_topic_mappings where topic_id=? and moment_id=bm.id) " +
			"and not exists (select id from blog_moment_deletes WHERE user_id = ? and blog_id = bm.id and deleted = 0) " +
			"and not exists (select id from blog_likes where blog_id=bm.id and deleted=0) " +
			"ORDER BY id DESC limit 0, ?"
	} else {
		query = "select id, user_id, blog_id, text, entities, video, photos, mention_uids, share_type, member_uids, has_geo, lat, `long`, address, likes, commits, `date`, topics, sort, `deleted` " +
			"from blog_moments bm where id<? and share_type = 0 and deleted = 0 " +
			"and exists (select moment_id from blog_topic_mappings where topic_id=? and moment_id=bm.id) " +
			"and not exists (select id from blog_moment_deletes WHERE user_id = ? and blog_id = bm.id and deleted = 0) " +
			"ORDER BY id DESC limit 0, ?"
	}
	rows, err = dao.db.Query(ctx, query, min_id, topicId, fromUserId, count)

	if err != nil {
		log.Errorf("queryx in SelectForwardByTopic(_), error: %v", err)
		return
	}

	defer rows.Close()

	rValue = make([]*dataobject.BlogMomentsDO, 0)
	for rows.Next() {
		do := &dataobject.BlogMomentsDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectForwardByTopic(_), error: %v", err)
			return
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *BlogMomentsDAO) SelectBackwardByTopic(ctx context.Context, fromUserId int32, min_id int32, count int32, topicId int32, excludeLiked bool) (rValue []*dataobject.BlogMomentsDO, err error) {
	var (
		query = ""
		rows  *sqlx.Rows
	)
	if excludeLiked {
		query = "select id, user_id, blog_id, text, entities, video, photos, mention_uids, share_type, member_uids, has_geo, lat, `long`, address, likes, commits, `date`, topics, sort, `deleted` " +
			"from blog_moments bm where id<? and share_type = 0 and deleted = 0 " +
			"and exists (select moment_id from blog_topic_mappings where topic_id=? and moment_id=bm.id) " +
			"and not exists (select id from blog_moment_deletes WHERE user_id = ? and blog_id = bm.id and deleted = 0) " +
			"and not exists (select id from blog_likes where blog_id=bm.id and deleted=0) " +
			"ORDER BY id DESC limit 0, ?"
	} else {
		query = "select id, user_id, blog_id, text, entities, video, photos, mention_uids, share_type, member_uids, has_geo, lat, `long`, address, likes, commits, `date`, topics, sort, `deleted` " +
			"from blog_moments bm where id<? and share_type = 0 and deleted = 0 " +
			"and exists (select moment_id from blog_topic_mappings where topic_id=? and moment_id=bm.id) " +
			"and not exists (select id from blog_moment_deletes WHERE user_id = ? and blog_id = bm.id and deleted = 0) " +
			"ORDER BY id DESC limit 0, ?"
	}
	rows, err = dao.db.Query(ctx, query, min_id, topicId, fromUserId, count)

	if err != nil {
		log.Errorf("queryx in SelectBackwardByPublic(_), error: %v", err)
		return
	}

	defer rows.Close()

	rValue = make([]*dataobject.BlogMomentsDO, 0)
	for rows.Next() {
		do := &dataobject.BlogMomentsDO{}
		err = rows.StructScan(do)
		if err != nil {
			log.Errorf("structScan in SelectBackwardByPublic(_), error: %v", err)
			return
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *BlogMomentsDAO) SelectLikedByTopic(ctx context.Context, fromUserId int32, topicId int32, limitDays int32) (rValue []*dataobject.BlogMomentsDO, err error) {
	var (
		query = ""
		rows  *sqlx.Rows
	)

	query = "select bm.id, bm.user_id, bm.blog_id, bm.text, bm.entities, bm.video, bm.photos, bm.mention_uids, bm.share_type, bm.member_uids, bm.has_geo, bm.lat, bm.`long`, bm.address, bm.likes, bm.commits, bm.`date`, bm.topics, bm.sort, bm.`deleted` "
	query += "from blog_moments bm "
	if limitDays == 0 {
		query += "join (select blog_id,count(*) as hot from blog_likes where deleted=0 group by blog_id) bl "
	} else {
		query += "join (select blog_id,count(*) as hot from blog_likes where created_at>now()-interval ? day and deleted=0 group by blog_id) bl "
	}
	query += `
on bm.id=bl.blog_id
where share_type = 0 and bm.deleted = 0 
and exists (select moment_id from blog_topic_mappings where topic_id=? and moment_id=bm.id) 
and not exists (select id from blog_moment_deletes WHERE user_id = ? and blog_id = bm.id and deleted = 0)
order by hot desc, id desc
`
	if limitDays == 0 {
		rows, err = dao.db.Query(ctx, query, topicId, fromUserId)
	} else {
		rows, err = dao.db.Query(ctx, query, limitDays, topicId, fromUserId)
	}

	if err != nil {
		log.Errorf("queryx in SelectLikedByTopic(_), error: %v", err)
		return
	}

	rValue, err = dao.makeDoList(rows)
	if err != nil {
		log.Errorf("structScan in SelectLikedByTopic(_), error: %v", err)
	}

	return
}

func (dao *BlogMomentsDAO) makeDoList(rows *sqlx.Rows) (rValue []*dataobject.BlogMomentsDO, err error) {
	defer rows.Close()

	rValue = make([]*dataobject.BlogMomentsDO, 0)
	for rows.Next() {
		do := &dataobject.BlogMomentsDO{}
		err = rows.StructScan(do)
		if err != nil {
			return
		} else {
			rValue = append(rValue, do)
		}
	}

	return
}

func (dao *BlogMomentsDAO) LikeTx(tx *sqlx.Tx, id int32, liked bool) (rowsAffected int64, err error) {
	var (
		query   = "update blog_moments set likes=likes+1 where id = ? and deleted = 0"
		rResult sql.Result
	)
	if !liked {
		query = "update blog_moments set likes=likes-1 where id = ? and deleted = 0"
	}
	rResult, err = tx.Exec(query, id)

	if err != nil {
		log.Errorf("exec in LikeTx(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in LikeTx(_), error: %v", err)
	}

	return
}

func (dao *BlogMomentsDAO) DeleteBlog(tx *sqlx.Tx, fromUserId int32, ids []int32) (rowsAffected int64, err error) {
	var (
		query   = "update blog_moments set deleted = 1 where id in (?) and user_id = ? and deleted = 0"
		a       []interface{}
		rResult sql.Result
	)
	if len(ids) == 0 {
		return 0, nil
	}
	query, a, err = sqlx.In(query, ids, fromUserId)
	if err != nil {
		log.Errorf("sqlx.In in DeleteBlog(_), error: %v", err)
		return
	}

	rResult, err = tx.Exec(query, a...)
	if err != nil {
		log.Errorf("exec in DeleteBlog(_), error: %v", err)
		return
	}

	rowsAffected, err = rResult.RowsAffected()
	if err != nil {
		log.Errorf("rowsAffected in DeleteBlog(_), error: %v", err)
	}

	return
}

func (dao *BlogMomentsDAO) SelectIdByOffset(ctx context.Context, uid int32, offset int32) (int32, error) {
	var (
		query = "select id from blog_moments where user_id = ? and deleted = 0 order by id desc limit ?,1"
	)

	var id int32
	err := dao.db.Get(ctx, &id, query, uid, offset)
	if err != nil {
		log.Errorf("get in SelectIdByOffset(%d, %d), error: %v", uid, offset, err)
		return 0, err
	}

	return id, nil
}

func (dao *BlogMomentsDAO) CountBlogMomentByUid(ctx context.Context, uid int32) (int32, error) {
	var (
		query = "select count(*) from blog_moments where user_id = ? and deleted=0"
	)

	var count int32
	err := dao.db.Get(ctx, &count, query, uid)
	if err != nil {
		log.Errorf("get in CountBlogMementByUid(%d), uid, error: %v", uid, err)
		return 0, err
	}

	return count, nil
}
