package core

import (
	"context"
	"time"

	"open.chat/app/pkg/env2"
	"open.chat/app/service/biz_service/channel/dal/dataobject"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
	"open.chat/pkg/random2"
	"open.chat/pkg/util"
)

func (m *ChannelCore) addParticipant(ctx context.Context, participant *dataobject.ChannelParticipantsDO, admin, kicked, banned bool) error {
	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		var (
			participantsCount int32 = 1
			adminsCount       int32
			kickedCount       int32
			bannedCount       int32
			err               error
		)

		if admin {
			adminsCount = -1
		}

		if kicked {
			kickedCount = -1
		}

		if banned {
			bannedCount = -1
		}

		participant.Id, _, err = m.ChannelParticipantsDAO.InsertOrUpdateTx(tx, participant)
		if err != nil {
			result.Err = err
			return
		}

		_, err = m.ChannelsDAO.UpdateParticipantCountTx(tx, participantsCount, adminsCount, kickedCount, bannedCount, participant.Date2, participant.ChannelId)
		if err != nil {
			result.Err = err
			return
		}
	})
	return tR.Err
}

func (m *ChannelCore) InviteToChannel(ctx context.Context, channelId, inviterId int32, id ...int32) (*model.MutableChannel, []int32, error) {
	var (
		err         error
		date        = int32(time.Now().Unix())
		invitedList = make([]int32, 0, len(id))
	)

	channel, err := m.GetMutableChannel(ctx, channelId, append(id, inviterId)...)
	if err != nil {
		log.Errorf("InviteToChannel error - %v: (inviter_id: %d, user_id: %v)", err, inviterId, id)
		return nil, nil, err
	}

	me := channel.GetImmutableChannelParticipant(inviterId)
	if me == nil || !me.IsStateOk() {
		err = mtproto.ErrChannelPrivate
		log.Errorf("InviteToChannel error - %v: (inviter_id: %d, user_id: %v)", err, inviterId, id)
		return nil, nil, err
	}

	if !me.CanInviteUsers(date) {
		err = mtproto.ErrChatWriteForbidden
		log.Errorf("InviteToChannel error - %v: (inviter_id: %d, user_id: %d)", err, inviterId, id)
		return nil, nil, err
	}

	inviteOne := len(id) == 1
	for i, id2 := range id {
		invited := channel.GetImmutableChannelParticipant(id2)
		if invited != nil {
			if invited.IsStateOk() {
				if inviteOne && i == 0 {
					err = mtproto.ErrUserAlreadyParticipant
					log.Errorf("InviteToChannel error - %v: (inviter_id: %d, user_id: %d)", err, inviterId, id2)
					return nil, nil, err
				}
				continue
			} else if invited.IsKicked() {
				if !me.CanAdminBanUsers() {
					if inviteOne && i == 0 {
						err = mtproto.ErrUserKicked
						log.Errorf("InviteToChannel error - %v: (inviter_id: %d, user_id: %d)", err, inviterId, id2)
						return nil, nil, err
					}
					continue
				}
			}
		}

		pDO := &dataobject.ChannelParticipantsDO{
			ChannelId:     channelId,
			UserId:        id2,
			InviterUserId: inviterId,
			Date2:         date,
		}

		if channel.Channel.HiddenPrehistory {
			pDO.AvailableMinId = channel.Channel.TopMessage
			pDO.AvailableMinPts = channel.Channel.Pts
		}

		err = m.addParticipant(ctx, pDO, false, false, false)
		if err != nil {
			log.Errorf("InviteToChannel error - %v: (inviter_id: %d, user_id: %d)", err, inviterId, id2)
			return nil, nil, err
		}

		channel.Channel.ParticipantsCount += 1
		channel.Channel.Version += 1
		channel.Channel.Date = date
		channel.AddChannelParticipant(makeImmutableChannelParticipant(channel.Channel, pDO))

		invitedList = append(invitedList, id2)
	}

	return channel, invitedList, nil
}

func (m *ChannelCore) JoinChannel(ctx context.Context, channelId, joinId int32, force bool) (*model.MutableChannel, error) {
	var (
		err  error
		date = int32(time.Now().Unix())
	)

	channel, err := m.GetMutableChannel(ctx, channelId, joinId)
	if err != nil {
		log.Errorf("joinChannel error - %s: (join %d)", err, joinId)
		return nil, err
	}

	if !force && channel.Channel.Username == "" && channel.Channel.Link == "" {
		err = mtproto.ErrChannelPrivate
		log.Errorf("joinChannel error - %s: (join %d)", err, joinId)
		return nil, err
	}

	me := channel.GetImmutableChannelParticipant(joinId)
	if me != nil {
		if me.IsKicked() {
			err = mtproto.ErrChannelPrivate
			log.Errorf("joinChannel error - %s: (join %d)", err, joinId)
			return nil, err
		} else if me.IsStateOk() {
			err = mtproto.ErrUserAlreadyParticipant
			log.Errorf("joinChannel error - %s: (join %d)", err, joinId)
			return nil, err
		}
	}

	pDO := &dataobject.ChannelParticipantsDO{
		ChannelId:     channelId,
		UserId:        joinId,
		InviterUserId: joinId,
		Date2:         date,
	}

	if channel.Channel.HiddenPrehistory {
		pDO.AvailableMinId = channel.Channel.TopMessage
		pDO.AvailableMinPts = channel.Channel.Pts
	}

	err = m.addParticipant(ctx, pDO, false, false, false)
	if err != nil {
		log.Errorf("joinChannel error - %s: (join %d)", err, joinId)
		return nil, err
	}

	channel.Channel.ParticipantsCount += 1
	channel.Channel.Version += 1
	channel.Channel.Date = date
	channel.AddChannelParticipant(makeImmutableChannelParticipant(channel.Channel, pDO))
	return channel, nil
}

