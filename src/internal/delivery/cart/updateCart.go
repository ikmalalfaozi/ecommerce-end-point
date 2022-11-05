package cart

import (
	"ecommerce/internal/entity"
	"ecommerce/internal/usecase/cart"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateCart(usecase cart.CartUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var cart entity.Cart
		err := ctx.ShouldBind(&cart)
		// check if input is not valid
		if (err != nil) {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		cart, err = usecase.UpdateCart(cart)
		// check if internal server error
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		// return response succeed
		ctx.JSON(http.StatusOK, cart)
	}
}