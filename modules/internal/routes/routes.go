package routes

import (
	"github.com/davidveg/goapi/modules/internal/database"
	"github.com/davidveg/goapi/modules/internal/service"
	"github.com/davidveg/goapi/modules/internal/webserver"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var db = database.GetConnection()

func CreateRoutes() *chi.Mux {
	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)

	productDB := database.NewProductDB(db)
	productService := service.NewProductService(*productDB, *categoryDB)

	webCategoryHandler := webserver.NewWebCategoryHandler(categoryService)
	webProductHandler := webserver.NewWebProductHandler(productService)

	routes := chi.NewRouter()
	routes.Use(middleware.Logger)
	routes.Use(middleware.Recoverer)
	routes.Get("/category/{id}", webCategoryHandler.GetCategory)
	routes.Get("/category", webCategoryHandler.GetCategories)
	routes.Post("/category", webCategoryHandler.CreateCategory)

	routes.Get("/product/{id}", webProductHandler.GetProduct)
	routes.Get("/product", webProductHandler.GetProducts)
	routes.Get("/product/category/{categoryID}", webProductHandler.GetProductsByCategoryId)
	routes.Post("/product", webProductHandler.CreateProduct)

	return routes
}
