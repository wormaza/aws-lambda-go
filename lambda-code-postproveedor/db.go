package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

//region para AWS
const region = "us-east-2"

//TableName Nombre de la tabla donde se almacenan las recomendaciones
const TableName = "RecomendacionesTest2"

// Declare a new DynamoDB instance. Note that this is safe for concurrent
// use.
var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion(region))

func postItem(item recomendaciones) (bool, error) {

	av, err := dynamodbattribute.MarshalMap(item)

	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Create item
	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(TableName),
	}

	_, err = db.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return true, nil
}
