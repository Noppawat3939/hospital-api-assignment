package main

import (
	"fmt"
	"hospital-api/config"
	"hospital-api/internal/database"
	"hospital-api/internal/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "local"
	}

	// load env on local only
	if env == "local" {
		if err := godotenv.Load(".env.local"); err != nil {
			log.Println(".env.local file found")
		}
	}

	err := godotenv.Load(".env.local")

	fmt.Println("Running in environment:", env)

	// load configs
	cfg := config.Load()

	// connect db with retry
	db, err := database.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// init server
	srv := server.New(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	// start server
	if err := srv.Start(port); err != nil {
		log.Fatal(err)
	}
}
