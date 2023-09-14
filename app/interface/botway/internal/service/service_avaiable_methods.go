package service

import (
	"context"
	"fmt"
	"github.com/gogo/protobuf/types"
	"math/rand"
	"open.chat/app/interface/botway/botapi"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/log"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func (s *Service) GetMe(ctx context.Context, token string, r *botapi.GetMe2) (user *botapi.User, err error) {
	var (
		resp mtproto.TLObject
		ok   bool
	)

	resp, err = s.Invoke(ctx, token, EncodeToTLMethod(r))
	if err != nil {
		log.Errorf("callSession error: %v", err)
	} else {
		if users, ok2 := resp.(*mtproto.Vector_User); ok2 {
			log.Debugf("users: %s", users.DebugString())
			if len(users.Datas) > 0 {
				if user, ok = DecodeToBotApi(users.Datas[0]).(*botapi.User); !ok {
				}
			}
		}
	}

	return
}

func ToMessage(updates *mtproto.Updates) (m *botapi.Message) {
	model.VisitUpdates(0, updates, map[string]model.UpdateVisitedFunc{
		mtproto.Predicate_updateNewMessage: func(userId int32,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32) {

			m = messageToMessage(update.GetMessage_MESSAGE(), users, chats)
		},
		mtproto.Predicate_updateNewChannelMessage: func(userId int32,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32) {

			m = messageToMessage(update.GetMessage_MESSAGE(), users, chats)
		},
		mtproto.Predicate_updateEditMessage: func(userId int32,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32) {

			m = messageToMessage(update.GetMessage_MESSAGE(), users, chats)
		},
		mtproto.Predicate_updateEditChannelMessage: func(userId int32,
			update *mtproto.Update,
			users []*mtproto.User,
			chats []*mtproto.Chat,
			date int32) {

			m = messageToMessage(update.GetMessage_MESSAGE(), users, chats)
		},
	})
	return
}

func (s *Service) SendMessage(ctx context.Context, token string, r *botapi.SendMessage2) (message *botapi.Message, err error) {
	var (
		resp mtproto.TLObject
		me   *mtproto.Updates
		ok   bool
	)

	i := &mtproto.TLMessagesSendMessage{
		Constructor:  mtproto.CRC32_messages_sendMessage_520c3870,
		NoWebpage:    r.DisableWebPagePreview,
		Silent:       r.DisableNotification,
		Background:   false,
		ClearDraft:   false,
		Peer:         MakePeer(r.ChatId),
		ReplyToMsgId: mtproto.MakeFlagsInt32(r.ReplyToMessageId), //
		Message:      r.Text,
		RandomId:     rand.Int63(),
		ReplyMarkup:  encodeToReplyMarkup(r.ReplyMarkup),
		Entities:     nil,
		ScheduleDate: nil,
	}

	log.Debugf("send: %s", i.DebugString())
	resp, err = s.Invoke(ctx, token, i)
	if err != nil {
		return nil, err
	}

	if me, ok = resp.(*mtproto.Updates); !ok {
		return nil, fmt.Errorf("not updates... %v", ok)
	}

	message = ToMessage(me)
	return
}

func (s *Service) ForwardMessage(ctx context.Context, token string, r *botapi.ForwardMessage2) (message *botapi.Message, err error) {
	var (
		resp mtproto.TLObject
		me   *mtproto.Updates
		ok   bool
	)

	i := &mtproto.TLMessagesForwardMessages{
		Constructor:  mtproto.CRC32_messages_forwardMessages_d9fee60e,
		Silent:       r.DisableNotification,
		Background:   false,
		WithMyScore:  false,
		Grouped:      false,
		FromPeer:     MakePeer(r.FromChatId),
		Id:           []int32{r.MessageId},
		RandomId:     []int64{rand.Int63()},
		ToPeer:       MakePeer(r.ChatId),
		ScheduleDate: nil,
	}

	resp, err = s.Invoke(ctx, token, i)

	if err != nil {
		return nil, err
	}

	if me, ok = resp.(*mtproto.Updates); !ok {
		return nil, fmt.Errorf("not updates... %v", ok)
	}

	message = ToMessage(me)
	return
}

