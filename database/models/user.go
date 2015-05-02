package models

import "time"

// A User is an administrator or developer who has access to modifying resources
// on Central.
type User struct {
	ID        uint `gorm:"primary_key"`
	Username  string
	Password  string
	CreatedAt time.Time `gorm:"column:createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt"`
}

// TableName gets the name of the database table to use for this model.
func (u User) TableName() string {
	return "Users"
}
