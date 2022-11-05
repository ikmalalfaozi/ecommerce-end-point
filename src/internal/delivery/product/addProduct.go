package product

import (
	"ecommerce/internal/entity"
	"ecommerce/internal/usecase/product"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddProduct(usecase product.ProductUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var product entity.Product
		err := c.ShouldBind(&product)
		
		// check if input not valid
		if (err != nil) {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}

		product, err = usecase.AddProduct(product)
		// check if Internal Server Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// return response succeed
		c.JSON(http.StatusOK, product)
	}
}