package mysqldao

import (
	"context"

	"open.chat/app/json/db/dbo"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

// DiscoverMenuDAO .
type DiscoverMenuDAO struct {
	db *sqlx.DB
}

// NewDiscoverMenuDAO .
func NewDiscoverMenuDAO(db *sqlx.DB) *DiscoverMenuDAO {
	return &DiscoverMenuDAO{db}
}

// SelectByChannelId .
func (dao *DiscoverMenuDAO) SelectByChannelId(ctx context.Context, channelID uint32) (rList map[uint32][]dbo.DiscoverMenus, err error) {
	var (
		query = "SELECT id, channel_id, group_id, category, title, logo, created_at, url FROM `discover_menus` WHERE channel_id = ? AND deleted = 0 AND `disable` = 0 ORDER BY `sort`"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, channelID)

	if err != nil {
		log.Errorf("queryx in SelectByChannelId(_), error: %v", err)
		return
	}

	defer rows.Close()

	rList = make(map[uint32][]dbo.DiscoverMenus)
	for rows.Next() {
		v := dbo.DiscoverMenus{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByChannelId(_), error: %v", err)
		}
		v.CreateAt = (uint32)(v.CreateTime.Unix())
		rList[v.GroupID] = append(rList[v.GroupID], v)
	}

	return
}
