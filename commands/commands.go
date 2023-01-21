package commands

import (
	"discord-bot-go/controllers"
	"discord-bot-go/models"
)

var Commands = []models.Command{
	{
		Name:        "Ping",
		Description: "Pings the bot.",
		Aliases:     []string{"ping", "p"},
		Handler:     controllers.Ping,
	},
}
