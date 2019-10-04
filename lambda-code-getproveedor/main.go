package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

//GetProveedorByLicitacionRequest type
//input event: Licitacion a consultar
type GetProveedorByLicitacionRequest struct {
	Code string `json:"code"`
}

//recomendaciones type
//output event: Listado de Proveedores asociadas a Licitacion
type recomendaciones struct {
	Rut        string  `json:"Rut"`
	Relevancia float32 `json:"Relevancia"`
}

//GetRecomendaciones func
func GetRecomendaciones(rut string) ([]recomendaciones, error) {
	bk, err := getItem(rut)
	if err != nil {
		return nil, err
	}

	return bk, nil
}

//HandleEvent func
func HandleEvent(ctx context.Context, event GetProveedorByLicitacionRequest) ([]recomendaciones, error) {
	lc, _ := lambdacontext.FromContext(ctx)
	log.Print(lc)
	log.Print(lc.AwsRequestID)
	log.Print(lc.InvokedFunctionArn)

	resultado, err := GetRecomendaciones(event.Code)

	if err != nil {
		fmt.Println("Dynamo call failed:")
		fmt.Println((err.Error()))
		os.Exit(1)
	}

	if len(resultado) == 0 {
		return []recomendaciones{{Rut: "-1", Relevancia: -1.0}}, nil
	}

	return resultado, nil
}
func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(HandleEvent)
}
