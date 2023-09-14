package mysqldao

import (
	"context"

	"open.chat/app/json/db/dbo"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/log"
)

// DiscoverDAO .
type DiscoverGroupDAO struct {
	db *sqlx.DB
}

// NewDiscoverDAO .
func NewDiscoverGroupDAO(db *sqlx.DB) *DiscoverGroupDAO {
	return &DiscoverGroupDAO{db}
}

// SelectByChannelId .
func (dao *DiscoverGroupDAO) SelectByChannelId(ctx context.Context, channelID uint32) (rList []dbo.DiscoverGroupDO, err error) {
	var (
		query = "SELECT id, channel_id, name, created_at FROM `discover_groups` WHERE channel_id = ? AND deleted = 0 AND `disable` = 0 ORDER BY `sort`"
		rows  *sqlx.Rows
	)
	rows, err = dao.db.Query(ctx, query, channelID)

	if err != nil {
		log.Errorf("queryx in SelectByChannelId(_), error: %v", err)
		return
	}

	defer rows.Close()

	rList = make([]dbo.DiscoverGroupDO, 0)
	for rows.Next() {
		v := dbo.DiscoverGroupDO{}
		err = rows.StructScan(&v)
		if err != nil {
			log.Errorf("structScan in SelectByChannelId(_), error: %v", err)
		}
		v.CreateAt = (uint32)(v.CreateTime.Unix())
		rList = append(rList, v)
	}

	return
}
