package main

import (
	"ebobot/config"
	"ebobot/discord"
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

// Bot parameters
var (
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

var s *discordgo.Session

func init() {
	config.SetConfig()
	var err error
	s, err = discordgo.New("Bot " + config.Token)

	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	discord.RegisterHandlers(s)

	log.Printf("Initialized...")
}

func main() {
	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	err := s.Open()
	if err != nil {
		log.Fatalf("Cannot open session: %v", err)
	}

	registeredCommands := discord.RegisterCommands(s)

	defer s.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	if *RemoveCommands {
		discord.RemoveCommands(s, registeredCommands)
	}

	log.Println("Gracefully shutting down")
}
