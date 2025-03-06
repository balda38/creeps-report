package database

import (
	"log"
	"sync"

	"github.com/balda38/creeps-report/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func EnableDBConnection() {
	once.Do(func() {
		var err error
		DB, err = gorm.Open(sqlite.Open("data/bot.db"), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database: ", err)
		}
	})
}

func RunMigrations() {
	err := DB.AutoMigrate(
		&models.Team{},
		&models.Subscription{},
	)
	if err != nil {
		log.Fatal("Migration failed: ", err)
	}
}

func CloseDBConnection() {
	db, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get DB: ", err)
	}
	err = db.Close()
	if err != nil {
		log.Fatal("Failed to close DB: ", err)
	}
}
