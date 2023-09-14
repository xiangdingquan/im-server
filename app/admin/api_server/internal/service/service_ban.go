package service

import (
	"context"

	"github.com/gogo/protobuf/types"

	"open.chat/app/admin/api_server/api"
	"open.chat/mtproto"
	"open.chat/pkg/log"
)

func (s *Service) ToggleBan(ctx context.Context, i *api.ToggleBan) (r *mtproto.PredefinedUser, err error) {
	log.Debugf("toggleBan - request: %s", i.DebugString())

	var (
		result mtproto.TLObject
		ok     bool
	)

	req := &mtproto.TLAuthToggleBan{
		Phone:      i.Phone,
		Predefined: i.Predefined,
		Expires:    nil,
		Reason:     nil,
	}
	if i.Expires != 0 {
		req.Expires = &types.Int32Value{Value: i.Expires}
		req.Reason = &types.StringValue{Value: i.Reason}
	}

	result, err = s.Invoke(ctx, 777000, req)
	if err != nil {
		log.Errorf("toggleBan - error: %v", err)
		return
	} else if r, ok = result.(*mtproto.PredefinedUser); !ok {
		log.Errorf("toggleBan - error: invalid Bool type")
		err = mtproto.ErrInternelServerError
		return
	}

	log.Debugf("toggleBan - reply: %s", r.DebugString())
	return
}
