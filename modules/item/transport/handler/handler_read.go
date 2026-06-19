package handler

import (
	"net/http"
	"http_api/common"
	"http_api/modules/item/model"
	"github.com/gin-gonic/gin"
	"http_api/modules/item/business"
	"http_api/modules/item/storage"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

func GetParameterByID(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		if _, err := uuid.Parse(id); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		var parameter model.Parameter

		store := storage.NewSqlStore(db)
		uc := business.NewGetUsecase(store)
		
		result, err := uc.GetParameterByID(c.Request.Context(), uuid.MustParse(id))
		if err != nil {
			if err == model.ErrDataNotFound {
				c.JSON(http.StatusNotFound, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		parameter = *result

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(parameter))
	}
}