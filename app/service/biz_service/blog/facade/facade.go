package blog_facade

import (
	"context"
	"fmt"
	"open.chat/model"
	"open.chat/mtproto"
)

type BlogFacade interface {
	SendBlog(ctx context.Context, fromUserId int32, randomId int64, visible *model.VisibleTypeUtil, text string, entities []*mtproto.MessageEntity, mentionUserIds []int32, content *mtproto.BlogContent, geo *mtproto.BlogGeoPoint, topicIds []int32) (model.IDList, *mtproto.MicroBlog, error)
	DeleteBlog(ctx context.Context, fromUserId int32, ids []int32) ([]int32, []int32, error)
	SendComment(ctx context.Context, fromUserId int32, fromAuthKeyId int64, id *mtproto.Blogs_IdType, text string) (int32, *mtproto.TLBlogsComment, error)
	Like(ctx context.Context, fromUserId int32, id *mtproto.Blogs_IdType, liked bool) (int32, *mtproto.Blogs_UserDate, error)
	Follow(ctx context.Context, fromUserId int32, userId int32, followed bool) (bool, *mtproto.Blogs_UserDate, error)
	//CanReward(ctx context.Context, fromUserId int32, userId int32, blogId int32) (bool, error)
	Reward(ctx context.Context, fromUserId int32, userId int32, blogId int32) (bool, error)

	GetCircleHistory(ctx context.Context, fromUserId int32, visible *model.VisibleTypeUtil, min_id int32, offset int32, limit int32, fnMutualUsers func() []int32, fnIsHisFriend func(int32) bool) (*mtproto.MicroBlogs, error)
	GetUserHistory(ctx context.Context, fromUserId int32, user_id int32, isHisFriend, isMutualFriend bool, min_id int32, offset int32, limit int32) (*mtproto.MicroBlogs, error)
	GetUser(ctx context.Context, fromUserId int32, targerId int32) (*mtproto.Blogs_User, error)
	GetComments(ctx context.Context, fromUserId int32, id *mtproto.Blogs_IdType, offsetId, limit int32) (*mtproto.Blogs_Comments, error)
	GetLikes(ctx context.Context, fromUserId int32, id *mtproto.Blogs_IdType, offsetId, limit int32) (*mtproto.Blogs_UserDates, error)

	GetGroupTags(ctx context.Context, fromUserId int32) (*mtproto.Blogs_GroupTags, error)
	GetFollows(ctx context.Context, fromUserId int32, userId int32, offset, limit int32) (*mtproto.Blogs_UserDates, error)
	GetFans(ctx context.Context, fromUserId int32, userId int32, offset, limit int32) (*mtproto.Blogs_UserDates, error)
	GetBlogs(ctx context.Context, fromUserId int32, ids []int32) (*mtproto.MicroBlogs, error)
	GmGetBlogs(ctx context.Context, ids []int32) (*mtproto.MicroBlogs, error)
	GetCommentList(ctx context.Context, fromUserId int32, ids []int32) (*mtproto.Blogs_Comments, error)

	CreateGroupTag(ctx context.Context, fromUserId int32, title string, userIds []int32) (*mtproto.Blogs_GroupTag, error)
	AddGroupTagMember(ctx context.Context, fromUserId int32, tagId int32, userIds []int32) ([]int32, error)
	DeleteGroupTagMember(ctx context.Context, fromUserId int32, tagId int32, userIds []int32) ([]int32, error)
	DeleteGroupTag(ctx context.Context, fromUserId int32, tagIds []int32) ([]int32, error)
	EditGroupTag(ctx context.Context, fromUserId int32, tagId int32, title string) (bool, error)

	GetGroupMembers(ctx context.Context, fromUserId int32, ids []int32) []int32
	ReadHistory(ctx context.Context, fromUserId int32, pts int32, maxId *mtproto.Blogs_IdType) (*mtproto.Blogs_IdType, error)

	GetUnreads(ctx context.Context, fromUserId int32, pts, limit int32) (*mtproto.Blogs_Unreads, error)

	SetPrivacy(ctx context.Context, userId int32, key int8, rules string) error
	ModifyPrivacyUsers(ctx context.Context, userId int32, key int8, uidList []int32, isAdding bool) error
	GetUserPrivacy(ctx context.Context, userId int32) (map[int8]string, error)

	TouchTopics(ctx context.Context, nameList []string) ([]*mtproto.Blogs_Topic, error)
	GetTopics(ctx context.Context, fromTopicId, limit int32) (*mtproto.Blogs_Topics, error)
	GetHotTopics(ctx context.Context, limit int32) (*mtproto.Blogs_Topics, error)
	GetTopicIdByName(ctx context.Context, name string) (int32, error)

	AddBannedUser(ctx context.Context, userId int32, banFrom, banTo int32) error
	GetBannedUsers(ctx context.Context, offset, limit int32) (map[int32][]int32, error)
	IsBannedUser(ctx context.Context, userId int32) (bool, error)
	GetBannedByUsers(ctx context.Context, uidList []int32) (map[int32][]int32, error)
}

type Instance func() BlogFacade

var instances = make(map[string]Instance)

func Register(name string, inst Instance) {
	if inst == nil {
		panic("register instance is nil")
	}
	if _, ok := instances[name]; ok {
		panic("register called twice for instance " + name)
	}
	instances[name] = inst
}

func NewBlogFacade(name string) (inst BlogFacade, err error) {
	instanceFunc, ok := instances[name]
	if !ok {
		err = fmt.Errorf("unknown adapter name %q (forgot to import?)", name)
		return
	}
	inst = instanceFunc()
	return
}
