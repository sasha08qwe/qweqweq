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

func (s *RatesService) FetchRates(ctx context.Context) error {
	titles, err := s.storage.GetTitles(ctx)
	if err != nil {
		return err
	}

	if len(titles) == 0 {
		return nil
	}

	coins, err := s.client.GetRates(ctx, titles)
	if err != nil {
		return err
	}

	if len(coins) == 0 {
		return nil
	}

	if err := s.storage.Store(ctx, coins); err != nil {
		return err
	}

	return nil
}

func (s *RatesService) GetLast(
	ctx context.Context,
	titles []string,
) ([]entities.Coin, error) {

	if len(titles) == 0 {
		return nil, errors.New("titles are empty")
	}

	// 1. Проверяем / создаём монеты
	if err := s.ensureCoinsExist(ctx, titles); err != nil {
		return nil, err
	}

	// 2. Получаем данные
	coins, err := s.storage.GetLast(ctx, titles)
	if err != nil {
		return nil, err
	}

	return coins, nil
}

func (s *RatesService) GetMax(
	ctx context.Context,
	titles []string,
) ([]entities.Coin, error) {

	if len(titles) == 0 {
		return nil, errors.New("titles are empty")
	}

	if err := s.ensureCoinsExist(ctx, titles); err != nil {
		return nil, err
	}

	return s.storage.GetMax(ctx, titles)
}

func (s *RatesService) GetMin(
	ctx context.Context,
	titles []string,
) ([]entities.Coin, error) {

	if len(titles) == 0 {
		return nil, errors.New("titles are empty")
	}

	if err := s.ensureCoinsExist(ctx, titles); err != nil {
		return nil, err
	}

	return s.storage.GetMin(ctx, titles)
}

func (s *RatesService) ensureCoinsExist(
	ctx context.Context,
	titles []string,
) error {

	for _, title := range titles {
		if title == "" {
			return errors.New("coin title is empty")
		}

		exists, err := s.storage.CoinExists(ctx, title)
		if err != nil {
			return err
		}

		if !exists {
			if err := s.storage.AddCoin(ctx, title); err != nil {
				return err
			}
		}
	}

	return nil
}
