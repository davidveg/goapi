package routes

import (
	"github.com/davidveg/goapi/modules/internal/database"
	"github.com/davidveg/goapi/modules/internal/database/connectors"
	"github.com/davidveg/goapi/modules/internal/entrypoints"
	"github.com/davidveg/goapi/modules/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var db = connectors.GetDBConnection()

func CreateRoutes() *chi.Mux {
	CategoryController, ProductController := CreateControllers()

	routes := chi.NewRouter()
	routes.Use(middleware.Logger)
	routes.Use(middleware.Recoverer)
	routes.Get("/category/{id}", CategoryController.GetCategory)
	routes.Get("/category", CategoryController.GetCategories)
	routes.Post("/category", CategoryController.CreateCategory)

	routes.Get("/product/{id}", ProductController.GetProduct)
	routes.Get("/product", ProductController.GetProducts)
	routes.Get("/product/category/{categoryID}", ProductController.GetProductsByCategoryId)
	routes.Post("/product", ProductController.CreateProduct)

	return routes
}

func CreateControllers() (*entrypoints.CategoryController, *entrypoints.ProductController) {
	categoryRepository := database.CreateCategoryRepository(db)
	categoryService := service.NewCategoryService(*categoryRepository)

	productRepository := database.CreateProductRepository(db)
	productService := service.NewProductService(*productRepository, *categoryRepository)

	categoryController := entrypoints.CreateCategoryController(categoryService)
	productController := entrypoints.CreateProductController(productService)
	return categoryController, productController
}
