package core

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/gogo/protobuf/types"
	"open.chat/app/service/biz_service/blog/internal/dal/dataobject"
	"open.chat/app/service/biz_service/blog/internal/dao"
	idgen "open.chat/app/service/idgen/client"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

const (
	BLOG_COMMENT_TYPE_BLOG int8 = iota
	BLOG_COMMENT_TYPE_COMMENT
)

const (
	PTS_UPDATE_TYPE_UNKNOWN = 0
	PTS_UPDATE_NEW_BLOG     = 21
	PTS_UPDATE_DELETE_BLOG  = 22
	PTS_UPDATE_BLOG_FOLLOW  = 23
	PTS_UPDATE_BLOG_COMMENT = 24
	PTS_UPDATE_BLOG_LIKE    = 25
)

type BlogCore struct {
	*dao.Dao
}

func New(dao *dao.Dao) *BlogCore {
	return &BlogCore{
		Dao: dao,
	}
}

func (m *BlogCore) getUpdateType(update *mtproto.Update) int8 {
	switch update.PredicateName {
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

func (m *BlogCore) makeMicroBlogs(ctx context.Context, fromUserId int32, bmDos []*dataobject.BlogMomentsDO, fromGm bool) (*mtproto.MicroBlogs, error) {
	var (
		result *mtproto.MicroBlogs = mtproto.MakeTLMicroBlogsNotModified(nil).To_MicroBlogs()
		err    error
	)
	if len(bmDos) == 0 {
		return result, nil
	}
	result = mtproto.MakeTLMicroBlogs(nil).To_MicroBlogs()
	result.Topics = mtproto.MakeTLBlogsTopics(&mtproto.Blogs_Topics{
		Count:  0,
		Topics: []*mtproto.Blogs_Topic{},
	}).To_Blogs_Topics()
	var blogIds []int32 = make([]int32, 0, len(bmDos))
	var list map[int32]*mtproto.MicroBlog = make(map[int32]*mtproto.MicroBlog)
	for _, bm := range bmDos {
		var userIds []int32 = make([]int32, 0)
		if len(bm.MemberUserIds) > 0 {
			json.Unmarshal([]byte(bm.MemberUserIds), &userIds)
		}
		visible := true
		if !fromGm && bm.UserId != fromUserId {
			if bm.ShareType == model.VisibleType_Allow || bm.ShareType == model.VisibleType_NotAllow {
				contains := false
				for _, f := range userIds {
					if f == fromUserId {
						contains = true
						break
					}
				}
				visible = contains
				if bm.ShareType == model.VisibleType_NotAllow {
					visible = !visible
				}
			}
		}
		if !visible {
			continue
		}

		blog := mtproto.MakeTLMicroBlog(&mtproto.MicroBlog{
			Id:     bm.Id,
			UserId: bm.UserId,
			Date:   bm.Date,
			Text:   bm.Text,
			Sort:   bm.Sort,
		}).To_MicroBlog()

		c := &mtproto.BlogContent{
			Document: &mtproto.Document{},
		}
		if len(bm.Video) > 0 {
			json.Unmarshal([]byte(bm.Video), c.Document)
			blog.Content = mtproto.MakeTLBlogContentDocument(c).To_BlogContent()
		} else if len(bm.Photos) > 0 {
			json.Unmarshal([]byte(bm.Photos), &c.Photos)
			blog.Content = mtproto.MakeTLBlogContentPhotos(c).To_BlogContent()
		}

		if len(bm.MentionUserIds) > 0 {
			var users []int32
			_ = json.Unmarshal([]byte(bm.MentionUserIds), &users)
			for _, uid := range users {
				if uid == fromUserId {
					blog.Mentioned = true
					break
				}
			}
		}

		if bm.HasGeo {
			blog.GeoPoint = mtproto.MakeTLBlogGeoPoint(&mtproto.BlogGeoPoint{
				GeoPoint: mtproto.MakeTLGeoPoint(&mtproto.GeoPoint{
					Long: bm.Long,
					Lat:  bm.Lat,
				}).To_GeoPoint(),
				Address: bm.Address,
			}).To_BlogGeoPoint()
		}

		topicIds := make([]int32, 0)
		err = json.Unmarshal([]byte(bm.Topics), &topicIds)
		if err != nil {
			log.Errorf("unmarshal bm.Topics failed, error: %v", err)
		}
		var topicDos []*dataobject.BlogTopicsDO
		if len(topicIds) > 0 {
			topicDos, _ = m.BlogTopicsDAO.SelectByIds(ctx, topicIds)
		}
		blog.Topics = m.topicListDoToMTProto(topicDos, int32(len(topicDos)))

		blog.LikeCount = int32(m.CommonDAO.CalcSize(ctx, "blog_likes", map[string]interface{}{
			"type":    BLOG_COMMENT_TYPE_BLOG,
			"blog_id": bm.Id,
			"deleted": 0,
		}))

		commentCount := int32(m.CommonDAO.CalcSize(ctx, "blog_comments", map[string]interface{}{
			"type":    BLOG_COMMENT_TYPE_BLOG,
			"blog_id": bm.Id,
			"deleted": 0,
		}))

		if commentCount > 0 {
			blog.CommentCount = &types.Int32Value{Value: commentCount}
		}

		//log.Debugf("makeMicroBlog, entities: %s", bm.Entities)
		err = json.Unmarshal([]byte(bm.Entities), &blog.Entities)
		if err != nil {
			log.Errorf("unmarshal bm.Entities failed, error: %v", err)
		}
		//log.Debugf("makeMicroBlog, entities: %v", blog.Entities)

		list[bm.Id] = blog
		blogIds = append(blogIds, bm.Id)
		result.Blogs = append(result.Blogs, blog)
	}
	if len(blogIds) > 0 {
		var blDos []*dataobject.BlogLikesDO
		blDos, err = m.BlogLikesDAO.SelectByUserAndBlogIds(ctx, fromUserId, blogIds)
		if err == nil {
			for _, bl := range blDos {
				list[bl.BlogId].Liked = true
			}
		}
	}
	result.Count = int32(len(result.Blogs))
	return result, nil
}

func (m *BlogCore) SendBlog(ctx context.Context, fromUserId int32, visible *model.VisibleTypeUtil, text string, entities []*mtproto.MessageEntity, mentionUserIds []int32, content *mtproto.BlogContent, geo *mtproto.BlogGeoPoint, topicIds []int32) (model.IDList, *mtproto.MicroBlog, error) {
	var (
		blogId       = int32(idgen.NextUserBlogId(ctx, fromUserId))
		now          = int32(time.Now().Unix())
		toJsonString = func(v interface{}) string {
			var result string
			if v != nil {
				b, _ := json.Marshal(v)
				result = string(b)
			}
			return result
		}
		mentionUids model.IDList
		filteUser   model.IDList
	)

	if visible.VisibleType == model.VisibleType_Allow || visible.VisibleType == model.VisibleType_NotAllow {
		filteUser = visible.UserIds
		if len(visible.GroupTags) > 0 {
			bgtDo, err := m.BlogGroupTagsDAO.SelectByUserAndTags(ctx, fromUserId, visible.GroupTags)
			if err != nil {
				return mentionUids, nil, mtproto.ErrInternelServerError
			}
			for _, tag := range bgtDo {
				if len(tag.MemberUserIds) > 0 {
					var users []int32
					err = json.Unmarshal([]byte(tag.MemberUserIds), &users)
					if err == nil && len(users) > 0 {
						filteUser.AddIfNot(users...)
						//filteUser = append(filteUser, users...)
					}
				}
			}
		}

		for _, mu := range mentionUserIds {
			contains := false
			for _, f := range filteUser {
				if f == mu {
					contains = true
					break
				}
			}

			if contains {
				if visible.VisibleType == model.VisibleType_Allow {
					mentionUids.AddIfNot(mu)
				}
			} else {
				if visible.VisibleType == model.VisibleType_NotAllow {
					mentionUids.AddIfNot(mu)
				}
			}
		}
	} else {
		mentionUids = mentionUserIds
	}

	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		blogDO := &dataobject.BlogMomentsDO{
			UserId:    fromUserId,
			BlogId:    blogId,
			Text:      text,
			ShareType: visible.VisibleType,
			Date:      now,
		}

		if content.GetDocument() != nil {
			blogDO.Video = toJsonString(content.GetDocument())
		}

		if len(content.GetPhotos()) > 0 {
			blogDO.Photos = toJsonString(content.GetPhotos())
		}

		if len(mentionUids) > 0 {
			blogDO.MentionUserIds = toJsonString(mentionUids)
		}

		if len(filteUser) > 0 {
			blogDO.MemberUserIds = toJsonString(filteUser)
		}

		if geo != nil {
			blogDO.HasGeo = true
			blogDO.Lat = geo.GetGeoPoint().GetLat()
			blogDO.Long = geo.GetGeoPoint().GetLong()
			blogDO.Address = geo.GetAddress()
		}

		blogDO.Topics = toJsonString(topicIds)

		entities, err := json.Marshal(entities)
		if err != nil {
			log.Errorf("SendBlog, marshal entities failed, error: %v", err)
		}
		blogDO.Entities = string(entities)

		lastInsertId, _, err := m.BlogMomentsDAO.InsertTx(tx, blogDO)
		if err != nil {
			result.Err = err
			return
		}

		if lastInsertId == 0 {
			result.Err = fmt.Errorf("insert error")
			return
		}

		m.BlogsDAO.IncMoments(tx, fromUserId, 1)
		result.Data = lastInsertId

		if len(topicIds) > 0 {
			doList := make([]*dataobject.BlogTopicMappingsDo, len(topicIds))
			for i, id := range topicIds {
				doList[i] = &dataobject.BlogTopicMappingsDo{
					TopicId:  id,
					MomentId: int32(lastInsertId),
				}
			}
			_, _, err = m.BlogTopicMappingsDAO.InsertTx(tx, doList)
			if err != nil {
				result.Err = err
				return
			}
		}
	})

	if tR.Err != nil {
		return mentionUids, nil, mtproto.ErrInternelServerError
	}

	makeTopicsByIds := func(topicIds []int32) *mtproto.Blogs_Topics {
		if len(topicIds) > 0 {
			dos, err := m.BlogTopicsDAO.SelectByIds(ctx, topicIds)
			if err != nil {
				log.Errorf("select topic failed, error: %v", err)
				return nil
			}
			return m.topicListDoToMTProto(dos, int32(len(dos)))
		}
		return nil
	}

	mentionUids.AddFrontIfNot(fromUserId)
	blog := mtproto.MakeTLMicroBlog(&mtproto.MicroBlog{
		Id:       int32(tR.Data.(int64)),
		UserId:   fromUserId,
		Date:     now,
		Text:     text,
		Entities: entities,
		Content:  content,
		GeoPoint: geo,
		Topics:   makeTopicsByIds(topicIds),
	}).To_MicroBlog()
	return mentionUids, blog, nil
}

