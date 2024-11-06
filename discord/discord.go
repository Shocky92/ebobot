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
	componentHandlers = MakeComponentHandlers()
	commands          = MakeCommandList()
	commandHandlers   = MakeCommandHandlers()
)

func RegisterHandlers(s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
				h(s, i)
			}
		case discordgo.InteractionMessageComponent:
			if h, ok := componentHandlers[i.MessageComponentData().CustomID]; ok {
				h(s, i)
			}
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
		log.Printf("Registered command: %v", v.Name)
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
		log.Printf("Removed command: %v", v.Name)
	}
}
