package server

import (
	"log"
	"net/http"
	"time"

	"github.com/amosehiguese/stock/app/config"
	"github.com/amosehiguese/stock/routes"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Unable to load env variables ->", err)
	}
}

func NewHttpServer(config config.Config) *http.Server {
	r := chi.NewRouter()

	useCors(r)
	useMiddlewares(r)

	routes.PublicRoutes(r)
	routes.PrivateRoutes(r)
	routes.AdminRoutes(r)

	return &http.Server{
		Addr: config.Address,
		ReadTimeout: time.Duration(config.ReadTimeout),
		WriteTimeout: time.Duration(config.WriteTimeout),
		MaxHeaderBytes: 1 << 20,
	}
}

func useCors(r *chi.Mux) {
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"htts://api/v1/*", "http://api/v1/*"},
		AllowedMethods: []string{"GET", "POST", "PUT","PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: true,
		MaxAge: 300,
	}))
}

func useMiddlewares(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.CleanPath)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
}




