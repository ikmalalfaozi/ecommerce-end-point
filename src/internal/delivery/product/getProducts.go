package product

import (
	"ecommerce/internal/usecase/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProducts(usecase product.ProductUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		page := 1
		// check if query page exist
		if (c.Query("page") != "") {
			var err error;
			page, err = strconv.Atoi(c.Query("page"))
			// check if query page not a number
			if (err != nil) {
				c.JSON(http.StatusBadRequest, "Query page must be a positive integer")
				return	
			}
		}

		limit := 10

		result, err := usecase.GetProducts(page, limit)
		// check if internal server error 
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// return response succeed
		c.JSON(http.StatusOK, result)
	}
}