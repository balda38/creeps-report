package opendotaclient

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/balda38/creeps-report/database/models"
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

func FetchTeams() []models.Team {
	response, err := http.Get(teamsAPI)
	if err != nil {
		log.Fatal("Error fetching teams:", err)
	}
	defer response.Body.Close()

	var teams []OpenDotaTeam
	if err := json.NewDecoder(response.Body).Decode(&teams); err != nil {
		log.Fatal("Error decoding JSON:", err)
	}

	var teamModels []models.Team
	for _, team := range teams {
		if team.Name == "" {
			continue
		}

		teamModels = append(teamModels, models.Team{
			ID:            team.ID,
			Label:         team.Name,
			LastMatchTime: team.LastMatchTime,
		})
	}

	return teamModels
}
