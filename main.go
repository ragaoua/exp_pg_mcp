package main

import (
	"log"
	"mcpg/server"
	"os"
)

func main() {
	db_url, var_exists := os.LookupEnv("DB_URL")
	if !var_exists {
		log.Fatalf("Variable DB_URL must be set")
	}

	err := server.Start(db_url)
	if err != nil {
		log.Fatalf("error while starting the MCPG server : %v", err)
	}
}
