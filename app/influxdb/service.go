package influxdb

import (
	"fmt"
	"log"
	"net/url"

	"github.com/influxdb/influxdb/client"
	"github.com/jacobstr/confer"
)

// ConnectionsService creates and manages connections to the application's main
// influx DB server.
type ConnectionsService struct {
	Config        *confer.Config `inject:""`
	defaultClient *client.Client
}

// Make creates a new connection to InfluxDB.
func (s *ConnectionsService) Make() *client.Client {
	if s.defaultClient != nil {
		return s.defaultClient
	}

	u, err := url.Parse(
		fmt.Sprintf(
			"%s://%s:%d",
			s.Config.GetString("influxdb.scheme"),
			s.Config.GetString("influxdb.host"),
			s.Config.GetInt("influxdb.port"),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	conf := client.Config{
		URL:      *u,
		Username: s.Config.GetString("influxdb.username"),
		Password: s.Config.GetString("influxdb.password"),
	}

	con, err := client.NewClient(conf)
	if err != nil {
		log.Fatal(err)
	}

	_, ver, err := con.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to InfluxDB version:", ver)

	s.defaultClient = con
	return con
}
