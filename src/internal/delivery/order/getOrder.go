package order

import (
	"ecommerce/internal/usecase/order"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOrder(usecase order.OrderUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customer_id, err := strconv.ParseInt(ctx.Param("customer_id"), 10, 64)
		// check if param order id is not a number
		if (err != nil) {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		order_id, err := strconv.ParseInt(ctx.Param("order_id"), 10, 64)
		// check if param order id is not a number
		if (err != nil) {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		
		order, err := usecase.GetOrder(customer_id, order_id)
		// check if internal server error
		if (err != nil) {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// return response succeed
		ctx.JSON(http.StatusOK, order)
	}
}