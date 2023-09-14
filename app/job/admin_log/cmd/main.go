package main

import (
	"open.chat/app/job/admin_log/internal/server"
	"open.chat/pkg/commands"
)

func main() {
	commands.Run(server.New())
}
