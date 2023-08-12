package routes

import (
	"github.com/amosehiguese/stock/app/api"
	"github.com/go-chi/chi/v5"
)


func PrivateRoutes(r *chi.Mux, api *api.Handler) {
	r.Group(func(r chi.Router) {
		r.Post("/api/stocks", api.CreateStock)
		r.Patch("/api/stock/{id}", api.UpdateStock)
		r.Delete("/api/stock/{id}", api.DeleteStock)

		r.Post("/api/sign-out", api.SignOut)

		r.Get("/api/user/{id}", api.GetUser)
		r.Get("/api/user/{id}/portfolio", api.GetUserPortfolio)
		r.Patch("/api/edit", api.UpdateUserProfile)
		r.Delete("/api/user/{id}", api.DeleteUser) //admin-level
	})
}


func PublicRoutes(r *chi.Mux, api *api.Handler) {
	r.Group(func(r chi.Router) {
		r.Post("/api/sign-in", api.SignIn)
		r.Post("/api/sign-up", api.SignUp)

		r.Get("/api/stocks", api.GetStocks)
		r.Get("/api/stock/{id}", api.RetrieveStock)
	})

}
