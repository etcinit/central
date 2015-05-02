package entities

import "time"

// A Document is a simple text file stored in Central.
type Document struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Contents  string    `json:"contents"`
	Alias     string    `json:"alias"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updated_at"`
}
