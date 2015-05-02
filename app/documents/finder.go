package documents

import (
	"github.com/etcinit/central/app/entities"
	"github.com/etcinit/central/database/models"
	"github.com/etcinit/ohmygorm"
)

// Finder finds documents using different methods.
type Finder struct {
	Connections *ohmygorm.ConnectionsService `inject:""`
	Repository  *ohmygorm.RepositoryService  `inject:""`
}

// FindAllByApplication gets all the documents that an application can currently
// access.
func (f *Finder) FindAllByApplication(app *models.Application) ([]entities.Document, error) {
	var documents []entities.Document

	db, err := f.Connections.Make()

	if err != nil {
		return nil, err
	}

	err = db.Table("Files").
		Select("Files.*, Grants.alias").
		Joins("LEFT JOIN Grants on Grants.FileId = Files.id").
		Where("Grants.ApplicationId = ?", app.ID).
		Scan(&documents).
		Error

	if err != nil {
		return nil, err
	}

	return documents, nil
}
