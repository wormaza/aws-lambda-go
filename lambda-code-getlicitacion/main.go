package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

//recomendaciones type
//output event: Licitaciones obtenidas
type recomendaciones struct {
	Code         string  `json:"Code"`
	ExternalCode string  `json:"ExternalCode"`
	Relevancia   float32 `json:"Relevancia"`
}

//GetRecomendaciones func
func GetRecomendaciones(rut string) ([]recomendaciones, error) {
	bk, err := getItem(rut)
	if err != nil {
		return nil, err
	}

	return bk, nil
}

//GetLicitacionesByProveedorRequest type
//input event: Rut Proveedor a consultar
type GetLicitacionesByProveedorRequest struct {
	RutProveedor string `json:"RutProveedor"`
}

//HandleEvent func
func HandleEvent(ctx context.Context, event GetLicitacionesByProveedorRequest) ([]recomendaciones, error) {
	lc, _ := lambdacontext.FromContext(ctx)
	log.Print(lc)
	log.Print(lc.AwsRequestID)
	log.Print(lc.InvokedFunctionArn)

	resultado, err := GetRecomendaciones(event.RutProveedor)

	if err != nil {
		fmt.Println("Dynamo call failed:")
		fmt.Println((err.Error()))
		os.Exit(1)
	}

	if len(resultado) == 0 {
		return []recomendaciones{{Code: "-1", ExternalCode: "", Relevancia: -1.0}}, nil
	}

	return resultado, nil
}
func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(HandleEvent)
}