func (m *BlogCore) LikeBlog(ctx context.Context, fromUserId int32, id int32, liked bool) (int32, *mtproto.Blogs_UserDate, error) {
	var (
		now              = int32(time.Now().Unix())
		pushUserId int32 = 0
	)
	var userDate *mtproto.Blogs_UserDate
	likeDo, err := m.BlogLikesDAO.SelectByBlogId(ctx, fromUserId, id)
	if err != nil {
		return 0, userDate, mtproto.ErrInternelServerError
	}
	if likeDo == nil && !liked {
		return 0, userDate, mtproto.ErrButtonTypeInvalid
	}

	if likeDo != nil {
		userDate = mtproto.MakeTLBlogsUserDate(&mtproto.Blogs_UserDate{
			UserId: likeDo.UserId,
			Date:   likeDo.Date,
		}).To_Blogs_UserDate()
		if likeDo.Deleted != liked {
			return 0, userDate, nil
		}
	}

	bmDo, err := m.BlogMomentsDAO.Select(ctx, id)
	if err != nil {
		return 0, userDate, mtproto.ErrInternelServerError
	}
	if bmDo == nil || bmDo.Deleted {
		return 0, userDate, mtproto.ErrButtonTypeInvalid
	}
	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		if likeDo == nil {
			do := &dataobject.BlogLikesDO{
				UserId: fromUserId,
				Type:   BLOG_COMMENT_TYPE_BLOG,
				BlogId: bmDo.Id,
				Date:   now,
			}
			lastInsertId, _, err := m.BlogLikesDAO.InsertTx(tx, do)
			if err != nil {
				result.Err = err
				return
			}
			result.Data = lastInsertId
			pushUserId = bmDo.UserId
		} else {
			result.Data = int64(likeDo.Id)
			_, err := m.BlogLikesDAO.UpdateLikeTx(tx, likeDo.Id, liked)
			if err != nil {
				result.Err = err
				return
			}
		}
		_, err := m.BlogMomentsDAO.LikeTx(tx, id, liked)
		if err != nil {
			result.Err = err
			return
		}
		if liked {
			m.BlogsDAO.IncLikes(tx, bmDo.UserId, 1)
		} else {
			m.BlogsDAO.DecLikes(tx, bmDo.UserId, 1)
		}
	})
	if tR.Err != nil {
		return 0, nil, mtproto.ErrInternelServerError
	}

	if liked {
		likeId := int32(tR.Data.(int64))
		likeDo, _ = m.BlogLikesDAO.Select(ctx, likeId)
		if likeDo != nil {
			userDate = mtproto.MakeTLBlogsUserDate(&mtproto.Blogs_UserDate{
				UserId: likeDo.UserId,
				Date:   likeDo.Date,
			}).To_Blogs_UserDate()
		}
	}
	return pushUserId, userDate, nil
}

func (m *BlogCore) LikeReply(ctx context.Context, fromUserId int32, id int32, liked bool) (int32, *mtproto.Blogs_UserDate, error) {
	var (
		now              = int32(time.Now().Unix())
		pushUserId int32 = 0
	)
	var userDate *mtproto.Blogs_UserDate
	likeDo, err := m.BlogLikesDAO.SelectByCommentId(ctx, fromUserId, id)
	if err != nil {
		return 0, userDate, mtproto.ErrInternelServerError
	}
	if likeDo == nil && !liked {
		return 0, userDate, mtproto.ErrButtonTypeInvalid
	}

	if likeDo != nil {
		userDate = mtproto.MakeTLBlogsUserDate(&mtproto.Blogs_UserDate{
			UserId: likeDo.UserId,
			Date:   likeDo.Date,
		}).To_Blogs_UserDate()
		if likeDo.Deleted != liked {
			return 0, userDate, nil
		}
	}
	bcDo, err := m.BlogCommentsDAO.Select(ctx, id)
	if err != nil {
		return 0, userDate, mtproto.ErrInternelServerError
	}
	if bcDo == nil || bcDo.Deleted {
		return 0, userDate, mtproto.ErrButtonTypeInvalid
	}
	bmDo, err := m.BlogMomentsDAO.Select(ctx, bcDo.BlogId)
	if err != nil {
		return 0, userDate, mtproto.ErrInternelServerError
	}
	if bmDo == nil || bmDo.Deleted {
		return 0, userDate, mtproto.ErrButtonTypeInvalid
	}
	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		if likeDo == nil {
			do := &dataobject.BlogLikesDO{
				UserId:    fromUserId,
				Type:      BLOG_COMMENT_TYPE_COMMENT,
				BlogId:    bmDo.Id,
				CommentId: id,
				Date:      now,
			}
			lastInsertId, _, err := m.BlogLikesDAO.InsertTx(tx, do)
			if err != nil {
				result.Err = err
				return
			}
			result.Data = lastInsertId
			pushUserId = bcDo.UserId
		} else {
			result.Data = int64(likeDo.Id)
			_, err := m.BlogLikesDAO.UpdateLikeTx(tx, likeDo.Id, liked)
			if err != nil {
				result.Err = err
				return
			}
		}
		_, err := m.BlogCommentsDAO.LikeTx(tx, id, liked)
		if err != nil {
			result.Err = err
			return
		}
		//if liked {
		//	m.BlogsDAO.IncLikes(tx, bmDo.UserId, 1)
		//} else {
		//	m.BlogsDAO.DecLikes(tx, bmDo.UserId, 1)
		//}
	})
	if tR.Err != nil {
		return 0, nil, mtproto.ErrInternelServerError
	}

	if liked {
		likeId := int32(tR.Data.(int64))
		likeDo, _ = m.BlogLikesDAO.Select(ctx, likeId)
		if likeDo != nil {
			userDate = mtproto.MakeTLBlogsUserDate(&mtproto.Blogs_UserDate{
				UserId: likeDo.UserId,
				Date:   likeDo.Date,
			}).To_Blogs_UserDate()
		}
	}
	return pushUserId, userDate, nil
}

