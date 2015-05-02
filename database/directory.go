package database

import "github.com/etcinit/central/database/models"

// Directory provides a database model directory.
type Directory struct{}

// GetModels gets a slice of all the models used by Central.
func (d *Directory) GetModels() []interface{} {
	return []interface{}{
		&models.Application{},
		&models.ApplicationToken{},
		&models.Grant{},
		&models.Document{},
		&models.Ping{},
		&models.User{},
	}
}
