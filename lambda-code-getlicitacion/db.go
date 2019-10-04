package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

//region para AWS
const region = "us-east-2"

//CodeVariable nombre del filtro para buscar por texto
const CodeVariable = "Code"

//RelevanciaVariable nombre de la variable que indica la relevancia
const RelevanciaVariable = "Relevancia"

//RutVariable indica el rut del proveedor
const RutVariable = "Rut"

//MinRelevancia Valor minimo definido para poder filtrar
const MinRelevancia = 50.0

//TableName Nombre de la tabla donde se almacenan las recomendaciones
const TableName = "RecomendacionesTest2"

//ActivoName Nombre del campo activo
const ActivoName = "Activo"

// Declare a new DynamoDB instance. Note that this is safe for concurrent
// use.
var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion(region))

func getItem(rut string) ([]recomendaciones, error) {

	filterCode := expression.Name(RutVariable).Equal(expression.Value(rut)).And(expression.Name(RelevanciaVariable).GreaterThan(expression.Value(MinRelevancia))).And(expression.Name(ActivoName).Equal(expression.Value(true)))

	proj := expression.NamesList(expression.Name("Code"), expression.Name("ExternalCode"), expression.Name("Relevancia"))

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

	var bk []recomendaciones

	for _, i := range result.Items {
		item := recomendaciones{}

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
