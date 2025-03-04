package opendotaclient

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"strconv"
	"strings"

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
	var teams []OpenDotaTeam
	page := 0
	for {
		response, err := http.Get(teamsAPI + "?page=" + strconv.Itoa(page))
		if err != nil {
			log.Fatal("Error fetching teams:", err)
		}
		defer response.Body.Close()

		var teamsOnPage []OpenDotaTeam
		if err := json.NewDecoder(response.Body).Decode(&teamsOnPage); err != nil {
			log.Fatal("Error decoding JSON:", err)
		}
		teams = append(teams, teamsOnPage...)

		if len(teamsOnPage) < 1000 {
			break
		}

		page++
	}

	var teamModels []models.Team
	for _, team := range teams {
		teamName := strings.TrimSpace(team.Name)
		if teamName == "" {
			continue
		}

		// If the team with the same name exists, update the last match time if it's newer
		existingTeamIndex := slices.IndexFunc(teamModels, func(teamModel models.Team) bool {
			return strings.EqualFold(teamName, teamModel.Label)
		})
		if existingTeamIndex == -1 {
			teamModels = append(teamModels, models.Team{
				ID:            team.ID,
				Label:         teamName,
				LastMatchTime: team.LastMatchTime,
			})
		} else if teamModels[existingTeamIndex].LastMatchTime < team.LastMatchTime {
			teamModels[existingTeamIndex] = models.Team{
				ID:            team.ID,
				Label:         teamName,
				LastMatchTime: team.LastMatchTime,
			}
		}
	}

	return teamModels
}
