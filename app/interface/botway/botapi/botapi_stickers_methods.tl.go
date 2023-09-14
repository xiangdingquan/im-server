package botapi

import (
	"strconv"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"

	"open.chat/mtproto"
)

type SendStickerReq struct {
	ChatId              string `json:"chat_id" form:"chat_id"`
	Sticker             string `json:"sticker" form:"sticker"`
	DisableNotification string `json:"disable_notification,omitempty" form:"disable_notification"`
	ReplyToMessageId    string `json:"reply_to_message_id,omitempty" form:"reply_to_message_id"`
	ReplyMarkup         string `json:"reply_markup,omitempty" form:"reply_markup"`
}

func (m *SendStickerReq) Method() string {
	return "sendSticker"
}

type SendSticker2 struct {
	ChatId              ChatId2               `json:"chat_id,omitempty"`
	Sticker             InputFile             `json:"sticker,omitempty"`
	DisableNotification bool                  `json:"disable_notification,omitempty"`
	ReplyToMessageId    int32                 `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (m *SendSticker2) NewRequest() BotApiRequest {
	return new(SendStickerReq)
}

func (m *SendSticker2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SendStickerReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}
	if len(req.Sticker) > 0 {
		m.Sticker = MakeInputFile(req.Sticker)
	}
	if len(req.DisableNotification) > 0 {
		if m.DisableNotification, err = strconv.ParseBool(req.DisableNotification); err != nil {
			return
		}
	}
	if len(req.ReplyToMessageId) > 0 {
		if i, err2 := strconv.ParseInt(req.ReplyToMessageId, 10, 32); err2 != nil {
			return err2
		} else {
			m.ReplyToMessageId = int32(i)
		}
	}
	if len(req.ReplyMarkup) > 0 {
		m.ReplyMarkup = new(InlineKeyboardMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

func (m *SendSticker2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SendStickerReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.ChatId) > 0 {
		m.ChatId = MakeChatId(req.ChatId)
	}

	if m.Sticker, err = MakeInputFile2(c, contentType, "sticker", req.Sticker); err != nil {
		return
	}

	if len(req.DisableNotification) > 0 {
		if m.DisableNotification, err = strconv.ParseBool(req.DisableNotification); err != nil {
			return
		}
	}
	if len(req.ReplyToMessageId) > 0 {
		if i, err2 := strconv.ParseInt(req.ReplyToMessageId, 10, 32); err2 != nil {
			return err2
		} else {
			m.ReplyToMessageId = int32(i)
		}
	}
	if len(req.ReplyMarkup) > 0 {
		m.ReplyMarkup = new(InlineKeyboardMarkup)
		if err = m.ReplyMarkup.Decode([]byte(req.ReplyMarkup)); err != nil {
			return
		}
	}

	return
}

type GetStickerSetReq struct {
	Name string `json:"name" form:"name"`
}

func (m *GetStickerSetReq) Method() string {
	return "getStickerSet"
}

type GetStickerSet2 struct {
	Name string `json:"name,omitempty"`
}

func (m *GetStickerSet2) NewRequest() BotApiRequest {
	return new(GetStickerSetReq)
}

func (m *GetStickerSet2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*GetStickerSetReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.Name) > 0 {
		m.Name = req.Name
	}

	return
}

func (m *GetStickerSet2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*GetStickerSetReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.Name) > 0 {
		m.Name = req.Name
	}

	return
}

type UploadStickerFileReq struct {
	UserId     string `json:"user_id" form:"user_id"`
	PngSticker string `json:"png_sticker" form:"png_sticker"`
}

func (m *UploadStickerFileReq) Method() string {
	return "uploadStickerFile"
}

type UploadStickerFile2 struct {
	UserId     int32     `json:"user_id,omitempty"`
	PngSticker InputFile `json:"png_sticker,omitempty"`
}

func (m *UploadStickerFile2) NewRequest() BotApiRequest {
	return new(UploadStickerFileReq)
}

func (m *UploadStickerFile2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*UploadStickerFileReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.UserId) > 0 {
		if i, err2 := strconv.ParseInt(req.UserId, 10, 32); err2 != nil {
			return err2
		} else {
			m.UserId = int32(i)
		}
	}
	if len(req.PngSticker) > 0 {
		m.PngSticker = MakeInputFile(req.PngSticker)
	}

	return
}

func (m *UploadStickerFile2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*UploadStickerFileReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.UserId) > 0 {
		if i, err2 := strconv.ParseInt(req.UserId, 10, 32); err2 != nil {
			return err2
		} else {
			m.UserId = int32(i)
		}
	}
	if m.PngSticker, err = MakeInputFile2(c, contentType, "png_sticker", req.PngSticker); err != nil {
		return
	}

	return
}

type CreateNewStickerSetReq struct {
	UserId        string `json:"user_id" form:"user_id"`
	Name          string `json:"name" form:"name"`
	Title         string `json:"title" form:"title"`
	PngSticker    string `json:"png_sticker,omitempty" form:"png_sticker"`
	TgsSticker    string `json:"tgs_sticker,omitempty" form:"tgs_sticker"`
	Emojis        string `json:"emojis" form:"emojis"`
	ContainsMasks string `json:"contains_masks,omitempty" form:"contains_masks"`
	MaskPosition  string `json:"mask_position,omitempty" form:"mask_position"`
}

func (m *CreateNewStickerSetReq) Method() string {
	return "createNewStickerSet"
}

type CreateNewStickerSet2 struct {
	UserId        int32         `json:"user_id,omitempty"`
	Name          string        `json:"name,omitempty"`
	Title         string        `json:"title,omitempty"`
	PngSticker    InputFile     `json:"png_sticker,omitempty"`
	TgsSticker    InputFile     `json:"tgs_sticker,omitempty"`
	Emojis        string        `json:"emojis,omitempty"`
	ContainsMasks bool          `json:"contains_masks,omitempty"`
	MaskPosition  *MaskPosition `json:"mask_position,omitempty"`
}

func (m *CreateNewStickerSet2) NewRequest() BotApiRequest {
	return new(CreateNewStickerSetReq)
}

func (m *CreateNewStickerSet2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*CreateNewStickerSetReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.UserId) > 0 {
		if i, err2 := strconv.ParseInt(req.UserId, 10, 32); err2 != nil {
			return err2
		} else {
			m.UserId = int32(i)
		}
	}
	if len(req.Name) > 0 {
		m.Name = req.Name
	}
	if len(req.Title) > 0 {
		m.Title = req.Title
	}
	if len(req.PngSticker) > 0 {
		m.PngSticker = MakeInputFile(req.PngSticker)
	}
	if len(req.TgsSticker) > 0 {
		m.TgsSticker = MakeInputFile(req.TgsSticker)
	}
	if len(req.Emojis) > 0 {
		m.Emojis = req.Emojis
	}
	if len(req.ContainsMasks) > 0 {
		if m.ContainsMasks, err = strconv.ParseBool(req.ContainsMasks); err != nil {
			return
		}
	}
	if len(req.MaskPosition) > 0 {
		m.MaskPosition = new(MaskPosition)
		if err = m.MaskPosition.Decode([]byte(req.MaskPosition)); err != nil {
			return
		}
	}

	return
}

func (m *CreateNewStickerSet2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*CreateNewStickerSetReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.UserId) > 0 {
		if i, err2 := strconv.ParseInt(req.UserId, 10, 32); err2 != nil {
			return err2
		} else {
			m.UserId = int32(i)
		}
	}
	if len(req.Name) > 0 {
		m.Name = req.Name
	}
	if len(req.Title) > 0 {
		m.Title = req.Title
	}
	if m.PngSticker, err = MakeInputFile2(c, contentType, "png_sticker", req.PngSticker); err != nil {
		return
	}

	if m.TgsSticker, err = MakeInputFile2(c, contentType, "tgs_sticker", req.TgsSticker); err != nil {
		return
	}

	if len(req.Emojis) > 0 {
		m.Emojis = req.Emojis
	}
	if len(req.ContainsMasks) > 0 {
		if m.ContainsMasks, err = strconv.ParseBool(req.ContainsMasks); err != nil {
			return
		}
	}
	if len(req.MaskPosition) > 0 {
		m.MaskPosition = new(MaskPosition)
		if err = m.MaskPosition.Decode([]byte(req.MaskPosition)); err != nil {
			return
		}
	}

	return
}

type AddStickerToSetReq struct {
	UserId       string `json:"user_id" form:"user_id"`
	Name         string `json:"name" form:"name"`
	PngSticker   string `json:"png_sticker" form:"png_sticker"`
	TgsSticker   string `json:"tgs_sticker,omitempty" form:"tgs_sticker"`
	Emojis       string `json:"emojis" form:"emojis"`
	MaskPosition string `json:"mask_position,omitempty" form:"mask_position"`
}

func (m *AddStickerToSetReq) Method() string {
	return "addStickerToSet"
}

type AddStickerToSet2 struct {
	UserId       int32         `json:"user_id,omitempty"`
	Name         string        `json:"name,omitempty"`
	PngSticker   InputFile     `json:"png_sticker,omitempty"`
	TgsSticker   InputFile     `json:"tgs_sticker,omitempty"`
	Emojis       string        `json:"emojis,omitempty"`
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
}

func (m *AddStickerToSet2) NewRequest() BotApiRequest {
	return new(AddStickerToSetReq)
}

func (m *AddStickerToSet2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*AddStickerToSetReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.UserId) > 0 {
		if i, err2 := strconv.ParseInt(req.UserId, 10, 32); err2 != nil {
			return err2
		} else {
			m.UserId = int32(i)
		}
	}
	if len(req.Name) > 0 {
		m.Name = req.Name
	}
	if len(req.PngSticker) > 0 {
		m.PngSticker = MakeInputFile(req.PngSticker)
	}
	if len(req.TgsSticker) > 0 {
		m.TgsSticker = MakeInputFile(req.TgsSticker)
	}
	if len(req.Emojis) > 0 {
		m.Emojis = req.Emojis
	}
	if len(req.MaskPosition) > 0 {
		m.MaskPosition = new(MaskPosition)
		if err = m.MaskPosition.Decode([]byte(req.MaskPosition)); err != nil {
			return
		}
	}

	return
}

func (m *AddStickerToSet2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*AddStickerToSetReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.UserId) > 0 {
		if i, err2 := strconv.ParseInt(req.UserId, 10, 32); err2 != nil {
			return err2
		} else {
			m.UserId = int32(i)
		}
	}
	if len(req.Name) > 0 {
		m.Name = req.Name
	}
	if m.PngSticker, err = MakeInputFile2(c, contentType, "png_sticker", req.PngSticker); err != nil {
		return
	}

	if m.TgsSticker, err = MakeInputFile2(c, contentType, "tgs_sticker", req.TgsSticker); err != nil {
		return
	}

	if len(req.Emojis) > 0 {
		m.Emojis = req.Emojis
	}
	if len(req.MaskPosition) > 0 {
		m.MaskPosition = new(MaskPosition)
		if err = m.MaskPosition.Decode([]byte(req.MaskPosition)); err != nil {
			return
		}
	}

	return
}

type SetStickerPositionInSetReq struct {
	Sticker  string `json:"sticker" form:"sticker"`
	Position string `json:"position" form:"position"`
}

func (m *SetStickerPositionInSetReq) Method() string {
	return "setStickerPositionInSet"
}

type SetStickerPositionInSet2 struct {
	Sticker  string `json:"sticker,omitempty"`
	Position int32  `json:"position,omitempty"`
}

func (m *SetStickerPositionInSet2) NewRequest() BotApiRequest {
	return new(SetStickerPositionInSetReq)
}

func (m *SetStickerPositionInSet2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SetStickerPositionInSetReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.Sticker) > 0 {
		m.Sticker = req.Sticker
	}
	if len(req.Position) > 0 {
		if i, err2 := strconv.ParseInt(req.Position, 10, 32); err2 != nil {
			return err2
		} else {
			m.Position = int32(i)
		}
	}

	return
}

func (m *SetStickerPositionInSet2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SetStickerPositionInSetReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.Sticker) > 0 {
		m.Sticker = req.Sticker
	}
	if len(req.Position) > 0 {
		if i, err2 := strconv.ParseInt(req.Position, 10, 32); err2 != nil {
			return err2
		} else {
			m.Position = int32(i)
		}
	}

	return
}

type DeleteStickerFromSetReq struct {
	Sticker string `json:"sticker" form:"sticker"`
}

func (m *DeleteStickerFromSetReq) Method() string {
	return "deleteStickerFromSet"
}

type DeleteStickerFromSet2 struct {
	Sticker string `json:"sticker,omitempty"`
}

func (m *DeleteStickerFromSet2) NewRequest() BotApiRequest {
	return new(DeleteStickerFromSetReq)
}

func (m *DeleteStickerFromSet2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*DeleteStickerFromSetReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.Sticker) > 0 {
		m.Sticker = req.Sticker
	}

	return
}

func (m *DeleteStickerFromSet2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*DeleteStickerFromSetReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.Sticker) > 0 {
		m.Sticker = req.Sticker
	}

	return
}

type SetStickerSetThumbReq struct {
	Name   string `json:"name" form:"name"`
	UserId string `json:"user_id" form:"user_id"`
	Thumb  string `json:"thumb,omitempty" form:"thumb"`
}

func (m *SetStickerSetThumbReq) Method() string {
	return "setStickerSetThumb"
}

type SetStickerSetThumb2 struct {
	Name   string    `json:"name,omitempty"`
	UserId int32     `json:"user_id,omitempty"`
	Thumb  InputFile `json:"thumb,omitempty"`
}

func (m *SetStickerSetThumb2) NewRequest() BotApiRequest {
	return new(SetStickerSetThumbReq)
}

func (m *SetStickerSetThumb2) Decode(r BotApiRequest) (err error) {
	req, ok := r.(*SetStickerSetThumbReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.Name) > 0 {
		m.Name = req.Name
	}
	if len(req.UserId) > 0 {
		if i, err2 := strconv.ParseInt(req.UserId, 10, 32); err2 != nil {
			return err2
		} else {
			m.UserId = int32(i)
		}
	}
	if len(req.Thumb) > 0 {
		m.Thumb = MakeInputFile(req.Thumb)
	}

	return
}

func (m *SetStickerSetThumb2) Decode2(c *bm.Context, contentType string, r BotApiRequest) (err error) {
	req, ok := r.(*SetStickerSetThumbReq)
	if !ok || req == nil {
		err = mtproto.ErrBadRequest
		return
	}

	if len(req.Name) > 0 {
		m.Name = req.Name
	}
	if len(req.UserId) > 0 {
		if i, err2 := strconv.ParseInt(req.UserId, 10, 32); err2 != nil {
			return err2
		} else {
			m.UserId = int32(i)
		}
	}
	if m.Thumb, err = MakeInputFile2(c, contentType, "thumb", req.Thumb); err != nil {
		return
	}

	return
}
