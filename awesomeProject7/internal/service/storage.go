package service

import (
	"awesomeProject7/internal/entities"
	"context"
)

type Storage interface {
	Store(ctx context.Context, coin entities.Coin) error

	GetLast(
		ctx context.Context,
		title []string,
	) (entities.Coin, error)

	GetMax(
		ctx context.Context,
		title []string,
	) (entities.Coin, error)

	GetMin(
		ctx context.Context,
		title []string,
	) (entities.Coin, error)
}

// слайсы, для пачек запросов
