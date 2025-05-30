-include .env

migrate-up:
	migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)" up

migrate-down:
	migrate -path migrations -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL_MODE)" down

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

build:
	go build -o goodschain main.go

mock-clean:
	rm -rf mock

mock: mock-clean
	mockgen -destination=mock/customer_repository_mock.go -package=mock github.com/GoodsChain/backend/repository CustomerRepository
	mockgen -destination=mock/customer_usecase_mock.go -package=mock github.com/GoodsChain/backend/usecase CustomerUsecase
	mockgen -destination=mock/supplier_repository_mock.go -package=mock github.com/GoodsChain/backend/repository SupplierRepository
	mockgen -destination=mock/supplier_usecase_mock.go -package=mock github.com/GoodsChain/backend/usecase SupplierUsecase
	mockgen -destination=mock/car_repository_mock.go -package=mock github.com/GoodsChain/backend/repository CarRepository
	mockgen -destination=mock/car_usecase_mock.go -package=mock github.com/GoodsChain/backend/usecase CarUsecase
	mockgen -destination=mock/customer_car_repository_mock.go -package=mock github.com/GoodsChain/backend/repository CustomerCarRepository
	mockgen -destination=mock/customer_car_usecase_mock.go -package=mock github.com/GoodsChain/backend/usecase CustomerCarUsecase

test:
	go test -v -cover ./... -count=1
