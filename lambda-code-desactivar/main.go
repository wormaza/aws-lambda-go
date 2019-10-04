package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//region para AWS
const region = "us-east-2"

//TableName Nombre de la tabla donde se almacenan las recomendaciones
const TableName = "RecomendacionesTest2"

// Declare a new DynamoDB instance. Note that this is safe for concurrent
// use.
var db = dynamodb.New(session.New(), aws.NewConfig().WithRegion(region))

//licitacion type
//output event: Licitaciones obtenidas
type licitacion struct {
	Code string `json:"Code"`
}

//licitacionUpdate type
type dataUpdate struct {
	Code string `json:"Code"`
	Rut  string `json:"Rut"`
}

//PutRecomendaciones func
//Ingresa nuevas recomendaciones
func PutRecomendaciones(re dataUpdate) (bool, error) {
	return putItem(re)
}

//PostResponse type
//output event: Response de la operacion
type PostResponse struct {
	Message string `json:"Message"`
}

//HandleEvent func
func HandleEvent(ctx context.Context, event licitacion) (PostResponse, error) {
	lc, _ := lambdacontext.FromContext(ctx)
	log.Print(lc)
	log.Print(lc.AwsRequestID)
	log.Print(lc.InvokedFunctionArn)

	bk, err := getItem(event.Code)
	if err != nil {
		fmt.Println("Dynamo call failed:")
		fmt.Println((err.Error()))
		os.Exit(1)
	}

	var mensaje = "Ok"

	for _, i := range bk {

		resultado, err := PutRecomendaciones(i)

		if err != nil {
			fmt.Println("Dynamo call failed:")
			fmt.Println((err.Error()))
			os.Exit(1)
		}

		if resultado == false {
			mensaje = "No es posible almacenar el registro"
		}
	}

	return PostResponse{Message: mensaje}, nil
}
func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(HandleEvent)
}
