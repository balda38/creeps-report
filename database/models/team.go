package models

type Team struct {
	ID            int `gorm:"primaryKey"`
	Label         string
	IsActive      bool
	LastMatchTime int64
	// TODO: more fields

	// Relationships
	Subscriptions []Subscription
}
