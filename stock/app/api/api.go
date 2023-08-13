package api

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

type Logic interface{}

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
