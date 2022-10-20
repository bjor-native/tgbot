package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	getInfoCovid "tgbot/api/covidSummaryAPI"
	getJoke "tgbot/api/jokesAPI"
	config "tgbot/utils"
)

var (
	buttonTD = []tgbotapi.KeyboardButton{{Text: "Total Death COVID-19"}}
	buttonTC = []tgbotapi.KeyboardButton{{Text: "Total Confirmed COVID-19"}}
	buttonNC = []tgbotapi.KeyboardButton{{Text: "New confirmed COVID-19"}}
	buttonND = []tgbotapi.KeyboardButton{{Text: "New Deaths COVID-19"}}
	buttonNR = []tgbotapi.KeyboardButton{{Text: "New recovered COVID-19"}}
	buttonTR = []tgbotapi.KeyboardButton{{Text: "Total recovered COVID-19"}}
)

func main() {
	tgbot()
}

func tgbot() {
	conf := config.GetConfig()

	bot, err := tgbotapi.NewBotAPI(conf.ApiKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		var message tgbotapi.MessageConfig
		log.Println("received text: ", update.Message.Text)

		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		switch update.Message.Text {
		case "Хочу шутку", "хочу шутку":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, getJoke.GetJoke())

		case "Total Death COVID-19":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, "Всего смертей в мире: "+getInfoCovid.GetCovidStatistic("totalDeaths"))
		case "Total Confirmed COVID-19":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, "Всего подтвержденных заражений в мире: "+getInfoCovid.GetCovidStatistic("totalConfirmed"))
		case "New confirmed COVID-19":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, "Новых заражений в мире: "+getInfoCovid.GetCovidStatistic("newConfirmed"))
		case "New Deaths COVID-19":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, "Новых смертей в мире: "+getInfoCovid.GetCovidStatistic("newDeaths"))
		case "New recovered COVID-19":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, "Вылечилось за сутки в мире: "+getInfoCovid.GetCovidStatistic("newRecovered"))
		case "Total recovered COVID-19":
			message = tgbotapi.NewMessage(update.Message.Chat.ID, "Всего вылечилось в мире: "+getInfoCovid.GetCovidStatistic("totalRecovered"))

		default:
			message = tgbotapi.NewMessage(update.Message.Chat.ID, `Я не понимаю что ты хочешь :(`+"\n"+`Иди нахуй`)
		}

		message.ReplyToMessageID = update.Message.MessageID
		message.ReplyMarkup = tgbotapi.NewReplyKeyboard(buttonTD, buttonTC, buttonTR, buttonND, buttonNC, buttonNR)

		bot.Send(message)
	}
}
