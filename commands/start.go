package commands

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type StartCommand struct{}

func (StartCommand) GetName() string {
	return "/start"
}

func (StartCommand) GetDescription() string {
	return ""
}

func (StartCommand) CommandMatchType() bot.MatchType {
	return bot.MatchTypeExact
}

func (StartCommand) Handler(ctx context.Context, botInstance *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	botInstance.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   "Welcome to CreepsReport! Stay tuned for live Dota 2 match stats.",
	})
}
