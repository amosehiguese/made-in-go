package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/amosehiguese/mogo/api"
	_ "github.com/amosehiguese/mogo/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)


func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/users", api.GetAllUsers)
	r.Get("/user/{id}", api.RetriveUser)
	r.Post("/users", api.CreateUser)
	r.Patch("/user/{id}", api.UpdateUser)
	r.Delete("/user/{id}", api.DeleteUser)

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

