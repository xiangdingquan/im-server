package model

const (
	LangPackAndroid  = "android"
	LangPackiOS      = "ios"
	LangPackmacOS    = "macos"
	LangPackTDesktop = "tdesktop"
	LangPackAndroidX = "android_x"
	LangPackEmojie   = "emojie"
)

func CheckLangPackInvalid(langPack string) (r bool) {
	r = true
	switch langPack {
	case LangPackAndroid:
	case LangPackiOS:
	case LangPackmacOS:
	case LangPackTDesktop:
	case LangPackAndroidX:
	default:
		r = false
	}
	return
}
