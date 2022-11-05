package cart

import (
	"ecommerce/internal/entity"

	"gorm.io/gorm"
)

type CartRepository interface {
	GetCartsByCustomerId(customerId int64) ([]entity.Cart, error)
	GetCartByCustomerIdAndProductId(customerId int64, productId int64) (entity.Cart, error)
	AddProductToCart(cart entity.Cart) (entity.Cart, error)
	FindCustomerById(id int64) (entity.Customer, error)
	FindProductById(id int64) (entity.Product, error)
	OrderProduct(order entity.Order) (entity.Order, error)
	FindCartById(customerid int64, productid int64) (entity.Cart, error)
	UpdateProduct(product entity.Product) (entity.Product, error)
	DeleteCart(cart entity.Cart) (entity.Cart, error)
	UpdateCart(cart entity.Cart) (entity.Cart, error) 
}

type cartRepository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) (cartRepository) {
	// return new cart repository
	return cartRepository{db: db}
}

func (r cartRepository) GetCartsByCustomerId(customerId int64) ([]entity.Cart, error) {
	// get list of carts from a specific customer id
	var carts []entity.Cart
	result := r.db.Find(&carts, &entity.Cart{CustomerID: customerId})
	return carts, result.Error
}

func (r cartRepository) GetCartByCustomerIdAndProductId(customerId int64, productId int64) (entity.Cart, error) {
	// get cart by customer id and product id
	var cart entity.Cart
	result := r.db.Find(&cart, &entity.Cart{CustomerID: customerId, ProductID: productId})
	return cart, result.Error
}

func (r cartRepository) AddProductToCart(cart entity.Cart) (entity.Cart, error) {
	// Add product to cart. If product already exist, then add quantity product in cart, else add new cart
	var oldCart entity.Cart
	result := r.db.Where(&entity.Cart{CustomerID: cart.CustomerID, ProductID: cart.ProductID}).First(&oldCart)
	if (oldCart != entity.Cart{}) {
		oldCart.Quantity += cart.Quantity
		r.db.Where(&entity.Cart{CustomerID: cart.CustomerID, ProductID: cart.ProductID}).Save(&oldCart)
		return oldCart, result.Error
	}

	result2 := r.db.Create(&cart)
	return cart, result2.Error  
}

func (r cartRepository) FindCustomerById(id int64) (entity.Customer, error) {
	// find customer by id
	var customer entity.Customer
	result := r.db.Find(&customer, id)
	return customer, result.Error
}

func (r cartRepository) FindProductById(id int64) (entity.Product, error) {
	// find product by id
	var product entity.Product
	result := r.db.Find(&product, id)
	return product, result.Error
}

func (r cartRepository) OrderProduct(order entity.Order) (entity.Order, error) {
	// order product from cart
	result := r.db.Create(&order)
	return order, result.Error
}

func (r cartRepository) FindCartById(customerid int64, productid int64) (entity.Cart, error) {
	// find cart by id
	var cart entity.Cart
	result := r.db.Where(&entity.Cart{ProductID: productid, CustomerID: customerid}).First(&cart)
	return cart, result.Error
}

func (r cartRepository) UpdateProduct(product entity.Product) (entity.Product, error) {
	// update product
	result := r.db.Save(&product);
	return product, result.Error
}

func (r cartRepository) DeleteCart(cart entity.Cart) (entity.Cart, error) {
	// delete cart
	result := r.db.Where(&entity.Cart{CustomerID: cart.CustomerID, ProductID: cart.ProductID}).Delete(&entity.Cart{})
	return cart, result.Error
}

func (r cartRepository) UpdateCart(cart entity.Cart) (entity.Cart, error) {
	// update cart
	result := r.db.Where(&entity.Cart{CustomerID: cart.CustomerID, ProductID: cart.ProductID}).Save(&cart)
	return cart, result.Error
}