func (m *ChannelCore) LeaveChannel(ctx context.Context, channelId, userId int32) (*model.MutableChannel, error) {
	var (
		err  error
		date = int32(time.Now().Unix())
	)

	channel, err := m.GetMutableChannel(ctx, channelId, userId)
	if err != nil {
		log.Errorf("joinChannel error - %s: (join %d)", err, userId)
		return nil, err
	}

	me := channel.GetImmutableChannelParticipant(userId)
	if me == nil || me.IsLeft() || me.IsKicked() {
		err = mtproto.ErrChannelPrivate
		log.Errorf("leaveChannel error - %s: (%d)", err, userId)
		return nil, err
	}

	// store
	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		_, err = m.ChannelParticipantsDAO.UpdateLeaveTx(tx, date, channelId, userId)
		if err != nil {
			result.Err = err
			return
		}

		_, err = m.ChannelsDAO.UpdateParticipantCountTx(tx, -1, 0, 0, 0, date, channelId)

		if err != nil {
			result.Err = err
		}
	})

	if tR.Err != nil {
		return nil, tR.Err
	}

	channel.Channel.ParticipantsCount -= 1
	me.State = model.ChatMemberStateLeft
	me.Date = date
	me.PromotedBy = 0

	return channel, nil
}

func (m *ChannelCore) EditTitle(ctx context.Context, channelId, editUserId int32, title string) (*model.MutableChannel, error) {
	if title == "" {
		err := mtproto.ErrChatTitleEmpty
		log.Errorf("editChannelTitle error - %s: {channelId:%d, editUserId:%d, title:%s}", err, channelId, editUserId, title)
		return nil, err
	}

	date := int32(time.Now().Unix())
	channel, err := m.GetMutableChannel(ctx, channelId, editUserId)
	if err != nil {
		log.Errorf("editChannelTitle error - %s: {channelId:%d, editUserId:%d, title:%s}", err, channelId, editUserId, title)
		return nil, err
	}

	if channel.Channel.Title == title {
		err := mtproto.ErrChatNotModified
		log.Errorf("editChannelTitle error - %s: {channelId:%d, editUserId:%d, title:%s}", err, channelId, editUserId, title)
		return nil, err
	}

	me := channel.GetImmutableChannelParticipant(editUserId)
	if me == nil || !me.IsStateOk() {
		err := mtproto.ErrChannelPrivate
		log.Errorf("editChannelTitle error - %s: {channelId:%d, editUserId:%d, title:%s}", err, channelId, editUserId, title)
		return nil, err
	}

	if !me.CanChangeInfo(date) {
		err := mtproto.ErrChatAdminRequired
		log.Errorf("editChannelTitle error - %s: {channelId:%d, editUserId:%d, title:%s}", err, channelId, editUserId, title)
		return nil, err
	}

	if _, err = m.ChannelsDAO.UpdateTitle(ctx, title, date, channelId); err != nil {
		log.Errorf("editChannelTitle error - %s: {channelId:%d, editUserId:%d, title:%s}", err, channelId, editUserId, title)
		return nil, err
	}

	channel.Channel.Title = title
	channel.Channel.Date = date

	return channel, nil
}

func (m *ChannelCore) EditAbout(ctx context.Context, channelId, aboutUserId int32, about string) (*model.MutableChannel, error) {
	log.Debugf("editAbout: {channel_id: %d, editId: %d, about: %s}", channelId, aboutUserId, about)

	date := int32(time.Now().Unix())
	channel, err := m.GetMutableChannel(ctx, channelId, aboutUserId)
	if err != nil {
		log.Errorf("editChannelAbout error - %s: {channelId:%d, editUserId:%d, about:%s}", err, channelId, aboutUserId, about)
		return nil, err
	}

	if channel.Channel.About == about {
		err := mtproto.ErrChatNotModified
		log.Errorf("editChannelAbout error - %s: {channelId:%d, editUserId:%d, about:%s}", err, channelId, aboutUserId, about)
		return nil, err
	}

	me := channel.GetImmutableChannelParticipant(aboutUserId)
	if me == nil || !me.IsStateOk() {
		err := mtproto.ErrChannelPrivate
		log.Errorf("editChannelAbout error - %s: {channelId:%d, editUserId:%d, about:%s}", err, channelId, aboutUserId, about)
		return nil, err
	}

	if !me.CanChangeInfo(date) {
		err := mtproto.ErrChatAdminRequired
		log.Errorf("editChannelAbout error - %s: {channelId:%d, editUserId:%d, about:%s}", err, channelId, aboutUserId, about)
		return nil, err
	}

	if _, err := m.ChannelsDAO.UpdateAbout(ctx, about, date, channelId); err != nil {
		log.Errorf("editChannelAbout error - %s: {channelId:%d, editUserId:%d, about:%s}", err, channelId, aboutUserId, about)
		return nil, err
	}

	channel.Channel.About = about
	channel.Channel.Date = date

	return channel, nil
}

