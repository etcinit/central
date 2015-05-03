package v1

import (
	"github.com/etcinit/central/app/pings"
	"github.com/etcinit/central/app/responses"
	"github.com/etcinit/central/app/v1/requests"
	"github.com/gin-gonic/gin"
	"gopkg.in/bluesuncorp/validator.v5"
)

// PingController handles instance pinging routes.
type PingController struct {
	Factory *pings.Factory `inject:""`
	Pusher  *pings.Pusher  `inject:""`
}

// Register registers all the routes for this controller.
func (p *PingController) Register(r *gin.RouterGroup) {
	r.POST("/ping", p.postPing)
}

func (p *PingController) postPing(c *gin.Context) {
	var json requests.PingJSON

	c.Bind(&json)

	validate := validator.New("validate", validator.BakedInValidators)
	errs := validate.Struct(&json)
	if errs != nil {
		responses.SendValidationMessages(c, errs, &json)
		return
	}

	ping := p.Factory.MakeFromV1(c.Request, &json)
	err := p.Pusher.Push(ping)

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"status": "success",
		"input":  json,
	})
}
