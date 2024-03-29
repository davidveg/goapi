package entrypoints

import (
	"encoding/json"
	"github.com/davidveg/goapi/modules/internal/entrypoints/dto"
	"github.com/davidveg/goapi/modules/internal/service"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ProductController struct {
	ProductService *service.ProductService
}

func CreateProductController(productService *service.ProductService) *ProductController {
	return &ProductController{ProductService: productService}
}

func (wph *ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := wph.ProductService.GetProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err1 := json.NewEncoder(w).Encode(products)
	if err1 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (wph *ProductController) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	product, err := wph.ProductService.GetProduct(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err1 := json.NewEncoder(w).Encode(product)
	if err1 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (wph *ProductController) GetProductsByCategoryId(w http.ResponseWriter, r *http.Request) {
	categoryID := chi.URLParam(r, "categoryID")
	if categoryID == "" {
		http.Error(w, "categoryID is required", http.StatusBadRequest)
		return
	}
	products, err := wph.ProductService.GetProductsByCategoryID(categoryID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err1 := json.NewEncoder(w).Encode(products)
	if err1 != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (wph *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var productReq dto.ProductRequest
	err := json.NewDecoder(r.Body).Decode(&productReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := wph.ProductService.CreateProduct(productReq.Name, productReq.Description, productReq.CategoryID, productReq.ImageURL, productReq.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err1 := json.NewEncoder(w).Encode(result)
	if err1 != nil {
		return
	}
}
