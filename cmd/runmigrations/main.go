package main

import "github.com/balda38/creeps-report/database"

func main() {
	database.EnableDBConnection()
	database.RunMigrations()
}
