package main

import (
	"log"
	"time"

	"github.com/balda38/creeps-report/database"
	"github.com/balda38/creeps-report/opendotaclient"
	"gorm.io/gorm/clause"
)

func main() {
	teamsToInsert := opendotaclient.FetchTeams()

	database.EnableDBConnection()
	if len(teamsToInsert) > 0 {
		for teamIndex, team := range teamsToInsert {
			// Set team as inactive if they have not played a match in last six months. TODO: probably, it should be configurable/less (?)
			// Probably, it should be also filtered by rating (but i don't know what amount should be used. By maybe 1000)
			teamsToInsert[teamIndex].IsActive = team.LastMatchTime >= time.Now().AddDate(0, -6, 0).Unix()
		}

		result := database.DB.Clauses(
			clause.OnConflict{DoNothing: true},
		).CreateInBatches(&teamsToInsert, 1000)
		if result.Error != nil {
			log.Fatal("Error inserting teams:", result.Error)
		}
	}

	log.Printf("Successfully stored %d teams in the database!\n", len(teamsToInsert))
}
