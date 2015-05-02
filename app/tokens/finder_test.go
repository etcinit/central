package tokens

import (
	"testing"
	"time"

	"github.com/etcinit/central/app/util"
	"github.com/etcinit/central/database/models"
	"github.com/etcinit/ohmygorm"
	"github.com/facebookgo/inject"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	var service Finder
	var connections ohmygorm.ConnectionsService

	util.PopulateDatabaseTest(
		t,
		&inject.Object{Value: &service},
		&inject.Object{Value: &connections},
		&inject.Object{Value: util.GetTestingConfig()},
	)

	instance, err := service.Find(1)
	assert.Nil(t, instance)
	assert.NotNil(t, err)

	db, _ := connections.Make()
	db.Create(&models.ApplicationToken{
		Token:          "blahblah",
		Comment:        "something",
		ExpirationDate: time.Now().Add(time.Hour),
		ApplicationID:  2,
	})

	instance, err = service.Find(1)
	assert.NotNil(t, instance)
	assert.Nil(t, err)
}

func TestFindByToken(t *testing.T) {
	var service Finder
	var connections ohmygorm.ConnectionsService

	util.PopulateDatabaseTest(
		t,
		&inject.Object{Value: &service},
		&inject.Object{Value: &connections},
		&inject.Object{Value: util.GetTestingConfig()},
	)

	instance, err := service.FindByToken("blahblah")
	assert.Nil(t, instance)
	assert.NotNil(t, err)

	db, _ := connections.Make()
	db.Create(&models.ApplicationToken{
		Token:          "blahblah",
		Comment:        "something",
		ExpirationDate: time.Now().Add(time.Hour),
		ApplicationID:  2,
	})

	instance, err = service.FindByToken("blahblah")
	assert.NotNil(t, instance)
	assert.Nil(t, err)
}
