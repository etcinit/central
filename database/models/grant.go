package models

import "time"

// A Grant allows an Application to access a certain file. Additionally, a
// grant can specify which name the application will see for the file (alias).
type Grant struct {
	ID            uint `gorm:"primary_key"`
	ApplicationID uint `gorm:"column:ApplicationId"`
	Application   Application
	DocumentID    uint `gorm:"column:FileId"`
	Document      Document
	Alias         string
	CreatedAt     time.Time `gorm:"column:createdAt"`
	UpdatedAt     time.Time `gorm:"column:updatedAt"`
}

// TableName gets the name of the database table to use for this model.
func (g Grant) TableName() string {
	return "Grants"
}
