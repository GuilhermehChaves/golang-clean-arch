package model

type ErrorResponse struct {
	//*RestResponse
	Message string `json:"message"`
	Details string `json:"details"`
}

func NewErrorResponse(message, details string) *ErrorResponse {
	return &ErrorResponse{
		//RestResponse: NewRestResponse(status, time.Now()),
		Message: message,
		Details: details,
	}
}