func (m *ChannelCore) EditNotice(ctx context.Context, channelId, noticeUserId int32, notice string) (*model.MutableChannel, error) {
	log.Debugf("editNotice: {channel_id: %d, editId: %d, notice: %s}", channelId, noticeUserId, notice)

	date := int32(time.Now().Unix())
	channel, err := m.GetMutableChannel(ctx, channelId, noticeUserId)
	if err != nil {
		log.Errorf("editChannelNotice error - %s: {channelId:%d, editUserId:%d, notice:%s}", err, channelId, noticeUserId, notice)
		return nil, err
	}

	if channel.Channel.Notice == notice {
		err := mtproto.ErrChatNotModified
		log.Errorf("editChannelNotice error - %s: {channelId:%d, editUserId:%d, notice:%s}", err, channelId, noticeUserId, notice)
		return nil, err
	}

	me := channel.GetImmutableChannelParticipant(noticeUserId)
	if me == nil || !me.IsStateOk() {
		err := mtproto.ErrChannelPrivate
		log.Errorf("editChannelNotice error - %s: {channelId:%d, editUserId:%d, notice:%s}", err, channelId, noticeUserId, notice)
		return nil, err
	}

	if !me.CanChangeInfo(date) {
		err := mtproto.ErrChatAdminRequired
		log.Errorf("editChannelNotice error - %s: {channelId:%d, editUserId:%d, notice:%s}", err, channelId, noticeUserId, notice)
		return nil, err
	}

	if _, err := m.ChannelsDAO.UpdateNotice(ctx, notice, date, channelId); err != nil {
		log.Errorf("editChannelNotice error - %s: {channelId:%d, editUserId:%d, notice:%s}", err, channelId, noticeUserId, notice)
		return nil, err
	}

	channel.Channel.Notice = notice
	channel.Channel.Date = date

	return channel, nil
}

func (m *ChannelCore) EditPhoto(ctx context.Context, channelId, editUserId int32, photo *mtproto.Photo) (*model.MutableChannel, error) {
	date := int32(time.Now().Unix())
	channel, err := m.GetMutableChannel(ctx, channelId, editUserId)
	if err != nil {
		log.Errorf("editChannelPhoto error - %s: {channelId:%d, editUserId:%d, title:%s}", err, channelId, editUserId, photo.GetId())
		return nil, err
	}

	me := channel.GetImmutableChannelParticipant(editUserId)
	if me == nil || !me.IsStateOk() {
		err := mtproto.ErrChannelPrivate
		log.Errorf("editChannelPhoto error - %s: {channelId:%d, editUserId:%d, title:%s}", err, channelId, editUserId, photo.GetId())
		return nil, err
	}

	if !me.CanChangeInfo(date) {
		err := mtproto.ErrChatAdminRequired
		log.Errorf("editChannelPhoto error - %s: {channelId:%d, editUserId:%d, title:%s}", err, channelId, editUserId, photo.GetId())
		return nil, err
	}

	if _, err := m.ChannelsDAO.UpdatePhoto(ctx, hack.String(model.TLObjectToJson(photo)), date, channelId); err != nil {
		log.Errorf("editChannelPhoto error - %s: {channelId:%d, editUserId:%d, title:%s}", err, channelId, editUserId, photo.GetId())
		return nil, err
	}

	channel.Channel.Date = date

	return channel, nil
}

