package order

import (
	"ecommerce/internal/usecase/order"
	"net/http"

	"github.com/gin-gonic/gin"
)


type changeOrderRequest struct{
	OrderID int64 `json:"order_id"`
	Status string `json:"status"`
}

func ChangeOrderStatus(usecase order.OrderUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req changeOrderRequest
		err := ctx.ShouldBind(&req)
		// check if input is not valid
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		order, err := usecase.ChangeOrderStatus(req.OrderID, req.Status)
		// check if internal server error
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// return response succeed
		ctx.JSON(http.StatusOK, order)
	}
}