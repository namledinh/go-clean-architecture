package business

import (
	"context"
	"http_api/common"
	"http_api/modules/item/model"
	"strings"
)


type ICreateStorage interface {
	CreateParameter(
		ctx context.Context,
		parameter *model.Parameter,
	) error
}


type iCreateUsecase struct {
	store ICreateStorage
}

func NewCreateUsecase(store ICreateStorage) *iCreateUsecase {
	return &iCreateUsecase{
		store: store,
	}
}

func (biz *iCreateUsecase) InsertParameter(
	ctx context.Context,
	parameter *model.Parameter,
) error {
	path := strings.TrimSpace(parameter.Path)
	if path == "" {
		return common.ErrBadRequest("path is empty")
	}

	dataType := strings.TrimSpace(parameter.DataType)
	if dataType == "" {
		return common.ErrBadRequest("data type is empty")
	}

	if err := biz.store.CreateParameter(ctx, parameter); err != nil {
		return common.ErrCannotCreateEntity("Parameter", err)
	}

	return nil
}
