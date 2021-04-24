package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Telegram struct {
	bot       *tgbotapi.BotAPI
	botConfig tgbotapi.UpdateConfig
}

var telegram Telegram //struct for using telegram
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

func (t *Telegram) CreateBot() (err error) {

	clientStatus.status = make(map[int64]int)
	clientStatus.shownCocktails = make(map[int64][]string)
	//database = *NewFDatabase()
	database.Init()

	t.bot, err = tgbotapi.NewBotAPI("1356963581:AAGPlUyAkofdhcehODZ-jvIv9Qu9T196pRQ")
	if err != nil {
		return err
	}
	t.botConfig = tgbotapi.NewUpdate(0)
	t.botConfig.Timeout = 60

	return nil
}

func (t *Telegram) GetResponseFromInline(chatID int64, input string, callbackQuerryID string) {
	switch input[0] {
	case 'l':
		if res, _ := database.IsLike(chatID, input[1:]); !res {
			database.Like(chatID, input[1:])
			t.bot.AnswerCallbackQuery(tgbotapi.NewCallback(callbackQuerryID, "like added"))
		} else {
			t.bot.AnswerCallbackQuery(tgbotapi.NewCallback(callbackQuerryID, "like was added before"))
		}
	case 'd':
		cocktail, _ := lookUpCocktailId(input[1:])
		t.SendDetailedCocktail(chatID, cocktail)
	case 's':
		t.SendMessage(chatID, "type a number of cocktail")
		clientStatus.status[chatID] = WFLIST
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
		t.SendCocktail(input.Chat.ID, temp)

	case "Search by ingredient":
		t.SendMessage(input.Chat.ID, "Type the ingredient")
		clientStatus.status[input.Chat.ID] = WFINGR

	case "Search by name":
		t.SendMessage(input.Chat.ID, "Type the name")
		clientStatus.status[input.Chat.ID] = WFNAME

	case "Get like list":
		res, _ := database.GetRangeOfLikes(input.Chat.ID)
		t.SendRangeOfCocktails(res, input.Chat.ID)

	default:
		switch clientStatus.status[input.Chat.ID] {
		case WFINGR:
			t.answerIngredient(input.Chat.ID, input.Text)

		case WFNAME:
			t.answerName(input.Chat.ID, input.Text)

		case WFLIST:
			t.answerList(input.Chat.ID, input.Text)

		}

	}
}
