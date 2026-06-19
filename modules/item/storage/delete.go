package storage

import (
	"context"
	"github.com/google/uuid"
	"http_api/modules/item/model"
	"time"
)

func (s *sqlStore) DeleteParameter(
	ctx context.Context,
	id uuid.UUID,
) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second * 10)
	defer cancel()

	if err := s.db.WithContext(ctx).Table(model.Parameter{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
		return err
	}

	return nil
}