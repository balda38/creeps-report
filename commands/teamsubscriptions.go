package commands

import (
	"context"
	"strings"

	"github.com/balda38/creeps-report/database"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	dbModels "github.com/balda38/creeps-report/database/models"
)

type TeamSubscriptionsCommand struct{}

func (TeamSubscriptionsCommand) GetName() string {
	return "/team_subscriptions"
}

func (TeamSubscriptionsCommand) GetDescription() string {
	return "List team subscriptions"
}

func (TeamSubscriptionsCommand) CommandMatchType() bot.MatchType {
	return bot.MatchTypeExact
}

func (TeamSubscriptionsCommand) Handler(ctx context.Context, botInstance *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID

	var teamSubscriptions []dbModels.Subscription
	teamSubscriptionsResult := database.DB.Where("chat_id = ?", chatID).
		Joins("Team").
		Order("Team.label").
		Find(&teamSubscriptions)

	if teamSubscriptionsResult.RowsAffected == 0 || teamSubscriptionsResult.Error != nil {
		botInstance.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "🤷‍♂️ You are not subscribed to any teams.",
		})
	} else {
		var teamList []string
		for _, subscription := range teamSubscriptions {
			teamList = append(teamList, subscription.Team.Label)
		}
		botInstance.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "📌 You are subscribed to the following teams:\n• " + strings.Join(teamList, "\n• "),
		})
	}
}
