package main

import (
	//"fmt"
	//"math/rand"
	//"strconv"
	//"strings"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var telegram Telegram
var clientStatus ClientStatus

var replyKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Lookup a random cocktail"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Get like list"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Search by name"),
		tgbotapi.NewKeyboardButton("Search by ingridient"),
	),
)

type Telegram struct {
	bot       *tgbotapi.BotAPI
	botConfig tgbotapi.UpdateConfig
}

func (t *Telegram) CreateBot() (err error) {
	clientStatus.status = make(map[int64]int)
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
	if input == "liked" {
		t.bot.AnswerCallbackQuery(tgbotapi.NewCallback(callbackQuerryID, "dekjhgj"))
	} else {
		cocktail, _ := lookUpCocktailId(input)
		t.SendDetailedCocktail(chatID, cocktail)
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
		temp, err := getRandomCocktail()
		fmt.Println(err)
		t.SendCocktail(input.Chat.ID, temp)

	case "Search by ingredient":
		t.SendMessage(input.Chat.ID, "Type the ingridient")
		clientStatus.status[input.Chat.ID] = WFINGR
	
	case "Search by name":
		t.SendMessage(input.Chat.ID, "Type the name")
		clientStatus.status[input.Chat.ID] = WFNAME

	default:
		if(clientStatus.status[input.Chat.ID] == WFINGR){
			cocktail, _:= searchByIngredient(input.Text)
			t.SendCocktail(input.Chat.ID, cocktail.Drinks[0])
			clientStatus.status[input.Chat.ID] = DONE
		}

		if(clientStatus.status[input.Chat.ID] == WFNAME){
			cocktail, _:= searchByIngredient(input.Text)
			t.SendCocktail(input.Chat.ID, cocktail.Drinks[0])
			clientStatus.status[input.Chat.ID] = DONE
		}

	}
}
