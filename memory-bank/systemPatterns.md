# System Patterns

## Architecture Overview
- Hexagonal architecture with domain/core layer separation
- REST API layer using Go net/http
- Database layer using SQL drivers

## Key Patterns
- Repository pattern for data access
- Service layer for business logic
- Middleware for cross-cutting concerns
- Event sourcing for transaction history

## Component Relationships
- API handlers -> Service layer -> Repositories -> Database
- Domain entities <-> Value objects <-> Aggregates

## Critical Paths
- Goods tracking workflow
- Transaction recording pipeline
- Inventory synchronization process
