package main

import (
	"open.chat/app/interface/botway/internal/server"
	"open.chat/pkg/commands"
)

func main() {
	commands.Run(server.New())
}
