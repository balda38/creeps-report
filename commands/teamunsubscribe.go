package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"github.com/balda38/creeps-report/database"
	dbModels "github.com/balda38/creeps-report/database/models"
)

type TeamUnsubscribeCommand struct{}

func (TeamUnsubscribeCommand) GetName() string {
	return "/team_unsubscribe"
}

func (TeamUnsubscribeCommand) GetDescription() string {
	return "Unsubscribe from team match updates"
}

func (TeamUnsubscribeCommand) CommandMatchType() bot.MatchType {
	return bot.MatchTypePrefix
}

func (command TeamUnsubscribeCommand) Handler(ctx context.Context, botInstance *bot.Bot, update *models.Update) {
	chatID := update.Message.Chat.ID
	team := strings.TrimSpace(strings.TrimPrefix(update.Message.Text, command.GetName()))

	if team == "" {
		botInstance.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   "📝 Usage: " + command.GetName() + " <team_name>",
		})
		return
	}

	var dbTeam = dbModels.Team{Label: team}
	dbTeamResult := database.DB.First(&dbTeam, dbTeam)
	if dbTeamResult.RowsAffected == 0 || dbTeamResult.Error != nil {
		botInstance.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   fmt.Sprintf("❌ Invalid team name: %s.", team),
		})
		return
	}

	var dbTeamSubsription = dbModels.Subscription{TeamID: dbTeam.ID, ChatID: chatID}
	dbTeamSubsriptionResult := database.DB.First(&dbTeamSubsription, dbTeamSubsription)
	if dbTeamSubsriptionResult.RowsAffected == 0 || dbTeamSubsriptionResult.Error != nil {
		botInstance.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   fmt.Sprintf("🤷‍♂️ You're not subscribed to %s.", team),
		})
		return
	}

	database.DB.Delete(&dbTeamSubsription)
	botInstance.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   fmt.Sprintf("💔 Unsubscribed from %s!", team),
	})
}
