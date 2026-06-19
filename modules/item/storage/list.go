package storage

import (
	"context"
	"http_api/modules/item/model"
	"http_api/common"
	"time"
)

func (s *sqlStore) ListParameters(
	ctx context.Context,
	filter *model.Filter,
	pagination *common.Paging,
	moreKeys ...string,
) ([]model.Parameter, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second * 10)
	defer cancel()

	var result []model.Parameter

	if f := filter; f != nil {
		if f.Path != "" {
			s.db = s.db.Where("path LIKE ?", "%"+f.Path+"%")
		}
		if f.DataType != "" {
			s.db = s.db.Where("type = ?", f.DataType)
		}
		if f.Status != "" {
			s.db = s.db.Where("status = ?", f.Status)
		}
	}

	if err := s.db.Table(model.Parameter{}.TableName()).Count(&pagination.Total).Error; err != nil {
		return nil, err
	}

	if err := s.db.Table(model.Parameter{}.TableName()).Order("created_at desc").
		Offset((pagination.Page - 1) * pagination.Limit).
		Limit(pagination.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}