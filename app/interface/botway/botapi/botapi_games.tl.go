package botapi

import (
	"encoding/json"
)

type CallbackGame struct {
}

func (m *CallbackGame) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *CallbackGame) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type GameHighScore struct {
	Position int32 `json:"position"`
	User     *User `json:"user"`
	Score    int32 `json:"score"`
}

func (m *GameHighScore) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *GameHighScore) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type Game struct {
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	Photo        []*PhotoSize     `json:"photo,omitempty"`
	Text         string           `json:"text,omitempty"`
	TextEntities []*MessageEntity `json:"text_entities,omitempty"`
	Animation    *Animation       `json:"animation,omitempty"`
}

func (m *Game) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *Game) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}
