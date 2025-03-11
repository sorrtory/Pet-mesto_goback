# Load .env variables
ifneq (,$(wildcard ./.env))
    include .env
    export $(shell sed 's/=.*//' .env)
endif

APP_NAME = mesto

all: run

build:
	go build -o bin/$(APP_NAME) cmd/main.go

run: build
	./bin/$(APP_NAME)

prod:
	docker compose --file deployment/compose.yaml --project-name 'mesto-goback' up -d --build

down:
	docker compose --file deployment/compose.yaml down

up-db:
	docker compose --file deployment/compose.yaml up -d --build 'db'

# Install migrate tool (~/go/bin/migrate by default)
migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate-create:
	~/go/bin/migrate create -seq -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

# cmd/migrate/main.go does the same as "migrate up/down"
migrate-up:
	go run cmd/migrate/main.go up

migrate-down:
	go run cmd/migrate/main.go down

# Use any migrate command with "make migrate <cmds...>"
migrate:
	~/go/bin/migrate -source file://cmd/migrate/migrations -database "postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/postgres?sslmode=disable" $(filter-out $@,$(MAKECMDGOALS))


