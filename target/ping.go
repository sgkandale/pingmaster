package target

// Ping is the response of each call to the target
type Ping struct {
	// TargetKey is the pool key of target
	TargetKey string `json:"target_key,omitempty"`

	// Timestamp in seconds when the ping was initiated
	Timestamp int64 `json:"-"`

	TimestampStr string `json:"timestamp,omitempty"`

	// Duration is in milliseconds
	Duration int `json:"duration,omitempty"`

	// Status Code of HTTP response
	StatusCode int `json:"status_code,omitempty"`

	// Error while calling the target
	Error error `json:"error,omitempty"`
}
