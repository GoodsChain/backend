# Tech Context

## Technologies
- **Language**: Go (1.24.2)
- **Database**: PostgreSQL 17.4
- **Drivers**: lib/pq for SQL connectivity
- **Migrations**: golang-migrate/migrate

## Development Tools
- **Build**: Makefile-based workflows
- **Testing**: Go test ecosystem
- **Formatting**: gofmt, goimports
- **Dependencies**: Go modules (go.mod)

## Technical Constraints
- Must support ACID transactions
- Immutable transaction records
- Cross-region deployment readiness
- GDPR compliance for data storage
