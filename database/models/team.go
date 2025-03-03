package models

type Team struct {
	ID    int `gorm:"primaryKey"`
	Label string
	// TODO: more fields
}
