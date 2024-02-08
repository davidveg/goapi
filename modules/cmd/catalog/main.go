package main

import (
	"fmt"
	"github.com/davidveg/goapi/modules/internal/database"
	"github.com/davidveg/goapi/modules/internal/routes"
	"net/http"
)

func main() {
	var r = routes.CreateRoutes()
	fmt.Println("Server is running on port 8080")
	err1 := http.ListenAndServe(":8080", r)
	if err1 != nil {
		fmt.Println("ERROR {}", err1)
		return
	}
	defer database.CloseDBConnection()
}
