package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func putItem(item dataUpdate) (bool, error) {

	liccode := item.Code
	rut := item.Rut
	activo := false

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				BOOL: aws.Bool(activo),
			},
		},
		TableName: aws.String(TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Code": {
				S: aws.String(liccode),
			},
			"Rut": {
				S: aws.String(rut),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Activo = :r"),
	}

	_, err := db.UpdateItem(input)
	if err != nil {
		fmt.Println("Got error calling UpdateItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return true, nil
}
