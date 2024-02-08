package service

import (
	"github.com/davidveg/goapi/modules/internal/database/repositories"
	"github.com/davidveg/goapi/modules/internal/entity"
)

type CategoryService struct {
	CategoryDB repositories.CategoryRepository
}

func NewCategoryService(categoryDB repositories.CategoryRepository) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (cs *CategoryService) GetCategories() ([]*entity.Category, error) {
	categories, err := cs.CategoryDB.GetCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (cs *CategoryService) CreateCategory(name, description string) (*entity.Category, error) {
	category := entity.NewCategory(name, description)
	_, err := cs.CategoryDB.CreateCategory(category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (cs *CategoryService) GetCategory(id string) (*entity.Category, error) {
	category, err := cs.CategoryDB.GetCategory(id)
	if err != nil {
		return nil, err
	}
	return category, nil
}
