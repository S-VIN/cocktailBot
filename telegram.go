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

var replyKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Lookup a random cocktail"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Search by ingredient"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Get like list"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Search ingredient by name"),
		tgbotapi.NewKeyboardButton("Search cocktail by name"),
	),
)

type Telegram struct {
	bot       *tgbotapi.BotAPI
	botConfig tgbotapi.UpdateConfig
}

func (t *Telegram) CreateBot() (err error) {
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

func (t Telegram) SendCocktail(chatID int64, cocktail Cocktail) error {

	var shortCocktailKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("details", "details"),
			tgbotapi.NewInlineKeyboardButtonData("🤎", "liked"),
		),
	)

	var temp = "*" + cocktail.StrDrink + "* " + "("+ cocktail.StrGlass + ")" + "\n"
	fmt.Println(cocktail)
	for i := 0; i < 15; i++ {
		if(cocktail.Ingridients[i]!=""){
			temp += "✅"
		}
		temp += string(cocktail.Ingridients[i])
		temp += "\n"
	}
	msg := tgbotapi.NewMessage(chatID, temp)
	msg.ReplyMarkup = shortCocktailKeyboard
	_, err := t.bot.Send(msg)
	return err
}

func (t Telegram) SendDetailedCocktail(chatID int64, cocktail Cocktail) error{
	var shortCocktailKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🤎", "liked"),
		),
	)
	
	var textAnswer = "*" + cocktail.StrDrink + "* " + "\n" + "\n"
	textAnswer += "glass:\n" + cocktail.StrGlass + "\n" + "\n"
	textAnswer += "instruction:\n" + cocktail.StrInstructions + "\n" + "\n"
	textAnswer += "ingridients:" + "\n"
	for i := 0; i < 15; i++ {
		if(cocktail.Ingridients[i]==""){
			break
		}
		textAnswer += "✅"
		textAnswer += string(cocktail.Ingridients[i]) 
		textAnswer += " - " + string(cocktail.Meashures[i])
		textAnswer += "\n"
	}


	msg := tgbotapi.NewPhotoUpload(chatID, nil)
	msg.FileID = cocktail.StrDrinkThumb
	msg.UseExisting = true

	msg.Caption = textAnswer
	
	msg.ReplyMarkup = shortCocktailKeyboard

	_, err := t.bot.Send(msg)
	
	return err
}

func (t *Telegram) GetResponseFromInline(chatID int64, input string, callbackQuerryID string) {
	if input == "details" {
		t.bot.AnswerCallbackQuery(tgbotapi.NewCallback(callbackQuerryID, "dekjhgj"))
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
			t.bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, "Молодец! Твой палец записан, куда надо."))
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
		t.SendDetailedCocktail(input.Chat.ID, temp)

	default:
		t.SendMessage(input.Chat.ID, "Unknown command")

	}
}
