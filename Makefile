# Load .env variables
ifneq (,$(wildcard ./.env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

APP_NAME = mesto
.PHONY: run, prod, build


all: run

# Recompile
build:
	go build -o bin/$(APP_NAME) cmd/main.go

# Run on localhost with already started DB
run: build
	BACKEND_PUBLIC=web/public POSTGRES_HOST=localhost BACKEND_HOST=localhost BACKEND_PORT=8080 ./bin/$(APP_NAME)

shutdown:
	docker compose down

up-db:
	docker compose up -d --build 'db'

down-db:
	docker compose down 'db'

# Clear DB and up it
restart-db: shutdown up-db migrate-down migrate-up

# Install migrate tool (~/go/bin/migrate by default)
migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate-create:
	~/go/bin/migrate create -seq -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

# cmd/migrate/migrate.go does the same as "migrate up/down"
migrate-up:
	MIGRATIONS_PATH=file://cmd/migrate/migrations go run cmd/migrate/migrate.go up

migrate-down:
	MIGRATIONS_PATH=file://cmd/migrate/migrations go run cmd/migrate/migrate.go down

# Use any migratinon command with "make migrate <cmds...>"
migrate:
	~/go/bin/migrate -source file://cmd/migrate/migrations -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):5432/postgres?sslmode=disable" $(filter-out $@,$(MAKECMDGOALS))


# Start containers
prod:
	docker compose up -d --build
