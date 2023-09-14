package core

import (
	"context"

	media_client "open.chat/app/service/media/client"
	"open.chat/mtproto"
)

func (m *UserCore) GetWallPaperList(ctx context.Context) ([]*mtproto.WallPaper, error) {
	doList, err := m.WallPapersDAO.SelectAll(ctx)

	walls := make([]*mtproto.WallPaper, 0, len(doList))

	for _, wallData := range doList {
		if wallData.Type == 0 {
			szList, _ := media_client.GetPhotoSizeList(wallData.PhotoId)
			wall := mtproto.MakeTLWallPaper(&mtproto.WallPaper{
				Id_INT32: wallData.Id,
				Title:    wallData.Title,
				Sizes:    szList,
				Color:    wallData.Color,
			})
			walls = append(walls, wall.To_WallPaper())
		} else {
			wall := mtproto.MakeTLWallPaperSolid(&mtproto.WallPaper{
				Id_INT32: wallData.Id,
				Title:    wallData.Title,
				Color:    wallData.Color,
				BgColor:  wallData.BgColor,
			})
			walls = append(walls, wall.To_WallPaper())
		}
	}

	return walls, err
}
