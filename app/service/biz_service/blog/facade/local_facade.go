package blog_facade

import (
	"context"
	"encoding/json"
	"open.chat/app/service/biz_service/blog/internal/core"
	"open.chat/app/service/biz_service/blog/internal/dao"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

type localBlogFacade struct {
	*core.BlogCore
}

func New() BlogFacade {
	return &localBlogFacade{
		BlogCore: core.New(dao.New()),
	}
}

func (b *localBlogFacade) GetGroupMembers(ctx context.Context, fromUserId int32, ids []int32) []int32 {
	var userIds model.IDList
	bgtDo, err := b.BlogCore.BlogGroupTagsDAO.SelectByUserAndTags(ctx, fromUserId, ids)
	if err != nil {
		return userIds
	}
	for _, bgt := range bgtDo {
		var uids []int32
		if len(bgt.MemberUserIds) > 0 {
			json.Unmarshal([]byte(bgt.MemberUserIds), &uids)
		}
		userIds.AddIfNot(uids...)
	}
	return userIds
}

func (b *localBlogFacade) SendBlog(ctx context.Context, fromUserId int32, randomId int64, visible *model.VisibleTypeUtil, text string, entities []*mtproto.MessageEntity, mentionUserIds []int32, content *mtproto.BlogContent, geo *mtproto.BlogGeoPoint, topicIds []int32) (model.IDList, *mtproto.MicroBlog, error) {
	//, fnMutualUsers func() (model.UserHelper, []int32)
	var (
		blog                *mtproto.MicroBlog
		err                 error
		hasDuplicateMessage bool
		pushUsers           model.IDList
	)

	hasDuplicateMessage, err = b.BlogCore.HasDuplicateBlog(ctx, fromUserId, randomId)
	if err != nil {
		log.Errorf("checkDuplicateMessage error - %v", err)
		return pushUsers, nil, err
	} else if hasDuplicateMessage {
		blog, err = b.BlogCore.GetDuplicateBlog(ctx, fromUserId, randomId)
		if err != nil {
			log.Errorf("checkDuplicateMessage error - %v", err)
			return pushUsers, nil, err
		} else if blog != nil {
			return pushUsers, blog, nil
		}
	}

	pushUsers, blog, err = b.BlogCore.SendBlog(ctx, fromUserId, visible, text, entities, mentionUserIds, content, geo, topicIds)
	if err != nil {
		return pushUsers, nil, err
	}
	/*
		var uf model.UserHelper = nil
		var pushUsers model.IDList
		if fnMutualUsers != nil {
			var mutualUsers []int32
			uf, _ = fnMutualUsers()
			if visible.VisibleType == model.VisibleType_Friend {
				pushUsers = mutualUsers
			} else if visible.VisibleType == model.VisibleType_Fans {
				pushUsers = b.BlogCore.GetAllFans(ctx, fromUserId)
			} else if visible.VisibleType == model.VisibleType_Allow || visible.VisibleType == model.VisibleType_NotAllow {
				var filteIds model.IDList
				if len(visible.GroupTags) > 0 {
					filteIds = b.GetGroupMembers(ctx, fromUserId, visible.GroupTags)
				}
				if len(visible.UserIds) > 0 {
					filteIds.AddIfNot(visible.UserIds...)
				}
				if visible.VisibleType == model.VisibleType_Allow {
					pushUsers = util.Int32Intersect(mutualUsers, filteIds.ToIDList())
				} else if visible.VisibleType == model.VisibleType_NotAllow {
					pushUsers = util.Int32Subtract(mutualUsers, filteIds.ToIDList())
				}
			}
		}
	*/
	b.BlogCore.PutDuplicateBlog(ctx, fromUserId, randomId, blog)
	return pushUsers, blog, nil
}

func (b *localBlogFacade) DeleteBlog(ctx context.Context, fromUserId int32, ids []int32) ([]int32, []int32, error) {
	return b.BlogCore.DeleteBlog(ctx, fromUserId, ids)
}

func (b *localBlogFacade) SendComment(ctx context.Context, fromUserId int32, fromAuthKeyId int64, id *mtproto.Blogs_IdType, text string) (int32, *mtproto.TLBlogsComment, error) {
	var (
		reply  *mtproto.TLBlogsComment
		err    error
		userId int32
	)
	if id.PredicateName == mtproto.Predicate_blogs_idTypeBlog {
		userId, reply, err = b.BlogCore.SendBlogComment(ctx, fromUserId, id.GetBlogId(), text)
	} else if id.PredicateName == mtproto.Predicate_blogs_idTypeComment {
		userId, reply, err = b.BlogCore.SendCommentReply(ctx, fromUserId, id.GetCommentId(), text)
	} else {
		err = mtproto.ErrButtonTypeInvalid
		log.Errorf("blogs.sendBlog: %v", err)
	}
	if err != nil {
		return 0, reply, err
	}

	return userId, reply, err
}

func (b *localBlogFacade) Follow(ctx context.Context, fromUserId int32, userId int32, followed bool) (bool, *mtproto.Blogs_UserDate, error) {
	return b.BlogCore.Follow(ctx, fromUserId, userId, followed)
}

func (b *localBlogFacade) Like(ctx context.Context, fromUserId int32, id *mtproto.Blogs_IdType, liked bool) (int32, *mtproto.Blogs_UserDate, error) {
	var (
		userId int32
		like   *mtproto.Blogs_UserDate
		err    error
	)
	if id.PredicateName == mtproto.Predicate_blogs_idTypeBlog {
		userId, like, err = b.BlogCore.LikeBlog(ctx, fromUserId, id.GetBlogId(), liked)
	} else if id.PredicateName == mtproto.Predicate_blogs_idTypeComment {
		userId, like, err = b.BlogCore.LikeReply(ctx, fromUserId, id.GetCommentId(), liked)
	} else {
		err = mtproto.ErrButtonTypeInvalid
		log.Errorf("blogs.sendBlog: %v", err)
	}
	return userId, like, err
}

func (b *localBlogFacade) GetUserHistory(ctx context.Context, fromUserId int32, user_id int32, isHisFriend, mutualFriend bool, min_id int32, offset int32, limit int32) (*mtproto.MicroBlogs, error) {
	var (
		blogs *mtproto.MicroBlogs
		err   error
	)
	if fromUserId == user_id {
		blogs, err = b.BlogCore.GetSelfHistory(ctx, fromUserId, min_id, offset, limit)
	} else {
		blogs, err = b.BlogCore.GetUserHistory(ctx, fromUserId, user_id, isHisFriend, mutualFriend, min_id, offset, limit)
	}
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (b *localBlogFacade) GetCircleHistory(ctx context.Context, fromUserId int32, visible *model.VisibleTypeUtil, min_id int32, offset int32, limit int32, fnMutualUsers func() []int32, fnIsHisFriend func(int32) bool) (*mtproto.MicroBlogs, error) {
	var (
		blogs *mtproto.MicroBlogs
		err   error
	)
	switch visible.VisibleType {
	case model.VisibleType_Public:
		blogs, err = b.BlogCore.GetPublicHistory(ctx, fromUserId, min_id, offset, limit, fnIsHisFriend)
	case model.VisibleType_Friend, model.VisibleType_Follow:
		users := []int32{}
		if visible.VisibleType == model.VisibleType_Follow {
			users = b.BlogCore.GetAllFollows(ctx, fromUserId)
		} else if fnMutualUsers != nil {
			users = fnMutualUsers()
			users = append(users, fromUserId)
		}
		blogs, err = b.BlogCore.GetUsersHistory(ctx, fromUserId, users, min_id, offset, limit, fnIsHisFriend)
	case model.VisibleType_Topic:
		blogs, err = b.BlogCore.GetTopicHistory(ctx, fromUserId, min_id, offset, limit, visible.TopicId, visible.Tab, fnIsHisFriend)
	}
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (b *localBlogFacade) GetBlogs(ctx context.Context, fromUserId int32, ids []int32) (*mtproto.MicroBlogs, error) {
	blogs, err := b.BlogCore.GetBlogs(ctx, fromUserId, ids)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (b *localBlogFacade) GmGetBlogs(ctx context.Context, ids []int32) (*mtproto.MicroBlogs, error) {
	blogs, err := b.BlogCore.GmGetBlogs(ctx, ids)
	if err != nil {
		return nil, err
	}
	return blogs, nil
}

func (b *localBlogFacade) GetUser(ctx context.Context, fromUserId int32, targerId int32) (*mtproto.Blogs_User, error) {
	return b.BlogCore.GetBlogUser(ctx, fromUserId, targerId)
}

func (b *localBlogFacade) GetComments(ctx context.Context, fromUserId int32, Id *mtproto.Blogs_IdType, offsetId int32, limit int32) (*mtproto.Blogs_Comments, error) {
	var (
		reply *mtproto.Blogs_Comments
		err   error
	)
	if Id.PredicateName == mtproto.Predicate_blogs_idTypeBlog {
		reply, err = b.BlogCore.GetBlogComments(ctx, fromUserId, Id.GetBlogId(), offsetId, limit)
	} else if Id.PredicateName == mtproto.Predicate_blogs_idTypeComment {
		reply, err = b.BlogCore.GetCommentAndReply(ctx, fromUserId, Id.GetCommentId(), offsetId, limit)
	} else {
		err = mtproto.ErrButtonTypeInvalid
		log.Errorf("blogs.sendBlog: %v", err)
	}
	return reply, err
}

func (b *localBlogFacade) GetCommentList(ctx context.Context, fromUserId int32, ids []int32) (*mtproto.Blogs_Comments, error) {
	return b.BlogCore.GetCommentList(ctx, fromUserId, ids)
}

func (b *localBlogFacade) GetLikes(ctx context.Context, fromUserId int32, Id *mtproto.Blogs_IdType, offsetId int32, limit int32) (*mtproto.Blogs_UserDates, error) {
	var (
		reply *mtproto.Blogs_UserDates
		err   error
	)
	if Id.PredicateName == mtproto.Predicate_blogs_idTypeBlog {
		reply, err = b.BlogCore.GetBlogLikes(ctx, fromUserId, Id.GetBlogId(), offsetId, limit)
	} else if Id.PredicateName == mtproto.Predicate_blogs_idTypeComment {
		reply, err = b.BlogCore.GetCommentLikes(ctx, fromUserId, Id.GetCommentId(), offsetId, limit)
	} else {
		err = mtproto.ErrButtonTypeInvalid
		log.Errorf("blogs.sendBlog: %v", err)
	}
	return reply, err
}

func (b *localBlogFacade) GetGroupTags(ctx context.Context, fromUserId int32) (*mtproto.Blogs_GroupTags, error) {
	var (
		reply *mtproto.Blogs_GroupTags
		err   error
	)
	reply, err = b.BlogCore.GetGroupTags(ctx, fromUserId)
	return reply, err
}

func (b *localBlogFacade) GetFollows(ctx context.Context, fromUserId int32, userId int32, offset int32, limit int32) (*mtproto.Blogs_UserDates, error) {
	var (
		reply *mtproto.Blogs_UserDates
		err   error
	)
	reply, err = b.BlogCore.GetFollows(ctx, fromUserId, userId, offset, limit)
	return reply, err
}

func (b *localBlogFacade) GetFans(ctx context.Context, fromUserId int32, userId int32, offset int32, limit int32) (*mtproto.Blogs_UserDates, error) {
	var (
		reply *mtproto.Blogs_UserDates
		err   error
	)
	reply, err = b.BlogCore.GetFans(ctx, fromUserId, userId, offset, limit)
	return reply, err
}

func (b *localBlogFacade) CreateGroupTag(ctx context.Context, fromUserId int32, title string, userIds []int32) (*mtproto.Blogs_GroupTag, error) {
	return b.BlogCore.CreateGroupTag(ctx, fromUserId, title, userIds)
}

func (b *localBlogFacade) AddGroupTagMember(ctx context.Context, fromUserId int32, tagId int32, userIds []int32) ([]int32, error) {
	return b.BlogCore.AddGroupTagMember(ctx, fromUserId, tagId, userIds)
}

func (b *localBlogFacade) DeleteGroupTagMember(ctx context.Context, fromUserId int32, tagId int32, userIds []int32) ([]int32, error) {
	return b.BlogCore.DeleteGroupTagMember(ctx, fromUserId, tagId, userIds)
}

func (b *localBlogFacade) DeleteGroupTag(ctx context.Context, fromUserId int32, tagIds []int32) ([]int32, error) {
	return b.BlogCore.DeleteGroupTag(ctx, fromUserId, tagIds)
}

func (b *localBlogFacade) EditGroupTag(ctx context.Context, fromUserId int32, tagId int32, title string) (bool, error) {
	return b.BlogCore.EditGroupTag(ctx, fromUserId, tagId, title)
}

func (b *localBlogFacade) ReadHistory(ctx context.Context, fromUserId int32, pts int32, maxId *mtproto.Blogs_IdType) (*mtproto.Blogs_IdType, error) {
	var (
		newMax *mtproto.Blogs_IdType
		err    error
	)
	if pts > 0 {
		_, err = b.BlogCore.ReadMaxUpdateId(ctx, fromUserId, pts)
	}
	if err != nil {
		return nil, err
	}
	if maxId != nil {
		if maxId.PredicateName == mtproto.Predicate_blogs_idTypeBlog {
			newMax, err = b.BlogCore.ReadMaxBlogId(ctx, fromUserId, maxId.GetBlogId())
		} else if maxId.PredicateName == mtproto.Predicate_blogs_idTypeComment {
			newMax, err = b.BlogCore.ReadMaxCommentId(ctx, fromUserId, maxId.GetCommentId())
		} else {
			err = mtproto.ErrButtonTypeInvalid
		}
	}

	return newMax, err
}

func (b *localBlogFacade) GetUnreads(ctx context.Context, fromUserId int32, pts, limit int32) (*mtproto.Blogs_Unreads, error) {
	return b.BlogCore.GetUnreads(ctx, fromUserId, pts, limit)
}

func (b *localBlogFacade) Reward(ctx context.Context, fromUserId int32, userId int32, blogId int32) (bool, error) {
	return true, nil
}

func (b *localBlogFacade) SetPrivacy(ctx context.Context, userId int32, key int8, rules string) error {
	return b.BlogCore.SetPrivacy(ctx, userId, key, rules)
}

func (b *localBlogFacade) GetUserPrivacy(ctx context.Context, userId int32) (map[int8]string, error) {
	return b.BlogCore.GetUserPrivacy(ctx, userId)
}

func (b *localBlogFacade) ModifyPrivacyUsers(ctx context.Context, userId int32, key int8, uidList []int32, isAdding bool) error {
	return b.BlogCore.ModifyPrivacyUsers(ctx, userId, key, uidList, isAdding)
}

func (b *localBlogFacade) TouchTopics(ctx context.Context, nameList []string) ([]*mtproto.Blogs_Topic, error) {
	return b.BlogCore.TouchTopics(ctx, nameList)
}

func (b *localBlogFacade) GetTopics(ctx context.Context, fromTopicId, limit int32) (*mtproto.Blogs_Topics, error) {
	return b.BlogCore.GetTopics(ctx, fromTopicId, limit)
}

func (b *localBlogFacade) GetHotTopics(ctx context.Context, limit int32) (*mtproto.Blogs_Topics, error) {
	return b.BlogCore.GetHotTopics(ctx, limit)
}

func (b *localBlogFacade) GetTopicIdByName(ctx context.Context, name string) (int32, error) {
	return b.BlogCore.GetTopicIdByName(ctx, name)
}

func (b *localBlogFacade) AddBannedUser(ctx context.Context, userId int32, banFrom, banTo int32) error {
	return b.BlogCore.AddBannedUser(ctx, userId, banFrom, banTo)
}

func (b *localBlogFacade) GetBannedUsers(ctx context.Context, offset, limit int32) (map[int32][]int32, error) {
	return b.BlogCore.GetBannedUsers(ctx, offset, limit)
}

func (b *localBlogFacade) IsBannedUser(ctx context.Context, userId int32) (bool, error) {
	return b.BlogCore.IsBannedUser(ctx, userId)
}

func (b *localBlogFacade) GetBannedByUsers(ctx context.Context, uidList []int32) (map[int32][]int32, error) {
	return b.BlogCore.GetBannedByUserIds(ctx, uidList)
}

func init() {
	Register("local", New)
}