func (m *BlogCore) GetSelfHistory(ctx context.Context, fromUserId int32, min_id int32, offset int32, limit int32) (*mtproto.MicroBlogs, error) {
	var (
		bmDos []*dataobject.BlogMomentsDO
		err   error
	)
	count := offset + limit
	if offset >= 0 {
		if min_id == 0 {
			min_id = math.MaxInt32
		}
		bmDos, err = m.BlogMomentsDAO.SelectBackwardBySelf(ctx, fromUserId, min_id, offset+limit)
		if err != nil {
			return nil, mtproto.ErrInternelServerError
		}
	} else {
		var bmDos1 []*dataobject.BlogMomentsDO
		bmDos1, err = m.BlogMomentsDAO.SelectForwardBySelf(ctx, fromUserId, min_id, -offset)
		if err != nil {
			return nil, mtproto.ErrInternelServerError
		}
		//for i, j := 0, len(bmDos1)-1; i < j; i, j = i+1, j-1 {
		//	bmDos1[i], bmDos1[j] = bmDos1[j], bmDos1[i]
		//}
		bmDos = append(bmDos, bmDos1...)
		if count > 0 {
			var bmDos2 []*dataobject.BlogMomentsDO
			bmDos2, err = m.BlogMomentsDAO.SelectBackwardBySelf(ctx, fromUserId, min_id, count)
			if err != nil {
				return nil, mtproto.ErrInternelServerError
			}
			bmDos = append(bmDos, bmDos2...)
		}
	}
	return m.makeMicroBlogs(ctx, fromUserId, bmDos, false)
}

func (m *BlogCore) GetPublicHistory(ctx context.Context, fromUserId int32, min_id int32, offset int32, limit int32, fnIsHisFriend func(int32) bool) (*mtproto.MicroBlogs, error) {
	var (
		bmDos       []*dataobject.BlogMomentsDO
		bmDosSorted []*dataobject.BlogMomentsDO
		err         error
	)
	if offset == 0 {
		bmDosSorted, err = m.BlogMomentsDAO.SelectSortedByPublic(ctx, fromUserId)
		if err != nil {
			log.Errorf("BlogCore.GetPublicHistory, SelectSortedByPublic, error: %v", err)
			return nil, mtproto.ErrInternelServerError
		}
	}
	count := offset + limit
	if offset >= 0 {
		if min_id == 0 {
			min_id = math.MaxInt32
		}
		bmDos, err = m.BlogMomentsDAO.SelectBackwardByPublic(ctx, fromUserId, min_id, count)
		if err != nil {
			log.Errorf("BlogCore.GetPublicHistory, SelectBackwardByPublic, error: %v", err)
			return nil, mtproto.ErrInternelServerError
		}
	} else {
		var bmDos1 []*dataobject.BlogMomentsDO
		bmDos1, err = m.BlogMomentsDAO.SelectForwardByPublic(ctx, fromUserId, min_id, -offset)
		if err != nil {
			log.Errorf("BlogCore.GetPublicHistory, SelectForwardByPublic, error: %v", err)
			return nil, mtproto.ErrInternelServerError
		}
		//for i, j := 0, len(bmDos1)-1; i < j; i, j = i+1, j-1 {
		//	bmDos1[i], bmDos1[j] = bmDos1[j], bmDos1[i]
		//}
		bmDos = append(bmDos, bmDos1...)
		if count > 0 {
			var bmDos2 []*dataobject.BlogMomentsDO
			bmDos2, err = m.BlogMomentsDAO.SelectBackwardByPublic(ctx, fromUserId, min_id, count)
			if err != nil {
				log.Errorf("BlogCore.GetPublicHistory, SelectBackwardByPublic, error: %v", err)
				return nil, mtproto.ErrInternelServerError
			}
			bmDos = append(bmDos, bmDos2...)
		}
	}
	if err != nil {
		log.Errorf("BlogCore.GetPublicHistory, error: %v", err)
		return nil, mtproto.ErrInternelServerError
	}

	bmDos = append(bmDosSorted, bmDos...)

	log.Debugf("[blogsPrivacy] GetPublicHistory, bmDos before filter, %v", bmDos)
	bmDos, err = m.filterBlogMomentsByPrivacy(ctx, fromUserId, bmDos, fnIsHisFriend)
	if err != nil {
		return nil, mtproto.ErrInternelServerError
	}
	//log.Debugf("[blogsPrivacy] GetPublicHistory, bmDos after filter, %v", bmDos)
	return m.makeMicroBlogs(ctx, fromUserId, bmDos, false)
}

func (m *BlogCore) GetUsersHistory(ctx context.Context, fromUserId int32, users []int32, min_id int32, offset int32, limit int32, fnIsHisFriend func(int32) bool) (*mtproto.MicroBlogs, error) {
	var (
		bmDos []*dataobject.BlogMomentsDO
		err   error
	)
	//朋友圈权限，过滤掉"不看他"的
	log.Debugf("[blogsPrivacy] GetUsersHistory, users before filter, %v", users)
	users, err = m.CheckPrivacySkipFromList(ctx, fromUserId, users)
	if err != nil {
		return nil, err
	}
	log.Debugf("[blogsPrivacy] GetUsersHistory, users after filter, %v", users)

	if len(users) == 0 {
		log.Debugf("GetUsersHistory, users is empty after privacy check, fromUserId:%d", fromUserId)
	}

	count := offset + limit
	if offset >= 0 {
		if min_id == 0 {
			min_id = math.MaxInt32
		}
		bmDos, err = m.BlogMomentsDAO.SelectBackwardByUsers(ctx, fromUserId, users, min_id, offset+limit)
		if err != nil {
			return nil, mtproto.ErrInternelServerError
		}
	} else {
		if offset < 0 {
			var bmDos1 []*dataobject.BlogMomentsDO
			bmDos1, err = m.BlogMomentsDAO.SelectForwardByUsers(ctx, fromUserId, users, min_id, -offset)
			if err != nil {
				return nil, mtproto.ErrInternelServerError
			}
			//for i, j := 0, len(bmDos1)-1; i < j; i, j = i+1, j-1 {
			//	bmDos1[i], bmDos1[j] = bmDos1[j], bmDos1[i]
			//}
			bmDos = append(bmDos, bmDos1...)
		}
		if count > 0 {
			var bmDos2 []*dataobject.BlogMomentsDO
			bmDos2, err = m.BlogMomentsDAO.SelectBackwardByUsers(ctx, fromUserId, users, min_id, count)
			if err != nil {
				return nil, mtproto.ErrInternelServerError
			}
			bmDos = append(bmDos, bmDos2...)
		}
	}
	if err != nil {
		return nil, mtproto.ErrInternelServerError
	}

	//log.Debugf("[blogsPrivacy] GetUsersHistory, bmDos before filter, %v", bmDos)
	bmDos, err = m.filterBlogMomentsByPrivacy(ctx, fromUserId, bmDos, fnIsHisFriend)
	if err != nil {
		return nil, mtproto.ErrInternelServerError
	}
	log.Debugf("[blogsPrivacy] GetUsersHistory, bmDos after filter, %v", bmDos)

	return m.makeMicroBlogs(ctx, fromUserId, bmDos, false)
}

