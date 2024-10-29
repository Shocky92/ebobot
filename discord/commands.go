package discord

import (
	"flag"
	"log"

	"github.com/bwmarrin/discordgo"
)

// Bot parameters
var (
	GuildID = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
)

// Commands
var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Test bot",
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"ping": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "pong",
				},
			})
		},
	}
)

func RegisterHandlers(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})
}

func RegisterCommands(s *discordgo.Session) []*discordgo.ApplicationCommand {
	log.Println("Register commands...")

	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, *GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	return registeredCommands
}

func RemoveCommands(s *discordgo.Session, registeredCommands []*discordgo.ApplicationCommand) {
	log.Println("Removing commands ...")

	for _, v := range registeredCommands {
		err := s.ApplicationCommandDelete(s.State.User.ID, *GuildID, v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}
}
