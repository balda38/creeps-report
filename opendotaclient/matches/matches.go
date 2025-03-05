package matches

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"

	dbModels "github.com/balda38/creeps-report/database/models"
	"github.com/balda38/creeps-report/notificator"
	"github.com/balda38/creeps-report/opendotaclient/types"
	"github.com/go-telegram/bot"
	"gorm.io/gorm"
)

const teamMatchesAPI = "https://api.opendota.com/api/proMatches"

func FetchRecentMatches(db *gorm.DB, botInstance *bot.Bot, lastMatchTime *int64) {
	response, err := http.Get(teamMatchesAPI)
	if err != nil {
		log.Fatal("Error fetching team results:", err)
	}
	defer response.Body.Close()

	var teamMatches []types.OpenDotaMatchShort
	if err := json.NewDecoder(response.Body).Decode(&teamMatches); err != nil {
		log.Fatal("Error decoding JSON:", err)
	}

	// TODO: is this possible that several matches end at the same time?
	if len(teamMatches) > 0 && teamMatches[0].StartTime > *lastMatchTime {
		var subscriptions []dbModels.Subscription
		db.Model(&dbModels.Subscription{}).Find(&subscriptions)

		for _, match := range teamMatches {
			if match.StartTime <= *lastMatchTime {
				break
			}

			updateResult := db.Model(&dbModels.Team{}).
				Where("id IN (?, ?)", match.RadiandTeamId, match.DireTeamId).
				Where("last_match_time < ?", match.StartTime).
				Updates(dbModels.Team{IsActive: true, LastMatchTime: match.StartTime})

			if updateResult.Error == nil && updateResult.RowsAffected > 0 {
				var relatedChats []int64
				for _, subsubscription := range subscriptions {
					if (subsubscription.TeamID == match.RadiandTeamId ||
						subsubscription.TeamID == match.DireTeamId) && !slices.Contains(relatedChats, subsubscription.ChatID) {
						relatedChats = append(relatedChats, subsubscription.ChatID)
					}
				}
				if len(relatedChats) > 0 {
					go notificator.NotifySubscribers(botInstance, relatedChats, match.ID)
				}
			}
		}

		*lastMatchTime = teamMatches[0].StartTime
	}
}
