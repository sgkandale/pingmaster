package target

// Ping is the response of each call to the target
type Ping struct {
	// Timestamp in seconds when the ping was initiated
	Timestamp int64

	// Duration is in milliseconds
	Duration int

	// Status Code of HTTP response
	StatusCode int

	// Error while calling the target
	Error error
}
