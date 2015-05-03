package responses

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	validator "gopkg.in/bluesuncorp/validator.v5"
)

// SendResourceNotFound sends an error message indicating that the specified
// resource was not found.
func SendResourceNotFound(c *gin.Context) {
	SendMessages(c, http.StatusNotFound, "The specified resource was not found.")
}

// SendValidationMessages sends a response with validation errors.
func SendValidationMessages(c *gin.Context, errors *validator.StructErrors, object interface{}) {
	var fields []string

	for fieldName := range errors.Errors {
		field, _ := reflect.TypeOf(object).Elem().FieldByName(fieldName)

		fields = append(fields, field.Tag.Get("json"))
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"status": "error",
		"messages": []string{
			"One or more fields are invalid or missing.",
			"Check documentation for proper usage.",
		},
		"fields": fields,
	})
}

// SendInvalidInputMessages sends a response with error messages.
func SendInvalidInputMessages(c *gin.Context, messages ...string) {
	SendMessages(c, http.StatusBadRequest, messages...)
}

// SendMessages sends a response with messages.
func SendMessages(c *gin.Context, code int, messages ...string) {
	status := "success"
	if code != http.StatusOK {
		status = "error"
	}

	c.JSON(code, gin.H{
		"status":   status,
		"messages": messages,
	})
}

// SendSingleResource sends a single resource to the client
func SendSingleResource(c *gin.Context, name string, instance interface{}) {
	c.JSON(200, gin.H{
		"status": "success",
		name:     instance,
	})
}
