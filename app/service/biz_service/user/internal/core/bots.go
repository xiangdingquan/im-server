package core

import (
	"context"

	"open.chat/app/service/biz_service/user/internal/dal/dataobject"
	"open.chat/mtproto"
	"open.chat/pkg/database/sqlx"
)

func (m *UserCore) GetBotInfo(ctx context.Context, botId int32) *mtproto.BotInfo {
	botsDO, _ := m.BotsDAO.Select(ctx, botId)
	if botsDO == nil {
		return nil
	}
	botInfo := mtproto.MakeTLBotInfo(&mtproto.BotInfo{
		UserId:      botId,
		Description: botsDO.Description,
	}).To_BotInfo()

	botCommandsDOList, _ := m.BotCommandsDAO.SelectList(ctx, botId)
	for i := 0; i < len(botCommandsDOList); i++ {
		botCommand := mtproto.MakeTLBotCommand(&mtproto.BotCommand{
			Command:     botCommandsDOList[i].Command,
			Description: botCommandsDOList[i].Description,
		})
		botInfo.Commands = append(botInfo.Commands, botCommand.To_BotCommand())
	}

	return botInfo
}

func (m *UserCore) SetBotCommands(ctx context.Context, botId int32, commands []*mtproto.BotCommand) error {
	if len(commands) == 0 {
		return nil
	}
	cList := make([]*dataobject.BotCommandsDO, 0, len(commands))
	for i := 0; i < len(commands); i++ {
		cList = append(cList, &dataobject.BotCommandsDO{
			BotId:       botId,
			Command:     commands[i].Command,
			Description: commands[i].Description,
		})
	}
	tR := sqlx.TxWrapper(ctx, m.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
		m.BotCommandsDAO.DeleteTx(tx, botId)
		m.BotCommandsDAO.InsertBulkTx(tx, cList)
	})

	return tR.Err
}

func (m *UserCore) IsBot(ctx context.Context, id int32) bool {
	do, _ := m.BotsDAO.Select(ctx, id)
	return do != nil
}