func (m *BlogCore) GetTopicHistory(ctx context.Context, fromUserId int32, min_id int32, offset int32, limit int32, topicId int32, tab int32, fnIsHisFriend func(int32) bool) (*mtproto.MicroBlogs, error) {
	//log.Debugf("BlogCore.GetTopicHistory, fromUserId:%d, min_id:%d, offset:%d, limit:%d, topicId:%d, tab:%d", fromUserId, min_id, offset, limit, topicId, tab)
	var (
		bmDos        []*dataobject.BlogMomentsDO
		bmDosLiked   []*dataobject.BlogMomentsDO
		err          error
		excludeLiked = tab == 2 || tab == 3
	)
	if offset == 0 && excludeLiked {
		days := int32(0)
		if tab == 3 {
			days = 1
		}
		bmDosLiked, err = m.BlogMomentsDAO.SelectLikedByTopic(ctx, fromUserId, topicId, days)
		if err != nil {
			log.Errorf("BlogCore.GetTopicHistory, SelectLikedByTopic, error: %v", err)
			return nil, mtproto.ErrInternelServerError
		}
	}

	count := offset + limit
	if offset >= 0 {
		if min_id == 0 {
			min_id = math.MaxInt32
		}
		bmDos, err = m.BlogMomentsDAO.SelectBackwardByTopic(ctx, fromUserId, min_id, count, topicId, excludeLiked)
		//log.Debugf("BlogCore.GetTopicHistory, bmDos1:%v", bmDos)
		if err != nil {
			log.Errorf("BlogCore.GetTopicHistory, SelectBackwardByTopic, error: %v", err)
			return nil, mtproto.ErrInternelServerError
		}
	} else {
		var bmDos1 []*dataobject.BlogMomentsDO
		bmDos1, err = m.BlogMomentsDAO.SelectForwardByTopic(ctx, fromUserId, min_id, -offset, topicId, excludeLiked)
		//log.Debugf("BlogCore.GetTopicHistory, bmDos2:%v", bmDos1)
		if err != nil {
			log.Errorf("BlogCore.GetTopicHistory, SelectForwardByTopic, error: %v", err)
			return nil, mtproto.ErrInternelServerError
		}
		//for i, j := 0, len(bmDos1)-1; i < j; i, j = i+1, j-1 {
		//	bmDos1[i], bmDos1[j] = bmDos1[j], bmDos1[i]
		//}
		bmDos = append(bmDos, bmDos1...)
		//log.Debugf("BlogCore.GetTopicHistory, bmDos3:%v", bmDos)
		if count > 0 {
			var bmDos2 []*dataobject.BlogMomentsDO
			bmDos2, err = m.BlogMomentsDAO.SelectBackwardByTopic(ctx, fromUserId, min_id, count, topicId, excludeLiked)
			//log.Debugf("BlogCore.GetTopicHistory, bmDos4:%v", bmDos2)
			if err != nil {
				log.Errorf("BlogCore.GetTopicHistory, SelectBackwardByTopic, error: %v", err)
				return nil, mtproto.ErrInternelServerError
			}
			bmDos = append(bmDos, bmDos2...)
			//log.Debugf("BlogCore.GetTopicHistory, bmDos5:%v", bmDos)
		}
	}
	if err != nil {
		log.Errorf("BlogCore.GetTopicHistory, error: %v", err)
		return nil, mtproto.ErrInternelServerError
	}

	bmDos = append(bmDosLiked, bmDos...)

	//log.Debugf("[blogsPrivacy] GetTopicHistory, bmDos before filter, %v", bmDos)
	bmDos, err = m.filterBlogMomentsByPrivacy(ctx, fromUserId, bmDos, fnIsHisFriend)
	if err != nil {
		return nil, mtproto.ErrInternelServerError
	}
	//log.Debugf("[blogsPrivacy] GetTopicHistory, bmDos after filter, %v", bmDos)
	return m.makeMicroBlogs(ctx, fromUserId, bmDos, false)
}

func (m *BlogCore) filterBlogMomentsByPrivacy(ctx context.Context, fromUserId int32, bmDos []*dataobject.BlogMomentsDO, fnIsHisFriend func(int32) bool) ([]*dataobject.BlogMomentsDO, error) {
	out := make([]*dataobject.BlogMomentsDO, 0, len(bmDos))
	for _, bm := range bmDos {
		uid := bm.UserId

		// uid 不给 fromUserId 看
		canSee, err := m.CheckPrivacyHideTo(ctx, uid, fromUserId)
		if err != nil {
			return nil, err
		}
		if !canSee {
			continue
		}

		// fromUserId 不看 uid
		canSee, err = m.CheckPrivacySkipFrom(ctx, fromUserId, uid)
		if err != nil {
			return nil, err
		}
		if !canSee {
			continue
		}

		if fnIsHisFriend(uid) {
			// 好友，显示最近几天
			days, err := m.GetPrivacyShowInDays(ctx, uid)
			if err != nil {
				return nil, err
			}
			if days > 0 {
				minDate := int32(time.Now().UTC().Unix()) - days*86400
				if bm.Date < minDate {
					continue
				}
			}
		} else {
			// 非好友，显示最近几条
			counts, err := m.GetPrivacyShowInCounts(ctx, uid)
			if err != nil {
				return nil, err
			}
			if counts > 0 {
				blogMomentCount, err := m.BlogMomentsDAO.CountBlogMomentByUid(ctx, uid)
				if blogMomentCount <= counts {
					continue
				}
				id, err := m.BlogMomentsDAO.SelectIdByOffset(ctx, uid, counts)
				if err != nil {
					return nil, err
				}
				if bm.Id <= id {
					continue
				}
			}
		}

		out = append(out, bm)
	}

	return out, nil
}

func (m *BlogCore) GetUserHistory(ctx context.Context, fromUserId int32, user_id int32, isHisFriend, isMutualFriend bool, min_id int32, offset int32, limit int32) (*mtproto.MicroBlogs, error) {
	var (
		bmDos    []*dataobject.BlogMomentsDO
		err      error
		followed bool
		minDate  uint32
	)

	//判断朋友圈权限，如果"不让他看"返回空
	canSee, err := m.CheckPrivacyHideTo(ctx, user_id, fromUserId)
	if err != nil {
		return nil, err
	}
	if !canSee {
		return m.makeMicroBlogs(ctx, fromUserId, bmDos, false)
	}

	//判断朋友圈权限
	if isHisFriend {
		days, err := m.GetPrivacyShowInDays(ctx, user_id)
		if err != nil {
			return nil, err
		}
		if days == -1 {
			minDate = 0
		} else {
			minDate = uint32(time.Now().UTC().Unix()) - uint32(days*86400)
		}
	} else {
		counts, err := m.GetPrivacyShowInCounts(ctx, user_id)
		if err != nil {
			return nil, err
		}
		if counts != -1 {
			min_id = 0
			offset = 0
			limit = counts
		}
	}

	//获取双方关系
	//isMutualFriend 可获取公开 好友(未过滤)
	bfDo, err := m.BlogFollowsDAO.SelectByUserAndTargetUser(ctx, fromUserId, user_id)
	if err == nil && bfDo != nil && !bfDo.Deleted {
		followed = true
	}
	//followed 可获取公开 粉丝
	visibles := []int8{model.VisibleType_Public}
	if followed {
		visibles = append(visibles, model.VisibleType_Fans)
	}
	if isMutualFriend {
		visibles = append(visibles, []int8{model.VisibleType_Friend, model.VisibleType_Allow, model.VisibleType_NotAllow}...)
	}

	count := offset + limit
	if offset >= 0 {
		if min_id == 0 {
			min_id = math.MaxInt32
		}
		bmDos, err = m.BlogMomentsDAO.SelectBackwardByUser(ctx, fromUserId, user_id, visibles, min_id, offset+limit, minDate)
		if err != nil {
			return nil, mtproto.ErrInternelServerError
		}
	} else {
		if offset < 0 {
			var bmDos1 []*dataobject.BlogMomentsDO
			bmDos1, err = m.BlogMomentsDAO.SelectForwardByUser(ctx, fromUserId, user_id, visibles, min_id, -offset, minDate)
			if err != nil {
				return nil, mtproto.ErrInternelServerError
			}
			bmDos = append(bmDos, bmDos1...)
		}
		if count > 0 {
			var bmDos2 []*dataobject.BlogMomentsDO
			bmDos2, err = m.BlogMomentsDAO.SelectBackwardByUser(ctx, fromUserId, user_id, visibles, min_id, count, minDate)
			if err != nil {
				return nil, mtproto.ErrInternelServerError
			}
			bmDos = append(bmDos, bmDos2...)
		}
	}
	if err != nil {
		return nil, mtproto.ErrInternelServerError
	}
	return m.makeMicroBlogs(ctx, fromUserId, bmDos, false)
}

func (m *BlogCore) GetBlogs(ctx context.Context, fromUserId int32, ids []int32) (*mtproto.MicroBlogs, error) {
	var (
		bmDos []*dataobject.BlogMomentsDO
		err   error
	)

	bmDos, err = m.BlogMomentsDAO.SelectList(ctx, ids)
	if err != nil {
		return nil, mtproto.ErrInternelServerError
	}
	return m.makeMicroBlogs(ctx, fromUserId, bmDos, false)
}

func (m *BlogCore) GmGetBlogs(ctx context.Context, ids []int32) (*mtproto.MicroBlogs, error) {
	var (
		bmDos []*dataobject.BlogMomentsDO
		err   error
	)

	bmDos, err = m.BlogMomentsDAO.SelectListIgnoreDeletedFlag(ctx, ids)
	if err != nil {
		return nil, mtproto.ErrInternelServerError
	}
	return m.makeMicroBlogs(ctx, 0, bmDos, true)
}

