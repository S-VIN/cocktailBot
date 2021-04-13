package main

import (
	//"fmt"
	"math/rand"
	"strconv"
	//"strings"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var telegram Telegram

var replyKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("СЛУЧАЙНЫЙ АНЕК"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("СЛУЧАЙНЫЙ СМЕШНОЙ АНЕК"),
		tgbotapi.NewKeyboardButton("СЛУЧАЙНЫЙ НЕСМЕШНОЙ АНЕК"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("СЛУЧАЙНЫЙ ИЗБРАННЫЙ АНЕК"),
		tgbotapi.NewKeyboardButton("СПИСОК ИЗБРАННЫХ АНЕКОВ"),
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
	msg := tgbotapi.NewMessage(chatID, "Чтобы было проще хихикать, пользуйся клавиатурой.")
	msg.ReplyMarkup = replyKeyboard
	_, err := t.bot.Send(msg)
	return err
}

func (t Telegram) SendAnek(chatID int64, id int) error {
	if id < 0 || id > len(database.arrayOfAneks) {
		return nil
	}

	var likesKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(strconv.Itoa(database.arrayOfAneks[id].GetLikes())+" 👍🏻", "l"+strconv.Itoa(id)),
			tgbotapi.NewInlineKeyboardButtonData(strconv.Itoa(database.arrayOfAneks[id].GetDislikes())+" 👎🏾", "d"+strconv.Itoa(id)),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("🤎", "f"+strconv.Itoa(id)),
		),
	)

	msg := tgbotapi.NewMessage(chatID, database.GetAnekText(id))
	msg.ReplyMarkup = likesKeyboard
	_, err := t.bot.Send(msg)
	return err
}

func (t *Telegram) GetResponseFromInline(chatID int64, input string, callbackQuerryID string) {
	temp, _ := strconv.Atoi(input[1:len(input)])
	switch os := input[0]; os {
	case 'l':
		if !database.IsLike(chatID, temp) {
			database.Like(chatID, temp)
		} else {
			t.bot.AnswerCallbackQuery(tgbotapi.NewCallback(callbackQuerryID, "Ну ты и шалун! Любишь шалить! Лайк то ты уже поставил."))
		}
	case 'd':
		if !database.IsDislike(chatID, temp) {
			database.Dislike(chatID, temp)
		} else {
			t.bot.AnswerCallbackQuery(tgbotapi.NewCallback(callbackQuerryID, "Ух, какой ты злой. Проверь сахар. Дизлайк уже стоял."))
		}
	case 'f':
		if !database.IsFavourite(chatID, temp) {
			database.Favourite(chatID, temp)
		} else {
			t.bot.AnswerCallbackQuery(tgbotapi.NewCallback(callbackQuerryID, "Анек не смешной, а ты его второй раз в избранное добавляешь."))
		}
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

	case "СЛУЧАЙНЫЙ АНЕК":
		t.SendAnek(input.Chat.ID, rand.Intn(anekQuantity))

	case "СЛУЧАЙНЫЙ СМЕШНОЙ АНЕК":
		textOfAnek, index := database.GetRandomLikedAnek(input.Chat.ID)
		if textOfAnek == "" {
			t.SendMessage(input.Chat.ID, "Смешных анеков нет. Можешь посмотреть в зеркало.")
		} else {
			t.SendAnek(input.Chat.ID, index)
		}

	case "СЛУЧАЙНЫЙ НЕСМЕШНОЙ АНЕК":
		textOfAnek, index := database.GetRandomDislikedAnek(input.Chat.ID)
		if textOfAnek == "" {
			t.SendMessage(input.Chat.ID, "Несмешных анеков нет. Смейся, любитель похохотать.")
		} else {
			t.SendAnek(input.Chat.ID, index)
		}

	case "СЛУЧАЙНЫЙ ИЗБРАННЫЙ АНЕК":
		textOfAnek, index := database.GetRandomFavouriteAnek(input.Chat.ID)
		if textOfAnek == "" {
			t.SendMessage(input.Chat.ID, "Ты ничего не любишь. Твоё сердце пусто. Обычно такие люди умирают в одиночестве.")
		} else {
			t.SendAnek(input.Chat.ID, index)
		}

	case "СПИСОК ИЗБРАННЫХ АНЕКОВ":
		temp, _ := database.GetRandomFavouriteAnek(input.Chat.ID)
		if temp == "" {
			t.SendMessage(input.Chat.ID, "А вот <3 тебе, а не список.")
		}
		t.SendMessage(input.Chat.ID, database.GetStringOfFavourites(input.Chat.ID))

	default:
		i, err := strconv.Atoi(input.Text)
		if err == nil && i >= 0 && i < anekQuantity {
			t.SendAnek(input.Chat.ID, i)
		} else {
			t.SendMessage(input.Chat.ID, "Ты что, дурачок? Нажимай на кнопки, либо пиши число. Число должно быть правильным, а не как обычно.")
		}
	}
}