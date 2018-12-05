package utils

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
}

func NewErrorResponse(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
	}
}
