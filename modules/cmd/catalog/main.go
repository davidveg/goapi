package main

import (
	"context"
	"fmt"
	"github.com/davidveg/goapi/modules/internal/data_providers/connectors"
	"github.com/davidveg/goapi/modules/internal/entrypoints/queues"
	"github.com/davidveg/goapi/modules/internal/routes"
	"net/http"
)

func main() {

	// Contexto para controle do listener
	ctx := context.Background()

	// Inicie o listener em uma goroutine
	go queues.ReceiveSQSMessages(ctx)

	// Mantenha a aplicação em execução
	fmt.Println("Listener iniciado. Aguardando mensagens...")

	var r = routes.CreateRoutes()
	fmt.Println("Server is running on port 8080")
	err1 := http.ListenAndServe(":8080", r)
	if err1 != nil {
		fmt.Println("ERROR {}", err1)
		return
	}

	<-ctx.Done()
	fmt.Println("Listener encerrado")

	defer connectors.CloseDBConnection()
}
