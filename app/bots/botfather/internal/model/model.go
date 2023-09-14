package model

const (
	BotFatherID = int32(6)
)

type CommandInfo struct {
	CmdName     string
	Description string
}

type BotFatherCommandStates struct {
	MainCmd            string            `json:"main_cmd"`
	NextSubCmd         string            `json:"next_sub_cmd,omitempty"`
	CacheSubCmdResults map[string]string `json:"cache_sub_cmd_results,omitempty"`
}

func NewBotFatherCommandStates() *BotFatherCommandStates {
	return &BotFatherCommandStates{
		MainCmd:            "",
		CacheSubCmdResults: make(map[string]string),
	}
}
