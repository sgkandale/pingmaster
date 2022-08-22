package server

const (
	ResponseStatus_Error   = "error"
	ResponseStatus_Success = "success"

	ResponseMessage_ReadRequestError    = "failed to read request"
	ResponseMessage_UsernameExistsError = "user name already exists"
)

type ServerResponse struct {
	// Status of the response
	Status string `json:"status,omitempty"`

	// Message to provide additional details of the response
	Message string `json:"message,omitempty"`

	// Response is the response of the server
	// when it has to some data
	Response interface{} `json:"response,omitempty"`
}
