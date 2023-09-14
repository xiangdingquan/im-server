package tpns

import "encoding/json"

// Message represents jpush request message
type (
	Payload struct {
		//推送目标：
		//all：全量推送
		//tag：标签推送
		//token：单设备推送
		//token_list：设备列表推送
		//account：单账号推送
		//account_list：账号列表推送
		//package_account_push：号码包推送
		//package_token_push：token 文件包推送
		AudienceType string   `json:"audience_type"`
		TokenList    []string `json:"token_list"`

		//消息类型：
		//notify：通知
		//message：透传消息/静默消息
		MessageType string  `json:"message_type"`
		Message     Message `json:"message,omitempty"`
	}

	Message struct {
		//消息标题
		Title string `json:"title"`
		//消息内容
		Content string `json:"content"`
	}

	// Response represents fcm response message - (tokens and topics)
	Response struct {
		StatusCode int
		//Seq        int    `json:"seq,omitempty"`
		PushId string `json:"push_id,omitempty"`
		//InvalidTargeList []string `json:"invalid_targe_list,omitempty"`
		RetCode int `json:"ret_code"`
		//Environment      string   `json:"environment,omitempty"`
		ErrMsg string `json:"err_msg,omitempty"`
		Result string `json:"result,omitempty"`
	}
)

func (p *Payload) Marshal() []byte {
	payload, _ := json.Marshal(p)
	return payload
}
