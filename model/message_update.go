package model

import (
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"open.chat/mtproto"
)

func MessageUpdate(message *mtproto.Message) (msg *mtproto.Message) {
	if message != nil {
		msg = proto.Clone(message).(*mtproto.Message)
		peer := msg.GetFromId_FLAGPEER()
		if peer != nil && peer.PredicateName == mtproto.Predicate_peerUser {
			msg.FromId_FLAGINT32 = &types.Int32Value{Value: peer.GetUserId()}
		}
		if msg.Views != nil && msg.Forwards == nil {
			msg.Forwards = &types.Int32Value{Value: msg.Views.GetValue()}
		}
		//peerId := msg.GetPeerId()
		//if peerId != nil {
		//	msg.ToId = proto.Clone(peerId).(*mtproto.Peer)
		//}
		reply := msg.GetReplyTo()
		if reply != nil {
			msg.ReplyToMsgId = &types.Int32Value{Value: reply.GetReplyToMsgId()}
		}
		fwdFrom := msg.GetFwdFrom()
		if fwdFrom != nil {
			peer = fwdFrom.GetFromId_FLAGPEER()
			if peer != nil && peer.PredicateName == mtproto.Predicate_peerUser {
				msg.FwdFrom.FromId_FLAGINT32 = &types.Int32Value{Value: peer.GetUserId()}
			}
		}

		//if msg.GetPeerId() != nil && msg.GetPeerId().PredicateName == mtproto.Predicate_peerUser {
		//	if msg.GetFromId_FLAGPEER() != nil && msg.GetFromId_FLAGPEER().PredicateName == mtproto.Predicate_peerUser {
		//		msg.GetPeerId().UserId = msg.GetFromId_FLAGPEER().UserId
		//	}
		//}
	}
	return
}
