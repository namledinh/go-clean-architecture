package handler

import (
	"net/http"
	"http_api/common"
	"http_api/modules/item/model"
	"github.com/gin-gonic/gin"
	"http_api/modules/item/business"
	"http_api/modules/item/storage"
	"gorm.io/gorm"
)

func ListParameters(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging
		if err := c.ShouldBindQuery(&paging); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}
		paging.Process()

		var filter model.Filter
		if err := c.ShouldBindQuery(&filter); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		var result []model.Parameter

		store := storage.NewSqlStore(db)
		uc := business.NewListUsecase(store)

		result, err := uc.ListParameters(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, nil))
	}
}