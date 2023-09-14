package model

import (
	"strings"
)

const (
	MinUsernameLen = 5
	MaxUsernameLen = 32
)

const (
	UsernameNotExisted   = 0
	UsernameExisted      = 1
	UsernameExistedNotMe = 2
	UsernameExistedIsMe  = 3
)

func CheckUsernameInvalid(username string) bool {
	if len(username) < MinUsernameLen || len(username) > MaxUsernameLen {
		return false
	}

	if strings.HasPrefix(username, "_") || strings.HasSuffix(username, "_") {
		return false
	}

	if username[0] >= '0' && username[0] <= '9' {
		return false
	}

	for _, ch := range username {
		if !(ch >= '0' && ch <= '9' || ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_') {
			return false
		}
	}
	return true
}
