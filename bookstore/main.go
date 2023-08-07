package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/amosehiguese/bookstore/routes"
	"github.com/amosehiguese/bookstore/storage"
	"github.com/amosehiguese/bookstore/types"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db := storage.Connect()
	db.AutoMigrate(&types.Book{})
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	routes.BookStoreRoutes(r)

	fmt.Printf("Listening on :8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}
