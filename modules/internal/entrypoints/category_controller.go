package entrypoints

import (
	"encoding/json"
	"github.com/davidveg/goapi/modules/internal/entity"
	"github.com/davidveg/goapi/modules/internal/service"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type CategoryController struct {
	CategoryService *service.CategoryService
}

func CreateCategoryController(categoryService *service.CategoryService) *CategoryController {
	return &CategoryController{CategoryService: categoryService}
}

func (wch *CategoryController) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := wch.CategoryService.GetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err1 := json.NewEncoder(w).Encode(categories)
	if err1 != nil {
		return
	}
}

func (wch *CategoryController) GetCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	category, err := wch.CategoryService.GetCategory(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err1 := json.NewEncoder(w).Encode(category)
	if err1 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (wch *CategoryController) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category entity.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := wch.CategoryService.CreateCategory(category.Name, category.Description)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err1 := json.NewEncoder(w).Encode(result)
	if err1 != nil {
		return
	}
}
