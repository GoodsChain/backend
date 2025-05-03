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

## Component Relationships
- Handler depends on Usecase interfaces
- Usecase implementation depends on Repository interfaces
- Repository implementation depends on database driver
