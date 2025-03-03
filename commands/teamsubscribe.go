package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/balda38/creeps-report/database"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	dbModels "github.com/balda38/creeps-report/database/models"
)

type TeamSubscribeCommand struct{}

func (TeamSubscribeCommand) GetName() string {
	return "/team_subscribe"
}

func (TeamSubscribeCommand) GetDescription() string {
	return "Subscribe to team match updates"
}

func (TeamSubscribeCommand) CommandMatchType() bot.MatchType {
	return bot.MatchTypePrefix
}

func (command TeamSubscribeCommand) Handler(ctx context.Context, botInstance *bot.Bot, update *models.Update) {
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
	if dbTeamSubsriptionResult.RowsAffected == 1 {
		botInstance.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: chatID,
			Text:   fmt.Sprintf("🤷‍♂️ You're already subscribed to %s.", team),
		})
		return
	}

	database.DB.Create(&dbTeamSubsription)
	botInstance.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: chatID,
		Text:   fmt.Sprintf("♥️ Subscribed to %s!", team),
	})
}
