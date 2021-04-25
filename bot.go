package main

import (
	"errors"
	"strconv"
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

	t.bot, err = tgbotapi.NewBotAPI("1356963581:AAGPlUyAkofdhcehODZ-jvIv9Qu9T196pRQ")
	if err != nil {
		return err
	}
	t.botConfig = tgbotapi.NewUpdate(0)
	t.botConfig.Timeout = 60

	return nil
}

func (t *Telegram) GetResponseFromInline(chatID int64, input string, callbackQuerryID string) error {
	if input == "" {
		return errors.New("empty inline input")
	}
	switch input[0] {
	case 'l':
		if res, err := database.IsLike(chatID, input[1:]); !res {
			if err != nil {
				return err
			}
			err = database.Like(chatID, input[1:])
			if err != nil {
				return err
			}
			t.bot.AnswerCallbackQuery(tgbotapi.NewCallback(callbackQuerryID, "like added"))
		} else {
			t.bot.AnswerCallbackQuery(tgbotapi.NewCallback(callbackQuerryID, "like was added before"))
		}
	case 'd':
		cocktail, err := lookUpCocktailId(input[1:])
		if err != nil {
			return err
		}
		err = t.SendDetailedCocktail(chatID, cocktail)
		if err != nil {
			return err
		}
	case 's':
		err := t.SendMessage(chatID, "type a number of cocktail")
		if err != nil {
			return err
		}
		clientStatus.status[chatID] = WFLIST
	default:
		return errors.New("wrong inline input, input = " + input)
	}
	return nil
}

func (t Telegram) CheckUpdates() error {
	updates, err := t.bot.GetUpdatesChan(t.botConfig)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.CallbackQuery != nil {
			err := t.GetResponseFromInline(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data, update.CallbackQuery.ID)
			if err != nil {
				return err
			}
			_, err = t.bot.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))
			if err != nil {
				return err
			}
		}

		if update.Message == nil {
			continue
		}
		t.CreateAnswer(*update.Message)
	}
	return nil
}

func (t Telegram) CreateAnswer(input tgbotapi.Message) {
	logInf.Println(strconv.FormatInt(input.Chat.ID, 10) + ": CreateAnswer, input: " + input.Text)
	switch input.Text {
	case "/start":
		err := t.SendReplyKeyboard(input.Chat.ID)
		if err != nil {
			logErr.Println(err.Error())
		}

	case "Lookup a random cocktail":
		temp, err := getRandomCocktail()
		if err != nil {
			logErr.Println(err.Error())
		}
		err = t.SendCocktail(input.Chat.ID, temp)
		if err != nil {
			logErr.Println(err.Error())
		}

	case "Search by ingredient":
		err := t.SendMessage(input.Chat.ID, "Type the ingredient")
		if err != nil {
			logErr.Println(err.Error())
		}
		clientStatus.status[input.Chat.ID] = WFINGR

	case "Search by name":
		err := t.SendMessage(input.Chat.ID, "Type the name")
		if err != nil {
			logErr.Println(err.Error())
		}
		clientStatus.status[input.Chat.ID] = WFNAME

	case "Get like list":
		res, err := database.GetRangeOfLikes(input.Chat.ID)
		if err != nil {
			logErr.Println(err.Error())
		}
		err = t.SendRangeOfCocktails(res, input.Chat.ID)
		if err != nil {
			logErr.Println(err.Error())
		}

	default:
		var err error
		switch clientStatus.status[input.Chat.ID] {
		case WFINGR:
			err = t.answerIngredient(input.Chat.ID, input.Text)

		case WFNAME:
			err = t.answerName(input.Chat.ID, input.Text)

		case WFLIST:
			err = t.answerList(input.Chat.ID, input.Text)
		default:
			err = t.SendMessage(input.Chat.ID, "unknown command")
		}
		if err != nil{
			logErr.Println(err.Error())
		}
	}
}
