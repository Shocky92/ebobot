package discord

import (
	"context"
	"ebobot/config"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
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
			ctx := context.Background()
			client, err := newYandexGPTClient(config.YandexAPIKey, config.YandexCatalogID)
			if err != nil {
				log.Println("Error creating Yandex GPT client:", err)
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Ошибка при создании клиента Yandex GPT.",
					},
				})
				return
			}

			query := i.ApplicationCommandData().Options[0].StringValue()
			request := newYandexGPTRequest(config.YandexCatalogID, query)

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags:   discordgo.MessageFlagsLoading,
					Content: "Обработка ответа, ожидайте...",
				},
			})

			response, err := client.CreateRequest(ctx, request)
			if err != nil {
				log.Println("Request error:", err)
				message := fmt.Sprintf("Request error: %v", err)
				s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
					Content: &message,
				})
				return
			}

			content := response.Result.Alternatives[0].Message.Text
			s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Content: &content,
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
