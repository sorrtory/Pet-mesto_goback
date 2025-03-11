package main

import (
	"log"
	"mesto-goback/cmd/api"
	"mesto-goback/internal/db"
	"os"
)

func getEnv(env string) string {
	v := os.Getenv(env)
	if v == "" {
		log.Printf("%v isn't set!\n", env)
	}
	return v
}

func main() {
	// Init DB connectdion
	DB_USER := getEnv("POSTGRES_USER")
	DB_PSWD := getEnv("POSTGRES_PASSWORD")

	store, err := db.NewConnection(DB_USER, DB_PSWD)
	if err != nil {
		log.Fatalf("Can't connect to DB: %v\n", err)
	}
	log.Println("Connected to DB")

	// Serve API
	a := api.NewAPI("localhost", "8080", store)
	a.Serve()

}
