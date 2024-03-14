package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/magiconair/properties"
	"log"
)

func CreateSQSSession() (*session.Session, error) {
	// Carregar o arquivo de propriedades
	p, err := properties.LoadFile("config.properties", properties.UTF8)
	if err != nil {
		log.Fatalf("Erro ao carregar arquivo de propriedades: %v", err)
	}

	// Acessar as propriedades
	credentialsId := p.GetString("sqs.credentials.id", "")
	credentialsSecret := p.GetString("sqs.credentials.secret_key", "")

	// Configuração das credenciais da AWS
	creds := credentials.NewStaticCredentials(credentialsId, credentialsSecret, "")

	// Configuração do cliente do SQS
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(p.GetString("sqs.queue.region", "")), // Defina a região que você deseja usar
		Credentials: creds,
	})
	if err != nil {
		panic(err)
		return nil, err
	}
	return sess, nil
}
