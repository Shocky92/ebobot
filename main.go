package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	// Читаем токен дискорд бота из env файла
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Не удалось загрузить .env файл")
		return
	}

	// Создаем новую сессию DiscordGo
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		fmt.Println("Ошибка при создании сессии DiscordGo: ", err)
		return
	}

	// Обработчик события готовности бота
	dg.AddHandlerOnce(onReady)

	// Обработчики команд бота
	dg.AddHandler(messageCreate)

	// Открываем сессию
	err = dg.Open()
	if err != nil {
		fmt.Println("Ошибка при открытии сессии: ", err)
		return
	}

	// Ожидаем сигналы останова (SIGINT/SIGTERM)
	fmt.Println("Бот запущен. Нажмите CTRL-C для завершения.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Закрываем сессию
	dg.Close()
}

func onReady(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("Бот готов к работе!")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Игнорируем сообщения от ботов и себя
	if m.Author.Bot || m.Author.ID == s.State.User.ID {
		return
	}

	// Отвечаем на сообщение
	s.ChannelMessageSend(m.ChannelID, "Привет, я бот!")
}
