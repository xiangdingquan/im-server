package blog

import (
	"context"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
	"open.chat/app/json/helper"
	svc "open.chat/app/json/service"
	"open.chat/app/json/service/http"
	"open.chat/app/messenger/biz_server/upload"
	blog_client "open.chat/app/service/biz_service/blog/facade"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type cls struct {
	blog_client.BlogFacade
	mtproto.RPCUploadServer
}

func New(s *svc.Service, rg *bm.RouterGroup) {
	service := &cls{
		RPCUploadServer: upload.New(),
	}

	var err error
	service.BlogFacade, err = blog_client.NewBlogFacade("local")
	checkErr(err)

	http.RegisterBlog(service, rg)
}

func (s *cls) GetContent(ctx context.Context, r *http.TBlogGetContent) *helper.ResultJSON {
	blogs, err := s.BlogFacade.GmGetBlogs(ctx, []int32{int32(r.BlogId)})
	if err != nil {
		log.Errorf("/blog/getContent failed, error:%v", err)
		return nil
	}

	if blogs.PredicateName != mtproto.Predicate_microBlogs {
		log.Errorf("/blog/getContent, blogs not found")
		return nil
	}

	var blog *mtproto.MicroBlog
	for _, v := range blogs.GetBlogs() {
		if v.PredicateName != mtproto.Predicate_microBlog {
			continue
		}
		if v.GetId() == int32(r.BlogId) {
			blog = v
			break
		}
	}

	if blog == nil {
		log.Errorf("/blog/getContent, blogs not found")
		return nil
	}

	blogContent := blog.GetContent()
	if blogContent == nil {
		log.Debugf("/blog/getContent, blog content is nil")
		return nil
	}

	var out []http.TBlogGetContentFile

	switch blogContent.PredicateName {
	case mtproto.Predicate_blogContentPhotos:
		log.Debugf("/blog/getContent, photos content, size:%d", len(blogContent.GetPhotos()))
		out = make([]http.TBlogGetContentFile, len(blogContent.GetPhotos()))
		for i, photo := range blogContent.GetPhotos() {
			if photo.PredicateName != mtproto.Predicate_photo {
				log.Errorf("/blog/getContent, invalid photo, %s", photo.PredicateName)
				return nil
			}

			out[i] = http.TBlogGetContentFile{
				Type:       1,
				Id:         photo.GetId(),
				AccessHash: photo.GetAccessHash(),
			}
		}
	case mtproto.Predicate_blogContentDocument:
		document := blogContent.GetDocument()
		if document.PredicateName == mtproto.Predicate_document {
			out = []http.TBlogGetContentFile{
				{
					Type:       2,
					Id:         document.GetId(),
					AccessHash: document.GetAccessHash(),
				},
			}
		}
	}

	return &helper.ResultJSON{
		Code: 200,
		Msg:  "success",
		Data: out,
	}
}

func (s *cls) GetContentFile(ctx context.Context, r *http.TBlogGetContentFile) *[]byte {
	log.Debugf("/blog/getContentFile, %v", r)

	var fl *mtproto.InputFileLocation
	switch r.Type {
	case 1:
		fl = mtproto.MakeTLInputPhotoFileLocation(&mtproto.InputFileLocation{
			Id:         r.Id,
			AccessHash: r.AccessHash,
		}).To_InputFileLocation()
	case 2:
		fl = mtproto.MakeTLInputDocumentFileLocation(&mtproto.InputFileLocation{
			Id:         r.Id,
			AccessHash: r.AccessHash,
		}).To_InputFileLocation()
	default:
		log.Errorf("/blog/getContentFile, unknown type: %d", r.Type)
		return nil
	}

	limit := 1024 * 100
	gf := &mtproto.TLUploadGetFile{
		Location: fl,
		Offset:   0,
		Limit:    int32(limit),
	}
	data := make([]byte, 0)
	uf, err := s.RPCUploadServer.UploadGetFile(ctx, gf)
	if err != nil {
		return nil
	}
	data = append(data, uf.Bytes...)
	for len(uf.Bytes) >= limit {
		gf.Offset += int32(limit)
		uf, err = s.RPCUploadServer.UploadGetFile(ctx, gf)
		if err != nil {
			return nil
		}
		data = append(data, uf.Bytes...)
	}

	return &data
}

func (s *cls) GetBannedUsers(ctx context.Context, r *http.TBlogGetBannedUsers) *helper.ResultJSON {
	m, err := s.BlogFacade.GetBannedUsers(ctx, r.Offset, r.Limit)
	if err != nil {
		log.Errorf("/blog/getBannedUsers, get failed, %v", err)
		return &helper.ResultJSON{Code: -1, Msg: "get failed"}
	}

	l := make([]*http.TBlogBannedUser, 0, len(m))
	for k, v := range m {
		l = append(l, &http.TBlogBannedUser{
			UserId:  k,
			BanFrom: v[0],
			BanTo:   v[1],
		})
	}

	return &helper.ResultJSON{
		Code: 200,
		Msg:  "success",
		Data: l,
	}
}

func (s *cls) BanUser(ctx context.Context, r *http.TBlogBannedUser) *helper.ResultJSON {
	err := s.BlogFacade.AddBannedUser(ctx, r.UserId, r.BanFrom, r.BanTo)
	if err != nil {
		log.Errorf("/blog/banUser, add failed, %v", err)
		return &helper.ResultJSON{Code: -1, Msg: "add failed"}
	}
	return &helper.ResultJSON{Code: 200, Msg: "success"}
}

func (s *cls) GetBannedByUsers(ctx context.Context, r *http.TBlogGetBannedByUsers) *helper.ResultJSON {
	m, err := s.BlogFacade.GetBannedByUsers(ctx, r.UidList)
	if err != nil {
		log.Errorf("/blog/getBannedByUsers, get failed, %v", err)
		return &helper.ResultJSON{Code: -1, Msg: "get failed"}
	}

	l := make([]*http.TBlogBannedUser, 0, len(m))
	for k, v := range m {
		l = append(l, &http.TBlogBannedUser{
			UserId:  k,
			BanFrom: v[0],
			BanTo:   v[1],
		})
	}

	return &helper.ResultJSON{
		Code: 200,
		Msg:  "success",
		Data: l,
	}
}
