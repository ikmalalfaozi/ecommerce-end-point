package cart

import (
	"ecommerce/internal/usecase/cart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderCartRequest struct {
	CustomerID int64 `json:"customer_id"`
	ProductID  int64 `json:"product_id"`
}

func OrderProduct(usecase cart.CartUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var orderCartRequest OrderCartRequest
		err := ctx.ShouldBind(&orderCartRequest)
		// check if input is not valid
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		order, err := usecase.OrderProduct(orderCartRequest.CustomerID, orderCartRequest.ProductID)
		// check if internal server error
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// return response succeed
		ctx.JSON(http.StatusOK, order)
	}
}