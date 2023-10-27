package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func mainSQS() {
	// Configura la sesión de AWS
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"), // Cambia a tu región
	})

	if err != nil {
		log.Fatal(err)
	}

	// Crea un nuevo cliente SQS
	svc := sqs.New(sess)

	// URL de la cola que creaste en SQS
	queueURL := "https://sqs.us-east-2.amazonaws.com/103652139467/TQ_Test_SQS"

	// Configura el tiempo de espera para recibir mensajes
	waitTimeSeconds := int64(20)

	// Recibe mensajes de la cola
	result, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:        &queueURL,
		WaitTimeSeconds: &waitTimeSeconds,
	})

	if err != nil {
		log.Fatal(err)
	}

	// Itera sobre los mensajes recibidos
	for _, message := range result.Messages {
		fmt.Printf("Mensaje ID: %s\n", *message.MessageId)
		fmt.Printf("Contenido del mensaje: %s\n", *message.Body)

		// Elimina el mensaje de la cola después de procesarlo
		// _, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
		// 	QueueUrl:      &queueURL,
		// 	ReceiptHandle: message.ReceiptHandle,
		// })

		if err != nil {
			log.Printf("Error al eliminar el mensaje: %v", err)
		}
	}
}
