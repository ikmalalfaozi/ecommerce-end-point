package product

import (
	"ecommerce/internal/entity"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProducts(offset int, limit int) ([]entity.Product, error)
	GetProductById(id int64) (entity.Product, error)
	AddProduct(product entity.Product) (entity.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) productRepository {
	// return new product repository
	return productRepository{db: db}
}

func (r productRepository) AddProduct(product entity.Product) (entity.Product, error) {
	// insert product to database
	result := r.db.Create(&product) 
	return product, result.Error
}

func (r productRepository) GetProducts(offset int, limit int) ([]entity.Product, error) {
	// get list of products with offset and limit
	var products []entity.Product 
	result := r.db.Offset(offset).Limit(limit).Find(&products)
	return products, result.Error
}

func (r productRepository) GetProductById(id int64) (entity.Product, error) {
	// get specific product by id in database
	var product entity.Product
	result := r.db.Find(&product, id)
	return product, result.Error
}