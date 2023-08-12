package api

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

type Logic interface{}
