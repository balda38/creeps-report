package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"slices"
	"strconv"
	"time"

	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	commandsCore "github.com/balda38/creeps-report/commands/core"
	"github.com/balda38/creeps-report/database"
	dbModels "github.com/balda38/creeps-report/database/models"
	"github.com/balda38/creeps-report/notificator"
	"github.com/balda38/creeps-report/opendotaclient"
)

func main() {
	dotEnvErr := godotenv.Load()
	if dotEnvErr != nil {
		log.Fatal(".env file doesn't exist")
	}

	database.EnableDBConnection()

	// TODO: add secret token for webhook
	creepsReportBot, botCreationErr := bot.New(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if botCreationErr != nil {
		log.Fatal("Failed to create bot: ", botCreationErr)
	}

	go StartWatchingRecentMatches(database.DB, creepsReportBot)

	commandsCore.RegisterForBot(creepsReportBot)

	go creepsReportBot.StartWebhook(context.Background())

	log.Fatal(http.ListenAndServe(
		os.Getenv("APP_URL")+":"+os.Getenv("PORT"),
		creepsReportBot.WebhookHandler()),
	)
}

// TODO: cron/system.d or something like that
func StartWatchingRecentMatches(db *gorm.DB, botInstance *bot.Bot) {
	timeout, _ := strconv.Atoi(os.Getenv("FETCH_RECENT_MATCHES_TIMEOUT"))
	timeoutDuration := time.Duration(timeout) * time.Second

	var lastSavedMatchTime int64
	db.Model(&dbModels.Team{}).Select("MAX(last_match_time)").Row().Scan(&lastSavedMatchTime)

	for {
		recentMatches := opendotaclient.FetchRecentMatches()

		if len(recentMatches) > 0 {
			var subscriptions []dbModels.Subscription
			db.Model(&dbModels.Subscription{}).Find(&subscriptions)

			recentMatchTime := recentMatches[0].StartTime

			// TODO: create team if it doesn't exist (?)
			for _, match := range recentMatches {
				updateResult := db.Model(&dbModels.Team{}).
					Debug().
					Where("id IN (?, ?)", match.RadiandTeamId, match.DireTeamId).
					Where("last_match_time < ?", match.StartTime).
					Updates(dbModels.Team{IsActive: true, LastMatchTime: match.StartTime})

				if updateResult.Error == nil && updateResult.RowsAffected > 0 {
					relatedChats := []int64{}
					for _, subsubscription := range subscriptions {
						if (subsubscription.TeamID == match.RadiandTeamId ||
							subsubscription.TeamID == match.DireTeamId) &&
							!slices.Contains(relatedChats, subsubscription.ChatID) {
							relatedChats = append(relatedChats, subsubscription.ChatID)
						}
					}
					if len(relatedChats) > 0 {
						go notificator.NotifySubscribers(botInstance, relatedChats, match.ID)
					}
				}

				if match.StartTime > recentMatchTime {
					recentMatchTime = match.StartTime
				}
			}

			lastSavedMatchTime = recentMatchTime
		}

		time.Sleep(timeoutDuration)
	}
}
