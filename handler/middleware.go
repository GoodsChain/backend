package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// ErrorResponse represents a standardized error response.
type ErrorResponse struct {
	Error string `json:"error"`
}

// ErrorHandlingMiddleware is a Gin middleware for centralized error handling.
// It catches errors that occur during request processing and formats them into
// a standardized JSON response. It also logs requests and errors.
func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		latency := time.Since(start)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path
		userAgent := c.Request.UserAgent()

		logEvent := log.Info() // Default to Info for successful requests
		if statusCode >= 400 && statusCode < 500 {
			logEvent = log.Warn()
		} else if statusCode >= 500 {
			logEvent = log.Error()
		}

		// After request processing, check for errors
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			// Log the error with more details
			logEvent.Err(err.Err). // Use err.Err to get the actual error interface
				Str("method", method).
				Str("path", path).
				Str("clientIP", clientIP).
				Int("statusCode", statusCode).
				Dur("latency", latency).
				Str("userAgent", userAgent).
				Msg("Request error")

			// Determine the status code for the response.
			// If a status code was set on the context previously (e.g. by c.AbortWithStatusJSON),
			// and it's an error code, prefer that. Otherwise, default to 500.
			responseStatusCode := http.StatusInternalServerError
			if c.Writer.Status() >= 400 {
				responseStatusCode = c.Writer.Status()
			}
			
			// Ensure the logged status code matches the response status code if it was an error
			if statusCode < 400 { // If Gin didn't set an error status, but we have c.Errors
				statusCode = responseStatusCode // Update statusCode for logging consistency
			}


			c.JSON(responseStatusCode, ErrorResponse{Error: err.Error()})
			return // Stop further processing if error handled
		}

		// Log the request
		logEvent.
			Str("method", method).
			Str("path", path).
			Str("clientIP", clientIP).
			Int("statusCode", statusCode).
			Dur("latency", latency).
			Str("userAgent", userAgent).
			Msg("Request processed")


		// If no errors from c.Errors, but status code indicates an error (e.g. 404 Not Found from router)
		// and no response has been written yet.
		if statusCode == http.StatusNotFound && !c.Writer.Written() {
			log.Warn().
				Str("method", method).
				Str("path", path).
				Str("clientIP", clientIP).
				Int("statusCode", statusCode).
				Dur("latency", latency).
				Str("userAgent", userAgent).
				Msg("Resource not found (404)")
			c.JSON(http.StatusNotFound, ErrorResponse{Error: "resource not found"})
			return
		}
	}
}
