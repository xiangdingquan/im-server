package model

import (
	"context"
	"open.chat/app/sysconfig"
	"open.chat/mtproto"
	"strings"
)

func FixBannedWord(ctx context.Context, text string) string {
	words := sysconfig.GetConfig2StringArray(ctx, sysconfig.ConfigKeysBanWords, nil, 0)
	replaces := make([]string, 0)
	for _, w := range words {
		if len(w) > 0 && strings.Contains(text, w) {
			replaces = append(replaces, w)
		}
	}

	for _, r := range replaces {
		s := ""
		for i := 0; i < len(r); i++ {
			s += "*"
		}
		text = strings.Replace(text, r, s, -1)
	}

	return text
}

func FixInputMentionNameEntity(org mtproto.MessageEntitySlice) (entities mtproto.MessageEntitySlice) {
	for _, entity := range org {
		switch entity.PredicateName {
		case mtproto.Predicate_inputMessageEntityMentionName:
			entityMentionName := mtproto.MakeTLMessageEntityMentionName(&mtproto.MessageEntity{
				Offset:       entity.Offset,
				Length:       entity.Length,
				UserId_INT32: entity.UserId_INPUTUSER.UserId,
			})
			entities = append(entities, entityMentionName.To_MessageEntity())
		default:
			entities = append(entities, entity)
		}
	}

	return
}
