package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendCocktail(chatID int64, cocktail Cocktail, bot *tgbotapi.BotAPI) error {

	var shortCocktailKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("details", cocktail.IdDrink),
			tgbotapi.NewInlineKeyboardButtonData("🤎", "liked"),
		),
	)

	var temp = "*" + cocktail.StrDrink + "* " + "(" + cocktail.StrGlass + ")" + "\n"
	fmt.Println(cocktail)
	for i := 0; i < 15; i++ {
		if cocktail.Ingridients[i] != "" {
			temp += "✅"
		}
		temp += string(cocktail.Ingridients[i])
		temp += "\n"
	}
	msg := tgbotapi.NewMessage(chatID, temp)
	msg.ReplyMarkup = shortCocktailKeyboard
	_, err := bot.Send(msg)
	return err
}

func SendDetailedCocktail(chatID int64, cocktail Cocktail, bot *tgbotapi.BotAPI) error {
	var shortCocktailKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🤎", "l" + cocktail.IdDrink),
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

	_, err := bot.Send(msg)

	return err
}

func SendRangeOfCocktails(inputIDS *[]string, chatID int64, bot *tgbotapi.BotAPI) error{
	var messageString string
	for iter, value := range(*inputIDS){
		cocktail, _ := lookUpCocktailId(value)
		messageString += string(iter) + ") "
		messageString += cocktail.StrDrink + "\n"
	}
	msg := tgbotapi.NewMessage(chatID, messageString)
	_, err := bot.Send(msg)
	return err
}