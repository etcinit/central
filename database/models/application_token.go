package models

import "time"

// An ApplicationToken is a simple API token which dictates whether a client
// has access to the Central and which application data it has access to.
type ApplicationToken struct {
	ID             uint `gorm:"primary_key"`
	Token          string
	Comment        string
	ApplicationID  uint `gorm:"column:ApplicationId"`
	ExpirationDate time.Time
	CreatedAt      time.Time `gorm:"column:createdAt"`
	UpdatedAt      time.Time `gorm:"column:updatedAt"`
}

// TableName gets the name of the database table to use for this model.
func (t ApplicationToken) TableName() string {
	return "ApplicationTokens"
}
