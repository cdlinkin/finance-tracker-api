package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cdlinkin/finance-tracker-api/internal/handler"
	"github.com/cdlinkin/finance-tracker-api/internal/repository"
	"github.com/cdlinkin/finance-tracker-api/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := repository.NewTransactionRepo()
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
