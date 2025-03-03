package core

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Command interface {
	GetName() string
	GetDescription() string
	CommandMatchType() bot.MatchType
	Handler(
		ctx context.Context,
		b *bot.Bot,
		update *models.Update,
	) // TODO: return type?
}
