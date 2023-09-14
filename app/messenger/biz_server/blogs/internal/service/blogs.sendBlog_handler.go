package service

import (
	"context"
	"github.com/pkg/errors"
	"time"

	"github.com/gogo/protobuf/proto"
	idgen "open.chat/app/service/idgen/client"
	media_client "open.chat/app/service/media/client"
	"open.chat/model"
	"open.chat/pkg/log"

	sync_client "open.chat/app/messenger/sync/client"
	"open.chat/mtproto"
	"open.chat/pkg/grpc_util"
)

func (s *Service) makeContentBySend(userId int32, authKeyId int64, content *mtproto.InputBlogContent) (*mtproto.BlogContent, error) {
	var (
		err         error = mtproto.ErrButtonTypeInvalid
		blogContent *mtproto.BlogContent
		now         = int32(time.Now().Unix())
	)
	switch content.GetPredicateName() {
	case mtproto.Predicate_inputBlogPhotos:
		inputPhotos := content.GetPhotos()
		blogContent = mtproto.MakeTLBlogContentPhotos(&mtproto.BlogContent{
			Photos: make([]*mtproto.Photo, 0, len(inputPhotos)),
		}).To_BlogContent()
		for _, photo := range inputPhotos {
			if photo == nil {
				continue
			}
			switch photo.GetPredicateName() {
			case mtproto.Predicate_inputBlogPhotoFile:
				file := photo.To_InputBlogPhotoFile()
				if file != nil {
					result, e := media_client.UploadPhotoFile(authKeyId, file.GetFile())
					if e != nil {
						err = e
						log.Errorf("UploadPhoto error: %v, by %s", e, photo.DebugString())
						continue
					}
					blogContent.Photos = append(blogContent.Photos, mtproto.MakeTLPhoto(&mtproto.Photo{
						Id:         result.PhotoId,
						AccessHash: result.AccessHash,
						Date:       now,
						Sizes:      result.SizeList,
						DcId:       2,
					}).To_Photo())
				}
			case mtproto.Predicate_inputBlogPhoto:
				fid := photo.To_InputBlogPhoto().GetPhoto()
				if fid != nil {
					sizeList, e := media_client.GetPhotoSizeList(fid.GetId())
					if e != nil {
						err = e
						log.Errorf("UploadPhoto error: %v, by %s", e, photo.DebugString())
						continue
					}
					blogContent.Photos = append(blogContent.Photos, mtproto.MakeTLPhoto(&mtproto.Photo{
						Id:         fid.GetId(),
						AccessHash: fid.GetAccessHash(),
						Date:       now,
						Sizes:      sizeList,
						DcId:       2,
					}).To_Photo())
				}
			default:
			}
		}
		if len(blogContent.Photos) == 0 {
			return blogContent, err
		}
	case mtproto.Predicate_inputBlogUploadVideo:
		video := content.To_InputBlogUploadVideo()
		var uploadedDocument = mtproto.MakeTLInputMediaUploadedDocument(&mtproto.InputMedia{
			PredicateName: "",
			Constructor:   0,
			File:          video.GetFile(),
			Thumb:         video.GetThumb(),
			MimeType:      video.GetMimeType(),
			Attributes: []*mtproto.DocumentAttribute{
				mtproto.MakeTLDocumentAttributeFilename(&mtproto.DocumentAttribute{
					FileName: video.GetFile().GetName(),
				}).To_DocumentAttribute(),
				mtproto.MakeTLDocumentAttributeVideo(&mtproto.DocumentAttribute{
					W:                 video.GetW(),
					H:                 video.GetH(),
					SupportsStreaming: video.GetSupportsStreaming(),
					Duration:          video.GetDuration(),
				}).To_DocumentAttribute(),
			},
		})
		media, e := media_client.UploadedDocumentMedia(authKeyId, uploadedDocument)
		if e != nil {
			log.Errorf("UploadVideo error: %v, by %s", e, uploadedDocument.DebugString())
			return nil, e
		}
		blogContent = mtproto.MakeTLBlogContentDocument(&mtproto.BlogContent{
			Document: media.GetDocument(),
		}).To_BlogContent()
	case mtproto.Predicate_inputBlogVideo:
		input_id := content.To_InputBlogVideo().GetId()
		document, e := media_client.GetDocumentById(input_id.GetId(), input_id.GetAccessHash())
		if e != nil {
			err = e
			log.Errorf("UploadPhoto error: %v, by %s", e, input_id.DebugString())
			return nil, err
		}
		blogContent = mtproto.MakeTLBlogContentDocument(&mtproto.BlogContent{
			Document: document,
		}).To_BlogContent()
	default:
		return nil, err
	}
	return blogContent, nil
	/*
		file := request.GetVideo()
		if file != nil {
			document, e := media_client.UploadVideoFile(authKeyId, file)
			if e != nil {
				log.Errorf("UploadVideo error: %v, by %s", e, file.DebugString())
				return nil, e
			}
			blogContent = mtproto.MakeTLBlogContentDocument(&mtproto.BlogContent{
				Document: document,
			}).To_BlogContent()
		} else if len(request.Photos) > 0 {
			now := int32(time.Now().Unix())
			for i := 0; i < len(request.Photos); i++ {
				file = request.Photos[i]
				if file != nil {
					result, e := media_client.UploadPhotoFile(authKeyId, file)
					if e != nil {
						err = e
						log.Errorf("UploadPhoto error: %v, by %s", e, file.DebugString())
						continue
					}
					blogContent.Photos = append(blogContent.Photos, mtproto.MakeTLPhoto(&mtproto.Photo{
						Id:         result.PhotoId,
						AccessHash: result.AccessHash,
						Date:       now,
						Sizes:      result.SizeList,
						DcId:       2,
					}).To_Photo())
				}
			}
			if len(blogContent.Photos) == 0 {
				return blogContent, err
			}
		}
		return blogContent, nil
	*/
}

