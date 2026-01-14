package service

import (
	"context"

	"awesomeProject7/internal/entities"
)

type RatesClient interface {
	GetRates(
		ctx context.Context,
		titles []string,
	) ([]entities.Coin, error)
}
