package customer

import (
	"ecommerce/internal/usecase/customer"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCustomers(usecase customer.CustomerUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		page := 1
		// check if query page is exist
		if (c.Query("page") != "") {
			var err error;
			page, err = strconv.Atoi(c.Query("page"))
			// check if query page is not a number
			if (err != nil) {
				c.JSON(http.StatusBadRequest, "Query page must be a positive integer")
				return	
			}
		}

		limit := 10

		result, err := usecase.GetCustomers(page, limit)
		// check if internal server error
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// return response succeed
		c.JSON(http.StatusOK, result)
	}
}