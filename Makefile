-include .env

# Database migration targets
migrate-up:
	migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)" up

migrate-down:
	migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)" down

# Database container targets
db:
	docker run --name goodschain-db \
		-e POSTGRES_USER=$(DB_USER) \
		-e POSTGRES_PASSWORD=$(DB_PASSWORD) \
		-e POSTGRES_DB=$(DB_DB) \
		-p $(DB_PORT):5432 \
		-d \
		postgres:17.4-alpine3.21

db-down:
	docker stop goodschain-db
	docker rm goodschain-db

run:
	go run main.go

# Mock generation and testing targets
mock:
	mockgen -destination=mock/customer_repository_mock.go -package=mock github.com/GoodsChain/backend/repository CustomerRepository
	mockgen -destination=mock/customer_usecase_mock.go -package=mock github.com/GoodsChain/backend/usecase CustomerUsecase
	mockgen -destination=mock/supplier_repository_mock.go -package=mock github.com/GoodsChain/backend/repository SupplierRepository
	mockgen -destination=mock/supplier_usecase_mock.go -package=mock github.com/GoodsChain/backend/usecase SupplierUsecase
	mockgen -destination=mock/car_repository_mock.go -package=mock github.com/GoodsChain/backend/repository CarRepository
	mockgen -destination=mock/car_usecase_mock.go -package=mock github.com/GoodsChain/backend/usecase CarUsecase

test:
	go test -v -cover ./... -count=1
