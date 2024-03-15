package queues

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/davidveg/goapi/modules/internal/config"
	"github.com/magiconair/properties"
	"log"
)

func ReceiveSQSMessages(ctx context.Context, p *properties.Properties) {
	sess, err := config.CreateSQSSession(p)
	if err != nil {
		log.Fatalf("Erro ao criar sessão do SQS: %v", err)
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("Encerrando a função de recebimento de mensagens...")
			return
		default:

			svc := sqs.New(sess)

			// Receba mensagens da fila SQS
			resultReceive, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
				QueueUrl:            aws.String(p.GetString("sqs.queue.url", "")),
				MaxNumberOfMessages: aws.Int64(p.GetInt64("sqs.queue.max_messages", 1)), // Defina o número máximo de mensagens a serem recebidas
			})
			if err != nil {
				log.Fatalf("Erro ao receber mensagem da fila: %v", err)
				return
			}

			if len(resultReceive.Messages) > 0 {
				log.Println("Mensagens recebidas:")
				for _, message := range resultReceive.Messages {
					log.Println("Corpo da mensagem:", *message.Body)

					// Excluir mensagem da fila SQS após processamento
					_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
						QueueUrl:      aws.String(p.GetString("sqs.queue.url", "")),
						ReceiptHandle: message.ReceiptHandle,
					})
					if err != nil {
						log.Fatalf("Erro ao excluir mensagem da fila: %v", err)
						return
					}
					log.Println("Mensagem excluída com sucesso")
				}
			}
		}
	}
}
