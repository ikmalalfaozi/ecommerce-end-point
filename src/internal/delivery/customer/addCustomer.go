package customer

import (
	"ecommerce/internal/entity"
	"ecommerce/internal/usecase/customer"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddCustomer(usecase customer.CustomerUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var customer entity.Customer
		err := c.ShouldBind(&customer)
		// check if input not valid
		if (err != nil) {
			c.JSON(http.StatusBadRequest, err.Error())
			return 
		}

		customer, err = usecase.AddCustomer(customer)
		// check if internal server error
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// return response succeed
		c.JSON(http.StatusOK, customer)
	}
}