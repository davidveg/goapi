package entities

import "github.com/google/uuid"

type Category struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewCategory(name, description string) *Category {
	return &Category{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
	}
}
