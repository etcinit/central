package app

import (
	"github.com/etcinit/central/app/controllers"
	"github.com/etcinit/central/app/middleware"
	"github.com/gin-gonic/gin"
)

// EngineService provides the API engine
type EngineService struct {
	Bearer *middleware.BearerGenerator  `inject:""`
	Front  *controllers.FrontController `inject:""`
	Fetch  *controllers.FetchController `inject:""`
}

// New creates a new instance of an API engine
func (e *EngineService) New() *gin.Engine {
	router := gin.Default()

	e.Front.Register(router)

	v1 := router.Group("/v1")
	{
		v1.Use(e.Bearer.NewMiddleware())
		e.Fetch.Register(v1)
	}

	return router
}
