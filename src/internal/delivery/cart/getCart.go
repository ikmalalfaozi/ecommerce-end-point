package cart

import (
	"ecommerce/internal/usecase/cart"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCart(usecase cart.CartUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customer_id, err := strconv.ParseInt(ctx.Param("customer_id"), 10, 64)
		// check if param customer id is not a number
		if (err != nil) {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		
		product_id, err := strconv.ParseInt(ctx.Param("product_id"), 10, 64)
		// check if param product id is not a number
		if (err != nil) {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		
		cart, err := usecase.GetCart(customer_id, product_id)
		// check if internal server error
		if (err != nil) {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// return response succeed
		ctx.JSON(http.StatusOK, cart)
	}
}