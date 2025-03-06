package commands

import (
	"context"
	"os"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type HelpCommand struct{}

func (HelpCommand) GetName() string {
	return "/help"
}

func (HelpCommand) GetDescription() string {
	return "Get help"
}

func (HelpCommand) CommandMatchType() bot.MatchType {
	return bot.MatchTypeExact
}

func (HelpCommand) Handler(ctx context.Context, botInstance *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	botAdmin := os.Getenv("TELEGRAM_BOT_ADMIN")
	if botAdmin == "" {
		return
	}
	botInstance.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   "Contact @" + botAdmin + " for help!",
	})
}
