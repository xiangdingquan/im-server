package model

import (
	"fmt"

	"open.chat/mtproto"
)

const (
	WalletRecordType_Invalid                  = 0  //无效类型
	WalletRecordType_Recharge                 = 1  //充值
	WalletRecordType_Withdrawal               = 2  //提现
	WalletRecordType_TransferIn               = 3  //转入
	WalletRecordType_TransferOut              = 4  //转出
	WalletRecordType_CreateRedPacket          = 5  //创建红包
	WalletRecordType_GetRedPacket             = 6  //领取红包
	WalletRecordType_GivebackRedPacket        = 7  //退回红包
	WalletRecordType_Modify                   = 8  //修改金额
	WalletRecordType_WithdrawFail             = 9  //提现失败
	WalletRecordType_BlogReward               = 10 //打赏
	WalletRecordType_BlogGotReward            = 11 //领取打赏
	WalletRecordType_RemitRemittance          = 12 //转账-付款
	WalletRecordType_ReceiveRemittance        = 13 //转账-收款
	WalletRecordType_RefundRemittanceByUser   = 14 //转账-用户退款
	WalletRecordType_RefundRemittanceBySystem = 15 //转账-系统退款
)

type WalletRecordTypeUtil struct {
	WalletRecordType     int8
	OrderId              string
	UserId               int32
	Peer                 *PeerUtil
	BlogId               int32
	RedPacketId          int32
	RemittanceId         int32
	RemittanceRefundType int32
	Reason               string
}

func (p WalletRecordTypeUtil) String() (s string) {
	switch p.WalletRecordType {
	case WalletRecordType_Invalid:
		return fmt.Sprintf("WalletRecordType_Invalid: {type: %d}", p.WalletRecordType)
	case WalletRecordType_Recharge:
		return fmt.Sprintf("WalletRecordType_Recharge: {type: %d}", p.WalletRecordType)
	case WalletRecordType_Withdrawal:
		return fmt.Sprintf("WalletRecordType_Withdrawal: {type: %d}", p.WalletRecordType)
	case WalletRecordType_TransferIn:
		return fmt.Sprintf("WalletRecordType_TransferIn: {type: %d}", p.WalletRecordType)
	case WalletRecordType_TransferOut:
		return fmt.Sprintf("WalletRecordType_TransferOut: {type: %d}", p.WalletRecordType)
	case WalletRecordType_CreateRedPacket:
		return fmt.Sprintf("WalletRecordType_CreateRedPacket: {type: %d}", p.WalletRecordType)
	case WalletRecordType_GetRedPacket:
		return fmt.Sprintf("WalletRecordType_GetRedPacket: {type: %d}", p.WalletRecordType)
	case WalletRecordType_GivebackRedPacket:
		return fmt.Sprintf("WalletRecordType_GivebackRedPacket: {type: %d}", p.WalletRecordType)
	case WalletRecordType_WithdrawFail:
		return fmt.Sprintf("WalletRecordType_WithdrawFail: {type: %d}", p.WalletRecordType)
	case WalletRecordType_BlogReward:
		return fmt.Sprintf("WalletRecordType_BlogReward: {type: %d}", p.WalletRecordType)
	case WalletRecordType_BlogGotReward:
		return fmt.Sprintf("WalletRecordType_BlogGotReward: {type: %d}", p.WalletRecordType)
	case WalletRecordType_RemitRemittance:
		return fmt.Sprintf("WalletRecordType_RemitRemittance: {type: %d}", p.WalletRecordType)
	case WalletRecordType_ReceiveRemittance:
		return fmt.Sprintf("WalletRecordType_ReceiveRemittance: {type: %d}", p.WalletRecordType)
	case WalletRecordType_RefundRemittanceByUser:
		return fmt.Sprintf("WalletRecordType_RefundRemittanceByUser: {type: %d}", p.WalletRecordType)
	case WalletRecordType_RefundRemittanceBySystem:
		return fmt.Sprintf("WalletRecordType_RefundRemittanceBySystem: {type: %d}", p.WalletRecordType)
	default:
		return fmt.Sprintf("WalletRecordType_UNKNOWN: {type: %d}", p.WalletRecordType)
	}
	// return
}

