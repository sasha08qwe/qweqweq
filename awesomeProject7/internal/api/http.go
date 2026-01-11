package api

import (
	"errors"

	"awesomeProject7/internal/service"
)

type HTTPHandler struct {
	service *service.RatesService
}

func NewHTTPHandler(service *service.RatesService) (*HTTPHandler, error) {
	if service == nil {
		return nil, errors.New("rates service is nil")
	}

	return &HTTPHandler{
		service: service,
	}, nil
}