func (s *Service) SendPhoto(ctx context.Context, token string, r *botapi.SendPhoto2) (message *botapi.Message, err error) {
	var (
		resp mtproto.TLObject
		me   *mtproto.Updates
		ok   bool
	)

	a := strings.Split(token, ":")
	if len(a) != 2 {
		return nil, fmt.Errorf("invalid tolen: %s", token)
	}

	auth, err := s.dao.GetCacheAuthUser(ctx, a[0], a[1])
	if err != nil {
		log.Errorf("getBotSessionByToken error: %v", err)
		return nil, err
	}

	i := &mtproto.TLMessagesSendMedia{
		Constructor:  mtproto.CRC32_messages_sendMedia_3491eba9,
		Silent:       r.DisableNotification,
		Background:   false,
		ClearDraft:   false,
		Peer:         MakePeer(r.ChatId),
		ReplyToMsgId: mtproto.MakeFlagsInt32(r.ReplyToMessageId),
		Media:        nil,
		Message:      r.Caption,
		RandomId:     rand.Int63(),
		ReplyMarkup:  encodeToReplyMarkup(r.ReplyMarkup),
		Entities:     nil,
		ScheduleDate: nil,
	}

	switch f := r.Photo.(type) {
	case botapi.InputFileUpload:
		if f.FileName == "" {
			f.FileName = fmt.Sprintf("image%d.jpg", time.Now().Unix())
		} else {
			f.FileName += ".jpg"
		}
		i.Media = mtproto.MakeTLInputMediaUploadedPhoto(&mtproto.InputMedia{
			File: ToInputFile(
				&f,
				func(fileId int64, filePart int32, bytes []byte) error {
					s.DfsFacade.WriteFilePartData(ctx, auth.AuthKeyId(), fileId, filePart, bytes)
					return nil
				}),
			Stickers:   nil,
			TtlSeconds: nil,
		}).To_InputMedia()
	}

	resp, err = s.Invoke(ctx, token, i)
	if err != nil {
		return nil, err
	}
	if me, ok = resp.(*mtproto.Updates); !ok {
		return nil, fmt.Errorf("not updates... %v", ok)
	}
	message = ToMessage(me)

	return
}

func (s *Service) SendAudio(ctx context.Context, token string, r *botapi.SendAudio2) (message *botapi.Message, err error) {
	var (
		resp mtproto.TLObject
		me   *mtproto.Updates
		ok   bool
	)

	a := strings.Split(token, ":")
	auth, err := s.dao.GetCacheAuthUser(ctx, a[0], a[1])
	if err != nil {
		log.Errorf("getBotSessionByToken error: %v", err)
		return nil, err
	}

	i := &mtproto.TLMessagesSendMedia{
		Silent:       r.DisableNotification,
		Background:   false,
		ClearDraft:   false,
		Peer:         MakePeer(r.ChatId),
		ReplyToMsgId: mtproto.MakeFlagsInt32(r.ReplyToMessageId),
		Media:        nil,
		Message:      r.Caption,
		RandomId:     rand.Int63(),
		ReplyMarkup:  encodeToReplyMarkup(r.ReplyMarkup),
		Entities:     nil,
		ScheduleDate: nil,
	}

	switch f := r.Audio.(type) {
	case botapi.InputFileUpload:
		if f.FileName == "" {
			f.FileName = fmt.Sprintf("voice%d.mp3", time.Now().Unix())
		}
		i.Media = mtproto.MakeTLInputMediaUploadedDocument(&mtproto.InputMedia{
			NosoundVideo: false,
			File: ToInputFile(
				&f,
				func(fileId int64, filePart int32, bytes []byte) error {
					s.DfsFacade.WriteFilePartData(ctx, auth.AuthKeyId(), fileId, filePart, bytes)
					return nil
				}),
			Thumb:    nil,
			MimeType: model.GuessMimeTypeByFileExtension(filepath.Ext(f.FileName)),
			Attributes: []*mtproto.DocumentAttribute{
				mtproto.MakeTLDocumentAttributeFilename(&mtproto.DocumentAttribute{
					FileName: f.FileName,
				}).To_DocumentAttribute(),
				mtproto.MakeTLDocumentAttributeAudio(&mtproto.DocumentAttribute{
					Voice:     true,
					Duration:  r.Duration,
					Title:     mtproto.MakeFlagsString(r.Title),
					Performer: mtproto.MakeFlagsString(r.Performer),
				}).To_DocumentAttribute(),
			},
			Stickers:   nil,
			TtlSeconds: nil,
		}).To_InputMedia()
	}

	resp, err = s.Invoke(ctx, token, i)
	if err != nil {
		return nil, err
	}
	if me, ok = resp.(*mtproto.Updates); !ok {
		return nil, fmt.Errorf("not updates... %v", ok)
	}
	message = ToMessage(me)

	return
}

