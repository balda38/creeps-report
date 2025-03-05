package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	commandsCore "github.com/balda38/creeps-report/commands/core"
	"github.com/balda38/creeps-report/database"
	dbModels "github.com/balda38/creeps-report/database/models"
	"github.com/balda38/creeps-report/opendotaclient/matches"
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

	go StartWatchingRecentMatches(database.DB, creepsReportBot)

	commandsCore.RegisterForBot(creepsReportBot)
	creepsReportBot.Start(context.Background())
}

// TODO: cron/system.d or something like that
func StartWatchingRecentMatches(db *gorm.DB, botInstance *bot.Bot) {
	var lastMatchTime int64
	db.Model(&dbModels.Team{}).Select("MAX(last_match_time)").Row().Scan(&lastMatchTime)

	for {
		matches.FetchRecentMatches(db, botInstance, &lastMatchTime)
		time.Sleep(5 * time.Second)
	}
}
