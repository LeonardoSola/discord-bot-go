package controllers

import (
	"fmt"
	"discord-bot-go/models"
)

func Ping(ctx models.Context) {
	message, err := ctx.Reply("Ping...")
	if err != nil {
		fmt.Println(err)
		return
	}

	miliseconds := message.Timestamp.Sub(ctx.Message.Timestamp).Milliseconds()

	ctx.Session.ChannelMessageEdit(ctx.Message.ChannelID, message.ID, fmt.Sprintf("Pong! (%dms)", miliseconds))
}
