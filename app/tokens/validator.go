package tokens

import (
	"time"

	"github.com/etcinit/central/database/models"
)

// Validator provides methods for validating an application token.
type Validator struct{}

// Expired returns whether or not a token has expired.
func (v *Validator) Expired(token *models.ApplicationToken) bool {
	return time.Now().After(token.ExpirationDate)
}
