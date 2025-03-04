package main

import (
	"os"
	"strings"
	"time"

	"github.com/balda38/creeps-report/database"
	"github.com/balda38/creeps-report/database/models"
)

// TODO: add support for multiple team labels
func main() {
	database.EnableDBConnection()

	if len(os.Args) < 2 {
		timestampToCompare := time.Now().AddDate(0, -6, 0).Unix()
		database.DB.Model(&models.Team{}).
			Where("is_active = ?", true).
			Where("last_match_time < ?", timestampToCompare).
			Update("is_active", false)
	} else {
		team := strings.ReplaceAll(os.Args[1], "_", " ")
		database.DB.Model(&models.Team{}).Where("label = ?", team).Update("is_active", false)
	}
}
