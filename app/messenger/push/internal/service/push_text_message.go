package service

import (
	"open.chat/model"
)

func (s *Service) handleNotificationMessage(mid uint32, sid uint32, sMessage string) string {
	if mid == 2 && sid == 100 {
		return "【收到一个红包】"
	} else if mid == 2 && sid == 200 {
		return "【领取了一个红包】"
	} else if mid == 3 && sid == 100 {
		return "【截屏了一次】"
	} else if mid == 3 && sid == 101 {
		return "【对方加你为好友了】"
	} else if mid == 5 && sid == 100 {
		return "【收到一笔转账】"
	} else if mid == 5 && sid == 200 {
		return "【领取了一笔转账】"
	} else if mid == 6 && sid == 100 {
		return "【有成员发送敏感词被踢出该群聊】"
	} else if mid == 7 { //oss消息
		switch sid {
		case 101:
			return "【图片】"
		case 102:
			return "【视频】"
		case 103:
			return "【动画】"
		case 104:
			return "【语音】"
		case 105, 106, 107:
			return "【文件】"
		}
	}
	return ""
}

func (s *Service) customNotification(sMsg string) string {
	msg, _ := model.ParseJsonMessage(sMsg)
	if msg != nil {
		sMsg = s.handleNotificationMessage(msg.Mid, msg.Sid, msg.JsonText)
	}
	return sMsg
}
