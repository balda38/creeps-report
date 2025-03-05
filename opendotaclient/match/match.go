package match

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/balda38/creeps-report/opendotaclient/types"
)

const matchAPI = "https://api.opendota.com/api/matches/%s"

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
