package core

import (
	"context"
	"open.chat/mtproto"
)

type entity struct {
	PredicateName string `json:"predicateName"`
	Offset        int32  `json:"offset"`
	Length        int32  `json:"length"`
	UserId        int32  `json:"userId"`
}

func (m *BlogCore) Marshal(ctx context.Context, entities []*mtproto.MessageEntity) (string, error) {
	return "", nil
}

func (m *BlogCore) Unmarshal(ctx context.Context, entitiesString string) ([]*mtproto.MessageEntity, error) {
	return nil, nil
}