func (m *ChannelCore) EditAdminRights(ctx context.Context, channelId, operatorId, editChannelAdminsId int32, adminRights model.ChatAdminRights, rank string) (*model.MutableChannel, bool, error) {
	var (
		err  error
		date = int32(time.Now().Unix())
	)

	if operatorId == editChannelAdminsId {
	}

	channel, err := m.GetMutableChannel(ctx, channelId, operatorId, editChannelAdminsId)
	if err != nil {
		log.Errorf("editChannelAdminRights error - %s: (%d - %d - %v)", err, operatorId, editChannelAdminsId, adminRights)
		return nil, false, err
	}

	me := channel.GetImmutableChannelParticipant(operatorId)
	if me == nil || !me.IsStateOk() {
		err = mtproto.ErrChannelPrivate
		log.Errorf("editChannelAdminRights error - %s: (%d - %d - %v)", err, operatorId, editChannelAdminsId, adminRights)
		return nil, false, err
	} else if me.IsChatMemberNormal() {
		err = mtproto.ErrChatAdminRequired
		log.Errorf("editChannelAdminRights error - %s: (%d - %d - %v)", err, operatorId, editChannelAdminsId, adminRights)
		return nil, false, err
	}

	if !me.CanAdminAddAdmins() {
		err = mtproto.ErrChatAdminInviteRequired
		log.Errorf("editChannelAdminRights error - %s: (%d - %d - %v)", err, operatorId, editChannelAdminsId, adminRights)
		return nil, false, err
	} else {
		if !me.CanAdminChangeInfo() && adminRights.CanChangeInfo() {
			err = mtproto.ErrRightForbidden
			log.Errorf("editChannelAdminRights error - %s: (%d - %d - %v)", err, operatorId, editChannelAdminsId, adminRights)
			return nil, false, err
		}
		if !me.CanAdminEditMessages() && adminRights.CanEditMessages() {
			err = mtproto.ErrRightForbidden
			log.Errorf("editChannelAdminRights error - %s: (%d - %d - %v)", err, operatorId, editChannelAdminsId, adminRights)
			return nil, false, err
		}
		if !me.CanAdminDeleteMessages() && adminRights.CanDeleteMessages() {
			err = mtproto.ErrRightForbidden
			log.Errorf("editChannelAdminRights error - %s: (%d - %d - %v)", err, operatorId, editChannelAdminsId, adminRights)
			return nil, false, err
		}
		if !me.CanAdminBanUsers() && adminRights.CanBanUsers() {
			err = mtproto.ErrRightForbidden
			log.Errorf("editChannelAdminRights error - %s: (%d - %d - %v)", err, operatorId, editChannelAdminsId, adminRights)
			return nil, false, err
		}
		if !me.CanAdminInviteUsers() && adminRights.CanInviteUsers() {
			err = mtproto.ErrRightForbidden
			log.Errorf("editChannelAdminRights error - %s: (%d - %d - %v)", err, operatorId, editChannelAdminsId, adminRights)
			return nil, false, err
		}
		if !me.CanAdminAddAdmins() && adminRights.CanAddAdmins() {
			err = mtproto.ErrRightForbidden
			log.Errorf("editChannelAdminRights error - %s: (%d - %d - %v)", err, operatorId, editChannelAdminsId, adminRights)
			return nil, false, err
		}
	}

	edited := channel.GetImmutableChannelParticipant(editChannelAdminsId)
	if edited == nil {
		pDO := &dataobject.ChannelParticipantsDO{
			ChannelId:   channelId,
			UserId:      editChannelAdminsId,
			PromotedBy:  operatorId,
			AdminRights: int32(adminRights),
			Date2:       date,
		}

		if channel.Channel.HiddenPrehistory {
			pDO.AvailableMinId = channel.Channel.TopMessage
			pDO.AvailableMinPts = channel.Channel.Pts
		}

		err = m.addParticipant(ctx, pDO, false, false, false)
		if err != nil {
			log.Errorf("editChannelAdminRights error - %s: (%d - %d - %v)", err, operatorId, editChannelAdminsId, adminRights)
			return nil, false, err
		}

		channel.Channel.ParticipantsCount += 1
		channel.Channel.Version += 1
		channel.Channel.Date = date
		channel.AddChannelParticipant(makeImmutableChannelParticipant(channel.Channel, pDO))
		return channel, true, nil
	} else {
		tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
			_, err = m.ChannelParticipantsDAO.UpdateAdminRightsTx(tx, operatorId, int32(adminRights), date, channelId, editChannelAdminsId)
			if err != nil {
				result.Err = err
				return
			}
			m.ChannelsDAO.UpdateParticipantCountTx(tx, 0, 1, 0, 0, date, channelId)
			channel.Channel.AdminsCount += 1
		})
		if tR.Err != nil {
			log.Errorf("editChannelAdminRights - edit creator: (%d - %d - %v)", operatorId, editChannelAdminsId, adminRights)
			return nil, false, tR.Err
		}

		edited.AdminRights = adminRights
		edited.PromotedBy = operatorId
		edited.Date = date
		return channel, false, nil
	}
}