func (s *Service) SendDocument(ctx context.Context, token string, r *botapi.SendDocument2) (message *botapi.Message, err error) {
	var (
		resp mtproto.TLObject
		me   *mtproto.Updates
		ok   bool
	)

	a := strings.Split(token, ":")
	auth, err := s.dao.GetCacheAuthUser(ctx, a[0], a[1])
	if err != nil {
		log.Errorf("getBotSessionByToken error: %v", err)
		return nil, err
	}

	i := &mtproto.TLMessagesSendMedia{
		Silent:       r.DisableNotification,
		Background:   false,
		ClearDraft:   false,
		Peer:         MakePeer(r.ChatId),
		ReplyToMsgId: mtproto.MakeFlagsInt32(r.ReplyToMessageId),
		Media:        nil,
		Message:      r.Caption,
		RandomId:     rand.Int63(),
		ReplyMarkup:  encodeToReplyMarkup(r.ReplyMarkup),
		Entities:     nil,
		ScheduleDate: nil,
	}

	switch f := r.Document.(type) {
	case botapi.InputFileUpload:
		if f.FileName == "" {
			f.FileName = fmt.Sprintf("doc%d", time.Now().Unix())
		}
		i.Media = mtproto.MakeTLInputMediaUploadedDocument(&mtproto.InputMedia{
			NosoundVideo: false,
			File: ToInputFile(
				&f,
				func(fileId int64, filePart int32, bytes []byte) error {
					s.DfsFacade.WriteFilePartData(ctx, auth.AuthKeyId(), fileId, filePart, bytes)
					return nil
				}),
			Thumb:    nil,
			MimeType: model.GuessMimeTypeByFileExtension(filepath.Ext(f.FileName)),
			Attributes: []*mtproto.DocumentAttribute{
				mtproto.MakeTLDocumentAttributeFilename(&mtproto.DocumentAttribute{
					FileName: f.FileName,
				}).To_DocumentAttribute(),
			},
			Stickers:   nil,
			TtlSeconds: nil,
		}).To_InputMedia()
	}

	resp, err = s.Invoke(ctx, token, i)
	if err != nil {
		return nil, err
	}
	if me, ok = resp.(*mtproto.Updates); !ok {
		return nil, fmt.Errorf("not updates... %v", ok)
	}
	message = ToMessage(me)

	return
}

func (s *Service) SendVideo(ctx context.Context, token string, r *botapi.SendVideo2) (message *botapi.Message, err error) {
	var (
		resp mtproto.TLObject
		me   *mtproto.Updates
		ok   bool
	)

	a := strings.Split(token, ":")
	auth, err := s.dao.GetCacheAuthUser(ctx, a[0], a[1])
	if err != nil {
		log.Errorf("getBotSessionByToken error: %v", err)
		return nil, err
	}

	i := &mtproto.TLMessagesSendMedia{
		Silent:       r.DisableNotification,
		Background:   false,
		ClearDraft:   false,
		Peer:         MakePeer(r.ChatId),
		ReplyToMsgId: mtproto.MakeFlagsInt32(r.ReplyToMessageId),
		Media:        nil,
		Message:      r.Caption,
		RandomId:     rand.Int63(),
		ReplyMarkup:  encodeToReplyMarkup(r.ReplyMarkup),
		Entities:     nil,
		ScheduleDate: nil,
	}

	switch f := r.Video.(type) {
	case botapi.InputFileUpload:
		if f.FileName == "" {
			f.FileName = fmt.Sprintf("video%d", time.Now().Unix())
		}
		i.Media = mtproto.MakeTLInputMediaUploadedDocument(&mtproto.InputMedia{
			NosoundVideo: false,
			File: ToInputFile(
				&f,
				func(fileId int64, filePart int32, bytes []byte) error {
					s.DfsFacade.WriteFilePartData(ctx, auth.AuthKeyId(), fileId, filePart, bytes)
					return nil
				}),
			Thumb:    nil,
			MimeType: model.GuessMimeTypeByFileExtension(filepath.Ext(f.FileName)),
			Attributes: []*mtproto.DocumentAttribute{
				mtproto.MakeTLDocumentAttributeFilename(&mtproto.DocumentAttribute{
					FileName: f.FileName,
				}).To_DocumentAttribute(),
				mtproto.MakeTLDocumentAttributeVideo(&mtproto.DocumentAttribute{
					RoundMessage:      false,
					SupportsStreaming: false,
					Duration:          r.Duration,
					W:                 r.Width,
					H:                 r.Height,
				}).To_DocumentAttribute(),
			},
			Stickers:   nil,
			TtlSeconds: nil,
		}).To_InputMedia()
	}

	resp, err = s.Invoke(ctx, token, i)
	if err != nil {
		return nil, err
	}
	if me, ok = resp.(*mtproto.Updates); !ok {
		return nil, fmt.Errorf("not updates... %v", ok)
	}
	message = ToMessage(me)

	return
}

