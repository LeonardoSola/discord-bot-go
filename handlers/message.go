package handlers

import (
	"discord-bot-go/commands"
	"discord-bot-go/models"
	"discord-bot-go/utils"
	"strings"

	"github.com/spf13/viper"
)

// MessageHandler handles all messages
func messageHandler(ctx models.Context) {
	if ctx.Message.Author.ID == ctx.Session.State.User.ID {
		return
	}

	arguments := strings.Split(ctx.Message.Content, " ")
	if len(arguments) == 0 {
		return
	}

	if arguments[0][:1] != viper.GetString("PREFIX") {
		return
	}

	command := arguments[0][1:]
	arguments = arguments[1:]
	argument := strings.Join(arguments, " ")

	ctx.Data = models.Data{
		Command:   command,
		Arguments: arguments,
		Argument:  argument,
	}

	// Check if the command is a valid command
	for _, c := range commands.Commands {
		if utils.Includes(c.Aliases, ctx.Data.Command) {
			c.Handler(ctx)
		}
	}
}
