package commands

import (
	"context"
	"slices"
	"time"

	"github.com/balda38/creeps-report/database"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	dbModels "github.com/balda38/creeps-report/database/models"
)

const maxMessageLength = 4096
const messagesPrefix = "📌 You can subscribe or unsubscribe from the team using the commands below:\n"

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

	var teams []dbModels.Team
	teamsResult := database.DB.Where("is_active = ?", true).
		Order("label COLLATE NOCASE").
		Find(&teams)

	var existingSubscriptions []dbModels.Subscription
	database.DB.Where("chat_id = ?", chatID).Find(&existingSubscriptions)

	if teamsResult.RowsAffected == 0 || teamsResult.Error != nil {
		botInstance.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "🤷‍♂️ No teams available for subscription.",
		})
	} else {
		var teamList []string
		for _, team := range teams {
			var subscriptionEmoji string
			var subscriptionCommand string
			if slices.ContainsFunc(existingSubscriptions, func(subscription dbModels.Subscription) bool {
				return subscription.TeamID == team.ID
			}) {
				subscriptionEmoji = "✅"
				subscriptionCommand = "to unsubscribe: <code>/team_unsubscribe "
			} else {
				subscriptionEmoji = "❌"
				subscriptionCommand = "to subscribe: <code>/team_subscribe "
			}
			subscriptionCommand = subscriptionEmoji + " " + team.Label + " - " + subscriptionCommand + team.Label + "</code>"
			teamList = append(teamList, subscriptionCommand)
		}
		// The only message could be too long. Telegram allows only 4096 characters per message
		// So we need to split the message into smaller parts
		messages := generateSubscriptionMessages(teamList)
		for _, message := range messages {
			botInstance.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:    chatID,
				Text:      message,
				ParseMode: models.ParseModeHTML,
			})
			// To avoid throttling
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func generateSubscriptionMessages(teamLines []string) []string {
	var messages []string
	message := messagesPrefix
	for _, teamLine := range teamLines {
		if (len(message) + len(teamLine+"\n")) < maxMessageLength {
			message += teamLine + "\n"
		} else {
			messages = append(messages, message)
			message = messagesPrefix + teamLine + "\n"
		}
	}
	messages = append(messages, message)

	return messages
}
