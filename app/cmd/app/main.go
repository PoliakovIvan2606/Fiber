package main

import (
	"github.com/PoliakovIvan2606/Fiber/internal/app"
	"github.com/PoliakovIvan2606/Fiber/internal/books"
	"github.com/PoliakovIvan2606/Fiber/internal/config"
	"github.com/PoliakovIvan2606/Fiber/pkg/db/postgres"
	"context"
	"fmt"
	"time"
)

func main() {
	cfg := config.MastLoad()

	app := app.NewApp()

	pgDsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.Postgres.DB_USER,
		cfg.Postgres.DB_PASS,
		cfg.Postgres.DB_HOST,
		cfg.Postgres.DB_PORT,
		cfg.Postgres.DB_NAME,
	)
	

	pgClient, err := postgres.NewClient(context.Background(), 5, 3*time.Second, pgDsn, false)
	if err != nil {
		panic(err)
	}
	defer pgClient.Close()


	repositoryBook := books.NewBookRepository(pgClient)
	usecaseBook := books.NewUseCase(repositoryBook)
	bookRputer := books.NewBookRouter(app, usecaseBook)
	bookRputer.Register()


	app.Listen(cfg.HttpServer.Host + cfg.HttpServer.Port)
}