package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"regexp"
	"strings"

	"open.chat/pkg/util"
)

type (
	tJsonMessage struct {
		Mid      uint32
		Sid      uint32
		Msg      string
		JsonText string
	}
)

func ParseJsonMessage(message string) (*tJsonMessage, error) {
	str := `^ct!([a-f\dA-F]{32})!(\d+)!(\d+)!([^$]+)`
	r := regexp.MustCompile(str)
	matchs := r.FindStringSubmatch(message)
	if len(matchs) != 5 {
		return nil, errors.New("parse fail")
	}
	jm := &tJsonMessage{Msg: message}
	jm.JsonText = matchs[4]
	if jm.JsonText == "" {
		return nil, errors.New("json fail")
	}
	vv := md5.Sum([]byte(jm.JsonText))
	sMd5 := hex.EncodeToString(vv[:])
	if sMd5 != strings.ToLower(matchs[1]) {
		return nil, errors.New("check fail")
	}
	jm.Mid, _ = util.StringToUint32(matchs[2])
	jm.Sid, _ = util.StringToUint32(matchs[3])
	return jm, nil
}

// ToString .
func MakeJsonMessage(mid uint32, sid uint32, v interface{}) (*tJsonMessage, error) {
	//Data map[string]interface{} `json:"data"`
	jm := &tJsonMessage{Msg: ""}
	if v == nil {
		v = struct{}{}
	}
	cbJson, err := json.Marshal(v)
	if err != nil {
		return nil, errors.New("make message fail")
	}
	vv := md5.Sum(cbJson)
	jm.JsonText = string(cbJson)
	sMd5 := hex.EncodeToString(vv[:])
	jm.Msg = "ct!" + sMd5
	jm.Msg += "!" + util.Int64ToString(int64(mid))
	jm.Msg += "!" + util.Int64ToString(int64(sid))
	jm.Msg += "!" + jm.JsonText
	return jm, nil
}

// ParsingMessage .
func (m *tJsonMessage) Parse(v interface{}) bool {
	if m == nil {
		return false
	}
	s, ok := v.(*string)
	if ok {
		*s = m.JsonText
	} else {
		err := json.Unmarshal([]byte(m.JsonText), v)
		if err != nil {
			return false
		}
	}
	return true
}
