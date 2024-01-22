package main

import (
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"
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
	dg.AddHandler(onReceiveMessage)

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
	if m.ChannelID == os.Getenv("TEST_CHANNEL_ID") {
		s.ChannelMessageSend(m.ChannelID, "Привет, я пидорас!")
	}
}

func onReceiveMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Игнорируем сообщения от ботов и себя
	if m.Author.Bot || m.Author.ID == s.State.User.ID {
		return
	}
	fmt.Println(m.Content)
	// // Устанавливаем соединение с сервером
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://ebobot_db_1:27017"))
	// if err != nil {
	// 	fmt.Println("Не создать клиент базы данных: ", err)
	// }

	// // Подключаем клиента к серверу
	// ctx := context.Background()
	// err = client.Connect(ctx)
	// if err != nil {
	// 	fmt.Println("Не удалось присоединиться к базе данных: ", err)
	// }

	// // Определяем имя базы данных и коллекции
	// db := client.Database("ebobot")
	// col := db.Collection("sentences")

	// // Очищаем предложение
	// cleanSentence := sanitizeSentence(m.Content)

	// if cleanSentence == "" {
	// 	return
	// }

	// // Создаем новый документ, содержащий очищенное предложение
	// doc := bson.D{{Key: "sentence", Value: cleanSentence}}

	// // Вставляем документ в коллекцию
	// _, err = col.InsertOne(ctx, doc)
	// if err != nil {
	// 	fmt.Println("Не удалось сохранить запись в коллекцию: ", err)
	// }
}

func sanitizeSentence(sentence string) string {
	// Удаляем символы, отличные от букв и цифр
	regex := regexp.MustCompile(`<@?[:[a-zA-Z0-9]*]*>|![a-zA-Z]*|http\S+`)
	sentence = regex.ReplaceAllString(sentence, " ")

	// Приводим к нижнему регистру и удаляем лишние пробелы
	sentence = strings.ToLower(strings.TrimSpace(sentence))

	return sentence
}
