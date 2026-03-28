package database

import (
	"fmt"
	"hospital-api/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	var db *gorm.DB
	var err error

	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn))
		if err == nil {
			fmt.Println("Connected to database")
			return db, nil
		}

		fmt.Printf("waiting for DB (%d/%d)\n", i+1, maxRetries)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("failed to connect to database after %d attempts: %w", maxRetries, err)
}
