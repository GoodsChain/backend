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

## Key Dependencies
- github.com/gin-gonic/gin - Web framework
- github.com/jmoiron/sqlx - Database interaction
- github.com/google/uuid - UUID generation
- github.com/lib/pq - PostgreSQL driver

## Tool Usage Patterns
- `make run` - Start the development server
- `go run main.go` - Directly run the application
- Proper error handling with context propagation
- Structured logging (not yet implemented)
