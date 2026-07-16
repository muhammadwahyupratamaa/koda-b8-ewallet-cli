package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func Connect() *pgx.Conn {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to load .env:", err)
	}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is empty")
	}

	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	return conn
}