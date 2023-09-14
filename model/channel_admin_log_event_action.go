package model

import "open.chat/mtproto"

type EventActionType int32

const (
	EventActionUnknown                EventActionType = 0
	EventActionChangeTitle            EventActionType = 1 << 0
	gEventActionChangeAbout           EventActionType = 1 << 1
	EventActionChangeUsername         EventActionType = 1 << 2
	EventActionChangePhoto            EventActionType = 1 << 3
	EventActionToggleInvites          EventActionType = 1 << 4
	EventActionToggleSignatures       EventActionType = 1 << 5
	EventActionUpdatePinned           EventActionType = 1 << 6
	EventActionEditMessage            EventActionType = 1 << 7
	EventActionDeleteMessage          EventActionType = 1 << 8
	EventActionParticipantJoin        EventActionType = 1 << 9
	EventActionParticipantLeave       EventActionType = 1 << 10
	EventActionParticipantInvite      EventActionType = 1 << 11
	EventActionParticipantToggleBan   EventActionType = 1 << 12
	EventActionParticipantToggleAdmin EventActionType = 1 << 13
	EventActionChangeStickerSet       EventActionType = 1 << 14
	EventActionTogglePreHistoryHidden EventActionType = 1 << 15
	EventActionDefaultBannedRights    EventActionType = 1 << 16
	EventActionStopPoll               EventActionType = 1 << 17
	EventActionChangeLinkedChat       EventActionType = 1 << 18
	EventActionChangeLocation         EventActionType = 1 << 19
	EventActionToggleSlowMode         EventActionType = 1 << 20
)

func FromChannelAdminLogEventAction(action *mtproto.ChannelAdminLogEventAction) (evt EventActionType) {
	switch action.GetPredicateName() {
	case mtproto.Predicate_channelAdminLogEventActionChangeTitle:
	case mtproto.Predicate_channelAdminLogEventActionChangeAbout:
	case mtproto.Predicate_channelAdminLogEventActionChangeUsername:
	case mtproto.Predicate_channelAdminLogEventActionChangePhoto:
	case mtproto.Predicate_channelAdminLogEventActionToggleInvites:
	case mtproto.Predicate_channelAdminLogEventActionToggleSignatures:
	case mtproto.Predicate_channelAdminLogEventActionUpdatePinned:
	case mtproto.Predicate_channelAdminLogEventActionEditMessage:
	case mtproto.Predicate_channelAdminLogEventActionDeleteMessage:
	case mtproto.Predicate_channelAdminLogEventActionParticipantJoin:
	case mtproto.Predicate_channelAdminLogEventActionParticipantLeave:
	case mtproto.Predicate_channelAdminLogEventActionParticipantInvite:
	case mtproto.Predicate_channelAdminLogEventActionParticipantToggleBan:
	case mtproto.Predicate_channelAdminLogEventActionParticipantToggleAdmin:
	case mtproto.Predicate_channelAdminLogEventActionChangeStickerSet:
	case mtproto.Predicate_channelAdminLogEventActionTogglePreHistoryHidden:
	case mtproto.Predicate_channelAdminLogEventActionDefaultBannedRights:
	case mtproto.Predicate_channelAdminLogEventActionStopPoll:
	case mtproto.Predicate_channelAdminLogEventActionChangeLinkedChat:
	case mtproto.Predicate_channelAdminLogEventActionChangeLocation:
	case mtproto.Predicate_channelAdminLogEventActionToggleSlowMode:
	default:
		evt = EventActionUnknown
	}
	return
}

func FromChannelAdminLogEventsFilter(filter *mtproto.ChannelAdminLogEventsFilter) (evt EventActionType) {
	evt = 0

	if filter.GetJoin() {
	}
	if filter.GetLeave() {
	}
	if filter.GetInvite() {
	}
	if filter.GetBan() {
	}
	if filter.GetUnban() {
	}
	if filter.GetKick() {
	}
	if filter.GetUnkick() {
	}
	if filter.GetPromote() {
	}
	if filter.GetDemote() {
	}
	if filter.GetInfo() {
	}
	if filter.GetSettings() {
	}
	if filter.GetPinned() {
	}
	if filter.GetEdit() {
	}
	if filter.GetDelete() {
	}

	return
}
