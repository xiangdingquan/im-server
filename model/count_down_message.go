package model

type CountDownMessage struct {
	UId      int32 `json:"uId"`
	PeerType int32 `json:"peerType"`
	PeerId   int32 `json:"peerId"`
	MsgId    int32 `json:"msgId"`
}
