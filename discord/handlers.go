package discord

import (
	"context"
	"ebobot/config"
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/sheeiavellie/go-yandexgpt"
)

var firstRowButtons = buttons[:5]
var secondRowButtons = buttons[5:]

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
			ctx := context.Background()
			service, err := newYoutubeService(ctx, config.YoutubeApiKey)
			if err != nil {
				log.Fatalf("Error creating new YouTube service: %v", err)
			}

			query := i.ApplicationCommandData().Options[0].StringValue()
			response, err := searchVideo(service, query)
			if err != nil {
				log.Fatalf("Error making API call: %v", err)
			}

			if len(response.Items) == 0 {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Видео не найдено.",
					},
				})
				return
			}

			url := createURL(response.Items[0])
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: url,
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
						createButtonRow(firstRowButtons),
						createButtonRow(secondRowButtons),
					},
				},
			})
		},
	}
}
