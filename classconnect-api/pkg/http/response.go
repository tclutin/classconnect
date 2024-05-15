package http

type APIErrorResponse struct {
	Error string `json:"error"`
}

type APIResponse struct {
	Message string `json:"message"`
}

func NewAPIErrorResponse(message string) APIErrorResponse {
	return APIErrorResponse{Error: message}
}

func NewAPIResponse(message string) APIResponse {
	return APIResponse{Message: message}
}
