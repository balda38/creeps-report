package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/balda38/creeps-report/database"
	"github.com/balda38/creeps-report/database/models"
)

// TODO: add support for multiple teams
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Unknown command")
	}
	team := strings.ReplaceAll(os.Args[1], "_", " ")
	database.EnableDBConnection()
	database.DB.Model(&models.Team{}).Where("label = ?", team).Update("is_active", false)
}
