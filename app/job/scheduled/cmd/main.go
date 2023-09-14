package main

import (
	"open.chat/app/job/scheduled/internal/server"
	"open.chat/pkg/commands"
)

func main() {
	commands.Run(server.New())
}
