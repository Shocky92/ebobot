package discord

import (
	"context"
	"ebobot/config"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sheeiavellie/go-yandexgpt"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
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
			Description: "Проверка бота",
		},
		{
			Name:        "yagpt",
			Description: "YandexGPT request",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "prompt",
					Description: "Запрос к нейросети",
					Required:    true,
				},
			},
		},
		{
			Name:        "youtube",
			Description: "Youtube video search",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "search",
					Description: "Поиск видео в Youtube",
					Required:    true,
				},
			},
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
		"yagpt": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			client := yandexgpt.NewYandexGPTClientWithAPIKey(config.YandexAPIKey)
			request := yandexgpt.YandexGPTRequest{
				ModelURI: yandexgpt.MakeModelURI(config.YandexCatalogID, yandexgpt.YandexGPT4ModelLite),
				CompletionOptions: yandexgpt.YandexGPTCompletionOptions{
					Stream:      false,
					Temperature: 0.7,
					MaxTokens:   2000,
				},
				Messages: []yandexgpt.YandexGPTMessage{
					{
						Role: yandexgpt.YandexGPTMessageRoleAssistant,
						Text: i.ApplicationCommandData().Options[0].StringValue(),
					},
				},
			}

			var content string
			response, err := client.CreateRequest(context.Background(), request)
			if err != nil {
				content = fmt.Sprintf("Request error: %v", err)
				log.Fatalf("Request error: %v", err)
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   discordgo.MessageFlagsLoading,
					Content: "Обработка ответа, ожидайте...",
				},
			})

			time.AfterFunc(time.Second, func() {
				content = response.Result.Alternatives[0].Message.Text
				_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
					Content: &content,
				})
			})
		},
		"youtube": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			service, err := youtube.NewService(context.Background(), option.WithAPIKey(config.YoutubeApiKey))

			if err != nil {
				log.Fatalf("Error creating new YouTube service: %v", err)
			}

			// Make the API call to YouTube.
			call := service.Search.
				List([]string{"id", "snippet"}).
				Q(i.ApplicationCommandData().Options[0].StringValue()).
				MaxResults(1)

			response, err := call.Do()
			if err != nil {
				log.Fatalf("Error making API call: $v", err)
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "https://www.youtube.com/watch?v=" + response.Items[0].Id.VideoId,
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
