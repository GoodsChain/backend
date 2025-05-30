# System Patterns

## Architecture Overview
The system implements Clean Architecture principles with a clear separation of concerns into distinct layers:

1. **Handler Layer** - Entry point for HTTP requests
   - Located in `handler/` directory
   - Handles HTTP routing and request/response formatting
   - Converts HTTP requests to domain models and vice versa

2. **Usecase Layer** - Business logic layer
   - Located in `usecase/` directory
   - Contains interface definitions and implementation
   - Orchestrates business rules and operations

3. **Repository Layer** - Data access layer
   - Located in `repository/` directory
   - Handles database operations
   - Implements data access interfaces

## Key Implementation Patterns
1. **Dependency Injection** - Usecase layer is injected into handlers, repository layer is injected into usecases
2. **Interface-based Design** - Clear interfaces between layers for loose coupling
3. **Error Handling** - Consistent error responses with appropriate HTTP status codes
4. **Validation** - Input validation at handler layer before business logic execution
5. **Graceful Shutdown** - Proper signal handling and resource cleanup during application termination
6. **API Versioning** - Implemented through router groups (e.g., /v1 prefix)
7. **Request Tracking** - Request IDs propagated through the system for observability
8. **Connection Pooling** - Configurable database connection pooling parameters
9. **Middleware Organization** - Global middleware registered at engine level, route-specific at group level

## Component Relationships
- Handler depends on Usecase interfaces
- Usecase implementation depends on Repository interfaces
- Repository implementation depends on database driver

## Server Lifecycle Management
1. **Initialization**
   - Configuration loading
   - Database connection
   - Dependency initialization and injection
   - HTTP server configuration

2. **Operation**
   - Non-blocking server start using goroutines
   - Signal handling for graceful termination

3. **Termination**
   - Dedicated function for graceful shutdown (handleGracefulShutdown)
   - Signal-triggered shutdown (SIGINT, SIGTERM)
   - Graceful shutdown with timeout for in-flight requests
   - Ordered resource cleanup (HTTP server, database connections)
   - Context cancellation for coordinated shutdown
