package database

import (
	"database/sql"
	"github.com/davidveg/goapi/app/internal/entity"
)

type CategoryDB struct {
	db *sql.DB
}

func NewCategoryDB(db *sql.DB) *CategoryDB {
	return &CategoryDB{db: db}
}

func (cd *CategoryDB) GetCategories() ([]*entity.Category, error) {
	rows, err := cd.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []*entity.Category
	for rows.Next() {
		var category entity.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}

func (cd *CategoryDB) GetCategory(id string) (*entity.Category, error) {
	var category entity.Category
	err := cd.db.QueryRow("SELECT id, name, description FROM categories WHERE id = ?", id).Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (cd *CategoryDB) CreateCategory(category *entity.Category) (string, error) {
	_, err := cd.db.Exec("INSERT INTO categories(id, name, description) VALUES (?,?,?)", category.ID, category.Name, category.Description)
	if err != nil {
		return "", err
	}
	return category.ID, nil
}