func (s *Service) SendAnimation(ctx context.Context, token string, r *botapi.SendAnimation2) (message *botapi.Message, err error) {
	var (
		resp mtproto.TLObject
		me   *mtproto.Updates
		ok   bool
	)

	a := strings.Split(token, ":")
	auth, err := s.dao.GetCacheAuthUser(ctx, a[0], a[1])
	if err != nil {
		log.Errorf("getBotSessionByToken error: %v", err)
		return nil, err
	}

	i := &mtproto.TLMessagesSendMedia{
		Silent:       r.DisableNotification,
		Background:   false,
		ClearDraft:   false,
		Peer:         MakePeer(r.ChatId),
		ReplyToMsgId: mtproto.MakeFlagsInt32(r.ReplyToMessageId),
		Media:        nil,
		Message:      r.Caption,
		RandomId:     rand.Int63(),
		ReplyMarkup:  encodeToReplyMarkup(r.ReplyMarkup),
		Entities:     nil,
		ScheduleDate: nil,
	}

	switch f := r.Animation.(type) {
	case botapi.InputFileUpload:
		if f.FileName == "" {
			f.FileName = fmt.Sprintf("video%d", time.Now().Unix())
		}
		i.Media = mtproto.MakeTLInputMediaUploadedDocument(&mtproto.InputMedia{
			NosoundVideo: false,
			File: ToInputFile(
				&f,
				func(fileId int64, filePart int32, bytes []byte) error {
					s.DfsFacade.WriteFilePartData(ctx, auth.AuthKeyId(), fileId, filePart, bytes)
					return nil
				}),
			Thumb:    nil,
			MimeType: model.GuessMimeTypeByFileExtension(filepath.Ext(f.FileName)),
			Attributes: []*mtproto.DocumentAttribute{
				mtproto.MakeTLDocumentAttributeFilename(&mtproto.DocumentAttribute{
					FileName: f.FileName,
				}).To_DocumentAttribute(),
				mtproto.MakeTLDocumentAttributeImageSize(&mtproto.DocumentAttribute{
					W: r.Width,
					H: r.Height,
				}).To_DocumentAttribute(),
				mtproto.MakeTLDocumentAttributeAnimated(&mtproto.DocumentAttribute{
					//
				}).To_DocumentAttribute(),
			},
			Stickers:   nil,
			TtlSeconds: nil,
		}).To_InputMedia()
	}

	resp, err = s.Invoke(ctx, token, i)
	if err != nil {
		return nil, err
	}
	if me, ok = resp.(*mtproto.Updates); !ok {
		return nil, fmt.Errorf("not updates... %v", ok)
	}
	message = ToMessage(me)

	return
}

