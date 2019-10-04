package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

//CodeVariable nombre del filtro para buscar por texto
const CodeVariable = "Code"

func getItem(code string) ([]dataUpdate, error) {

	filterCode := expression.Name(CodeVariable).Equal(expression.Value(code))

	proj := expression.NamesList(expression.Name("Code"), expression.Name("Rut"))

	expr, err := expression.NewBuilder().WithFilter(filterCode).WithProjection(proj).Build()

	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(TableName),
	}

	// Make the DynamoDB Query API call
	result, err := db.Scan(params)

	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println((err.Error()))
		os.Exit(1)
	}

	var bk []dataUpdate

	for _, i := range result.Items {
		item := dataUpdate{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		bk = append(bk, item)
	}

	return bk, nil
}
