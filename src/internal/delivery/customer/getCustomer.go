package customer

import (
	"ecommerce/internal/entity"
	"ecommerce/internal/usecase/customer"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCustomer(usecase customer.CustomerUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		// check if param id is not a number
		if err != nil {
			c.JSON(http.StatusBadRequest, "Param id must be a positive integer")
			return
		}

		result, err := usecase.GetCustomer(id)
		// check if internal server error
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		
		// check if id customer not found
		if (result == entity.Customer{}) {
			c.JSON(http.StatusBadRequest, "Id customer not found")
			return
		}

		// return response succeed
		c.JSON(http.StatusOK, result)
	}
}