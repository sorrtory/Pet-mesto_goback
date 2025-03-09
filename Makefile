APP_NAME = mesto

all: run

build:
	go build -o dist/$(APP_NAME) src/cmd/main.go

run: build
	./dist/$(APP_NAME)


prod:
	docker compose --file compose.yaml --project-name 'mesto-goback' -d up --build

stop:
	docker compose --file compose.yaml --project-name 'mesto-goback' down
