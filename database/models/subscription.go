package models

// TODO: foreign key constraint?
// TODO: unique key constraint
type Subscription struct {
	ID     int64 `gorm:"primaryKey"`
	TeamID int
	ChatID int64 `gorm:"index"`
	// TODO: more fields?

	// Relationships
	Team Team `gorm:"foreignKey:TeamID"`
}
