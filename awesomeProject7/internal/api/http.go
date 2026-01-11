package api

import "awesomeProject7/internal/service"

type HTTPHandler struct {
	service *service.RatesService
}

func NewHTTPHandler(service *service.RatesService) *HTTPHandler {
	return &HTTPHandler{service: service}
}
