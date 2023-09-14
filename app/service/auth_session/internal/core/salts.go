package core

import (
	"context"
	"math/rand"
	"time"

	"open.chat/app/service/auth_session/internal/model"
	"open.chat/mtproto"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func (m *AuthSessionCore) getOrNotInsertSaltList(ctx context.Context, keyId int64, size int32) ([]*mtproto.TLFutureSalt, error) {
	var (
		salts = make([]*mtproto.TLFutureSalt, 0, size)

		date           = int32(time.Now().Unix())
		lastValidUntil = date
		saltsData      []*mtproto.TLFutureSalt
		lastSalt       *mtproto.TLFutureSalt
	)

	saltList, err := m.Redis.GetSalts(ctx, keyId)
	if err != nil {
		return nil, err
	}

	if len(saltList) > 0 {
		hasLastSalt := false
		for idx, salt := range saltList {
			if salt.GetValidUntil() >= date {
				if !hasLastSalt {
					if idx > 0 {
						lastSalt = saltList[idx-1]
					}
					hasLastSalt = true
				}
				saltsData = append(saltsData, salt)
				if lastValidUntil < salt.GetValidUntil() {
					lastValidUntil = salt.GetValidUntil()
				}
			}
		}
		if !hasLastSalt {
			lastSalt = saltList[len(saltList)-1]
		}

		if lastSalt != nil && lastSalt.GetValidUntil()+300 < date {
			lastSalt = nil
		}
	}

	left := size - int32(len(saltsData))
	if left > 0 {
		for i := int32(0); i < size; i++ {
			salt := mtproto.MakeTLFutureSalt(&mtproto.FutureSalt{
				ValidSince: lastValidUntil,
				ValidUntil: lastValidUntil + model.SALT_TIMEOUT,
				Salt:       rand.Int63(),
			})
			saltsData = append(saltsData, salt)
			lastValidUntil += model.SALT_TIMEOUT
		}
	}

	for i := int32(0); i < size; i++ {
		salts = append(salts, saltsData[i])
	}

	var (
		salts2     []*mtproto.TLFutureSalt
		saltsData2 []*mtproto.TLFutureSalt
	)

	if lastSalt != nil {
		salts2 = append(salts2, lastSalt)
		saltsData2 = append(saltsData2, lastSalt)
	}

	salts2 = append(salts2, salts...)
	saltsData2 = append(saltsData2, saltsData...)

	if left > 0 {
		err = m.Redis.PutSalts(ctx, keyId, saltsData2)
		if err != nil {
			return nil, err
		}
	}
	return salts2, nil
}

func (m *AuthSessionCore) putSaltCache(ctx context.Context, keyId int64, salt *mtproto.TLFutureSalt) error {
	return m.Redis.PutSalts(ctx, keyId, []*mtproto.TLFutureSalt{salt})
}
