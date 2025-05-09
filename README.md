# GoodsChain Backend

A robust RESTful API backend for supply chain management built with Clean Architecture principles in Go.

## Overview

GoodsChain Backend is a Go-based RESTful API service that provides comprehensive management of customers, suppliers, cars, and the relationships between customers and cars. The application follows Clean Architecture principles and SOLID design patterns to ensure maintainability, testability, and scalability.

## Features

- **Customer Management**: Full CRUD operations for customer data
- **Supplier Management**: Full CRUD operations for supplier data
- **Car Management**: Full CRUD operations for car data
- **Customer-Car Relationship Management**: Manage associations between customers and cars
- **Clean Architecture**: Clear separation of concerns with handler, usecase, and repository layers
- **PostgreSQL Integration**: Reliable data persistence with PostgreSQL
- **Input Validation**: Request payload validation using Gin's built-in validator
- **Structured Logging**: Comprehensive logging with zerolog
- **API Documentation**: Interactive API documentation with Swagger/OpenAPI
- **Graceful Shutdown**: Proper handling of termination signals
- **Comprehensive Testing**: Unit tests with high coverage across all layers
- **Structured Error Handling**: Consistent error responses with error codes
- **Request Tracking**: Request IDs for tracing requests through logs
- **CI/CD**: GitHub Actions workflow for automated build and test
- **Connection Pooling**: Configurable database connection pool
- **API Versioning**: Versioned API endpoints for backward compatibility
- **Health Check Endpoint**: Dedicated endpoint for monitoring service health

## Architecture

The application implements Clean Architecture with the following layers:

1. **Handler Layer** (Delivery/API Layer)
   - HTTP request/response handling
   - Input validation
   - Route definitions

2. **Usecase Layer** (Business Logic Layer)
   - Core business rules and logic
   - Orchestration of data operations

3. **Repository Layer** (Data Access Layer)
   - Database interactions
   - Data persistence

## API Endpoints

### Health Check
- `GET /health` - API health check endpoint

### Customer Endpoints (prefixed with API version, e.g., `/v1`)
- `POST /v1/customers` - Create a new customer
- `GET /v1/customers` - List all customers
- `GET /v1/customers/:id` - Get customer by ID
- `PUT /v1/customers/:id` - Update customer by ID
- `DELETE /v1/customers/:id` - Delete customer by ID
- `GET /v1/customers/:customer_id/cars` - Get all cars owned by a customer

### Supplier Endpoints
- `POST /v1/suppliers` - Create a new supplier
- `GET /v1/suppliers` - List all suppliers
- `GET /v1/suppliers/:id` - Get supplier by ID
- `PUT /v1/suppliers/:id` - Update supplier by ID
- `DELETE /v1/suppliers/:id` - Delete supplier by ID

### Car Endpoints
- `POST /v1/cars` - Create a new car
- `GET /v1/cars` - List all cars
- `GET /v1/cars/:id` - Get car by ID
- `PUT /v1/cars/:id` - Update car by ID
- `DELETE /v1/cars/:id` - Delete car by ID
- `GET /v1/cars/:car_id/customers` - Get all customers who own a specific car

### Customer-Car Relationship Endpoints
- `POST /v1/customer-cars` - Create a new customer-car relationship
- `GET /v1/customer-cars` - List all customer-car relationships
- `GET /v1/customer-cars/:id` - Get customer-car relationship by ID
- `PUT /v1/customer-cars/:id` - Update customer-car relationship by ID
- `DELETE /v1/customer-cars/:id` - Delete customer-car relationship by ID

### Documentation
- `GET /swagger/*any` - Swagger UI for API documentation and testing

## Technologies Used

- **Language**: Go (Golang)
- **Web Framework**: Gin (Gin-Gonic)
- **Database**: PostgreSQL
- **Database Driver**: github.com/jmoiron/sqlx
- **UUID Generation**: github.com/google/uuid
- **Logging**: github.com/rs/zerolog
- **API Documentation**: github.com/swaggo/gin-swagger
- **Testing**: Standard library testing, github.com/golang/mock, github.com/DATA-DOG/go-sqlmock

## Getting Started

### Prerequisites

- Go 1.16 or higher
- PostgreSQL
- Make (for using the Makefile commands)

### Configuration

The application uses environment variables for configuration with intelligent defaults:

- Database settings:
  - `DB_HOST` - PostgreSQL host (default: localhost)
  - `DB_PORT` - PostgreSQL port (default: 5432)
  - `DB_USER` - PostgreSQL username (default: postgres)
  - `DB_PASSWORD` - PostgreSQL password
  - `DB_NAME` - PostgreSQL database name (default: goodschain)
  - `DB_SSL_MODE` - PostgreSQL SSL mode (default: disable)
  - `DB_MAX_OPEN_CONNS` - Maximum number of open connections (default: 25)
  - `DB_MAX_IDLE_CONNS` - Maximum number of idle connections (default: 5)
  - `DB_CONN_MAX_LIFE` - Maximum connection lifetime in seconds (default: 300)

- API settings:
  - `API_PORT` - Application port (default: 3000)
  - `API_VERSION` - API version for URL prefix (default: v1)
  - `API_READ_TIMEOUT` - HTTP read timeout in seconds (default: 15)
  - `API_WRITE_TIMEOUT` - HTTP write timeout in seconds (default: 15)
  - `API_IDLE_TIMEOUT` - HTTP idle timeout in seconds (default: 60)
  - `API_SHUTDOWN_TIMEOUT` - Graceful shutdown timeout in seconds (default: 30)

You can set these in a `.env` file or directly in your environment.

### Running the Application

1. Clone the repository
2. Set up environment variables
3. Run the application:

```bash
# Run directly
go run main.go

# Or use the Makefile
make run
```

### API Documentation

The API documentation is available at `/swagger/index.html` when the application is running.

## Testing

The project includes comprehensive unit tests for all layers:

```bash
# Run all tests
make test

# Generate mocks (required before running tests for the first time)
make mock
```

## Project Structure

```
├── config/             # Configuration handling
├── docs/               # Swagger documentation
├── handler/            # HTTP handlers and routing
├── logger/             # Logging setup
├── migrations/         # Database migration files
├── mock/               # Generated mock implementations
├── model/              # Data models and DTOs
├── repository/         # Data access layer
├── usecase/            # Business logic layer
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
├── main.go             # Application entry point
├── Makefile            # Development task automation
└── README.md
```

## License

This project is licensed under the terms of the license included in the LICENSE file.