func FromWalletRecordTypeUtil(recordType *mtproto.Wallet_RecordType) (v *WalletRecordTypeUtil) {
	v = &WalletRecordTypeUtil{}
	switch recordType.PredicateName {
	case mtproto.Predicate_wallet_recordTypeUnknown:
		v.WalletRecordType = WalletRecordType_Invalid
	case mtproto.Predicate_wallet_recordTypeManual:
		v.WalletRecordType = WalletRecordType_Modify
	case mtproto.Predicate_wallet_recordTypeRecharge:
		v.WalletRecordType = WalletRecordType_Recharge
		v.OrderId = recordType.GetOrderId()
	case mtproto.Predicate_wallet_recordTypeWithdrawal:
		v.WalletRecordType = WalletRecordType_Withdrawal
		v.OrderId = recordType.GetOrderId()
	case mtproto.Predicate_wallet_recordTypeWithdrawRefunded:
		v.WalletRecordType = WalletRecordType_WithdrawFail
		v.OrderId = recordType.GetOrderId()
	case mtproto.Predicate_wallet_recordTypeTransferIn:
		v.WalletRecordType = WalletRecordType_TransferIn
		v.UserId = recordType.GetUserId()
	case mtproto.Predicate_wallet_recordTypeTransferOut:
		v.WalletRecordType = WalletRecordType_TransferOut
		v.UserId = recordType.GetUserId()
	case mtproto.Predicate_wallet_recordTypeRedpacketDeduct:
		v.WalletRecordType = WalletRecordType_CreateRedPacket
		v.Peer = FromPeer(recordType.GetPeer())
		v.RedPacketId = recordType.GetRid()
	case mtproto.Predicate_wallet_recordTypeRedpacketGot:
		v.WalletRecordType = WalletRecordType_GetRedPacket
		v.Peer = FromPeer(recordType.GetPeer())
		v.RedPacketId = recordType.GetRid()
	case mtproto.Predicate_wallet_recordTypeRedpacketRefund:
		v.WalletRecordType = WalletRecordType_GivebackRedPacket
		v.Peer = FromPeer(recordType.GetPeer())
		v.RedPacketId = recordType.GetRid()
	case mtproto.Predicate_wallet_recordTypeBlogReward:
		v.WalletRecordType = WalletRecordType_BlogReward
		v.UserId = recordType.GetUserId()
		v.BlogId = recordType.GetBlogId()
	case mtproto.Predicate_wallet_recordTypeBlogGotReward:
		v.WalletRecordType = WalletRecordType_BlogGotReward
		v.UserId = recordType.GetUserId()
		v.BlogId = recordType.GetBlogId()
	case mtproto.Predicate_wallet_recordTypeRemittanceRemit:
		v.WalletRecordType = WalletRecordType_RemitRemittance
		v.RemittanceId = recordType.GetRemittanceId()
		v.UserId = recordType.GetUserId()
	case mtproto.Predicate_wallet_recordTypeRemittanceReceive:
		v.WalletRecordType = WalletRecordType_ReceiveRemittance
		v.RemittanceId = recordType.GetRemittanceId()
		v.UserId = recordType.GetUserId()
	case mtproto.Predicate_wallet_recordTypeRemittanceRefund:
		if recordType.GetType() == 1 {
			v.WalletRecordType = WalletRecordType_RefundRemittanceByUser
		} else {
			v.WalletRecordType = WalletRecordType_RefundRemittanceBySystem
		}
		v.RemittanceId = recordType.GetRemittanceId()
		v.UserId = recordType.GetUserId()
		v.RemittanceRefundType = recordType.GetType()
		v.Reason = recordType.GetReason()
	default:
		panic(fmt.Sprintf("FromWalletRecordTypeUtil(%v) error!", v))
	}
	return
}

