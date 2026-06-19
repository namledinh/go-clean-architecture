package business

import (
	"context"
	"github.com/google/uuid"
	"http_api/modules/item/model"
)


type IGetStorage interface {
	GetParameter(
		ctx context.Context,
		id uuid.UUID,
	) (*model.Parameter, error)
}

type iGetUsecase struct {
	store IGetStorage
}

func NewGetUsecase(store IGetStorage) *iGetUsecase {
	return &iGetUsecase{
		store: store,
	}
}

func (biz *iGetUsecase) GetParameterByID(
	ctx context.Context,
	id uuid.UUID,
) (*model.Parameter, error) {
	if id == uuid.Nil {
		return nil, model.ErrInvalidID
	}

	parameter, err := biz.store.GetParameter(ctx, id)
	if err != nil {
		return nil, err
	}

	if parameter == nil {
		return nil, model.ErrDataNotFound
	}

	return parameter, nil
}
