package api

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
