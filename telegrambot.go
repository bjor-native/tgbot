package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	getInfoCovid "ace-h/tgbot/api/covidSummaryAPI"
	getJoke "ace-h/tgbot/api/jokesAPI"
	db "ace-h/tgbot/db"
)

var (
	buttonTD = []tgbotapi.KeyboardButton{tgbotapi.KeyboardButton{Text: "Total Death COVID-19"}}
	buttonTC = []tgbotapi.KeyboardButton{tgbotapi.KeyboardButton{Text: "Total Confirmed COVID-19"}}
	buttonNC = []tgbotapi.KeyboardButton{tgbotapi.KeyboardButton{Text: "New confirmed COVID-19"}}
	buttonND = []tgbotapi.KeyboardButton{tgbotapi.KeyboardButton{Text: "New Deaths COVID-19"}}
	buttonNR = []tgbotapi.KeyboardButton{tgbotapi.KeyboardButton{Text: "New recovered COVID-19"}}
	buttonTR = []tgbotapi.KeyboardButton{tgbotapi.KeyboardButton{Text: "Total recovered COVID-19"}}
)

func main() {
	tgbot()
}

func tgbot() {
	bot, err := tgbotapi.NewBotAPI("813814117:AAEy5T8hws-wU86USfOQOcHQ_kOZFu8-x68")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		var message tgbotapi.MessageConfig
		log.Println("received text: ", update.Message.Text)

		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		switch update.Message.Text {
		case "Hello", "hello", "Привет", "привет":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, db.HelloWords())
		case "Иди нахуй", "иди нахуй", "Иди на хуй", "иди на хуй", "Пидарас", "пидарас":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, db.DirtyWords())
		case "Расскажи о себе", "расскажи о себе":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, db.AboutBot())
		case "Расскажи что-нибудь", "расскажи что-нибудь", "Расскажи что нибудь", "расскажи что нибудь", "Расскажи чтонибудь", "расскажи чтонибудь":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, db.HistoryWords())

		case "Хочу шутку", "хочу шутку":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, getJoke.GetJoke())

		case "Total Death COVID-19":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, "Всего смертей в мире: "+getInfoCovid.TotalDeath())
		case "Total Confirmed COVID-19":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, "Всего подтвержденных заражений в мире: "+getInfoCovid.TotalConfirmed())
		case "New confirmed COVID-19":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, "Новых заражений в мире: "+getInfoCovid.NewConfirmed())
		case "New Deaths COVID-19":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, "Новых смертей в мире: "+getInfoCovid.NewDeaths())
		case "New recovered COVID-19":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, "Вылечилось за сутки в мире: "+getInfoCovid.NewRecovered())
		case "Total recovered COVID-19":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, "Всего вылечилось в мире: "+getInfoCovid.TotalRecovered())

		default:
			message = tgbotapi.NewMessage(update.Message.Chat.ID, `Я не понимаю что ты хочешь :(`+"\n"+`Напиши @boot_fail`)
		}

		message.ReplyToMessageID = update.Message.MessageID
		message.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttonTD, buttonTC, buttonTR, buttonND, buttonNC, buttonNR)

		bot.Send(message)
	}
}
