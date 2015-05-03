package v1

import (
	"github.com/etcinit/central/app/responses"
	"github.com/gin-gonic/gin"
	"gopkg.in/bluesuncorp/validator.v5"
)

// PingController handles instance pinging routes.
type PingController struct{}

// Register registers all the routes for this controller.
func (p *PingController) Register(r *gin.RouterGroup) {
	r.POST("/ping", p.postPing)
}

func (p *PingController) postPing(c *gin.Context) {
	var json PingJSON

	c.Bind(&json)

	validate := validator.New("validate", validator.BakedInValidators)
	errs := validate.Struct(&json)
	if errs != nil {
		responses.SendValidationMessages(c, errs, &json)
		return
	}

	c.JSON(200, json)
}
