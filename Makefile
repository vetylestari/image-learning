#!make
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

default_migration_cmd := -dir database/migrations -table ${MIGRATION_TABLE_NAME} postgres "host=${DB_HOST} port=${DB_PORT} user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_DATABASE} sslmode=disable search_path=${DB_SCHEMA}"
default_seeder_cmd := -no-versioning -dir database/seeders -table ${MIGRATION_TABLE_NAME} postgres "host=${DB_HOST} port=${DB_PORT} user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_DATABASE} sslmode=disable search_path=${DB_SCHEMA}"

module-new:
	touch domain/$$name.go && mkdir -p $$name/delivery/http && mkdir -p $$name/repositories && mkdir -p $$name/usecases && touch $$name/root.go
	
run:
	go run main.go

dev:
	air

init-schema:
	go run console/command/command.go init-schema

build:
	go build -o go-starter-template main.go

migrate-up:
	goose $(default_migration_cmd) up

migrate-up-by-one:
	goose $(default_migration_cmd) up-by-one

migrate-up-to:
	goose $(default_migration_cmd) up-to $$version

migrate-down:
	goose $(default_migration_cmd) down

migrate-down-to:
	goose $(default_migration_cmd) down-to $$version

migrate-redo:
	goose $(default_migration_cmd) redo

migrate-reset:
	goose $(default_migration_cmd) reset

migrate-status:
	goose $(default_migration_cmd) status

migrate-version:
	goose $(default_migration_cmd) version

migrate-create:
	goose $(default_migration_cmd) create $$name sql

migrate-fix:
	goose $(default_migration_cmd) fix

seed-up:
	goose $(default_seeder_cmd) up

seed-up-by-one:
	goose $(default_seeder_cmd) up-by-one

seed-up-to:
	goose $(default_seeder_cmd) up-to $$version

seed-down:
	goose $(default_seeder_cmd) down

seed-down-to:
	goose $(default_seeder_cmd) down-to $$version

seed-redo:
	goose $(default_seeder_cmd) redo

seed-reset:
	goose $(default_seeder_cmd) reset

seed-status:
	goose $(default_seeder_cmd) status

seed-version:
	goose $(default_seeder_cmd) version

seed-create:
	goose $(default_seeder_cmd) create $$name sql

seed-fix:
	goose $(default_seeder_cmd) fix