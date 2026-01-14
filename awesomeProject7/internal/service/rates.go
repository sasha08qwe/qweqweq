package service

import (
	"context"
	"errors"

	"awesomeProject7/internal/entities"
)

type RatesService struct {
	storage Storage
	client  RatesClient
}

func NewRatesService(
	storage Storage,
	client RatesClient,
) (*RatesService, error) {

	if storage == nil {
		return nil, errors.New("storage is nil")
	}

	if client == nil {
		return nil, errors.New("rates client is nil")
	}

	return &RatesService{
		storage: storage,
		client:  client,
	}, nil
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
	title []string,
) (entities.Coin, error) {
	coin, err := s.storage.GetLast(ctx, title)
	if err != nil {
		return entities.Coin{}, err
	}
	return coin, err
}

func (s *RatesService) GetMax(
	ctx context.Context,
	title []string,

) (entities.Coin, error) {
	coin, err := s.storage.GetMax(ctx, title)
	if err != nil {
		return entities.Coin{}, err
	}
	return coin, err
}

func (s *RatesService) GetMin(
	ctx context.Context,
	title []string,
) (entities.Coin, error) {
	coin, err := s.storage.GetMin(ctx, title)
	if err != nil {
		return entities.Coin{}, err
	}
	return coin, err
}

//
// слайсы, ретерниить переменную, функции типа проверки перед походом в бд
