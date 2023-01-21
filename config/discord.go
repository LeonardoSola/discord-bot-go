package config

import (
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

var (
	// Discord is the Discord session.
	Discord *discordgo.Session
)

// Create a new Discord session using the provided bot token.
func DiscordSetup() error {
	discordToken := viper.GetString("DISCORD_TOKEN")

	discord, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		return err
	}

	Discord = discord

	return nil
}

// DiscordOpen opens the Discord session.
func DiscordOpen() error {
	err := Discord.Open()
	if err != nil {
		return err
	}

	return nil
}

// DiscordClose closes the Discord session.
func DiscordClose() {
	Discord.Close()
}
