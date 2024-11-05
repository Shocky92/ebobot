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
	componentHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"badumtss": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "https://cdn.discordapp.com/attachments/209759858856165377/394033551294726154/7445930.jpg",
				},
			})
		},
		"feron": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "http://image.prntscr.com/image/2bf970c168d14558a5a5c94f7fbb2c36.png",
				},
			})
		},
		"goblin": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "https://youtu.be/TfV0QGthMmc",
				},
			})
		},
		"happyment": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "https://pp.userapi.com/c830400/v830400244/18f4d1/UszOuoziVEc.jpg",
				},
			})
		},
		"kertcoin": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "https://i.imgur.com/po5nbTF.jpg",
				},
			})
		},
		"omegalul": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "https://youtu.be/5ejsttVNPLA",
				},
			})
		},
		"pidaras": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "https://media.discordapp.net/attachments/209759858856165377/642729886728847370/unknown.png?ex=672b7ada&is=672a295a&hm=d25106434000f23f88a322820270f6215b87552695f17268a49848ae59ba1f3d&=&format=webp&quality=lossless",
				},
			})
		},
		"pizda": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "https://youtu.be/a3SarF-TOMc",
				},
			})
		},
		"who": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "https://media.discordapp.net/attachments/209759858856165377/444549331199066122/who.PNG",
				},
			})
		},
	}

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
		{
			Name:        "custom",
			Description: "custom commands",
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
		"custom": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "custom command",
					Flags:   discordgo.MessageFlagsEphemeral,
					Components: []discordgo.MessageComponent{
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.Button{
									Label:    "badumtss",
									Style:    discordgo.SuccessButton,
									Disabled: false,
									CustomID: "badumtss",
								},
								discordgo.Button{
									Label:    "feron",
									Style:    discordgo.SuccessButton,
									Disabled: false,
									CustomID: "feron",
								},
								discordgo.Button{
									Label:    "goblin",
									Style:    discordgo.SuccessButton,
									Disabled: false,
									CustomID: "goblin",
								},
								discordgo.Button{
									Label:    "happyment",
									Style:    discordgo.SuccessButton,
									Disabled: false,
									CustomID: "happyment",
								},
								discordgo.Button{
									Label:    "pizda",
									Style:    discordgo.SuccessButton,
									Disabled: false,
									CustomID: "pizda",
								},
							},
						},
						discordgo.ActionsRow{
							Components: []discordgo.MessageComponent{
								discordgo.Button{
									Label:    "kertcoin",
									Style:    discordgo.SuccessButton,
									Disabled: false,
									CustomID: "kertcoin",
								},
								discordgo.Button{
									Label:    "omegalul",
									Style:    discordgo.SuccessButton,
									Disabled: false,
									CustomID: "omegalul",
								},
								discordgo.Button{
									Label:    "pidaras",
									Style:    discordgo.SuccessButton,
									Disabled: false,
									CustomID: "pidaras",
								},
								discordgo.Button{
									Label:    "who",
									Style:    discordgo.SuccessButton,
									Disabled: false,
									CustomID: "who",
								},
							},
						},
					},
				},
			})
		},
	}
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
