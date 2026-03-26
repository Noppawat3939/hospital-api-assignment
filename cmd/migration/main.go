package main

import (
	"flag"
	"fmt"
	"hospital-api/config"
	"hospital-api/internal/database"
	"hospital-api/internal/migration"
	"log"
)

/**
	Step migration
	- make migration
	- input file name
**/

func main() {
	fileName := flag.String("file", "", "migration file name")
	flag.Parse()
	fmt.Println("> file name", fileName)

	if *fileName == "" {
		log.Fatalf("file name is required")
	}

	cfg := config.Load()

	db, err := database.New(cfg)
	if err != nil {
		log.Fatalf("Failed connect to database: %v", err)
	}

	fmt.Printf("running migration: %s\n", *fileName)
	migration.RunMigration(db, *fileName)
}
