package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/balda38/creeps-report/database"
	"github.com/balda38/creeps-report/database/models"
	"gorm.io/gorm/clause"
)

const teamsAPI = "https://api.opendota.com/api/teams"

type OpenDotaTeam struct {
	ID            int     `json:"team_id"`
	Name          string  `json:"name"`
	Tag           string  `json:"tag"`
	Win           int     `json:"wins"`
	Loss          int     `json:"losses"`
	Rating        float64 `json:"rating"`
	LastMatchTime int64   `json:"last_match_time"`
	LogoUrl       string  `json:"logo_url"`
}

func main() {
	resp, err := http.Get(teamsAPI)
	if err != nil {
		log.Fatal("Error fetching teams:", err)
	}
	defer resp.Body.Close()

	var teams []OpenDotaTeam
	if err := json.NewDecoder(resp.Body).Decode(&teams); err != nil {
		log.Fatal("Error decoding JSON:", err)
	}

	var teamsToInsert []models.Team
	for _, team := range teams {
		// Do not parse teams that have not played a match in last 2 years. TODO: probably, it should be configurable
		// Probably, it should be also filtered by rating (but i don't know what amount should be used)
		if team.LastMatchTime < time.Now().AddDate(-2, 0, 0).Unix() || team.Name == "" {
			continue
		}
		teamsToInsert = append(teamsToInsert, models.Team{
			ID:    team.ID,
			Label: team.Name,
		})
	}

	database.EnableDBConnection()
	if len(teamsToInsert) > 0 {
		result := database.DB.Clauses(
			clause.OnConflict{DoNothing: true},
		).Create(&teamsToInsert)
		if result.Error != nil {
			log.Fatal("Error inserting teams:", result.Error)
		}
	}

	log.Printf("Successfully stored %d teams in the database!\n", len(teamsToInsert))
}
