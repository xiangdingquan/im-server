package dao

import (
	"context"
	"fmt"
	"math/rand"

	"open.chat/app/bots/botfather/internal/dal/dataobject"
	"open.chat/app/bots/botfather/internal/model"
	"open.chat/app/service/auth_session/authsessionpb"
	model2 "open.chat/model"
	"open.chat/pkg/crypto"
	"open.chat/pkg/database/sqlx"
)

func (d *Dao) CreateNewBot(ctx context.Context, creator int32, botName, username, token string) (int32, error) {
	key := crypto.CreateAuthKey()
	_, err := d.RPCSessionClient.SessionSetAuthKey(ctx, &authsessionpb.TLSessionSetAuthKey{
		AuthKey: &authsessionpb.AuthKeyInfo{
			AuthKeyId:          key.AuthKeyId(),
			AuthKey:            key.AuthKey(),
			AuthKeyType:        model2.AuthKeyTypePerm,
			PermAuthKeyId:      key.AuthKeyId(),
			TempAuthKeyId:      0,
			MediaTempAuthKeyId: 0,
		},
		FutureSalt: nil,
	})
	if err != nil {
		return 0, err
	}

	tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		usersDO := &dataobject.UsersDO{
			UserType:    2,
			AccessHash:  rand.Int63(),
			SecretKeyId: key.AuthKeyId(),
			FirstName:   botName,
			Username:    username,
			Phone:       username,
			CountryCode: "CN",
		}
		botId, _, err := d.UsersDAO.InsertTx(tx, usersDO)
		if err != nil {
			result.Err = err
			return
		}

		botsDO := &dataobject.BotsDO{
			BotId:         int32(botId),
			BotType:       0,
			CreatorUserId: creator,
			Token:         token,
		}
		_, _, err = d.BotsDAO.InsertOrUpdateTx(tx, botsDO)
		if err != nil {
			result.Err = err
			return
		}

		result.Data = botsDO.BotId
	})

	if tR.Err != nil {
		return 0, tR.Err
	}

	return tR.Data.(int32), nil
}

func (d *Dao) WalkMyBots(ctx context.Context, creator int32, enumF func(botId int32, botUserName string)) error {
	doList, err := d.UsersDAO.SelectListByCreator(ctx, creator)
	if err != nil {
		return err
	}

	for i := 0; i < len(doList); i++ {
		enumF(doList[i].Id, doList[i].Username)
	}

	return nil
}

func (d *Dao) GetBotToken(ctx context.Context, creator, botId int32) (string, error) {
	botsDO, err := d.BotsDAO.Select(ctx, botId)
	if err != nil {
		return "", err
	}

	if botsDO.CreatorUserId != creator {
		err = fmt.Errorf("invalid creator (%d, %d)", creator, botId)
		return "", err
	}

	return botsDO.Token, nil
}

func (d *Dao) UpdateBotToken(ctx context.Context, botId int32, token string) (err error) {
	for i := 0; i < 3; i++ {
		if _, err = d.BotsDAO.UpdateToken(ctx, token, botId); err == nil {
			break
		}
	}
	return
}

func (d *Dao) UpdateBotDescription(ctx context.Context, botId int32, description string) error {
	_, err := d.BotsDAO.UpdateDescription(ctx, description, botId)
	return err
}

func (d *Dao) SetBotCommands(ctx context.Context, botId int32, cmdList []model.CommandInfo) error {
	if len(cmdList) == 0 {
		return nil
	}
	cList := make([]*dataobject.BotCommandsDO, 0, len(cmdList))
	for i := 0; i < len(cmdList); i++ {
		cList = append(cList, &dataobject.BotCommandsDO{
			BotId:       botId,
			Command:     cmdList[i].CmdName,
			Description: cmdList[i].Description,
		})
	}
	tR := sqlx.TxWrapper(ctx, d.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		d.BotCommandsDAO.DeleteTx(tx, botId)
		d.InsertBulkTx(tx, cList)
	})

	return tR.Err
}
