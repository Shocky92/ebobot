package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	Token           string
	YandexAPIKey    string
	YandexCatalogID string
	YoutubeApiKey   string
)

func SetConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Не удалось загрузить .env файл")
		return
	}

	// Discord
	Token = os.Getenv("DISCORD_TOKEN")
	// Yandex
	YandexAPIKey = os.Getenv("YANDEX_API_KEY")
	YandexCatalogID = os.Getenv("YANDEX_CATALOG_ID")
	// Youtube
	YoutubeApiKey = os.Getenv("YOUTUBE_API_KEY")
}
