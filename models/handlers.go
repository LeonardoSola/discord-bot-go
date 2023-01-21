package models

import "github.com/bwmarrin/discordgo"

type Context struct {
	Session *discordgo.Session
	Message *discordgo.MessageCreate
	Data
}

type Data struct {
	Command   string
	Arguments []string
	Argument  string
}

type Command struct {
	Name        string
	Description string
	Aliases     []string
	Handler     func(ctx Context)
}

func (ctx *Context) GetMessageReference() *discordgo.MessageReference {
	return &discordgo.MessageReference{
		MessageID: ctx.Message.ID,
		ChannelID: ctx.Message.ChannelID,
		GuildID:   ctx.Message.GuildID,
	}
}

func (ctx *Context) Reply(message string) (*discordgo.Message, error) {
	return ctx.Session.ChannelMessageSendReply(ctx.Message.ChannelID, message, ctx.GetMessageReference())
}
