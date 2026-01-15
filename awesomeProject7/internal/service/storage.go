package service

import (
	"awesomeProject7/internal/entities"
	"context"
)

type Storage interface {
	// таблица coins
	CoinExists(ctx context.Context, title string) (bool, error)
	AddCoin(ctx context.Context, title string) error
	GetTitles(ctx context.Context) ([]string, error)

	// таблица rates
	Store(ctx context.Context, coins []entities.Coin) error

	GetLast(ctx context.Context, titles []string) ([]entities.Coin, error)
	GetMax(ctx context.Context, titles []string) ([]entities.Coin, error)
	GetMin(ctx context.Context, titles []string) ([]entities.Coin, error)
}
