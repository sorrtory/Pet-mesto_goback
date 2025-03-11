package main

import (
	"log"
	"mesto-goback/cmd/api"
	"mesto-goback/internal/common"
	"mesto-goback/internal/db"
)

func main() {
	// Init DB connectdion
	DB_USER := common.GetEnv("POSTGRES_USER")
	DB_PSWD := common.GetEnv("POSTGRES_PASSWORD")

	store, err := db.NewConnection(DB_USER, DB_PSWD)
	if err != nil {
		log.Fatalf("Can't connect to DB: %v\n", err)
	}
	log.Println("Connected to DB")

	// Serve API
	a := api.NewAPI("localhost", "8080", store)
	a.Serve()

}
