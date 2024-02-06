package entity

import "github.com/google/uuid"

type Product struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Category    Category `json:"category"`
	ImageURL    string   `json:"image_url"`
}

func NewProduct(name, description, imageURL string, price float64, category Category) *Product {
	return &Product{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		Price:       price,
		Category:    category,
		ImageURL:    imageURL,
	}
}
