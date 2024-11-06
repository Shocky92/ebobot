package discord

import (
	"context"
	"ebobot/config"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sheeiavellie/go-yandexgpt"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func MakeCommandHandlers() map[string]func(*discordgo.Session, *discordgo.InteractionCreate) {
	return map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
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
					Temperature: 1.0,
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
}
