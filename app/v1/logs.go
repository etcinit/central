package v1

import (
	"github.com/etcinit/central/app/logs"
	"github.com/etcinit/central/app/responses"
	"github.com/etcinit/central/app/v1/requests"
	"github.com/gin-gonic/gin"
	"gopkg.in/bluesuncorp/validator.v5"
)

// LogsController handles log pushing routes.
type LogsController struct {
	Factory *logs.Factory `inject:""`
	Pusher  *logs.Pusher  `inject:""`
}

// Register registers all the routes for this controller.
func (l *LogsController) Register(r *gin.RouterGroup) {
	r.POST("/logs", l.postLogs)
}

func (l *LogsController) postLogs(c *gin.Context) {
	var json requests.LogEntryJSON

	c.Bind(&json)

	validate := validator.New("validate", validator.BakedInValidators)
	errs := validate.Struct(&json)
	if errs != nil {
		responses.SendValidationMessages(c, errs, &json)
		return
	}

	entries := l.Factory.MakeFromV1(&json)
	err := l.Pusher.Push(entries)

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"input": json,
	})
}
