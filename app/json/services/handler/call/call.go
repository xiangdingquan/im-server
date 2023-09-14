package call

import (
	"context"
	"time"

	"open.chat/app/json/consts"
	svc "open.chat/app/json/service"
	rtctokenbuilder "open.chat/app/json/services/handler/call/RtcTokenBuilder"
	"open.chat/app/json/services/handler/call/core"
	"open.chat/pkg/grpc_util"
	"open.chat/pkg/util"

	"open.chat/app/json/services/handler/call/dao"

	"open.chat/app/json/helper"
	"open.chat/app/json/service/handler"

	user_client "open.chat/app/service/biz_service/user/client"
)

type (
	ThirdAgoraConfig struct {
		Appid string
		Token string
	}

	cls struct {
		*core.AVCallCore
		user_client.UserFacade
	}
)

var G_AgoraToken *ThirdAgoraConfig = nil

// New .
func New(s *svc.Service) {
	c := &cls{
		AVCallCore: core.New(nil),
	}
	s.AppendServices(handler.RegisterCall(c))
	var err error
	c.UserFacade, err = user_client.NewUserFacade("local")
	helper.CheckErr(err)
}

func (s *cls) CreateToken(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TAvCallToken) *helper.ResultJSON {
	appID := G_AgoraToken.Appid                                       //"26c305baf6d64e9babbfc55912368e17" //"161db627fec34f12b6080582682546cb"                       //"970CA35de60c44645bbae8a215061b33"
	appCertificate := G_AgoraToken.Token                              //"50c9604682f2410584a61586a5895290" //"f8da7545822341d08e919eaefe1d98e0"              //"5CFd2fd1755d40ecb72977518be15d3b"
	expireTimestamp := uint32(time.Now().UTC().Unix()) + uint32(1800) //30分钟
	token, _ := rtctokenbuilder.BuildTokenWithUID(appID, appCertificate, r.ChannelName, r.UserID, rtctokenbuilder.RoleAttendee, expireTimestamp)
	var data = struct {
		UserID      uint32 `json:"-"` // userId
		ChannelName string `json:"-"` // channelName
		Token       string `json:"token"`
	}{
		UserID:      r.UserID,
		ChannelName: r.ChannelName,
		Token:       token,
	}
	return &helper.ResultJSON{Code: 200, Msg: "success", Data: data}
}

func (s *cls) Create(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TAvCallCreate) *helper.ResultJSON {
	if len(r.To) < 1 {
		return &helper.ResultJSON{Code: -1, Msg: "recipient cannot be empty"}
	}
	//to不能包含所有者
	var ownerUID = uint32(md.UserId)
	var tos []uint32
	for _, uid := range r.To {
		if uid != ownerUID {
			//检查对方是否为黑名单
			if s.IsBlockedByUser(ctx, int32(uid), md.UserId) {
				return &helper.ResultJSON{Code: 501, Msg: "user is blacklist,please check!"}
			}
			tos = append(tos, uid)
		}
	}
	if len(tos) < 1 {
		return &helper.ResultJSON{Code: -2, Msg: "not call you self"}
	}

	callID, err := s.AVCallCore.Create(ctx, r.ChannelName, r.ChatID, ownerUID, tos, r.IsMeetingAV, r.IsVideo)
	if callID == 0 || err != nil {
		return &helper.ResultJSON{Code: -3, Msg: "create av call fail"}
	}
	cdo, err := s.AVCallCore.SelectByCallID(ctx, callID)
	if err != nil {
		return &helper.ResultJSON{Code: -4, Msg: "sql error"}
	} else if cdo == nil {
		return &helper.ResultJSON{Code: -5, Msg: "call id not exist"}
	}
	//通知用户被邀请
	var data = handler.TAvOnInvite{
		CallID:      cdo.ID,
		ChannelName: cdo.ChannelName,
		From:        cdo.OwnerUID,
		To:          cdo.Members,
		ChatID:      cdo.ChatID,
		IsMeetingAV: cdo.IsMeet,
		IsVideo:     cdo.IsVideo,
		CreateAt:    cdo.CreateAt,
	}
	rd := helper.TrelayData{
		Action: consts.ActionCallOnInvite,
		From:   uint32(md.UserId),
		To:     r.To,
	}
	err = rd.PushUpdates(ctx, data)
	if err != nil {
		return &helper.ResultJSON{Code: -6, Msg: "create notice message fail"}
	}

	return &helper.ResultJSON{Code: 200, Msg: "success", Data: data}
}

