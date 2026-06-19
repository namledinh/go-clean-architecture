package business

import (
	"context"
	"http_api/modules/item/model"
	"http_api/common"
)


type IListStorage interface {
	ListParameters(
		ctx context.Context,
		filter *model.Filter,
		pagination *common.Paging,
		moreKeys ...string,
	) ([]model.Parameter, error)
}

type iListUsecase struct {
	store IListStorage
}

func NewListUsecase(store IListStorage) *iListUsecase {
	return &iListUsecase{
		store: store,
	}
}

func (biz *iListUsecase) ListParameters(
	ctx context.Context,
	filter *model.Filter,
	pagination *common.Paging,
) ([]model.Parameter, error) {

	data, err := biz.store.ListParameters(ctx, filter, pagination)
	if err != nil {
		return nil, err
	}
	
	return data, nil
}
