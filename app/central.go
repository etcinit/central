package app

import (
	"github.com/etcinit/ohmygorm"
	"github.com/jacobstr/confer"
)

// Central is the root node of the DI graph
type Central struct {
	Config      *confer.Config               `inject:""`
	Connections *ohmygorm.ConnectionsService `inject:""`
	Engine      *EngineService               `inject:""`
	Serve       *ServeService                `inject:""`
}
