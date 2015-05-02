package applications

import (
	"github.com/etcinit/central/database/models"
	"github.com/etcinit/ohmygorm"
)

// Finder finds applications using different methods.
type Finder struct {
	Connections *ohmygorm.ConnectionsService `inject:""`
	Repository  *ohmygorm.RepositoryService  `inject:""`
}

// FindByToken finds an application by using a related token.
func (f *Finder) FindByToken(token *models.ApplicationToken) (*models.Application, error) {
	var app models.Application

	db, err := f.Connections.Make()
	err = f.Repository.FirstOrFail(
		&app,
		db.Where(&models.Application{
			ID: token.ApplicationID,
		}),
	)

	if err != nil {
		return nil, err
	}

	return &app, nil
}
