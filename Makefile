# This version-strategy uses git tags to set the version string
APP_NAME = sagara_crud
VERSION := $(shell git describe --tags --always --dirty)
GIT_COMMIT := $(shell git rev-list -1 HEAD)

# DATABASE
DB_USER=postgres
DB_PASSWORD=bismillah
DB_HOST=127.0.0.1
DB_PORT=5432
DB_SSL=disable

DB_NAME_AUTH=sagara_crud_auth
DB_NAME_PRODUCT=sagara_crud_product


version: ## Show version
	@echo $(APP_NAME) $(VERSION) \(git commit: $(GIT_COMMIT)\)

local:
	air -c .air.toml

run:
	go run main.go

install:
	go get -u github.com/swaggo/swag/cmd/swag && swag init 
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

setup:
	make migrate-latest

migrate-lastest:
	migrate -source file:./cmd/auth/migrations/ -verbose -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME_AUTH}?sslmode=${DB_SSL} up
	migrate -source file:./cmd/product/migrations/ -verbose -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME_PRODUCT}?sslmode=${DB_SSL} up

migrate-drop:
	migrate -source file:./cmd/auth/migrations/ -verbose -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME_AUTH}?sslmode=${DB_SSL} down
	migrate -source file:./cmd/product/migrations/ -verbose -database postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME_PRODUCT}?sslmode=${DB_SSL} down
