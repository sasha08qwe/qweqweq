package entities

import (
	"errors"
	"time"
)

type Coin struct {
	Title     string
	Rate      float64
	CreatedAt time.Time
}

func NewCoin(
	title string,
	rate float64,
	createdAt time.Time,
) (*Coin, error) {

	if title == "" {
		return nil, errors.New("title is required")
	}

	if rate <= 0 {
		return nil, errors.New("rate must be greater than zero")
	}

	if createdAt.IsZero() {
		return nil, errors.New("createdAt is required")
	}

	return &Coin{
		Title:     title,
		Rate:      rate,
		CreatedAt: createdAt,
	}, nil
}
