package core

import (
	"context"

	"open.chat/app/service/biz_service/username/internal/dal/dataobject"
	"open.chat/app/service/biz_service/username/internal/dao"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/database/sqlx"
)

type UsernameCore struct {
	*dao.Dao
}

func New(dao *dao.Dao) *UsernameCore {
	return &UsernameCore{Dao: dao}
}

func (m *UsernameCore) GetListByUsernameList(ctx context.Context, names []string) (map[string]*dataobject.UsernameDO, error) {
	if doList, err := m.UsernameDAO.SelectList(ctx, names); err != nil {
		return nil, err
	} else {
		m2 := make(map[string]*dataobject.UsernameDO, len(doList))
		for i := 0; i < len(doList); i++ {
			m2[doList[i].Username] = &doList[i]
		}
		return m2, nil
	}
}

func (m *UsernameCore) CheckAccountUsername(ctx context.Context, userId int32, username string) (int, error) {
	var (
		checked = model.UsernameNotExisted
	)

	usernameDO, err := m.UsernameDAO.SelectByUsername(ctx, username)
	if usernameDO != nil {
		if usernameDO.PeerType == int8(model.PEER_USER) && usernameDO.PeerId == userId {
			checked = model.UsernameExistedIsMe
		} else {
			checked = model.UsernameExistedNotMe
		}
	}

	return checked, err
}

func (m *UsernameCore) CheckChannelUsername(ctx context.Context, channelId int32, username string) (int, error) {
	var (
		checked = model.UsernameNotExisted
	)

	usernameDO, err := m.UsernameDAO.SelectByUsername(ctx, username)
	if usernameDO != nil {
		if usernameDO.PeerType == int8(model.PEER_CHANNEL) && usernameDO.PeerId == channelId {
			checked = model.UsernameExistedIsMe
		} else {
			checked = model.UsernameExistedNotMe
		}
	}

	return checked, err
}

func (m *UsernameCore) UpdateUsernameByPeer(ctx context.Context, peerType, peerId int32, username string) (bool, error) {
	if username == "" {
		m.UsernameDAO.Delete2(ctx, int8(peerType), peerId)
		return true, nil
	} else {
		tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
			do := &dataobject.UsernameDO{
				PeerType: int8(peerType),
				PeerId:   peerId,
				Username: username,
			}
			m.UsernameDAO.Delete2Tx(tx, int8(peerType), peerId)
			_, _, err := m.UsernameDAO.InsertTx(tx, do)
			result.Err = err
		})

		if tR.Err != nil {
			if sqlx.IsDuplicate(tR.Err) {
				return false, nil
			} else {
				return false, tR.Err
			}
		}

		return true, nil
	}
}

func (m *UsernameCore) GetAccountUsername(ctx context.Context, userId int32) (username string, err error) {
	do, _ := m.UsernameDAO.SelectByPeer(ctx, int8(model.PEER_USER), userId)
	if do != nil {
		username = do.Username
	}
	return
}

func (m *UsernameCore) GetChannelUsername(ctx context.Context, channelId int32) (username string, err error) {
	do, _ := m.UsernameDAO.SelectByPeer(ctx, int8(model.PEER_CHANNEL), channelId)
	if do != nil {
		username = do.Username
	}
	return
}

func (m *UsernameCore) CheckUsername(ctx context.Context, username string) (int, error) {
	var checked = model.UsernameNotExisted

	usernameDO, err := m.UsernameDAO.SelectByUsername(ctx, username)
	if usernameDO != nil {
		checked = model.UsernameExisted
	}

	return checked, err
}

func (m *UsernameCore) UpdateUsername(ctx context.Context, peerType, peerId int32, username string) (bool, error) {
	do := &dataobject.UsernameDO{
		Username: username,
		PeerType: int8(peerType),
		PeerId:   peerId,
	}
	_, _, err := m.UsernameDAO.Insert(ctx, do)
	if err != nil {
		if sqlx.IsDuplicate(err) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func (m *UsernameCore) DeleteUsername(ctx context.Context, username string) (bool, error) {
	rowsAffected, err := m.UsernameDAO.Delete(ctx, username)
	if err != nil {
		return false, err
	}
	return rowsAffected > 0, nil
}

func (m *UsernameCore) ResolveUsername(ctx context.Context, username string) (*model.PeerUtil, error) {
	var (
		peer *model.PeerUtil
		err  error
	)

	switch username {
	case "gif":
	case "pic":
	case "bing":
	case "":
	}
	if len(username) >= 5 {
		usernameDO, _ := m.UsernameDAO.SelectByUsername(ctx, username)
		if usernameDO != nil {
			if usernameDO.PeerType == model.PEER_USER || usernameDO.PeerType == model.PEER_CHANNEL {
				peer = &model.PeerUtil{
					PeerType: int32(usernameDO.PeerType),
					PeerId:   usernameDO.PeerId,
				}
			} else {
				err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USERNAME_NOT_OCCUPIED)
			}
		} else {
			err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USERNAME_NOT_OCCUPIED)
		}
	} else {
		err = mtproto.NewRpcError2(mtproto.TLRpcErrorCodes_USERNAME_INVALID)
	}

	return peer, err
}

func (m *UsernameCore) DeleteUsernameByPeer(ctx context.Context, peerType, peerId int32) error {
	_, err := m.UsernameDAO.Delete2(ctx, int8(peerType), peerId)
	return err
}
