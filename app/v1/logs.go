package v1

import (
	"github.com/etcinit/central/app/responses"
	"github.com/gin-gonic/gin"
	"gopkg.in/bluesuncorp/validator.v5"
)

// LogsController handles log pushing routes.
type LogsController struct{}

// Register registers all the routes for this controller.
func (l *LogsController) Register(r *gin.RouterGroup) {
	r.POST("/logs", l.postLogs)
}

func (l *LogsController) postLogs(c *gin.Context) {
	var json LogEntryJSON

	c.Bind(&json)

	validate := validator.New("validate", validator.BakedInValidators)
	errs := validate.Struct(&json)
	if errs != nil {
		responses.SendValidationMessages(c, errs, &json)
		return
	}

	c.JSON(200, json)
}
