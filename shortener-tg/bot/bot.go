package bot

import (
	"context"
	"log"
	"log/slog"
	grpc "shortener-tg/grpc/clients/api"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func Start(BotToken string, client *grpc.Client, ctx context.Context, timeout time.Duration) {
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}

	for update := range updates {
		if update.Message != nil {
			slog.Info("message received", slog.String("message", update.Message.Text))

			shortenedURL, err := client.ShortenURL(ctx, update.Message.Text)
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			}

			slog.Info("Request recieved: ", slog.String("url", shortenedURL))

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, shortenedURL)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}
}
