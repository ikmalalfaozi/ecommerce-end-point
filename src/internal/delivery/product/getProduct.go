package product

import (
	"ecommerce/internal/entity"
	"ecommerce/internal/usecase/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProduct(usecase product.ProductUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		// check if input is not number
		if err != nil {
			c.JSON(http.StatusBadRequest, "Param id must be a positive integer")
			return
		}

		result, err := usecase.GetProduct(id)
		// check if internal server error
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		
		// check if product not found
		if (result == entity.Product{}) {
			c.JSON(http.StatusBadRequest, "Id product not found")
			return
		}

		// return response succeed
		c.JSON(http.StatusOK, result)
	}
}