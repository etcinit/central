package entities

// LogEntry represents as single log entry.
type LogEntry struct {
	InstanceID string `json:"instance_id"`
	Name       string `json:"name"`
	Line       string `json:"line"`
}
