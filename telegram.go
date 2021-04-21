package main

import (
	//"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

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
	_, err := t.bot.Send(msg)
	return err
}

func (t Telegram) SendDetailedCocktail(chatID int64, cocktail Cocktail) error {
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
	_, err := t.bot.Send(msg)

	return err
}

func (t Telegram) SendRangeOfCocktails(inputIDS []string, chatID int64) error {
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
	_, err := t.bot.Send(msg)
	clientStatus.shownCocktails[chatID] = inputIDS
	return err
}

func (t Telegram) answerIngredient(chatID int64, textFromKeyboard string) {
	cocktails, _ := searchByIngredient(textFromKeyboard)
	var cocktailIDS []string
	for _, value := range cocktails {
		cocktailIDS = append(cocktailIDS, value.IdDrink)
	}
	t.SendRangeOfCocktails(cocktailIDS, chatID)
	clientStatus.status[chatID] = WFLIST
}

func (t Telegram) answerName(chatID int64, textFromKeyboard string) {
	cocktails, _ := searchCocktailByName(textFromKeyboard)
	var cocktailIDS []string
	for _, value := range cocktails {
		cocktailIDS = append(cocktailIDS, value.IdDrink)
	}
	t.SendRangeOfCocktails(cocktailIDS, chatID)
	clientStatus.status[chatID] = WFLIST
}

func (t Telegram) answerList(chatID int64, numberFromKeyboard string) {
	index, _ := strconv.Atoi(numberFromKeyboard)
	cocktailID := clientStatus.shownCocktails[chatID][index]
	cocktail, _ := lookUpFullCocktailDetailById(cocktailID)
	t.SendDetailedCocktail(chatID, cocktail)
	clientStatus.status[chatID] = DONE
}
