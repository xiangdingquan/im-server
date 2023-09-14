package core

import (
	"context"
	"math"
	"open.chat/app/service/biz_service/blog/internal/dal/dataobject"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

func (m *BlogCore) TouchTopics(ctx context.Context, nameList []string) ([]*mtproto.Blogs_Topic, error) {
	log.Debugf("BlogCore.TouchTopics, nameList:%v", nameList)
	l := make([]*dataobject.BlogTopicsDO, len(nameList))
	for i, name := range nameList {
		l[i] = &dataobject.BlogTopicsDO{
			Name: name,
		}
	}

	err := m.BlogTopicsDAO.InsertOrUpdate(ctx, l)
	if err != nil {
		return nil, err
	}

	l, err = m.BlogTopicsDAO.SelectByName(ctx, nameList)
	if err != nil {
		return nil, err
	}

	out := make([]*mtproto.Blogs_Topic, len(l))
	for i, do := range l {
		out[i] = mtproto.MakeTLBlogsTopic(&mtproto.Blogs_Topic{
			Id:      do.Id,
			Name:    do.Name,
			Ranking: do.Ranking,
		}).To_Blogs_Topic()
	}
	return out, nil
}

func (m *BlogCore) GetTopics(ctx context.Context, fromTopicId, limit int32) (*mtproto.Blogs_Topics, error) {
	if fromTopicId == 0 {
		fromTopicId = math.MaxInt32
	}
	l, err := m.BlogTopicsDAO.Select(ctx, fromTopicId, limit)
	if err != nil {
		return nil, mtproto.ErrInternelServerError
	}

	c, err := m.BlogTopicsDAO.Count(ctx)
	if err != nil {
		return nil, mtproto.ErrInternelServerError
	}

	return m.topicListDoToMTProto(l, c), nil
}

func (m *BlogCore) topicDoToMTProto(do *dataobject.BlogTopicsDO) *mtproto.Blogs_Topic {
	return mtproto.MakeTLBlogsTopic(&mtproto.Blogs_Topic{
		Id:      do.Id,
		Name:    do.Name,
		Ranking: do.Ranking,
	}).To_Blogs_Topic()
}

func (m *BlogCore) topicListDoToMTProto(doList []*dataobject.BlogTopicsDO, count int32) *mtproto.Blogs_Topics {
	out := mtproto.MakeTLBlogsTopics(&mtproto.Blogs_Topics{
		Count:  count,
		Topics: make([]*mtproto.Blogs_Topic, len(doList)),
	}).To_Blogs_Topics()
	for i, do := range doList {
		out.Topics[i] = m.topicDoToMTProto(do)
	}
	return out
}

func (m *BlogCore) GetHotTopics(ctx context.Context, limit int32) (*mtproto.Blogs_Topics, error) {
	l, err := m.BlogTopicsDAO.SelectOrdered(ctx, limit)
	if err != nil {
		return nil, mtproto.ErrInternelServerError
	}

	c, err := m.BlogTopicsDAO.Count(ctx)
	if err != nil {
		return nil, mtproto.ErrInternelServerError
	}

	return m.topicListDoToMTProto(l, c), nil
}

func (m *BlogCore) GetTopicIdByName(ctx context.Context, name string) (int32, error) {
	l, err := m.BlogTopicsDAO.SelectByName(ctx, []string{name})
	if err != nil {
		return 0, err
	}

	if len(l) == 0 {
		return 0, nil
	}

	return l[0].Id, nil
}
