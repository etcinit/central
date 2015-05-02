package tokens

import (
	"github.com/etcinit/central/database/models"
	"github.com/etcinit/ohmygorm"
)

// Finder finds tokens using different methods.
type Finder struct {
	Connections *ohmygorm.ConnectionsService `inject:""`
	Repository  *ohmygorm.RepositoryService  `inject:""`
}

// Find looks for a specific token by its database ID.
func (f *Finder) Find(id int) (*models.ApplicationToken, error) {
	token := models.ApplicationToken{}

	if err := f.Repository.Find(&token, id); err != nil {
		return nil, err
	}

	return &token, nil
}

// FindByToken attempts to find a token model instance by its token string.
func (f *Finder) FindByToken(token string) (*models.ApplicationToken, error) {
	db, err := f.Connections.Make()

	instance := models.ApplicationToken{}

	err = f.Repository.FirstOrFail(
		&instance,
		db.Model(&models.ApplicationToken{}).Where("token = ?", token),
	)
	if err != nil {
		return nil, err
	}

	return &instance, nil
}
