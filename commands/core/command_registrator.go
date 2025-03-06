package core

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"github.com/balda38/creeps-report/commands"
)

func GetCommandsToRegister() []Command {
	return []Command{
		&commands.StartCommand{},
		&commands.HelpCommand{},
		&commands.TeamSubscribeCommand{},
		&commands.TeamUnsubscribeCommand{},
		&commands.TeamSubscriptionsCommand{},
	}
}

func RegisterForBot(botInstance *bot.Bot) {
	var myCommands = []models.BotCommand{}
	for _, command := range GetCommandsToRegister() {
		botInstance.RegisterHandler(
			bot.HandlerTypeMessageText,
			command.GetName(),
			command.CommandMatchType(),
			command.Handler,
		)
		myCommands = append(myCommands, models.BotCommand{
			Command:     command.GetName(),
			Description: command.GetDescription(),
		})
	}

	botInstance.SetMyCommands(
		context.Background(),
		&bot.SetMyCommandsParams{
			Commands: myCommands,
		},
	)
}
