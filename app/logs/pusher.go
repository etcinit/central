package logs

import (
	"time"

	"github.com/etcinit/central/app/entities"
	"github.com/etcinit/central/app/influxdb"
	"github.com/influxdb/influxdb/client"
	"github.com/jacobstr/confer"
	"github.com/kr/pretty"
)

// Pusher provides functions for sending log data to the main InfluxDB server.
type Pusher struct {
	Connections *influxdb.ConnectionsService `inject:""`
	Config      *confer.Config               `inject:""`
}

// Push sends log data to the main InfluxDB server.
func (p *Pusher) Push(entries []*entities.LogEntry) error {
	var points []client.Point

	for _, entry := range entries {
		points = append(
			points,
			client.Point{
				Name: "logs",
				Fields: map[string]interface{}{
					"instance_id": entry.InstanceID,
					"name":        entry.Name,
					"line":        entry.Line,
				},
				Timestamp: time.Now(),
				Precision: "s",
			},
		)
	}

	batch := client.BatchPoints{
		RetentionPolicy: "default",
		Database:        p.Config.GetString("influxdb.database"),
		Points:          points,
	}

	_, err := p.Connections.Make().Write(batch)
	pretty.Println(err)

	return err
}
