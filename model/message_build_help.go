package model

import "open.chat/mtproto"

type MessageBuildEntry struct {
	Text           string
	Param          string
	EntityType     string
	EntityUrl      string
	EntityUserId   int32
	EntityLanguage string
}

type MessageBuildHelper []MessageBuildEntry

func MakeTextAndMessageEntities(m MessageBuildHelper) (text string, entities []*mtproto.MessageEntity) {
	if len(m) == 0 {
		return
	}

	var (
		offset int
		length int
	)
	for i := 0; i < len(m); i++ {
		text += m[i].Text
		offset = len(text)
		length = len(m[i].Param)
		if length > 0 {
			switch m[i].EntityType {
			case mtproto.Predicate_messageEntityUnknown:
				entities = append(entities, mtproto.MakeTLMessageEntityUnknown(&mtproto.MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case mtproto.Predicate_messageEntityMention:
				entities = append(entities, mtproto.MakeTLMessageEntityMention(&mtproto.MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case mtproto.Predicate_messageEntityHashtag:
				entities = append(entities, mtproto.MakeTLMessageEntityHashtag(&mtproto.MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case mtproto.Predicate_messageEntityBotCommand:
				entities = append(entities, mtproto.MakeTLMessageEntityBotCommand(&mtproto.MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case mtproto.Predicate_messageEntityUrl:
				entities = append(entities, mtproto.MakeTLMessageEntityUrl(&mtproto.MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case mtproto.Predicate_messageEntityEmail:
				entities = append(entities, mtproto.MakeTLMessageEntityEmail(&mtproto.MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case mtproto.Predicate_messageEntityBold:
				entities = append(entities, mtproto.MakeTLMessageEntityBold(&mtproto.MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case mtproto.Predicate_messageEntityItalic:
				entities = append(entities, mtproto.MakeTLMessageEntityItalic(&mtproto.MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case mtproto.Predicate_messageEntityCode:
				entities = append(entities, mtproto.MakeTLMessageEntityCode(&mtproto.MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case mtproto.Predicate_messageEntityPre:
				entities = append(entities, mtproto.MakeTLMessageEntityPre(&mtproto.MessageEntity{
					Offset:   int32(offset),
					Length:   int32(length),
					Language: m[i].EntityLanguage,
				}).To_MessageEntity())
			case mtproto.Predicate_messageEntityTextUrl:
				entities = append(entities, mtproto.MakeTLMessageEntityTextUrl(&mtproto.MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
					Url:    m[i].EntityUrl,
				}).To_MessageEntity())
			case mtproto.Predicate_messageEntityMentionName:
				entities = append(entities, mtproto.MakeTLMessageEntityMentionName(&mtproto.MessageEntity{
					Offset:       int32(offset),
					Length:       int32(length),
					UserId_INT32: m[i].EntityUserId,
				}).To_MessageEntity())
			case mtproto.Predicate_messageEntityPhone:
				entities = append(entities, mtproto.MakeTLMessageEntityPhone(&mtproto.MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case mtproto.Predicate_messageEntityCashtag:
				entities = append(entities, mtproto.MakeTLMessageEntityCashtag(&mtproto.MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case mtproto.Predicate_messageEntityUnderline:
				entities = append(entities, mtproto.MakeTLMessageEntityUnderline(&mtproto.MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case mtproto.Predicate_messageEntityStrike:
				entities = append(entities, mtproto.MakeTLMessageEntityStrike(&mtproto.MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case mtproto.Predicate_messageEntityBlockquote:
				entities = append(entities, mtproto.MakeTLMessageEntityBlockquote(&mtproto.MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			}
		}
		text = text + m[i].Param
	}

	return
}