func (v *WalletRecordTypeUtil) ToWalletRecordType() (recordType *mtproto.Wallet_RecordType) {
	switch v.WalletRecordType {
	case WalletRecordType_Invalid:
		recordType = mtproto.MakeTLWalletRecordTypeUnknown(nil).To_Wallet_RecordType()
	case WalletRecordType_Recharge:
		recordType = mtproto.MakeTLWalletRecordTypeRecharge(&mtproto.Wallet_RecordType{
			OrderId: v.OrderId,
		}).To_Wallet_RecordType()
	case WalletRecordType_Withdrawal:
		recordType = mtproto.MakeTLWalletRecordTypeWithdrawal(&mtproto.Wallet_RecordType{
			OrderId: v.OrderId,
		}).To_Wallet_RecordType()
	case WalletRecordType_TransferIn:
		recordType = mtproto.MakeTLWalletRecordTypeTransferIn(&mtproto.Wallet_RecordType{
			UserId: v.UserId,
		}).To_Wallet_RecordType()
	case WalletRecordType_TransferOut:
		recordType = mtproto.MakeTLWalletRecordTypeTransferOut(&mtproto.Wallet_RecordType{
			UserId: v.UserId,
		}).To_Wallet_RecordType()
	case WalletRecordType_CreateRedPacket:
		recordType = mtproto.MakeTLWalletRecordTypeRedpacketDeduct(&mtproto.Wallet_RecordType{
			Peer: v.Peer.ToPeer(),
			Rid:  v.RedPacketId,
		}).To_Wallet_RecordType()
	case WalletRecordType_GetRedPacket:
		recordType = mtproto.MakeTLWalletRecordTypeRedpacketGot(&mtproto.Wallet_RecordType{
			Peer: v.Peer.ToPeer(),
			Rid:  v.RedPacketId,
		}).To_Wallet_RecordType()
	case WalletRecordType_GivebackRedPacket:
		recordType = mtproto.MakeTLWalletRecordTypeRedpacketRefund(&mtproto.Wallet_RecordType{
			Peer: v.Peer.ToPeer(),
			Rid:  v.RedPacketId,
		}).To_Wallet_RecordType()
	case WalletRecordType_WithdrawFail:
		recordType = mtproto.MakeTLWalletRecordTypeWithdrawRefunded(&mtproto.Wallet_RecordType{
			OrderId: v.OrderId,
		}).To_Wallet_RecordType()
	case WalletRecordType_BlogReward:
		recordType = mtproto.MakeTLWalletRecordTypeBlogReward(&mtproto.Wallet_RecordType{
			UserId: v.UserId,
			BlogId: v.BlogId,
		}).To_Wallet_RecordType()
	case WalletRecordType_BlogGotReward:
		recordType = mtproto.MakeTLWalletRecordTypeBlogGotReward(&mtproto.Wallet_RecordType{
			UserId: v.UserId,
			BlogId: v.BlogId,
		}).To_Wallet_RecordType()
	case WalletRecordType_RemitRemittance:
		recordType = mtproto.MakeTLWalletRecordTypeRemittanceRemit(&mtproto.Wallet_RecordType{
			UserId:       v.UserId,
			RemittanceId: v.RemittanceId,
		}).To_Wallet_RecordType()
	case WalletRecordType_ReceiveRemittance:
		recordType = mtproto.MakeTLWalletRecordTypeRemittanceReceive(&mtproto.Wallet_RecordType{
			UserId:       v.UserId,
			RemittanceId: v.RemittanceId,
		}).To_Wallet_RecordType()
	case WalletRecordType_RefundRemittanceByUser:
		recordType = mtproto.MakeTLWalletRecordTypeRemittanceRefund(&mtproto.Wallet_RecordType{
			UserId:       v.UserId,
			RemittanceId: v.RemittanceId,
			Type:         v.RemittanceRefundType,
			Reason:       v.Reason,
		}).To_Wallet_RecordType()
	case WalletRecordType_RefundRemittanceBySystem:
		recordType = mtproto.MakeTLWalletRecordTypeRemittanceRefund(&mtproto.Wallet_RecordType{
			UserId:       v.UserId,
			RemittanceId: v.RemittanceId,
			Type:         v.RemittanceRefundType,
			Reason:       v.Reason,
		}).To_Wallet_RecordType()
	default:
		panic(fmt.Sprintf("ToWalletRecordType(%v) error!", v))
	}
	return
}

