package server

import (
	"context"
	"fmt"

	"open.chat/app/messenger/biz_server/blogs/internal/service"
	user_client "open.chat/app/service/biz_service/user/client"
	"open.chat/app/sysconfig"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

// ///////////////////////////////////////////////////////////////////////////
type Server struct {
	svc *service.Service
}

func New() *Server {
	return new(Server)
}

func (s *Server) Initialize() error {
	s.svc = service.New()
	return nil
}

func (s *Server) RunLoop() {
	var (
		fromUserId int32 = 136817712
		ctx              = context.TODO()
		blogFacade       = s.svc.BlogFacade
		userFacade user_client.UserFacade
	)
	userFacade, _ = user_client.NewUserFacade("local")
	code := "21292"
	verifyCode := sysconfig.GetConfig2String(ctx, sysconfig.ConfigKeysCommonSmsCode, "", 0)
	if len(verifyCode) == 0 || code != verifyCode {
		log.Errorf("verifySmsCode - code invalid")
	}
	return
	func() {
		return
		var (
			targerId int32 = fromUserId
		)

		if targerId != fromUserId {
			user, err := userFacade.GetUserById(ctx, fromUserId, targerId)
			if err != nil || user == nil || user.Deleted {
				if err == nil {
					err = mtproto.ErrUserIdInvalid
				}
				fmt.Printf("blogFacade.Follow - error: %s\n", err.Error())
				return
			}
		}

		reply, err := blogFacade.GetUser(ctx, fromUserId, targerId)
		if err != nil {
			fmt.Printf("blogFacade.GetUser - error: %s\n", err.Error())
		}
		if reply != nil {
			fmt.Printf("blogFacade.GetUser - reply: %s\n", reply.DebugString())
		}
	}()
	func() {
		var (
			visible                = model.MakeVisibleTypeAllow([]int32{2}, []int32{136817715})
			randomId       int64   = 0
			text           string  = "test blog"
			mentionUserIds []int32 // = []int32{136817715}
		)
		/*
			func() (model.UserHelper, []int32) {
				mutualUsers := []int32{}
				if visible.VisibleType == model.VisibleType_Friend || visible.VisibleType == model.VisibleType_Allow || visible.VisibleType == model.VisibleType_NotAllow {
					mutualUsers = userFacade.GetMutualContactUserIdList(ctx, true, fromUserId)
				}
				if visible.VisibleType != model.VisibleType_Public && visible.VisibleType != model.VisibleType_Private {
					//获取自己双向用户
					mutualUsers = userFacade.GetMutualContactUserIdList(ctx, true, fromUserId)
				}
				return userFacade, mutualUsers
			}
		*/
		_, reply, err := blogFacade.SendBlog(ctx, fromUserId, randomId, visible, text, mentionUserIds, nil, nil)
		if err != nil {
			fmt.Printf("blogFacade.SendBlog - error: %s\n", err.Error())
		}
		if reply != nil {
			fmt.Printf("blogFacade.SendBlog - reply: %s\n", reply.DebugString())
		}
	}()
	func() {
		var (
			ids []int32 = []int32{1}
		)
		deleteIds, shieldIds, err := blogFacade.DeleteBlog(ctx, fromUserId, ids)
		if err != nil {
			fmt.Printf("blogFacade.DeleteBlog - error: %s\n", err.Error())
		}
		fmt.Printf("blogFacade.DeleteBlog - reply: %v,%v\n", deleteIds, shieldIds)
	}()
	func() {
		return
		var (
			userId   int32 = 136817712
			followed bool  = true
		)
		if fromUserId == userId {
			err := mtproto.ErrPeerIdInvalid
			fmt.Printf("blogFacade.Follow - error: %s\n", err.Error())
			return
		}

		if followed {
			user, err := userFacade.GetUserById(ctx, fromUserId, userId)
			if err != nil || user == nil || user.Deleted {
				if err == nil {
					err = mtproto.ErrUserIdInvalid
				}
				fmt.Printf("blogFacade.Follow - error: %s\n", err.Error())
				return
			}
		}

		_, reply, err := blogFacade.Follow(ctx, fromUserId, userId, followed)
		if err != nil {
			fmt.Printf("blogFacade.Follow - error: %s\n", err.Error())
		}
		fmt.Printf("blogFacade.Follow - reply: %v\n", reply)
	}()
	func() {
		return
		var (
			userId int32 = fromUserId
			offset int32 = 0
			limit  int32 = 100
		)

		if userId != fromUserId {
			user, err := userFacade.GetUserById(ctx, fromUserId, userId)
			if err != nil || user == nil || user.Deleted {
				if err == nil {
					err = mtproto.ErrUserIdInvalid
				}
				fmt.Printf("blogFacade.Follow - error: %s\n", err.Error())
				return
			}
		}

		reply, err := blogFacade.GetFollows(ctx, fromUserId, userId, offset, limit)
		if err != nil {
			fmt.Printf("blogFacade.GetFollows - error: %s\n", err.Error())
		}
		if reply != nil {
			fmt.Printf("blogFacade.GetFollows - reply: %s\n", reply.DebugString())
		}
	}()
	func() {
		return
		var (
			userId int32 = fromUserId
			offset int32 = 0
			limit  int32 = 100
		)

		if userId != fromUserId {
			user, err := userFacade.GetUserById(ctx, fromUserId, userId)
			if err != nil || user == nil || user.Deleted {
				if err == nil {
					err = mtproto.ErrUserIdInvalid
				}
				fmt.Printf("blogFacade.Follow - error: %s\n", err.Error())
				return
			}
		}

		reply, err := blogFacade.GetFans(ctx, fromUserId, userId, offset, limit)
		if err != nil {
			fmt.Printf("blogFacade.GetFans - error: %s\n", err.Error())
		}
		if reply != nil {
			fmt.Printf("blogFacade.GetFans - reply: %s\n", reply.DebugString())
		}
	}()
	func() {
		return
		var (
			visibleType       = model.MakeVisibleTypePublic()
			min_id      int32 = 0
			offset      int32 = 0
			limit       int32 = 100
		)
		//blogFacade.GetUserHistory
		reply, err := blogFacade.GetCircleHistory(ctx, fromUserId, visibleType, min_id, offset, limit, func() []int32 {
			return userFacade.GetMutualContactUserIdList(ctx, true, fromUserId)
		})
		if err != nil {
			fmt.Printf("blogFacade.GetHistory - error: %s\n", err.Error())
		}
		if reply != nil {
			fmt.Printf("blogFacade.GetHistory - reply: %s\n", reply.DebugString())
		}
	}()
	func() {
		return
		var (
			pts    int32                 = 0
			max_id *mtproto.Blogs_IdType = mtproto.MakeTLBlogsIdTypeBlog(&mtproto.Blogs_IdType{
				BlogId: 1,
			}).To_Blogs_IdType()
		)
		reply, err := blogFacade.ReadHistory(ctx, fromUserId, pts, max_id)
		if err != nil {
			fmt.Printf("blogFacade.ReadHistory - error: %s\n", err.Error())
		}
		fmt.Printf("blogFacade.ReadHistory - reply: %v\n", reply)
	}()
	func() {
		return
		var (
			Id = mtproto.MakeTLBlogsIdTypeBlog(&mtproto.Blogs_IdType{
				BlogId: 1,
			}).To_Blogs_IdType()
			text string = "test comment"
		)
		_, replyId, err := blogFacade.SendComment(ctx, fromUserId, 0, Id, text)
		if err != nil {
			fmt.Printf("blogFacade.Comment - error: %s\n", err.Error())
		}
		fmt.Printf("blogFacade.Comment - reply: %v\n", replyId)
	}()
	func() {
		return
		var (
			Id = mtproto.MakeTLBlogsIdTypeBlog(&mtproto.Blogs_IdType{
				BlogId: 1,
			}).To_Blogs_IdType()
			liked bool = false
		)
		_, reply, err := blogFacade.Like(ctx, fromUserId, Id, liked)
		if err != nil {
			fmt.Printf("blogFacade.Like - error: %s\n", err.Error())
		}
		fmt.Printf("blogFacade.Like - reply: %v\n", reply)
	}()
	func() {
		return
		var (
			Id = mtproto.MakeTLBlogsIdTypeBlog(&mtproto.Blogs_IdType{
				BlogId: 2,
			}).To_Blogs_IdType()
			offset int32 = 0
			limit  int32 = 100
		)
		reply, err := blogFacade.GetComments(ctx, fromUserId, Id, offset, limit)
		if err != nil {
			fmt.Printf("blogFacade.GetComments - error: %s\n", err.Error())
		}
		if reply != nil {
			fmt.Printf("blogFacade.GetComments - reply: %s\n", reply.DebugString())
		}
	}()
	func() {
		return
		var (
			Id = mtproto.MakeTLBlogsIdTypeBlog(&mtproto.Blogs_IdType{
				BlogId: 1,
			}).To_Blogs_IdType()
			offset int32 = 0
			limit  int32 = 100
		)
		reply, err := blogFacade.GetLikes(ctx, fromUserId, Id, offset, limit)
		if err != nil {
			fmt.Printf("blogFacade.GetLikes - error: %s\n", err.Error())
		}
		if reply != nil {
			fmt.Printf("blogFacade.GetLikes - reply: %s\n", reply.DebugString())
		}
	}()
	func() {
		return
		var (
			ids []int32 = []int32{1, 2, 3, 4}
		)
		reply, err := blogFacade.GetBlogs(ctx, fromUserId, ids)
		if err != nil {
			fmt.Printf("blogFacade.GetBlogs - error: %s\n", err.Error())
		}
		if reply != nil {
			fmt.Printf("blogFacade.GetBlogs - reply: %s\n", reply.DebugString())
		}
	}()
	func() {
		return
		var (
			title   string  = "test group tag"
			userIds []int32 = []int32{1, 2, 136817712, 136817715}
		)

		users := userFacade.GetUserListByIdList(ctx, fromUserId, userIds)
		userIds = make([]int32, 0, len(users))
		for _, user := range users {
			if !user.Deleted && user.GetId() != fromUserId {
				userIds = append(userIds, user.GetId())
			}
		}

		if len(userIds) == 0 {
			err := mtproto.ErrUserIdInvalid
			fmt.Printf("blogFacade.Follow - error: %s\n", err.Error())
			return
		}

		reply, err := blogFacade.CreateGroupTag(ctx, fromUserId, title, userIds)
		if err != nil {
			fmt.Printf("blogFacade.CreateGroupTag - error: %s\n", err.Error())
		}
		if reply != nil {
			fmt.Printf("blogFacade.CreateGroupTag - reply: %s\n", reply.DebugString())
		}
	}()
	func() {
		return
		reply, err := blogFacade.GetGroupTags(ctx, fromUserId)
		if err != nil {
			fmt.Printf("blogFacade.GetGroupTags - error: %s\n", err.Error())
		}
		if reply != nil {
			fmt.Printf("blogFacade.GetGroupTags - reply: %s\n", reply.DebugString())
		}
	}()
	func() {
		return
		var (
			tagId   int32   = 1
			userIds []int32 = []int32{1, 2, 136817712, 136817715}
		)

		users := userFacade.GetUserListByIdList(ctx, fromUserId, userIds)
		userIds = make([]int32, 0, len(users))
		for _, user := range users {
			if !user.Deleted && user.GetId() != fromUserId {
				userIds = append(userIds, user.GetId())
			}
		}

		if len(userIds) == 0 {
			err := mtproto.ErrUserIdInvalid
			fmt.Printf("blogFacade.Follow - error: %s\n", err.Error())
			return
		}

		reply, err := blogFacade.AddGroupTagMember(ctx, fromUserId, tagId, userIds)
		if err != nil {
			fmt.Printf("blogFacade.AddGroupTagMember - error: %s\n", err.Error())
		}
		fmt.Printf("blogFacade.AddGroupTagMember - reply: %v\n", reply)
	}()
	func() {
		return
		var (
			tagId   int32   = 1
			userIds []int32 = []int32{1, 2, 136817712, 136817715}
		)

		users := userFacade.GetUserListByIdList(ctx, fromUserId, userIds)
		userIds = make([]int32, 0, len(users))
		for _, user := range users {
			if user.GetId() != fromUserId {
				userIds = append(userIds, user.GetId())
			}
		}

		if len(userIds) == 0 {
			err := mtproto.ErrUserIdInvalid
			fmt.Printf("blogFacade.Follow - error: %s\n", err.Error())
			return
		}

		reply, err := blogFacade.DeleteGroupTagMember(ctx, fromUserId, tagId, userIds)
		if err != nil {
			fmt.Printf("blogFacade.DeleteGroupTagMember - error: %s\n", err.Error())
		}
		fmt.Printf("blogFacade.DeleteGroupTagMember - reply: %v\n", reply)
	}()
	func() {
		return
		var (
			tagId int32  = 1
			title string = "test group tag 2"
		)
		reply, err := blogFacade.EditGroupTag(ctx, fromUserId, tagId, title)
		if err != nil {
			fmt.Printf("blogFacade.EditGroupTag - error: %s\n", err.Error())
		}
		fmt.Printf("blogFacade.EditGroupTag - reply: %v\n", reply)
	}()
	func() {
		return
		var (
			tagIds []int32 = []int32{1, 2, 3, 4}
		)
		reply, err := blogFacade.DeleteGroupTag(ctx, fromUserId, tagIds)
		if err != nil {
			fmt.Printf("blogFacade.DeleteGroupTag - error: %s\n", err.Error())
		}
		fmt.Printf("blogFacade.DeleteGroupTag - reply: %v\n", reply)
	}()
}

func (s *Server) Destroy() {
}
