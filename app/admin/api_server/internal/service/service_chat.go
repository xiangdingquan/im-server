package service

import (
	"context"

	"open.chat/app/admin/api_server/api"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

func (s *Service) GetAllChats(ctx context.Context, i *api.GetAllChats) (r *mtproto.Messages_Chats, err error) {
	log.Debugf("getAllChats - request: %s", i.DebugString())

	var (
		result mtproto.TLObject
		ok     bool
	)

	req := &mtproto.TLMessagesGetAllChats{
		ExceptIds: []int32{},
	}

	result, err = s.Invoke(ctx, i.Id, req)
	if err != nil {
		log.Errorf("getAllChats - error: %v", err)
		return
	} else if r, ok = result.(*mtproto.Messages_Chats); !ok {
		log.Errorf("getAllChats - error: invalid mtproto.Messages_Chats type")
		err = mtproto.ErrInternelServerError
		return
	}

	if r == nil {
		log.Errorf("getAllChats - error: invalid mtproto.Messages_Chats type")
		err = mtproto.ErrInternelServerError
		return
	}

	log.Debugf("getAllChats - reply: %s", r.DebugString())
	return
}
