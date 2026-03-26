package main

import (
	"fmt"
	"hospital-api/config"
	"hospital-api/internal/database"
	"hospital-api/internal/server"
	"log"
)

func main() {
	cfg := config.Load()

	db, err := database.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	r := server.New(db)

	fmt.Println("Server started")
	r.Run(":" + cfg.Port)
}
