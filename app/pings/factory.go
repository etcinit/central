package pings

import (
	"net/http"

	"github.com/etcinit/central/app/entities"
	"github.com/etcinit/central/app/v1/requests"
	"github.com/etcinit/speedbump/ginbump"
)

// Factory provides factory functions for creating instances of Ping entities.
type Factory struct{}

// MakeFromV1 create a new Ping entitie from a API v1 request.
func (f *Factory) MakeFromV1(request *http.Request, input *requests.PingJSON) *entities.Ping {
	ip := ginbump.GetRequesterAddress(request)

	return &entities.Ping{
		InstanceID: input.InstanceID,
		Message:    input.Message,
		IP:         ip,
		Version:    "",
	}
}