func (m *ChannelCore) EditBanned(ctx context.Context,
	channelId, operatorId, bannedUserId int32,
	bannedRights model.ChatBannedRights) (*model.MutableChannel, bool, error) {
	var (
		err         error
		date        = int32(time.Now().Unix())
		editedIsNil bool
		needKick    bool
	)

	// editChatAdminId not creator
	if bannedUserId == operatorId {
		err = mtproto.ErrUserIdInvalid
		log.Errorf("editBanned error - %s: (%d - %d - %v)", err, operatorId, bannedUserId, bannedRights)
		return nil, false, err
	}

	channel, err := m.GetMutableChannel(ctx, channelId, operatorId, bannedUserId)
	if err != nil {
		log.Errorf("editBanned error - %s: (%d - %d - %v)", err, operatorId, bannedUserId, bannedRights)
		return nil, false, err
	}

	if channel.Channel.IsBroadcast() {
	}

	me := channel.GetImmutableChannelParticipant(operatorId)
	if me == nil || !me.IsStateOk() {
		err = mtproto.ErrChannelPrivate
		log.Errorf("editBanned error - %s: (%d - %d - %v)", err, operatorId, bannedUserId, bannedRights)
		return nil, false, err
	} else if me.IsChatMemberNormal() {
		err = mtproto.ErrUserAdminInvalid
		log.Errorf("editBanned error - %s: (%d - %d - %v)", err, operatorId, bannedUserId, bannedRights)
		return nil, false, err
	}

	if !me.CanAdminBanUsers() {
		err = mtproto.ErrChatAdminRequired
		log.Errorf("editBanned error - %s: (%d - %d - %v)", err, operatorId, bannedUserId, bannedRights)
		return nil, false, err
	}

	edited := channel.GetImmutableChannelParticipant(bannedUserId)
	if edited != nil && edited.Creator {
		err = mtproto.ErrUserAdminInvalid
		log.Errorf("editBanned error - %s: (%d - %d - %v)", err, operatorId, bannedUserId, bannedRights)
		return nil, false, err
	}

	if edited == nil {
		edited = &model.ImmutableChannelParticipant{
			Channel:             channel.Channel,
			Creator:             false,
			State:               model.ChatMemberStateLeft,
			Id:                  0,
			UserId:              bannedUserId,
			Date:                date,
			InviterId:           0,
			Rank:                "",
			CanEdit:             false,
			PromotedBy:          0,
			AdminRights:         0,
			KickedBy:            operatorId,
			BannedRights:        bannedRights,
			Pinned:              false,
			UnreadMark:          false,
			ChannelId:           channelId,
			TopMessage:          channel.Channel.TopMessage,
			ReadInboxMaxId:      0,
			UnreadCount:         0,
			UnreadMentionsCount: 0,
			NotifySettings:      nil,
			Draft:               nil,
			FolderId:            0,
		}
		if bannedRights.IsKick() {
			edited.State = model.ChatMemberStateKicked
		} else {
			edited.State = model.ChatMemberStateLeft
		}
		editedIsNil = true
	} else if edited.IsLeft() {
		if bannedRights.IsKick() {
			edited.State = model.ChatMemberStateKicked
		}
	} else if edited.IsKicked() {
		if bannedRights.NoBanRights() {
			edited.State = model.ChatMemberStateLeft
		}
	} else {
		if bannedRights.IsKick() {
			edited.State = model.ChatMemberStateKicked
		}
	}

	edited.KickedBy = operatorId
	edited.Date = date
	edited.BannedRights = bannedRights
	needKick = edited.State == model.ChatMemberStateLeft || edited.State == model.ChatMemberStateKicked

	// USER_NOT_PARTICIPANT
	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		if editedIsNil {
			pDO := &dataobject.ChannelParticipantsDO{
				ChannelId:       channelId,
				UserId:          bannedUserId,
				InviterUserId:   0,
				Date2:           date,
				KickedBy:        operatorId,
				BannedRights:    bannedRights.Rights,
				BannedUntilDate: bannedRights.UntilDate,
				State:           int8(edited.State),
			}
			_, _, err = m.ChannelParticipantsDAO.InsertOrUpdateTx(tx, pDO)
			if err != nil {
				log.Errorf("editBanned error - %v: (inviter_id: %d, user_id: %d)", err, operatorId, bannedUserId)
				result.Err = err
				return
			}
		} else {
			_, err = m.ChannelParticipantsDAO.UpdateBannedRightsTx(tx, operatorId, bannedRights.Rights, bannedRights.UntilDate, int8(edited.State), date, channelId, bannedUserId)
			if err != nil {
				result.Err = err
				return
			}

			if needKick {
				m.ChannelsDAO.UpdateParticipantCountTx(tx, -1, 0, 0, 0, date, channelId)
				channel.Channel.ParticipantsCount -= 1
			}
		}
	})
	if tR.Err != nil {
		return nil, false, tR.Err
	}

	return channel, needKick, nil
}

func (m *ChannelCore) ExportChannelInvite(ctx context.Context, channelId, operatorId int32) (*model.MutableChannel, error) {
	channel, err := m.GetMutableChannel(ctx, channelId, operatorId)
	if err != nil {
		log.Errorf("not found channel: %d", channelId)
		return nil, err
	}

	me := channel.GetImmutableChannelParticipant(operatorId)
	if me == nil || !me.IsStateOk() {
		err = mtproto.ErrChannelPrivate
		log.Errorf("exportedChatInvite error - operatorId(%d) %v", operatorId, err)
		return nil, err
	}

	if !me.CanAdminInviteUsers() {
		err = mtproto.ErrChatAdminInviteRequired
		log.Errorf("exportedChatInvite error - operatorId(%d) %v", operatorId, err)
		return nil, err
	}

	// retry 5
	for i := 0; i < 5; i++ {
		link := random2.RandomAlphanumeric(22)
		_, err = m.ChannelsDAO.UpdateLink(ctx, link, int32(time.Now().Unix()), channelId)

		if err != nil {
			if sqlx.IsDuplicate(err) {
				continue
			} else {
				return nil, err
			}
		} else {
			channel.Channel.Link = env2.T_ME + "/joinchat?link=" + link
			channel.Channel.Username = ""
			break
		}
	}

	if err != nil {
		return nil, err
	}

	return channel, err
}

