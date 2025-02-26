package main

import (
	"ProjectBot1/weather"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
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

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			str := weather.Get(weatherApi, update.Message.Text)
			ans := fmt.Sprintf(fmt.Sprintf("%s, %s\nДата и время: %s\nТемпература: %.1f°C\nОщущается как: %.1f°C,\nПогодные условия: %s\n",
				str.Location.Name, str.Location.Country, str.Location.Localtime, str.Current.TempCelsius, str.Current.FeelsLikeCelsius, str.Current.Condition.Text))

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, ans)

			tgBot.Send(msg)
		}
	}
}
