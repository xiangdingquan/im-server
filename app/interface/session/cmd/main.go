package main

import (
	"open.chat/app/interface/session/internal/server"
	_ "open.chat/mtproto"
	"open.chat/pkg/commands"
)

func main() {
	commands.Run(server.New())
}
