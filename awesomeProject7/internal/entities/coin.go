package entities

import (
	"errors"
	"time"
)

type Coin struct {
	title     string
	rate      float64
	createdAt time.Time
}

func NewCoin(
	title string,
	rate float64,
	createdAt time.Time,
) (Coin, error) {

	if title == "" {
		return Coin{}, errors.New("title is required")
	}

	if rate <= 0 {
		return Coin{}, errors.New("rate must be greater than zero")
	}

	if createdAt.IsZero() {
		return Coin{}, errors.New("createdAt is required")
	}

	return Coin{
		title:     title,
		rate:      rate,
		createdAt: createdAt,
	}, nil
}
