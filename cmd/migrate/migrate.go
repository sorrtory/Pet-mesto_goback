package main

import (
	"log"
	"mesto-goback/internal/common"
	"mesto-goback/internal/db"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	ALLOW_MIGRATION := common.GetEnv("ALLOW_MIGRATION")
	if ALLOW_MIGRATION == "true" || ALLOW_MIGRATION == "yes" || ALLOW_MIGRATION == "y" {
		// Init DB connectdion
		DB_HOST := common.GetEnv("POSTGRES_HOST")
		DB_USER := common.GetEnv("POSTGRES_USER")
		DB_PSWD := common.GetEnv("POSTGRES_PASSWORD")

		MIGRATIONS_PATH := common.GetEnv("MIGRATIONS_PATH")

		store, err := db.NewConnection(DB_HOST, DB_USER, DB_PSWD)
		if err != nil {
			log.Fatalf("Can't connect to DB: %v\n", err)
		}
		log.Println("Connected to DB")

		// Init migration connection
		driver, err := postgres.WithInstance(store.DB, &postgres.Config{})
		if err != nil {
			log.Fatalf("Can't connect to driver: %v\n", err)
		}

		m, err := migrate.NewWithDatabaseInstance(
			MIGRATIONS_PATH,
			"postgres", driver)

		if err != nil {
			log.Fatalf("Can't connect to DB in migration: %v\n", err)
		}

		// Handle the flags
		cmd := os.Args[len(os.Args)-1]
		if cmd == "up" {
			if err := m.Up(); err != nil && err != migrate.ErrNoChange {
				log.Fatal(err)
			}
		} else if cmd == "down" {
			if err := m.Down(); err != nil && err != migrate.ErrNoChange {
				log.Fatal(err)
			}
		} else {
			log.Fatalln("Use up/down argument")
		}
	}

}