func (s *cls) GetInfo(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TAvCallID) *helper.ResultJSON {
	cdo, err := s.AVCallCore.SelectByCallID(ctx, r.CallID)
	if err != nil {
		return &helper.ResultJSON{Code: -1, Msg: "sql error"}
	} else if cdo == nil {
		return &helper.ResultJSON{Code: -2, Msg: "call id not exist"}
	}
	//fmt.Printf("%v", do)
	var data = struct {
		CallID      uint32   `json:"callId"`      //记录标识
		ChannelName string   `json:"channelName"` //通话唯一标识
		From        uint32   `json:"from"`        //通话发起者 userid
		To          []uint32 `json:"to"`          //接收者 userids
		ChatID      uint32   `json:"chatId"`      //归属会话 可为0
		IsMeetingAV bool     `json:"isMeetingAV"` //会话是否为多人
		IsVideo     bool     `json:"isVideo"`     //是否开启视频
		IsClose     bool     `json:"isClose"`     //是否已取消通话
		CreateAt    uint32   `json:"createAt"`    //创建成功的时间
	}{
		CallID:      cdo.ID,
		ChannelName: cdo.ChannelName,
		From:        cdo.OwnerUID,
		To:          cdo.Members,
		ChatID:      cdo.ChatID,
		IsMeetingAV: cdo.IsMeet,
		IsVideo:     cdo.IsVideo,
		IsClose:     cdo.CloseAt != 0,
		CreateAt:    cdo.CreateAt,
	}

	return &helper.ResultJSON{Code: 200, Msg: "success", Data: data}
}

func (s *cls) Cancel(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TAvCallID) *helper.ResultJSON {
	cdo, err := s.AVCallCore.SelectByCallID(ctx, r.CallID)
	if err != nil {
		return &helper.ResultJSON{Code: -1, Msg: "sql error"}
	} else if cdo == nil {
		return &helper.ResultJSON{Code: -2, Msg: "call id does not exist"}
	}
	if (uint32)(md.UserId) != cdo.OwnerUID {
		return &helper.ResultJSON{Code: -3, Msg: "you're not the owner"}
	}
	if cdo.CloseAt != 0 {
		return &helper.ResultJSON{Code: -4, Msg: "cannot be cancel closed in the call"}
	}

	err = s.AVCallCore.Cancel(ctx, cdo)
	if err != nil {
		return &helper.ResultJSON{Code: -5, Msg: "cancel fail"}
	}

	//通知用户取消了
	var data = struct {
		CallID      uint32 `json:"callId"`      //记录标识
		ChannelName string `json:"channelName"` //通话标识
		ChatID      uint32 `json:"chatId"`      //归属会话 可为0
		CreateAt    uint32 `json:"createAt"`    //创建成功的时间
	}{
		CallID:      cdo.ID,
		ChannelName: cdo.ChannelName,
		ChatID:      cdo.ChatID,
		CreateAt:    cdo.CreateAt,
	}
	rd := helper.TrelayData{
		Action: consts.ActionCallOnCancel,
		From:   cdo.OwnerUID,
		To:     cdo.Members,
	}
	err = rd.PushUpdates(ctx, data)
	if err != nil {
		return &helper.ResultJSON{Code: -7, Msg: "notice message fail"}
	}

	return &helper.ResultJSON{Code: 200, Msg: "success"}
}

func (s *cls) Start(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TAvCallID) *helper.ResultJSON {
	cdo, err := s.AVCallCore.SelectByCallID(ctx, r.CallID)
	if err != nil {
		return &helper.ResultJSON{Code: -1, Msg: "sql error"}
	}
	if cdo == nil {
		return &helper.ResultJSON{Code: -2, Msg: "call id does not exist"}
	}
	if cdo.CloseAt != 0 {
		return &helper.ResultJSON{Code: -3, Msg: "current call is closed"}
	}

	rdo, err := s.AvcallsRecordsDAO.Select(ctx, cdo.ID, (uint32)(md.UserId))
	if err != nil {
		return &helper.ResultJSON{Code: -4, Msg: "sql error"}
	}
	if rdo != nil {
		return &helper.ResultJSON{Code: -5, Msg: "you are already in a call"}
	}

	err = s.AVCallCore.Start(ctx, (uint32)(md.UserId), cdo)
	if err != nil {
		return &helper.ResultJSON{Code: -6, Msg: "start fail"}
	}

	return &helper.ResultJSON{Code: 200, Msg: "success"}
}

