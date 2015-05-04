package pings

import (
	"time"

	"github.com/etcinit/central/app/entities"
	"github.com/etcinit/central/app/influxdb"
	"github.com/influxdb/influxdb/client"
	"github.com/jacobstr/confer"
)

// Pusher provides functions for sending ping data to the main InfluxDB server.
type Pusher struct {
	Connections *influxdb.ConnectionsService `inject:""`
	Config      *confer.Config               `inject:""`
}

// Push sends ping data to the main InfluxDB server.
func (p *Pusher) Push(ping *entities.Ping) error {
	batch := client.BatchPoints{
		RetentionPolicy: "default",
		Database:        p.Config.GetString("influxdb.database"),
		Points: []client.Point{
			client.Point{
				Name: "pings",
				Tags: map[string]string{
					"instance_id": ping.InstanceID,
					"ip":          ping.IP,
				},
				Fields: map[string]interface{}{
					"instance_id": ping.InstanceID,
					"message":     ping.Message,
					"ip":          ping.IP,
				},
				Timestamp: time.Now(),
				Precision: "s",
			},
		},
	}

	_, err := p.Connections.Make().Write(batch)

	return err
}
