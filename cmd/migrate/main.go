package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Generalsimus/go-boilerplate/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func main() {
	dbURL := config.Cfg.DATABASE_URL
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Tell Goose where your .sql files are stored
	goose.SetBaseFS(os.DirFS("db/migrations"))

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Failed to set dialect: %v", err)
	}

	// Run all pending migrations
	if err := goose.Up(db, "."); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("✅ All database migrations applied successfully!")
}