// blogs.sendBlog#3fd43460 flags:# visible_type:VisibleType mention_users:Vector<int> text:flags.0?string random_id:long content:InputBlogContent geo_point:flags.1?InputGeoPoint address:flags.1?string = Updates;
func (s *Service) BlogsSendBlog(ctx context.Context, request *mtproto.TLBlogsSendBlog) (*mtproto.Updates, error) {
	md := grpc_util.RpcMetadataFromIncoming(ctx)
	log.Debugf("blogs.sendBlog - metadata: %s, request: %s", md.DebugString(), request.DebugString())

	var (
		mentionUsers model.IDList
		err          error
	)

	isBanned, err := s.BlogFacade.IsBannedUser(ctx, md.UserId)
	if err != nil {
		log.Errorf("blogs.sendBlog - check banned user, %v", err)
		return nil, err
	}
	if isBanned {
		log.Errorf("blogs.sendBlog - %d banned", md.UserId)
		return nil, errors.New("banned")
	}

	visible := model.FromVisibleType(md.GetUserId(), request.VisibleType)
	switch visible.VisibleType {
	case model.VisibleType_Public:
	case model.VisibleType_Private:
	case model.VisibleType_Friend:
	case model.VisibleType_Fans:
	case model.VisibleType_Allow:
	case model.VisibleType_NotAllow:
		{
			break
		}
	default:
		{
			err = mtproto.ErrButtonTypeInvalid
			log.Errorf("blogs.sendBlog: %v", err)
			return nil, err
		}
	}

	text := request.GetText().GetValue()
	if len(text) > 4000 {
		err = mtproto.ErrMessageTooLong
		log.Errorf("blogs.sendBlog: %v", err)
		return nil, err
	}
	text = model.FixString(ctx, text)

	me := md.GetUserId()
	blogContent, err := s.makeContentBySend(me, md.GetAuthId(), request.GetContent())
	if err != nil {
		return nil, err
	}

	if visible.VisibleType != model.VisibleType_Private {
		users := s.UserFacade.GetUserListByIdList(ctx, me, request.GetMentionUsers())
		mentionUsers = make([]int32, 0, len(users))
		for _, user := range users {
			if !user.Deleted && user.GetId() != me {
				mentionUsers = append(mentionUsers, user.GetId())
			}
		}
	}

	var gep *mtproto.BlogGeoPoint
	if request.GetGeoPoint() != nil {
		gep = mtproto.MakeTLBlogGeoPoint(&mtproto.BlogGeoPoint{
			GeoPoint: model.MakeGeoPointByInput(request.GetGeoPoint()),
			Address:  request.GetAddress().GetValue(),
		}).To_BlogGeoPoint()
	}

	var topics []*mtproto.Blogs_Topic
	var topicIds []int32
	topics, topicIds, err = s.handleTopics(ctx, request)
	if err != nil {
		return nil, err
	}

	entities := model.FixInputMentionNameEntity(request.Entities)

	/*
		func() (model.UserHelper, []int32) {
			mutualUsers := []int32{}
			if visible.VisibleType == model.VisibleType_Friend || visible.VisibleType == model.VisibleType_Allow || visible.VisibleType == model.VisibleType_NotAllow {
				mutualUsers = s.UserFacade.GetMutualContactUserIdList(ctx, true, me)
			}
			return s.UserFacade, mutualUsers
		}
	*/
	randomId := request.GetRandomId()
	mentionUsers, blog, err := s.BlogFacade.SendBlog(ctx, me, randomId, visible, text, entities, mentionUsers, blogContent, gep, topicIds)
	if err != nil {
		log.Errorf("blogs.sendBlog#3fd43460 - error: %v", err)
		return nil, err
	}

	up := mtproto.MakeTLUpdateNewBlog(&mtproto.Update{
		Blog:      blog,
		RandomId:  randomId,
		Pts_INT32: int32(idgen.NextBlogPtsId(ctx, me)),
		PtsCount:  1,
	}).To_Update()
	updateMe := model.MakeUpdatesHelper(up)
	reply := updateMe.ToReplyUpdates(ctx, me, s.UserFacade, nil, nil)

	for _, topic := range topics {
		updateMe.PushBackUpdate(mtproto.MakeTLUpdateTopic(&mtproto.Update{
			Topic: topic,
		}).To_Update())
	}

	for _, uid := range mentionUsers {
		if uid == me {
			sync_client.SyncUpdatesNotMe(ctx, me, md.GetAuthId(), reply)
		} else {
			up2 := proto.Clone(up).(*mtproto.Update)
			up2.Pts_INT32 = int32(idgen.NextBlogPtsId(ctx, uid))
			if up2.Blog != nil {
				up2.Blog.Mentioned = true
			}
			updateUser := model.MakeUpdatesHelper(up2)
			ups := updateUser.ToPushUpdates(ctx, uid, s.UserFacade, nil, nil)
			sync_client.PushUpdates(ctx, uid, ups)
		}
	}

	log.Debugf("blogs.sendBlog#3fd43460 - reply: %s", reply.DebugString())
	return reply, err
}

func (s *Service) handleTopics(ctx context.Context, request *mtproto.TLBlogsSendBlog) ([]*mtproto.Blogs_Topic, []int32, error) {
	var err error
	var topics []*mtproto.Blogs_Topic
	if request.GetTopics() != nil {
		log.Debugf("blogs.sendBlog, handleTopics, topics: %v", request.GetTopics())
		topics, err = s.BlogFacade.TouchTopics(ctx, request.GetTopics())
		if err != nil {
			log.Errorf("blogs.sendBlog#3fd43460, touch topics error: %v", err)
			return nil, nil, err
		}
	}
	topicIds := make([]int32, len(topics))
	for i, t := range topics {
		topicIds[i] = t.Id
	}
	return topics, topicIds, nil
	//return nil, nil, nil
}
