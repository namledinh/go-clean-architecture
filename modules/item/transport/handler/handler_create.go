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

func CreateParameter(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var reqInfo model.ParameterCreateRequest
		if err := c.ShouldBindJSON(&reqInfo); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		parameter := model.Parameter{
			Path:        reqInfo.Path,
			DataType:    reqInfo.DataType,
			Description: reqInfo.Description,
			Status:     model.ItemStatusENABLE,
			UpdatedBy:   reqInfo.UpdatedBy,
		}
		store := storage.NewSqlStore(db)
		uc := business.NewCreateUsecase(store)

		if err := uc.InsertParameter(c.Request.Context(), &parameter); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(parameter.Id))
	}
}