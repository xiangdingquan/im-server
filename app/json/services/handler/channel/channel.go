package channel

import (
	"context"
	"encoding/json"
	user_client "open.chat/app/service/biz_service/user/client"
	"open.chat/mtproto"
	"time"

	"open.chat/app/json/helper"
	svc "open.chat/app/json/service"
	"open.chat/model"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/log"

	"open.chat/app/json/service/handler"

	msg_facade "open.chat/app/messenger/msg/facade"
	channel_facade "open.chat/app/service/biz_service/channel/facade"
)

type cls struct {
	msg_facade.MsgFacade
	channel_facade.ChannelFacade
	user_client.UserFacade
}

// New .
func New(s *svc.Service) {
	service := &cls{}
	var err error
	service.MsgFacade, err = msg_facade.NewMsgFacade("emsg")
	if err != nil {
		panic(err)
	}
	service.ChannelFacade, err = channel_facade.NewChannelFacade("local")
	if err != nil {
		panic(err)
	}
	service.UserFacade, err = user_client.NewUserFacade("local")
	if err != nil {
		panic(err)
	}
	s.AppendServices(handler.RegisterChannel(service))
}

func (s *cls) CleanMessages(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TChannelDeleteSelfMessage) *helper.ResultJSON {
	msgIds := make([]int32, len(r.MessageIds))
	for i, id := range r.MessageIds {
		msgIds[i] = int32(id)
	}

	_, err := s.DeleteMessages(ctx, md.UserId, 0, model.MakeChannelPeerUtil(int32(r.ChannelID)), false, msgIds)
	if err != nil {
		log.Errorf("channels.deleteMessages - error: %v", err)
		return &helper.ResultJSON{Code: 200, Msg: "clean message fail"}
	}

	return &helper.ResultJSON{Code: 200, Msg: "success"}
}

func (s *cls) CountOnline(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TChannelCountOnline) *helper.ResultJSON {
	if data, err := json.Marshal(r); err != nil {
		log.Errorf("channel.countOnline - metadata: %s, request: %v", md.DebugString(), err)
	} else {
		log.Debugf("channel.countOnline - metadata: %s, request: %s", md.DebugString(), data)
	}
	channel, err := s.ChannelFacade.GetMutableChannel(ctx, r.ChannelID)
	if err != nil {
		log.Errorf("channel.countOnline - error: %v", err)
		return &helper.ResultJSON{Code: 200, Msg: "count online failed"}
	}

	filter := &mtproto.ChannelParticipantsFilter{PredicateName: mtproto.Predicate_channelParticipantsRecent}

	var (
		onlineCount  int32
		total        int32
		participants []*model.ImmutableChannelParticipant
	)

	f := func(offset int32, limit int32) {
		total, participants = s.ChannelFacade.GetChannelParticipants(ctx, channel.Channel, filter, offset, limit)
		idList := make([]int32, len(participants))
		for i, p := range participants {
			idList[i] = p.UserId
		}
		t := uint32(time.Now().UTC().Unix()) - uint32(120) // 120������������
		m := s.UserFacade.GetLastSeenList(ctx, idList)
		for k, v := range m {
			log.Debugf("channel.countOnline - uid:%d, last_seen_at:%d, dead_line:%d", k, v, t)
			if uint32(v) > t {
				onlineCount++
			}
		}
	}

	offset := int32(0)
	limit := int32(1000)
	f(offset, limit)
	offset += limit
	for offset < total {
		f(offset, limit)
		offset += limit
	}

	var data = struct {
		Count int32 `json:"count"`
	}{
		Count: onlineCount,
	}
	return &helper.ResultJSON{Code: 200, Msg: "success", Data: data}
}
