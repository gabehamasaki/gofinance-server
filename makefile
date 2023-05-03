.PHONY: default run build test docs clean

# Variables
APP_NAME=gofinance

# Tasks 
default: run

run:
	@go run ./cmd/server/main.go
build:
	@go build -o ./dist/$(APP_NAME) ./cmd/server/main.go
	@cp ./.env ./dist
run-build:
	@make build
	@./dist/$(APP_NAME)
test:
	@go test ./ ...
clean:
	@rm -rf ./dist
	@rm -rf ./docs