func (s *cls) Stop(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TAvCallID) *helper.ResultJSON {
	cdo, err := s.AVCallCore.SelectByCallID(ctx, r.CallID)
	if err != nil {
		return &helper.ResultJSON{Code: -1, Msg: "sql error"}
	} else if cdo == nil {
		return &helper.ResultJSON{Code: -2, Msg: "call id does not exist"}
	}
	if cdo.CloseAt != 0 {
		return &helper.ResultJSON{Code: -3, Msg: "the call has ended"}
	}

	rdo, err := s.AvcallsRecordsDAO.Select(ctx, cdo.ID, (uint32)(md.UserId))
	if err != nil {
		return &helper.ResultJSON{Code: -4, Msg: "sql error"}
	}

	if rdo != nil && rdo.LeaveAt != 0 {
		return &helper.ResultJSON{Code: -6, Msg: "you have left this call"}
	}

	var notice bool = true
	err = s.AVCallCore.Stop(ctx, (uint32)(md.UserId), cdo)
	if err != nil {
		return &helper.ResultJSON{Code: -7, Msg: "stop fail"}
	}

	if notice {
		//通知用户某人离开了
		var data = struct {
			From        uint32   `json:"from"`        //发起人
			To          []uint32 `json:"to"`          //邀请的成员
			CallID      uint32   `json:"callId"`      //记录标识
			ChannelName string   `json:"channelName"` //通话标识
			ChatID      uint32   `json:"chatId"`      //归属会话 可为0
			CreateAt    uint32   `json:"createAt"`    //创建成功的时间
		}{
			From:        cdo.OwnerUID,
			To:          cdo.Members,
			CallID:      cdo.ID,
			ChannelName: cdo.ChannelName,
			ChatID:      cdo.ChatID,
			CreateAt:    cdo.CreateAt,
		}
		rd := helper.TrelayData{
			Action: consts.ActionCallOnLeave,
			From:   uint32(md.UserId),
		}
		//除了自己都需要通知
		if rd.From != cdo.OwnerUID {
			rd.To = append(rd.To, cdo.OwnerUID)
		}
		for _, uid := range cdo.Members {
			if uid != rd.From {
				rd.To = append(rd.To, uid)
			}
		}
		err = rd.PushUpdates(ctx, data)
		if err != nil {
			return &helper.ResultJSON{Code: -10, Msg: "notice message fail"}
		}
	}

	return &helper.ResultJSON{Code: 200, Msg: "success"}
}

func (s *cls) QueryOffline(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TAvCallTimeOut) *helper.ResultJSON {
	crdos, err := s.AVCallCore.QueryMeOffline(ctx, (uint32)(md.UserId), r.Timeout)
	if err != nil {
		return &helper.ResultJSON{Code: -1, Msg: "sql error"}
	}
	var data = struct {
		Count   uint32                 `json:"count"`   //数量
		Records []dao.TAvCallAndRecord `json:"records"` //记录信息
	}{
		Count: (uint32)(len(crdos)),
	}
	data.Records = append(data.Records, crdos...)
	if data.Count > 0 {
		s.AvcallsRecordsDAO.UpdateRead(ctx, uint32(md.UserId))
	}
	return &helper.ResultJSON{Code: 200, Msg: "success", Data: data}
}

func (s *cls) QueryRecord(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TAvRecordPage) *helper.ResultJSON {
	rdos, err := s.AVCallCore.QueryMeRecords(ctx, (uint32)(md.UserId), r.Type, r.Count, r.Page)
	if err != nil {
		return &helper.ResultJSON{Code: -1, Msg: "sql error"}
	}
	var data = struct {
		Page      uint32                 `json:"page"`
		PageCount uint32                 `json:"pageCount"`
		Count     uint32                 `json:"count"` //数量
		Records   []dao.TAvCallAndRecord `json:"records"`
	}{
		Page:      r.Page,
		PageCount: 0,
	}
	data.Records = append(data.Records, rdos...)
	data.Count = (uint32)(len(data.Records))
	return &helper.ResultJSON{Code: 200, Msg: "success", Data: rdos}
}

func (s *cls) AckInvite(ctx context.Context, md *grpc_util.RpcMetadata, r *handler.TAvCallID) *helper.ResultJSON {
	rdo, err := s.AVCallCore.AvcallsRecordsDAO.Select(ctx, r.CallID, (uint32)(md.UserId))
	if err != nil {
		return &helper.ResultJSON{Code: -1, Msg: "sql error"}
	} else if rdo == nil {
		return &helper.ResultJSON{Code: -2, Msg: "user record does not exist"}
	}

	if rdo.IsRead {
		return &helper.ResultJSON{Code: 200, Msg: "you're act not to repeat"}
	}

	rowsAffected, _ := s.AVCallCore.AvcallsRecordsDAO.UpdateWithID(ctx, map[string]interface{}{
		"is_read": util.BoolToInt8(true),
	}, rdo.ID)
	if rowsAffected == 0 {
		return &helper.ResultJSON{Code: -1, Msg: "update offline fail"}
	}

	return &helper.ResultJSON{Code: 200, Msg: "success"}
}
