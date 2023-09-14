package discover

import (
	"context"

	"open.chat/app/json/helper"
	svc "open.chat/app/json/service"
	"open.chat/app/json/services/handler/discover/core"
	"open.chat/pkg/grpc_util"

	"open.chat/app/json/service/handler"
)

type (
	menuInfo struct {
		Title string `json:"title"`
		Icon  string `json:"icon"`
		Url   string `json:"url"`
	}

	menuGroup struct {
		Title string     `json:"title"`
		Menus []menuInfo `json:"menus"`
	}

	cls struct {
		*core.DiscoverCore
	}
)

// New .
func New(s *svc.Service) {
	service := &cls{
		DiscoverCore: core.New(nil),
	}
	s.AppendServices(handler.RegisterDiscover(service))
}

func (s *cls) List(ctx context.Context, md *grpc_util.RpcMetadata) *helper.ResultJSON {
	var channelId uint32 = 0
	dgs, err := s.DiscoverGroupDAO.SelectByChannelId(ctx, channelId)
	if err != nil {
		return &helper.ResultJSON{Code: -1, Msg: "read group fail"}
	}

	dmmap, err := s.DiscoverMenuDAO.SelectByChannelId(ctx, channelId)
	if err != nil {
		return &helper.ResultJSON{Code: -2, Msg: "read group fail"}
	}

	var data = make([]menuGroup, len(dgs))
	for i, dg := range dgs {
		dms := dmmap[dg.ID]
		data[i] = menuGroup{
			Title: dg.Name,
			Menus: make([]menuInfo, len(dms)),
		}
		for j, dm := range dms {
			data[i].Menus[j] = menuInfo{
				Title: dm.Title,
				Icon:  dm.Logo,
				Url:   dm.Url,
			}
		}
	}
	return &helper.ResultJSON{Code: 200, Msg: "success", Data: data}
}
