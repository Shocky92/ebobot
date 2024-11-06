package discord

import "github.com/bwmarrin/discordgo"

func MakeCommandList() []*discordgo.ApplicationCommand {
	return []*discordgo.ApplicationCommand{
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
}
