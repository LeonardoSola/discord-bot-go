package handlers

import (
	"discord-bot-go/config"
	"discord-bot-go/models"

	"github.com/bwmarrin/discordgo"
)

// SetupHandlers sets up all the handlers for the bot.
func SetupHandlers() {
	// Message
	config.Discord.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		messageHandler(models.Context{
			Session: s,
			Message: m,
		})
	})
}
