package consts

type SmsCodeType uint8

const (
	SmsCodeType_Invalid           SmsCodeType = iota //无效类型
	SmsCodeType_SetWalletPassword                    //修改钱包密码
	SmsCodeType_RegistAccount                        //注册账号
	SmsCodeType_DeleteAccount                        //注销账号
)
