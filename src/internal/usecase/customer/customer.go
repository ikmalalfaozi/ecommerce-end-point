package customer

import (
	"ecommerce/internal/entity"
	"ecommerce/internal/repository/customer"

	"golang.org/x/crypto/bcrypt"
)

type CustomerUseCase interface {
	AddCustomer(customer entity.Customer) (entity.Customer, error)
	GetCustomers(page int, limit int) ([]entity.Customer, error)
	GetCustomer(id int64) (entity.Customer, error)
}

type customerUsecase struct {
	repo customer.CustomerRepository
}

func NewUseCase(repo customer.CustomerRepository) (customerUsecase) {
	// return new customer usecase
	return customerUsecase{repo: repo}
}

func (u customerUsecase) AddCustomer(customer entity.Customer) (entity.Customer, error) {
	// hash password with bcrypt
	hashPasword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
	// check if hash password not succeed
	if (err != nil) {
		return entity.Customer{}, err
	}

	// convert hash password to string
	customer.Password = string(hashPasword);
	// return added customer and error
	return u.repo.AddCustomer(customer)
}

func (u customerUsecase) GetCustomers(page int, limit int) ([]entity.Customer, error) {
	// return list of customer and error
	return u.repo.GetCustomers((page-1)*10, limit)
}

func (u customerUsecase) GetCustomer(id int64) (entity.Customer, error) {
	// return specific customer and error
	return u.repo.GetCustomerById(id)
}