package repositories

import (
	"database/sql"
	"github.com/davidveg/goapi/modules/internal/data_providers/connectors"
	"github.com/davidveg/goapi/modules/internal/entities"
)

type ProductRepository struct {
	db *sql.DB
}

func CreateProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (pd *ProductRepository) GetProducts() ([]*entities.Product, error) {
	rows, err := pd.db.Query("SELECT id, name,  description, price, category_id, image_url FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entities.Product
	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Category.ID, &product.ImageURL); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (pd *ProductRepository) GetProduct(id string) (*entities.Product, error) {
	var product entities.Product
	err := pd.db.QueryRow("SELECT id, name, description, price, category_id, image_url FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Category.ID, &product.ImageURL)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (pd *ProductRepository) GetProductByCategoryID(categoryID string) ([]*entities.Product, error) {
	rows, err := pd.db.Query("SELECT id, name, description, price, category_id, image_url from products where category_id = ?", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*entities.Product
	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Category.ID, &product.ImageURL); err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (pd *ProductRepository) CreateProduct(product *entities.Product) (*entities.Product, error) {
	_, err := pd.db.Exec("INSERT INTO products(id, name, description, price, category_id, image_url) VALUES (?,?,?,?,?,?)", product.ID, product.Name, product.Description, product.Price, product.Category.ID, product.ImageURL)
	if err != nil {
		return nil, err
	}
	defer connectors.CloseDBConnection()
	return product, nil
}
