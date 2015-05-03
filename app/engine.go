package app

import (
	"github.com/etcinit/central/app/controllers"
	"github.com/etcinit/central/app/middleware"
	"github.com/etcinit/central/app/v1"
	"github.com/gin-gonic/gin"
)

// EngineService provides the API engine
type EngineService struct {
	Bearer *middleware.BearerGenerator  `inject:""`
	Front  *controllers.FrontController `inject:""`
	Fetch  *v1.FetchController          `inject:""`
	Logs   *v1.LogsController           `inject:""`
	Ping   *v1.PingController           `inject:""`
}

// New creates a new instance of an API engine
func (e *EngineService) New() *gin.Engine {
	router := gin.Default()

	e.Front.Register(router)

	v1Api := router.Group("/v1")
	{
		v1Api.Use(e.Bearer.NewMiddleware())
		e.Fetch.Register(v1Api)
		e.Logs.Register(v1Api)
		e.Ping.Register(v1Api)
	}

	return router
}
