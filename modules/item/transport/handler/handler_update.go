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

func UpdateParameterByID(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		parseId, err := uuid.Parse(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid id",
			})
			return
		}

		var reqInfo model.ParameterUpdateRequest
		if err := c.ShouldBindJSON(&reqInfo); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}
		
		store := storage.NewSqlStore(db)
		uc := business.NewUpdateUsecase(store)
		
		if err := uc.UpdateParameterById(c.Request.Context(), parseId, &reqInfo); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse("successfully updated"))
	}
}