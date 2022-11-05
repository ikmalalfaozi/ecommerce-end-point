package main

import (
	"ecommerce/internal/delivery"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Hello World")
	})

	delivery.Router(v1)

	r.Run();
}