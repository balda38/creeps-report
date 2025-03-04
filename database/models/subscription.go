package models

type Subscription struct {
	ID     int64 `gorm:"primaryKey"`
	TeamID int   `gorm:"index;uniqueIndex:chat_subscription"`
	ChatID int64 `gorm:"index;uniqueIndex:chat_subscription"`
	// TODO: more fields?

	// Relationships
	Team Team `gorm:"foreignKey:TeamID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
}