func (s *Service) SendVoice(ctx context.Context, token string, r *botapi.SendVoice2) (message *botapi.Message, err error) {
	var (
		resp mtproto.TLObject
		me   *mtproto.Updates
		ok   bool
	)

	a := strings.Split(token, ":")
	auth, err := s.dao.GetCacheAuthUser(ctx, a[0], a[1])
	if err != nil {
		log.Errorf("getBotSessionByToken error: %v", err)
		return nil, err
	}

	i := &mtproto.TLMessagesSendMedia{
		Silent:       r.DisableNotification,
		Background:   false,
		ClearDraft:   false,
		Peer:         MakePeer(r.ChatId),
		ReplyToMsgId: mtproto.MakeFlagsInt32(r.ReplyToMessageId),
		Media:        nil,
		Message:      r.Caption,
		RandomId:     rand.Int63(),
		ReplyMarkup:  encodeToReplyMarkup(r.ReplyMarkup),
		Entities:     nil,
		ScheduleDate: nil,
	}

	switch f := r.Voice.(type) {
	case botapi.InputFileUpload:
		if f.FileName == "" {
			f.FileName = fmt.Sprintf("audio%d.ogg", time.Now().Unix())
		}
		i.Media = mtproto.MakeTLInputMediaUploadedDocument(&mtproto.InputMedia{
			NosoundVideo: false,
			File: ToInputFile(
				&f,
				func(fileId int64, filePart int32, bytes []byte) error {
					s.DfsFacade.WriteFilePartData(ctx, auth.AuthKeyId(), fileId, filePart, bytes)
					return nil
				}),
			Thumb:    nil,
			MimeType: model.GuessMimeTypeByFileExtension(filepath.Ext(f.FileName)),
			Attributes: []*mtproto.DocumentAttribute{
				mtproto.MakeTLDocumentAttributeFilename(&mtproto.DocumentAttribute{
					FileName: f.FileName,
				}).To_DocumentAttribute(),
				mtproto.MakeTLDocumentAttributeAudio(&mtproto.DocumentAttribute{
					Voice:     true,
					Duration:  r.Duration,
					Title:     nil,
					Performer: nil,
					Waveform:  getWaveform2(f.FileUpload),
				}).To_DocumentAttribute(),
			},
			Stickers:   nil,
			TtlSeconds: nil,
		}).To_InputMedia()
	}

	resp, err = s.Invoke(ctx, token, i)
	if err != nil {
		return nil, err
	}
	if me, ok = resp.(*mtproto.Updates); !ok {
		return nil, fmt.Errorf("not updates... %v", ok)
	}
	message = ToMessage(me)

	return

	return nil, mtproto.ErrMethodNotImpl
}

func (s *Service) SendVideoNote(ctx context.Context, token string, r *botapi.SendVideoNote2) (message *botapi.Message, err error) {
	log.Warnf("sendVideoNote - method not impl")
	return nil, mtproto.ErrMethodNotImpl
}

func (s *Service) SendMediaGroup(ctx context.Context, token string, r *botapi.SendMediaGroup2) (message *botapi.Message, err error) {
	log.Warnf("sendMediaGroup - method not impl")
	return nil, mtproto.ErrMethodNotImpl
}

func (s *Service) SendLocation(ctx context.Context, token string, r *botapi.SendLocation2) (message *botapi.Message, err error) {
	log.Warnf("sendLocation - method not impl")
	return nil, mtproto.ErrMethodNotImpl
}

func (s *Service) EditMessageLiveLocation(ctx context.Context, token string, r *botapi.EditMessageLiveLocation2) (message *botapi.Message, err error) {
	log.Warnf("editMessageLiveLocation - method not impl")
	return nil, mtproto.ErrMethodNotImpl
}

func (s *Service) StopMessageLiveLocation(ctx context.Context, token string, r *botapi.StopMessageLiveLocation2) (message *botapi.Message, err error) {
	log.Warnf("stopMessageLiveLocation - method not impl")
	return nil, mtproto.ErrMethodNotImpl
}

func (s *Service) SendVenue(ctx context.Context, token string, r *botapi.SendVenue2) (message *botapi.Message, err error) {
	log.Warnf("sendVenue - method not impl")
	return nil, mtproto.ErrMethodNotImpl
}

func (s *Service) SendContact(ctx context.Context, token string, r *botapi.SendContact2) (message *botapi.Message, err error) {
	var (
		resp mtproto.TLObject
		me   *mtproto.Updates
		ok   bool
	)

	i := &mtproto.TLMessagesSendMedia{
		Silent:       r.DisableNotification,
		Background:   false,
		ClearDraft:   false,
		Peer:         MakePeer(r.ChatId),
		ReplyToMsgId: mtproto.MakeFlagsInt32(r.ReplyToMessageId),
		Media:        nil,
		Message:      "",
		RandomId:     rand.Int63(),
		ReplyMarkup:  encodeToReplyMarkup(r.ReplyMarkup),
		Entities:     nil,
		ScheduleDate: nil,
	}

	i.Media = mtproto.MakeTLInputMediaContact(&mtproto.InputMedia{
		PhoneNumber: r.PhoneNumber,
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		Vcard:       r.Vcard,
	}).To_InputMedia()

	log.Debugf("send: %s", i.DebugString())
	resp, err = s.Invoke(ctx, token, i)
	if err != nil {
		return nil, err
	}

	if me, ok = resp.(*mtproto.Updates); !ok {
		return nil, fmt.Errorf("not updates... %v", ok)
	}

	message = ToMessage(me)
	return
}