func (m *ChannelCore) ToggleSignatures(ctx context.Context, channelId, operatorId int32, enabled bool) (*model.MutableChannel, error) {
	channel, err := m.GetMutableChannel(ctx, channelId, operatorId)
	if err != nil {
		log.Errorf("not found channel: %d", channelId)
		return nil, err
	}

	if !channel.Channel.IsBroadcast() {
		log.Errorf("toggleSignatures megagroup not signatures - %s: (%d - %v)", err, operatorId, enabled)

		err := mtproto.ErrChatWriteForbidden
		return nil, err
	}

	me := channel.GetImmutableChannelParticipant(operatorId)
	if me == nil || !me.IsStateOk() {
		err = mtproto.ErrChannelPrivate
		log.Errorf("toggleSignatures error - %s: (%d - %v)", err, operatorId, enabled)
		return nil, err
	}

	if !me.IsAdmin() && !me.IsCreator() {
		err = mtproto.ErrUserAdminInvalid
		log.Errorf("toggleSignatures error - %s: (%d - %v)", err, operatorId, enabled)
		return nil, err
	}

	channel.Channel.Signatures = enabled
	_, err = m.ChannelsDAO.UpdateSignatures(ctx, util.BoolToInt8(enabled), channel.Channel.Date, channelId)
	if err != nil {
		return nil, err
	}
	return channel, err
}

func (m *ChannelCore) ToggleInvites(ctx context.Context, channelId, operatorId int32, enabled bool) (*model.MutableChannel, error) {
	channel, err := m.GetMutableChannel(ctx, channelId, operatorId)
	if err != nil {
		log.Errorf("not found channel: %d", channelId)
		return nil, err
	}

	me := channel.GetImmutableChannelParticipant(operatorId)
	if me == nil || !me.IsStateOk() {
		err = mtproto.ErrChannelPrivate
		log.Errorf("toggleSignatures error - %s: (%d - %v)", err, operatorId, enabled)
		return nil, err
	}

	if !me.CanAdminInviteUsers() {
		err = mtproto.ErrChatAdminInviteRequired
		log.Errorf("toggleSignatures error - %s: (%d - %v)", err, operatorId, enabled)
		return nil, err
	}

	channel.Channel.Democracy = enabled
	_, err = m.ChannelsDAO.UpdateDemocracy(ctx, util.BoolToInt8(enabled), channel.Channel.Date, channelId)
	if err != nil {
		log.Errorf("toggleSignatures error - %s: (%d - %v)", err, operatorId, enabled)
		return nil, err
	}

	return channel, nil
}

func (m *ChannelCore) UpdateUsername(ctx context.Context, channelId, operatorId int32, username string) (*model.MutableChannel, error) {
	var (
		date = int32(time.Now().Unix())
	)

	channel, err := m.GetMutableChannel(ctx, channelId, operatorId)
	if err != nil {
		log.Errorf("not found channel: %d", channelId)
		return nil, err
	}

	me := channel.GetImmutableChannelParticipant(operatorId)
	if me == nil || !me.IsStateOk() {
		err = mtproto.ErrChannelPrivate
		log.Errorf("updateUsername error - %s: (%d - %s)", err, operatorId, username)
		return nil, err
	}

	if !me.CanAdminInviteUsers() {
		err = mtproto.ErrChatAdminInviteRequired
		log.Errorf("updateUsername error - %s: (%d - %s)", err, operatorId, username)
		return nil, err
	}

	channel.Channel.Date = date
	channel.Channel.Username = username
	if _, err = m.ChannelsDAO.UpdateUsername(ctx, username, date, channelId); err != nil {
		log.Errorf("updateUsername error - %s: (%d - %s)", err, operatorId, username)
		return nil, err
	}

	return channel, nil
}

func (m *ChannelCore) TogglePreHistoryHidden(ctx context.Context, channelId, operatorId int32, enabled bool) (*model.MutableChannel, error) {
	channel, err := m.GetMutableChannel(ctx, channelId, operatorId)
	if err != nil {
		log.Errorf("not found channel: %d", channelId)
		return nil, err
	}

	me := channel.GetImmutableChannelParticipant(operatorId)
	if me == nil || !me.IsStateOk() {
		err = mtproto.ErrChannelPrivate
		log.Errorf("togglePreHistoryHidden error - %s: (%d - %v)", err, operatorId, enabled)
		return nil, err
	}

	if !me.CanAdminInviteUsers() {
		err = mtproto.ErrChatAdminInviteRequired
		log.Errorf("togglePreHistoryHidden error - %s: (%d - %v)", err, operatorId, enabled)
		return nil, err
	}

	if _, err = m.ChannelsDAO.Update(ctx, map[string]interface{}{
		"pre_history_hidden": util.BoolToInt8(enabled),
	}, channelId); err != nil {
		log.Errorf("togglePreHistoryHidden error - %s: (%d - %v)", err, operatorId, enabled)
		return nil, err
	}

	channel.Channel.HiddenPrehistory = enabled
	return channel, nil
}

