package queues

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/davidveg/goapi/modules/internal/config"
	"github.com/davidveg/goapi/modules/internal/entrypoints/dto"
	"github.com/magiconair/properties"
	"log"
)

func SendSQSMessages(sqsMessage *dto.SQSMessageRequest, p *properties.Properties) error {

	sess, err := config.CreateSQSSession(p)
	if err != nil {
		log.Fatalf("Erro ao criar sessão do SQS: %v", err)
		return err
	}

	svc := sqs.New(sess)

	// Envie uma mensagem para a fila SQS
	result, err := svc.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(sqsMessage.Message),
		QueueUrl:    aws.String(p.GetString("sqs.queue.url", "")),
	})
	if err != nil {
		log.Fatalf("Erro ao enviar mensagem para a fila: %v", err)
		return err
	}

	log.Println("Mensagem enviada com sucesso com ID:", *result.MessageId)

	return nil
}
