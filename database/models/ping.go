package models

import "time"

// A Ping is a legacy model for storing information about a "ping back home"
// that each instance of an application did.
type Ping struct {
	ID            uint   `gorm:"primary_key"`
	InstanceName  string `gorm:"column:instanceName"`
	ApplicationID uint   `gorm:"column:ApplicationId"`
	Message       string
	CreatedAt     time.Time `gorm:"column:createdAt"`
	UpdatedAt     time.Time `gorm:"column:updatedAt"`
}

// TableName gets the name of the database table to use for this model.
func (p Ping) TableName() string {
	return "Pings"
}
