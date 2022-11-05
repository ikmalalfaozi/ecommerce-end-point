package product

import (
	"ecommerce/internal/entity"
	"ecommerce/internal/repository/product"
)

type ProductUseCase interface {
	AddProduct(product entity.Product) (entity.Product, error)
	GetProducts(page int, limit int) ([]entity.Product, error)
	GetProduct(id int64) (entity.Product, error)
}

type productUsecase struct {
	repo product.ProductRepository
}

func NewUseCase(repo product.ProductRepository) (productUsecase) {
	// return new product usecase
	return productUsecase{repo: repo}
}

func (u productUsecase) AddProduct(product entity.Product) (entity.Product, error) {
	// return added product and error
	return u.repo.AddProduct(product)
}

func (u productUsecase) GetProducts(page int, limit int) ([]entity.Product, error) {
	// return list of product and error
	return u.repo.GetProducts((page-1)*10, limit)
}

func (u productUsecase) GetProduct(id int64) (entity.Product, error) {
	// return specific product and error
	return u.repo.GetProductById(id)
}