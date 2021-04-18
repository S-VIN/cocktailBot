package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var telegram Telegram
var clientStatus ClientStatus
var database Database

var replyKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Lookup a random cocktail"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Get like list"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Search by name"),
		tgbotapi.NewKeyboardButton("Search by ingredient"),
	),
)

type Telegram struct {
	bot       *tgbotapi.BotAPI
	botConfig tgbotapi.UpdateConfig
}

func (t *Telegram) CreateBot() (err error) {
	clientStatus.status = make(map[int64]int)
	database = *NewDatabase()
	t.bot, err = tgbotapi.NewBotAPI("1356963581:AAGPlUyAkofdhcehODZ-jvIv9Qu9T196pRQ")
	if err != nil {
		return err
	}
	t.botConfig = tgbotapi.NewUpdate(0)
	t.botConfig.Timeout = 60
	return nil
}

func (t Telegram) SendMessage(chatID int64, input string) error {
	_, err := t.bot.Send(tgbotapi.NewMessage(chatID, input))
	return err
}

func (t Telegram) SendReplyKeyboard(chatID int64) error {
	msg := tgbotapi.NewMessage(chatID, "Use keyboard for commands.")
	msg.ReplyMarkup = replyKeyboard
	_, err := t.bot.Send(msg)
	return err
}

func (t *Telegram) GetResponseFromInline(chatID int64, input string, callbackQuerryID string) {
	switch input[0]{
	case 'l':
		if database.isLike(chatID, input[1:]) {
			database.like(chatID, input[1:])
			t.bot.AnswerCallbackQuery(tgbotapi.NewCallback(callbackQuerryID, "like added"))
		} else {
			t.bot.AnswerCallbackQuery(tgbotapi.NewCallback(callbackQuerryID, "like was added before"))
		}
	case 'd':
		cocktail, _ := lookUpCocktailId(input[1:])
		SendDetailedCocktail(chatID, cocktail, t.bot)
	case 's':
		t.SendMessage(chatID, "type a number of cocktail")
		clientStatus.status[chatID] = WFLLIST
}
}

func (t Telegram) CheckUpdates() error {
	updates, err := t.bot.GetUpdatesChan(t.botConfig)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.CallbackQuery != nil {
			t.GetResponseFromInline(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data, update.CallbackQuery.ID)
			t.bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))
		}

		if update.Message == nil {
			continue
		}
		t.CreateAnswer(*update.Message)
	}
	return nil
}

func (t Telegram) CreateAnswer(input tgbotapi.Message) {
	switch input.Text {
	case "/start":
		t.SendReplyKeyboard(input.Chat.ID)

	case "Lookup a random cocktail":
		temp, _ := getRandomCocktail()
		SendCocktail(input.Chat.ID, temp, t.bot)

	case "Search by ingredient":
		t.SendMessage(input.Chat.ID, "Type the ingredient")
		clientStatus.status[input.Chat.ID] = WFINGR

	case "Search by name":
		t.SendMessage(input.Chat.ID, "Type the name")
		clientStatus.status[input.Chat.ID] = WFNAME

	case "Get like list":
		fmt.Println(database.getRangeOfLikes(input.Chat.ID))
		SendRangeOfCocktails(database.getRangeOfLikes(input.Chat.ID), input.Chat.ID, t.bot)

	default:
		if clientStatus.status[input.Chat.ID] == WFINGR {
			cocktail, _ := searchByIngredient(input.Text)
			SendCocktail(input.Chat.ID, cocktail.Drinks[0], t.bot)
			clientStatus.status[input.Chat.ID] = DONE
		}

		if clientStatus.status[input.Chat.ID] == WFNAME {
			cocktail, _ := searchByIngredient(input.Text)
			SendCocktail(input.Chat.ID, cocktail.Drinks[0], t.bot)
			clientStatus.status[input.Chat.ID] = DONE
		}

		if clientStatus.status[input.Chat.ID] == WFLLIST {
			//cocktail, _ := lookUpFullCocktailDetailById(input.Text)
		}

	}
}
