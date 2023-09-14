package botapi

import (
	"encoding/json"
)

type StickerSet struct {
	Name          string     `json:"name"`
	Title         string     `json:"title"`
	IsAnimated    bool       `json:"is_animated"`
	ContainsMasks bool       `json:"contains_masks"`
	Stickers      []*Sticker `json:"stickers"`
	Thumb         *PhotoSize `json:"thumb,omitempty"`
}

func (m *StickerSet) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *StickerSet) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

func (m *MaskPosition) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *MaskPosition) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}

type Sticker struct {
	FileId       string        `json:"file_id"`
	Width        int32         `json:"width"`
	Height       int32         `json:"height"`
	IsAnimated   bool          `json:"is_animated"`
	Thumb        *PhotoSize    `json:"thumb,omitempty"`
	Emoji        string        `json:"emoji,omitempty"`
	SetName      string        `json:"set_name,omitempty"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	FileSize     int32         `json:"file_size,omitempty"`
}

func (m *Sticker) Decode(b []byte) (err error) {
	err = json.Unmarshal(b, m)
	return
}

func (m *Sticker) Encode() (b []byte, err error) {
	b, err = json.Marshal(m)
	return
}