func (m *BlogCore) GetBlogUser(ctx context.Context, fromUserId int32, userId int32) (*mtproto.Blogs_User, error) {
	blogsDo, err := m.BlogsDAO.SelectByUser(ctx, userId)
	if err != nil {
		return nil, mtproto.ErrInternelServerError
	}
	user := mtproto.MakeTLBlogsUser(&mtproto.Blogs_User{
		UserId:   userId,
		Followed: false,
	}).To_Blogs_User()
	if blogsDo != nil {
		user.Blogs = blogsDo.Moments
		user.Follows = blogsDo.Follows
		user.Fans = blogsDo.Fans
		user.Likes = blogsDo.Likes
	}
	if fromUserId == userId {
		user.State = mtproto.MakeTLBlogsState(&mtproto.Blogs_State{
			Pts:           int32(idgen.CurrentBlogPtsId(ctx, fromUserId)),
			ReadBlogId:    blogsDo.BlogMaxId,
			ReadCommentId: blogsDo.CommentMaxId,
			Date:          int32(time.Now().Unix()),
		}).To_Blogs_State()
		if user.State.Pts > blogsDo.UpdateMaxId {
			user.State.UnreadCount = user.State.Pts - blogsDo.UpdateMaxId
		}
	} else {
		bfDo, err := m.BlogFollowsDAO.SelectByUserAndTargetUser(ctx, fromUserId, userId)
		if err == nil && bfDo != nil && !bfDo.Deleted {
			user.Followed = true
		}
	}
	return user, nil
}

func (m *BlogCore) makeComments(ctx context.Context, fromUserId int32, bcDos []*dataobject.BlogCommentsDO) (*mtproto.Blogs_Comments, error) {
	var (
		result  *mtproto.Blogs_Comments = mtproto.MakeTLBlogsCommentNotModified(nil).To_Blogs_Comments()
		doCount                         = len(bcDos)
		blogId  *mtproto.Blogs_IdType
	)
	if doCount == 0 {
		return result, nil
	}
	result = result.To_BlogsComments().To_Blogs_Comments()
	result.Comments = make([]*mtproto.Blogs_Comment, 0, doCount)
	var commentIds []int32 = make([]int32, 0, doCount)
	for _, do := range bcDos {
		commentIds = append(commentIds, do.Id)
		blogId = mtproto.MakeTLBlogsIdTypeBlog(&mtproto.Blogs_IdType{
			BlogId: do.BlogId,
		}).To_Blogs_IdType()
		if do.Type == BLOG_COMMENT_TYPE_COMMENT {
			blogId.CommentId = do.CommentId
			blogId = blogId.To_BlogsIdTypeComment().To_Blogs_IdType()
		}
		comment := mtproto.MakeTLBlogsComment(&mtproto.Blogs_Comment{
			BlogId: blogId,
			Id:     do.Id,
			Text:   do.Text,
			UserId: do.UserId,
			Date:   do.Date,
		}).To_Blogs_Comment()

		comment.LikeCount = int32(m.CommonDAO.CalcSize(ctx, "blog_likes", map[string]interface{}{
			"type":       BLOG_COMMENT_TYPE_COMMENT,
			"comment_id": do.Id,
			"deleted":    0,
		}))
		commentCount := int32(m.CommonDAO.CalcSize(ctx, "blog_comments", map[string]interface{}{
			"type":       BLOG_COMMENT_TYPE_COMMENT,
			"comment_id": do.Id,
			"deleted":    0,
		}))
		if commentCount > 0 {
			comment.CommentCount = &types.Int32Value{Value: commentCount}
		}
		if do.ReplyId != 0 {
			comment.ReplyId = &types.Int32Value{Value: do.ReplyId}
		}
		result.Comments = append(result.Comments, comment)
	}
	if len(commentIds) > 0 {
		blDos, err := m.BlogLikesDAO.SelectByUserAndCommentIds(ctx, fromUserId, commentIds)
		if err != nil {
			return result, mtproto.ErrInternelServerError
		}
		for _, r := range result.Comments {
			for _, bl := range blDos {
				if r.Id == bl.CommentId {
					r.Liked = true
					r.LikeDate = &types.Int32Value{Value: bl.Date}
				}
			}
		}
	}
	return result, nil
}

func (m *BlogCore) GetBlogComments(ctx context.Context, fromUserId, id, offsetId, limit int32) (*mtproto.Blogs_Comments, error) {
	result := mtproto.MakeTLBlogsCommentNotModified(nil).To_Blogs_Comments()
	if offsetId == 0 {
		offsetId = math.MaxInt32
	}
	bcDos, err := m.BlogCommentsDAO.SelectListByBlogId(ctx, id, offsetId, limit)
	if err != nil {
		return result, mtproto.ErrInternelServerError
	}
	result.Count = int32(m.CommonDAO.CalcSize(ctx, "blog_comments", map[string]interface{}{
		"type":    BLOG_COMMENT_TYPE_BLOG,
		"blog_id": id,
		"deleted": 0,
	}))
	return m.makeComments(ctx, fromUserId, bcDos)
}

func (m *BlogCore) GetCommentAndReply(ctx context.Context, fromUserId, id, offsetId, limit int32) (*mtproto.Blogs_Comments, error) {
	result := mtproto.MakeTLBlogsCommentNotModified(nil).To_Blogs_Comments()
	bcDos, err := m.BlogCommentsDAO.SelectListByCommentId(ctx, id, offsetId, limit)
	if err != nil {
		return result, mtproto.ErrInternelServerError
	}
	result.Count = int32(m.CommonDAO.CalcSize(ctx, "blog_comments", map[string]interface{}{
		"type":       BLOG_COMMENT_TYPE_COMMENT,
		"comment_id": id,
		"deleted":    0,
	}))
	return m.makeComments(ctx, fromUserId, bcDos)
}

func (m *BlogCore) GetCommentList(ctx context.Context, fromUserId int32, ids []int32) (*mtproto.Blogs_Comments, error) {
	result := mtproto.MakeTLBlogsCommentNotModified(nil).To_Blogs_Comments()
	bcDos, err := m.BlogCommentsDAO.SelectList(ctx, ids)
	if err != nil {
		return result, mtproto.ErrInternelServerError
	}
	return m.makeComments(ctx, fromUserId, bcDos)
}

func (m *BlogCore) GetBlogLikes(ctx context.Context, fromUserId, id, offsetId, limit int32) (*mtproto.Blogs_UserDates, error) {
	result := mtproto.MakeTLBlogsUserDatesNotModified(nil).To_Blogs_UserDates()
	blDos, err := m.BlogLikesDAO.SelectListByBlogId(ctx, id, offsetId, limit)
	if err != nil {
		return result, mtproto.ErrInternelServerError
	}
	result.Count = int32(m.CommonDAO.CalcSize(ctx, "blog_likes", map[string]interface{}{
		"type":    BLOG_COMMENT_TYPE_BLOG,
		"blog_id": id,
		"deleted": 0,
	}))
	if len(blDos) > 0 {
		result = result.To_BlogsUserDates().To_Blogs_UserDates()
		result.Users = make([]*mtproto.Blogs_UserDate, 0, len(blDos))
		for _, bl := range blDos {
			result.Users = append(result.Users, mtproto.MakeTLBlogsUserDate(&mtproto.Blogs_UserDate{
				UserId: bl.UserId,
				Date:   bl.Date,
			}).To_Blogs_UserDate())
		}
	}
	return result, nil
}

func (m *BlogCore) GetCommentLikes(ctx context.Context, fromUserId, id, offsetId, limit int32) (*mtproto.Blogs_UserDates, error) {
	result := mtproto.MakeTLBlogsUserDatesNotModified(nil).To_Blogs_UserDates()
	blDos, err := m.BlogLikesDAO.SelectListByCommentId(ctx, id, offsetId, limit)
	if err != nil {
		return result, mtproto.ErrInternelServerError
	}
	result.Count = int32(m.CommonDAO.CalcSize(ctx, "blog_likes", map[string]interface{}{
		"type":       BLOG_COMMENT_TYPE_COMMENT,
		"comment_id": id,
		"deleted":    0,
	}))
	if len(blDos) > 0 {
		result = result.To_BlogsUserDates().To_Blogs_UserDates()
		result.Users = make([]*mtproto.Blogs_UserDate, 0, len(blDos))
		for _, bl := range blDos {
			result.Users = append(result.Users, mtproto.MakeTLBlogsUserDate(&mtproto.Blogs_UserDate{
				UserId: bl.UserId,
				Date:   bl.Date,
			}).To_Blogs_UserDate())
		}
	}
	return result, nil
}

