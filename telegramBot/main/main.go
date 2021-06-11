package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

var bot, err_ = tgbotapi.NewBotAPI("1648899629:AAEpJWaRSxBBL0IoW5yJQ4_0uoHr7Qg1cyg")

func main() {
	if err_ != nil {
		panic(err_)
	}
	bot.Debug = true
	upd := tgbotapi.NewUpdate(0)
	upd.Timeout = 60
	updates, err := bot.GetUpdatesChan(upd)
	if err != nil {
		panic(err)
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		err = GetRespond(update.Message.Text, update)
		if err != nil {
			panic(err)
		}
	}
}
