package main

import (
	"database/sql"
	"fmt"
	"github.com/davidveg/goapi/app/internal/database"
	"github.com/davidveg/goapi/app/internal/service"
	"github.com/davidveg/goapi/app/internal/webserver"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/imersao17")
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

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

	fmt.Println("Server is running on port 8080")
	err1 := http.ListenAndServe(":8080", routes)
	if err1 != nil {
		fmt.Println("ERROR {}", err1)
		return
	}

}
