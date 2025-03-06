package opendotaclient

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/balda38/creeps-report/opendotaclient/types"
)

const recentMatchesAPI = "https://api.opendota.com/api/proMatches"
const leagueMatchesAPI = "https://api.opendota.com/api/leagues/%s/matches"
const matchAPI = "https://api.opendota.com/api/matches/%s"

func FetchRecentMatches() []types.OpenDotaMatchShort {
	response, err := http.Get(recentMatchesAPI)
	if err != nil {
		log.Fatal("Error fetching recent matches:", err)
	}
	defer response.Body.Close()

	var recentMatches []types.OpenDotaMatchShort
	if err := json.NewDecoder(response.Body).Decode(&recentMatches); err != nil {
		log.Fatal("Error decoding JSON:", err)
	}

	return recentMatches
}

func FetchLeagueMatches(leagueId int) []types.OpenDotaLeagueMatch {
	response, err := http.Get(fmt.Sprintf(leagueMatchesAPI, strconv.Itoa(leagueId)))
	if err != nil {
		log.Fatal("Error fetching league matches:", err)
	}
	defer response.Body.Close()

	var leagueMatches []types.OpenDotaLeagueMatch
	if err := json.NewDecoder(response.Body).Decode(&leagueMatches); err != nil {
		log.Fatal("Error decoding JSON:", err)
	}

	return leagueMatches
}

func FetchMatch(matchId int) types.OpenDotaMatchExtended {
	response, err := http.Get(fmt.Sprintf(matchAPI, strconv.Itoa(matchId)))
	if err != nil {
		log.Fatal("Error fetching match results:", err)
	}
	defer response.Body.Close()

	var match types.OpenDotaMatchExtended
	if err := json.NewDecoder(response.Body).Decode(&match); err != nil {
		log.Fatal("Error decoding JSON:", err)
	}

	return match
}
