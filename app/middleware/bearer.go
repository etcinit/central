package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/etcinit/central/app/tokens"
	"github.com/gin-gonic/gin"
)

// BearerGenerator provides functions for generating bearer auth middleware
// handlers.
type BearerGenerator struct {
	Finder    *tokens.Finder    `inject:""`
	Validator *tokens.Validator `inject:""`
}

// NewMiddleware creates a new bearer auth middleware handler.
func (b *BearerGenerator) NewMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Be pessimistic
		ok := false

		// Check that the token is present and valid
		if tokenStr, err := FromAuthHeader(c.Request); err == nil {
			token, err := b.Finder.FindByToken(tokenStr)

			if err == nil && !b.Validator.Expired(token) {
				ok = true
				c.Set("token", token)
			}
		}

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":   "error",
				"messages": []string{"Invalid or expired authorization token"},
			})
			c.Abort()
		} else {
			c.Next()
		}
	}
}

// FromAuthHeader gets the bearer token from the auth header.
func FromAuthHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("No token provided")
	}

	authHeaderParts := strings.Split(authHeader, " ")

	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", fmt.Errorf("Authorization header format must be Bearer {token}")
	}

	return authHeaderParts[1], nil
}
