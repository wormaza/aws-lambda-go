package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

//Proveedor type
//input event: Datos de Proveedor
type Proveedor struct {
	Rut       string  `json:"Rutproveedor"`
	Relevance float32 `json:"Relevancia"`
}

//PostLicitacionRequest type
//input event: Licitacion a guardar
type PostLicitacionRequest struct {
	IdLicitacion    int  		`json:"IdLicitacion"`
	Codigo 			string 		`json:"CodigoLicitacion"`
	Proveedores		[]Proveedor `json:"Proveedores"`
}

//PostLicitacionResponse type
//output event: Response de la operacion
type PostLicitacionResponse struct {
	Message       string  `json:"Message"`	
}

//HandleEvent func
func HandleEvent(ctx context.Context, event PostLicitacionRequest) ([]PostLicitacionResponse, error) {
	lc, _ := lambdacontext.FromContext(ctx)
	log.Print(lc)
	log.Print(lc.AwsRequestID)
	log.Print(lc.InvokedFunctionArn)

	return []PostLicitacionResponse{
		{Message: "Ok"},
	}, nil
}
func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(HandleEvent)
}
