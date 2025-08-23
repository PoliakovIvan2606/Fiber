package main

import (
	"github.com/PoliakovIvan2606/Fiber/internal/config"
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var migrationsPath, migrationsTable string

	flag.StringVar(&migrationsPath, "migrations-path", "", "path to migrations")
	flag.StringVar(&migrationsTable, "migrations-table", "migrations", "name of migrations table")
	flag.Parse()

	cfg := config.MastLoad()

	if migrationsPath == "" {
		panic("migrations-path is required")
	}

	// строка подключения через pgx
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&x-migrations-table=%s",
		cfg.Postgres.DB_USER, cfg.Postgres.DB_PASS, cfg.Postgres.DB_HOST, 
		cfg.Postgres.DB_PORT, cfg.Postgres.DB_NAME, migrationsTable,
	)

	m, err := migrate.New(
		"file://"+migrationsPath,
		dsn,
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")
			return
		}
		panic(err)
	}

	fmt.Println("migrations applied")
}
