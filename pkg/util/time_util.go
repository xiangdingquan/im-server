package util

import (
	"time"
)

// Duration be used toml unmarshal string time, like 1s, 500ms.
type Duration time.Duration

func (d *Duration) UnmarshalText(text []byte) error {
	tmp, err := time.ParseDuration(string(text))
	if err == nil {
		*d = Duration(tmp)
	}
	return err
}

// ///////////////////////////////////////////////////////////
func NowFormatYMDHMS() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
