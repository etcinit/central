package app

import (
	"github.com/etcinit/central/app/cache"
	"github.com/etcinit/central/app/controllers"
	"github.com/etcinit/central/app/middleware"
	"github.com/etcinit/central/app/v1"
	"github.com/etcinit/speedbump"
	"github.com/etcinit/speedbump/ginbump"
	"github.com/gin-gonic/gin"
	"github.com/jacobstr/confer"
)

// EngineService provides the API engine
type EngineService struct {
	Config *confer.Config               `inject:""`
	Bearer *middleware.BearerGenerator  `inject:""`
	Front  *controllers.FrontController `inject:""`
	Fetch  *v1.FetchController          `inject:""`
	Logs   *v1.LogsController           `inject:""`
	Ping   *v1.PingController           `inject:""`
	Cache  *cache.RedisService          `inject:""`
}

// New creates a new instance of an API engine
func (e *EngineService) New() *gin.Engine {
	router := gin.Default()

	// Setup some rate limits to avoid having one client overload the server.
	redis := e.Cache.Make()
	router.Use(
		ginbump.RateLimit(
			redis,
			speedbump.PerMinuteHasher{},
			int64(e.Config.GetInt("server.limit")),
		),
	)

	// Setup the main routes.
	e.Front.Register(router)

	// Setup the V1 API routes
	v1Api := router.Group("/v1")
	{
		v1Api.Use(e.Bearer.NewMiddleware())
		e.Fetch.Register(v1Api)
		e.Logs.Register(v1Api)
		e.Ping.Register(v1Api)
	}

	return router
}