func (m *BlogCore) GetGroupTags(ctx context.Context, fromUserId int32) (*mtproto.Blogs_GroupTags, error) {
	bgDos, err := m.BlogGroupTagsDAO.SelectByUser(ctx, fromUserId)
	if err != nil {
		return nil, mtproto.ErrInternelServerError
	}
	var result *mtproto.Blogs_GroupTags = mtproto.MakeTLBlogsGroupTags(&mtproto.Blogs_GroupTags{
		Count: int32(len(bgDos)),
		Tags:  make([]*mtproto.Blogs_GroupTag, 0, len(bgDos)),
	}).To_Blogs_GroupTags()
	for _, bg := range bgDos {
		userIds := make([]int32, 0)
		if len(bg.MemberUserIds) > 0 {
			json.Unmarshal([]byte(bg.MemberUserIds), &userIds)
		}
		result.Tags = append(result.Tags, mtproto.MakeTLBlogsGroupTag(&mtproto.Blogs_GroupTag{
			TagId: bg.Id,
			Title: bg.Title,
			Users: userIds,
			Date:  bg.Date,
		}).To_Blogs_GroupTag())
	}
	return result, nil
}

func (m *BlogCore) GetFollows(ctx context.Context, fromUserId int32, userId int32, offset int32, limit int32) (*mtproto.Blogs_UserDates, error) {
	var result *mtproto.Blogs_UserDates = mtproto.MakeTLBlogsUserDatesNotModified(nil).To_Blogs_UserDates()
	bfDos, err := m.BlogFollowsDAO.SelectByUser(ctx, userId, offset, limit)
	if err != nil {
		return result, mtproto.ErrInternelServerError
	}
	result.Count = int32(m.CommonDAO.CalcSize(ctx, "blog_follows", map[string]interface{}{
		"user_id": userId,
		"deleted": 0,
	}))
	if len(bfDos) > 0 {
		result = result.To_BlogsUserDates().To_Blogs_UserDates()
		result.Users = make([]*mtproto.Blogs_UserDate, 0, len(bfDos))
		for _, bf := range bfDos {
			result.Users = append(result.Users, mtproto.MakeTLBlogsUserDate(&mtproto.Blogs_UserDate{
				UserId: bf.TargetUserId,
				Date:   bf.Date,
			}).To_Blogs_UserDate())
		}

	}
	return result, nil
}

func (m *BlogCore) GetFans(ctx context.Context, fromUserId int32, userId int32, offset int32, limit int32) (*mtproto.Blogs_UserDates, error) {
	var result *mtproto.Blogs_UserDates = mtproto.MakeTLBlogsUserDatesNotModified(nil).To_Blogs_UserDates()
	bfDos, err := m.BlogFollowsDAO.SelectByTargetUser(ctx, userId, offset, limit)
	if err != nil {
		return result, mtproto.ErrInternelServerError
	}
	result.Count = int32(m.CommonDAO.CalcSize(ctx, "blog_follows", map[string]interface{}{
		"target_uid": userId,
		"deleted":    0,
	}))
	if len(bfDos) > 0 {
		result = result.To_BlogsUserDates().To_Blogs_UserDates()
		for _, bf := range bfDos {
			result.Users = append(result.Users, mtproto.MakeTLBlogsUserDate(&mtproto.Blogs_UserDate{
				UserId: bf.UserId,
				Date:   bf.Date,
			}).To_Blogs_UserDate())
		}
	}
	return result, nil
}

func (m *BlogCore) CreateGroupTag(ctx context.Context, fromUserId int32, title string, userIds []int32) (*mtproto.Blogs_GroupTag, error) {
	var now = int32(time.Now().Unix())
	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		b, _ := json.Marshal(userIds)
		blogDO := &dataobject.BlogGroupTagsDO{
			UserId:        fromUserId,
			Title:         title,
			MemberUserIds: string(b),
			Date:          now,
		}
		lastInsertId, _, err := m.BlogGroupTagsDAO.InsertTx(tx, blogDO)
		if err != nil {
			result.Err = err
			return
		}

		if lastInsertId > 0 {
			result.Data = lastInsertId
			return
		} else {
			result.Err = fmt.Errorf("insert error")
			return
		}
	})

	if tR.Err != nil {
		return nil, tR.Err
	}

	groupTag := mtproto.MakeTLBlogsGroupTag(&mtproto.Blogs_GroupTag{
		TagId: int32(tR.Data.(int64)),
		Title: title,
		Users: userIds,
		Date:  now,
	}).To_Blogs_GroupTag()
	return groupTag, nil
}

func (m *BlogCore) AddGroupTagMember(ctx context.Context, fromUserId int32, tagId int32, userIds []int32) ([]int32, error) {
	var newIds []int32 = make([]int32, 0)
	bgtDos, err := m.BlogGroupTagsDAO.SelectByUserAndTags(ctx, fromUserId, []int32{tagId})
	if err != nil {
		return newIds, mtproto.ErrInternelServerError
	}
	if len(bgtDos) == 0 {
		return newIds, mtproto.ErrButtonTypeInvalid
	}

	bgtDo := bgtDos[0]
	var existIds []int32 = make([]int32, 0)
	newIds = make([]int32, 0, len(userIds))
	if len(bgtDo.MemberUserIds) > 0 {
		json.Unmarshal([]byte(bgtDo.MemberUserIds), &existIds)
		for _, uid := range userIds {
			exist := false
			for _, existid := range existIds {
				if existid == uid {
					exist = true
				}
			}
			if !exist {
				newIds = append(newIds, uid)
			}
		}
		existIds = append(existIds, newIds...)
	}
	b, _ := json.Marshal(existIds)
	bgtDo.MemberUserIds = string(b)

	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		rowsAffected, err := m.BlogGroupTagsDAO.UpdateMembersTx(tx, fromUserId, bgtDo.Id, bgtDo.MemberUserIds)
		if err != nil {
			result.Err = err
			return
		}
		result.Data = rowsAffected
	})

	if tR.Err != nil {
		return nil, tR.Err
	}
	return newIds, nil
}

func (m *BlogCore) DeleteGroupTagMember(ctx context.Context, fromUserId int32, tagId int32, userIds []int32) ([]int32, error) {
	var deleteIds []int32 = make([]int32, 0)
	bgtDos, err := m.BlogGroupTagsDAO.SelectByUserAndTags(ctx, fromUserId, []int32{tagId})
	if err != nil {
		return deleteIds, mtproto.ErrInternelServerError
	}
	if len(bgtDos) == 0 {
		return deleteIds, mtproto.ErrButtonTypeInvalid
	}

	bgtDo := bgtDos[0]
	var newExistIds []int32 = make([]int32, 0)
	deleteIds = make([]int32, 0, len(userIds))
	if len(bgtDo.MemberUserIds) > 0 {
		var existIds []int32 = make([]int32, 0)
		json.Unmarshal([]byte(bgtDo.MemberUserIds), &existIds)
		for _, existid := range existIds {
			exist := false
			for _, uid := range userIds {
				if existid == uid {
					exist = true
				}
			}
			if exist {
				deleteIds = append(deleteIds, existid)
			} else {
				newExistIds = append(newExistIds, existid)
			}
		}
	}
	b, _ := json.Marshal(newExistIds)
	bgtDo.MemberUserIds = string(b)

	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		rowsAffected, err := m.BlogGroupTagsDAO.UpdateMembersTx(tx, fromUserId, bgtDo.Id, bgtDo.MemberUserIds)
		if err != nil {
			result.Err = err
			return
		}
		result.Data = rowsAffected
	})

	if tR.Err != nil {
		return nil, mtproto.ErrInternelServerError
	}
	return deleteIds, nil
}

func (m *BlogCore) DeleteGroupTag(ctx context.Context, fromUserId int32, Ids []int32) ([]int32, error) {
	deleteIds := make([]int32, 0)
	bgtDos, err := m.BlogGroupTagsDAO.SelectByUserAndTags(ctx, fromUserId, Ids)
	if err != nil {
		return deleteIds, mtproto.ErrInternelServerError
	}
	if len(bgtDos) == 0 {
		return deleteIds, mtproto.ErrButtonTypeInvalid
	}

	deleteIds = make([]int32, 0, len(bgtDos))
	for _, bgt := range bgtDos {
		deleteIds = append(deleteIds, bgt.Id)
	}

	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		rowsAffected, err := m.BlogGroupTagsDAO.DeleteGroupTags(tx, fromUserId, deleteIds)
		if err != nil {
			result.Err = err
			return
		}
		result.Data = rowsAffected
	})
	if tR.Err != nil {
		return nil, mtproto.ErrInternelServerError
	}
	return deleteIds, nil
}

