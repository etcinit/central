package documents

import (
	"testing"

	"github.com/etcinit/central/app/util"
	"github.com/etcinit/central/database/models"
	"github.com/etcinit/ohmygorm"
	"github.com/facebookgo/inject"
	"github.com/stretchr/testify/assert"
)

func TestFindAllByApplication(t *testing.T) {
	var service Finder
	var connections ohmygorm.ConnectionsService

	util.PopulateDatabaseTest(
		t,
		&inject.Object{Value: &service},
		&inject.Object{Value: &connections},
		&inject.Object{Value: util.GetTestingConfig()},
	)

	app := models.Application{
		Name:        "testApp",
		Description: "Testing application",
	}

	db, _ := connections.Make()
	db.LogMode(true)
	db.Create(&app)

	results, err := service.FindAllByApplication(&app)
	assert.True(t, len(results) == 0)
	assert.Nil(t, err)

	document := models.Document{
		Name:     "needsmore.jpg",
		Contents: "something",
	}
	db.Create(&document)
	db.Create(&models.Grant{
		ApplicationID: app.ID,
		DocumentID:    document.ID,
		Alias:         "needsmore.png",
	})

	results, err = service.FindAllByApplication(&app)
	assert.True(t, len(results) > 0)
	assert.Nil(t, err)
}
