package business

import (
	"context"
	"github.com/google/uuid"
	"http_api/modules/item/model"
)


type IUpdateStorage interface {
	GetParameter(
		ctx context.Context,
		id uuid.UUID,
	) (*model.Parameter, error)

	UpdateParameter(
		ctx context.Context,
		id uuid.UUID,
		req *model.ParameterUpdateRequest,
	) error
}

type iUpdateUsecase struct {
	store IUpdateStorage
}

func NewUpdateUsecase(store IUpdateStorage) *iUpdateUsecase {
	return &iUpdateUsecase{
		store: store,
	}
}

func (biz *iUpdateUsecase) UpdateParameterById(
	ctx context.Context,
	id uuid.UUID,
	req *model.ParameterUpdateRequest,
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
	
	if err := biz.store.UpdateParameter(ctx, id, req); err != nil {
		return err
	}

	return  nil
}
