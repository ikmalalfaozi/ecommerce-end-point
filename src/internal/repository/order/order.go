package order

import (
	"ecommerce/internal/entity"

	"gorm.io/gorm"
)

type OrderRepository interface {
	Order(order entity.Order) (entity.Order, error)
	FindCustomerById(id int64) (entity.Customer, error)
	FindProductById(id int64) (entity.Product, error)
	FindOrderById(orderId int64) (entity.Order, error)
	UpdateProduct(product entity.Product) (entity.Product, error)
	GetOrderByCustomerId(customerId int64) ([]entity.Order, error)
	GetOrderByCustomerIdAndOrderId(customerId int64, orderId int64) (entity.Order, error)
	UpdateOrder(order entity.Order) (entity.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) (orderRepository) {
	// return new order repository
	return orderRepository{db: db}
}

func (r orderRepository) Order(order entity.Order) (entity.Order, error) {
	// insert new order to database
	result := r.db.Create(&order)
	return order, result.Error
}

func (r orderRepository) FindCustomerById(id int64) (entity.Customer, error) {
	// get specific customer by id
	var customer entity.Customer
	result := r.db.Find(&customer, id)
	return customer, result.Error
}

func (r orderRepository) FindProductById(id int64) (entity.Product, error) {
	// get specific product by id
	var product entity.Product
	result := r.db.Find(&product, id)
	return product, result.Error
}

func (r orderRepository) FindOrderById(orderId int64) (entity.Order, error) {
	// find order by order id
	var order entity.Order
	result := r.db.Find(&order, &entity.Order{OrderID: orderId})
	return order, result.Error
}

func (r orderRepository) UpdateProduct(product entity.Product) (entity.Product, error) {
	// update product
	result := r.db.Save(&product);
	return product, result.Error
}

func (r orderRepository) GetOrderByCustomerId(customerId int64) ([]entity.Order, error) {
	// get list of order from a specific customer id
	var orders []entity.Order
	result := r.db.Find(&orders, &entity.Order{CustomerID: customerId})
	return orders, result.Error
}

func (r orderRepository) GetOrderByCustomerIdAndOrderId(customerId int64, orderId int64) (entity.Order, error) {
	// get order by order id and customer id
	var order entity.Order
	result := r.db.Find(&order, &entity.Order{OrderID: orderId, CustomerID: customerId})
	return order, result.Error
}

func (r orderRepository) UpdateOrder(order entity.Order) (entity.Order, error) {
	// update order
	result := r.db.Save(&order);
	return order, result.Error
}