func (s *Service) SendPoll(ctx context.Context, token string, r *botapi.SendPoll2) (message *botapi.Message, err error) {
	log.Warnf("sendPoll - method not impl")
	return nil, mtproto.ErrMethodNotImpl
}

func (s *Service) SendChatAction(ctx context.Context, token string, r *botapi.SendChatAction2) (result bool, err error) {
	var (
		resp mtproto.TLObject
		ok   bool
	)

	i := &mtproto.TLMessagesSetTyping{
		Peer:   MakePeer(r.ChatId),
		Action: nil,
	}

	switch botapi.ChatAction(r.Action) {
	case botapi.Typing:
		i.Action = mtproto.MakeTLSendMessageTypingAction(&mtproto.SendMessageAction{}).To_SendMessageAction()
	case botapi.UploadingPhoto:
		i.Action = mtproto.MakeTLSendMessageUploadPhotoAction(&mtproto.SendMessageAction{}).To_SendMessageAction()
	case botapi.UploadingVideo:
		i.Action = mtproto.MakeTLSendMessageUploadVideoAction(&mtproto.SendMessageAction{}).To_SendMessageAction()
	case botapi.UploadingAudio:
		i.Action = mtproto.MakeTLSendMessageUploadAudioAction(&mtproto.SendMessageAction{}).To_SendMessageAction()
	case botapi.UploadingDocument:
		i.Action = mtproto.MakeTLSendMessageUploadDocumentAction(&mtproto.SendMessageAction{}).To_SendMessageAction()
	case botapi.UploadingVNote:
		i.Action = mtproto.MakeTLSendMessageUploadVideoAction(&mtproto.SendMessageAction{}).To_SendMessageAction()
	case botapi.RecordingVideo:
		i.Action = mtproto.MakeTLSendMessageRecordVideoAction(&mtproto.SendMessageAction{}).To_SendMessageAction()
	case botapi.RecordingAudio:
		i.Action = mtproto.MakeTLSendMessageRecordAudioAction(&mtproto.SendMessageAction{}).To_SendMessageAction()
	case botapi.FindingLocation:
		i.Action = mtproto.MakeTLSendMessageGeoLocationAction(&mtproto.SendMessageAction{}).To_SendMessageAction()
	default:
		log.Errorf("invalid action: %v", i)
		return false, mtproto.ErrBadRequest
	}

	log.Debugf("send: %s", i.DebugString())
	resp, err = s.Invoke(ctx, token, i)
	if err != nil {
		return false, err
	}

	if _, ok = resp.(*mtproto.Bool); !ok {
		return false, fmt.Errorf("get error... %v", ok)
	}

	return true, nil
}

func (s *Service) GetUserProfilePhotos(ctx context.Context, token string, r *botapi.GetUserProfilePhotos2) (*botapi.UserProfilePhotos, error) {
	log.Warnf("getUserProfilePhotos - method not impl")
	return nil, mtproto.ErrMethodNotImpl
}

func (s *Service) GetFile(ctx context.Context, token string, r *botapi.GetFile2) (*botapi.File, error) {
	log.Warnf("getFile - method not impl")
	return nil, mtproto.ErrMethodNotImpl
}

func (s *Service) KickChatMember(ctx context.Context, token string, r *botapi.KickChatMember2) (bool, error) {
	log.Warnf("kickChatMember - method not impl")
	return false, mtproto.ErrMethodNotImpl
}

func (s *Service) UnbanChatMember(ctx context.Context, token string, r *botapi.UnbanChatMember2) (bool, error) {
	log.Warnf("unbanChatMember - method not impl")
	return false, mtproto.ErrMethodNotImpl
}

func (s *Service) RestrictChatMember(ctx context.Context, token string, r *botapi.RestrictChatMember2) (bool, error) {
	log.Warnf("restrictChatMember - method not impl")
	return false, mtproto.ErrMethodNotImpl
}

func (s *Service) PromoteChatMember(ctx context.Context, token string, r *botapi.PromoteChatMember2) (bool, error) {
	log.Warnf("promoteChatMember - method not impl")
	return false, mtproto.ErrMethodNotImpl
}

