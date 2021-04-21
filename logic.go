package main

import (
	//"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendCocktail(chatID int64, cocktail Cocktail, bot *tgbotapi.BotAPI) error {

	var shortCocktailKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("details", "d"+cocktail.IdDrink),
			tgbotapi.NewInlineKeyboardButtonData("🤎", "l"+cocktail.IdDrink),
		),
	)

	var temp = "*" + cocktail.StrDrink + "* " + "(" + cocktail.StrGlass + ")" + "\n"
	for i := 0; i < 15; i++ {
		if cocktail.Ingridients[i] != "" {
			temp += "✅"
		}
		temp += string(cocktail.Ingridients[i])
		temp += "\n"
	}
	msg := tgbotapi.NewMessage(chatID, temp)
	msg.ReplyMarkup = shortCocktailKeyboard
	msg.ParseMode = tgbotapi.ModeMarkdown
	_, err := bot.Send(msg)
	return err
}

func SendDetailedCocktail(chatID int64, cocktail Cocktail, bot *tgbotapi.BotAPI) error {
	var shortCocktailKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🤎", "l"+cocktail.IdDrink),
		),
	)

	var textAnswer = "*" + cocktail.StrDrink + "* " + "\n" + "\n"
	textAnswer += "glass:\n" + cocktail.StrGlass + "\n" + "\n"
	textAnswer += "instruction:\n" + cocktail.StrInstructions + "\n" + "\n"
	textAnswer += "ingridients:" + "\n"
	for i := 0; i < 15; i++ {
		if cocktail.Ingridients[i] == "" {
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
	msg.ParseMode = tgbotapi.ModeMarkdown
	_, err := bot.Send(msg)

	return err
}

func SendRangeOfCocktails(inputIDS []string, chatID int64, bot *tgbotapi.BotAPI) error {
	var messageString string
	for iter, value := range inputIDS {
		cocktail, _ := lookUpCocktailId(value)
		messageString += strconv.Itoa(iter) + ") "
		messageString += cocktail.StrDrink + "\n"
	}
	msg := tgbotapi.NewMessage(chatID, messageString)
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("show one", "s"),
		),
	)
	_, err := bot.Send(msg)
	clientStatus.shownCocktails[chatID] = inputIDS
	return err
}

