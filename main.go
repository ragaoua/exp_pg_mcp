package main

import (
	"log"

	"mcpg/server"
)

func main() {
	err := server.Start()
	if err != nil {
		log.Fatalf("error while starting the MCPG server : %v", err)
	}
}
