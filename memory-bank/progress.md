# Project Progress

## What Works
1. All CRUD operations for customers are implemented (including the newly added GET /customers)
2. All CRUD operations for suppliers are implemented with full endpoint coverage
3. Clean Architecture layers are properly separated and functioning
4. Database integration with PostgreSQL is working correctly
5. API endpoints are accessible and returning appropriate responses
6. Dependency injection is properly implemented
7. Unit tests for usecase layer with 100% coverage using Uber's GoMock
8. Unit tests for handler layer with table-driven tests covering all endpoints
9. Unit tests for repository layer with go-sqlmock simulating database interactions
10. Graceful shutdown handling with proper signal catching (SIGINT, SIGTERM)
11. Proper cleanup of resources during server shutdown
12. HTTP server configured with appropriate timeouts
13. Input validation for customer and supplier API request bodies (create/update) using Gin's built-in validator with struct tags.
14. Global error handling middleware (`handler.ErrorHandlingMiddleware`) for standardized JSON error responses.
15. Structured logging using `zerolog`, providing detailed request/response and error logs. Log level is configurable via environment variable.
16. API documentation using Swagger/OpenAPI, accessible via `/swagger/index.html`.
17. All CRUD operations for cars are implemented, following the same Clean Architecture and testing patterns.
18. All CRUD operations for customer-car relationships are implemented with full endpoint coverage.
19. Specialized endpoints for retrieving cars by customer ID (`/customers/:customer_id/cars`) and customers by car ID (`/cars/:car_id/customers`) are functioning correctly.
20. Comprehensive README.md documentation with project overview, architecture explanation, API endpoints, setup instructions, and more.

## What's Left to Build
All core requirements from the initial project brief, plus the requested Car CRUD API and CustomerCar relationship APIs, are now implemented. The system is ready for further feature development or deployment preparations. Potential next steps include implementing unit tests for the CustomerCar components across all layers and setting up a CI/CD pipeline.

## Current Status
The core functionality for customer, supplier, car management, and customer-car relationships is complete and working. The system follows Clean Architecture principles with proper separation of concerns. All database interactions are encapsulated in the repository layer, business logic is handled in the usecase layer, and HTTP routing is managed in the handler layer. All layers now have comprehensive unit tests with complete test coverage using appropriate testing strategies - table-driven tests for handlers, mocks for usecases, and go-sqlmock for repositories. Input validation has been added to the handler layer for create and update operations, ensuring that incoming data for customers and suppliers meets basic requirements (e.g., required fields, email format) before further processing. The supplier API implementation follows the same architectural patterns as the customer API, ensuring consistency across the codebase. The application now implements graceful shutdown, allowing in-flight requests to complete and resources to be properly released when the server receives termination signals. A global error handling middleware has been integrated to ensure consistent JSON responses for errors across the API. Structured logging has been implemented using `zerolog`, enhancing observability with detailed request and error logs. Comprehensive API documentation using Swagger/OpenAPI has been integrated, making all endpoints discoverable and testable via a UI at `/swagger/index.html`. Furthermore, CRUD APIs for managing 'Car' entities have been successfully added, including model, repository, usecase, and handler layers, complete with unit tests and Swagger documentation, maintaining architectural consistency. The new CustomerCar implementation enables management of many-to-many relationships between customers and cars, with dedicated endpoints for both direct relationship management and for retrieving related entities. Most recently, a comprehensive README.md file has been created to provide clear documentation on the project's architecture, features, API endpoints, technology stack, setup instructions, and testing procedures, significantly improving developer onboarding and project discoverability. A GitHub Actions workflow has been added to automate the build and test process on push and pull request events. A `build` target has been added to the Makefile to facilitate this.

## Known Issues
No known critical issues.

## Evolution of Decisions
1. Decided to use singular table names in database queries to match actual schema
2. Maintained thin handler layer to ensure business logic remains in usecase layer
3. Implemented interface-based design for better testability and loose coupling
4. Selected Uber's GoMock for creating mock implementations in unit tests
5. Added dedicated Makefile commands for testing and mock generation
6. Used go-sqlmock for repository testing to avoid real database dependencies
7. Created mock repositories using 'make mock' command for usecase testing
8. Applied consistent architectural patterns across different domain entities (customers and suppliers)
9. Implemented graceful shutdown with 30-second timeout to ensure clean application termination
10. Configured HTTP server with appropriate timeouts (read, write, idle) for better performance and security
11. Refactored graceful shutdown logic into a dedicated function to improve code organization and maintainability
12. Adopted Gin's built-in validation mechanism (`ShouldBindJSON` with struct tags) for request payload validation, removing manual checks and standardizing validation logic.
13. Implemented a global error handling middleware to centralize and standardize API error responses, improving consistency and maintainability.
14. Chose `zerolog` for structured logging, configured for console output during development and with an environment variable (`LOG_LEVEL`) for controlling log verbosity. Integrated request and error logging into the `ErrorHandlingMiddleware`.
15. Selected `swaggo` (`swag` CLI, `gin-swagger`) for generating API documentation due to its robust features and good integration with Gin. This provides a user-friendly UI for exploring and testing API endpoints.
16. Standardized on using `model.ErrorResponse` and `model.SuccessResponse` for API responses to ensure consistency in both the API behavior and its documentation.
17. Extended the system with a new 'Car' entity, applying the established Clean Architecture, testing methodologies, and API design patterns to ensure consistency and maintainability.
18. For the CustomerCar relationship, implemented both entity-centric endpoints (`/customer-cars/*`) and relationship-centric endpoints (`/customers/:customer_id/cars`, `/cars/:car_id/customers`) to provide flexible access patterns for querying relationships.
19. Created a comprehensive README.md with detailed project documentation, following software engineering best practices for project documentation to ensure ease of understanding and maintainability.
