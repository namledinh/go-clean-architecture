package storage

import (
	"context"
	"github.com/google/uuid"
	"http_api/modules/item/model"
	"time"
)

func (s *sqlStore) GetParameter(
	ctx context.Context,
	id uuid.UUID,
) (*model.Parameter, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second * 10)
	defer cancel()

	var parameter model.Parameter

	if err := s.db.WithContext(ctx).Where("id = ?", id).First(&parameter).Error; err != nil {
		return nil, err
	}

	return &parameter, nil
}