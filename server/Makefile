.PHONY: generate run migrate-up migrate-down migrate-status migrate-new install generate-mocks generate-models db-setup

generate:
	gqlgen generate
	@node scripts/merge-schema.js

mockgen:
	go generate ./repository/...
	mockgen -destination=repository/mocks/mock_sqlmock.go -package=mocks github.com/DATA-DOG/go-sqlmock Sqlmock

run:
	go run server.go

migrate-up:
	sql-migrate up

migrate-down:
	sql-migrate down

migrate-status:
	sql-migrate status

migrate-new:
	@read -p "Enter migration name: " name; \
	timestamp=$$(date +%Y%m%d%H%M%S); \
	sql-migrate new $${timestamp}_$$name

install:
	go mod download
	go install github.com/99designs/gqlgen@latest
	go install github.com/rubenv/sql-migrate/...@latest
	go install go.uber.org/mock/mockgen@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go get github.com/DATA-DOG/go-sqlmock

generate-models:
	sqlc generate

db-setup:
	@echo "Running migrations..."
	make migrate-up
	@echo "Generating models from schema..."
	make generate-models
	@echo "Database setup complete!"