func MakeWalletRecordTypeFriend() (v *VisibleTypeUtil) {
	v = &VisibleTypeUtil{
		VisibleType: VisibleType_Friend,
	}
	return
}

func WalletRecordTypeToRemark(recordType int8) LocalizationWords {
	switch recordType {
	case WalletRecordType_Invalid:
		return LocalizationWords{
			LocalizationCN:      "无效",
			LocalizationEN:      "invalid",
			LocalizationDefault: "invalid",
		}
	case WalletRecordType_Recharge:
		return LocalizationWords{
			LocalizationCN:      "充值",
			LocalizationEN:      "recharge",
			LocalizationDefault: "recharge",
		}
	case WalletRecordType_Withdrawal:
		return LocalizationWords{
			LocalizationCN:      "提现",
			LocalizationEN:      "withdrawal",
			LocalizationDefault: "withdrawal",
		}
	case WalletRecordType_TransferIn:
		return LocalizationWords{
			LocalizationCN:      "转入",
			LocalizationEN:      "transfer In",
			LocalizationDefault: "transfer In",
		}
	case WalletRecordType_TransferOut:
		return LocalizationWords{
			LocalizationCN:      "转出",
			LocalizationEN:      "transfer out",
			LocalizationDefault: "transfer out",
		}
	case WalletRecordType_CreateRedPacket:
		return LocalizationWords{
			LocalizationCN:      "发红包",
			LocalizationEN:      "create red packet",
			LocalizationDefault: "create red packet",
		}
	case WalletRecordType_GetRedPacket:
		return LocalizationWords{
			LocalizationCN:      "抢红包",
			LocalizationEN:      "get red packet",
			LocalizationDefault: "get red packet",
		}
	case WalletRecordType_GivebackRedPacket:
		return LocalizationWords{
			LocalizationCN:      "退红包",
			LocalizationEN:      "giveback red packet",
			LocalizationDefault: "giveback red packet",
		}
	case WalletRecordType_WithdrawFail:
		return LocalizationWords{
			LocalizationCN:      "提现失败",
			LocalizationEN:      "withdraw fail",
			LocalizationDefault: "withdraw fail",
		}
	case WalletRecordType_BlogReward:
		return LocalizationWords{
			LocalizationCN:      "朋友圈打赏",
			LocalizationEN:      "blog reward",
			LocalizationDefault: "blog reward",
		}
	case WalletRecordType_BlogGotReward:
		return LocalizationWords{
			LocalizationCN:      "朋友圈领取打赏",
			LocalizationEN:      "blog got reward",
			LocalizationDefault: "blog got reward",
		}
	case WalletRecordType_RemitRemittance:
		return LocalizationWords{
			LocalizationCN:      "转账",
			LocalizationEN:      "remit",
			LocalizationDefault: "remit",
		}
	case WalletRecordType_ReceiveRemittance:
		return LocalizationWords{
			LocalizationCN:      "收款",
			LocalizationEN:      "receive",
			LocalizationDefault: "receive",
		}
	case WalletRecordType_RefundRemittanceByUser:
		return LocalizationWords{
			LocalizationCN:      "用户退款",
			LocalizationEN:      "refund by user",
			LocalizationDefault: "refund by user",
		}
	case WalletRecordType_RefundRemittanceBySystem:
		return LocalizationWords{
			LocalizationCN:      "系统退款",
			LocalizationEN:      "refund by system",
			LocalizationDefault: "refund by system",
		}
	default:
		return LocalizationWords{
			LocalizationCN:      "未知",
			LocalizationEN:      "unknown",
			LocalizationDefault: "unknown",
		}
	}
}
