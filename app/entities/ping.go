package entities

// Ping represents an instance ping.
type Ping struct {
	InstanceID string `json:"instance_id"`
	Message    string `json:"message"`
	IP         string `json:"ip"`
	Version    string `json:"version"`
}
