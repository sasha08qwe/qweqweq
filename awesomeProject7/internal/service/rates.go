package service

import (
	"context"
	"time"

	"awesomeProject7/internal/client"
	"awesomeProject7/internal/entities"
	"awesomeProject7/internal/storage"
)

type RatesService struct {
	storage storage.Storage
	client  client.RatesClient
}

func NewRatesService(
	storage storage.Storage,
	client client.RatesClient,
) *RatesService {
	return &RatesService{
		storage: storage,
		client:  client,
	}
}

func (s *RatesService) UpdateRates(ctx context.Context) error {
	titles := []string{"BTC", "ETH"}

	coins, err := s.client.GetRates(ctx, titles)
	if err != nil {
		return err
	}

	for _, coin := range coins {
		if err := s.storage.Store(ctx, coin); err != nil {
			return err
		}
	}

	return nil
}

func (s *RatesService) GetLast(
	ctx context.Context,
	title string,
) (entities.Coin, error) {
	return s.storage.GetLast(ctx, title)
}

func (s *RatesService) GetMax(
	ctx context.Context,
	title string,
	day time.Time,
) (entities.Coin, error) {
	return s.storage.GetMax(ctx, title, day)
}

func (s *RatesService) GetMin(
	ctx context.Context,
	title string,
	day time.Time,
) (entities.Coin, error) {
	return s.storage.GetMin(ctx, title, day)
}
