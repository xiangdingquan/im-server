package core

import (
	"context"

	"open.chat/app/json/services/handler/chats/dao"
)

// ChatsCore .
type (
	ChatsCore struct {
		*dao.Dao
	}

	BannedRights struct {
		BanWhisper         bool `json:"banWhisper"`         //禁止私聊
		BanSendWebLink     bool `json:"banSendWebLink"`     //禁止发送网页链接
		BanSendQRcode      bool `json:"banSendQRcode"`      //禁止发送二维码
		BanSendKeyword     bool `json:"banSendKeyword"`     //禁止发送关键字
		BanSendDmMention   bool `json:"banSendDmMention"`   //禁止发送dm@
		KickWhoSendKeyword bool `json:"kickWhoSendKeyword"` //发送敏感词移出群聊
		ShowKickMessage    bool `json:"showKickMessage"`    //敏感词移出群聊提示
	}
)

func makeRightsObject(right uint32) *BannedRights {
	return &BannedRights{
		BanWhisper:         (right & (1 << 0)) != 0,
		BanSendWebLink:     (right & (1 << 1)) != 0,
		BanSendQRcode:      (right & (1 << 2)) != 0,
		BanSendKeyword:     (right & (1 << 3)) != 0,
		BanSendDmMention:   (right & (1 << 4)) != 0,
		KickWhoSendKeyword: (right & (1 << 5)) != 0,
		ShowKickMessage:    (right & (1 << 6)) != 0,
	}
}

func makeRightsNumber(r *BannedRights) (right uint32) {
	right = 0
	if r.BanWhisper {
		right |= (1 << 0)
	}
	if r.BanSendWebLink {
		right |= (1 << 1)
	}
	if r.BanSendQRcode {
		right |= (1 << 2)
	}
	if r.BanSendKeyword {
		right |= (1 << 3)
	}
	if r.BanSendDmMention {
		right |= (1 << 4)
	}
	if r.KickWhoSendKeyword {
		right |= (1 << 5)
	}
	if r.ShowKickMessage {
		right |= (1 << 6)
	}
	return
}

// New .
func New(d *dao.Dao) *ChatsCore {
	if d == nil {
		d = dao.New()
	}
	return &ChatsCore{d}
}

func (c *ChatsCore) GetChatBannedRights(ctx context.Context, cid uint32) *BannedRights {
	right, err := c.ChatsDAO.SelectChatBannedRights(ctx, cid)
	if err != nil {
		return nil
	}
	return makeRightsObject(right)
}

func (c *ChatsCore) UpdateChatBannedRights(ctx context.Context, cid uint32, r *BannedRights) (err error) {
	if r == nil {
		r = makeRightsObject(0)
	}
	err = c.ChatsDAO.UpdateChatBannedRights(ctx, cid, makeRightsNumber(r))
	return
}

func (c *ChatsCore) GetChannelBannedRights(ctx context.Context, cid uint32) *BannedRights {
	right, err := c.ChatsDAO.SelectChannelBannedRights(ctx, cid)
	if err != nil {
		return nil
	}
	return makeRightsObject(right)
}

func (c *ChatsCore) UpdateChannelBannedRights(ctx context.Context, cid uint32, r *BannedRights) (err error) {
	if r == nil {
		r = makeRightsObject(0)
	}
	err = c.ChatsDAO.UpdateChannelBannedRights(ctx, cid, makeRightsNumber(r))
	return
}
