package model

// ErrorResponse represents a generic error response.
type ErrorResponse struct {
	Code    string `json:"code" example:"NOT_FOUND" description:"Error code for programmatic handling"`
	Message string `json:"message" example:"Resource not found" description:"Human-readable error description"`
}

// SuccessResponse represents a generic success response.
// Often used for operations like DELETE where no other data is returned.
type SuccessResponse struct {
	Message string `json:"message" example:"Operation successful" description:"Success message"`
}

// PaginatedResponse represents a paginated list response
type PaginatedResponse struct {
	Data       interface{} `json:"data" description:"List of items"`
	TotalCount int         `json:"total_count" example:"100" description:"Total number of items"`
	PageSize   int         `json:"page_size" example:"10" description:"Number of items per page"`
	Page       int         `json:"page" example:"1" description:"Current page number"`
	TotalPages int         `json:"total_pages" example:"10" description:"Total number of pages"`
}
