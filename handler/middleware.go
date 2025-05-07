package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrorResponse represents a standardized error response.
type ErrorResponse struct {
	Error string `json:"error"`
}

// ErrorHandlingMiddleware is a Gin middleware for centralized error handling.
// It catches errors that occur during request processing and formats them into
// a standardized JSON response.
func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Process request

		// After request processing, check for errors
		if len(c.Errors) > 0 {
			// For now, we take the last error.
			// More sophisticated error handling can be added here,
			// such as logging, distinguishing error types, etc.
			err := c.Errors.Last()

			// Determine the status code.
			// If the error is of a type that has a status code, use it.
			// Otherwise, default to 500 Internal Server Error.
			// This part can be expanded to handle custom error types.
			statusCode := http.StatusInternalServerError
			if ginErr, ok := err.Err.(*gin.Error); ok {
				if ginErr.IsType(gin.ErrorTypePublic) {
					// For public errors, we might want to use a different status code
					// or expose more details, but for now, we keep it simple.
				}
			}
			
			// If a status code was set on the context previously (e.g. by c.AbortWithStatusJSON),
			// and it's an error code, prefer that.
			if c.Writer.Status() >= 400 {
				statusCode = c.Writer.Status()
			}

			c.JSON(statusCode, ErrorResponse{Error: err.Error()})
			return // Stop further processing if error handled
		}

		// If no errors, but status code indicates an error (e.g. 404 Not Found from router)
		// and no response has been written yet.
		// This is a basic way to catch 404s that didn't hit a specific route handler
		// and ensure they also get a JSON response.
		// Note: Gin's default 404 handler might already send a plain text response.
		// For full control, one might replace Gin's NoRoute handler.
		if c.Writer.Status() == http.StatusNotFound && !c.Writer.Written() {
			c.JSON(http.StatusNotFound, ErrorResponse{Error: "resource not found"})
			return
		}
	}
}
