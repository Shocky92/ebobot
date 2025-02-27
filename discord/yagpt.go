package discord

import "github.com/sheeiavellie/go-yandexgpt"

func newYandexGPTClient(apiKey string, catalogID string) (*yandexgpt.YandexGPTClient, error) {
	client := yandexgpt.NewYandexGPTClientWithAPIKey(apiKey)
	return client, nil
}

func newYandexGPTRequest(catalogID string, query string) yandexgpt.YandexGPTRequest {
	return yandexgpt.YandexGPTRequest{
		ModelURI: yandexgpt.MakeModelURI(catalogID, yandexgpt.YandexGPT4Model),
		CompletionOptions: yandexgpt.YandexGPTCompletionOptions{
			Stream:      false,
			Temperature: 1.0,
			MaxTokens:   2000,
		},
		Messages: []yandexgpt.YandexGPTMessage{
			{
				Role: yandexgpt.YandexGPTMessageRoleUser,
				Text: query,
			},
		},
	}
}
