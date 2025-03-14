package main

import (
	"log"
	"mesto-goback/cmd/server"
	"mesto-goback/internal/common"
	"mesto-goback/internal/db"
)

func main() {
    // Disable DB reset on container restart
    common.SetEnv("ALLOW_MIGRATION", "no")

	// Init DB connectdion
	DB_HOST := common.GetEnv("POSTGRES_HOST")
	DB_USER := common.GetEnv("POSTGRES_USER")
	DB_PSWD := common.GetEnv("POSTGRES_PASSWORD")

	store, err := db.NewConnection(DB_HOST, DB_USER, DB_PSWD)
	if err != nil {
		log.Fatalf("Can't connect to DB: %v\n", err)
	}
	log.Println("Connected to DB")

	// Serve API
	BE_HOST := common.GetEnv("BACKEND_HOST")
	BE_PORT := common.GetEnv("BACKEND_PORT")

	a := server.NewServer(BE_HOST, BE_PORT, store)
	a.Serve()

}
