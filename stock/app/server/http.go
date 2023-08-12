package server

import (
	"log"
	"net/http"
	"time"

	"github.com/amosehiguese/stock/app/api"
	"github.com/amosehiguese/stock/app/config"
	"github.com/amosehiguese/stock/routes"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to load env variables")
	}
}

func NewHttpServer(config config.Config) *http.Server {
	api := api.NewHandler()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	routes.PublicRoutes(r, api)
	routes.PrivateRoutes(r, api)

	return &http.Server{
		Addr: config.Address,
		ReadTimeout: time.Duration(config.ReadTimeout),
		WriteTimeout: time.Duration(config.WriteTimeout),
		MaxHeaderBytes: 1 << 20,
	}
}


type response struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

func NewResponse(success bool, data any) *response {
	return &response{
		Success: success,
		Data:    data,
	}
}

type errorResp struct {
	Success   bool   `json:"success"`
	ErrorCode int    `json:"error"`
	Message   string `json:"message"`
}

func NewError(success bool, errCode int, msg string) *errorResp {
	return &errorResp{
		Success:   success,
		ErrorCode: errCode,
		Message:   msg,
	}
}

