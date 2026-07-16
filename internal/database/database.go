package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func Connect() *pgxpool.Pool {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load .env:", err)
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is empty")
	}

	pool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	return pool
}