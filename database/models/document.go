package models

import "time"

// A Document is a simple text file stored in a database.
type Document struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Contents  string    `json:"contents"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updated_at"`
}

// TableName gets the name of the database table to use for this model.
func (d Document) TableName() string {
	return "Files"
}
