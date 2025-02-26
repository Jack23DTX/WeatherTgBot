package main

import (
	"ProjectBot1/weather"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	api, exists := os.LookupEnv("TELEGRAM_TOKEN")
	if !exists {
		log.Fatal("TELEGRAM_TOKEN environment variable not set")
	}

	weatherApi, exists := os.LookupEnv("WEATHER_API")
	if !exists {
		log.Fatal("WEATHER_API environment variable not set")
	}

	tgBot, err := tgbotapi.NewBotAPI(api)
	if err != nil {
		fmt.Printf("Ошибка запуска бота: %s\n", err)
	}
	log.Printf("Authorized on account %s", tgBot.Self.UserName)

	tgBot.Debug = false

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := tgBot.GetUpdatesChan(u)
	for {
		for update := range updates {
			if update.Message != nil {
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, weather.Get(weatherApi, update.Message.Text))

				tgBot.Send(msg)
			}
		}
	}
}
