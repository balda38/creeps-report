package main

import (
	"os"
	"strings"
	"time"

	"github.com/balda38/creeps-report/database"
	"github.com/balda38/creeps-report/database/models"
)

func main() {
	database.EnableDBConnection()

	if len(os.Args) < 2 {
		timestampToCompare := time.Now().AddDate(0, -6, 0).Unix()
		database.DB.Model(&models.Team{}).
			Where("is_active = ?", false).
			Where("last_match_time >= ?", timestampToCompare).
			Update("is_active", true)
	} else {
		var teams []string
		for _, team := range os.Args[1:] {
			teams = append(teams, strings.ReplaceAll(team, "_", " "))
		}
		database.DB.Model(&models.Team{}).Where("label IN (?)", teams).Update("is_active", true)
	}
}