func (m *ChannelCore) DeleteChannel(ctx context.Context, channelId, operatorId int32) (*model.MutableChannel, error) {
	channel, err := m.GetMutableChannel(ctx, channelId, operatorId)
	if err != nil {
		log.Errorf("not found channel: %d", channelId)
		return nil, err
	}

	me := channel.GetImmutableChannelParticipant(operatorId)
	if me == nil || !me.IsStateOk() {
		err = mtproto.ErrChannelPrivate
		log.Errorf("deleteChannel error - %s: (%d - %d)", err, channelId, operatorId)
		return nil, err
	}

	if !me.IsCreator() {
		err = mtproto.ErrChatAdminRequired
		log.Errorf("deleteChannel error - %s: (%d - %d)", err, channelId, operatorId)
		return nil, err
	}

	_, err = m.ChannelsDAO.Delete(ctx, channelId)
	if err != nil {
		log.Errorf("deleteChannel error - %s: (%d - %d)", err, channelId, operatorId)
		return nil, err
	}

	channel.Channel.Deleted = true
	return channel, nil
}

func (m *ChannelCore) EditChatDefaultBannedRights(ctx context.Context, channelId, operatorId int32, bannedRights *mtproto.ChatBannedRights) (*model.MutableChannel, error) {
	var (
		date = int32(time.Now().Unix())
	)

	channel, err := m.GetMutableChannel(ctx, channelId, operatorId)
	if err != nil {
		log.Errorf("not found channel: %d", channelId)
		return nil, err
	}

	me := channel.GetImmutableChannelParticipant(operatorId)
	if me == nil || !me.IsStateOk() {
		err = mtproto.ErrChannelPrivate
		log.Errorf("deleteChannel error - %s: (%d - %d)", err, channelId, operatorId)
		return nil, err
	}

	if !me.CanAdminBanUsers() {
		err = mtproto.ErrChatAdminRequired
		log.Errorf("deleteChannel error - %s: (%d - %d)", err, channelId, operatorId)
		return nil, err
	}

	rights := model.MakeChatBannedRights(bannedRights)

	if _, err = m.ChannelsDAO.UpdateDefaultBannedRights(ctx, rights.Rights, date, channelId); err != nil {
		log.Errorf("deleteChannel error - %s: (%d - %d)", err, channelId, operatorId)
		return nil, err
	}

	channel.Channel.DefaultBannedRights = rights
	channel.Channel.Date = date
	return channel, nil
}

func (m *ChannelCore) ToggleSlowMode(ctx context.Context, channelId, operatorId, seconds int32) (*model.MutableChannel, error) {
	var (
		date = int32(time.Now().Unix())
	)

	channel, err := m.GetMutableChannel(ctx, channelId, operatorId)
	if err != nil {
		log.Errorf("not found channel: %d", channelId)
		return nil, err
	}

	me := channel.GetImmutableChannelParticipant(operatorId)
	if me == nil || !me.IsStateOk() {
		err = mtproto.ErrChannelPrivate
		log.Errorf("toggleSlowMode{channel_id:%d, id:%d, seconds:%d} error - %v", channelId, operatorId, err)
		return nil, err
	}

	if !me.CanAdminBanUsers() {
		err = mtproto.ErrChatAdminRequired
		log.Errorf("toggleSlowMode{channel_id:%d, id:%d, seconds:%d} error - %v", channelId, operatorId, err)
		return nil, err
	}

	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		_, err = m.ChannelsDAO.UpdateSlowModeTx(tx, util.BoolToInt8(seconds != 0), seconds, channelId)
		if err != nil {
			result.Err = err
			return
		}
	})

	if tR.Err != nil {
		log.Errorf("ToggleSlowMode : (%d)", operatorId)
		return nil, tR.Err
	}

	channel.Channel.SlowmodeEnabled = seconds != 0
	channel.Channel.SlowmodeSeconds = seconds
	channel.Channel.Date = date
	return channel, nil
}

func (m *ChannelCore) DeleteMyHistory(ctx context.Context, channelId, operatorId, maxId int32) (*model.MutableChannel, error) {
	channel, err := m.GetMutableChannel(ctx, channelId, operatorId)
	if err != nil {
		log.Errorf("not found channel: %d", channelId)
		return nil, err
	}

	if !channel.Channel.IsMegagroup() {
		log.Errorf("deleteMyHistory{channel_id:%d, id:%d, max_id:%d} error - %v", channelId, operatorId, maxId, err)
		err = mtproto.ErrChannelInvalid
		return nil, err
	}

	me := channel.GetImmutableChannelParticipant(operatorId)
	if me == nil {
		log.Errorf("deleteMyHistory{channel_id:%d, id:%d, max_id:%d} error - %v", channelId, operatorId, maxId, err)
		err = mtproto.ErrUserNotParticipant
		return nil, err
	}

	_, err = m.ChannelParticipantsDAO.UpdateAvailableMinId(ctx, maxId, operatorId, channelId)
	if err != nil {
		log.Errorf("deleteMyHistory{channel_id:%d, id:%d, max_id:%d} error - %v", channelId, operatorId, maxId, err)
		return nil, err
	}

	me.AvailableMinId = maxId
	return channel, nil
}

