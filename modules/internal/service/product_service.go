package service

import (
	"github.com/davidveg/goapi/modules/internal/dataproviders/repositories"
	"github.com/davidveg/goapi/modules/internal/entity"
)

type ProductService struct {
	ProductDB  repositories.ProductRepository
	CategoryDB repositories.CategoryRepository
}

func NewProductService(productDB repositories.ProductRepository, categoryDB repositories.CategoryRepository) *ProductService {
	return &ProductService{
		ProductDB:  productDB,
		CategoryDB: categoryDB,
	}
}

func (ps *ProductService) GetProducts() ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) GetProduct(id string) (*entity.Product, error) {
	product, err := ps.ProductDB.GetProduct(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (ps *ProductService) GetProductsByCategoryID(categoryID string) ([]*entity.Product, error) {
	products, err := ps.ProductDB.GetProductByCategoryID(categoryID)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ps *ProductService) CreateProduct(name, description, categoryId, imageUrl string, price float64) (*entity.Product, error) {
	category, _ := ps.CategoryDB.GetCategory(categoryId)
	product := entity.NewProduct(name, description, imageUrl, price, *category)
	p, err := ps.ProductDB.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return p, err
}
