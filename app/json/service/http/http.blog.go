package http

import (
	"context"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"open.chat/app/json/helper"
)

type (
	TBlogGetContent struct {
		BlogId int32 `json:"blogId"`
	}

	TBlogGetContentFile struct {
		Type       int32 `json:"type"`
		Id         int64 `json:"id"`
		AccessHash int64 `json:"accessHash"`
		Offset     int32 `json:"offset"`
	}

	TBlogGetBannedUsers struct {
		Offset int32 `json:"offset"`
		Limit  int32 `json:"limit"`
	}

	TBlogBannedUser struct {
		UserId  int32 `json:"userId"`
		BanFrom int32 `json:"banFrom"`
		BanTo   int32 `json:"banTo"`
	}

	TBlogGetBannedByUsers struct {
		UidList []int32 `json:"uidList"`
	}

	ServiceBlog interface {
		GetContent(context.Context, *TBlogGetContent) *helper.ResultJSON
		GetContentFile(ctx context.Context, file *TBlogGetContentFile) *[]byte
		GetBannedUsers(ctx context.Context, r *TBlogGetBannedUsers) *helper.ResultJSON
		BanUser(ctx context.Context, r *TBlogBannedUser) *helper.ResultJSON
		GetBannedByUsers(ctx context.Context, r *TBlogGetBannedByUsers) *helper.ResultJSON
	}
)

func RegisterBlog(s ServiceBlog, rg *bm.RouterGroup) {
	rg2 := rg.Group("/blog")

	rg2.POST("/getContent", func(c *bm.Context) {
		helper.DoHttpJson(c, &TBlogGetContent{}, func(ctx context.Context, data interface{}) *helper.ResultJSON {
			return s.GetContent(ctx, data.(*TBlogGetContent))
		})
	})
	rg2.POST("/getContentFile", func(c *bm.Context) {
		helper.DoHttpDownload(c, &TBlogGetContentFile{}, func(ctx context.Context, data interface{}) *[]byte {
			return s.GetContentFile(ctx, data.(*TBlogGetContentFile))
		})
	})

	rg2.POST("/getBannedUsers", func(c *bm.Context) {
		helper.DoHttpJson(c, &TBlogGetBannedUsers{}, func(ctx context.Context, data interface{}) *helper.ResultJSON {
			return s.GetBannedUsers(ctx, data.(*TBlogGetBannedUsers))
		})
	})
	rg2.POST("/banUser", func(c *bm.Context) {
		helper.DoHttpJson(c, &TBlogBannedUser{}, func(ctx context.Context, data interface{}) *helper.ResultJSON {
			return s.BanUser(ctx, data.(*TBlogBannedUser))
		})
	})
	rg2.POST("/getBannedByUsers", func(c *bm.Context) {
		helper.DoHttpJson(c, &TBlogGetBannedByUsers{}, func(ctx context.Context, data interface{}) *helper.ResultJSON {
			return s.GetBannedByUsers(ctx, data.(*TBlogGetBannedByUsers))
		})
	})
}
