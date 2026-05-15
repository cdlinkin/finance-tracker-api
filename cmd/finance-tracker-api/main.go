package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/cdlinkin/finance-tracker-api/internal/config"
	"github.com/cdlinkin/finance-tracker-api/internal/handler"
	"github.com/cdlinkin/finance-tracker-api/internal/repository"
	"github.com/cdlinkin/finance-tracker-api/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file")
	}
	storage := flag.String("storage", "memory", "storage type: memory or postgres")
	flag.Parse()

	var r repository.Repository

	cfg := config.ConfigLoad()

	source := fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=%s sslmode=disable",
		cfg.User, cfg.Password, cfg.DBName)

	switch *storage {
	case "postgres":
		db, err := sqlx.Connect("postgres", source)
		if err != nil {
			log.Fatal("failed to connect to postgres:", err)
		}
		defer db.Close()
		fmt.Println("storage: PostgreSQL")
		r = repository.NewPostgresRepo(db)
	default:
		fmt.Println("storage: in-memory")
		r = repository.NewTransactionRepo()
	}

	s := service.NewTransactionService(r)
	h := handler.NewTransactionHandler(s)

	chiRouter := chi.NewRouter()

	chiRouter.Use(middleware.Logger)
	chiRouter.Use(middleware.Recoverer)

	chiRouter.Post("/transaction", h.Create)
	chiRouter.Get("/transaction", h.GetAll)
	chiRouter.Get("/transaction/summary", h.Summary)
	chiRouter.Delete("/transaction/{id}", h.Delete)

	fmt.Println("server running")
	if err := http.ListenAndServe(":3000", chiRouter); err != nil {
		log.Fatal(err)
	}
}
