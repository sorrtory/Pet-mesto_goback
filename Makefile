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


