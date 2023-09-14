package model

import "open.chat/mtproto"

// //////////////////////////////////////////////////////////////////////
type DialogPinnedExt struct {
	Order int64
	PeerUtil
}

type DialogPinnedExtList []DialogPinnedExt

func (m DialogPinnedExtList) Add(peerType, peerId int32, order int64) DialogPinnedExtList {
	return append(m, DialogPinnedExt{
		Order: order,
		PeerUtil: PeerUtil{
			PeerType: peerType,
			PeerId:   peerId,
		},
	})
}

// sort
func (m DialogPinnedExtList) Len() int {
	return len(m)
}
func (m DialogPinnedExtList) Swap(i, j int) {
	m[j], m[i] = m[i], m[j]
}
func (m DialogPinnedExtList) Less(i, j int) bool {
	return m[i].Order < m[j].Order
}

// //////////////////////////////////////////////////////////////////////
type DialogExt struct {
	Order int64
	*mtproto.Dialog
	AvailableMinId int32 // channel
	Date           int32
}

type DialogExtList []*DialogExt

// sort
func (m DialogExtList) Len() int {
	return len(m)
}
func (m DialogExtList) Swap(i, j int) {
	m[j], m[i] = m[i], m[j]
}
func (m DialogExtList) Less(i, j int) bool {
	return m[i].Order < m[j].Order
}

func (m DialogExtList) ToDialogList() []*mtproto.Dialog {
	dialogs := make([]*mtproto.Dialog, 0, len(m))
	for _, d := range m {
		dialogs = append(dialogs, d.Dialog)
	}
	return dialogs
}
