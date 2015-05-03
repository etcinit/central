package v1

// PingJSON contains the JSON input of a POST /v1/ping call.
type PingJSON struct {
	InstanceID string `json:"name" binding:"required" valid:"" validate:"required,max=128"`
	Message    string `json:"message" binding:"required" validate:"max=255"`
}

// LogEntryJSON contains the JSON input of a POST /v1/logs call.
type LogEntryJSON struct {
	InstanceID string   `json:"instanceName" binding:"required" validate:"required,max=128"`
	Name       string   `json:"filename" binding:"required" validate:"required,max=128"`
	Lines      []string `json:"lines" binding:"required" validate:"required,min=1"`
}
