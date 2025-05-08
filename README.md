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

### Customer Endpoints
- `POST /customers` - Create a new customer
- `GET /customers` - List all customers
- `GET /customers/:id` - Get customer by ID
- `PUT /customers/:id` - Update customer by ID
- `DELETE /customers/:id` - Delete customer by ID
- `GET /customers/:customer_id/cars` - Get all cars owned by a customer

### Supplier Endpoints
- `POST /suppliers` - Create a new supplier
- `GET /suppliers` - List all suppliers
- `GET /suppliers/:id` - Get supplier by ID
- `PUT /suppliers/:id` - Update supplier by ID
- `DELETE /suppliers/:id` - Delete supplier by ID

### Car Endpoints
- `POST /cars` - Create a new car
- `GET /cars` - List all cars
- `GET /cars/:id` - Get car by ID
- `PUT /cars/:id` - Update car by ID
- `DELETE /cars/:id` - Delete car by ID
- `GET /cars/:car_id/customers` - Get all customers who own a specific car

### Customer-Car Relationship Endpoints
- `POST /customer-cars` - Create a new customer-car relationship
- `GET /customer-cars` - List all customer-car relationships
- `GET /customer-cars/:id` - Get customer-car relationship by ID
- `PUT /customer-cars/:id` - Update customer-car relationship by ID
- `DELETE /customer-cars/:id` - Delete customer-car relationship by ID

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

The application uses environment variables for configuration:

- `DB_HOST` - PostgreSQL host
- `DB_PORT` - PostgreSQL port
- `DB_USER` - PostgreSQL username
- `DB_PASSWORD` - PostgreSQL password
- `DB_NAME` - PostgreSQL database name
- `PORT` - Application port (default: 3000)
- `LOG_LEVEL` - Logging level (default: info)

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
