package delivery

import (
	cartHandler "ecommerce/internal/delivery/cart"
	customerHandler "ecommerce/internal/delivery/customer"
	orderHandler "ecommerce/internal/delivery/order"
	productHandler "ecommerce/internal/delivery/product"
	cartrepo "ecommerce/internal/repository/cart"
	customerrepo "ecommerce/internal/repository/customer"
	orderrepo "ecommerce/internal/repository/order"
	productrepo "ecommerce/internal/repository/product"
	cartusecase "ecommerce/internal/usecase/cart"
	customerusecase "ecommerce/internal/usecase/customer"
	orderusecase "ecommerce/internal/usecase/order"
	productusecase "ecommerce/internal/usecase/product"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Router(r *gin.RouterGroup) {
	dsn := "root:@tcp(127.0.0.1:3306)/ecommerce-end-point?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
  	if err != nil {
	  log.Fatal("DB connection error")
	}
	
	productRepository := productrepo.NewRepo(db)
	productUsecase := productusecase.NewUseCase(productRepository)
	customerRepository := customerrepo.NewRepo(db)
	customerUsecase := customerusecase.NewUseCase(customerRepository)
	cartRepository := cartrepo.NewRepo(db)
	cartUsecase := cartusecase.NewUseCase(cartRepository)
	orderRepository := orderrepo.NewRepo(db)
	orderUsecase := orderusecase.NewUseCase(orderRepository)
	
	fmt.Println("Database Connected")
	// GET list of product
	r.GET("/products", productHandler.GetProducts(productUsecase))
	
	// GET specific product by id
	r.GET("/products/:id", productHandler.GetProduct(productUsecase))

	// POST new product and add to database
	r.POST("/products", productHandler.AddProduct(productUsecase))

	// GET list of customers
	r.GET("/customers", customerHandler.GetCustomers(customerUsecase))
	
	// GET specific customer by id 
	r.GET("/customers/:id", customerHandler.GetCustomer(customerUsecase))
	
	// POST new customers and add to database
	r.POST("/customers", customerHandler.AddCustomer(customerUsecase))
	
	// Checkout product by product id 
	r.POST("/products/checkout", cartHandler.AddProductToCart(cartUsecase))

	// buy products directly without checkout
	r.POST("/products/buy", orderHandler.Order(orderUsecase))

	// GET list of product in the cart
	r.GET("/carts/:customer_id", cartHandler.GetCarts(cartUsecase))

	// GET product in the chart by id
	r.GET("/carts/:customer_id/:product_id", cartHandler.GetCart(cartUsecase))

	// buy the product in the cart by id
	r.POST("/carts/buy", cartHandler.OrderProduct(cartUsecase))

	// update the number of products in the cart if exist, else add new cart
	r.PUT("/carts", cartHandler.UpdateCart(cartUsecase))
	
	// update order status
	r.PUT("/orders", orderHandler.ChangeOrderStatus(orderUsecase))

	// GET list of orders
	r.GET("/orders/:customer_id", orderHandler.GetOrders(orderUsecase))
	
	// GET specific order by ID
	r.GET("/orders/:customer_id/:order_id", orderHandler.GetOrder(orderUsecase))
}