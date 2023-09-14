package model

import (
	"strings"

	"open.chat/mtproto"
)

const (
	BotFatherId   = int32(6)
	BotFatherName = "BotFather"
)

const (
	BotGifId          = int32(101)
	BotGifName        = "gif"
	BotVidId          = int32(102)
	BotVidName        = "vid"
	BotPicId          = int32(103)
	BotPicName        = "pic"
	BotBingId         = int32(104)
	BotBingName       = "bing"
	BotWikiId         = int32(105)
	BotWikiName       = "wiki"
	BotImdbId         = int32(106)
	BotImdbName       = "imdb"
	BotBoldId         = int32(107)
	BotBoldName       = "bold"
	BotYoutubeId      = int32(108)
	BotYoutubeName    = "youtube"
	BotMusicId        = int32(109)
	BotMusicName      = "music"
	BotFoursquareId   = int32(110)
	BotFoursquareName = "foursquare"
	BotStickerId      = int32(111)
	BotStickerName    = "sticker"
)

var (
	botIdNameMap = map[int32]string{
		BotFatherId:     BotFatherName,
		BotGifId:        BotGifName,
		BotVidId:        BotVidName,
		BotPicId:        BotPicName,
		BotBingId:       BotBingName,
		BotWikiId:       BotWikiName,
		BotImdbId:       BotImdbName,
		BotBoldId:       BotBoldName,
		BotYoutubeId:    BotYoutubeName,
		BotMusicId:      BotMusicName,
		BotFoursquareId: BotFoursquareName,
		BotStickerId:    BotStickerName,
	}

	botNameIdMap = map[string]int32{
		BotFatherName:     BotFatherId,
		BotGifName:        BotGifId,
		BotVidName:        BotVidId,
		BotPicName:        BotPicId,
		BotBingName:       BotBingId,
		BotWikiName:       BotWikiId,
		BotImdbName:       BotImdbId,
		BotBoldName:       BotBoldId,
		BotYoutubeName:    BotYoutubeId,
		BotMusicName:      BotMusicId,
		BotFoursquareName: BotFoursquareId,
		BotStickerName:    BotStickerId,
	}
)

func GetBotNameById(id int32) (n string) {
	n, _ = botIdNameMap[id]
	return
}

func GetBotIdByName(n string) (id int32) {
	id, _ = botNameIdMap[n]
	return
}

func IsBotFather(id int32) bool {
	return id == BotFatherId
}

func IsBotBing(id int32) bool {
	return id == BotBingId
}

func IsBotPic(id int32) bool {
	return id == BotPicId
}

func IsBotGif(id int32) bool {
	return id == BotGifId
}

func IsBotFoursquare(id int32) bool {
	return id == BotFoursquareId
}

// ////////////////////////////////////////////////////////////////////////////////
type BotCommand struct {
	CommandName string   `json:"command_name"`
	Params      []string `json:"params"`
}

func GetBotCommandByMessage(m *mtproto.Message) *BotCommand {
	commandLine := m.GetMessage()
	if len(commandLine) > 0 && commandLine[0] == '/' {
		commands := strings.Split(commandLine, " ")
		cmdName := commands[0][1:]
		if len(cmdName) > 0 {
			return &BotCommand{
				CommandName: cmdName,
				Params:      commands[1:],
			}
		}
	}

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////
type BotInlineIdResult struct {
	BotId           int32                    `json:"bot_id"`
	BotInlineResult *mtproto.BotInlineResult `json:"bot_inline_result"`
}

func (m *BotInlineIdResult) ID() string {
	if m.BotInlineResult == nil {
		return ""
	}

	return m.BotInlineResult.GetId()
}
