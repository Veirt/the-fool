package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/veirt/sena/types"
)

var ping = types.Command{
	Name:        "ping",
	Category:    "utility",
	Description: "Check if bot is online or not",
	Run: func(s *discordgo.Session, m *discordgo.MessageCreate) {
		s.ChannelMessageSendReply(m.ChannelID, "pong!", m.Reference())
	},
}

func init() {
	RegisterCommand(ping)
}
