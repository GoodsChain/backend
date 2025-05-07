package logger

import (
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// InitLogger initializes the global zerolog logger.
// It sets the log level based on the LOG_LEVEL environment variable.
// Supported levels: "debug", "info", "warn", "error", "fatal", "panic".
// Defaults to "info" if LOG_LEVEL is not set or invalid.
func InitLogger() {
	logLevelStr := strings.ToLower(os.Getenv("LOG_LEVEL"))
	var level zerolog.Level

	switch logLevelStr {
	case "debug":
		level = zerolog.DebugLevel
	case "info":
		level = zerolog.InfoLevel
	case "warn":
		level = zerolog.WarnLevel
	case "error":
		level = zerolog.ErrorLevel
	case "fatal":
		level = zerolog.FatalLevel
	case "panic":
		level = zerolog.PanicLevel
	default:
		level = zerolog.InfoLevel // Default level
	}

	zerolog.SetGlobalLevel(level)

	// Use console writer for human-readable output during development
	// In a production environment, you might want to use JSON output
	// and send logs to a centralized logging system.
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log.Logger = log.Output(output).With().Timestamp().Logger()

	log.Info().Msg("Logger initialized")
	if logLevelStr != "" && level.String() != logLevelStr {
		log.Warn().Str("requestedLogLevel", logLevelStr).Str("actualLogLevel", level.String()).Msg("Invalid LOG_LEVEL specified, using default or inferred level.")
	} else {
		log.Info().Str("logLevel", level.String()).Msg("Log level set")
	}
}
