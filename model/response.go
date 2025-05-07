package model

// ErrorResponse represents a generic error response.
type ErrorResponse struct {
	Error string `json:"error" example:"Error message description"`
}

// SuccessResponse represents a generic success response.
// Often used for operations like DELETE where no other data is returned.
type SuccessResponse struct {
	Message string `json:"message" example:"Operation successful"`
}
