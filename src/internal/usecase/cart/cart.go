package cart

import (
	"ecommerce/internal/entity"
	"ecommerce/internal/repository/cart"
	"errors"
	"time"
)

type CartUseCase interface {
	GetCarts(customer_id int64) ([]entity.Cart, error)
	GetCart(customer_id int64, product_id int64) (entity.Cart, error)
	AddProductToCart(cart entity.Cart) (entity.Cart, error)
	OrderProduct(customer_id int64, product_id int64) (entity.Order, error)
	UpdateCart(cart entity.Cart) (entity.Cart, error)
}

type cartUsecase struct {
	repo cart.CartRepository
}

func NewUseCase(repo cart.CartRepository) (cartUsecase) {
	// return new cart usecase
	return cartUsecase{repo: repo}
}

func (u cartUsecase) GetCarts(customer_id int64) ([]entity.Cart, error) {
	// check if customer_id = 0
	if (customer_id == 0) {
		return []entity.Cart{}, nil
	}
	return u.repo.GetCartsByCustomerId(customer_id)
}

func (u cartUsecase) GetCart(customer_id int64, product_id int64) (entity.Cart, error) {
	// check if customer id = 0 or product_id = 0
	if (customer_id == 0 || product_id == 0) {
		return entity.Cart{}, errors.New("cart not found")
	}

	cart, err := u.repo.GetCartByCustomerIdAndProductId(customer_id, product_id)
	// check if err
	if err != nil {
		return entity.Cart{}, err
	}

	// check if cart is not found
	if (cart == entity.Cart{}) {
		return cart, errors.New("cart not found")
	} 
	
	return cart, nil
}

func (u cartUsecase) AddProductToCart(cart entity.Cart) (entity.Cart, error) {
	customer, _ := u.repo.FindCustomerById(cart.CustomerID)
	product, _ := u.repo.FindProductById(cart.ProductID)
	
	// check if customer or product not found in database 
	if (customer == entity.Customer{} || product == entity.Product{}) {
		return entity.Cart{}, errors.New("customer id or product id not found")
	}

	// add product to cart then return added product and error
	return u.repo.AddProductToCart(cart)
}

func (u cartUsecase) OrderProduct(customer_id int64, product_id int64) (entity.Order, error) {
	cart, _ := u.repo.FindCartById(customer_id, product_id)
	product, _ := u.repo.FindProductById(product_id)
	customer, _ := u.repo.FindCustomerById(customer_id)
	
	// Check if cart or product or customer not found in database
	if (cart == entity.Cart{} || product == entity.Product{} || customer == entity.Customer{}) {
		return entity.Order{}, errors.New("cart or product or customer not found")
	}

	// Check quantity product
	if (product.Quantity < cart.Quantity) {
		return entity.Order{}, errors.New("Product stock is less than the quantity ordered")
	}

	// Update product
	product.Quantity -= cart.Quantity
	u.repo.UpdateProduct(product)

	// Delete cart by product id
	u.repo.DeleteCart(cart)

	// make order
	order := entity.Order{CustomerID: customer_id, ProductID: product_id, Quantity: cart.Quantity, Status: "Processed", OrderDate: time.Now()}
	return u.repo.OrderProduct(order)
}

func (u cartUsecase) UpdateCart(cart entity.Cart) (entity.Cart, error) {
	oldCart, _ := u.repo.FindCartById(cart.CustomerID, cart.ProductID)

	// check if oldCart is not exist
	if (oldCart == entity.Cart{}) {
		if (cart.Quantity == 0) {
			return entity.Cart{}, errors.New("the number of products cannot be zero")
		}
		return u.repo.AddProductToCart(cart)
	}

	// check if cart quantity is zero
	if (cart.Quantity == 0) {
		return u.repo.DeleteCart(cart)
	}

	// update cart
	return u.repo.UpdateCart(cart)
}