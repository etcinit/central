package logs

import (
	"github.com/etcinit/central/app/entities"
	"github.com/etcinit/central/app/v1/requests"
)

// Factory provides factory functions for creating instances of LogEntry
// entities.
type Factory struct{}

// MakeFromV1 converts a request from the v1 API into multiple entry entities.
func (f *Factory) MakeFromV1(input *requests.LogEntryJSON) []*entities.LogEntry {
	var entries []*entities.LogEntry

	for _, line := range input.Lines {
		entries = append(
			entries,
			&entities.LogEntry{
				InstanceID: input.InstanceID,
				Name:       input.Name,
				Line:       line,
			},
		)
	}

	return entries
}
