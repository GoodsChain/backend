# Technical Context

## Technology Stack
- **Language**: Go (Golang)
- **Web Framework**: Gin (Gin-Gonic)
- **Database**: PostgreSQL
- **Database Driver**: github.com/jmoiron/sqlx
- **Routing**: Gin's native routing
- **Error Handling**: Standard HTTP status codes with consistent JSON error responses

## Development Setup
- **Dependencies**: Managed through Go modules
- **Database Connection**: Configured through environment variables
- **Build Tool**: Makefile for development tasks
- **Code Organization**: Layered architecture with clear separation of concerns

## Technical Constraints
- Must follow Clean Architecture principles
- All database interactions must go through repository layer
- Business logic must be encapsulated in usecase layer
- Handler layer must remain thin, focusing only on HTTP concerns

## Operational Details
- **Request Tracking**: Implemented through middleware with X-Request-ID header
- **Connection Pooling**: Configured via sqlx with max open/idle connections and lifetime
- **Graceful Shutdown**: Handles SIGINT/SIGTERM with configurable timeout
- **API Versioning**: Implemented through router groups (e.g., /v1 prefix)
- **Configuration**: Centralized config with environment variable fallbacks

## Key Dependencies
- github.com/gin-gonic/gin - Web framework
- github.com/jmoiron/sqlx - Database interaction
- github.com/google/uuid - UUID generation
- github.com/lib/pq - PostgreSQL driver
- github.com/rs/zerolog - Structured logging
- github.com/swaggo/gin-swagger - Swagger UI for Gin
- github.com/swaggo/files - Swagger UI static files
- github.com/swaggo/swag - Swagger core library and CLI tool

## Tool Usage Patterns
- `make run` - Directly run the application
- `make db` - Create PostgreSQL database docker container
- `make migrate-up` - Apply database migration
- `make mock` - Generate mock for testing
- `make test` - Run unit test
- `swag init` - Generate/update Swagger API documentation
- Proper error handling with context propagation
- Structured logging implemented using `zerolog`
