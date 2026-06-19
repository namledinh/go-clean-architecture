package storage

import (
	"context"
	"time"
	"http_api/common"
	"http_api/modules/item/model"
)

func (s *sqlStore) CreateParameter(
	ctx context.Context,
	parameter *model.Parameter,
) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second * 10)
	defer cancel()

	if err := s.db.WithContext(ctx).Create(parameter).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
