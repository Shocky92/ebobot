package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	Token string
)

func SetConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Не удалось загрузить .env файл")
		return
	}

	Token = os.Getenv("DISCORD_TOKEN")
}
