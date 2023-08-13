package routes

import (
	"github.com/amosehiguese/stock/app/api"
	"github.com/go-chi/chi/v5"
)



func PublicRoutes(r *chi.Mux, api *api.Handler) {
	r.Route("/api/v1",func(r chi.Router) {
		r.Post("/sign-in", api.SignIn)
		r.Post("/sign-up", api.SignUp)

		r.Get("/stocks", api.GetStocks)
		r.Get("/stock/{id}", api.RetrieveStock)
	})

}


func PrivateRoutes(r *chi.Mux, api *api.Handler) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Use(AuthMiddleware)

		r.Post("/stocks", api.CreateStock)
		r.Patch("/stock/{id}", api.UpdateStock)
		r.Delete("/stock/{id}", api.DeleteStock)

		r.Post("/sign-out", api.SignOut)

		r.Get("/user/{id}", api.GetUser)
		r.Get("/user/{id}/portfolio", api.GetUserPortfolio)
		r.Patch("/edit", api.UpdateUserProfile)
		r.Delete("/user/{id}", api.DeleteUser) //admin-level
	})


}


