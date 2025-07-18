package main

import (
	"flag"
	"log"
	"mcpg/server"
	"os"
)

func main() {
	db_url, var_exists := os.LookupEnv("DB_URL")
	if !var_exists {
		db_url_ptr := flag.String("db-url", "", "URL to connect to the cluster db")
		flag.Parse()
		db_url = *db_url_ptr

		if db_url == "" {
			log.Fatalf("Variable DB_URL or option --db-url must be set")
		}
	}

	err := server.Start(db_url)
	if err != nil {
		log.Fatalf("error while starting the MCPG server : %v", err)
	}
}
