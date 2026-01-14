package service

import (
	"context"
	"errors"
)

func (s *RatesService) EnsureCoinsExist(
	ctx context.Context,
	titles []string,
) error {

	if len(titles) == 0 {
		return errors.New("no coin titles provided")
	}

	for _, title := range titles {
		if title == "" {
			return errors.New("coin title is empty")
		}

		// Проверка через хранилище
		// Если монеты нет — Storage вернёт ошибку
		_, err := s.storage.GetLast(ctx, []string{title})
		if err != nil {
			return err
		}
	}

	return nil
}
