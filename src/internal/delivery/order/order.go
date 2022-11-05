package order

import (
	"ecommerce/internal/entity"
	"ecommerce/internal/usecase/order"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Order(usecase order.OrderUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var order entity.Order
		err := c.ShouldBind(&order)
		// check if input not valid
		if (err != nil) {
			c.JSON(http.StatusBadRequest, err.Error())
			return 
		}

		order, err = usecase.Order(order)
		// check if internal server error
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// return response succeed
		c.JSON(http.StatusOK, order)
	}
}