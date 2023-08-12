package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/amosehiguese/stock/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load env variables")
	}
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	routes.PublicRoutes(r)

	fmt.Println("Server listenig on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
