package main

import(
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (t Telegram) SendCocktail(chatID int64, cocktail Cocktail) error {

	var shortCocktailKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("details", cocktail.IdDrink),
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