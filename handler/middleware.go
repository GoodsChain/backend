package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	appErrors "github.com/GoodsChain/backend/errors"
	"github.com/GoodsChain/backend/model"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// ErrorHandlingMiddleware is a Gin middleware for centralized error handling.
// It catches errors that occur during request processing and formats them into
// a standardized JSON response. It also logs requests and errors.
func ErrorHandlingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set request ID for better tracing
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = generateRequestID()
			c.Request.Header.Set("X-Request-ID", requestID)
			c.Writer.Header().Set("X-Request-ID", requestID)
		}

		// Record start time
		start := time.Now()

		// Create a logger with the request ID
		contextLogger := log.With().Str("request_id", requestID).Logger()
		
		// Process request
		c.Next()

		// Collect metrics
		latency := time.Since(start)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		method := c.Request.Method
		path := c.Request.URL.Path
		userAgent := c.Request.UserAgent()

		// Determine log level based on status code
		var logEvent *zerolog.Event
		if statusCode >= 500 {
			logEvent = contextLogger.Error()
		} else if statusCode >= 400 {
			logEvent = contextLogger.Warn()
		} else {
			logEvent = contextLogger.Info()
		}

		// After request processing, check for errors
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			
			// Determine appropriate response based on error type
			var appError *appErrors.AppError
			responseStatusCode := http.StatusInternalServerError
			errorCode := string(appErrors.ErrInternal)
			errorMessage := "Internal Server Error"
			
			// Check if the error is an AppError
			if errors.As(err, &appError) {
				// Use the HTTP status code and error details from the AppError
				responseStatusCode = appError.HTTPCode
				errorCode = string(appError.Code)
				errorMessage = appError.Message
			} else if c.Writer.Status() >= 400 {
				// If the handler already set an error status code, use that
				responseStatusCode = c.Writer.Status()
			}

			// Log the error with details
			logEvent.Err(err).
				Str("method", method).
				Str("path", path).
				Str("clientIP", clientIP).
				Int("statusCode", responseStatusCode).
				Dur("latency", latency).
				Str("userAgent", userAgent).
				Str("error_code", errorCode).
				Msg(errorMessage)

			// Return standardized error response
			c.JSON(responseStatusCode, model.ErrorResponse{
				Code:    errorCode,
				Message: errorMessage,
			})
			return // Stop further processing if error handled
		}

		// Log successful request
		logEvent.
			Str("method", method).
			Str("path", path).
			Str("clientIP", clientIP).
			Int("statusCode", statusCode).
			Dur("latency", latency).
			Str("userAgent", userAgent).
			Msg("Request processed")

		// Handle 404 Not Found if no response has been written yet
		if statusCode == http.StatusNotFound && !c.Writer.Written() {
			contextLogger.Warn().
				Str("method", method).
				Str("path", path).
				Str("clientIP", clientIP).
				Int("statusCode", statusCode).
				Dur("latency", latency).
				Str("userAgent", userAgent).
				Msg("Resource not found")
				
			c.JSON(http.StatusNotFound, model.ErrorResponse{
				Code:    string(appErrors.ErrNotFound),
				Message: "Resource not found",
			})
		}
	}
}

// generateRequestID creates a unique request ID
// In a production environment, this could use a more sophisticated
// approach like UUID generation
func generateRequestID() string {
	now := time.Now().UnixNano()
	return fmt.Sprintf("%d", now)
}
