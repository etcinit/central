package tokens

import (
	"testing"
	"time"

	"github.com/etcinit/central/database/models"
	"github.com/stretchr/testify/assert"
)

func TestExpired(t *testing.T) {
	token := models.ApplicationToken{
		ExpirationDate: time.Now().Add(time.Hour),
	}

	validator := Validator{}

	assert.False(t, validator.Expired(&token))

	token.ExpirationDate = time.Now().Add(-time.Hour)
	assert.True(t, validator.Expired(&token))
}