func (m *ChannelCore) UpdatePinnedMessage(ctx context.Context, channelId, operatorId, id int32) (*model.MutableChannel, error) {
	channel, err := m.GetMutableChannel(ctx, channelId, operatorId)
	if err != nil {
		log.Errorf("not found channel: %d", channelId)
		return nil, err
	}

	date := int32(time.Now().Unix())
	me := channel.GetImmutableChannelParticipant(operatorId)
	if me == nil || !me.IsStateOk() {
		err = mtproto.ErrChannelPrivate
		log.Errorf("updatePinnedMessage{channel_id:%d, operator_id:%d, id:%d} error - %v", channelId, operatorId, id, err)
		return nil, err
	}

	if !me.CanPinMessages(date) {
		err = mtproto.ErrChatAdminRequired
		log.Errorf("updatePinnedMessage{channel_id:%d, operator_id:%d, id:%d} error - %v", channelId, operatorId, id, err)
		return nil, err
	}

	_, err = m.ChannelsDAO.Update(ctx, map[string]interface{}{
		"pinned_msg_id": id,
	}, channelId)

	if err != nil {
		log.Errorf("updatePinnedMessage{channel_id:%d, operator_id:%d, id:%d} error - %v", channelId, operatorId, id, err)
		return nil, err
	}

	channel.Channel.PinnedMsgId = id
	return channel, nil
}

func (m *ChannelCore) EditLocation(ctx context.Context, channelId, operatorId int32, geo *mtproto.InputGeoPoint, address string) (bool, error) {
	var (
		err error
	)

	channel, err := m.GetMutableChannel(ctx, channelId, operatorId)
	if err != nil {
		log.Errorf("editLocation error - %s: (%d - %d)", err, operatorId)
		return false, err
	}

	me := channel.GetImmutableChannelParticipant(operatorId)
	if me == nil || !me.IsStateOk() {
		err = mtproto.ErrChannelPrivate
		log.Errorf("editLocation error - %s: (%d)", err, operatorId)
		return false, err
	} else if me.IsChatMemberNormal() {
		err = mtproto.ErrChatAdminRequired
		log.Errorf("editLocation error - %s: (%d)", err, operatorId)
		return false, err
	}

	if !me.CanAdminChangeInfo() {
		err = mtproto.ErrRightForbidden
		log.Errorf("editLocation error - %s: (%d)", err, operatorId)
		return false, err
	}

	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		_, err = m.ChannelsDAO.UpdateGeoAddressTx(tx, util.BoolToInt8(geo != nil), geo.GetLat(), geo.GetLong(), 0, address, channelId)
		if err != nil {
			result.Err = err
			return
		}
	})

	if tR.Err != nil {
		log.Errorf("EditLocation - edit creator: (%d)", operatorId)
		return false, tR.Err
	}

	return true, nil
}

func (m *ChannelCore) SetDiscussionGroup(ctx context.Context, userId, broadcastId, groupId int32) error {
	broadcast, err := m.GetMutableChannel(ctx, broadcastId, userId)
	group, err := m.GetMutableChannel(ctx, groupId, userId)

	// 400	BOT_METHOD_INVALID	This method can't be used by a bot
	// 400	BROADCAST_ID_INVALID	Broadcast ID invalid
	// 400	LINK_NOT_MODIFIED	Discussion link not modified
	// 400	MEGAGROUP_ID_INVALID	Invalid supergroup ID
	//
	_ = broadcast
	_ = group
	_ = err

	// broadcastId
	sqlx.TxWrapper(ctx, m.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		if groupId > 0 {
			if broadcast.Channel.LinkedChatId > 0 {
				m.ChannelsDAO.ClearLinkedChatIdTx(tx, broadcast.Channel.LinkedChatId)
			}
			m.ChannelsDAO.UpdateLinkedChatIdTx(tx, groupId, 1, broadcastId)
			m.ChannelsDAO.UpdateLinkedChatIdTx(tx, broadcastId, 1, groupId)
		} else {
			m.ChannelsDAO.UpdateLinkedChatIdTx(tx, 0, 0, broadcastId)
			m.ChannelsDAO.UpdateLinkedChatIdTx(tx, 0, 0, broadcast.Channel.LinkedChatId)
		}
	})

	return nil
}

func (m *ChannelCore) GetFilterKeywords(ctx context.Context, id uint32) (keywords []string, err error) {
	return m.Dao.SelectChannelBannedKeywords(ctx, id)
}
