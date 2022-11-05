package customer

import (
	"ecommerce/internal/entity"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetCustomers(offset int, limit int) ([]entity.Customer, error)
	GetCustomerById(id int64) (entity.Customer, error)
	AddCustomer(customer entity.Customer) (entity.Customer, error)
}

type customerRepository struct {
	db *gorm.DB
}


func NewRepo(db *gorm.DB) customerRepository {
	// return new customer repository
	return customerRepository{db: db}
}

func (r customerRepository) AddCustomer(customer entity.Customer) (entity.Customer, error) {
	// insert new customer to database
	result := r.db.Create(&customer)
	return customer, result.Error
}

func (r customerRepository) GetCustomers(offset int, limit int) ([]entity.Customer, error) {
	// get list of customers with offset and limit
	var customers []entity.Customer 
	result := r.db.Offset(offset).Limit(limit).Find(&customers)
	return customers, result.Error
}

func (r customerRepository) GetCustomerById(id int64) (entity.Customer, error) {
	// get specific customer by id
	var customer entity.Customer
	result := r.db.Find(&customer, id)
	return customer, result.Error
}