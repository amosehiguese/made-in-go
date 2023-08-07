package routes

import (
	"github.com/amosehiguese/bookstore/api"
	"github.com/go-chi/chi/v5"
)

func BookStoreRoutes(r *chi.Mux) {
	r.Post("/books", api.CreateBook)
	r.Get("/books", api.GetBooks)
	r.Get("/books/{bookId}", api.RetrieveBook)
	r.Patch("/books/{bookId}", api.UpdateBook)
	r.Delete("/books/{bookId}", api.DeleteBook)
}
