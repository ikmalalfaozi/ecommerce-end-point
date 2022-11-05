package order

import (
	"ecommerce/internal/entity"
	"ecommerce/internal/repository/order"
	"errors"
	"strings"
	"time"
)

type OrderUseCase interface {
	Order(order entity.Order) (entity.Order, error)
	GetOrders(customerId int64) ([]entity.Order, error)
	GetOrder(customerId int64, orderId int64) (entity.Order, error)
	ChangeOrderStatus(orderId int64, status string) (entity.Order, error)
}

type orderUsecase struct {
	repo order.OrderRepository
}

func NewUseCase(repo order.OrderRepository) (orderUsecase) {
	// return new order usecase
	return orderUsecase{repo: repo}
}

func (u orderUsecase) Order(order entity.Order) (entity.Order, error) {
	customer, _ := u.repo.FindCustomerById(order.CustomerID)
	product, _ := u.repo.FindProductById(order.ProductID)
	
	// check is empty customer or product
	if (customer == entity.Customer{} || product == entity.Product{}) {
		return entity.Order{}, errors.New("customer id or product id not found")
	}
	
	// check quantity product
	if (product.Quantity < order.Quantity) {
		return entity.Order{}, errors.New("Product stock is less than the quantity ordered")
	}
	
	// update product
	product.Quantity -= order.Quantity
	u.repo.UpdateProduct(product)
	
	// make order
	order.Status = "Processed"
	order.OrderDate = time.Now()
	// return new order
	return u.repo.Order(order)
}

func (u orderUsecase) GetOrders(customerId int64) ([]entity.Order, error) {
		// check if customer id = 0
		if (customerId == 0) {
			return []entity.Order{}, nil
		}
		return u.repo.GetOrderByCustomerId(customerId)
}

func (u orderUsecase) GetOrder(customerId int64, orderId int64) (entity.Order, error) {
	// check if customer id = 0 or orderId = 0
	if (customerId == 0 || orderId == 0) {
		return entity.Order{}, errors.New("order not found")
	}

	order, err := u.repo.GetOrderByCustomerIdAndOrderId(customerId, orderId)
	// check if err
	if err != nil {
		return entity.Order{}, err
	}

	// check if order is not found
	if (order == entity.Order{}) {
		return order, errors.New("order not found")
	} 
	
	return order, nil
}

func (u orderUsecase) ChangeOrderStatus(orderId int64, status string) (entity.Order, error) {
	// check if orderId = 0
	if (orderId == 0) {
		return entity.Order{}, errors.New("order not found")
	} 

	order, err := u.repo.FindOrderById(orderId)
	// check if error
	if (err != nil) {
		return entity.Order{}, err
	}
	// check if order is not found
	if (order == entity.Order{}) {
		return entity.Order{}, errors.New("order not found")
	}
	
	status = strings.Title(status)
	// check status [processed, shipped, sent, cancelled]
	switch status {
	case "Shipped":
		// shipped status is valid if the previous status was sent
		if  order.Status == "Sent" {
			order.Status = "Shipped"
			order.ShippedDate = time.Now()
			return u.repo.UpdateOrder(order)
		}
		return entity.Order{}, errors.New("Can't change order status to Shipped")
	case "Sent":
		// sent status is valid if the previous status was processed
		if (order.Status == "Processed") {
			order.Status = "Sent"
			return u.repo.UpdateOrder(order)
		}
		return entity.Order{}, errors.New("Can't change order status to Sent")
	case "Cancelled":
		// cancelled status is valid if the previous status was processed
		if (order.Status == "Processed") {
			order.Status = "Cancelled"
			product, _ := u.repo.FindProductById(order.ProductID)
			product.Quantity += order.Quantity
			u.repo.UpdateProduct(product)
			return u.repo.UpdateOrder(order)
		}
		return entity.Order{}, errors.New("Can't change order status to Cancelled")
	case "Processed":
		// can't change order status to processed
		return entity.Order{}, errors.New("Can't change order status to Processed")
	default:
		// order status is invalid
		return entity.Order{}, errors.New("Invalid order status")
	}
}