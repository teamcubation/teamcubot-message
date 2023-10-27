package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	// Crea una nueva sesión de AWS
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"), // Cambia a tu región
	})

	if err != nil {
		log.Fatal(err)
	}

	// Crea un nuevo cliente de DynamoDB
	svc := dynamodb.New(sess)

	// Nombre de la tabla DynamoDB
	tableName := "TQ_Table_Test"

	// Prepara la consulta para escanear todos los items de la tabla
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	// Escanea la tabla
	result, err := svc.Scan(input)
	if err != nil {
		log.Fatal(err)
	}

	// Imprime los items
	for _, item := range result.Items {
		fmt.Println(item)
	}
}