func (m *BlogCore) EditGroupTag(ctx context.Context, fromUserId int32, tagId int32, title string) (bool, error) {
	bgtDo, err := m.BlogGroupTagsDAO.Select(ctx, tagId)
	if err != nil {
		return false, mtproto.ErrInternelServerError
	}
	if bgtDo == nil || bgtDo.UserId != fromUserId || bgtDo.Deleted {
		return false, mtproto.ErrButtonTypeInvalid
	}

	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		rowsAffected, err := m.BlogGroupTagsDAO.UpdateTx(tx, bgtDo.Id, map[string]interface{}{
			"title": title,
		})
		if err != nil {
			result.Err = err
			return
		}
		result.Data = rowsAffected
	})
	if tR.Err != nil {
		return false, mtproto.ErrInternelServerError
	}
	return true, nil
}

func (m *BlogCore) SendBlogComment(ctx context.Context, fromUserId int32, id int32, text string) (int32, *mtproto.TLBlogsComment, error) {
	var (
		now       = int32(time.Now().Unix())
		uid int32 = 0
	)
	bmDo, err := m.BlogMomentsDAO.Select(ctx, id)
	if err != nil {
		return uid, nil, mtproto.ErrInternelServerError
	}
	if bmDo == nil || bmDo.Deleted {
		return uid, nil, mtproto.ErrButtonTypeInvalid
	}
	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		do := &dataobject.BlogCommentsDO{
			UserId: fromUserId,
			Type:   BLOG_COMMENT_TYPE_BLOG,
			Text:   text,
			BlogId: bmDo.Id,
			Date:   now,
		}
		lastInsertId, _, err := m.BlogCommentsDAO.InsertTx(tx, do)
		if err != nil {
			result.Err = err
			return
		}

		if lastInsertId > 0 {
			m.BlogsDAO.IncComments(tx, bmDo.UserId, 1)
			result.Data = lastInsertId
			return
		} else {
			result.Err = fmt.Errorf("insert error")
			return
		}
	})

	if tR.Err != nil {
		return uid, nil, mtproto.ErrInternelServerError
	}
	replyId := int32(tR.Data.(int64))
	bcDo, err := m.BlogCommentsDAO.Select(ctx, replyId)
	if err != nil {
		return uid, nil, mtproto.ErrInternelServerError
	}
	uid = bmDo.UserId
	reply := mtproto.MakeTLBlogsComment(&mtproto.Blogs_Comment{
		BlogId: mtproto.MakeTLBlogsIdTypeBlog(&mtproto.Blogs_IdType{BlogId: id}).To_Blogs_IdType(),
		Id:     bcDo.Id,
		Text:   bcDo.Text,
		UserId: bcDo.UserId,
		Date:   bcDo.Date,
	})
	return uid, reply, nil
}

func (m *BlogCore) SendCommentReply(ctx context.Context, fromUserId int32, id int32, text string) (int32, *mtproto.TLBlogsComment, error) {
	var (
		now              = int32(time.Now().Unix())
		targetUser int32 = 0
		commentId  int32
		replyId    int32
	)
	bcDo, err := m.BlogCommentsDAO.Select(ctx, id)
	if err != nil {
		return 0, nil, mtproto.ErrInternelServerError
	}
	if bcDo == nil || bcDo.Deleted {
		log.Errorf("SendCommentReply, bcDo invalid")
		return 0, nil, mtproto.ErrButtonTypeInvalid
	}

	bmDo, err := m.BlogMomentsDAO.Select(ctx, bcDo.BlogId)
	if err != nil {
		return 0, nil, mtproto.ErrInternelServerError
	}
	if bmDo == nil || bmDo.Deleted {
		log.Errorf("SendCommentReply, bmDo invalid")
		return 0, nil, mtproto.ErrButtonTypeInvalid
	}
	targetUser = bmDo.UserId
	commentId = bcDo.Id
	if bcDo.Type == BLOG_COMMENT_TYPE_COMMENT {
		replyId = bcDo.Id
		commentId = bcDo.CommentId
		targetUser = bcDo.UserId
	}
	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		do := &dataobject.BlogCommentsDO{
			UserId:    fromUserId,
			Type:      BLOG_COMMENT_TYPE_COMMENT,
			Text:      text,
			BlogId:    bmDo.Id,
			CommentId: commentId,
			Date:      now,
		}
		if replyId > 0 {
			do.ReplyId = replyId
		}
		lastInsertId, _, err := m.BlogCommentsDAO.InsertTx(tx, do)
		if err != nil {
			result.Err = err
			return
		}

		if lastInsertId > 0 {
			result.Data = lastInsertId
			return
		} else {
			result.Err = fmt.Errorf("insert error")
			return
		}
	})

	if tR.Err != nil {
		return 0, nil, mtproto.ErrInternelServerError
	}
	commentId = int32(tR.Data.(int64))
	bcDo2, err := m.BlogCommentsDAO.Select(ctx, commentId)
	if err != nil {
		return 0, nil, mtproto.ErrInternelServerError
	}
	comment := mtproto.MakeTLBlogsComment(&mtproto.Blogs_Comment{
		BlogId: mtproto.MakeTLBlogsIdTypeComment(&mtproto.Blogs_IdType{CommentId: id}).To_Blogs_IdType(),
		Id:     bcDo2.Id,
		Text:   bcDo2.Text,
		UserId: bcDo2.UserId,
		Date:   bcDo2.Date,
	})
	return targetUser, comment, nil
}

func (m *BlogCore) Follow(ctx context.Context, fromUserId int32, userId int32, followed bool) (bool, *mtproto.Blogs_UserDate, error) {
	var now = int32(time.Now().Unix())
	var (
		follow *mtproto.Blogs_UserDate
		ok     bool = false
	)
	bfDo, err := m.BlogFollowsDAO.SelectByUserAndTargetUser(ctx, fromUserId, userId)
	if err != nil {
		return false, follow, mtproto.ErrInternelServerError
	}
	if bfDo == nil && !followed {
		return false, follow, mtproto.ErrButtonTypeInvalid
	}

	if bfDo != nil {
		follow = mtproto.MakeTLBlogsUserDate(&mtproto.Blogs_UserDate{
			UserId: bfDo.UserId,
			Date:   bfDo.Date,
		}).To_Blogs_UserDate()
		if bfDo.Deleted != followed {
			return false, follow, nil
		}
	}

	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		if bfDo == nil {
			do := &dataobject.BlogFollowsDO{
				UserId:       fromUserId,
				TargetUserId: userId,
				Date:         now,
			}
			lastInsertId, _, err := m.BlogFollowsDAO.InsertTx(tx, do)
			if err != nil {
				result.Err = err
				return
			}

			if lastInsertId > 0 {
				m.BlogsDAO.IncFollows(tx, do.UserId, 1)
				m.BlogsDAO.IncFans(tx, do.TargetUserId, 1)
				result.Data = lastInsertId
				ok = true
				return
			} else {
				result.Err = fmt.Errorf("insert error")
				return
			}
		} else {
			result.Data = int64(bfDo.Id)
			_, err := m.BlogFollowsDAO.UpdateFollowTx(tx, fromUserId, bfDo.Id, followed)
			if err != nil {
				result.Err = err
				return
			}
			if followed {
				m.BlogsDAO.IncFollows(tx, bfDo.UserId, 1)
				m.BlogsDAO.IncFans(tx, bfDo.TargetUserId, 1)
			} else {
				m.BlogsDAO.DecFollows(tx, bfDo.UserId, 1)
				m.BlogsDAO.DecFans(tx, bfDo.TargetUserId, 1)
			}
		}
	})

	if tR.Err != nil {
		return false, nil, mtproto.ErrInternelServerError
	}

	if followed {
		fId := int32(tR.Data.(int64))
		bfDo, _ = m.BlogFollowsDAO.Select(ctx, fId)
		if bfDo != nil {
			follow = mtproto.MakeTLBlogsUserDate(&mtproto.Blogs_UserDate{
				UserId: bfDo.UserId,
				Date:   bfDo.Date,
			}).To_Blogs_UserDate()
		}
	}

	return ok, follow, nil
}

