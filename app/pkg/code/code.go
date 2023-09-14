package code

type SmsVerifyCodeConfig struct {
	Name          string
	SendCodeUrl   string
	VerifyCodeUrl string
	Signed        string
	UserName      string
	Password      string
	Key           string
	Secret        string
}
