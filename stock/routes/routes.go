package routes

import (
	"github.com/amosehiguese/stock/app/api"
	"github.com/go-chi/chi/v5"
)



func PublicRoutes(r *chi.Mux) {
	r.Route("/api/v1",func(r chi.Router) {
		r.Post("/sign-in", api.SignIn)
		r.Post("/sign-up", api.SignUp)

		r.Get("/stocks", api.GetStocks)
		r.Get("/stock/{id}", api.RetrieveStock)
	})

}


func PrivateRoutes(r *chi.Mux) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Use(api.AuthMiddleware)

		r.Post("/buy", api.BuyStock)
		r.Post("/sell", api.SellStock)

		r.Get("/user/{id}/portfolio", api.GetUserPortfolio)
		r.Patch("/user/{id}/edit", api.UpdateUserProfile)
		r.Post("/sign-out", api.SignOut)
	})
}

func AdminRoutes(r *chi.Mux) {
	r.Route("/api/v1/admin", func(r chi.Router) {

		r.Use(api.AdminMiddleware)

		r.Post("/stocks", api.CreateStock)  // admin-level
		r.Patch("/stock/{id}", api.UpdateStock) //admin-level
		r.Delete("/stock/{id}", api.DeleteStock)  //admin-level

		r.Get("/user/{id}", api.GetUser)
		r.Get("/users", api.GetAllUsers)
		r.Delete("/user/{id}", api.DeleteUser) //admin-level
	})
}


