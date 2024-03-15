package main

import (
	"context"
	"github.com/davidveg/goapi/modules/internal/config"
	"github.com/davidveg/goapi/modules/internal/data_providers/connectors"
	"github.com/davidveg/goapi/modules/internal/entrypoints/queues"
	"github.com/davidveg/goapi/modules/internal/routes"
	"log"
	"net/http"
)

func main() {

	// Contexto para controle do listener
	ctx := context.Background()

	// Inicie o listener em uma goroutine
	go queues.ReceiveSQSMessages(ctx, config.CreateProperties())

	// Mantenha a aplicação em execução
	log.Println("Listener iniciado. Aguardando mensagens...")

	// Inicie o WebServer
	var r = routes.CreateRoutes()
	log.Println("Server is running on port 8080")
	err1 := http.ListenAndServe(":8080", r)
	if err1 != nil {
		log.Fatalf("ERROR : %v", err1)
		return
	}

	<-ctx.Done()

	log.Println("Listener encerrado")

	defer connectors.CloseDBConnection()
}
