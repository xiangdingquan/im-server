package model

import (
	"open.chat/mtproto"
)

type Theme struct {
	*mtproto.Theme
	DocumentId int64
}

func MakeTheme(theme *mtproto.Theme, id int64) *Theme {
	return &Theme{
		Theme:      theme,
		DocumentId: id,
	}
}

func (m *Theme) GetTheme() *mtproto.Theme {
	return m.Theme
}

func (m *Theme) SetDocument(v *mtproto.Document) {
	m.Theme.Document = v
}