func (m *BlogCore) DeleteBlog(ctx context.Context, fromUserId int32, Ids []int32) ([]int32, []int32, error) {
	var now = int32(time.Now().Unix())
	deleteIds := []int32{}
	shieldIds := []int32{}
	bmDos, err := m.BlogMomentsDAO.SelectList(ctx, Ids)
	if err != nil {
		return deleteIds, shieldIds, mtproto.ErrInternelServerError
	}
	if len(bmDos) == 0 {
		return deleteIds, shieldIds, mtproto.ErrButtonTypeInvalid
	}
	deleteIds = make([]int32, 0, len(bmDos))
	shieldIds = make([]int32, 0, len(bmDos))
	for _, bm := range bmDos {
		if bm.UserId == fromUserId {
			deleteIds = append(deleteIds, bm.Id)
		} else {
			shieldIds = append(shieldIds, bm.Id)
		}
	}
	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		if len(deleteIds) > 0 {
			rowsAffected, err := m.BlogMomentsDAO.DeleteBlog(tx, fromUserId, deleteIds)
			if err != nil {
				result.Err = err
				return
			}
			result.Data = rowsAffected
			m.BlogsDAO.DecMoments(tx, fromUserId, int32(rowsAffected))
			//是否还需要删除所有评论和点赞
		}
		for _, shieldId := range shieldIds {
			do := &dataobject.BlogMomentDeletesDO{
				UserId: fromUserId,
				BlogId: shieldId,
				Date:   now,
			}

			lastInsertId, _, err := m.BlogMomentDeletesDAO.InsertTx(tx, do)
			if err != nil {
				result.Err = err
				return
			}

			if lastInsertId <= 0 {
				result.Err = fmt.Errorf("insert error")
			}
		}
	})
	if tR.Err != nil {
		return deleteIds, shieldIds, mtproto.ErrInternelServerError
	}
	return deleteIds, shieldIds, nil
}

func (m *BlogCore) GetAllFans(ctx context.Context, me int32) []int32 {
	var uids []int32
	bfDos, err := m.BlogFollowsDAO.SelectAllByTargetUser(ctx, me)
	if err != nil {
		return uids
	}
	uids = make([]int32, 0, len(bfDos))
	for _, bf := range bfDos {
		if !bf.Deleted {
			uids = append(uids, bf.UserId)
		}
	}
	return uids
}

func (m *BlogCore) GetAllFollows(ctx context.Context, me int32) []int32 {
	var uids []int32
	bfDos, err := m.BlogFollowsDAO.SelectAllByUser(ctx, me)
	if err != nil {
		return uids
	}
	uids = make([]int32, 0, len(bfDos))
	for _, bf := range bfDos {
		if !bf.Deleted {
			uids = append(uids, bf.TargetUserId)
		}
	}
	return uids
}

func (m *BlogCore) ReadMaxUpdateId(ctx context.Context, me int32, maxId int32) (bool, error) {
	bDo, err := m.BlogsDAO.SelectByUser(ctx, me)
	if err != nil {
		return false, mtproto.ErrInternelServerError
	}
	if bDo == nil {
		return false, mtproto.ErrButtonTypeInvalid
	}
	if maxId > bDo.UpdateMaxId {
		tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
			rowsAffected, err := m.BlogsDAO.UpdateTx(tx, bDo.Id, map[string]interface{}{
				"read_update_id": maxId,
			})
			if err != nil {
				result.Err = err
				return
			}
			result.Data = rowsAffected
		})
		if tR.Err != nil {
			return false, mtproto.ErrInternelServerError
		}
	}
	return true, nil
}

func (m *BlogCore) ReadMaxBlogId(ctx context.Context, me int32, maxId int32) (*mtproto.Blogs_IdType, error) {
	bDo, err := m.BlogsDAO.SelectByUser(ctx, me)
	if err != nil {
		return nil, mtproto.ErrInternelServerError
	}
	bmDo, err := m.BlogMomentDeletesDAO.Select(ctx, maxId)
	if err != nil {
		return nil, mtproto.ErrInternelServerError
	}
	if bDo == nil || bmDo == nil {
		return nil, mtproto.ErrButtonTypeInvalid
	}
	if maxId > bDo.BlogMaxId {
		tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
			rowsAffected, err := m.BlogsDAO.UpdateTx(tx, bDo.Id, map[string]interface{}{
				"read_blog_id": maxId,
			})
			if err != nil {
				result.Err = err
				return
			}
			result.Data = rowsAffected
		})
		if tR.Err != nil {
			return nil, mtproto.ErrInternelServerError
		}
	}
	return mtproto.MakeTLBlogsIdTypeBlog(&mtproto.Blogs_IdType{BlogId: maxId}).To_Blogs_IdType(), nil
}

func (m *BlogCore) ReadMaxCommentId(ctx context.Context, me int32, maxId int32) (*mtproto.Blogs_IdType, error) {
	bDo, err := m.BlogsDAO.SelectByUser(ctx, me)
	if err != nil {
		return nil, mtproto.ErrInternelServerError
	}
	bcDo, err := m.BlogCommentsDAO.Select(ctx, maxId)
	if err != nil {
		return nil, mtproto.ErrInternelServerError
	}
	if bDo == nil || bcDo == nil {
		return nil, mtproto.ErrButtonTypeInvalid
	}
	if maxId > bDo.CommentMaxId {
		tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
			rowsAffected, err := m.BlogsDAO.UpdateTx(tx, bDo.Id, map[string]interface{}{
				"read_comment_id": maxId,
			})
			if err != nil {
				result.Err = err
				return
			}
			result.Data = rowsAffected
		})
		if tR.Err != nil {
			return nil, mtproto.ErrInternelServerError
		}
	}
	return mtproto.MakeTLBlogsIdTypeComment(&mtproto.Blogs_IdType{CommentId: maxId}).To_Blogs_IdType(), nil
}

func (m *BlogCore) getUpdateListByGtPts(ctx context.Context, userId, pts int32, limit int32) []*mtproto.Update {
	doList, _ := m.BlogPtsUpdatesDAO.SelectByGtPts(ctx, userId, pts, limit)
	if len(doList) == 0 {
		return []*mtproto.Update{}
	}
	updates := make([]*mtproto.Update, 0, len(doList))
	for _, do := range doList {
		update := &mtproto.Update{}
		err := json.Unmarshal([]byte(do.UpdateData), update)
		if err != nil {
			log.Errorf("unmarshal pts's update(%d)error: %v", do.Id, err)
			continue
		}
		if m.getUpdateType(update) != do.UpdateType {
			log.Errorf("update data error.")
			continue
		}
		updates = append(updates, update)
	}
	return updates
}

func (m *BlogCore) GetUnreads(ctx context.Context, fromUserId int32, pts, limit int32) (*mtproto.Blogs_Unreads, error) {
	var (
		unreads *mtproto.Blogs_Unreads
	)

	me, err := m.GetBlogUser(ctx, fromUserId, fromUserId)
	if err != nil {
		return nil, err
	}

	remain := me.State.Pts - pts
	if limit > 1000 || remain > limit {
		return mtproto.MakeTLBlogsUnreadTooLong(&mtproto.Blogs_Unreads{
			State: me.State,
		}).To_Blogs_Unreads(), nil
	}

	updateList := m.getUpdateListByGtPts(ctx, fromUserId, pts, limit)
	if len(updateList) == 0 {
		unreads = mtproto.MakeTLBlogsUnreadEmpty(&mtproto.Blogs_Unreads{
			Pts: me.State.GetPts(),
		}).To_Blogs_Unreads()
	} else {
		unreads = mtproto.MakeTLBlogsUnreads(&mtproto.Blogs_Unreads{
			NewUpdates: updateList,
			Users:      []*mtproto.User{},
			State:      me.State,
		}).To_Blogs_Unreads()
	}

	return unreads, nil
}

func (m *BlogCore) CanReward(ctx context.Context, fromUserId int32, userId int32, blogId int32) (bool, error) {
	bmDo, err := m.BlogMomentsDAO.Select(ctx, blogId)
	if err != nil {
		return false, mtproto.ErrInternelServerError
	}
	if bmDo == nil || bmDo.Deleted {
		return false, mtproto.ErrButtonTypeInvalid
	}
	if bmDo.UserId != fromUserId {
		return false, mtproto.ErrButtonTypeInvalid
	}
	return true, nil
}
