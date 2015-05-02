package models

import "time"

// An Application is an abstract representation of the product or server that
// will be using Central to obtain its configuration.
type Application struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Grants      []Grant   `json:"-"`
	CreatedAt   time.Time `gorm:"column:createdAt" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updatedAt" json:"updated_at"`
}

// TableName gets the name of the database table to use for this model.
func (a Application) TableName() string {
	return "Applications"
}