func (s *Service) SetChatPermissions(ctx context.Context, token string, r *botapi.SetChatPermissions2) (bool, error) {
	log.Warnf("setChatPermissions - method not impl")
	return false, mtproto.ErrMethodNotImpl
}

func (s *Service) ExportChatInviteLink(ctx context.Context, token string, r *botapi.ExportChatInviteLink2) (string, error) {
	log.Warnf("exportChatInviteLink - method not impl")
	return "", mtproto.ErrMethodNotImpl
}

func (s *Service) SetChatPhoto(ctx context.Context, token string, r *botapi.SetChatPhoto2) (bool, error) {
	log.Warnf("setChatPhoto - method not impl")
	return false, mtproto.ErrMethodNotImpl
}

func (s *Service) DeleteChatPhoto(ctx context.Context, token string, r *botapi.DeleteChatPhoto2) (bool, error) {
	log.Warnf("deleteChatPhoto - method not impl")
	return false, mtproto.ErrMethodNotImpl
}

func (s *Service) SetChatTitle(ctx context.Context, token string, r *botapi.SetChatTitle2) (bool, error) {
	log.Warnf("setChatTitle - method not impl")
	return false, mtproto.ErrMethodNotImpl
}

func (s *Service) SetChatDescription(ctx context.Context, token string, r *botapi.SetChatDescription2) (bool, error) {
	log.Warnf("setChatDescription - method not impl")
	return false, mtproto.ErrMethodNotImpl
}

func (s *Service) PinChatMessage(ctx context.Context, token string, r *botapi.PinChatMessage2) (bool, error) {
	log.Warnf("pinChatMessage - method not impl")
	return false, mtproto.ErrMethodNotImpl
}

func (s *Service) UnpinChatMessage(ctx context.Context, token string, r *botapi.UnpinChatMessage2) (bool, error) {
	log.Warnf("unpinChatMessage - method not impl")
	return false, mtproto.ErrMethodNotImpl
}

func (s *Service) LeaveChat(ctx context.Context, token string, r *botapi.LeaveChat2) (bool, error) {
	log.Warnf("leaveChat - method not impl")
	return false, mtproto.ErrMethodNotImpl
}

func (s *Service) GetChat(ctx context.Context, token string, r *botapi.GetChat2) (*botapi.Chat, error) {
	log.Warnf("getChat - method not impl")
	return nil, mtproto.ErrMethodNotImpl
}

func (s *Service) GetChatAdministrators(ctx context.Context, token string, r *botapi.GetChatAdministrators2) ([]*botapi.ChatMember, error) {
	log.Warnf("getChatAdministrators - method not impl")
	return nil, mtproto.ErrMethodNotImpl
}

func (s *Service) GetChatMembersCount(ctx context.Context, token string, r *botapi.GetChatMembersCount2) (int32, error) {
	log.Warnf("getChatMembersCount - method not impl")
	return 0, mtproto.ErrMethodNotImpl
}

func (s *Service) GetChatMember(ctx context.Context, token string, r *botapi.GetChatMember2) ([]*botapi.ChatMember, error) {
	log.Warnf("getChatMember - method not impl")
	return nil, mtproto.ErrMethodNotImpl
}

func (s *Service) SetChatStickerSet(ctx context.Context, token string, r *botapi.SetChatStickerSet2) (bool, error) {
	log.Warnf("setChatStickerSet - method not impl")
	return false, mtproto.ErrMethodNotImpl
}

func (s *Service) DeleteChatStickerSet(ctx context.Context, token string, r *botapi.DeleteChatStickerSet2) (bool, error) {
	log.Warnf("deleteChatStickerSet - method not impl")
	return false, mtproto.ErrMethodNotImpl
}

func (s *Service) AnswerCallbackQuery(ctx context.Context, token string, r *botapi.AnswerCallbackQuery2) (bool, error) {
	i := &mtproto.TLMessagesSetBotCallbackAnswer{
		Constructor: mtproto.CRC32_messages_setBotCallbackAnswer,
		Alert:       r.ShowAlert,
		QueryId:     0,
		Message:     nil,
		Url:         nil,
	}

	i.QueryId, _ = strconv.ParseInt(r.CallbackQueryId, 10, 64)
	if r.Text != "" {
		i.Message = &types.StringValue{Value: r.Text}
	}
	log.Debugf("send: %s", i.DebugString())
	_, err := s.Invoke(ctx, token, i)
	if err != nil {
		return false, err
	}

	return true, nil
}
