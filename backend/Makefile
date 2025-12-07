.PHONY: generate run migrate-up migrate-down migrate-status migrate-new install

generate:
	gqlgen generate
	@node scripts/merge-schema.js

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
	sql-migrate new $$name

install:
	go mod download
	go install github.com/99designs/gqlgen@latest
	go install github.com/rubenv/sql-migrate/...@latest
