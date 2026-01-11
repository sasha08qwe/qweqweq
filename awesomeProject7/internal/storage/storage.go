package storage

import (
	"context"
	"time"

	"awesomeProject7/internal/entities"
)

type Storage interface {
	Store(ctx context.Context, coin entities.Coin) error

	GetLast(
		ctx context.Context,
		title string,
	) (entities.Coin, error)

	GetMax(
		ctx context.Context,
		title string,
		day time.Time,
	) (entities.Coin, error)

	GetMin(
		ctx context.Context,
		title string,
		day time.Time,
	) (entities.Coin, error)
}
