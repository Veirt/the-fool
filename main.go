package main

import (
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/veirt/sena/commands"
	"github.com/veirt/sena/handlers"

	_ "github.com/joho/godotenv/autoload"
)

const PREFIX = "."

func main() {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	s.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.ChannelID == handlers.MistralChannelId {
			handlers.HandleMistralMessage(s, m)
			return
		}

		if !strings.HasPrefix(m.Content, PREFIX) {
			return
		}

		commands := commands.GetCommands()

		for _, cmd := range commands {
			if strings.HasPrefix(m.Content, PREFIX+cmd.Name) {
				cmd.Run(s, m)
			}
		}
	})

	err = s.Open()
	if err != nil {
		log.Fatal("error when opening session: ", err)
	}
	defer s.Close()

	// Run until code is terminated
	log.Println("Sena is online!")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
