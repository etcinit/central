package applications

import (
	"testing"

	"github.com/etcinit/central/app/util"
	"github.com/etcinit/central/database/models"
	"github.com/etcinit/ohmygorm"
	"github.com/facebookgo/inject"
	"github.com/stretchr/testify/assert"
)

func TestFindByToken(t *testing.T) {
	var service Finder
	var connections ohmygorm.ConnectionsService

	util.PopulateDatabaseTest(
		t,
		&inject.Object{Value: &service},
		&inject.Object{Value: &connections},
		&inject.Object{Value: util.GetTestingConfig()},
	)

	token := models.ApplicationToken{
		ApplicationID: 1,
	}

	instance, err := service.FindByToken(&token)
	assert.Nil(t, instance)
	assert.NotNil(t, err)

	db, _ := connections.Make()
	db.Create(&models.Application{
		Name:        "blahblah",
		Description: "something",
	})

	instance, err = service.FindByToken(&token)
	assert.NotNil(t, instance)
	assert.Nil(t, err)
}
