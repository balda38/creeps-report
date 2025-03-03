package main

import (
	"context"
	"log"
	"os"

	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"

	commandsCore "github.com/balda38/creeps-report/commands/core"
	"github.com/balda38/creeps-report/database"
)

func main() {
	dotEnvErr := godotenv.Load()
	if dotEnvErr != nil {
		log.Fatal(".env file doesn't exist")
	}

	database.EnableDBConnection()

	creepsReportBot, botCreationErr := bot.New(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if botCreationErr != nil {
		log.Fatal("Failed to create bot: ", botCreationErr)
	}
	commandsCore.RegisterForBot(creepsReportBot)
	creepsReportBot.Start(context.Background())
}
