package business

import (
	"context"
	"github.com/google/uuid"
	"http_api/modules/item/model"
)


type IDeleteStorage interface {
	GetParameter(
		ctx context.Context,
		id uuid.UUID,
	) (*model.Parameter, error)

	DeleteParameter(
		ctx context.Context,
		id uuid.UUID,
	) error
}


type iDeleteUsecase struct {
	store IDeleteStorage
}

func NewDeleteUsecase(store IDeleteStorage) *iDeleteUsecase {
	return &iDeleteUsecase{
		store: store,
	}
}

func (biz *iDeleteUsecase) DeleteParameterByID(
	ctx context.Context,
	id uuid.UUID,
) error {
	if id == uuid.Nil {
		return model.ErrInvalidID
	}
	parameter, err := biz.store.GetParameter(ctx, id)
	if err != nil {
		return err
	}

	if parameter == nil {
		return model.ErrDataNotFound
	}
	
	if err := biz.store.DeleteParameter(ctx, id); err != nil {
		return err
	}

	return nil
}
